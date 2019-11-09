// Exercose 3.2
//
// Stack Min
//
// How would you design a stack which, in addition to push and pop, has a
// function min which returns the minimum element? Push, pop and min should all
// operate in 0(1) time.

package main

import "testing"

type Stack struct {
	Storage  []int
	MinIndex int
	Index    int
}

func New() *Stack {
	storage := make([]int, 10)
	return &Stack{
		Storage:  storage,
		MinIndex: 0,
		Index:    0,
	}
}

func (s *Stack) Push(v int) {
	if v < s.Storage[s.MinIndex] {
		s.MinIndex = s.Index
	}

	s.Storage[s.Index] = v
	s.Index++
}

func (s *Stack) Pop() int {
	popped := s.Storage[s.Index]
	s.Index--

	return popped
}

func (s *Stack) Min() int {
	return s.Storage[s.MinIndex]
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

}
