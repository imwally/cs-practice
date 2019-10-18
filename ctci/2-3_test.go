// Exercise 2.3 - Delete Middle Node
//
// Implement an algorithm to delete a node in the middle (i.e., any node but
// the first and last node, not necessarily the exact middle) of a singly
// linked list, given only access to that node.
//
// EXAMPLE
// Input: the node c from the linked list a -> b-> c -> d -> e -> f
// Result: nothing is returned, but the new linked list looks like a -> b-> d -> e -> f

package main

import (
	"errors"
	"fmt"
	"testing"

	"../data-structures/linkedlist"
)

func RemoveMiddleNode(ll *linkedlist.LinkedList, pos int) error {
	if pos == 0 || pos == ll.Size()-1 {
		return errors.New("index out of bounds")
	}

	for n, i := ll.Head, 0; n != nil; n, i = n.Next, i+1 {
		if i >= pos {
			n.Value = n.Next.Value
			if n.Next.Next == nil {
				n.Next = nil
			}
		}
	}

	return nil
}

func TestRemoveMiddleNode(t *testing.T) {
	ll := &linkedlist.LinkedList{}

	for i := 1; i <= 4; i++ {
		ll.Append(i)
	}

	expected := "1 -> 2 -> 3 -> 4"
	got := fmt.Sprint(ll)

	if got != expected {
		t.Errorf("error: got %s, expected %s", got, expected)
	}

	err := RemoveMiddleNode(ll, 1)
	if err != nil {
		t.Error(err)
	}

	expected = "1 -> 3 -> 4"
	got = fmt.Sprint(ll)

	if got != expected {
		t.Errorf("error: got %s, expected %s", got, expected)
	}

	err = RemoveMiddleNode(ll, 1)
	if err != nil {
		t.Error(err)
	}
	expected = "1 -> 4"
	got = fmt.Sprint(ll)

	if got != expected {
		t.Errorf("error: got %s, expected %s", got, expected)
	}
}
