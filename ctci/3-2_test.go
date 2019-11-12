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
	Storage    []int
	MinStorage []int
	Index      int
}

func New() *Stack {
	storage := make([]int, 10)
	return &Stack{
		Storage: storage,
		Index:   0,
	}
}

func (s *Stack) Push(v int) {
	if len(s.MinStorage) == 0 {
		s.MinStorage = append(s.MinStorage, v)
	} else {
		if v < s.MinStorage[len(s.MinStorage)-1] {
			s.MinStorage = append(s.MinStorage, v)
		}
	}

	s.Storage[s.Index] = v
	s.Index++
}

// Technically using append() doesn't run in O(1) time.
func (s *Stack) Pop() int {
	s.Index--
	popped := s.Storage[s.Index]

	if popped == s.Min() {
		s.MinStorage = append(s.MinStorage[:len(s.MinStorage)-1], s.MinStorage[len(s.MinStorage):]...)
	}

	return popped
}

func (s *Stack) Min() int {
	return s.MinStorage[len(s.MinStorage)-1]
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

}
