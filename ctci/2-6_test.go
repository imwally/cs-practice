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

	ll2 := &linkedlist.LinkedList{}
	for _, r2 := range "racecar" {
		ll.Append(int(r2))
	}

	expected = true
	got = Palindrome(ll2)

	if got != expected {
		t.Errorf("Palindrome error: got %v, expected %v", got, expected)
	}

	ll3 := &linkedlist.LinkedList{}
	for _, r3 := range "wally" {
		ll3.Append(int(r3))
	}

	expected = false
	got = Palindrome(ll3)

	if got != expected {
		t.Errorf("Palindrome error: got %v, expected %v", got, expected)
	}

	ll4 := &linkedlist.LinkedList{}
	for _, r4 := range "aba" {
		ll4.Append(int(r4))
	}

	expected = true
	got = Palindrome(ll4)

	if got != expected {
		t.Errorf("Palindrome error: got %v, expected %v", got, expected)
	}

	ll5 := &linkedlist.LinkedList{}
	for _, r5 := range "aaa" {
		ll5.Append(int(r5))
	}

	expected = true
	got = Palindrome(ll5)

	if got != expected {
		t.Errorf("Palindrome error: got %v, expected %v", got, expected)
	}

}
