package main

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestBPLeafHashing(t *testing.T) {
	var tests = []struct {
		leaf         BPLeaf
		expectedHash common.Hash
	}{
		{
			leaf: BPLeaf{
				index:               big.NewInt(0),
				beneficiary:         common.HexToAddress("0x823b92d6a4b2AED4b15675c7917c9f922ea8ADAD"),
				projectTokenAmount:  big.NewInt(1000),
				terminalTokenAmount: big.NewInt(500),
			},
			expectedHash: common.HexToHash("0xf9a8b58bcc5c9af6169ceffcaedb846ba4b57da24a42783dd3efb92641d993c0"),
		},
		{
			leaf: BPLeaf{
				index:               big.NewInt(0),
				beneficiary:         common.HexToAddress("0xAF28bcB48C40dBC86f52D459A6562F658fc94B1e"),
				projectTokenAmount:  big.NewInt(5000),
				terminalTokenAmount: big.NewInt(2500),
			},
			expectedHash: common.HexToHash("0x925615b3b23aad3e3d18b2bea82dfa3b7efd7881efb993d063c8741008bc1a39"),
		},
		{
			leaf: BPLeaf{
				index:               big.NewInt(0),
				beneficiary:         common.HexToAddress("0x30670D81E487c80b9EDc54370e6EaF943B6EAB39"),
				projectTokenAmount:  big.NewInt(10000),
				terminalTokenAmount: big.NewInt(7500),
			},
			expectedHash: common.HexToHash("0x61b21ed31b56f26f070d6724c8ebddf16a342426638c3e2683bf2349f4e03fca"),
		},
	}

	for _, test := range tests {
		hash := test.leaf.hash()
		if hash != test.expectedHash {
			t.Errorf("expected %s, got %s", test.expectedHash.Hex(), hash.Hex())
		}
	}
}
