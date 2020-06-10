package main

import (
	"fmt"
	"testing"

	"../../data-structures/linkedlist"
	"../../data-structures/queue"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func minimalBST(s []int, left int, right int) *Node {
	if left > right {
		return nil
	}

	mid := left + (right-left)/2

	current := &Node{s[mid], nil, nil}

	current.Left = minimalBST(s, left, mid-1)
	current.Right = minimalBST(s, mid+1, right)

	return current
}

func MinimalBST(s []int) *Node {
	return minimalBST(s, 0, len(s)-1)
}

func DepthLinkedList(current *Node) []*linkedlist.LinkedList {
	q := new(queue.Queue)
	q.Enqueue(current)

	// Slice to hold linked lists at each level
	var llSlice []*linkedlist.LinkedList

	// Loop over items in the queue
	for q.Length() > 0 {

		// Number of nodes at this current depth
		nodesAtDepth := q.Length()

		// Create a new linked list for each depth
		ll := new(linkedlist.LinkedList)

		// Loop over each node at depth
		for nodesAtDepth > 0 {

			// Process node at current depth
			current = q.Dequeue().(*Node)

			ll.Append(current.Value)

			if current.Left != nil {
				q.Enqueue(current.Left)
			}

			if current.Right != nil {
				q.Enqueue(current.Right)
			}

			// Decrement counter after processing node
			nodesAtDepth--
		}

		// All nodes have been processed at depth, add to llSlice
		llSlice = append(llSlice, ll)
	}

	return llSlice
}

func TestDepthLinkedList(t *testing.T) {
	var s []int
	for i := 1; i <= 8; i++ {
		s = append(s, i)
	}

	root := MinimalBST(s)

	lists := DepthLinkedList(root)

	got, expected := fmt.Sprintf("%s", lists[0]), "4"
	if got != expected {
		t.Errorf("error: got %v, expected %v", got, expected)
	}

	got, expected = fmt.Sprintf("%s", lists[1]), "2 -> 6"
	if got != expected {
		t.Errorf("error: got %v, expected %v", got, expected)
	}

	got, expected = fmt.Sprintf("%s", lists[2]), "1 -> 3 -> 5 -> 7"
	if got != expected {
		t.Errorf("error: got %v, expected %v", got, expected)
	}

	got, expected = fmt.Sprintf("%s", lists[3]), "8"
	if got != expected {
		t.Errorf("error: got %v, expected %v", got, expected)
	}
}
