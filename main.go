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
	"strconv"

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
type DBLeaf struct {
	ChainId             string `db:"chain_id"`
	ContractAddress     string `db:"contract_address"`
	TokenAddress        string `db:"token_address"`
	Index               string `db:"leaf_index"`
	Beneficiary         string `db:"leaf_beneficiary"`
	ProjectTokenAmount  string `db:"leaf_project_token_amount"`
	TerminalTokenAmount string `db:"leaf_terminal_token_amount"`
	LeafHash            string `db:"leaf_hash"`
	IsClaimed           bool   `db:"is_claimed"`
}

// A specification for an inbox tree on a sucker contract
type InboxTree struct {
	ChainId       *big.Int
	SuckerAddress common.Address
	TokenAddress  common.Address
	Root          [32]byte
}

// Schema for incoming claims requests
type ClaimsRequest struct {
	ChainId     *big.Int       `json:"chainId"`     // The chain ID of the sucker contract
	Sucker      common.Address `json:"sucker"`      // The sucker contract address
	Token       common.Address `json:"token"`       // The token address of the inbox tree being claimed from
	Beneficiary common.Address `json:"beneficiary"` // The address of the beneficiary to get the claims for
}

// Map of chain IDs (as strings) to ETH clients
var clients = make(map[string]*ethclient.Client)

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
		clients[network.ChainId.String()] = client
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

	http.HandleFunc("POST /claims", claims)
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Juicerkle running. Send claims requests to /claims."))
	})
	log.Printf("Listening on http://localhost:%s\n", port)
	http.ListenAndServe(":"+port, nil)
}

var (
	MaxIndex = big.NewInt(1 << 32)
	ZeroRoot = make([]byte, 32)
)

