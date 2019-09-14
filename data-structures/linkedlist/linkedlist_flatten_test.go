package linkedlist

import (
	"fmt"
	"reflect"
	"testing"
)

func Merge(a, b *LinkedList) *LinkedList {
	if a == nil {
		return b
	}

	if b == nil {
		return a
	}

	ac := a.Head
	bc := b.Head

	// Merge lists in order
	merged := &LinkedList{}
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

func flatten(n *Node) *LinkedList {
	ll := n.Value.(*LinkedList)
	if n.Next == nil {
		return Merge(ll, nil)
	}

	return Merge(ll, flatten(n.Next))
}

func Flatten(ll *LinkedList) *LinkedList {
	if reflect.TypeOf(ll.Head.Value).String() != "*LinkedList" {
		return ll
	}

	return flatten(ll.Head)
}

func TestFlatten(t *testing.T) {
	lle := &LinkedList{}
	lle.Append(11)
	lle.Append(12)
	lle.Append(13)
	lle.Append(14)

	lld := &LinkedList{}
	lld.Append(1)
	lld.Append(2)
	lld.Append(3)
	lld.Append(10)

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
	lla.Append(lle)

	flat := Flatten(lla)

	got := fmt.Sprintf("%s", flat)
	expected := "1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> 8 -> 9 -> 10 -> 11 -> 12 -> 13 -> 14"
	if got != expected {
		t.Errorf("Flatten error: got %v, expected %v", got, expected)
	}

	flat = Flatten(lld)

	got = fmt.Sprintf("%s", flat)
	expected = "1 -> 2 -> 3 -> 10"
	if got != expected {
		t.Errorf("Flatten error: got %v, expected %v", got, expected)
	}

}
