package main

import (
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
)

const TREE_DEPTH = 32
const MAX_LEAVES = 1<<TREE_DEPTH - 1

type IncrementalTree struct {
	LeftDigests [TREE_DEPTH][]byte
	ZeroDigests [TREE_DEPTH][]byte

	RootDigest    []byte
	NextLeafIndex uint32
}

type Leaf struct {
	Beneficiary         string // address
	ProjectTokenAmount  string // uint256
	TerminalTokenAmount string // uint256
}

var ZERO_DIGESTS [TREE_DEPTH][]byte // hash values at different heights for a binary tree with leaves equal to 0 (keccak256)

func init() {
	ZERO_DIGESTS[0] = make([]byte, 32)
	for i := 1; i < TREE_DEPTH; i++ {
		ZERO_DIGESTS[i] = crypto.Keccak256(ZERO_DIGESTS[i-1], ZERO_DIGESTS[i-1])
	}
}

func main() {
	/* myInt := 368
	fmt.Println()
	for i := 0; i < TREE_DEPTH; i++ {
		fmt.Print(myInt >> i & 1)
	}
	fmt.Println()
	for i := 0; i < TREE_DEPTH; i++ {
		fmt.Print(myInt >> (TREE_DEPTH - 1 - i) & 1)
	}
	os.Exit(0)*/

	t := &IncrementalTree{
		ZeroDigests: ZERO_DIGESTS,
		LeftDigests: [TREE_DEPTH][]byte{},

		RootDigest:    ZERO_DIGESTS[TREE_DEPTH-1],
		NextLeafIndex: 0,
	}

	// fmt.Println(hex.EncodeToString(t.RootDigest))
	fmt.Println("Adding leaves...")

	// l, _ := hex.DecodeString("f4cff1055989ad136597d6d081b574c479f5434e04f81af6f6009c7f4c84fc7f")
	// t.AddLeaf(l)
	// fmt.Println(hex.EncodeToString(t.RootDigest))

	l1, _ := hex.DecodeString("d84691a17cd171b6bb464b0161f2b0c7773f5b73a9b67962e7fba4f478cb9c80")
	l2, _ := hex.DecodeString("7ee645549845dbdb0a5ca8460b206e0340b07c79674d83ba91013e124c219ffc")
	l3, _ := hex.DecodeString("90fbd9ec3d536c70a2c30ac2f8bdfbc16455819d6d36bdcb4d76efa8049dfe3c")

	t.AddLeaf(l1)
	t.AddLeaf(l2)
	t.AddLeaf(l3)

	// fmt.Println(hex.EncodeToString(t.RootDigest))

	var proof [TREE_DEPTH][]byte
	proof[0], _ = hex.DecodeString("0000000000000000000000000000000000000000000000000000000000000000")
	proof[1], _ = hex.DecodeString("15d682413b76d69d3a9f37321d80938e54900b74e3f119c027ed87dcec1a935b")
	proof[2], _ = hex.DecodeString("b4c11951957c6f8f642c4af61cd6b24640fec6dc7fc607ee8206a99e92410d30")
	proof[3], _ = hex.DecodeString("21ddb9a356815c3fac1026b6dec5df3124afbadb485c9ba5a3e3398a04b7ba85")
	proof[4], _ = hex.DecodeString("e58769b32a1beaf1ea27375a44095a0d1fb664ce2dd358e7fcbfb78c26a19344")
	proof[5], _ = hex.DecodeString("0eb01ebfc9ed27500cd4dfc979272d1f0913cc9f66540d7e8005811109e1cf2d")
	proof[6], _ = hex.DecodeString("887c22bd8750d34016ac3c66b5ff102dacdd73f6b014e710b51e8022af9a1968")
	proof[7], _ = hex.DecodeString("ffd70157e48063fc33c97a050f7f640233bf646cc98d9524c6b92bcf3ab56f83")
	proof[8], _ = hex.DecodeString("9867cc5f7f196b93bae1e27e6320742445d290f2263827498b54fec539f756af")
	proof[9], _ = hex.DecodeString("cefad4e508c098b9a7e1d8feb19955fb02ba9675585078710969d3440f5054e0")
	proof[10], _ = hex.DecodeString("f9dc3e7fe016e050eff260334f18a5d4fe391d82092319f5964f2e2eb7c1c3a5")
	proof[11], _ = hex.DecodeString("f8b13a49e282f609c317a833fb8d976d11517c571d1221a265d25af778ecf892")
	proof[12], _ = hex.DecodeString("3490c6ceeb450aecdc82e28293031d10c7d73bf85e57bf041a97360aa2c5d99c")
	proof[13], _ = hex.DecodeString("c1df82d9c4b87413eae2ef048f94b4d3554cea73d92b0f7af96e0271c691e2bb")
	proof[14], _ = hex.DecodeString("5c67add7c6caf302256adedf7ab114da0acfe870d449a3a489f781d659e8becc")
	proof[15], _ = hex.DecodeString("da7bce9f4e8618b6bd2f4132ce798cdc7a60e7e1460a7299e3c6342a579626d2")
	proof[16], _ = hex.DecodeString("2733e50f526ec2fa19a22b31e8ed50f23cd1fdf94c9154ed3a7609a2f1ff981f")
	proof[17], _ = hex.DecodeString("e1d3b5c807b281e4683cc6d6315cf95b9ade8641defcb32372f1c126e398ef7a")
	proof[18], _ = hex.DecodeString("5a2dce0a8a7f68bb74560f8f71837c2c2ebbcbf7fffb42ae1896f13f7c7479a0")
	proof[19], _ = hex.DecodeString("b46a28b6f55540f89444f63de0378e3d121be09e06cc9ded1c20e65876d36aa0")
	proof[20], _ = hex.DecodeString("c65e9645644786b620e2dd2ad648ddfcbf4a7e5b1a3a4ecfe7f64667a3f0b7e2")
	proof[21], _ = hex.DecodeString("f4418588ed35a2458cffeb39b93d26f18d2ab13bdce6aee58e7b99359ec2dfd9")
	proof[22], _ = hex.DecodeString("5a9c16dc00d6ef18b7933a6f8dc65ccb55667138776f7dea101070dc8796e377")
	proof[23], _ = hex.DecodeString("4df84f40ae0c8229d0d6069e5c8f39a7c299677a09d367fc7b05e3bc380ee652")
	proof[24], _ = hex.DecodeString("cdc72595f74c7b1043d0e1ffbab734648c838dfb0527d971b602bc216c9619ef")
	proof[25], _ = hex.DecodeString("0abf5ac974a1ed57f4050aa510dd9c74f508277b39d7973bb2dfccc5eeb0618d")
	proof[26], _ = hex.DecodeString("b8cd74046ff337f0a7bf2c8e03e10f642c1886798d71806ab1e888d9e5ee87d0")
	proof[27], _ = hex.DecodeString("838c5655cb21c6cb83313b5a631175dff4963772cce9108188b34ac87c81c41e")
	proof[28], _ = hex.DecodeString("662ee4dd2dd7b2bc707961b1e646c4047669dcb6584f0d8d770daf5d7e7deb2e")
	proof[29], _ = hex.DecodeString("388ab20e2573d171a88108e79d820e98f26c0b84aa8b2f4aa4968dbb818ea322")
	proof[30], _ = hex.DecodeString("93237c50ba75ee485f4c22adf2f741400bdf8d6a9cc7df7ecae576221665d735")
	proof[31], _ = hex.DecodeString("8448818bb4ae4562849e949e17ac16e0be16688e156b5cf15e098c627c0056a9")

	fmt.Println("Expected: ", hex.EncodeToString(t.RootDigest))
	VerifyProof(1, proof, l1, t.RootDigest)
}

// TODO: Accept a Leaf parameter and hash in here
func (tree *IncrementalTree) AddLeaf(leaf []byte) error {
	if tree.NextLeafIndex >= MAX_LEAVES-1 {
		return errors.New("tree is full")
	}

	latestDigest := leaf

	// Iterate through the tree from the bottom
	for i := 0; i < TREE_DEPTH; i++ {
		var left, right []byte

		// If the index is even, we're on a left node
		// This bit math is equivalent to dividing the index by 2 each iteration (to traverse up the tree) and checking if it is even.
		if tree.NextLeafIndex>>(TREE_DEPTH-1-i)&1 == 0 {
			left = latestDigest
			// Right is always the zero digest
			right = tree.ZeroDigests[i]
			tree.LeftDigests[i] = latestDigest
		} else {
			left = tree.LeftDigests[i]
			right = latestDigest
		}

		latestDigest = crypto.Keccak256(left, right)
	}

	tree.RootDigest = latestDigest
	tree.NextLeafIndex++

	return nil
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

	fmt.Println("Got:", hex.EncodeToString(latestDigest))
	return true, nil
}

func (tree *IncrementalTree) GetProof(index uint64) (proof [TREE_DEPTH][]byte, err error) {

	return
}
