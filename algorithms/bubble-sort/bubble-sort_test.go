package main

import "testing"

func swap(s []int, i, j int) {
	s[i], s[j] = s[j], s[i]
}

func BubbleSort(s []int) {
	for i := 0; i < len(s)-1; {
		if s[i+1] < s[i] {
			swap(s, i, i+1)
			i = 0
		} else {
			i++
		}
	}
}

func TestBubbleSort(t *testing.T) {
	var s []int

	for i := 1000; i > 0; i-- {
		s = append(s, i)
	}

	BubbleSort(s)

	for i := 0; i < 1000; i++ {
		expected := i + 1
		got := s[i]

		if got != expected {
			t.Errorf("BubbleSort error: got %v, expected %v", got, expected)
		}
	}

}
