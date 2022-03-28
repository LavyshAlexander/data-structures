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

func TestHeapHasIndexes(t *testing.T) {
	comparer := func(a, b int) bool { return a <= b }

	h := New(comparer)
	h.size = 3

	hasLeftChild := h.hasLeftChild(0)
	if !hasLeftChild {
		t.Error("Left child does not exist while should be.")
	}

	hasLeftChild = h.hasLeftChild(100)
	if hasLeftChild {
		t.Error("Left child does exist while should not be")
	}

	hasRightChild := h.hasRightChild(0)
	if !hasRightChild {
		t.Error("Right child does not exist while should be.")
	}

	hasRightChild = h.hasRightChild(100)
	if hasRightChild {
		t.Error("Right child does exist while should not be")
	}

	hasParent := h.hasParent(2)
	if !hasParent {
		t.Error("Parent does not exist while should be.")
	}

	hasParent = h.hasParent(0)
	if hasLeftChild {
		t.Error("Parent does exist while should not be")
	}
}

func TestHeapGetNode(t *testing.T) {
	comparer := func(a, b int) bool { return a <= b }

	h := New(comparer)
	h.size = 3
	h.items = []int{1, 2, 3, 0, 0, 0, 0, 0, 0, 0}

	leftChild := h.leftChild(0)
	if leftChild != 2 {
		t.Errorf("Left child has unexpected value %v, should be %v.", leftChild, 2)
	}

	rightChild := h.rightChild(0)
	if rightChild != 3 {
		t.Errorf("Right child has unexpected value %v, should be %v.", rightChild, 3)
	}

	parent := h.parent(1)
	if parent != 1 {
		t.Errorf("Parent has unexpected value %v, should be %v.", parent, 1)
	}

	parent = h.parent(2)
	if parent != 1 {
		t.Errorf("Parent has unexpected value %v, should be %v.", parent, 1)
	}
}