func claims(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	var toClaim ClaimsRequest
	err := json.NewDecoder(req.Body).Decode(&toClaim)
	if err != nil {
		errStr := fmt.Sprintf("Failed to parse claims request body: %v", err)
		log.Println(errStr)
		http.Error(w, errStr, http.StatusBadRequest)
		return
	}

	// Set up context
	// ctx, cancel := context.WithTimeout(req.Context(), 15*time.Second)
	ctx, cancel := context.WithCancel(req.Context()) // Use cancel instead of timeout for debugging
	defer cancel()

	logDescription := fmt.Sprintf("beneficiary %s; inbox %s of sucker %s on chain %s",
		toClaim.Beneficiary, toClaim.Token, toClaim.Sucker, toClaim.ChainId)

	client, ok := clients[toClaim.ChainId.String()]
	if !ok {
		errStr := fmt.Sprintf("No RPC for chain ID %s. Contact the developer to have it added.", toClaim.ChainId)
		log.Println(errStr + " For " + logDescription)
		http.Error(w, errStr, http.StatusNotFound)
		return
	}

	localSucker, err := NewBPSucker(toClaim.Sucker, client)
	if err != nil {
		errStr := fmt.Sprintf("Failed to instantiate a BPSucker for %s: %v", logDescription, err)
		log.Println(errStr)
		http.Error(w, errStr, http.StatusInternalServerError)
		return
	}

	// callOpts with our context
	callOpts := &bind.CallOpts{Context: ctx}
	localInboxTree, err := localSucker.Inbox(callOpts, toClaim.Token)
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
		http.Error(w, errStr, http.StatusBadRequest)
		return
	}

	// Get the claims
	claims, err := dbClaims(ctx, toClaim.Beneficiary, InboxTree{
		ChainId:       toClaim.ChainId,
		SuckerAddress: toClaim.Sucker,
		TokenAddress:  toClaim.Token,
		Root:          localInboxTree.Root,
	})
	if err != nil {
		errStr := fmt.Sprintf("Failed to get claims for %s: %v", logDescription, err)
		log.Println(errStr)
		http.Error(w, errStr, http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(claims)
	if err != nil {
		errStr := fmt.Sprintf("Failed to marshal claims for %s: %v", logDescription, err)
		log.Println(errStr)
		http.Error(w, errStr, http.StatusInternalServerError)
		return
	}

	w.Write(b)
}

// Get the currently available BPClaims for a beneficiary from the database.
// If the database is out of date for the given inbox tree, update it.
func dbClaims(ctx context.Context, beneficiary common.Address, inboxTree InboxTree) ([]BPClaim, error) {
	logDescription := fmt.Sprintf("inbox %s of sucker %s on chain %s",
		inboxTree.TokenAddress, inboxTree.SuckerAddress, inboxTree.ChainId)

	// Get the current root from the database (which may be out of date)
	row := db.QueryRowContext(ctx, `SELECT current_root FROM trees
		WHERE chain_id = ? AND contract_address = ? AND token_address = ?`,
		inboxTree.ChainId.String(), inboxTree.SuckerAddress.String(), inboxTree.TokenAddress.String())

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

	// Update the database with the latest claims for this inbox tree.
	if err := updateClaims(ctx, inboxTree); err != nil {
		return nil, fmt.Errorf("failed to update claims for %s: %v", logDescription, err)
	}

	rows, err := db.QueryContext(ctx, `SELECT leaf_hash, leaf_index, leaf_beneficiary,
		leaf_project_token_amount, leaf_terminal_token_amount, is_claimed
		FROM leaves WHERE chain_id = ? AND contract_address = ? AND token_address = ?
		ORDER BY leaf_index ASC`,
		inboxTree.ChainId.String(), inboxTree.SuckerAddress.String(), inboxTree.TokenAddress.String())
	if err != nil {
		// Note: this includes sql.ErrNoRows
		return nil, fmt.Errorf("failed to read leaves for %s from the database: %v", logDescription, err)
	}

	// Build the tree and claims from the database leaves
	claims := make([]BPClaim, 0)
	leaves := make([][]byte, 0)

	// TODO: Check whether the leaf has already been claimed.
	defer rows.Close()
	for rows.Next() {
		var leafHash, index, leafBeneficiary, projectTokenAmount, terminalTokenAmount string
		var isClaimed bool
		err := rows.Scan(&leafHash, &index, &leafBeneficiary, &projectTokenAmount, &terminalTokenAmount, &isClaimed)
		if err != nil {
			return nil, fmt.Errorf("failed to scan leaf hash %s for %s: %v", leafHash, logDescription, err)
		}

		h, err := hex.DecodeString(leafHash)
		if err != nil {
			return nil, fmt.Errorf("failed to decode leaf hash %s for %s: %v", leafHash, logDescription, err)
		}

		leaves = append(leaves, h)

		// If this leaf has already been claimed, move on.
		if isClaimed {
			continue
		}

		// Check if the leaf is for the beneficiary we're looking for, and if so, add it to the claims
		if !common.IsHexAddress(leafBeneficiary) {
			return nil, fmt.Errorf("db leaf beneficiary %s is not a valid address for %s", leafBeneficiary, logDescription)
		}
		b := common.HexToAddress(leafBeneficiary)

		if beneficiary.Cmp(b) == 0 {
			idx, success := big.NewInt(0).SetString(index, 10)
			if !success {
				return nil, fmt.Errorf("failed to parse index %s for %s: %v", index, logDescription, err)
			}
			pt, success := big.NewInt(0).SetString(projectTokenAmount, 10)
			if !success {
				return nil, fmt.Errorf("failed to parse projectTokenAmount %s for %s: %v", projectTokenAmount, logDescription, err)
			}
			tt, success := big.NewInt(0).SetString(terminalTokenAmount, 10)
			if !success {
				return nil, fmt.Errorf("failed to parse terminalTokenAmount %s for %s: %v", terminalTokenAmount, logDescription, err)
			}

			claims = append(claims, BPClaim{
				Token: inboxTree.TokenAddress,
				Leaf: BPLeaf{
					Index:               idx,
					Beneficiary:         beneficiary,
					ProjectTokenAmount:  pt,
					TerminalTokenAmount: tt,
				},
			})
		}
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

	// Add the proofs to the claims
	for i := range claims {
		// We know the index is within uint bounds for 32/64-bit platforms because there can only be 2^32 leaves
		proofIndex := uint(claims[i].Leaf.Index.Uint64())
		p, err := t.Proof(proofIndex)
		if err != nil {
			return nil, fmt.Errorf("failed to get proof at index %d for %s: %v", proofIndex, logDescription, err)
		}
		proofArray, err := proofSliceToArray(p)
		if err != nil {
			return nil, fmt.Errorf("failed to convert proof to arrays for %s: %v", logDescription, err)
		}
		claims[i].Proof = proofArray
	}

	return claims, nil
}

// Update the leaves in the database for a specific inbox tree
func updateLeaves(ctx context.Context, inboxTree InboxTree) error {
	logDescription := fmt.Sprintf("inbox %s of sucker %s on chain %s",
		inboxTree.TokenAddress, inboxTree.SuckerAddress, inboxTree.ChainId)

	client := clients[inboxTree.ChainId.String()] // chain was checked in the handler function

	// Get the latest leaf hash for the tree from the db
	row := db.QueryRowContext(ctx, `SELECT leaf_hash FROM leaves
		WHERE chain_id = ? AND contract_address = ? AND token_address = ?
		ORDER BY leaf_index DESC LIMIT 1`,
		inboxTree.ChainId.String(), inboxTree.SuckerAddress.String(), inboxTree.TokenAddress.String())

	var latestHash string
	err := row.Scan(&latestHash)
	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("failed to query latest leaf from database for %s: %v", logDescription, err)
	}

	seenLatestDbLeaf := false
	var latestHashBytes []byte
	if err == sql.ErrNoRows {
		// Start parsing events from the beginning if our query didn't return any leaf
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
	peerClient, ok := clients[peerSuckerChainId.String()]
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

	leavesToInsert := make([]DBLeaf, 0)
	defer outboxIterator.Close()
	for outboxIterator.Next() {
		// TODO: Remove this once we know it's working
		log.Printf("Leaf index: %s; project token amount: %s; terminal token amount: %s",
			outboxIterator.Event.Index.String(), outboxIterator.Event.ProjectTokenAmount.String(), outboxIterator.Event.TerminalTokenAmount.String())

		// Keep skipping until we pass the latest hash
		if !seenLatestDbLeaf {
			if bytes.Equal(outboxIterator.Event.Hashed[:], latestHashBytes) {
				seenLatestDbLeaf = true
			}
			continue
		}

		// Add the remaining leaves to the list to insert
		// Leaves are associated with their inbox tree
		leavesToInsert = append(leavesToInsert, DBLeaf{
			ChainId:             inboxTree.ChainId.String(),
			ContractAddress:     inboxTree.SuckerAddress.String(),
			TokenAddress:        inboxTree.TokenAddress.String(),
			Index:               outboxIterator.Event.Index.String(),
			Beneficiary:         outboxIterator.Event.Beneficiary.String(),
			ProjectTokenAmount:  outboxIterator.Event.ProjectTokenAmount.String(),
			TerminalTokenAmount: outboxIterator.Event.TerminalTokenAmount.String(),
			LeafHash:            hex.EncodeToString(outboxIterator.Event.Hashed[:]),
			IsClaimed:           false,
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

	stmt, err := tx.PrepareContext(ctx, `INSERT INTO leaves 
		(chain_id, contract_address, token_address, leaf_index, leaf_beneficiary,
		leaf_project_token_amount, leaf_terminal_token_amount, leaf_hash, is_claimed)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return fmt.Errorf("failed to prepare sqlite statement: %v", err)
	}
	defer stmt.Close()

	// Insert the leaves
	for _, leaf := range leavesToInsert {
		_, err := stmt.ExecContext(ctx, leaf.ChainId, leaf.ContractAddress, leaf.TokenAddress, leaf.Index, leaf.Beneficiary,
			leaf.ProjectTokenAmount, leaf.TerminalTokenAmount, leaf.LeafHash, leaf.IsClaimed)
		if err != nil {
			tx.Rollback() // Rollback the transaction if an error occurs
			return fmt.Errorf("failed to insert leaf into sqlite for %s: %v", logDescription, err)
		}
	}

	// Update the inbox tree root
	if _, err = tx.ExecContext(ctx, `INSERT OR REPLACE INTO trees(chain_id, contract_address, token_address, current_root, block_claims_last_updated)
		VALUES (?, ?, ?, ?, ?)`,
		inboxTree.ChainId.String(), inboxTree.SuckerAddress.String(), inboxTree.TokenAddress.String(), hex.EncodeToString(inboxTree.Root[:]), "0"); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update inbox tree root for %s: %v", logDescription, err)
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit sqlite transaction for %s: %v", logDescription, err)
	}

	return nil
}

// Updates the database with the latest claims for a specific inbox tree.
func updateClaims(ctx context.Context, inboxTree InboxTree) error {
	logDescription := fmt.Sprintf("inbox %s of sucker %s on chain %s",
		inboxTree.TokenAddress, inboxTree.SuckerAddress, inboxTree.ChainId)

	// Read the latest block that we've checked for claims in from the db.
	row := db.QueryRowContext(ctx, `SELECT block_claims_last_updated FROM trees
		WHERE chain_id = ? AND contract_address = ? AND token_address = ?`,
		inboxTree.ChainId.String(), inboxTree.SuckerAddress.String(), inboxTree.TokenAddress.String())

	var lastUpdatedBlock string
	err := row.Scan(&lastUpdatedBlock)
	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("failed to query block_claims_last_updated from database for %s: %v", logDescription, err)
	}
	if err == sql.ErrNoRows {
		lastUpdatedBlock = "0"
	}
	startBlock, err := strconv.ParseUint(lastUpdatedBlock, 10, 64)
	if err != nil {
		return fmt.Errorf("failed to parse last updated block %s for %s: %v", lastUpdatedBlock, logDescription, err)
	}

	// Get the sucker and an iterator for Claimed events.
	client := clients[inboxTree.ChainId.String()]
	sucker, err := NewBPSucker(inboxTree.SuckerAddress, client)
	if err != nil {
		return fmt.Errorf("failed to instantiate local sucker contract for %s: %v", logDescription, err)
	}
	claimIterator, err := sucker.FilterClaimed(&bind.FilterOpts{Context: ctx, Start: startBlock})
	if err != nil {
		return fmt.Errorf("failed to instantiate Claimed iterator for %s: %v", logDescription, err)
	}
	defer claimIterator.Close()

	// Iterate through the claimed events and update the leaves in the db accordingly
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to start sqlite transaction for %s: %v", logDescription, err)
	}

	for claimIterator.Next() {
		// If this event is for a different inbox tree, skip it
		if claimIterator.Event.Token.Cmp(inboxTree.TokenAddress) != 0 {
			continue
		}

		_, err := tx.ExecContext(ctx, `UPDATE leaves SET is_claimed = 1
			WHERE chain_id = ? AND contract_address = ? AND token_address = ? AND leaf_index = ?`,
			inboxTree.ChainId.String(), inboxTree.SuckerAddress.String(), inboxTree.TokenAddress.String(), claimIterator.Event.Index.String())
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed while executing update sqlite transaction for %s: %v", logDescription, err)
		}
	}
	if err := claimIterator.Error(); err != nil {
		return fmt.Errorf("failed while iterating through claimed events for %s: %v", logDescription, err)
	}

	// Update the last block that we've checked for claims in
	tx.ExecContext(ctx, `UPDATE trees SET block_claims_last_updated = ?
		WHERE chain_id = ? AND contract_address = ? AND token_address = ?`,
		strconv.FormatUint(claimIterator.Event.Raw.BlockNumber, 10),
		inboxTree.ChainId.String(), inboxTree.SuckerAddress.String(), inboxTree.TokenAddress.String())
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed while updating block_claims_last_updated for %s: %v", logDescription, err)
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit sqlite transaction for %s: %v", logDescription, err)
	}

	return nil
}

// Convert a proof from slices to arrays, and make sure it's the right length.
func proofSliceToArray(input [][]byte) ([32][32]byte, error) {
	var output [32][32]byte

	if len(input) != 32 {
		return output, fmt.Errorf("input does not have exactly 32 elements")
	}

	for i, slice := range input {
		if len(slice) != 32 {
			return output, fmt.Errorf("slice %d is not exactly 32 bytes long", i)
		}
		copy(output[i][:], slice)
	}

	return output, nil
}
