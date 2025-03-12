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
				't': 5,
				'h': 1,
				'i': 2,
				's': 4,
			},
			want: binarytree.HuffTree{
				Root: binarytree.InternalNode{
					Weight: 12,
					Left: binarytree.LeafNode{
						Character: 't',
						Weight:    5,
					},
					Right: binarytree.InternalNode{
						Weight: 7,
						Left: binarytree.InternalNode{
							Weight: 3,
							Left: binarytree.LeafNode{
								Weight:    1,
								Character: 'h',
							},
							Right: binarytree.LeafNode{
								Weight:    2,
								Character: 'i',
							},
						},
						Right: binarytree.LeafNode{
							Weight:    4,
							Character: 's',
						},
					},
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

			if !bt.Same(&test.want) {
				t.Fatalf("Not the same trees\n")
			}
		})
	}
}
