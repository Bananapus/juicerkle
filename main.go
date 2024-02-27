package main

import (
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
)

const TREE_DEPTH = 32

type IncrementalTree struct {
	LeftDigests [TREE_DEPTH][]byte
	ZeroDigests [TREE_DEPTH][]byte

	RootDigest    []byte
	NextLeafIndex uint64
	MaxLeaves     uint64
}

type Leaf struct {
	Beneficiary         string // address
	ProjectTokenAmount  string // uint256
	TerminalTokenAmount string // uint256
}

var ZERO_DIGESTS [32][]byte // hash values at different heights for a binary tree with leaves equal to 0 (keccak256)

func init() {
	ZERO_DIGESTS[0] = make([]byte, 32)
	for i := 1; i < len(ZERO_DIGESTS); i++ {
		ZERO_DIGESTS[i] = crypto.Keccak256(ZERO_DIGESTS[i-1], ZERO_DIGESTS[i-1])
	}
}

func main() {
	t := &IncrementalTree{
		ZeroDigests: ZERO_DIGESTS,
		LeftDigests: [TREE_DEPTH][]byte{},

		MaxLeaves:     1 << TREE_DEPTH,
		RootDigest:    ZERO_DIGESTS[TREE_DEPTH-1],
		NextLeafIndex: 0,
	}

	fmt.Println(hex.EncodeToString(t.RootDigest))
	fmt.Println("Adding leaves...")

	l1, _ := hex.DecodeString("d84691a17cd171b6bb464b0161f2b0c7773f5b73a9b67962e7fba4f478cb9c80")
	l2, _ := hex.DecodeString("7ee645549845dbdb0a5ca8460b206e0340b07c79674d83ba91013e124c219ffc")
	l3, _ := hex.DecodeString("90fbd9ec3d536c70a2c30ac2f8bdfbc16455819d6d36bdcb4d76efa8049dfe3c")

	t.AddLeaf(l1)
	t.AddLeaf(l2)
	t.AddLeaf(l3)

	fmt.Println(hex.EncodeToString(t.RootDigest))
}

// TODO: Accept a Leaf parameter and hash in here
func (tree *IncrementalTree) AddLeaf(leaf []byte) error {
	if tree.NextLeafIndex >= tree.MaxLeaves {
		return errors.New("tree is full")
	}

	leftRightIndex := tree.NextLeafIndex
	latestDigest := leaf

	// Iterate through the tree from the bottom
	for i := 0; i < TREE_DEPTH; i++ {
		var left, right []byte

		// If the index is even, we're on a left node
		if leftRightIndex%2 == 0 {
			left = latestDigest
			// Right is always the zero digest
			right = tree.ZeroDigests[i]
			tree.LeftDigests[i] = latestDigest
		} else {
			left = tree.LeftDigests[i]
			right = latestDigest
		}

		latestDigest = crypto.Keccak256(left, right)
		// Divide the index by two to traverse up the tree
		leftRightIndex >>= 1 // same as leftRightIndex = leftRightIndex / 2
	}

	tree.RootDigest = latestDigest
	tree.NextLeafIndex++

	return nil
}
