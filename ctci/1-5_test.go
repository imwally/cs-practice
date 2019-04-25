// String Compression: Implement a method to perform basic string
// compression using the counts of repeated characters. For example,
// the string aabcccccaaa would become a2b1c5a3. If the "compressed"
// string would not become smaller than the original string, your
// method should return the original string. You can assume the string
// has only uppercase and lowercase letters (a - z).

package main

import (
	"fmt"
	"testing"
)

func CompressString(s string) string {
	var compressed string
	seen := make(map[rune]int)

	for i, char := range s {
		seen[char] += 1

		if i == len(s)-1 || rune(s[i+1]) != char {
			compressed += fmt.Sprintf("%c%d", char, seen[char])
			seen = make(map[rune]int)
		}
	}

	return compressed
}

func TestCompressedString(t *testing.T) {
	str1 := "aabcccccaaa"
	ans1 := "a2b1c5a3"
	compressed := CompressString(str1)

	if compressed != ans1 {
		t.Errorf("%s did not compress to %s, got %s instead", str1, ans1, compressed)
	}
}