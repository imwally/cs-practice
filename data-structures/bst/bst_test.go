package bst

import "testing"

func TestInsert(t *testing.T) {
	root := New(10)

	values := []int{8, 12, 100, 1, 45, 6}

	for _, v := range values {
		root.Insert(v)
	}

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

	expected = 100
	got = root.Right.Right.Value
	if got != expected {
		t.Errorf("insert error: got %v, expected %v", got, expected)
	}

	expected = 1
	got = root.Left.Left.Value
	if got != expected {
		t.Errorf("insert error: got %v, expected %v", got, expected)
	}

	expected = 45
	got = root.Right.Right.Left.Value
	if got != expected {
		t.Errorf("insert error: got %v, expected %v", got, expected)
	}

	expected = 6
	got = root.Left.Left.Right.Value
	if got != expected {
		t.Errorf("insert error: got %v, expected %v", got, expected)
	}
}

func TestFind(t *testing.T) {
	root := New(10)

	values := []int{8, 12, 100, 1, 45, 6}

	for _, v := range values {
		root.Insert(v)
	}

	for _, v := range values {
		expected := v
		got := root.Find(v).Value
		if got != expected {
			t.Errorf("find error: got %v, expected %v", got, expected)
		}
	}

	got := root.Find(99)
	if got != nil {
		t.Errorf("find error: got %v, expected nil", got)
	}

}

func TestPreOrder(t *testing.T) {
	root := New(10)

	values := []int{8, 12, 100, 1, 45, 6}
	for _, v := range values {
		root.Insert(v)
	}

	expected := []int{10, 8, 1, 6, 12, 100, 45}
	for i, got := range PreOrder(root) {
		if got != expected[i] {
			t.Errorf("PreOrder error: got %v, expected %v", got, expected[i])
		}
	}
}

func TestInOrder(t *testing.T) {
	root := New(10)

	values := []int{8, 12, 100, 1, 45, 6}
	for _, v := range values {
		root.Insert(v)
	}

	expected := []int{1, 6, 8, 10, 12, 45, 100}
	for i, got := range InOrder(root) {
		if got != expected[i] {
			t.Errorf("InOrder error: got %v, expected %v", got, expected[i])
		}
	}
}

func TestPostOrder(t *testing.T) {
	root := New(10)

	values := []int{8, 12, 100, 1, 45, 6}
	for _, v := range values {
		root.Insert(v)
	}

	expected := []int{6, 1, 8, 45, 100, 12, 10}
	for i, got := range PostOrder(root) {
		if got != expected[i] {
			t.Errorf("PostOrder error: got %v, expected %v", got, expected[i])
		}
	}
}
