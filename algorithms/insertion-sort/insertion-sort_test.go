package main

import "testing"

func InsertionSort(s []int) {
	for i := 0; i < len(s); i++ {
		j := i
		for j > 0 && s[j-1] > s[j] {
			s[j], s[j-1] = s[j-1], s[j]
			j--
		}
	}
}

func TestInsertionSort(t *testing.T) {
	s := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	InsertionSort(s)

	for i, v := range s {
		if v != expected[i] {
			t.Errorf("InsertionSort error: got %d, expected %d", v, expected[i])
		}
	}
}
