package main

import (
	"math/rand"
	"testing"
)

func swap(s []int, i, j int) {
	s[i], s[j] = s[j], s[i]
}

func partition(s []int, lo, hi int) int {
	// Set pivot to mid point
	p := s[lo+(hi-lo)/2]

	// Essentially the low end and high end indices close in on each other
	// until an inverted pair is found. Those elements are then
	// finally swapped.
	for {
		// Walk foward through slice incrementing the lo index until an
		// element that is greater than the pivot is found.
		for s[lo] < p {
			lo++
		}

		// Walk backwards through the slice decrementing the hi index
		// until an element that is lower than the pivot is found.
		for s[hi] > p {
			hi--
		}

		// Return the hi index when indices meet.
		if lo >= hi {
			return hi
		}

		// If the indices don't meet swap elements at the lo and hi
		// index. This should be an inverted pair. The element at index
		// lo is greater than the pivot and the element at index hi is
		// less than the pivot.
		swap(s, lo, hi)

		// Increment lo and decrement hi, otherwise the loop would run
		// forever.
		lo++
		hi--
	}

	// This should never be reached. It's only here to satisfy the return
	// type as Go requires a return at the end of a function.
	return 0
}

func quickSort(s []int, lo, hi int) {
	if lo < hi {
		p := partition(s, lo, hi)
		quickSort(s, lo, p)
		quickSort(s, p+1, hi)
	}
}

func QuickSort(s []int) {
	quickSort(s, 0, len(s)-1)
}

func TestQuickSort(t *testing.T) {
	size := 10000
	s := rand.Perm(size)
	for i := 0; i < size; i++ {
		s[i]++
	}

	QuickSort(s)

	for i := 0; i < size; i++ {
		expected := i + 1
		got := s[i]

		if got != expected {
			t.Errorf("QuickSort error: got %v, expected %v", got, expected)
		}
	}
}
