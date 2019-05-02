// Is Unique: Implement an algorithm to determine if a string has all
// unique characters. What if you cannot use additional data
// structures?
package main

import "testing"

// Using additional data structures (map or hash table)
func IsUnique1(s string) bool {
	seen := make(map[rune]int)
	for _, r := range s {
		seen[r] += 1
		if seen[r] > 1 {
			return false
		}
	}

	return true
}

// No additional data structures (but runs in O(n^2) time...)
func IsUnique2(s string) bool {
	for i, r1 := range s {
		for _, r2 := range s[i+1:] {
			if r1 == r2 {
				return false
			}
		}
	}

	return true
}

// Is an array considered an additional data structure?
func IsUnique3(s string) bool {
	seen := make([]int, 1024)

	for _, r := range s {
		if seen[int(r)] == 1 {
			return false
		}
		seen[int(r)] = 1
	}

	return true
}

func TestIsUnique1(t *testing.T) {
	s1 := "foobar"
	if IsUnique1(s1) == true {
		t.Errorf("%s is not unique", s1)
	}

	s2 := "abcdefghijklmnopqrstuvwxyz"
	if IsUnique1(s2) != true {
		t.Errorf("%s should be unique", s2)
	}

}

func TestIsUnique2(t *testing.T) {
	s1 := "foobar"
	if IsUnique2(s1) == true {
		t.Errorf("%s is not unique", s1)
	}

	s2 := "abcdefghijklmnopqrstuvwxyz"
	if IsUnique2(s2) != true {
		t.Errorf("%s should be unique", s2)
	}
}

func TestIsUnique3(t *testing.T) {
	s1 := "foobar"
	if IsUnique3(s1) == true {
		t.Errorf("%s is not unique", s1)
	}

	s2 := "abcdefghijklmnopqrstuvwxyz"
	if IsUnique3(s2) != true {
		t.Errorf("%s should be unique", s2)
	}
}
