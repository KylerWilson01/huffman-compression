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

func walk(t binarytree.BaseNode, code []byte, ch chan PrefixNode) {
	if t == nil {
		return
	}

	walk(t.Left(), append(code, 0), ch)
	if t.IsLeaf() {
		bn := t.(binarytree.LeafNode)
		ch <- PrefixNode{Character: bn.Character, Code: code, Frequency: bn.Weight}
	}
	walk(t.Right(), append(code, 1), ch)
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

	c1 := make(chan PrefixNode)
	go walk(ht.Root, []byte{}, c1)
	for v1 := range c1 {
		table := ft.GetFrequency()

		_, ok := table[v1.Character]

		if !ok {
			return nil, errors.New("Could not find the character in the given table")
		}

		pt = append(pt, v1)
	}

	return pt, nil
}
