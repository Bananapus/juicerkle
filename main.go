package main

import (
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

func (tree *Tree) getProof(index int) (proof [][]byte, err error) {
	// Find log2 of the tree size. This is the depth of the non-zero subtree.
	i, n := 0, tree.count
	for n != 0 {
		n >>= 1
		i++
	}

	if i > 31 {
		return nil, errTreeIsFull
	}

	proof = make([][]byte, treeDepth)
	// Copy the zero hashes into the proof. All siblings above the non-zero subtree are zero hashes.
	copy(proof, zeroDigests[i:])

	// Find siblings at remaining depths moving up from the bottom of the tree
	for depth := 0; depth < i; depth++ {
		leavesToHash := 1 << depth            // 2 to the power of depth
		startingIndex := index | (1 << depth) // starting index of the leaves to hash

		// If we're outside the non-zero subtree, we can use the zero hash.
		if startingIndex >= tree.count>>depth { // Bit shift equivalent to dividing by 2^depth
			proof[depth] = zeroDigests[depth]
			continue
		}

		// Get the leaves to hash from the tree
		toHash := tree.leaves[startingIndex:min(startingIndex+leavesToHash, tree.count)]

		// Iteratively hash up the subtree
		for subtreeDepth := 0; subtreeDepth < depth; subtreeDepth++ {
			if leavesToHash == 1 {
				break
			}

			// Divide the number of leaves to hash by 2 as we move up the tree
			leavesToHash >>= 1
			nextLayer := make([][]byte, (len(toHash)+1)/2) // Use half of len(toHash) to skip zero hashes

			// Use the zero hashes if we've reached the end of the defined leaves
			for i := 0; i < len(nextLayer); i++ {
				// We don't need to check if i*2 >= len(toHash) because we're using half of len(toHash).
				// If we go outside the bounds of toHash, hash with the appropriate zero digest.
				if i*2+1 >= len(toHash) {
					nextLayer[i] = crypto.Keccak256(toHash[i*2], zeroDigests[subtreeDepth])
				} else {
					nextLayer[i] = crypto.Keccak256(toHash[i*2], toHash[i*2+1])
				}
			}

			toHash = nextLayer
		}

		// There should only be one hash left in toHash
		proof[depth] = toHash[0]
	}

	return
}

func (tree *Tree) subtreeHash(layer, index int) []byte {
	if layer == 0 {
		return tree.leaves[index]
	}

	return nil
}
