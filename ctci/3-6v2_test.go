package main

import (
	"testing"
	"time"

	"../data-structures/linkedlist"
)

type Animal struct {
	Type  string
	Name  string
	Added time.Time
}

type Queue struct {
	Dogs linkedlist.LinkedList
	Cats linkedlist.LinkedList
}

func (q *Queue) Enqueue(a Animal) {
	a.Added = time.Now()
	if a.Type == "Cat" {
		q.Cats.Append(a)
	}
	if a.Type == "Dog" {
		q.Dogs.Append(a)
	}
}

func (q *Queue) DequeueDog() *Animal {
	a := q.Dogs.Head.Value.(Animal)
	q.Dogs.Delete(0)

	return &a
}
func (q *Queue) DequeueCat() *Animal {
	a := q.Cats.Head.Value.(Animal)
	q.Cats.Delete(0)

	return &a
}

func (q *Queue) DequeueAny() *Animal {
	c, d := q.Cats.Head.Value.(Animal), q.Dogs.Head.Value.(Animal)

	if c.Added.Before(d.Added) {
		return q.DequeueCat()
	} else {
		return q.DequeueDog()
	}

	return nil
}

func TestAnimalQueue(t *testing.T) {
	q := new(Queue)

	q.Enqueue(Animal{"Dog", "Tux", time.Time{}})
	q.Enqueue(Animal{"Dog", "Billie", time.Time{}})
	q.Enqueue(Animal{"Cat", "Angel", time.Time{}})
	q.Enqueue(Animal{"Dog", "Tank", time.Time{}})
	q.Enqueue(Animal{"Cat", "Alex", time.Time{}})
	q.Enqueue(Animal{"Cat", "Pheobe", time.Time{}})

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
