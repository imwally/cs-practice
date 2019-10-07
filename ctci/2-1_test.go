package main

import (
	"errors"
	"testing"

	"../data-structures/linkedlist"
)

func KthToLast(ll *linkedlist.LinkedList, k int) (int, error) {
	if k < 0 || k > ll.Size()-1 {
		return 0, errors.New("error: %d out of bounds")
	}

	current := ll.Head
	for i := 1; i < ll.Size()-k; i++ {
		current = current.Next
	}

	return current.Value.(int), nil
}

func TestKthToLast(t *testing.T) {
	l := &linkedlist.LinkedList{}

	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)

	got, err := KthToLast(l, 2)
	if err != nil {
		t.Error(err)
	}

	if got != 2 {
		t.Errorf("error: expected 2, got %d\n", got)
	}
}
