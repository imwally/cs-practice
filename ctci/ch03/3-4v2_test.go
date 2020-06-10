package main

import (
	"testing"
)

type Queue struct {
	s1, s2 Stack
}

func (q *Queue) Enqueue(v int) {
	q.s1.Push(v)
}

func (q *Queue) Dequeu() int {
	for len(q.s1.Data) > 0 {
		q.s2.Push(q.s1.Pop())
	}

	v := q.s2.Pop()

	for len(q.s2.Data) > 0 {
		q.s1.Push(q.s2.Pop())
	}

	return v
}

type Stack struct {
	Data []int
}

func (s *Stack) Push(v int) {
	s.Data = append(s.Data, v)
}

func (s *Stack) Pop() int {
	v := s.Data[len(s.Data)-1]
	s.Data = s.Data[:len(s.Data)-1]

	return v
}

func TestPush(t *testing.T) {
	s := new(Stack)

	s.Push(4)
	s.Push(3)
	s.Push(2)
	s.Push(1)

	for i := 0; i < len(s.Data); i++ {
		got, expected := s.Data[i], len(s.Data)-i
		if got != expected {
			t.Errorf("push error: got %v, expected %v", got, expected)
		}
	}
}

func TestPop(t *testing.T) {
	s := new(Stack)

	s.Push(4)
	s.Push(3)
	s.Push(2)
	s.Push(1)

	for i := 0; len(s.Data) > 0; i++ {
		got, expected := s.Pop(), i+1
		if got != expected {
			t.Errorf("pop error: got %v, expected %v", got, expected)
		}
	}

	got, expected := len(s.Data), 0
	if got != expected {
		t.Errorf("len error: got %v, expected %v", got, expected)
	}
}

func TestEnqueue(t *testing.T) {
	q := new(Queue)

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)

	for i := 0; i < len(q.s1.Data); i++ {
		got, expected := q.s1.Data[i], i+1
		if got != expected {
			t.Errorf("enqueue error: got %v, expected %v", got, expected)
		}
	}
}

func TestDequeue(t *testing.T) {
	q := new(Queue)

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)

	for i := 0; len(q.s1.Data) > 0; i++ {
		got, expected := q.Dequeu(), i+1
		if got != expected {
			t.Errorf("dequeue error: got %v, expected %v", got, expected)
		}
	}
}
