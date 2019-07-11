package main

import "testing"

func FindRecursive(x int, slice []int) int {
	mid := len(slice) / 2
	if slice[mid] == x {
		return slice[mid]
	}

	if mid > 0 {
		if x < slice[mid] {
			return FindRecursive(x, slice[:mid])
		} else {
			return FindRecursive(x, slice[mid:])
		}
	}

	return -1
}

func FindIterative(x int, slice []int) int {
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

func TestFindIterative(t *testing.T) {
	var slice []int
	for i := 1; i <= 1000000; i++ {
		slice = append(slice, i)
	}

	got := FindIterative(8270, slice)
	expected := 8270
	if got != expected {
		t.Errorf("find error: got %d, expected %d\n", got, expected)
	}

	got = FindIterative(1, slice)
	expected = 1
	if got != expected {
		t.Errorf("find error: got %d, expected %d\n", got, expected)
	}

	got = FindIterative(999999, slice)
	expected = 999999
	if got != expected {
		t.Errorf("find error: got %d, expected %d\n", got, expected)
	}

	got = FindIterative(1000001, slice)
	expected = -1
	if got != expected {
		t.Errorf("find error: got %d, expected %d\n", got, expected)
	}
}

func TestFindRecursive(t *testing.T) {
	var slice []int
	for i := 1; i <= 1000000; i++ {
		slice = append(slice, i)
	}

	got := FindRecursive(8270, slice)
	expected := 8270
	if got != expected {
		t.Errorf("find error: got %d, expected %d\n", got, expected)
	}

	got = FindRecursive(1, slice)
	expected = 1
	if got != expected {
		t.Errorf("find error: got %d, expected %d\n", got, expected)
	}

	got = FindRecursive(999999, slice)
	expected = 999999
	if got != expected {
		t.Errorf("find error: got %d, expected %d\n", got, expected)
	}

	got = FindRecursive(1000001, slice)
	expected = -1
	if got != expected {
		t.Errorf("find error: got %d, expected %d\n", got, expected)
	}
}
