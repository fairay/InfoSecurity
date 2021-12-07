package main

import "fmt"

type Node struct {
	val   *byte
	count int
	n0    *Node
	n1    *Node
}

func NewNode(d byte, count int) *Node {
	return &Node{&d, count, nil, nil}
}
func MergeNodes(n0, n1 *Node) *Node {
	return &Node{nil, n0.count + n1.count, n0, n1}
}
func (this *Node) Print(code string) {
	if this.val != nil {
		fmt.Printf("%s \"%c\" %v\n", code, *this.val, this.count)
	} else {
		this.n0.Print(code + "0")
		this.n1.Print(code + "1")
	}
}

func (this *Node) ToMap(bits *BitSet, m *CmpMap) {
	if this.val != nil {
		m.Table[*this.val] = bits
	} else {
		bit0 := bits.Copy()
		bit0.ForwardBit(false)
		this.n0.ToMap(bit0, m)

		bit1 := bits.Copy()
		bit1.ForwardBit(true)
		this.n1.ToMap(bit1, m)
	}
}

type prefixTree struct {
	root *Node
	arr  []*Node
}

func NewPrefixTree(stat *statMap) *prefixTree {
	tree := new(prefixTree)
	tree.arr = make([]*Node, 0)
	for k, v := range stat.st {
		tree.arr = append(tree.arr, NewNode(k, v))
		tree.BackSort()
	}

	tree.Collapse()
	return tree
}
func (this *prefixTree) BackSort() {
	for i := len(this.arr) - 1; i > 0; i-- {
		if this.arr[i].count > this.arr[i-1].count {
			this.arr[i], this.arr[i-1] = this.arr[i-1], this.arr[i]
		} else {
			break
		}
	}
}
func (this *prefixTree) Collapse() {
	for i := len(this.arr) - 1; i > 0; i-- {
		merged := MergeNodes(this.arr[i], this.arr[i-1])
		this.arr[i-1] = merged
		this.arr = this.arr[:i]
		this.BackSort()
	}

	this.root = this.arr[0]
}
func (this *prefixTree) ToMap() *CmpMap {
	m := new(CmpMap)
	m.Table = make(map[byte]*BitSet)
	this.root.ToMap(EmptyBitSet(), m)
	return m
}
