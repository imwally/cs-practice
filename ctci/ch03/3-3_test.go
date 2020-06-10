// Exercise 3.3
//
// Stack of Plates
//
// Imagine a (literal) stack of plates. If the stack gets too high, it might
// topple. Therefore, in real life, we would likely start a new stack when the
// previous stack exceeds some threshold. Implement a data structure
// SetOfStacks that mimics this. SetOfStacks should be composed of several
// stacks and should create a new stack once the previous one exceeds capacity.
// SetOfStacks. push() and SetOfStacks. pop() should behave identically to a
// single stack (that is, pop() should return the same values as it would if
// there were just a single stack).

package main

import (
	"testing"

	"../../data-structures/stack"
)

type SetOfStacks struct {
	Stacks []*stack.Stack
	Index  int
}

func (s *SetOfStacks) Push(v int) {
	currentStack := s.Stacks[s.Index]

	if currentStack.Size() == 10 {
		s.Index++
		s.Stacks = append(s.Stacks, &stack.Stack{})
		currentStack = s.Stacks[s.Index]
	}

	currentStack.Push(v)
}

func (s *SetOfStacks) Pop() int {
	if s.Stacks[s.Index].Size() == 0 {
		s.Index--
	}

	return s.Stacks[s.Index].Pop()
}

func New() *SetOfStacks {
	var stacks []*stack.Stack
	stacks = append(stacks, &stack.Stack{})
	return &SetOfStacks{
		stacks,
		0,
	}
}

func TestSetOfStacks(t *testing.T) {
	s := New()

	for i := 1; i < 1001; i++ {
		s.Push(i)
	}

	for j := 1000; j > 0; j-- {
		expected := j
		got := s.Pop()

		if got != expected {
			t.Errorf("error: got %v, expected %v", got, expected)
		}
	}

	expected := 100
	got := len(s.Stacks)

	if got != expected {
		t.Errorf("error: got %v, expected %v", got, expected)
	}

}
