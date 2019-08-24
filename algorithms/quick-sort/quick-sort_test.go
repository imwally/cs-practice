package main

import "testing"

func swap(s []int, i, j int) {
	s[i], s[j] = s[j], s[i]
}

func partition(s []int, lo, hi int) int {
	p := s[lo+(hi-lo)/2]

	for {
		for s[lo] < p {
			lo++
		}

		for s[hi] > p {
			hi--
		}

		if lo >= hi {
			return hi
		}

		swap(s, lo, hi)

		lo++
		hi--
	}

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
	var s []int
	for i := size; i > 0; i-- {
		s = append(s, i)
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
