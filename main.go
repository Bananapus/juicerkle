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
	"slices"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	_ "github.com/mattn/go-sqlite3"
)

type chainId int

const (
	mainnet  chainId = 1
	optimism chainId = 10
	base     chainId = 8453
	arbitrum chainId = 42161
)

const (
	sepolia   chainId = 11155111
	opSepolia chainId = 11155420
)

// Map of chainId to ethclient.Client
var clients = make(map[chainId]*ethclient.Client)

var networks = []struct {
	name    string
	chainId chainId
	rpcUrl  string
}{
	{
		name:    "mainnet",
		chainId: mainnet,
		rpcUrl:  "https://rpc.ankr.com/eth",
	},
	{
		name:    "sepolia",
		chainId: sepolia,
		rpcUrl:  "https://rpc.sepolia.org",
	},
	{
		name:    "optimism",
		chainId: optimism,
		rpcUrl:  "https://rpc.ankr.com/optimism",
	},
	{
		name:    "optimism sepolia",
		chainId: opSepolia,
		rpcUrl:  "https://sepolia.optimism.io",
	},
	{
		name:    "base",
		chainId: base,
		rpcUrl:  "https://rpc.ankr.com/base",
	},
	{
		name:    "arbitrum",
		chainId: arbitrum,
		rpcUrl:  "https://rpc.ankr.com/arbitrum",
	},
}

