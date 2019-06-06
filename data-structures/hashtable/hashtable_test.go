package hashtable

import (
	"testing"
)

func TestInsert(t *testing.T) {
	ht := new(Table)
	k, v := "go", "Go Language"
	ht.Insert(k, v)

	var record *Record
	index := hash(k)
	for record = ht.Storage[index]; record != nil; record = record.Next {
		if record.Value == v {
			break
		}
	}

	if record.Value != v {
		t.Errorf("Insert error: %s does not point to %s", k, v)
	}
}
