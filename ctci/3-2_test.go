// Exercise 3.2
//
// Stack Min
//
// How would you design a stack which, in addition to push and pop, has a
// function min which returns the minimum element? Push, pop and min should all
// operate in 0(1) time.

package main

import (
	"testing"
)

type Node struct {
	Value interface{}
	Min   int
}

type Stack struct {
	Storage []Node
	Index   int
}

func New() *Stack {
	storage := make([]Node, 10)
	return &Stack{
		Storage: storage,
		Index:   0,
	}
}

func (s *Stack) Size() int {
	return s.Index
}

// Doesn't resize
func (s *Stack) Push(v int) {
	var n Node
	n.Value = v

	if s.Size() == 0 {
		n.Min = 0
	} else {
		// Get the last node and
		lastNode := s.Storage[s.Size()-1]

		// Last node's Min holds the index of the minimum value seen so
		// far. Therefore the node at that index will have the minimum
		// value seen.
		minValue := s.Storage[lastNode.Min].Value

		// If the currently added value is less that the minimum seen
		// so far, it's Min value will be the current index, as in it
		// points to itself.
		if v < minValue.(int) {
			n.Min = s.Index
		} else {
			// Otherwise it's Min value will be the minimum seen so
			// far, which is located in last node's Min. This gets
			// copied over.
			n.Min = lastNode.Min
		}
	}

	s.Storage[s.Index] = n
	s.Index++
}

func (s *Stack) Pop() int {
	if s.Size() < 0 {
		panic("stack underflow")
	}

	s.Index--
	popped := s.Storage[s.Index]

	return popped.Value.(int)
}

func (s *Stack) Min() int {
	lastNode := s.Storage[s.Index-1]
	minValue := s.Storage[lastNode.Min].Value

	return minValue.(int)
}

func TestMin(t *testing.T) {
	s := New()
	s.Push(30)
	s.Push(100)
	s.Push(20)
	s.Push(50)
	s.Push(10)

	expected := 10
	got := s.Min()

	if got != expected {
		t.Errorf("Min error: got %v, expected %v", got, expected)
	}

	s.Push(5)

	expected = 5
	got = s.Min()

	if got != expected {
		t.Errorf("Min error: got %v, expected %v", got, expected)
	}

	expected = 5
	got = s.Pop()

	if got != expected {
		t.Errorf("Min error: got %v, expected %v", got, expected)
	}

	expected = 10
	got = s.Min()

	if got != expected {
		t.Errorf("Min error: got %v, expected %v", got, expected)
	}

	expected = 10
	got = s.Pop()

	if got != expected {
		t.Errorf("Min error: got %v, expected %v", got, expected)
	}

	expected = 20
	got = s.Min()

	if got != expected {
		t.Errorf("Min error: got %v, expected %v", got, expected)
	}

	expected = 50
	got = s.Pop()

	if got != expected {
		t.Errorf("Min error: got %v, expected %v", got, expected)
	}

	expected = 20
	got = s.Min()

	if got != expected {
		t.Errorf("Min error: got %v, expected %v", got, expected)
	}

	expected = 20
	got = s.Pop()

	if got != expected {
		t.Errorf("Min error: got %v, expected %v", got, expected)
	}

	expected = 30
	got = s.Min()

	if got != expected {
		t.Errorf("Min error: got %v, expected %v", got, expected)
	}
}

func TestStackUnderflow(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Pop error: no stack underflow panic")
		}
	}()

	s := New()
	s.Push(1)
	s.Push(2)

	expected := 1
	got := s.Min()

	if got != expected {
		t.Errorf("Min error: got %v, expected %v", got, expected)
	}

	s.Pop()
	s.Pop()

	s.Pop()
}
