package heap

import "testing"

func TestPush(t *testing.T) {
	h := New()

	expected := 10
	h.Push(expected)
	got := h.Data[0]
	if got != expected {
		t.Errorf("Push failed: expected %d, got %d", expected, got)
	}

	expected = 5
	h.Push(expected)
	got = h.Data[0]
	if got != expected {
		t.Errorf("Push failed: expected %d, got %d", expected, got)
	}

	expected = 1
	h.Push(expected)
	got = h.Data[0]
	if got != expected {
		t.Errorf("Push failed: expected %d, got %d", expected, got)
	}
}

func TestPop(t *testing.T) {
	h := New()
	size := 40

	for i := size; i > 0; i-- {
		h.Push(i)
	}

	for i := 1; i <= size; i++ {
		expected := i
		got := h.Pop()
		if got != expected {
			t.Errorf("Pop failed: expected %d, got %d", expected, got)
		}
	}
}

func TeatPeek(t *testing.T) {
	h := New()

	h.Push(5)
	expected := 5
	got := h.Peek()
	if got != expected {
		t.Errorf("Peek failed: expected %d, got %d", expected, got)
	}

	h.Push(4)
	expected = 4
	got = h.Peek()
	if got != expected {
		t.Errorf("Peek failed: expected %d, got %d", expected, got)
	}

	h.Push(6)
	expected = 4
	got = h.Peek()
	if got != expected {
		t.Errorf("Peek failed: expected %d, got %d", expected, got)
	}

	h.Push(1)
	expected = 1
	got = h.Peek()
	if got != expected {
		t.Errorf("Peek failed: expected %d, got %d", expected, got)
	}
}
