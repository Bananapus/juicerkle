package main

import (
	"encoding/hex"
	"errors"

	"github.com/ethereum/go-ethereum/crypto"
)

const (
	treeDepth = 32
	maxLeaves = 1<<treeDepth - 1 // 2^TREE_DEPTH - 1
	debugging = true
)

type Tree struct {
	leaves [][]byte // The leaves, in order of insertion.
	count  int      // The current number of leaves in the tree.
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

func (tree *Tree) insert(leaf []byte) error {
	if tree.count >= maxLeaves {
		return errTreeIsFull
	}

	tree.leaves = append(tree.leaves, leaf)
	tree.count++
	return nil
}

// Calculate the root hash of the tree
func (tree *Tree) root() []byte {
	if tree.count == 0 {
		return z32
	}

	i := tree.nonZeroDepth()
	current := tree.subtreeRoot(i, 0)

	// For the rest of the tree, hash the current node with sibling zero hashes until we reach the root
	for ; i < treeDepth; i++ {
		current = crypto.Keccak256(current, zeroDigests[i])
	}

	return current
}

// Find log2 of the tree's count. This is the depth of the non-zero subtree.
func (tree *Tree) nonZeroDepth() int {
	i, n := 0, tree.count
	for n != 0 {
		n >>= 1
		i++
	}
	return i
}

// Get the root of the subtree at the given depth and starting index.
// depth is the depth of the subtree root (0 is the bottom of the tree, and 31 is the top).
// startingIndex is the index of the first leaf in the subtree (on the left at depth 0).
func (tree *Tree) subtreeRoot(depth, startingIndex int) []byte {
	// If we would start outside the defined sub-tree, return the zero hash for that depth.
	if startingIndex > tree.count {
		return zeroDigests[depth]
	}

	leavesToHash := 1 << depth // 2 to the power of depth
	toHash := tree.leaves[startingIndex:min(startingIndex+leavesToHash, tree.count)]

	// If there's nothing to hash, return the zero hash for that depth.
	if len(toHash) == 0 {
		return zeroDigests[depth]
	}

	// Iteratively hash up the subtree
	for subtreeDepth := 0; subtreeDepth < depth; subtreeDepth++ {
		// If we've reached the top of the subtree, break
		if leavesToHash == 1 {
			break
		}

		// Divide the number of leaves to hash by 2 as we move up the tree
		leavesToHash >>= 1
		nextLayer := make([][]byte, (len(toHash)+1)/2) // Use half of len(toHash) to skip zero hashes

		for i := 0; i < len(nextLayer); i++ {
			// We don't need to check if i*2 >= len(toHash) because we're using half of len(toHash).
			if i*2+1 >= len(toHash) {
				// If we go outside the bounds of toHash, hash with the appropriate zero digest.
				nextLayer[i] = crypto.Keccak256(toHash[i*2], zeroDigests[subtreeDepth])
			} else {
				nextLayer[i] = crypto.Keccak256(toHash[i*2], toHash[i*2+1])
			}
		}

		toHash = nextLayer
	}

	return toHash[0]
}

func (tree *Tree) getProof(index int) (proof [][]byte, err error) {
	i := tree.nonZeroDepth()

	if i > 31 {
		return nil, errTreeIsFull
	}

	proof = make([][]byte, treeDepth)
	// Copy the zero hashes into the proof. All siblings above the non-zero subtree are zero hashes.
	copy(proof[i:], zeroDigests[i:])

	// Find siblings at remaining depths moving up from the bottom of the tree
	for depth := 0; depth < i; depth++ {
		startingIndex := (index/(1<<depth) ^ 1) * (1 << depth) // starting index of the leaves to hash
		proof[depth] = tree.subtreeRoot(depth, startingIndex)
	}

	return
}

func VerifyProof(index int, proof [][]byte, leaf []byte, expectedRoot []byte) (bool, error) {
	latestDigest := leaf

	for i := 0; i < treeDepth; i++ {
		// If the index is even, we're on a left node
		// This bit math multiplies the index by 2 each iteration (to traverse down the tree) and checks if the index is even.
		if index>>i&1 == 0 {
			latestDigest = crypto.Keccak256(latestDigest, proof[i])
		} else {
			latestDigest = crypto.Keccak256(proof[i], latestDigest)
		}
	}

	if hex.EncodeToString(latestDigest) != hex.EncodeToString(expectedRoot) {
		return false, errors.New("proof did not match expected root")
	}

	return true, nil
}
