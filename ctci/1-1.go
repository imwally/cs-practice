// Is Unique: Implement an algorithm to determine if a string has all
// unique characters. What if you cannot use additional data
// structures?

package main

import "fmt"

// Using additional data structures
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

func main() {
	fmt.Println(IsUnique1("foobar"))
	fmt.Println(IsUnique1("abcdefghijklmnopqrstuvwxyz"))

	fmt.Println(IsUnique2("foobar"))
	fmt.Println(IsUnique2("abcdefghijklmnopqrstuvwxyz"))
}
