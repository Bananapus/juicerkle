package main

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"juicerkle/tree"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	_ "github.com/mattn/go-sqlite3"
)

// An EVM chain ID
type chainId int

// An item in config.json
type ConfigItem struct {
	Name    string  `json:"name"`
	ChainId chainId `json:"chainId"`
	RpcUrl  string  `json:"rpcUrl"`
}

// A BPLeaf as stored in sqlite
type Leaf struct {
	ChainId         int    `db:"chain_id"`
	ContractAddress string `db:"contract_address"`
	TokenAddress    string `db:"token_address"`
	LeafIndex       uint   `db:"leaf_index"`
	LeafHash        string `db:"leaf_hash"`
}

// A specification for an inbox tree on a sucker contract
type InboxTree struct {
	ChainId       chainId
	SuckerAddress common.Address
	TokenAddress  common.Address
	Root          [32]byte
}

// Schema for incoming proof requests
type ProofRequest struct {
	ChainId chainId        `json:"chainId"` // The chain ID of the sucker contract
	Sucker  common.Address `json:"sucker"`  // The sucker contract address
	Token   common.Address `json:"token"`   // The address of the token being claimed
	Index   uint           `json:"index"`   // The index of the leaf to prove on the sucker contract
}

// Map of chainId to ethclient.Client
var clients = make(map[chainId]*ethclient.Client)

func main() {
	// Read config
	var config []ConfigItem
	configFile, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatalf("Could not read config.json: %v\n", err)
	}

	if err := json.Unmarshal(configFile, &config); err != nil {
		log.Fatalf("Failed to unmarshal config.json: %v\n", err)
	}

	// Set up ETH clients
	for _, network := range config {
		client, err := ethclient.Dial(network.RpcUrl)
		if err != nil {
			log.Fatalf("Failed to connect to the %s network: %v", network.Name, err)
		}
		clients[network.ChainId] = client
	}

	// Set up DB
	if err := initDb(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("POST /proof", proof)
	log.Printf("Listening on http://localhost:%s\n", port)
	http.ListenAndServe(":"+port, nil)
}

