// Exercise 3.1
//
// Three in One
//
// Describe how you could use a single array to implement three stacks.

package main

import "testing"

type Stacks struct {
	Storage  []int
	Capacity int
	One      int
	Two      int
	Three    int
}

func New() *Stacks {
	cap := 30
	storage := make([]int, cap)
	return &Stacks{
		Storage:  storage,
		Capacity: cap,
		One:      0,
		Two:      0,
		Three:    0,
	}
}

func (s *Stacks) Push(which, v int) {
	i := 0
	switch which {
	case 1:
		i = s.One
		s.One++
	case 2:
		i = (s.Capacity / 3) + s.Two
		s.Two++
	case 3:
		i = (s.Capacity / 3) + (s.Capacity / 3) + s.Three
		s.Three++
	}

	s.Storage[i] = v
}

func (s *Stacks) Pop(which int) int {
	i := 0
	switch which {
	case 1:
		s.One--
		i = s.One
	case 2:
		s.Two--
		i = (s.Capacity / 3) + s.Two
	case 3:
		s.Three--
		i = (s.Capacity / 3) + (s.Capacity / 3) + s.Three
	}

	return s.Storage[i]
}

func TestPush(t *testing.T) {

	stacks := New()

	// Testing Stack 1
	stacks.Push(1, 1)
	stacks.Push(1, 2)
	stacks.Push(1, 3)

	expected := 3
	got := stacks.Pop(1)

	if got != expected {
		t.Errorf("Pop error: got %v, expected %v", got, expected)
	}

	// Testing Stack 2
	stacks.Push(2, 10)
	stacks.Push(2, 11)
	stacks.Push(2, 12)

	expected = 12
	got = stacks.Pop(2)

	if got != expected {
		t.Errorf("Pop error: got %v, expected %v", got, expected)
	}

	// Testing Stack 3
	stacks.Push(3, 20)
	stacks.Push(3, 21)
	stacks.Push(3, 22)

	expected = 22
	got = stacks.Pop(3)

	if got != expected {
		t.Errorf("Pop error: got %v, expected %v", got, expected)
	}

}
