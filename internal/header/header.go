// Package header is responible for handling the header for the huffman compression
package header

import (
	"errors"

	binarytree "github.com/KylerWilson01/huffman-compression/internal/binary-tree"
)

// PrefixNode allows us to create a prefix header for encoding the file
type PrefixNode struct {
	Character rune
	Code      []int
	Frequency int
}

func walk(t binarytree.BaseNode, code []int, ch chan PrefixNode) {
	if t == nil {
		return
	}

	walk(t.Left(), append(code, 0b0), ch)
	if t.IsLeaf() {
		bn := t.(binarytree.LeafNode)
		ch <- PrefixNode{Character: bn.Character, Code: code, Frequency: bn.Weight}
	}
	walk(t.Right(), append(code, 0b1), ch)
}

func walker(root binarytree.BaseNode, ch chan PrefixNode) {
	code := []int{}
	walk(root, code, ch)
	close(ch)
}

// CreatePrefixTableFromTree creates a prefix table from the binary tree that was previously created
func CreatePrefixTableFromTree(
	ht *binarytree.HuffTree,
	ft map[rune]int,
) ([]PrefixNode, error) {
	if ht == nil {
		return nil, errors.New("the tree is nil")
	}
	pt := []PrefixNode{}

	c1 := make(chan PrefixNode)
	go walker(ht.Root, c1)
	for v1 := range c1 {
		_, ok := ft[v1.Character]

		if !ok {
			return nil, errors.New("Could not find the character in the given table")
		}

		pt = append(pt, v1)
	}

	return pt, nil
}
