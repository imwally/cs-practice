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
