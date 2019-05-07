// Given an image represented by an NxN matrix, where each pixel in
// the image is 4 bytes, write a method to rotate the image by 90
// degrees. Can you do this in place?

package main

import (
	"fmt"
	"testing"
)

const SIZE = 4

type matrix [SIZE][SIZE]int

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

	fmt.Println(m)
	m = RotateMatrix(m)
	fmt.Println(m)
}