func proof(w http.ResponseWriter, req *http.Request) {
	var toProve ProofRequest
	err := json.NewDecoder(req.Body).Decode(&toProve)
	if err != nil {
		log.Printf("Failed to parse request body: %v\n", err)
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	if toProve.Index > 1<<32 {
		http.Error(w, "Invalid leaf index (too large)", http.StatusBadRequest)
		return
	}

	// Set up cancellation context
	ctx, cancel := context.WithTimeout(req.Context(), 15*time.Second)
	defer cancel()

	// Calculate the proof off of the main thread
	proofCh := make(chan [][]byte)
	go func() {
		client, ok := clients[toProve.ChainId]
		if !ok {
			log.Printf("Chain %d not supported\n", toProve.ChainId)
			http.Error(w, "Chain not supported", http.StatusBadRequest)
			return
		}

		localSucker, err := NewBPSucker(toProve.Sucker, client)
		if err != nil {
			log.Printf("Failed to instantiate the sucker contract: %v\n", err)
			http.Error(w, "Failed to instantiate the sucker contract", http.StatusInternalServerError)
			return
		}

		inboxTree, err := localSucker.Inbox(&bind.CallOpts{Context: ctx}, toProve.Token)
		if err != nil {
			log.Printf("Failed to get inbox tree '%d:%s:%s' root: %v\n",
				toProve.ChainId, toProve.Sucker.String(), toProve.Token.String(), err)
			http.Error(w, "Failed to get inbox tree", http.StatusInternalServerError)
			return
		}

		// Get the proof
		proof, err := dbProof(ctx, toProve.Index, InboxTree{
			ChainId:       toProve.ChainId,
			SuckerAddress: toProve.Sucker,
			TokenAddress:  toProve.Token,
			Root:          inboxTree.Root,
		})
		if err != nil {
			log.Printf("Failed to get proof: %v\n", err)
			http.Error(w, "Failed to get proof", http.StatusInternalServerError)
			return
		}

		proofCh <- proof
	}()

	// Wait for the proof or a cancellation
	select {
	case <-ctx.Done():
		log.Printf("Request cancelled: %v\n", ctx.Err())
		http.Error(w, "Request cancelled", http.StatusRequestTimeout)
		return
	case p := <-proofCh:
		b, err := json.Marshal(p)
		if err != nil {
			log.Printf("Failed to marshal proof: %v\n", err)
			http.Error(w, "Failed to marshal proof", http.StatusInternalServerError)
			return
		}

		w.Write(b)
	}
}

// Update the leaves in the database for a specific sucker
func updateLeaves(ctx context.Context, inboxTree InboxTree) error {
	client, ok := clients[inboxTree.ChainId]
	if !ok {
		return fmt.Errorf("chain %d not supported", inboxTree.ChainId)
	}

	sucker, err := NewBPSucker(inboxTree.SuckerAddress, client)
	if err != nil {
		return fmt.Errorf("failed to instantiate the sucker contract: %v", err)
	}

	// Get the latest leaf hash for the tree from the db
	var latestHash string
	err = db.QueryRowContext(ctx, `SELECT leaf_hash FROM leaves
		WHERE chain_id = ? AND contract_address = ? AND token_address = ?
		ORDER BY leaf_index DESC LIMIT 1`,
		inboxTree.ChainId, inboxTree.SuckerAddress.String(), inboxTree.TokenAddress.String()).Scan(&latestHash)
	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("failed to query database: %v", err)
	}

	seenLatestDbLeaf := false
	var latestHashBytes []byte
	if err == sql.ErrNoRows {
		// Start from the beginning if there are no leaves in the db
		seenLatestDbLeaf = true
	} else {
		if latestHashBytes, err = hex.DecodeString(latestHash); err != nil {
			return fmt.Errorf("failed to decode leaf hash '%s': %v", latestHash, err)
		}
	}

	// Get peer chain ID and address
	peerSuckerChainId, err := sucker.PeerChainID(&bind.CallOpts{Context: ctx})
	if err != nil {
		return fmt.Errorf("failed to get peer chain ID: %v", err)
	}
	peerChainId := chainId(peerSuckerChainId.Uint64())

	peerSuckerAddr, err := sucker.PEER(&bind.CallOpts{Context: ctx})
	if err != nil {
		return fmt.Errorf("failed to get peer: %v", err)
	}

	peerClient, ok := clients[peerChainId]
	if !ok {
		return fmt.Errorf("peer chain %d not supported", peerChainId)
	}

	// Instantiate peer sucker and iterate through insertions to its outbox
	peerSucker, err := NewBPSucker(peerSuckerAddr, peerClient)
	if err != nil {
		return fmt.Errorf("failed to instantiate the peer sucker contract '%s' on chain %d: %v", peerSuckerAddr.String(), peerChainId, err)
	}

	outboxIterator, err := peerSucker.FilterInsertToOutboxTree(&bind.FilterOpts{Context: ctx}, nil, []common.Address{inboxTree.TokenAddress})
	if err != nil {
		return fmt.Errorf("failed to instantiate outbox iterator: %v", err)
	}
	defer outboxIterator.Close()

	leavesToInsert := make([]Leaf, 0)
	for outboxIterator.Next() {
		// Keep iterating until we pass the latest hash
		if !seenLatestDbLeaf {
			if bytes.Equal(outboxIterator.Event.Hashed[:], latestHashBytes) {
				seenLatestDbLeaf = true
			}
			continue
		}

		// Add the remaining leaves to the list to insert
		leavesToInsert = append(leavesToInsert, Leaf{
			ChainId:         int(inboxTree.ChainId),
			ContractAddress: inboxTree.SuckerAddress.String(),
			LeafIndex:       uint(outboxIterator.Event.Index.Uint64()),
			TokenAddress:    inboxTree.TokenAddress.String(),
			LeafHash:        hex.EncodeToString(outboxIterator.Event.Hashed[:]),
		})

		// If we've gotten to the latest root, break
		if inboxTree.Root == outboxIterator.Event.Root {
			break
		}
	}

	if err := outboxIterator.Error(); err != nil {
		return fmt.Errorf("failed while iterating through outbox insertions: %v", err)
	}

	// Start sqlite transaction to insert leaves
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to start sqlite transaction: %v", err)
	}

	stmt, err := tx.PrepareContext(ctx, `INSERT INTO leaves (chain_id, contract_address, token_address, leaf_index, leaf_hash)
		VALUES (?, ?, ?, ?, ?)`)
	if err != nil {
		return fmt.Errorf("failed to prepare sqlite statement: %v", err)
	}
	defer stmt.Close()

	for _, leaf := range leavesToInsert {
		_, err := stmt.ExecContext(ctx, leaf.ChainId, leaf.ContractAddress, leaf.TokenAddress, leaf.LeafIndex, leaf.LeafHash)
		if err != nil {
			tx.Rollback() // Rollback the transaction if an error occurs
			return fmt.Errorf("failed to insert leaf into sqlite: %v", err)
		}
	}

	finalCount := leavesToInsert[len(leavesToInsert)-1].LeafIndex

	tx.ExecContext(ctx, `INSERT OR REPLACE INTO trees (chain_id, contract_address, token_address, current_root, count)
		VALUES (?, ?, ?, ?, ?)`, inboxTree.ChainId, inboxTree.SuckerAddress.String(), inboxTree.TokenAddress.String(),
		hex.EncodeToString(inboxTree.Root[:]), finalCount)

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit sqlite transaction: %v", err)
	}

	return nil
}

