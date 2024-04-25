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

	logDescription := fmt.Sprintf("inbox %s of sucker %s on chain %s",
		toProve.Token, toProve.Sucker, toProve.ChainId)

	client, ok := clients[toProve.ChainId]
	if !ok {
		errStr := fmt.Sprintf("No RPC for chain ID %s. Contact the developer to have it added.", toProve.ChainId)
		log.Println(errStr + logDescription)
		http.Error(w, errStr, http.StatusNotFound)
		return
	}

	localSucker, err := NewBPSucker(toProve.Sucker, client)
	if err != nil {
		errStr := fmt.Sprintf("Failed to instantiate a BPSucker for %s: %v", logDescription, err)
		log.Println(errStr)
		http.Error(w, errStr, http.StatusInternalServerError)
		return
	}

	// Default callOpts
	callOpts := &bind.CallOpts{Context: ctx}

	localInboxTree, err := localSucker.Inbox(callOpts, toProve.Token)
	if err != nil {
		errStr := fmt.Sprintf("Failed to get inbox tree root for %s: %v", logDescription, err)
		log.Println(errStr)
		http.Error(w, errStr, http.StatusInternalServerError)
		return
	}

	// Check if the inbox tree root is empty.
	if bytes.Equal(localInboxTree.Root[:], ZeroRoot) {
		errStr := fmt.Sprintf("Inbox tree root for %s is empty", logDescription)
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
		errStr := fmt.Sprintf("Failed to get proof for %s: %v", logDescription, err)
		log.Println(errStr)
		http.Error(w, errStr, http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(proof)
	if err != nil {
		errStr := fmt.Sprintf("Failed to marshal proof for %s: %v", logDescription, err)
		log.Println(errStr)
		http.Error(w, errStr, http.StatusInternalServerError)
		return
	}

	w.Write(b)
}

// Get the proof for a leaf from the database. If the database is out of date, update it.
func dbProof(ctx context.Context, index *big.Int, inboxTree InboxTree) ([][]byte, error) {
	logDescription := fmt.Sprintf("inbox %s of sucker %s on chain %s",
		inboxTree.SuckerAddress, inboxTree.TokenAddress, inboxTree.ChainId)

	// Get the current root from the database (which may be out of date)
	row := db.QueryRowContext(ctx, `SELECT current_root FROM trees
		WHERE chain_id = ? AND contract_address = ? AND token_address = ?`,
		inboxTree.ChainId, inboxTree.SuckerAddress, inboxTree.TokenAddress)

	var dbRoot string
	if err := row.Scan(&dbRoot); err != nil && err != sql.ErrNoRows {
		return nil, fmt.Errorf("failed to query root from database for %s: %v", logDescription, err)
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

// Update the leaves in the database for a specific inbox tree
func updateLeaves(ctx context.Context, inboxTree InboxTree) error {
	logDescription := fmt.Sprintf("inbox %s of sucker %s on chain %s",
		inboxTree.SuckerAddress, inboxTree.TokenAddress, inboxTree.ChainId)

	client := clients[inboxTree.ChainId] // chain was checked in the proof function

	// Get the latest leaf hash for the tree from the db
	row := db.QueryRowContext(ctx, `SELECT leaf_hash FROM leaves
		WHERE chain_id = ? AND contract_address = ? AND token_address = ?
		ORDER BY leaf_index DESC LIMIT 1`,
		inboxTree.ChainId, inboxTree.SuckerAddress, inboxTree.TokenAddress)

	var latestHash string
	err := row.Scan(&latestHash)
	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("failed to query latest leaf from database for %s: %v", logDescription, err)
	}

	seenLatestDbLeaf := false
	var latestHashBytes []byte
	if err == sql.ErrNoRows {
		// Start parsing events from the beginning if there are no leaves in the db
		seenLatestDbLeaf = true
	} else {
		if latestHashBytes, err = hex.DecodeString(latestHash); err != nil {
			return fmt.Errorf("failed to decode leaf hash %s from db for %s: %v", latestHash, logDescription, err)
		}
	}

	// Instantiate the local sucker
	localSucker, err := NewBPSucker(inboxTree.SuckerAddress, client)
	if err != nil {
		return fmt.Errorf("failed to instantiate the sucker contract for %s: %v", logDescription, err)
	}

	// TODO: We can parallelize some of these reads if we need to reduce latency.

	// We need to read the InsertToOutboxTree events from the peer sucker.
	// First, get the peer sucker's chain ID and address.
	callOpts := &bind.CallOpts{Context: ctx}
	peerSuckerChainId, err := localSucker.PeerChainID(callOpts)
	if err != nil {
		return fmt.Errorf("failed to get peer chain ID for %s: %v", logDescription, err)
	}
	peerSuckerAddr, err := localSucker.PEER(callOpts)
	if err != nil {
		return fmt.Errorf("failed to get peer address for %s: %v", logDescription, err)
	}
	peerClient, ok := clients[peerSuckerChainId]
	if !ok {
		return fmt.Errorf("no RPC for peer chain %d, peer for %s; contact the developer to have it added", peerSuckerChainId, logDescription)
	}

	// Instantiate the peer sucker
	peerSucker, err := NewBPSucker(peerSuckerAddr, peerClient)
	if err != nil {
		return fmt.Errorf("failed to instantiate the peer sucker at %s on chain %d for %s: %v", peerSuckerAddr, peerSuckerChainId, logDescription, err)
	}

	// Add peer information to further logging
	logDescription += fmt.Sprintf(" with peer sucker %s on chain %d", peerSuckerAddr, peerSuckerChainId)

	// Make sure the peers match
	peerAddressOfPeerSucker, err := peerSucker.PEER(callOpts)
	if err != nil {
		return fmt.Errorf("failed to get peer address of peer sucker: %v", err)
	}
	if peerAddressOfPeerSucker.Cmp(inboxTree.SuckerAddress) != 0 {
		return fmt.Errorf("peer address of peer sucker %s on chain %d is %s, which does not match local sucker %s on chain %d",
			peerSuckerAddr, peerSuckerChainId, peerAddressOfPeerSucker, inboxTree.SuckerAddress, inboxTree.ChainId)
	}

	// We also need to know what peer outbox token address the local inbox tree corresponds to
	remoteToken, err := localSucker.RemoteTokenFor(callOpts, inboxTree.TokenAddress)
	if err != nil {
		return fmt.Errorf("failed to get remote token for %s: %v", logDescription, err)
	}

	// Iterate through insertions to the peer sucker's outbox
	outboxIterator, err := peerSucker.FilterInsertToOutboxTree(
		&bind.FilterOpts{Context: ctx},
		nil,
		[]common.Address{remoteToken.Addr}, // Only get logs where the terminal token matches the correct remote token
	)
	if err != nil {
		return fmt.Errorf("failed to instantiate peer outbox iterator for peer in request for %s: %v", logDescription, err)
	}
	defer outboxIterator.Close()

	leavesToInsert := make([]Leaf, 0)
	for outboxIterator.Next() {
		log.Printf("Event: %+v", outboxIterator.Event)

		// Keep iterating until we pass the latest hash
		if !seenLatestDbLeaf {
			if bytes.Equal(outboxIterator.Event.Hashed[:], latestHashBytes) {
				seenLatestDbLeaf = true
			}
			continue
		}

		// Add the remaining leaves to the list to insert
		// Leaves are associated with their inbox tree
		leavesToInsert = append(leavesToInsert, Leaf{
			ChainId:         inboxTree.ChainId.String(),
			ContractAddress: inboxTree.SuckerAddress.String(),
			LeafIndex:       outboxIterator.Event.Index.String(),
			TokenAddress:    inboxTree.TokenAddress.String(),
			LeafHash:        hex.EncodeToString(outboxIterator.Event.Hashed[:]),
		})

		// If we've gotten to the latest inbox root, break
		if inboxTree.Root == outboxIterator.Event.Root {
			break
		}
	}

	if err := outboxIterator.Error(); err != nil {
		return fmt.Errorf("failed while iterating through outbox insertions for %s: %v", logDescription, err)
	}

	if len(leavesToInsert) == 0 {
		log.Printf("found no leaves to insert for %s", logDescription)
		return nil
	}

	// Start sqlite transaction to insert leaves
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to start sqlite transaction for %s: %v", logDescription, err)
	}

	stmt, err := tx.PrepareContext(ctx, `INSERT INTO leaves (chain_id, contract_address, token_address, leaf_index, leaf_hash)
		VALUES (?, ?, ?, ?, ?)`)
	if err != nil {
		return fmt.Errorf("failed to prepare sqlite statement: %v", err)
	}
	defer stmt.Close()

	// Insert the leaves
	for _, leaf := range leavesToInsert {
		_, err := stmt.ExecContext(ctx, leaf.ChainId, leaf.ContractAddress, leaf.TokenAddress, leaf.LeafIndex, leaf.LeafHash)
		if err != nil {
			tx.Rollback() // Rollback the transaction if an error occurs
			return fmt.Errorf("failed to insert leaf into sqlite for %s: %v", logDescription, err)
		}
	}

	// Update the inbox tree root
	finalCount := leavesToInsert[len(leavesToInsert)-1].LeafIndex // we checked that len != 0 above
	tx.ExecContext(ctx, `INSERT OR REPLACE INTO trees (chain_id, contract_address, token_address, current_root, count)
		VALUES (?, ?, ?, ?, ?)`, inboxTree.ChainId, inboxTree.SuckerAddress, inboxTree.TokenAddress,
		hex.EncodeToString(inboxTree.Root[:]), finalCount)

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit sqlite transaction for %s: %v", logDescription, err)
	}

	return nil
}
