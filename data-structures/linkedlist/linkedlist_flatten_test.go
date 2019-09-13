package linkedlist

import (
	"fmt"
	"testing"
)

func Merge(a, b *LinkedList) *LinkedList {
	merged := &LinkedList{}
	if a == nil {
		return b
	}

	if b == nil {
		return a
	}

	ac := a.Head
	bc := b.Head

	// Merge lists in order
	for ac != nil && bc != nil {
		if ac.Value.(int) <= bc.Value.(int) {
			merged.Append(ac.Value)
			ac = ac.Next
		} else {
			merged.Append(bc.Value)
			bc = bc.Next
		}
	}

	// Add remaining elements from b list
	for bc != nil {
		merged.Append(bc.Value)
		bc = bc.Next
	}

	// Add remaining elements from a list
	for ac != nil {
		merged.Append(ac.Value)
		ac = ac.Next
	}

	return merged
}

func Flatten(n *Node) *LinkedList {
	if n.Next == nil {
		return Merge(n.Value.(*LinkedList), nil)
	}

	return Merge(n.Value.(*LinkedList), Flatten(n.Next))
}

func TestFlatten(t *testing.T) {
	lld := &LinkedList{}
	lld.Append(1)
	lld.Append(2)
	lld.Append(3)

	llc := &LinkedList{}
	llc.Append(4)
	llc.Append(5)
	llc.Append(6)

	llb := &LinkedList{}
	llb.Append(7)
	llb.Append(8)
	llb.Append(9)

	lla := &LinkedList{}
	lla.Append(llb)
	lla.Append(llc)
	lla.Append(lld)

	flat := Flatten(lla.Head)

	got := fmt.Sprintf("%s", flat)
	expected := "1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> 8 -> 9"
	if got != expected {
		t.Errorf("Flatten error: got %v, expected %v", got, expected)
	}
}
