// Given an image represented by an NxN matrix, where each pixel in
// the image is 4 bytes, write a method to rotate the image by 90
// degrees. (an you do this in place?

package main

import (
	"fmt"
	"testing"
)

func RotateMatrix(m [][]int) {
	for i := 0; i < len(m)-1; i++ {
		for j := 0; j < len(m)-1; j++ {
			temp := m[i][j+1]
			m[i][j+1] = m[i+1][j]
			m[i+1][j] = temp
		}
	}
}

func TestRotateMatrix(t *testing.T) {

	m := [][]int{
		{1, 2},
		{3, 4},
	}

	fmt.Println(m)
	RotateMatrix(m)
	fmt.Println(m)
}
