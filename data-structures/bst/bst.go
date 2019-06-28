package main

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

func main() {
	bst := New(10)
	bst.Insert(8)
	bst.Insert(11)

	fmt.Println(bst.Value)
	fmt.Println(bst.Left.Value)
	fmt.Println(bst.Right.Value)
}
