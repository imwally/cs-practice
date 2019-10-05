package main

import (
	"math/rand"
	"testing"
)

func InsertionSort(s []int) {
	// Loop over every item in the slice
	for i := 0; i < len(s); i++ {
		// Starting at the i'th element and working backwards, if the
		// i-1 element is larger, swap them
		for j := i; j > 0 && s[j-1] > s[j]; j-- {
			s[j], s[j-1] = s[j-1], s[j]
		}
	}
}

func TestInsertionSort(t *testing.T) {
	size := 10000
	s := rand.Perm(size)

	expected := make([]int, size)
	for i, _ := range s {
		expected[i] = i
	}

	InsertionSort(s)

	for i, v := range s {
		if v != expected[i] {
			t.Errorf("InsertionSort error: got %d, expected %d", v, expected[i])
		}
	}
}
