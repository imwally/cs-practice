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

	// Why use
	//
	// left + (right-left)/2
	// vs.
	// (left+right)/2
	//
	// Because the first version never grows larger than the size
	// of the input. This could eventually lead to overflows given
	// a very large input size.
	//
	// Ex with input size of 65535 and last recursive call:
	//
	// 65534 + (65534-65534)
	// left + (right-left) = 65534 / 2 = 65534
	// vs.
	// (65534+65534)
	// (left+right) = 131068 / 2 = 65534
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

	var s []int
	for i := 1; i <= 3; i++ {
		s = append(s, i)
	}

	tree := new(BST)
	tree.Root = MinimalBST(s, 0, len(s)-1)

	got, expected := TreeHeight(tree.Root), 2
	if got != expected {
		t.Errorf("error: got %v, expected %v", got, expected)
	}

	s = []int{}
	for i := 1; i <= 8; i++ {
		s = append(s, i)
	}
	tree = new(BST)
	tree.Root = MinimalBST(s, 0, len(s)-1)

	got, expected = TreeHeight(tree.Root), 4
	if got != expected {
		t.Errorf("error: got %v, expected %v", got, expected)
	}
}
