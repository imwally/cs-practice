package hashtable

const ARRAYSIZE int = 16

type Record struct {
	Key   string
	Value string
	Next  *Record
}

type Table struct {
	Storage [ARRAYSIZE]*Record
}

func hash(s string) int {
	sum := 0
	for _, c := range s {
		sum += int(c) * 31
	}

	return sum % ARRAYSIZE
}

func (t *Table) Insert(key, s string) {
	h := hash(key)
	record := &Record{key, s, nil}

	for r := t.Storage[h]; r != nil; r = r.Next {
		if r.Next == nil {
			r.Next = record
			return
		}
	}

	t.Storage[h] = record
}

func (t *Table) Get(key string) string {
	h := hash(key)

	for r := t.Storage[h]; r != nil; r = r.Next {
		if r.Key == key {
			return r.Value
		}
	}

	return ""
}
