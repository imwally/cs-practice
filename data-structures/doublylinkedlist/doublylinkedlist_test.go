package doublylinkedlist

import (
	"fmt"
	"testing"
)

func TestAppend(t *testing.T) {
	dll := new(DoublyLinkedList)

	for i := 1; i < 6; i++ {
		dll.Append(i)
	}

	for i, cur := 1, dll.Head; cur != nil; i, cur = i+1, cur.Next {
		if cur.Value != i {
			t.Errorf("%d != %d", cur.Value, i)
		}
	}
}

func TestPrint(t *testing.T) {
	dll := new(DoublyLinkedList)

	dll.Append(1)
	if fmt.Sprintf("%s", dll) != "1" {
		t.Errorf("Print failed: printed \"%s\" instead of \"1\"", dll)
	}

	dll.Append(2)
	if fmt.Sprintf("%s", dll) != "1 -> 2" {
		t.Errorf("Print failed: printed \"%s\" instead \"1 -> 2\"", dll)
	}

	dll.Append(3)
	if fmt.Sprintf("%s", dll) != "1 -> 2 -> 3" {
		t.Errorf("Print failed: printed \"%s\" instead \"1 -> 2 -> 3\"", dll)
	}
}

func TestSize(t *testing.T) {
	dll := new(DoublyLinkedList)

	dll.Append(1)
	dll.Append(2)
	if dll.Size() != 2 {
		t.Errorf("Size failed: got %d, expected %d", dll.Size(), 2)
	}

	dll.Append(3)
	dll.Append(4)
	if dll.Size() != 4 {
		t.Errorf("Size failed: got %d, expected %d", dll.Size(), 2)
	}
}

func TestRemove(t *testing.T) {
	dll := new(DoublyLinkedList)

	dll.Append(1)
	dll.Append(2)
	dll.Append(3)

	dll.Remove(2)

	if fmt.Sprintf("%s", dll) != "1 -> 3" {
		t.Errorf("Remove failed: got \"%s\", expected \"1 -> 3\"", dll)
	}

	if dll.Head.Value != 1 {
		t.Errorf("Remove failed: list Head is %d, expected 1", dll.Head.Value)
	}

	if dll.Tail.Value != 3 {
		t.Errorf("Remove failed: list Head is %d, expected 3", dll.Head.Value)
	}

	dll.Remove(1)

	if fmt.Sprintf("%s", dll) != "3" {
		t.Errorf("Remove failed: got \"%s\", expected \"3\"", dll)
	}

	if dll.Head.Value != 3 && dll.Tail.Value != 3 {
		t.Errorf("Remove failed: list Head and Tail point to %d and %d, expected both to point to 3", dll.Head.Value, dll.Tail.Value)
	}
}
