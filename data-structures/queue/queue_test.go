package queue

import "testing"

func TestEnqueue(t *testing.T) {
	q := New()

	for i := 1; i <= 10; i++ {
		q.Enqueue(i)
	}

	for i := 1; i <= 10; i++ {
		expected := i
		got := q.Dequeue()

		if got != expected {
			t.Errorf("Enqueue error: got %v, expected %v", got, expected)
		}
	}
}
