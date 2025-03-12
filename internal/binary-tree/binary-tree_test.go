package binarytree_test

import (
	"testing"

	binarytree "github.com/KylerWilson01/huffman-compression/internal/binary-tree"
)

func TestCreateBinTreeFromMap(t *testing.T) {
	tests := []struct {
		name  string
		input map[rune]int
	}{
		{name: "pass with symbols", input: map[rune]int{
			't': 3,
			'h': 1,
			'i': 2,
			's': 3,
			'a': 1,
			'e': 1,
			'!': 2,
			',': 2,
			'"': 1,
			';': 1,
			'.': 3,
		}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var totalWeight int
			for _, wt := range test.input {
				totalWeight += wt
			}

			bt := binarytree.CreateBinaryTreeFromMap(test.input)

			if bt.Weight() != totalWeight {
				t.Fatalf("Got %v, wanted %v\n", bt.Weight(), totalWeight)
			}
		})
	}
}
