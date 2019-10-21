// Exercise 2.4

package main

import (
	"fmt"
	"testing"

	"../data-structures/linkedlist"
)

func Partition(ll *linkedlist.LinkedList, v int) *linkedlist.LinkedList {
	less := &linkedlist.LinkedList{}
	greater := &linkedlist.LinkedList{}

	for c := ll.Head; c != nil; c = c.Next {
		if c.Value.(int) < v {
			less.Append(c.Value)
		} else {
			greater.Append(c.Value)
		}
	}

	newLL := &linkedlist.LinkedList{}

	for lc := less.Head; lc != nil; lc = lc.Next {
		newLL.Append(lc.Value)
	}

	for gc := greater.Head; gc != nil; gc = gc.Next {
		newLL.Append(gc.Value)
	}

	return newLL
}

func TestPartition(t *testing.T) {
	ll := &linkedlist.LinkedList{}
	ll.Append(3)
	ll.Append(5)
	ll.Append(8)
	ll.Append(5)
	ll.Append(10)
	ll.Append(2)
	ll.Append(1)

	expected := "3 -> 5 -> 8 -> 5 -> 10 -> 2 -> 1"
	got := fmt.Sprint(ll)

	if got != expected {
		t.Errorf("error: got %s, expected %s", got, expected)
	}

	expected = "3 -> 2 -> 1 -> 5 -> 8 -> 5 -> 10"
	got = fmt.Sprint(Partition(ll, 5))

	if got != expected {
		t.Errorf("error: got %s, expected %s", got, expected)
	}
}
