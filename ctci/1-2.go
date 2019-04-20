// Check Permutation: Given two strings, write a method to decide if
// one is a permutation of the other.
package main

import "fmt"

func Permutation(s1, s2 string) bool {
	seen := make(map[rune]int)

	for _, r1 := range s1 {
		seen[r1] += 1
	}

	for _, r2 := range s2 {
		seen[r2] -= 1
	}

	for _, times := range seen {
		if times != 0 {
			return false
		}
	}

	return true
}

func main() {

	s1 := "abcdef"
	s2 := "fedcba"
	s3 := "foobar"
	s4 := "barfoo"

	fmt.Println(Permutation(s1, s2))
	fmt.Println(Permutation(s2, s3))
	fmt.Println(Permutation(s1, s3))
	fmt.Println(Permutation(s3, s4))
}
