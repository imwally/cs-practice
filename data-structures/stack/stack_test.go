package stack

import "testing"

func TestPush(t *testing.T) {
	s := new(Stack)

	for i := 1; i <= len(s.Storage); i++ {
		if s.Storage[i] != i {
			t.Errorf("Push failed: %d does not match %d", i, s.Storage[i])
		}
	}
}
