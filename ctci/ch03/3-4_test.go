// Exercise 3.4
//
// Queue via Stacks
//
// Implement a MyQueue class which implements a queue using two stacks.

package main

import (
	"testing"

	"../../data-structures/stack"
)

type Queue struct {
	InStack  stack.Stack
	OutStack stack.Stack
}

func (q *Queue) Enqueue(v int) {
	for q.OutStack.Size() > 0 {
		q.InStack.Push(q.OutStack.Pop())
	}

	q.InStack.Push(v)
}

func (q *Queue) Dequeue() int {
	for q.InStack.Size() > 0 {
		q.OutStack.Push(q.InStack.Pop())
	}

	v := q.OutStack.Pop()

	return v
}

func TestQueueFromStacks(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected stack underflow")
		}
	}()

	q := &Queue{}

	for i := 1; i < 11; i++ {
		q.Enqueue(i)
	}

	for i := 1; i < 11; i++ {
		expected := i
		got := q.Dequeue()

		if got != expected {
			t.Errorf("error: got %v, expected %v", got, expected)
		}
	}

	q.Enqueue(11)
	q.Enqueue(12)
	q.Enqueue(13)

	expected := 11
	got := q.Dequeue()

	if got != expected {
		t.Errorf("error: got %v, expected %v", got, expected)
	}

	q.Enqueue(14)

	expected = 12
	got = q.Dequeue()

	if got != expected {
		t.Errorf("error: got %v, expected %v", got, expected)
	}

	q.Enqueue(15)
	q.Enqueue(16)

	expected = 13
	got = q.Dequeue()

	if got != expected {
		t.Errorf("error: got %v, expected %v", got, expected)
	}

	q.Enqueue(17)
	q.Enqueue(18)

	expected = 14
	got = q.Dequeue()

	if got != expected {
		t.Errorf("error: got %v, expected %v", got, expected)
	}

	expected = 15
	got = q.Dequeue()

	if got != expected {
		t.Errorf("error: got %v, expected %v", got, expected)
	}

	expected = 16
	got = q.Dequeue()

	if got != expected {
		t.Errorf("error: got %v, expected %v", got, expected)
	}

	expected = 17
	got = q.Dequeue()

	if got != expected {
		t.Errorf("error: got %v, expected %v", got, expected)
	}

	expected = 18
	got = q.Dequeue()

	if got != expected {
		t.Errorf("error: got %v, expected %v", got, expected)
	}

	q.Dequeue()
}
