// Package header is responible for handling the header for the huffman compression
package header

import (
	"errors"

	binarytree "github.com/KylerWilson01/huffman-compression/internal/binary-tree"
	frequencymap "github.com/KylerWilson01/huffman-compression/internal/frequency-map"
)

// PrefixNode allows us to create a prefix header for encoding the file
type PrefixNode struct {
	Character rune
	Code      []byte
	Frequency int
}

// CreatePrefixTableFromTree creates a prefix table from the binary tree that was previously created
func CreatePrefixTableFromTree(
	ht *binarytree.HuffTree,
	ft *frequencymap.FrequencyMap,
) ([]PrefixNode, error) {
	if ht == nil {
		return nil, errors.New("the tree is nil")
	}
	pt := []PrefixNode{}

	c1 := make(chan binarytree.BaseNode)
	go ht.Walker(c1)
	for v1 := range c1 {
		if v1.IsLeaf() {
			ln := v1.(binarytree.LeafNode)
			table := ft.GetFrequency()

			ftc, ok := table[ln.Character]

			if !ok {
				return nil, errors.New("Could not find the character in the given table")
			}

			pt = append(pt, PrefixNode{
				Character: ln.Character,
				Code:      ln.Code,
				Frequency: ftc,
			})
		}
	}

	return pt, nil
}
