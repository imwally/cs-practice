// Is Unique: Implement an algorithm to determine if a string has all
// unique characters. What if you cannot use additional data
// structures?
package main

import "fmt"

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

func main() {
	fmt.Println(IsUnique1("foobar"))
	fmt.Println(IsUnique1("abcdefghijklmnopqrstuvwxyz"))
	fmt.Println("---")
	fmt.Println(IsUnique2("foobar"))
	fmt.Println(IsUnique2("abcdefghijklmnopqrstuvwxyz"))
	fmt.Println("---")
	fmt.Println(IsUnique3("foobar"))
	fmt.Println(IsUnique3("abcdefghijklmnopqrstuvwxyz"))
}
