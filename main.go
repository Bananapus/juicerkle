package main

import (
	"encoding/hex"
	"errors"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/crypto"
)

/*type Leaf struct {
	Beneficiary         string // address
	ProjectTokenAmount  string // uint256
	TerminalTokenAmount string // uint256
}*/

const TREE_DEPTH = 32
const MAX_LEAVES = 1<<TREE_DEPTH - 1 // 2^TREE_DEPTH - 1

type Tree struct {
	branch [TREE_DEPTH][]byte
	count  uint32
}

var ZERO_DIGESTS [TREE_DEPTH][]byte // hash values at different heights for a binary tree with leaves equal to 0 (keccak256)
var Z_32 []byte                     // The initial root hash of the tree with leaves all equal to 0.

var (
	treeIsFullError = errors.New("tree is full")
)

func init() {
	ZERO_DIGESTS[0] = make([]byte, 32)
	for i := 1; i < TREE_DEPTH; i++ {
		ZERO_DIGESTS[i] = crypto.Keccak256(ZERO_DIGESTS[i-1], ZERO_DIGESTS[i-1])
	}

	Z_32 = crypto.Keccak256(ZERO_DIGESTS[TREE_DEPTH-1], ZERO_DIGESTS[TREE_DEPTH-1])
}

func main() {
	t := Tree{}

	l, _ := hex.DecodeString("f4cff1055989ad136597d6d081b574c479f5434e04f81af6f6009c7f4c84fc7f")
	t.insert(l)

	fmt.Println("New root:", hex.EncodeToString(t.root()))
	os.Exit(0)

	// fmt.Println(hex.EncodeToString(t.RootDigest))
	fmt.Println("Adding leaves...")

	// l, _ := hex.DecodeString("f4cff1055989ad136597d6d081b574c479f5434e04f81af6f6009c7f4c84fc7f")
	// t.AddLeaf(l)
	// fmt.Println(hex.EncodeToString(t.RootDigest))

	l1, _ := hex.DecodeString("d84691a17cd171b6bb464b0161f2b0c7773f5b73a9b67962e7fba4f478cb9c80")
	l2, _ := hex.DecodeString("7ee645549845dbdb0a5ca8460b206e0340b07c79674d83ba91013e124c219ffc")
	l3, _ := hex.DecodeString("90fbd9ec3d536c70a2c30ac2f8bdfbc16455819d6d36bdcb4d76efa8049dfe3c")

	t.insert(l1)
	t.insert(l2)
	t.insert(l3)

	// fmt.Println(hex.EncodeToString(t.RootDigest))

	proofStrs := []string{"d84691a17cd171b6bb464b0161f2b0c7773f5b73a9b67962e7fba4f478cb9c80", "2d1fa8e18e0e2f8fee037dd3d52daa2ff9e61aefe194e5b0877b7c9e0e8c8817",
		"b4c11951957c6f8f642c4af61cd6b24640fec6dc7fc607ee8206a99e92410d30", "21ddb9a356815c3fac1026b6dec5df3124afbadb485c9ba5a3e3398a04b7ba85",
		"e58769b32a1beaf1ea27375a44095a0d1fb664ce2dd358e7fcbfb78c26a19344", "0eb01ebfc9ed27500cd4dfc979272d1f0913cc9f66540d7e8005811109e1cf2d",
		"887c22bd8750d34016ac3c66b5ff102dacdd73f6b014e710b51e8022af9a1968", "ffd70157e48063fc33c97a050f7f640233bf646cc98d9524c6b92bcf3ab56f83",
		"9867cc5f7f196b93bae1e27e6320742445d290f2263827498b54fec539f756af", "cefad4e508c098b9a7e1d8feb19955fb02ba9675585078710969d3440f5054e0",
		"f9dc3e7fe016e050eff260334f18a5d4fe391d82092319f5964f2e2eb7c1c3a5", "f8b13a49e282f609c317a833fb8d976d11517c571d1221a265d25af778ecf892",
		"3490c6ceeb450aecdc82e28293031d10c7d73bf85e57bf041a97360aa2c5d99c", "c1df82d9c4b87413eae2ef048f94b4d3554cea73d92b0f7af96e0271c691e2bb",
		"5c67add7c6caf302256adedf7ab114da0acfe870d449a3a489f781d659e8becc", "da7bce9f4e8618b6bd2f4132ce798cdc7a60e7e1460a7299e3c6342a579626d2",
		"2733e50f526ec2fa19a22b31e8ed50f23cd1fdf94c9154ed3a7609a2f1ff981f", "e1d3b5c807b281e4683cc6d6315cf95b9ade8641defcb32372f1c126e398ef7a",
		"5a2dce0a8a7f68bb74560f8f71837c2c2ebbcbf7fffb42ae1896f13f7c7479a0", "b46a28b6f55540f89444f63de0378e3d121be09e06cc9ded1c20e65876d36aa0",
		"c65e9645644786b620e2dd2ad648ddfcbf4a7e5b1a3a4ecfe7f64667a3f0b7e2", "f4418588ed35a2458cffeb39b93d26f18d2ab13bdce6aee58e7b99359ec2dfd9",
		"5a9c16dc00d6ef18b7933a6f8dc65ccb55667138776f7dea101070dc8796e377", "4df84f40ae0c8229d0d6069e5c8f39a7c299677a09d367fc7b05e3bc380ee652",
		"cdc72595f74c7b1043d0e1ffbab734648c838dfb0527d971b602bc216c9619ef", "0abf5ac974a1ed57f4050aa510dd9c74f508277b39d7973bb2dfccc5eeb0618d",
		"b8cd74046ff337f0a7bf2c8e03e10f642c1886798d71806ab1e888d9e5ee87d0", "838c5655cb21c6cb83313b5a631175dff4963772cce9108188b34ac87c81c41e",
		"662ee4dd2dd7b2bc707961b1e646c4047669dcb6584f0d8d770daf5d7e7deb2e", "388ab20e2573d171a88108e79d820e98f26c0b84aa8b2f4aa4968dbb818ea322",
		"93237c50ba75ee485f4c22adf2f741400bdf8d6a9cc7df7ecae576221665d735", "8448818bb4ae4562849e949e17ac16e0be16688e156b5cf15e098c627c0056a9"}

	var proof [TREE_DEPTH][]byte
	for i, s := range proofStrs {
		proof[i], _ = hex.DecodeString(s)
	}
	VerifyProof(1, proof, l2, l3)
}

