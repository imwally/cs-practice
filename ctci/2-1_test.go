package main

import (
	"testing"

	"../data-structures/linkedlist"
)

func KthToLast(ll *linkedlist.LinkedList, k int) int {
	current := ll.Head
	for i := 1; i < ll.Size()-k; i++ {
		current = current.Next
	}

	return current.Value.(int)
}

func TestKthToLast(t *testing.T) {
	l := &linkedlist.LinkedList{}

	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)

	got := KthToLast(l, 2)
	if got != 2 {
		t.Errorf("error: expected 2, got %d\n", got)
	}
}
