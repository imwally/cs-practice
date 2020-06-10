// Exercise 2.5
//
// Sum Lists: You have two numbers represented by a linked list, where each
// node contains a single digit. The digits are stored in reverse order, such
// that the 1's digit is at the head of the list. Write a function that adds
// the two numbers and returns the sum as a linked list.
//
// EXAMPLE:
//
// Input: (7 -> 1 -> 6) + (5 -> 9 -> 2).
// That is, 617 + 295.
// Output: 2 -> 1 -> 9.
// That is, 912.

package main

import (
	"fmt"
	"strconv"
	"testing"

	"../../data-structures/linkedlist"
	"../../data-structures/stack"
)

func SumList(ll1, ll2 *linkedlist.LinkedList) *linkedlist.LinkedList {
	s1 := &stack.Stack{}
	for c1 := ll1.Head; c1 != nil; c1 = c1.Next {
		s1.Push(c1.Value.(int))
	}

	s2 := &stack.Stack{}
	for c2 := ll2.Head; c2 != nil; c2 = c2.Next {
		s2.Push(c2.Value.(int))
	}

	numString1 := ""
	for s1.Size() > 0 {
		numString1 += strconv.Itoa(s1.Pop())
	}

	numString2 := ""
	for s2.Size() > 0 {
		numString2 += strconv.Itoa(s2.Pop())
	}

	num1, _ := strconv.Atoi(numString1)
	num2, _ := strconv.Atoi(numString2)

	sum := num1 + num2

	sumStack := &stack.Stack{}
	for _, v := range strconv.Itoa(sum) {
		sumNum, _ := strconv.Atoi(string(v))
		sumStack.Push(sumNum)
	}

	sumList := &linkedlist.LinkedList{}
	for sumStack.Size() > 0 {
		sumList.Append(sumStack.Pop())
	}

	return sumList
}

func TestSumList(t *testing.T) {
	ll1 := &linkedlist.LinkedList{}
	ll1.Append(7)
	ll1.Append(1)
	ll1.Append(6)

	ll2 := &linkedlist.LinkedList{}
	ll2.Append(5)
	ll2.Append(9)
	ll2.Append(2)

	expected := "2 -> 1 -> 9"
	got := fmt.Sprint(SumList(ll1, ll2))

	if got != expected {
		t.Errorf("SumList error: got %s, expected %s", got, expected)
	}
}
