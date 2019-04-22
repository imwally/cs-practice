// URLify: Write a method to replace all spaces in a string with '%20:
// You may assume that the string has sufficient space at the end to
// hold the additional characters, and that you are given the "true"
// length of the string.

package main

import "fmt"

func URLify(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes)-1; i++ {
		if runes[i] == ' ' {
			copy(runes[i+3:], runes[i+1:])
			runes[i] = '%'
			runes[i+1] = '2'
			runes[i+2] = '0'
			i = i + 2
		}
	}

	return string(runes)
}

func main() {
	str1 := "Hello World  "
	str2 := "A much longer test string with a lot more spaces                        "
	fmt.Println(URLify(str1))
	fmt.Println(URLify(str2))
}
