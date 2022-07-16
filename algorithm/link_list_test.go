package algorithm

import (
	"fmt"
	"testing"
)

type node struct {
	data int
	next *node
}

func showNode(p *node) {
	for p != nil {
		fmt.Println(*p)
		p = p.next
	}
}

func Test_Link(t *testing.T) {
	var head = new(node)
	head.data = 1
	var node1 = new(node)
	node1.data = 2

	head.next = node1
	var node2 = new(node)
	node2.data = 3
	node1.next = node2

	showNode(head)
}
