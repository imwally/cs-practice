package stack

import "testing"

func TestPush(t *testing.T) {
	s := new(Stack)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)

	for i := 0; i < len(s.Storage); i++ {
		if s.Storage[i] != i+1 {
			t.Errorf("Push failed: %d does not match %d", i, s.Storage[i])
		}
	}
}

func TestPeek(t *testing.T) {
	s := new(Stack)

	s.Push(1)
	if s.Peek() != 1 {
		t.Error("Peek failed: should be 1")
	}

	s.Push(2)
	if s.Peek() != 2 {
		t.Error("Peek failed: should be 2")
	}

	s.Push(3)
	if s.Peek() != 3 {
		t.Error("Peek failed: should be 3")
	}
}
