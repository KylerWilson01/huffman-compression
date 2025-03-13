// Package binarytree is used to hold all utils for a binary tree
package binarytree

import (
	"cmp"
	"slices"
)

// BaseNode is the interface for the node
type BaseNode interface {
	weight() int
	IsLeaf() bool
	Left() BaseNode
	Right() BaseNode
}

// LeafNode is the leaf to the tree
type LeafNode struct {
	Character rune
	Weight    int
	L         BaseNode
	R         BaseNode
}

func (n LeafNode) weight() int {
	return n.Weight
}

// IsLeaf determines if the node is a leaf node or not
func (n LeafNode) IsLeaf() bool {
	return true
}

func (n LeafNode) Left() BaseNode {
	return n.L
}

func (n LeafNode) Right() BaseNode {
	return n.R
}

// InternalNode is the struct that connects two leaf nodes
type InternalNode struct {
	Weight int
	L      BaseNode
	R      BaseNode
}

func (in InternalNode) weight() int {
	return in.Weight
}

// IsLeaf determines if the node is a leaf node or not
func (in InternalNode) IsLeaf() bool {
	return false
}

func (in InternalNode) Left() BaseNode {
	return in.L
}

func (in InternalNode) Right() BaseNode {
	return in.R
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
			L:      l,
			R:      r,
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
					Character: char, Weight: freq, L: nil,
					R: nil,
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

// walk walks the tree
func walk(t BaseNode, ch chan BaseNode) {
	if t == nil {
		return
	}

	walk(t.Left(), ch)
	ch <- t
	walk(t.Right(), ch)
}

// Walker starts the walk of the tree
func (ht *HuffTree) Walker(ch chan BaseNode) {
	walk(ht.Root, ch)
	close(ch)
}

// Same checks to see if the trees are the same
func (ht *HuffTree) Same(ht2 *HuffTree) bool {
	if ht == nil || ht2 == nil {
		return false
	}

	c1, c2 := make(chan BaseNode), make(chan BaseNode)
	go ht.Walker(c1)
	go ht2.Walker(c2)
	for v1 := range c1 {
		v2, ok := <-c2
		if !ok {
			return false
		}
		if v1 != v2 {
			return false
		}
	}
	return true
}