// Get the proof for a leaf from the database
func dbProof(ctx context.Context, index uint, inboxTree InboxTree) ([][]byte, error) {
	var dbRoot string
	var err error

	if err = db.QueryRowContext(ctx, `SELECT current_root FROM trees
		WHERE chain_id = ? AND contract_address = ? AND token_address = ?`,
		inboxTree.ChainId, inboxTree.SuckerAddress.String(), inboxTree.TokenAddress.String()).
		Scan(&dbRoot); err != nil && err != sql.ErrNoRows {
		return nil, fmt.Errorf("failed to query database")
	}

	dbRootBytes, err := hex.DecodeString(dbRoot)
	if err != nil {
		return nil, fmt.Errorf("failed to decode dbRoot '%s': %v", dbRoot, err)
	}

	// If no rows were found, or the counts/roots don't match, we need to update the database.
	if err == sql.ErrNoRows || !bytes.Equal(dbRootBytes, inboxTree.Root[:]) {
		if err := updateLeaves(ctx, inboxTree); err != nil {
			return nil, fmt.Errorf("failed to update leaves: %v", err)
		}
	}

	rows, err := db.QueryContext(ctx, `SELECT leaf_hash FROM leaves
		WHERE chain_id = ? AND contract_address = ? AND token_address = ?`,
		inboxTree.ChainId, inboxTree.SuckerAddress.String(), inboxTree.TokenAddress.String())
	if err != nil {
		// Note: this includes sql.ErrNoRows
		return nil, fmt.Errorf("failed to get '%d:%s:%s' leaves from the database: %v",
			inboxTree.ChainId, inboxTree.SuckerAddress.String(), inboxTree.TokenAddress.String(), err)
	}

	leaves := make([][]byte, 0)
	for rows.Next() {
		var leafHash string
		err := rows.Scan(&leafHash)
		if err != nil {
			return nil, fmt.Errorf("failed to scan leaf hash %s: %v", leafHash, err)
		}

		b, err := hex.DecodeString(leafHash)
		if err != nil {
			return nil, fmt.Errorf("failed to decode leaf hash %s: %v", leafHash, err)
		}

		leaves = append(leaves, b)
	}

	// Construct a tree with the retrieved leaves
	t := tree.NewTree(leaves)
	treeRoot := t.Root()

	// Sanity check the tree root
	if !bytes.Equal(treeRoot, inboxTree.Root[:]) {
		return nil, fmt.Errorf("calculated tree root '%s' does not match expected tree root '%s'",
			hex.EncodeToString(treeRoot), hex.EncodeToString(inboxTree.Root[:]))
	}

	// Bounds are checked in the proof function
	p, err := t.Proof(index)
	if err != nil {
		return nil, fmt.Errorf("failed to get proof: %v", err)
	}

	return p, nil
}
