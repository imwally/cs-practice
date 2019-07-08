package main

import "testing"

func find(x int, slice []int) int {
	i, j := 0, len(slice)-1

	for i <= j {
		mid := int((i + j) / 2)
		if x == slice[mid] {
			return slice[mid]
		}
		if x < slice[mid] {
			j = mid - 1
		} else {
			i = mid + 1
		}
	}

	return -1
}

func TestFind(t *testing.T) {
	var slice []int
	for i := 1; i <= 1000000; i++ {
		slice = append(slice, i)
	}

	got := find(8270, slice)
	expected := 8270
	if got != expected {
		t.Errorf("find error: got %d, expected %d\n", got, expected)
	}

	got = find(1, slice)
	expected = 1
	if got != expected {
		t.Errorf("find error: got %d, expected %d\n", got, expected)
	}

	got = find(999999, slice)
	expected = 999999
	if got != expected {
		t.Errorf("find error: got %d, expected %d\n", got, expected)
	}

	got = find(1000001, slice)
	expected = -1
	if got != expected {
		t.Errorf("find error: got %d, expected %d\n", got, expected)
	}
}
