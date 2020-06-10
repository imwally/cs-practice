package main

import (
	"fmt"
	"testing"

	"../../data-structures/linkedlist"
)

type Animal struct {
	Type string
	Name string
}

type Queue struct {
	Animals linkedlist.LinkedList
}

func (q *Queue) Enqueue(t interface{}) {
	q.Animals.Append(t)
}

func (q *Queue) dequeueType(t string) *Animal {
	if q.Animals.Head.Value.(Animal).Type == t {
		return q.DequeueAny()
	}

	for c := q.Animals.Head; c.Next != nil; c = c.Next {
		next := c.Next.Value.(Animal)
		if next.Type == t {
			c.Next = c.Next.Next

			return &next
		}
	}

	return nil
}

func (q *Queue) DequeueDog() *Animal {
	return q.dequeueType("Dog")
}
func (q *Queue) DequeueCat() *Animal {
	return q.dequeueType("Cat")
}

func (q *Queue) DequeueAny() *Animal {
	a := q.Animals.Head.Value.(Animal)
	q.Animals.Delete(0)

	return &a
}

func (q *Queue) String() string {
	s := ""
	for c := q.Animals.Head; c != nil; c = c.Next {
		s += fmt.Sprintf("%v", c.Value)
		if c.Next != nil {
			s += fmt.Sprintf(" -> ")
		}
	}

	return s
}

func TestAnimalQueue(t *testing.T) {
	q := new(Queue)

	q.Enqueue(Animal{"Dog", "Tux"})
	q.Enqueue(Animal{"Dog", "Billie"})
	q.Enqueue(Animal{"Cat", "Angel"})
	q.Enqueue(Animal{"Dog", "Tank"})
	q.Enqueue(Animal{"Cat", "Alex"})
	q.Enqueue(Animal{"Cat", "Pheobe"})

	got, expected := q.DequeueCat().Name, "Angel"
	if got != expected {
		t.Errorf("error: got %v, expected %v", got, expected)
	}

	got, expected = q.DequeueCat().Name, "Alex"
	if got != expected {
		t.Errorf("error: got %v, expected %v", got, expected)
	}

	got, expected = q.DequeueAny().Name, "Tux"
	if got != expected {
		t.Errorf("error: got %v, expected %v", got, expected)
	}

	got, expected = q.DequeueDog().Name, "Billie"
	if got != expected {
		t.Errorf("error: got %v, expected %v", got, expected)
	}

	got, expected = q.DequeueCat().Name, "Pheobe"
	if got != expected {
		t.Errorf("error: got %v, expected %v", got, expected)
	}
}
