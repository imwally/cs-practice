package main

import (
	"math/rand"
	"testing"
)

func Merge(a, b []int) []int {
	merged := make([]int, len(a)+len(b))

	i, j, k := 0, 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			merged[k] = a[i]
			i++
		} else {
			merged[k] = b[j]
			j++
		}
		k++
	}

	copy(merged[k:], a[i:])
	copy(merged[k:], b[j:])

	return merged
}

func MergeSort(s []int) []int {
	if len(s) <= 1 {
		return s
	}

	mid := int(len(s) / 2)

	return Merge(MergeSort(s[:mid]), MergeSort(s[mid:]))
}

func TestMergeSort(t *testing.T) {
	size := 10000
	s := rand.Perm(size)
	for i := 0; i < size; i++ {
		s[i]++
	}

	sorted := MergeSort(s)
	for i := 0; i < size; i++ {
		expected := i + 1
		got := sorted[i]

		if got != expected {
			t.Errorf("MergeSort error: got %v, expected %v", got, expected)
		}
	}
}
