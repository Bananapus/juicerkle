package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
	optimism         = 10
	base             = 8453
	arbitrum         = 42161
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
		name:    "optimism",
		chainId: optimism,
		rpcUrl:  "https://rpc.ankr.com/optimism",
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

/*type BPLeaf struct {
	Index               *big.Int       `json:"index"`
	Beneficiary         common.Address `json:"beneficiary"`
	ProjectTokenAmount  *big.Int       `json:"projectTokenAmount"`
	TerminalTokenAmount *big.Int       `json:"terminalTokenAmount"`
}*/

type ProofRequest struct {
	ChainId int    `json:"chainId"` // The chain ID of the sucker contract
	Sucker  string `json:"sucker"`  // The sucker contract address
	Leaf    BPLeaf `json:"leaf"`    // The leaf to prove on the sucker contract
}

func proof(w http.ResponseWriter, req *http.Request) {
	var toProve ProofRequest
	err := json.NewDecoder(req.Body).Decode(&toProve)
	if err != nil {
		log.Printf("Failed to parse request body: %v\n", err)
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	// Set up cancellation context
	ctx, cancel := context.WithTimeout(req.Context(), 15*time.Second)
	defer cancel()

	// Set up channel to receive the proof
	proof := make(chan [][]byte)

	// Calculate the proof off of the main thread
	go func() {
		client, ok := clients[chainId(toProve.ChainId)]
		if !ok {
			log.Printf("Chain %d not supported\n", toProve.ChainId)
			http.Error(w, "Chain not supported", http.StatusBadRequest)
			return
		}

		localSucker, err := NewBPSucker(common.HexToAddress(toProve.Sucker), client)
		if err != nil {
			log.Printf("Failed to instantiate the sucker contract: %v\n", err)
			http.Error(w, "Failed to instantiate the sucker contract", http.StatusInternalServerError)
			return
		}

		peerChainId, peerContractAddr, err := peer(toProve.ChainId, toProve.Sucker, localSucker, ctx)
		if err != nil {
			http.Error(w, "Failed to get peer", http.StatusInternalServerError)
			return
		}

		proof <- nil
	}()

	// Wait for the proof or a cancellation
	select {
	case <-ctx.Done():
		log.Printf("Request cancelled: %v\n", ctx.Err())
		http.Error(w, "Request cancelled", http.StatusRequestTimeout)
		return
	case p := <-proof:
		b, err := json.Marshal(p)
		if err != nil {
			log.Printf("Failed to marshal proof: %v\n", err)
			http.Error(w, "Failed to marshal proof", http.StatusInternalServerError)
			return
		}

		w.Write(b)
	}
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

// Get the peer chain ID and contract address for a sucker
func peer(suckerChainId int, suckerAddress string, localSucker *BPSucker, ctx context.Context) (chainId, common.Address, error) {
	var dbPeerChainId int
	var dbPeerContractAddr string

	peerCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// Check the db for the sucker contract's peer
	err := db.QueryRowContext(peerCtx, `SELECT peer_chain_id, peer_contract_address FROM suckers
		WHERE chain_id = ? AND contract_address = ?`,
		suckerChainId, suckerAddress).Scan(&dbPeerChainId, &dbPeerContractAddr)

	select {
	case <-peerCtx.Done():
		log.Printf("Request cancelled: %v\n", peerCtx.Err())
		return 0, common.Address{}, fmt.Errorf("request cancelled")
	default:
		// Continue
	}

	// If the sucker was found in the db
	if err == nil {
		// Convert the results to the correct types and return
		return chainId(dbPeerChainId), common.HexToAddress(dbPeerContractAddr), nil
	} else if err == sql.ErrNoRows {
		// If the sucker wasn't found in the db, we have to read the data from the blockchain
		// TODO: Read the peer chain ID from the sucker. Mocking for now
		peerChainId := chainId(1)

		peerContractAddr, err := localSucker.PEER(&bind.CallOpts{Context: peerCtx})
		if err != nil {
			log.Printf("Failed to get peer from sucker: %v\n", err)
			return 0, common.Address{}, fmt.Errorf("failed to get peer from sucker")
		}

		// Store the sucker in the database for next time.
		db.Exec(`INSERT INTO suckers (?, ?, ?, ?)`,
			suckerChainId, suckerAddress,
			int(peerChainId), peerContractAddr.String(),
		)

		return peerChainId, peerContractAddr, nil
	} else {
		return 0, common.Address{}, fmt.Errorf("failed to query database")
	}
}