func main() {
	// Set up ETH clients
	for _, network := range networks {
		client, err := ethclient.Dial(network.rpcUrl)
		if err != nil {
			log.Fatalf("Failed to connect to the %s network: %v", network.name, err)
		}
		clients[network.chainId] = client
	}

	// Set up DB
	if err := initDb(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	http.HandleFunc("POST /proof", proof)
	http.ListenAndServe(":8080", nil)
}

type ProofRequest struct {
	ChainId chainId        `json:"chainId"` // The chain ID of the sucker contract
	Sucker  common.Address `json:"sucker"`  // The sucker contract address
	Token   common.Address `json:"token"`   // The address of the token being claimed
	Leaf    BPLeaf         `json:"leaf"`    // The leaf to prove on the sucker contract
}

func proof(w http.ResponseWriter, req *http.Request) {
	var toProve ProofRequest
	err := json.NewDecoder(req.Body).Decode(&toProve)
	if err != nil {
		log.Printf("Failed to parse request body: %v\n", err)
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	if toProve.Leaf.Index.Cmp(big.NewInt(0)) < 0 || toProve.Leaf.Index.Cmp(big.NewInt(1<<32)) >= 0 {
		http.Error(w, "Invalid leaf index", http.StatusBadRequest)
		return
	}

	// Set up cancellation context
	ctx, cancel := context.WithTimeout(req.Context(), 15*time.Second)
	defer cancel()

	// Set up channel to receive the proofCh
	proofCh := make(chan [][]byte)

	// Calculate the proof off of the main thread
	go func() {
		client, ok := clients[chainId(toProve.ChainId)]
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
			log.Printf("Failed to get tree: %v\n", err)
			http.Error(w, "Failed to get tree", http.StatusInternalServerError)
			return
		}

		proof, err := dbProof(ctx, toProve.Leaf, InboxTree{
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

// A BPLeaf as stored in sqlite
type Leaf struct {
	ChainId         int    `db:"chain_id"`
	ContractAddress string `db:"contract_address"`
	TokenAddress    string `db:"token_address"`
	LeafIndex       uint   `db:"leaf_index"`
	LeafHash        string `db:"leaf_hash"`
}

// Update the leaves in the database for a specific sucker
func updateLeaves(ctx context.Context, tokenTree InboxTree) error {
	client, ok := clients[tokenTree.ChainId]
	if !ok {
		return fmt.Errorf("chain %d not supported", tokenTree.ChainId)
	}

	sucker, err := NewBPSucker(tokenTree.SuckerAddress, client)
	if err != nil {
		return fmt.Errorf("failed to instantiate the sucker contract: %v", err)
	}

	// Get the latest hash from the db
	var latestHash string
	err = db.QueryRowContext(ctx, `SELECT leaf_hash FROM leaves
		WHERE chain_id = ? AND contract_address = ? AND token_address = ?
		ORDER BY leaf_index DESC LIMIT 1`,
		tokenTree.ChainId, tokenTree.SuckerAddress.String(), tokenTree.TokenAddress.String()).Scan(&latestHash)
	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("failed to query database: %v", err)
	}

	seenLatestDbLeaf := false
	var latestHashBytes []byte
	// Start from the beginning if there are no leaves in the db
	if err == sql.ErrNoRows {
		seenLatestDbLeaf = true
	} else {
		if latestHashBytes, err = hex.DecodeString(latestHash); err != nil {
			return fmt.Errorf("failed to decode leaf hash '%s': %v", latestHash, err)
		}
	}

	inboxTreeRoot, err := sucker.Inbox(&bind.CallOpts{Context: ctx}, tokenTree.TokenAddress)
	if err != nil {
		return fmt.Errorf("failed to get inbox tree root for token '%s': %v",
			tokenTree.TokenAddress.String(), err)
	}

	peerSuckerAddr, err := sucker.PEER(&bind.CallOpts{Context: ctx})
	if err != nil {
		return fmt.Errorf("failed to get peer: %v", err)
	}

	var peerChainId chainId
	switch tokenTree.ChainId {
	case sepolia:
		peerChainId = opSepolia
	case opSepolia:
		peerChainId = sepolia
	default:
		return fmt.Errorf("peer chain not supported")
	}

	peerClient, ok := clients[peerChainId]
	if !ok {
		return fmt.Errorf("peer chain %d not supported", peerChainId)
	}

	peerSucker, err := NewBPSucker(peerSuckerAddr, peerClient)
	if err != nil {
		return fmt.Errorf("failed to instantiate the peer sucker contract '%s' on chain %d: %v", peerSuckerAddr.String(), peerChainId, err)
	}

	outboxIterator, err := peerSucker.FilterInsertToOutboxTree(&bind.FilterOpts{Context: ctx}, nil, []common.Address{tokenTree.TokenAddress})
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
			ChainId:         int(tokenTree.ChainId),
			ContractAddress: tokenTree.SuckerAddress.String(),
			LeafIndex:       uint(outboxIterator.Event.Index.Uint64()),
			TokenAddress:    tokenTree.TokenAddress.String(),
			LeafHash:        hex.EncodeToString(outboxIterator.Event.Hashed[:]),
		})

		// If we've gotten to the latest root, break
		if inboxTreeRoot.Root == outboxIterator.Event.Root {
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
		VALUES (?, ?, ?, ?, ?)`, tokenTree.ChainId, tokenTree.SuckerAddress.String(), tokenTree.TokenAddress.String(),
		hex.EncodeToString(inboxTreeRoot.Root[:]), finalCount)

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit sqlite transaction: %v", err)
	}

	return nil
}

func (leaf BPLeaf) hash() (common.Hash, error) {
	uint256Ty, _ := abi.NewType("uint256", "", nil)
	addressTy, _ := abi.NewType("address", "", nil)

	args := abi.Arguments{
		{Name: "projectTokenAmount", Type: uint256Ty},
		{Name: "terminalTokenAmount", Type: uint256Ty},
		{Name: "beneficiary", Type: addressTy},
	}

	bytes, err := args.Pack(leaf.ProjectTokenAmount, leaf.TerminalTokenAmount, leaf.Beneficiary)
	if err != nil {
		return common.Hash{}, err
	}

	hash := crypto.Keccak256(bytes)
	return common.BytesToHash(hash), nil
}

// A specification for an inbox tree on a sucker contract
type InboxTree struct {
	ChainId       chainId
	SuckerAddress common.Address
	TokenAddress  common.Address
	Root          common.Hash
}

// Get the proof for a leaf from the database
func dbProof(ctx context.Context, leaf BPLeaf, inboxTree InboxTree) ([][]byte, error) {
	var dbRoot string
	var err error

	if err = db.QueryRowContext(ctx, `SELECT current_root FROM trees
		WHERE chain_id = ? AND contract_address = ? AND token_address = ?`,
		inboxTree.ChainId, inboxTree.SuckerAddress.String(), inboxTree.TokenAddress.String()).
		Scan(&dbRoot); err != nil && err != sql.ErrNoRows {
		return nil, fmt.Errorf("failed to query database")
	}

	// If no rows were found, or the counts/roots don't match, we need to update the database.
	if err == sql.ErrNoRows || dbRoot != inboxTree.Root.String() {
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
	if !slices.Equal(treeRoot, inboxTree.Root.Bytes()) {
		return nil, fmt.Errorf("calculated tree root '%s' does not match expected tree root '%s'",
			hex.EncodeToString(treeRoot), inboxTree.Root.String())
	}

	// Bounds are checked in the proof function
	p, err := t.Proof(uint(leaf.Index.Uint64()))
	if err != nil {
		return nil, fmt.Errorf("failed to get proof: %v", err)
	}

	return p, nil
}
