package main

import (
	"testing"
)

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

func ValidateBST(node *Node, min int, max int) bool {
	if node == nil {
		return true
	}

	if min > node.Data || node.Data > max {
		return false
	}

	return ValidateBST(node.Left, min, node.Data) && ValidateBST(node.Right, node.Data, max)
}

func TestValidateBST(t *testing.T) {
	root := &Node{4, nil, nil}

	root.Left = &Node{3, nil, nil}
	root.Left.Left = &Node{2, nil, nil}

	root.Right = &Node{5, nil, nil}
	root.Right.Left = &Node{4, nil, nil}
	root.Right.Right = &Node{6, nil, nil}

	got, expected := ValidateBST(root, -999, 999), true
	if got != expected {
		t.Errorf("error: got %v, expected %v", got, expected)
	}

	root.Left.Right = &Node{6, nil, nil}

	got, expected = ValidateBST(root, -999, 999), false
	if got != expected {
		t.Errorf("error: got %v, expected %v", got, expected)
	}
}
