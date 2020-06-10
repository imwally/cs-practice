// Check Permutation: Given two strings, write a method to decide if
// one is a permutation of the other.
package main

import "testing"

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

func TestPermutation(t *testing.T) {
	s1 := "abcdef"
	s2 := "fedcba"
	s3 := "foobar"
	s4 := "barfoo"
	s5 := "barfooo"

	if Permutation(s1, s2) != true {
		t.Errorf("%s is not a permutation of %s", s1, s2)
	}

	if Permutation(s3, s4) != true {
		t.Errorf("%s is not a permutation of %s", s3, s4)
	}

	if Permutation(s4, s5) != false {
		t.Errorf("%s should not be a permutation of %s", s4, s5)
	}
}
