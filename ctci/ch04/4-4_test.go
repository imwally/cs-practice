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

func height(current *Node) int {
	if current == nil {
		return 0
	}

	l := height(current.Left)
	r := height(current.Right)

	deeper := 0

	if l > r {
		deeper = l + 1
	} else {
		deeper = r + 1
	}

	return deeper
}

func Balanced(node *Node) bool {
	if node == nil {
		return true
	}

	lh := height(node.Left)
	rh := height(node.Right)

	if Abs(lh-rh) > 1 {
		return false
	}

	return Balanced(node.Left) && Balanced(node.Right)
}

func TestBalanced(t *testing.T) {

	got, expected := Balanced(nil), true
	if got != expected {
		t.Errorf("error: got %v, expected %v", got, expected)
	}

	tree := &Node{1, nil, nil}
	tree.Right = &Node{2, nil, nil}
	tree.Right.Right = &Node{3, nil, nil}

	got, expected = Balanced(tree), false
	if got != expected {
		t.Errorf("error: got %v, expected %v", got, expected)
	}

	// Try an equal amount of height coming off of the last right node.
	//
	//         1
	//          \
	//           2
	//            \
	//             3
	//            /
	//           4
	//          /
	//         5
	//        /
	//       6
	tree.Right.Right.Left = &Node{4, nil, nil}
	tree.Right.Right.Left.Left = &Node{5, nil, nil}
	tree.Right.Right.Left.Left.Left = &Node{6, nil, nil}

	got, expected = Balanced(tree), false
	if got != expected {
		t.Errorf("error: got %v, expected %v", got, expected)
	}
}
