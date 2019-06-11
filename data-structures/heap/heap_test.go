package heap

import (
	"testing"
)

func TestInsert(t *testing.T) {
	h := New()

	expected := 10
	h.Insert(expected)
	got := h.Data[0]
	if got != expected {
		t.Errorf("Insert failed: expected %d, got %d", expected, got)
	}

	expected = 5
	h.Insert(expected)
	got = h.Data[0]
	if got != expected {
		t.Errorf("Insert failed: expected %d, got %d", expected, got)
	}

	expected = 1
	h.Insert(expected)
	got = h.Data[0]
	if got != expected {
		t.Errorf("Insert failed: expected %d, got %d", expected, got)
	}
}

func TestRemove(t *testing.T) {
	h := New()
	size := 10

	for i := size; i > 0; i-- {
		h.Insert(i)
	}

	for i := 1; i <= size; i++ {
		expected := i
		got := h.Remove()
		if got != expected {
			t.Errorf("Remove failed: expected %d, got %d", expected, got)
		}
	}
}
