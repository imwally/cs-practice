// Exercise 2.7
//
// Intersection
//
// Given two (singly) linked lists, determine if the two lists intersect.
// Return the inter- secting node. Note that the intersection is defined based
// on reference, not value. That is, if the kth node of the first linked list
// is the exact same node (by reference) as the jth node of the second linked
// list, then they are intersecting.

package main

import (
	"testing"

	"../data-structures/linkedlist"
)

func Intersection(ll1, ll2 *linkedlist.LinkedList) bool {
	seen := make(map[*linkedlist.Node]int)

	for c1 := ll1.Head; c1 != nil; c1 = c1.Next {
		seen[c1]++
	}

	for c2 := ll2.Head; c2 != nil; c2 = c2.Next {
		seen[c2]++
		if seen[c2] > 1 {
			return true
		}
	}

	return false
}

func TestIntersection(t *testing.T) {
	// Construct first linked list from scractch
	ll1 := &linkedlist.LinkedList{}

	n1 := &linkedlist.Node{1, nil}
	n2 := &linkedlist.Node{2, nil}
	n3 := &linkedlist.Node{3, nil}

	n1.Next = n2
	n2.Next = n3

	ll1.Head = n1
	ll1.Tail = n3

	// Construct second linked list from scratch
	ll2 := &linkedlist.LinkedList{}

	n4 := &linkedlist.Node{4, nil}
	n5 := &linkedlist.Node{5, nil}

	n4.Next = n5
	n5.Next = n3

	ll2.Head = n4
	ll2.Tail = n3

	expected := true
	got := Intersection(ll1, ll2)

	if got != expected {
		t.Errorf("Intersection error: got %v, expected %v", got, expected)
	}

	ll3 := &linkedlist.LinkedList{}
	ll3.Append(8)
	ll3.Append(9)
	ll3.Append(10)

	expected = false
	got = Intersection(ll1, ll3)

	if got != expected {
		t.Errorf("Intersection error: got %v, expected %v", got, expected)
	}
}
