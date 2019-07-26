package array

import "testing"

func TestAppend(t *testing.T) {
	a := New()
	for i := 1; i <= 10; i++ {
		a.Append(i)
	}

	for i := 0; i < 10; i++ {
		expected := i + 1
		got := a.Data[i]

		if expected != got {
			t.Errorf("append error: got %v, expected  %v", got, expected)
		}
	}
}

func TestPop(t *testing.T) {
	a := New()
	for i := 1; i <= 10; i++ {
		a.Append(i)
	}

	for i := 10; i > 0; i-- {
		expected := i
		got := a.Pop()

		if expected != got {
			t.Errorf("append error: got %v, expected  %v", got, expected)
		}
	}
}
