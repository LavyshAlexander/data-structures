package Heap

import (
	"fmt"
	"testing"
)

func TestMinAdd(t *testing.T) {
	h := NewMin[int]()

	h.Add(10)
	h.Add(1)
	h.Add(5)

	if fmt.Sprint(h.items) != "[1 10 5 0 0 0 0 0 0 0]" {
		t.Errorf("Items after Add has wrong value: %v", h.items)
	}
}
