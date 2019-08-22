package main

import "testing"

func swap(s []int, i, j int) {
	s[i], s[j] = s[j], s[i]
}

func partition(s []int, lo, hi int) int {
	pivot := s[hi]
	i, j := lo, hi
	for i < hi {
		if s[i] > pivot {
			swap(s, i, j)
			j = i

		}
		i++
	}

	return i
}

func quickSort(s []int, lo, hi int) {
	if lo < hi {
		p := partition(s, lo, hi)
		quickSort(s, lo, p-1)
		quickSort(s, p+1, hi)
	}
}

func QuickSort(s []int) {
	quickSort(s, 0, len(s)-1)
}

func TestQuickSort(t *testing.T) {
	var s []int

	for i := 10000; i > 0; i-- {
		s = append(s, i)
	}

	QuickSort(s)

	for i := 0; i < 10000; i++ {
		expected := i + 1
		got := s[i]

		if got != expected {
			t.Errorf("QuickSort error: got %v, expected %v", got, expected)
		}
	}
}
