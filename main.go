package main

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

var networks = []struct {
	name    string
	chainId int
	rpcUrl  string
}{
	{
		name:    "mainnet",
		chainId: 1,
		rpcUrl:  "https://rpc.ankr.com/eth",
	},
	{
		name:    "optimism",
		chainId: 10,
		rpcUrl:  "https://rpc.ankr.com/optimism",
	},
	{
		name:    "base",
		chainId: 8453,
		rpcUrl:  "https://rpc.ankr.com/base",
	},
}

type BPLeaf struct {
	index               *big.Int
	beneficiary         common.Address
	projectTokenAmount  *big.Int
	terminalTokenAmount *big.Int
}

func (leaf BPLeaf) hash() common.Hash {
	return common.Hash{}
}

func proof(chaindId int, sucker common.Address, leaf BPLeaf) (proof [][]byte, err error) {

	return
}
