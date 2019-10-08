package main

import (
	"fmt"
	"testing"

	"../data-structures/linkedlist"
)

func KthToLast(ll *linkedlist.LinkedList, k int) (int, error) {
	if k < 0 || k > ll.Size()-1 {
		return 0, fmt.Errorf("error: %d out of bounds", k)
	}

	current := ll.Head
	for i := 1; i < ll.Size()-k; i++ {
		current = current.Next
	}

	return current.Value.(int), nil
}

func TestKthToLast(t *testing.T) {
	ll := &linkedlist.LinkedList{}

	for i := 1; i <= 10; i++ {
		ll.Append(i)
	}

	got, err := KthToLast(ll, 1)
	if err != nil {
		t.Error(err)
	}

	expected := 9
	if got != expected {
		t.Errorf("error: expected %d, got %d\n", expected, got)
	}
}
