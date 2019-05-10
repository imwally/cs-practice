// Given an image represented by an NxN matrix, where each pixel in
// the image is 4 bytes, write a method to rotate the image by 90
// degrees. Can you do this in place?

package main

import "testing"

const SIZE = 4

type matrix [SIZE][SIZE]int

// This is not done in place and uses a fixed size matrix.
func RotateMatrix(m matrix) matrix {
	w := len(m)
	var rotated matrix
	for i := 0; i < w; i++ {
		for j := 0; j < w; j++ {
			rotated[i][j] = m[w-j-1][i]
		}
	}

	return rotated
}

func TestRotateMatrix(t *testing.T) {
	m := matrix{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	answer := matrix{
		{13, 9, 5, 1},
		{14, 10, 6, 2},
		{15, 11, 7, 3},
		{16, 12, 8, 4},
	}

	if RotateMatrix(m) != answer {
		t.Errorf("%v should be\n %v after rotation", m, answer)
	}

	m = matrix{
		{1, 2, 3, 4},
		{1, 2, 3, 4},
		{1, 2, 3, 4},
		{1, 2, 3, 4},
	}
	answer = matrix{
		{1, 1, 1, 1},
		{2, 2, 2, 2},
		{3, 3, 3, 3},
		{4, 4, 4, 4},
	}

	if RotateMatrix(m) != answer {
		t.Errorf("%v should be\n %v after rotation", m, answer)
	}
}
