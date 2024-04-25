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
	"math/big"
	"net/http"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	_ "github.com/mattn/go-sqlite3"
)

// An item in config.json
type NetworkConfig struct {
	Name    string   `json:"name"`
	ChainId *big.Int `json:"chainId"`
	RpcUrl  string   `json:"rpcUrl"`
}

// A BPLeaf as stored in sqlite
type Leaf struct {
	ChainId         string `db:"chain_id"`
	ContractAddress string `db:"contract_address"`
	TokenAddress    string `db:"token_address"`
	LeafIndex       string `db:"leaf_index"`
	LeafHash        string `db:"leaf_hash"`
}

// A specification for an inbox tree on a sucker contract
type InboxTree struct {
	ChainId       *big.Int
	SuckerAddress common.Address
	TokenAddress  common.Address
	Root          [32]byte
}

// Schema for incoming proof requests
type ProofRequest struct {
	ChainId *big.Int       `json:"chainId"` // The chain ID of the sucker contract
	Sucker  common.Address `json:"sucker"`  // The sucker contract address
	Token   common.Address `json:"token"`   // The address of the token being claimed
	Index   *big.Int       `json:"index"`   // The index of the leaf to prove on the sucker contract
}

// Map of chain IDs to ETH clients
var clients = make(map[*big.Int]*ethclient.Client)

func main() {
	// Read networks
	var networks []NetworkConfig
	configFile, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatalf("Could not read config.json: %v\n", err)
	}
	if err := json.Unmarshal(configFile, &networks); err != nil {
		log.Fatalf("Failed to unmarshal config.json: %v\n", err)
	}

	// Set up ETH clients
	for _, network := range networks {
		client, err := ethclient.Dial(network.RpcUrl)
		if err != nil {
			log.Fatalf("Failed to connect to the %s network: %v\n", network.Name, err)
		}
		clients[network.ChainId] = client
	}

	// Set up DB
	if err := initDb(); err != nil {
		log.Fatalf("Failed to initialize database: %v\n", err)
	}
	defer db.Close()

	// Default to 8080 if no port is specified
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("POST /proof", proof)
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Juicerkle running. Send requests to /proof."))
	})
	log.Printf("Listening on http://localhost:%s\n", port)
	http.ListenAndServe(":"+port, nil)
}

var (
	MaxIndex = big.NewInt(1 << 32)
	ZeroRoot = make([]byte, 32)
)

