package main

import (
	"math/rand"
	"testing"
)

func swap(s []int, i, j int) {
	s[i], s[j] = s[j], s[i]
}

func BubbleSort(s []int) {
	var swapped bool
	for i := 0; i < len(s)-1; {
		if s[i+1] < s[i] {
			swap(s, i, i+1)
			swapped = true
		}

		if i+1 == len(s)-1 && swapped {
			i = 0
			swapped = false
		} else {
			i++
		}
	}
}

func TestBubbleSort(t *testing.T) {
	s := rand.Perm(1000)
	for i, _ := range s {
		s[i]++
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
