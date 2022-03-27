package Heap

import "testing"

func TestNew(t *testing.T) {
	comparer := func(a, b int) bool { return a <= b }

	h := New(comparer)

	if h.Size() != 0 {
		t.Error("Size of newly created Heap is not zero.")
	}

	if h.Capacity() != 10 {
		t.Error("Capacity of newly created Heap is not 10.")
	}

	if h.compare == nil {
		t.Error("Compare function in newly created Heap is nil.")
	}
}
