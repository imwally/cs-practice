// Exercise 3.1
//
// Three in One
//
// Describe how you could use a single array to implement three stacks.

package main

import (
	"testing"
)

type Stacks struct {
	Storage  []int
	Capacity int
	Third    int
	One      int
	Two      int
	Three    int
}

func New() *Stacks {
	cap := 30
	third := cap / 3
	storage := make([]int, cap)
	return &Stacks{
		Storage:  storage,
		Capacity: cap,
		Third:    third,
		One:      0,
		Two:      0,
		Three:    0,
	}
}

func (s *Stacks) Shrink() {
	s.Capacity = s.Capacity / 2
	newThird := s.Capacity / 3

	newStorage := make([]int, s.Capacity)

	// Copy first third
	copy(newStorage[:newThird], s.Storage[:s.Third])

	// Copy middle
	copy(newStorage[newThird:newThird*2], s.Storage[s.Third:s.Third*2])

	// Copy last third
	copy(newStorage[newThird*2:], s.Storage[s.Third*2:])

	s.Storage = newStorage
	s.Third = newThird
}

func (s *Stacks) Grow() {
	s.Capacity = s.Capacity * 2
	newThird := s.Capacity / 3

	newStorage := make([]int, s.Capacity)

	// Copy first third
	copy(newStorage[:newThird], s.Storage[:s.Third])

	// Copy middle
	copy(newStorage[newThird:newThird*2], s.Storage[s.Third:s.Third*2])

	// Copy last third
	copy(newStorage[newThird*2:], s.Storage[s.Third*2:])

	s.Storage = newStorage
	s.Third = newThird
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
			i = s.Third + s.Two
		}
		s.Two++
	case 3:
		i = third*2 + s.Three
		if i == s.Capacity {
			s.Grow()
			i = s.Third*2 + s.Three
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

	shrunkThird := (s.Capacity / 2) / 3
	if s.One < shrunkThird && s.Two < shrunkThird && s.Three < shrunkThird {
		s.Shrink()
	}

	return popped
}

func (s *Stacks) Size(which int) int {
	i := 0
	switch which {
	case 1:
		i = s.One
	case 2:
		i = s.Two
	case 3:
		i = s.Three
	}

	return i
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

func TestPop(t *testing.T) {
	stacks := New()

	stacks.Push(1, 1)
	stacks.Push(1, 2)
	stacks.Push(1, 3)
	stacks.Push(1, 4)
	for i := 4; stacks.Size(1) > 0; i-- {
		expected := i
		got := stacks.Pop(1)

		if got != expected {
			t.Errorf("Pop error: got %v, expected %v", got, expected)
		}
	}

	stacks.Push(2, 5)
	stacks.Push(2, 6)
	stacks.Push(2, 7)
	stacks.Push(2, 8)
	for i := 8; stacks.Size(2) > 0; i-- {
		expected := i
		got := stacks.Pop(2)

		if got != expected {
			t.Errorf("Pop error: got %v, expected %v", got, expected)
		}
	}

	stacks.Push(3, 9)
	stacks.Push(3, 10)
	stacks.Push(3, 11)
	stacks.Push(3, 12)
	for i := 12; stacks.Size(3) > 0; i-- {
		expected := i
		got := stacks.Pop(3)

		if got != expected {
			t.Errorf("Pop error: got %v, expected %v", got, expected)
		}
	}

}

func TestSize(t *testing.T) {
	stacks := New()

	for i := 1; i < 5; i++ {
		stacks.Push(1, i)
	}

	expected := 4
	got := stacks.Size(1)

	if got != expected {
		t.Errorf("Size error: got %v, expected %v", got, expected)
	}
}
