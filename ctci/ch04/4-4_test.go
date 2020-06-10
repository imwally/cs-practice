package main

import (
	"testing"
)

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}

func treeHeight(node *Node) (int, int) {
	if node == nil {
		return 0, 0
	}

	l, _ := treeHeight(node.Left)
	_, r := treeHeight(node.Right)

	l++
	r++

	return l, r
}

func Balanced(node *Node) bool {
	l, r := treeHeight(node)

	if Abs(l-r) > 1 {
		return false
	}

	return true
}

func TestBalanced(t *testing.T) {
	root := &Node{1, nil, nil}
	root.Left = &Node{2, nil, nil}
	root.Right = &Node{3, nil, nil}

	root.Left.Left = &Node{4, nil, nil}
	root.Left.Right = &Node{5, nil, nil}

	root.Right.Left = &Node{6, nil, nil}
	root.Right.Right = &Node{7, nil, nil}
	root.Right.Right.Right = &Node{8, nil, nil}

	got, expected := Balanced(root), true
	if got != expected {
		t.Errorf("error: got %v, expected %v", got, expected)
	}

	// Increase right subtree
	root.Right.Right.Right.Right = &Node{9, nil, nil}

	got, expected = Balanced(root), false
	if got != expected {
		t.Errorf("error: got %v, expected %v", got, expected)
	}
}
