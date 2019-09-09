package linkedlist

import (
	"fmt"
	"testing"
)

func TestFlatten(t *testing.T) {
	llc := &LinkedList{}
	llc.Append(4)
	llc.Append(5)
	llc.Append(6)

	llb := &LinkedList{}
	llb.Append(7)
	llb.Append(8)
	llb.Append(9)

	lla := &LinkedList{}
	lla.Append(1)
	lla.Append(llb)
	lla.Append(2)
	lla.Append(llc)
	lla.Append(3)

	for current := lla.Head; current != nil; current = current.Next {
		fmt.Println(current.Value)
	}
}
