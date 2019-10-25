// Exercise 2.6
//
// Palindrome: Implement a function to check if a linked list is a palindrome.

package main

import (
	"testing"

	"../data-structures/linkedlist"
	"../data-structures/stack"
)

func Palindrome(ll *linkedlist.LinkedList) bool {
	s := &stack.Stack{}
	for i, c := 0, ll.Head; c != nil; i, c = i+1, c.Next {
		if i < ll.Size()/2 {
			s.Push(c.Value.(int))
		}
		if i > ll.Size()/2 {
			if s.Pop() != c.Value.(int) {
				return false
			}
		}
	}

	return true
}

func TestPalindrome(t *testing.T) {
	ll := &linkedlist.LinkedList{}
	for _, r := range "a toyotas a toyota" {
		if r != 32 {
			ll.Append(int(r))
		}
	}

	expected := true
	got := Palindrome(ll)

	if got != expected {
		t.Errorf("Palindrome error: got %v, expected %v", got, expected)
	}
}
