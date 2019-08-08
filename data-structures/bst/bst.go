package bst

import "fmt"

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

func PrintPreOrder(n *Node) {
	if n == nil {
		return
	}

	fmt.Println(n.Value)

	PrintPreOrder(n.Left)
	PrintPreOrder(n.Right)
}

func PrintInOrder(n *Node) {
	if n == nil {
		return
	}

	PrintInOrder(n.Left)
	fmt.Println(n.Value)
	PrintInOrder(n.Right)
}
