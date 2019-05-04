// One Away: There are three types of edits that can be performed on
// strings: insert a character, remove a character, or replace a
// character. Given two strings, write a function to check if they are
// one edit (or zero edits) away.

package main

import "testing"

func OneAway(s1, s2 string) bool {
	i := 0
	for i < len(s1) && i < len(s2) {
		if s1[:i] == s2[:i] && s1[i+1:] == s2[i+1:] {
			return true
		}
		i++
	}

	if (len(s1)-i) == 1 || (len(s2)-i) == 1 {
		return true
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
