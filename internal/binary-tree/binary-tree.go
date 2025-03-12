// Package binarytree is used to hold all utils for a binary tree
package binarytree

import (
	"cmp"
	"slices"
)

// BaseNode is the interface for the node
type BaseNode interface {
	weight() int
	isLeaf() bool
}

// LeafNode is the leaf to the tree
type LeafNode struct {
	Character rune
	Weight    int
	Left      *BaseNode
	Right     *BaseNode
}

func (n LeafNode) weight() int {
	return n.Weight
}

func (n LeafNode) isLeaf() bool {
	return true
}

// InternalNode is the struct that connects two leaf nodes
type InternalNode struct {
	Weight int
	Left   *BaseNode
	Right  *BaseNode
}

func (in InternalNode) weight() int {
	return in.Weight
}

func (in InternalNode) isLeaf() bool {
	return false
}

// HuffTree is the struct for the tree
type HuffTree struct {
	Root BaseNode
}

// Weight returns the total weight of the huffman tree
func (ht *HuffTree) Weight() int {
	return ht.Root.weight()
}

func newHuffInternalTree(l, r BaseNode, wt int) HuffTree {
	return HuffTree{
		Root: InternalNode{
			Left:   &l,
			Right:  &r,
			Weight: wt,
		},
	}
}

func removeMinElement(ht []HuffTree) (HuffTree, []HuffTree) {
	return ht[0], ht[1:]
}

// CreateBinaryTreeFromMap creates a binary tree from a frequency map
func CreateBinaryTreeFromMap(m map[rune]int) HuffTree {
	leafNodeArr := []HuffTree{}
	for char, freq := range m {
		leafNodeArr = append(
			leafNodeArr,
			HuffTree{
				Root: LeafNode{
					Character: char, Weight: freq, Left: nil,
					Right: nil,
				},
			},
		)
	}

	var tmp1, tmp2, tmp3 HuffTree

	for len(leafNodeArr) > 1 {
		slices.SortFunc(leafNodeArr, func(i, j HuffTree) int {
			return cmp.Compare(i.Weight(), j.Weight())
		})

		tmp1, leafNodeArr = removeMinElement(leafNodeArr)
		tmp2, leafNodeArr = removeMinElement(leafNodeArr)
		tmp3 = newHuffInternalTree(tmp1.Root, tmp2.Root, tmp1.Weight()+tmp2.Weight())
		leafNodeArr = append(leafNodeArr, tmp3)
	}

	return tmp3
}
