// Exercise 3.5
//
// Sort Stack
//
// Write a program to sort a stack such that the smallest items are
// on the top. You can use an additional temporary stack, but you may not copy
// the elements into any other data structure (such as an array). The stack
// supports the following operations: push, pop, peek, and isEmpty.

package main

import (
	"testing"

	"../../data-structures/stack"
)

// Uses two additional stacks...
func SortStack(s *stack.Stack) *stack.Stack {
	sorted := &stack.Stack{}
	origSize := s.Size()
	for i := 0; i < origSize; i++ {
		max := 0
		for s.Size() > 0 {
			if s.Peek() > max {
				max = s.Peek()
			}
			sorted.Push(s.Pop())

		}

		for j := 0; j < origSize-i; j++ {
			if sorted.Peek() == max {
				sorted.Pop()
			} else {
				s.Push(sorted.Pop())
			}
		}

		sorted.Push(max)
	}

	return sorted
}

func TestSortStack(t *testing.T) {
	s := &stack.Stack{}
	for i := 1; i < 1001; i++ {
		s.Push(i)
	}

	s = SortStack(s)
	for i := 1; i < 1001; i++ {
		expected := i
		got := s.Pop()

		if got != expected {
			t.Errorf("error: got %v, expected %v", got, expected)
		}
	}
}
