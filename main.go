package main

import (
	"fmt"
	"github.com/timon-schelling/gohuffmancode/gohuffmancode"
	"os"
	"strconv"
)

func main() {

	text := os.Args[1]

	cm := make(map[int32]int)

	for _, c := range text {
		if _, ok := cm[c]; ok {
			cm[c]++
		} else {
			cm[c] = 1
		}
	}

	var leaves []*gohuffmancode.Node

	for k, v := range cm {
		leaves = append(leaves, &gohuffmancode.Node{Value: gohuffmancode.ValueType(k), Count: v})
	}

	root := gohuffmancode.Build(leaves)
	Print(root)
}

func Print(root *gohuffmancode.Node) {
	var traverse func(n *gohuffmancode.Node, code uint64, bits byte)
	traverse = func(n *gohuffmancode.Node, code uint64, bits byte) {
		if n.Left == nil {
			fmt.Printf("'%c': %0"+strconv.Itoa(int(bits))+"b\n", n.Value, code)
			return
		}
		bits++
		traverse(n.Left, code<<1, bits)
		traverse(n.Right, code<<1+1, bits)
	}
	traverse(root, 0, 0)
}
