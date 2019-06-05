package trie

import (
	"testing"
)

func TestInsert(t *testing.T) {
	t1 := new(Trie)
	t1.Insert("he")
	t1.Insert("hell")
	t1.Insert("hello")
	t1.Insert("helloworld")

	if !t1.Exists("he") {
		t.Error("Insert failed: he doesn't exist")
	}

	if t1.Exists("help") {
		t.Error("Insert failed: help shouldn't exist")
	}

	if !t1.Exists("hell") {
		t.Error("Insert failed: hell doesn't exist")
	}

	if !t1.Exists("hello") {
		t.Error("Insert failed: hello doesn't exist")
	}

	if !t1.Exists("helloworld") {
		t.Error("Insert failed: helloworld doesn't exist")
	}
}
