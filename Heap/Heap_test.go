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

func TestHeapGetIndexes(t *testing.T) {
	comparer := func(a, b int) bool { return a <= b }

	h := New(comparer)

	leftIndex := h.getLeftChildIndex(0)
	leftIndexExpexted := 1

	if leftIndex != leftIndexExpexted {
		t.Errorf("Left index calculated with error: %v, but expected %v", leftIndex, leftIndexExpexted)
	}

	leftIndex = h.getLeftChildIndex(100)
	leftIndexExpexted = 201

	if leftIndex != leftIndexExpexted {
		t.Errorf("Left index calculated with error: %v, but expected %v", leftIndex, leftIndexExpexted)
	}

	rightIndex := h.getRightChildIndex(0)
	rightIndexExpexted := 2

	if rightIndex != rightIndexExpexted {
		t.Errorf("Right index calculated with error: %v, but expected %v", rightIndex, rightIndexExpexted)
	}

	rightIndex = h.getRightChildIndex(100)
	rightIndexExpexted = 202

	if rightIndex != rightIndexExpexted {
		t.Errorf("Right index calculated with error: %v, but expected %v", rightIndex, rightIndexExpexted)
	}

	parentIndex := h.getParentIndex(0)
	parentIndexExpexted := 0

	if parentIndex != parentIndexExpexted {
		t.Errorf("Parent index calculated with error: %v, but expected %v", parentIndex, parentIndexExpexted)
	}

	parentIndex = h.getParentIndex(101)
	parentIndexExpexted = 50

	if parentIndex != parentIndexExpexted {
		t.Errorf("Parent index calculated with error: %v, but expected %v", parentIndex, parentIndexExpexted)
	}

	parentIndex = h.getParentIndex(102)
	parentIndexExpexted = 50

	if parentIndex != parentIndexExpexted {
		t.Errorf("Parent index calculated with error: %v, but expected %v", parentIndex, parentIndexExpexted)
	}
}
