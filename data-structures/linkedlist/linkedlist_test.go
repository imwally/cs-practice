package linkedlist

import (
	"fmt"
	"testing"
)

func TestAppend(t *testing.T) {
	l := new(LinkedList)
	l.Append(1)

	if l.Head.Value != 1 {
		t.Error("Append failed: head should be 1")
	}

	if l.Tail.Value != 1 {
		t.Error("Append failed: tail should be 1")
	}

	l.Append(2)

	if l.Head.Value != 1 {
		t.Error("Append failed: head should be 1")
	}

	if l.Tail.Value != 2 {
		t.Error("Append failed: tail should be 2")
	}
}

func TestInsert(t *testing.T) {
	l := new(LinkedList)

	for i := 1; i < 5; i++ {
		l.Append(i)
	}

	l.Insert(0, 10)

	printed := fmt.Sprintf("%s", l)
	expected := "10 -> 1 -> 2 -> 3 -> 4"
	if printed != expected {
		t.Errorf("Insert failed: expected \"%s\", got \"%s\"", expected, printed)
	}

	if l.Head.Value != 10 {
		t.Errorf("Insert failed: expected \"%v\", got \"%v\"", 10, l.Head.Value)
	}

	l.Insert(4, 11)
	printed = fmt.Sprintf("%s", l)
	expected = "10 -> 1 -> 2 -> 3 -> 11 -> 4"
	if printed != expected {
		t.Errorf("Insert failed: expected \"%s\", got \"%s\"", expected, printed)
	}

	if l.Tail.Value != 4 {
		t.Errorf("Insert failed: expected \"%v\", got \"%v\"", 4, l.Tail.Value)
	}

	l.Insert(6, 100)
	printed = fmt.Sprintf("%s", l)
	expected = "10 -> 1 -> 2 -> 3 -> 11 -> 4 -> 100"
	if printed != expected {
		t.Errorf("Insert failed: expected \"%s\", got \"%s\"", expected, printed)
	}

	if l.Tail.Value != 100 {
		t.Errorf("Insert failed: expected \"%v\", got \"%v\"", 100, l.Tail.Value)
	}
}

func TestString(t *testing.T) {
	l := new(LinkedList)

	for i := 1; i < 5; i++ {
		l.Append(i)
	}

	printed := fmt.Sprintf("%s", l)
	expected := "1 -> 2 -> 3 -> 4"
	if printed != expected {
		t.Errorf("String failed: expected \"%s\", got \"%s\"", expected, printed)
	}
}

func TestSize(t *testing.T) {
	l := new(LinkedList)

	for i := 1; i < 5; i++ {
		l.Append(i)
	}

	size := l.Size()
	expected := 4
	if size != expected {
		t.Errorf("Size failed: expected %d, got %d", expected, size)
	}
}
