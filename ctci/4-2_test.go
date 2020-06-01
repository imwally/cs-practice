package main

import (
	"testing"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type BST struct {
	Root *Node
}

func MinimalBST(s []int, left int, right int) *Node {

	if left > right {
		return nil
	}

	mid := left + (right-left)/2

	current := &Node{s[mid], nil, nil}
	current.Left = MinimalBST(s, left, mid-1)
	current.Right = MinimalBST(s, mid+1, right)

	return current
}

func TreeHeight(current *Node) int {
	if current == nil {
		return 0

	}

	l := TreeHeight(current.Left)
	r := TreeHeight(current.Right)

	deeper := 0
	if l > r {
		deeper = l + 1
	} else {
		deeper = r + 1
	}

	return deeper
}

func TestMinimaBST(t *testing.T) {

	s := []int{1, 2, 3}

	tree := new(BST)
	tree.Root = MinimalBST(s, 0, len(s)-1)

	got, expected := TreeHeight(tree.Root), 2
	if got != expected {
		t.Errorf("error: got %v, expected %v", got, expected)
	}

	s = []int{1, 2, 3, 4, 5, 6, 7, 8}

	tree = new(BST)
	tree.Root = MinimalBST(s, 0, len(s)-1)

	got, expected = TreeHeight(tree.Root), 4
	if got != expected {
		t.Errorf("error: got %v, expected %v", got, expected)
	}
}
