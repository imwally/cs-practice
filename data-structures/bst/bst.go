package bst

import (
	"fmt"
	"strings"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func New(v int) *Node {
	return &Node{v, nil, nil}
}

func (n *Node) Insert(v int) {
	newNode := New(v)
	for current := n; current != nil; {
		if v < current.Value {
			if current.Left == nil {
				current.Left = newNode
				return
			}
			current = current.Left

		} else {
			if current.Right == nil {
				current.Right = newNode
				return
			}
			current = current.Right
		}
	}
}

func (n *Node) Find(v int) *Node {
	for current := n; current != nil; {
		if current.Value == v {
			return current
		}
		if v < current.Value {
			current = current.Left
		} else {
			current = current.Right
		}
	}

	return nil
}

func printPre(n *Node, s string) string {
	if n == nil {
		return ""
	}

	s += fmt.Sprintf("%d ", n.Value)

	if n.Left != nil {
		s = printPre(n.Left, s)
	}

	if n.Right != nil {
		s = printPre(n.Right, s)
	}

	return s
}

func PrintPreOrder(n *Node, s string) string {
	return strings.TrimRight(printPre(n, s), " ")
}

func printIn(n *Node, s string) string {
	if n == nil {
		return ""
	}

	if n.Left != nil {
		s = printIn(n.Left, s)
	}

	s += fmt.Sprintf("%d ", n.Value)

	if n.Right != nil {
		s = printIn(n.Right, s)
	}

	return s
}

func PrintInOrder(n *Node, s string) string {
	return strings.TrimRight(printIn(n, s), " ")
}
