package main

import (
	"testing"

	"../../data-structures/linkedlist"
)

func LoopDetection(ll *linkedlist.LinkedList) *linkedlist.Node {
	seen := make(map[*linkedlist.Node]int)
	for c := ll.Head; c != nil; c = c.Next {
		seen[c]++
		if seen[c] > 1 {
			return c
		}
	}

	return nil
}

func TestLoopDetection(t *testing.T) {
	// Manually construct linked list
	ll := &linkedlist.LinkedList{}

	n1 := &linkedlist.Node{1, nil}
	n2 := &linkedlist.Node{2, nil}
	n3 := &linkedlist.Node{3, nil}
	n4 := &linkedlist.Node{4, nil}
	n5 := &linkedlist.Node{5, nil}
	n6 := &linkedlist.Node{6, nil}

	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	n5.Next = n6
	n6.Next = n3

	ll.Head = n1
	ll.Tail = n6

	expected := n3
	got := LoopDetection(ll)

	if got != expected {
		t.Errorf("LoopDetection error: got %v, expected %v", got, expected)
	}

	ll2 := &linkedlist.LinkedList{}
	for i := 1; i < 6; i++ {
		ll2.Append(i)
	}

	expected = nil
	got = LoopDetection(ll2)

	if got != expected {
		t.Errorf("LoopDetection error: got %v, expected %v", got, expected)
	}
}
