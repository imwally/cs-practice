// Exercise 3.4
//
// Queue via Stacks
//
// Implement a MyQueue class which implements a queue using two stacks.

package main

import (
	"fmt"
	"testing"

	"../data-structures/stack"
)

type Queue struct {
	InStack  stack.Stack
	OutStack stack.Stack
}

func (q *Queue) Enqueue(v int) {
	q.InStack.Push(v)
}

func (q *Queue) Dequeue() int {
	for q.InStack.Size() > 0 {
		q.OutStack.Push(q.InStack.Pop())
	}

	v := q.OutStack.Pop()

	for q.OutStack.Size() > 0 {
		q.InStack.Push(q.OutStack.Pop())
	}

	return v
}

func TestQueueFromStacks(t *testing.T) {
	q := &Queue{}

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)

	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
}
