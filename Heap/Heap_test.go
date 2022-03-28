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

func TestSwap(t *testing.T) {
	comparer := func(a, b int) bool { return a <= b }

	h := New(comparer)
	h.size = 3
	h.items = []int{1, 2, 3, 0, 0, 0, 0, 0, 0, 0}

	first := h.items[1]
	second := h.items[2]

	h.swap(1, 2)

	if first != h.items[2] {
		t.Errorf("After swap first element has unexpected value %v, should be %v.", first, second)
	}

	if second != h.items[1] {
		t.Errorf("After swap second element has unexpected value %v, should be %v.", second, first)
	}
}

func TestEnsureExtraCapacity(t *testing.T) {
	comparer := func(a, b int) bool { return a <= b }

	h := New(comparer)
	h.size = 10
	h.items = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	itemsExpected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	h.ensureExtraCapacity()

	if h.capacity != 20 {
		t.Errorf("Capacity value %v is not equal to expected %v.", h.capacity, 20)
	}

	for i, v := range h.items {
		if v != itemsExpected[i] {
			t.Errorf("Items has unexpected value %v instead of %v on index %d.", v, itemsExpected[i], i)
		}
	}
}

func TestPeek(t *testing.T) {
	comparer := func(a, b int) bool { return a <= b }

	h := New(comparer)

	peeked := h.Peek()
	if peeked != 0 {
		t.Errorf("Peeked value %v is not equal to expected %v.", peeked, 0)
	}

	h.size = 10
	h.items = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	peeked = h.Peek()
	if peeked != 1 {
		t.Errorf("Peeked value %v is not equal to expected %v.", peeked, 0)
	}
}

func TestAdd(t *testing.T) {
	comparer := func(a, b int) bool { return a < b }

	h := New(comparer)
	h.size = 4
	h.items = []int{10, 15, 20, 17, 0, 0, 0, 0, 0, 0}
	itemsExpected := []int{8, 10, 20, 17, 15, 0, 0, 0, 0, 0}

	h.Add(8)

	if h.size != 5 {
		t.Errorf("After adding new element size %v didn't changed to %v.", h.size, 5)
	}

	for i, v := range h.items {
		if v != itemsExpected[i] {
			t.Errorf("Items has unexpected value %v instead of %v on index %d.", v, itemsExpected[i], i)
		}
	}
}