// TODO: Accept a Leaf parameter and hash in here
func (tree *Tree) insert(node []byte) error {
	// Increment the tree's count
	tree.count++
	if tree.count > MAX_LEAVES {
		return treeIsFullError
	}

	size := tree.count

	// Iterate through the tree from the bottom
	for i := 0; i < TREE_DEPTH; i++ {
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

	return treeIsFullError
}

func (tree *Tree) root() []byte {
	index := tree.count
	if index == 0 {
		return Z_32
	}

	current := make([]byte, 32)
	i := 0
	for ; i < TREE_DEPTH; i++ {
		if index&(1<<i) == 1 {
			current = crypto.Keccak256(tree.branch[i], ZERO_DIGESTS[i])
			break
		}
	}

	if i == TREE_DEPTH {
		current = Z_32
	}

	if i > 30 {
		return current
	}

	for ; i < TREE_DEPTH-1; i++ {
		if index&(1<<(i+1)) == 0 {
			// Combine with the pre-defined zero hash because the sibling is an empty node.
			current = crypto.Keccak256(current, ZERO_DIGESTS[i+1])
		} else {
			// Combine with the next non-empty node at this level.
			current = crypto.Keccak256(tree.branch[i+1], current)
		}
	}

	return current
}

// TODO: Accept a leaf parameter
func VerifyProof(index uint32, proof [TREE_DEPTH][]byte, leaf []byte, expectedRoot []byte) (bool, error) {
	latestDigest := leaf

	for i := 0; i < TREE_DEPTH; i++ {
		// If the index is even, we're on a left node
		// This bit math multiplies the index by 2 each iteration (to traverse down the tree) and checks if the index is even.
		if index>>i&1 == 0 {
			fmt.Print(
				"0: hashed ",
				hex.EncodeToString(latestDigest),
				" with ",
				hex.EncodeToString(proof[i]),
				" to get ",
			)
			latestDigest = crypto.Keccak256(latestDigest, proof[i])
			fmt.Print(hex.EncodeToString(latestDigest), "\n")
		} else {
			fmt.Print(
				"1: hashed ",
				hex.EncodeToString(proof[i]),
				" with ",
				hex.EncodeToString(latestDigest),
				" to get ",
			)
			latestDigest = crypto.Keccak256(proof[i], latestDigest)
			fmt.Print(hex.EncodeToString(latestDigest), "\n")
		}

	}

	fmt.Println("Expected: ", hex.EncodeToString(expectedRoot))
	fmt.Println("Got:", hex.EncodeToString(latestDigest))
	if hex.EncodeToString(latestDigest) != hex.EncodeToString(expectedRoot) {
		fmt.Println("DID NOT MATCH")
		return false, errors.New("proof did not match expected root")
	}
	return true, nil
}

func (tree *Tree) GetProof(index uint64) (proof [TREE_DEPTH][]byte, err error) {

	return
}
