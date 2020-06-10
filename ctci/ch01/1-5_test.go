// One Away: There are three types of edits that can be performed on
// strings: insert a character, remove a character, or replace a
// character. Given two strings, write a function to check if they are
// one edit (or zero edits) away.

package main

import "testing"

func StringWithoutIndex(s string, i int) string {
	if i > len(s) || i < 0 {
		return ""
	}

	return s[:i] + s[i+1:]
}

func OneAway(s1, s2 string) bool {
	// Do either of the strings match if the last character is
	// removed?
	// Ex: pales, pale -> true
	if s1[:len(s1)-1] == s2 || s2[:len(s2)-1] == s1 {
		return true
	}

	for i := 0; i < len(s1) && i < len(s2); i++ {
		// Does the first string without the i'th charater
		// match the second?
		// Ex: pale, ale -> true
		if StringWithoutIndex(s1, i) == s2 {
			return true
		}
		// Same as above if the strings were swapped.
		// Ex: ale, pale -> true
		if StringWithoutIndex(s2, i) == s1 {
			return true
		}
		// Do both match if both i'th characters are removed?
		// Ex: pale, bale -> true
		if StringWithoutIndex(s1, i) == StringWithoutIndex(s2, i) {
			return true
		}
	}

	return false
}

func TestOneAway(t *testing.T) {
	s1 := "pale"
	s2 := "ple"
	if !OneAway(s1, s2) {
		t.Errorf("%s and %s should be true", s1, s2)
	}

	s1 = "pale"
	s2 = "ale"
	if !OneAway(s1, s2) {
		t.Errorf("%s and %s should be true", s1, s2)
	}

	s1 = "pale"
	s2 = "bale"
	if !OneAway(s1, s2) {
		t.Errorf("%s and %s should be true", s1, s2)
	}

	s1 = "pales"
	s2 = "pale"
	if !OneAway(s1, s2) {
		t.Errorf("%s and %s should be true", s1, s2)
	}

	s1 = "pale"
	s2 = "bake"
	if OneAway(s1, s2) {
		t.Errorf("%s and %s should be false", s1, s2)
	}

	s1 = "hello world"
	s2 = "helloworld"
	if !OneAway(s1, s2) {
		t.Errorf("%s and %s should be true", s1, s2)
	}

	s1 = "hello world"
	s2 = "helloworlds"
	if OneAway(s1, s2) {
		t.Errorf("%s and %s should be false", s1, s2)
	}

}
