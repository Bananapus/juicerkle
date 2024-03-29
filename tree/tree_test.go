package tree

import (
	"encoding/hex"
	"testing"
)

// Test that the starting leaf for subtree hashing is calculated correctly.
func TestStartingIndex(t *testing.T) {
	var tests = []struct {
		index                   uint
		expectedStartingIndices []uint
	}{
		{
			index:                   0,
			expectedStartingIndices: []uint{1, 2, 4, 8, 16, 32},
		},
		{
			index:                   1,
			expectedStartingIndices: []uint{0, 2, 4, 8, 16, 32},
		},
		{
			index:                   5,
			expectedStartingIndices: []uint{4, 6, 0, 8, 16, 32},
		},
		{
			index:                   16,
			expectedStartingIndices: []uint{17, 18, 20, 24, 0, 32},
		},
	}

	for _, test := range tests {
		for depth := 0; depth < len(test.expectedStartingIndices); depth++ {
			// The line being tested.
			startingIndex := uint((test.index/(1<<depth) ^ 1) * (1 << depth))
			if startingIndex != test.expectedStartingIndices[depth] {
				t.Errorf("For index %v, expected starting index %v, but got %v", test.index, test.expectedStartingIndices[depth], startingIndex)
			}
		}
	}
}

func TestProof(t *testing.T) {
	seventeenLeaves := []string{
		"8493eba7305be4ad7134adaf22748ebcc0db40056d78786e94ed35bcdeb3064d",
		"c14b723895c8e3c01cac0c4a51bafa968307c2c7fd61293cbbc65831d3de6219",
		"1bf1281a36f29f2c6ab1b7945420b78755afec33465ed1eda95bc8df484d5768",
		"183d2d3288f4c1511815f1a7be8239ccaeab2425720d3bf3c47c2165ac02c690",
		"490b1f43511f199d6a718c9387f9118f14b2dbcd575b6955711f176ce6e8c23a",
		"0b4c34478f59cf1059d40a95105994e1b5e626281faa7269f5567442d2ae02cd",
		"fb9f903700176f64000350b5aa73edfe05c49103f239fafb340ac5ffaea0cdc9",
		"962983f79d6b628524e6f721e3d72a0772e78c162ded0b017e37ff62ea15be91",
		"9b10b56115c8c303733c5e51a3b7548b1967a724b09146e635de0b5ebb1d450b",
		"be1a818599052df1eb800c78e771ce20dd260f2427a0c41211beb5331fd2c5db",
		"553e4c210e837b97d877aff12f63960039cd3dcccda5f723ec1749a7eac85ae8",
		"853366273b58270c3634ebe835bef5426e8087f529a0580c8ad3035c139c43f5",
		"ff4ee05878bff191b2097f1d5d9c49ac93f97724b217f55d1a5c922f0e255c5b",
		"8ded287963d1da7fcfb844b25730eeee6d3c4048b891dec236bb5d55e1eb74dd",
		"8bb515d204cedc015286e6458fc8e18627c11b4f15a10531485e23694bbe2bd4",
		"4fe58ff93a8b25a92ea990fce374ca47f63ceab01451fa65b9c98ac6f135a668",
		"a4eca02fd3701a051cd0c1494bbbe37d165405f325d408ad01f02f40573ca0cc",
	}

	var tests = []struct {
		leaves        []string
		index         uint
		expectedProof []string
		expectedRoot  string
	}{
		{
			leaves:       seventeenLeaves,
			index:        4,
			expectedRoot: "d427830563d6c336971d6ce4bb12cf5710b1fe9fafda7cca01b9e33fdb5287ca",
			expectedProof: []string{
				"0b4c34478f59cf1059d40a95105994e1b5e626281faa7269f5567442d2ae02cd", "9f279ac5daa089e3d343081e56f944a31ec39ce6d681f2cb4df79a0bee6541d3",
				"8db8d069533c556ac2e8defd7e1bcd861c35d8582b57edb1e4bf833afae903c5", "71af814b71bf6cdd1d7e24eee976d988ebca6ccb0667c3b23b56cef51b873443",
				"44aadd27fe442ff7988264938e4f950bdc138668d90e46eed78af041155db2de", "0eb01ebfc9ed27500cd4dfc979272d1f0913cc9f66540d7e8005811109e1cf2d",
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
				"93237c50ba75ee485f4c22adf2f741400bdf8d6a9cc7df7ecae576221665d735", "8448818bb4ae4562849e949e17ac16e0be16688e156b5cf15e098c627c0056a9",
			},
		},
		{
			leaves:       seventeenLeaves,
			index:        16,
			expectedRoot: "d427830563d6c336971d6ce4bb12cf5710b1fe9fafda7cca01b9e33fdb5287ca",
			expectedProof: []string{
				"0000000000000000000000000000000000000000000000000000000000000000", "ad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5",
				"b4c11951957c6f8f642c4af61cd6b24640fec6dc7fc607ee8206a99e92410d30", "21ddb9a356815c3fac1026b6dec5df3124afbadb485c9ba5a3e3398a04b7ba85",
				"7226d1796d5fb58c2a5f8ed5b9a89953976cb1523d5421e2675d5382bdb53b2c", "0eb01ebfc9ed27500cd4dfc979272d1f0913cc9f66540d7e8005811109e1cf2d",
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
				"93237c50ba75ee485f4c22adf2f741400bdf8d6a9cc7df7ecae576221665d735", "8448818bb4ae4562849e949e17ac16e0be16688e156b5cf15e098c627c0056a9",
			},
		},
		{
			leaves:       seventeenLeaves,
			index:        7,
			expectedRoot: "d427830563d6c336971d6ce4bb12cf5710b1fe9fafda7cca01b9e33fdb5287ca",
			expectedProof: []string{
				"fb9f903700176f64000350b5aa73edfe05c49103f239fafb340ac5ffaea0cdc9", "a7422347d37dbc23d05066192db9ab94db10a50c0fc41679524a619eb29fb6bf",
				"8db8d069533c556ac2e8defd7e1bcd861c35d8582b57edb1e4bf833afae903c5", "71af814b71bf6cdd1d7e24eee976d988ebca6ccb0667c3b23b56cef51b873443",
				"44aadd27fe442ff7988264938e4f950bdc138668d90e46eed78af041155db2de", "0eb01ebfc9ed27500cd4dfc979272d1f0913cc9f66540d7e8005811109e1cf2d",
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
				"93237c50ba75ee485f4c22adf2f741400bdf8d6a9cc7df7ecae576221665d735", "8448818bb4ae4562849e949e17ac16e0be16688e156b5cf15e098c627c0056a9",
			},
		},
		{
			leaves: []string{"d84691a17cd171b6bb464b0161f2b0c7773f5b73a9b67962e7fba4f478cb9c80", "7ee645549845dbdb0a5ca8460b206e0340b07c79674d83ba91013e124c219ffc", "90fbd9ec3d536c70a2c30ac2f8bdfbc16455819d6d36bdcb4d76efa8049dfe3c"},
			index:  1,
			expectedProof: []string{
				"d84691a17cd171b6bb464b0161f2b0c7773f5b73a9b67962e7fba4f478cb9c80", "2d1fa8e18e0e2f8fee037dd3d52daa2ff9e61aefe194e5b0877b7c9e0e8c8817",
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
				"93237c50ba75ee485f4c22adf2f741400bdf8d6a9cc7df7ecae576221665d735", "8448818bb4ae4562849e949e17ac16e0be16688e156b5cf15e098c627c0056a9",
			},
			expectedRoot: "70352a98d1692a499bcf052869fa679b7862630764f5e4a856ea5dc7e0c7f99c",
		},
	}

	for _, test := range tests {
		t.Logf("Testing proof for index %v, expected root %s", test.index, test.expectedRoot)

		if len(test.expectedProof) != 32 {
			t.Errorf("Expected proof length is not 32")
		}

		// Set up the tree
		tree := Tree{}
		for _, l := range test.leaves {
			decoded, err := hex.DecodeString(l)
			if err != nil {
				t.Errorf("Error decoding hex: %v", err)
			}
			tree.Insert(decoded)
		}

		// Check the root
		root := hex.EncodeToString(tree.Root())
		if root != test.expectedRoot {
			t.Errorf("Root did not match expected root (%s != %s)", root, test.expectedRoot)
		}

		// Check the proof
		proof, err := tree.Proof(test.index)
		if err != nil {
			t.Errorf("Error getting proof: %v", err)
		}

		for i, p := range proof {
			if hex.EncodeToString(p) != test.expectedProof[i] {
				t.Errorf("Proof %d did not match expected proof (%v != %v)", i, hex.EncodeToString(p), test.expectedProof[i])
			}
		}

		// Verify the proof
		expectedRoot, err := hex.DecodeString(test.expectedRoot)
		if err != nil {
			t.Errorf("Error decoding expected root: %v", err)
		}

		verified, err := VerifyProof(test.index, proof, tree.leaves[test.index], expectedRoot)
		if err != nil {
			t.Errorf("Error verifying proof: %v", err)
		}

		t.Logf("Proof verified: %v", verified)
	}
}
