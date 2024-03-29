// Exercise 2.1
//
// Remove Dups: Write code to remove duplicates from an unsorted li nked list.

package main

import (
	"fmt"
	"testing"

	"../../data-structures/linkedlist"
)

func RemoveDups(ll *linkedlist.LinkedList) {
	seen := make(map[interface{}]int)

	last := ll.Head
	for current := ll.Head; current != nil; current = current.Next {
		seen[current.Value]++
		if seen[current.Value] > 1 {
			last.Next = current.Next
			continue
		}
		last = current
	}
}

func TestRemoveDups(t *testing.T) {
	ll := &linkedlist.LinkedList{}

	ll.Append(1)
	ll.Append(1)
	ll.Append(2)
	ll.Append(2)
	ll.Append(2)
	ll.Append(3)
	ll.Append(3)
	ll.Append(4)
	ll.Append(1)
	ll.Append(4)
	ll.Append(2)
	ll.Append(5)
	ll.Append(5)
	ll.Append(2)
	ll.Append(1)

	expected := "1 -> 1 -> 2 -> 2 -> 2 -> 3 -> 3 -> 4 -> 1 -> 4 -> 2 -> 5 -> 5 -> 2 -> 1"
	got := fmt.Sprint(ll)

	if got != expected {
		t.Errorf("error: got %s, expected %s", got, expected)
	}

	RemoveDups(ll)

	expected = "1 -> 2 -> 3 -> 4 -> 5"
	got = fmt.Sprint(ll)

	if got != expected {
		t.Errorf("error: got %s, expected %s", got, expected)
	}
}
