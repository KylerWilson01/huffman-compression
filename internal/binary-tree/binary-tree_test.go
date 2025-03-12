package binarytree_test

import (
	"testing"

	binarytree "github.com/KylerWilson01/huffman-compression/internal/binary-tree"
)

func TestCreateBinTreeFromMap(t *testing.T) {
	tests := []struct {
		name  string
		input map[rune]int
		want  binarytree.HuffTree
	}{
		{
			name: "pass with symbols", input: map[rune]int{
				't': 4,
				'h': 1,
				'i': 2,
				's': 3,
			},
			want: binarytree.HuffTree{
				Root: binarytree.InternalNode{
					Weight: 10,
					Left:   nil,
					Right:  nil,
				},
			},
		},
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
