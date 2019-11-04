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

func (s *Stacks) Shrink() {
	oldThird := s.Capacity / 3
	s.Capacity = s.Capacity / 2
	newThird := s.Capacity / 3

	newStorage := make([]int, s.Capacity)

	// Copy first third
	copy(newStorage[:newThird], s.Storage[:oldThird])

	// Copy middle
	copy(newStorage[newThird:newThird*2], s.Storage[oldThird:oldThird*2])

	// Copy last third
	copy(newStorage[newThird*2:], s.Storage[oldThird*2:])

	s.Storage = newStorage

}

func (s *Stacks) Grow() {
	oldThird := s.Capacity / 3
	s.Capacity = s.Capacity * 2
	newThird := s.Capacity / 3

	newStorage := make([]int, s.Capacity)

	// Copy first third
	copy(newStorage[:newThird], s.Storage[:oldThird])

	// Copy middle
	copy(newStorage[newThird:newThird*2], s.Storage[oldThird:oldThird*2])

	// Copy last third
	copy(newStorage[newThird*2:], s.Storage[oldThird*2:])

	s.Storage = newStorage
}

func (s *Stacks) Push(which, v int) {
	i := 0
	third := (s.Capacity / 3)
	switch which {
	case 1:
		i = s.One
		if i == third {
			s.Grow()
		}
		s.One++
	case 2:
		i = third + s.Two
		if i == third*2 {
			s.Grow()
		}
		s.Two++
	case 3:
		i = third*2 + s.Three
		if i == s.Capacity {
			s.Grow()
		}
		s.Three++
	}

	s.Storage[i] = v
}

func (s *Stacks) Pop(which int) int {
	i := 0
	third := s.Capacity / 3

	switch which {
	case 1:
		s.One--
		i = s.One
	case 2:
		s.Two--
		i = third + s.Two
	case 3:
		s.Three--
		i = (third * 2) + s.Three
	}

	popped := s.Storage[i]

	shrunkThird := (s.Capacity / 3) / 3
	if s.One < shrunkThird && s.Two < shrunkThird && s.Three < shrunkThird {
		s.Shrink()
	}

	return popped
}

func TestPush(t *testing.T) {
	stacks := New()

	// Testing Stack 1
	stacks.Push(1, 1)
	stacks.Push(1, 2)
	stacks.Push(1, 3)
	for i := 0; i < 3; i++ {
		expected := i + 1
		got := stacks.Storage[i]

		if got != expected {
			t.Errorf("Pop error: got %v, expected %v", got, expected)
		}
	}

	// Testing Stack 2
	stacks.Push(2, 11)
	stacks.Push(2, 12)
	stacks.Push(2, 13)
	for i := 10; i < 13; i++ {
		expected := i + 1
		got := stacks.Storage[i]

		if got != expected {
			t.Errorf("Pop error: got %v, expected %v", got, expected)
		}
	}

	// Testing Stack 3
	stacks.Push(3, 21)
	stacks.Push(3, 22)
	stacks.Push(3, 23)
	for i := 20; i < 23; i++ {
		expected := i + 1
		got := stacks.Storage[i]

		if got != expected {
			t.Errorf("Pop error: got %v, expected %v", got, expected)
		}
	}
}
