package header_test

import (
	"reflect"
	"testing"

	binarytree "github.com/KylerWilson01/huffman-compression/internal/binary-tree"
	"github.com/KylerWilson01/huffman-compression/internal/header"
)

func TestCreatePrefixFromTree(t *testing.T) {
	tests := []struct {
		name  string
		input struct {
			ht binarytree.HuffTree
			ft map[rune]int
		}
		want []header.PrefixNode
	}{
		{
			name: "pass",
			input: struct {
				ht binarytree.HuffTree
				ft map[rune]int
			}{
				ht: binarytree.HuffTree{
					Root: binarytree.InternalNode{
						Weight: 12,
						L: binarytree.LeafNode{
							Character: 't',
							Weight:    5,
						},
						R: binarytree.InternalNode{
							Weight: 7,
							L: binarytree.InternalNode{
								Weight: 3,
								L: binarytree.LeafNode{
									Weight:    1,
									Character: 'h',
								},
								R: binarytree.LeafNode{
									Weight:    2,
									Character: 'i',
								},
							},
							R: binarytree.LeafNode{
								Weight:    4,
								Character: 's',
							},
						},
					},
				},
				ft: map[rune]int{
					't': 5,
					'h': 1,
					'i': 2,
					's': 4,
				},
			},
			want: []header.PrefixNode{
				{Character: 't', Code: []int{0b0}, Frequency: 5},
				{Character: 'h', Code: []int{0b1, 0b0, 0b0}, Frequency: 1},
				{Character: 'i', Code: []int{0b1, 0b0, 0b1}, Frequency: 2},
				{Character: 's', Code: []int{0b1, 0b1}, Frequency: 4},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			pt, err := header.CreatePrefixTableFromTree(&test.input.ht, test.input.ft)
			if err != nil {
				t.Fatalf("Something happened: %v", err)
			}

			if !reflect.DeepEqual(pt, test.want) {
				t.Fatalf("Got: %v \n\t Want: %v", pt, test.want)
			}
		})
	}
}