func proof(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	var toProve ProofRequest
	err := json.NewDecoder(req.Body).Decode(&toProve)
	if err != nil {
		errStr := fmt.Sprintf("Failed to parse proof request body: %v", err)
		log.Println(errStr)
		http.Error(w, errStr, http.StatusBadRequest)
		return
	}

	if toProve.Index.Cmp(MaxIndex) > 0 {
		errStr := fmt.Sprintf("Invalid leaf index: %d (too large)", toProve.Index)
		log.Println(errStr)
		http.Error(w, errStr, http.StatusBadRequest)
		return
	}

	// Set up context
	// ctx, cancel := context.WithTimeout(req.Context(), 15*time.Second)
	ctx, cancel := context.WithCancel(req.Context()) // Use cancel instead of timeout for debugging
	defer cancel()

	client, ok := clients[toProve.ChainId]
	if !ok {
		errStr := fmt.Sprintf("No RPC for chain ID %s. Contact the developer to have it added.", toProve.ChainId)
		log.Println(errStr)
		http.Error(w, errStr, http.StatusNotFound)
		return
	}

	localSucker, err := NewBPSucker(toProve.Sucker, client)
	if err != nil {
		errStr := fmt.Sprintf("Failed to instantiate a BPSucker at address %s on network #%s: %v",
			toProve.Sucker, toProve.ChainId, err)
		log.Println(errStr)
		http.Error(w, errStr, http.StatusInternalServerError)
		return
	}

	// Default callOpts
	callOpts := &bind.CallOpts{Context: ctx}

	localInboxTree, err := localSucker.Inbox(callOpts, toProve.Token)
	if err != nil {
		errStr := fmt.Sprintf("Failed to get inbox tree root for token %s, sucker %s, and chain %d: %v",
			toProve.Token, toProve.Sucker, toProve.ChainId, err)
		log.Println(errStr)
		http.Error(w, errStr, http.StatusInternalServerError)
		return
	}

	// Check if the inbox tree root is empty.
	if bytes.Equal(localInboxTree.Root[:], ZeroRoot) {
		errStr := fmt.Sprintf("Inbox tree root for token %s, sucker %s, and chain %d is empty",
			toProve.Token, toProve.Sucker, toProve.ChainId)
		log.Println(errStr)
		http.Error(w, errStr, http.StatusInternalServerError)
		return
	}

	// Get the proof
	proof, err := dbProof(ctx, toProve.Index, InboxTree{
		ChainId:       toProve.ChainId,
		SuckerAddress: toProve.Sucker,
		TokenAddress:  toProve.Token,
		Root:          localInboxTree.Root,
	})
	if err != nil {
		log.Printf("Failed to get proof: %v\n", err)
		http.Error(w, "Failed to get proof", http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(proof)
	if err != nil {
		log.Printf("Failed to marshal proof: %v\n", err)
		http.Error(w, "Failed to marshal proof", http.StatusInternalServerError)
		return
	}

	w.Write(b)
}

// Get the proof for a leaf from the database. If the database is out of date, update it.
func dbProof(ctx context.Context, index *big.Int, inboxTree InboxTree) ([][]byte, error) {
	var dbRoot string
	var err error

	logDescription := fmt.Sprintf("inbox %s of sucker %s on chain %s",
		inboxTree.SuckerAddress, inboxTree.TokenAddress, inboxTree.ChainId)

	row := db.QueryRowContext(ctx, `SELECT current_root FROM trees
		WHERE chain_id = ? AND contract_address = ? AND token_address = ?`,
		inboxTree.ChainId, inboxTree.SuckerAddress, inboxTree.TokenAddress)

	if err = row.Scan(&dbRoot); err != nil && err != sql.ErrNoRows {
		return nil, fmt.Errorf("failed to query database for %s: %v", logDescription, err)
	}
	dbRootBytes, err := hex.DecodeString(dbRoot)
	// This should never happen, but just in case:
	if err != nil {
		return nil, fmt.Errorf("failed to decode dbRoot '%s' for %s: %v", dbRoot, logDescription, err)
	}

	// If no rows were found, or the roots don't match, we need to update the database.
	if err == sql.ErrNoRows || !bytes.Equal(dbRootBytes, inboxTree.Root[:]) {
		log.Printf("Updating the database for %s", logDescription)
		if err := updateLeaves(ctx, inboxTree); err != nil {
			return nil, fmt.Errorf("failed to update leaves for %s: %v", logDescription, err)
		}
	}

	rows, err := db.QueryContext(ctx, `SELECT leaf_hash FROM leaves
		WHERE chain_id = ? AND contract_address = ? AND token_address = ?`,
		inboxTree.ChainId, inboxTree.SuckerAddress, inboxTree.TokenAddress)
	if err != nil {
		// Note: this includes sql.ErrNoRows
		return nil, fmt.Errorf("failed to read leaves for %s from the database: %v", logDescription, err)
	}

	// Build the leaves at the bottom of the tree
	leaves := make([][]byte, 0)
	for rows.Next() {
		var leafHash string
		err := rows.Scan(&leafHash)
		if err != nil {
			return nil, fmt.Errorf("failed to scan leaf hash %s for %s: %v", leafHash, logDescription, err)
		}

		b, err := hex.DecodeString(leafHash)
		if err != nil {
			return nil, fmt.Errorf("failed to decode leaf hash %s for %s: %v", leafHash, logDescription, err)
		}

		leaves = append(leaves, b)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed while scanning leaves for %s from the database: %v", logDescription, err)
	}

	// Construct a tree with the retrieved leaves
	t := tree.NewTree(leaves)
	treeRoot := t.Root()

	// Sanity check the tree root
	if !bytes.Equal(treeRoot, inboxTree.Root[:]) {
		return nil, fmt.Errorf("constructed tree has root %s, does not match onchain tree root %s for %s",
			hex.EncodeToString(treeRoot), hex.EncodeToString(inboxTree.Root[:]), logDescription)
	}

	// We know the index is within uint bounds for 32/64-bit platforms because we check in the proof function.
	proofIndex := uint(index.Uint64())
	if proofIndex == 0 {
		return nil, fmt.Errorf("index %s is out of bounds for %s", index, logDescription)
	}

	// Get and return the proof
	p, err := t.Proof(proofIndex)
	if err != nil {
		return nil, fmt.Errorf("failed to get proof for %s: %v", logDescription, err)
	}

	return p, nil
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
	log.Printf("Local sucker: %s on chain %d; Peer sucker: %s on chain %d\n",
		inboxTree.SuckerAddress.String(), inboxTree.ChainId, peerSuckerAddr.String(), peerChainId)
	peerSucker, err := NewBPSucker(peerSuckerAddr, peerClient)
	if err != nil {
		return fmt.Errorf("failed to instantiate the peer sucker contract '%s' on chain %d: %v", peerSuckerAddr.String(), peerChainId, err)
	}

	// Make sure the peers match
	peerAddressOfPeerSucker, err := peerSucker.PEER(&bind.CallOpts{Context: ctx})
	if err != nil {
		return fmt.Errorf("failed to get peer address of peer sucker: %v", err)
	}
	if peerAddressOfPeerSucker.Cmp(inboxTree.SuckerAddress) != 0 {
		return fmt.Errorf("peer address of peer sucker '%s' is '%s', which does not match local sucker '%s'",
			peerSuckerAddr.String(), peerAddressOfPeerSucker.String(), inboxTree.SuckerAddress.String())
	}

	// Iterate through insertions to the peer sucker's outbox
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

	if len(leavesToInsert) == 0 {
		log.Printf("Found no leaves to insert for '%d:%s:%s'", inboxTree.ChainId, inboxTree.SuckerAddress.String(), inboxTree.TokenAddress.String())
		return nil
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
