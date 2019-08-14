package queue

import "testing"

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
