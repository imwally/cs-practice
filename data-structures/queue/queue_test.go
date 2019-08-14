package queue

import (
	"fmt"
	"testing"
)

func TestEnqueue(t *testing.T) {
	q := New()

	for i := 1; i <= 10; i++ {
		q.Enqueue(i)
	}

	for i, v := range q.Storage {
		expected := i + 1
		got := v

		if got != expected {
			t.Errorf("Enqueue error: got %v, expected %v", got, expected)
		}
	}
}

func TestDequeue(t *testing.T) {
	q := New()

	for i := 1; i <= 10; i++ {
		q.Enqueue(i)
	}

	for i := 1; i <= 10; i++ {
		expected := i
		got := q.Dequeue()

		if got != expected {
			t.Errorf("Dequeue error: got %v, expected %v", got, expected)
		}
	}
}

func TestLength(t *testing.T) {
	q := New()

	for i := 1; i <= 10; i++ {
		q.Enqueue(i)
	}

	expected := 10
	got := q.Length()

	if got != expected {
		t.Errorf("Length error: got %v, expected %v", got, expected)
	}
}

func TestString(t *testing.T) {
	q := New()

	for i := 1; i <= 10; i++ {
		q.Enqueue(i)
	}

	expected := "10, 9, 8, 7, 6, 5, 4, 3, 2, 1"
	got := fmt.Sprintf("%s", q)

	if got != expected {
		t.Errorf("String error: got %v, expected %v", got, expected)
	}
}
