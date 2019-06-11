package heap

import "testing"

func TestInsert(t *testing.T) {
	h := new(Heap)

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
