package bst

import "testing"

func TestInsert(t *testing.T) {
	root := New(10)
	root.Insert(8)
	root.Insert(12)

	expected := 10
	got := root.Value
	if got != expected {
		t.Errorf("insert error: got %v, expected %v", got, expected)
	}

	expected = 8
	got = root.Left.Value
	if got != expected {
		t.Errorf("insert error: got %v, expected %v", got, expected)
	}

	expected = 12
	got = root.Right.Value
	if got != expected {
		t.Errorf("insert error: got %v, expected %v", got, expected)
	}
}

func TestFind(t *testing.T) {
	root := New(10)
	root.Insert(8)
	root.Insert(12)
	root.Insert(100)
	root.Insert(1)
	root.Insert(45)
	root.Insert(6)

	expected := 100
	got := root.Find(100).Value
	if got != expected {
		t.Errorf("find error: got %v, expected %v", got, expected)
	}
}
