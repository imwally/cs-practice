// URLify: Write a method to replace all spaces in a string with '%20:
// You may assume that the string has sufficient space at the end to
// hold the additional characters, and that you are given the "true"
// length of the string.

package main

import "testing"

func URLify(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes)-1; i++ {
		if runes[i] == ' ' {
			// Shift everything after the space ahead by
			// two characters leaving enough room for "%20".
			copy(runes[i+3:], runes[i+1:])
			copy(runes[i:i+3], []rune("%20"))
			i = i + 2
		}
	}

	return string(runes)
}

func TestURLify(t *testing.T) {
	str1 := "Hello World  "
	ans1 := "Hello%20World"

	str2 := "A much longer test string with a lot more spaces                  "
	ans2 := "A%20much%20longer%20test%20string%20with%20a%20lot%20more%20spaces"

	str3 := "One more for good measure        "
	ans3 := "One%20more%20for%20good%20measure"

	if ans1 != URLify(str1) {
		t.Errorf("Expected \"%s\" but got \"%s\"", ans1, URLify(str1))
	}

	if ans2 != URLify(str2) {
		t.Errorf("Expected \"%s\" but got \"%s\"", ans2, URLify(str2))
	}

	if ans3 != URLify(str3) {
		t.Errorf("Expected \"%s\" but got \"%s\"", ans3, URLify(str3))
	}
}
