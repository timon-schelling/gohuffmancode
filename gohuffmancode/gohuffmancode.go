package gohuffmancode

import (
	"sort"
)

type ValueType int32

type Node struct {
	Parent *Node
	Left   *Node
	Right  *Node
	Count  int
	Value  ValueType
}

func (n *Node) Code() (r uint64, bits byte) {
	for parent := n.Parent; parent != nil; n, parent = parent, parent.Parent {
		if parent.Right == n {
			r |= 1 << bits
		}
		bits++
	}
	return
}

type SortNodes []*Node

func (sn SortNodes) Len() int           { return len(sn) }
func (sn SortNodes) Less(i, j int) bool { return sn[i].Count < sn[j].Count }
func (sn SortNodes) Swap(i, j int)      { sn[i], sn[j] = sn[j], sn[i] }

func Build(leaves []*Node) *Node {
	sort.Stable(SortNodes(leaves))
	return BuildSorted(leaves)
}

func BuildSorted(leaves []*Node) *Node {
	if len(leaves) == 0 {
		return nil
	}

	for len(leaves) > 1 {
		left, right := leaves[0], leaves[1]
		parentCount := left.Count + right.Count
		parent := &Node{Left: left, Right: right, Count: parentCount}
		left.Parent = parent
		right.Parent = parent

		ls := leaves[2:]
		idx := sort.Search(len(ls), func(i int) bool { return ls[i].Count >= parentCount })
		idx += 2

		copy(leaves[1:], leaves[2:idx])
		leaves[idx-1] = parent
		leaves = leaves[1:]
	}

	return leaves[0]
}
