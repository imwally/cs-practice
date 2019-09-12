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

	for ac != nil || bc != nil {
		if ac == nil {
			merged.Append(bc.Value)
			bc = bc.Next
			continue
		}
		if bc == nil {
			merged.Append(ac.Value)
			ac = ac.Next
			continue
		}
		if ac.Value.(int) <= bc.Value.(int) {
			merged.Append(ac.Value)
			ac = ac.Next
			continue
		} else {
			merged.Append(bc.Value)
			bc = bc.Next
		}
	}

	return merged
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

	m := Merge(Merge(llb, llc), lld)

	got := fmt.Sprintf("%s", m)
	expected := "1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> 8 -> 9"
	if got != expected {
		t.Errorf("Flatten error: got %v, expected %v", got, expected)
	}
}
