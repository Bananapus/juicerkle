package main

import (
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
)

/*type Leaf struct {
	Beneficiary         string // address
	ProjectTokenAmount  string // uint256
	TerminalTokenAmount string // uint256
}*/

const (
	treeDepth = 32
	maxLeaves = 1<<treeDepth - 1 // 2^TREE_DEPTH - 1
	debugging = true
)

type Tree struct {
	branch [treeDepth][]byte
	count  uint32
}

var zeroDigests [treeDepth][]byte // hash values at different heights for a binary tree with leaves equal to 0 (keccak256)
var z32 []byte                    // The initial root hash of the tree with leaves all equal to 0.

var (
	errTreeIsFull = errors.New("tree is full")
)

func init() {
	zeroDigests[0] = make([]byte, 32)
	for i := 1; i < treeDepth; i++ {
		zeroDigests[i] = crypto.Keccak256(zeroDigests[i-1], zeroDigests[i-1])
	}

	z32 = crypto.Keccak256(zeroDigests[treeDepth-1], zeroDigests[treeDepth-1])
}

func main() {}

// TODO: Accept a Leaf parameter and hash in here
func (tree *Tree) insert(node []byte) error {
	// Increment the tree's count
	tree.count++
	if tree.count > maxLeaves {
		return errTreeIsFull
	}

	size := tree.count

	// Iterate through the tree from the bottom
	for i := 0; i < treeDepth; i++ {
		// If we've hit an odd size, set this index in the branch to the current node and return
		if size&1 == 1 {
			tree.branch[i] = make([]byte, 32)
			copy(tree.branch[i], node)
			return nil
		}

		// If the size isn't odd, hash to move up
		node = crypto.Keccak256(tree.branch[i], node)
		// Divide the size by 2 to move up the tree
		size >>= 1
	}

	return errTreeIsFull
}

func (tree *Tree) root() []byte {
	index := tree.count
	if index == 0 {
		return z32
	}

	current := make([]byte, 32)
	i := 0
	for ; i < treeDepth; i++ {
		if index&(1<<i) == 1 {
			current = crypto.Keccak256(tree.branch[i], zeroDigests[i])
			break
		}
	}

	if i == treeDepth {
		current = z32
	}

	if i > 30 {
		return current
	}

	for ; i < treeDepth-1; i++ {
		if index&(1<<(i+1)) == 0 {
			// Combine with the pre-defined zero hash because the sibling is an empty node.
			current = crypto.Keccak256(current, zeroDigests[i+1])
		} else {
			// Combine with the next non-empty node at this level.
			current = crypto.Keccak256(tree.branch[i+1], current)
		}
	}

	return current
}

// TODO: Accept a leaf parameter
func VerifyProof(index uint32, proof [treeDepth][]byte, leaf []byte, expectedRoot []byte) (bool, error) {
	latestDigest := leaf

	for i := 0; i < treeDepth; i++ {
		// If the index is even, we're on a left node
		// This bit math multiplies the index by 2 each iteration (to traverse down the tree) and checks if the index is even.
		if index>>i&1 == 0 {
			if debugging {
				fmt.Print(
					"0: hashed ",
					hex.EncodeToString(latestDigest),
					" with ",
					hex.EncodeToString(proof[i]),
					" to get ",
				)
			}

			latestDigest = crypto.Keccak256(latestDigest, proof[i])

			if debugging {
				fmt.Print(hex.EncodeToString(latestDigest), "\n")
			}
		} else {
			if debugging {
				fmt.Print(
					"1: hashed ",
					hex.EncodeToString(proof[i]),
					" with ",
					hex.EncodeToString(latestDigest),
					" to get ",
				)
			}

			latestDigest = crypto.Keccak256(proof[i], latestDigest)

			if debugging {
				fmt.Print(hex.EncodeToString(latestDigest), "\n")
			}
		}

	}

	if debugging {
		fmt.Println("Expected: ", hex.EncodeToString(expectedRoot))
		fmt.Println("Got:", hex.EncodeToString(latestDigest))
	}

	if hex.EncodeToString(latestDigest) != hex.EncodeToString(expectedRoot) {
		if debugging {
			fmt.Println("DID NOT MATCH")
		}
		return false, errors.New("proof did not match expected root")
	}
	return true, nil
}

func (tree *Tree) GetProof(index uint64) (proof [treeDepth][]byte, err error) {

	return
}
