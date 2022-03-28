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

	h.Add(25)
	h.Add(20)
	h.Add(100)
	h.Add(123)

	if fmt.Sprint(h.items) != "[1 10 5 25 20 100 123 0 0 0]" {
		t.Errorf("Items after Add has wrong value: %v", h.items)
	}

	h.Add(42)
	h.Add(200)
	h.Add(300)

	if fmt.Sprint(h.items) != "[1 10 5 25 20 100 123 42 200 300]" {
		t.Errorf("Items after Add has wrong value: %v", h.items)
	}

	h.Add(400)

	if fmt.Sprint(h.items) != "[1 10 5 25 20 100 123 42 200 300 400 0 0 0 0 0 0 0 0 0]" {
		t.Errorf("Items after Add has wrong value: %v", h.items)
	}

	h.Add(0)

	if fmt.Sprint(h.items) != "[0 10 1 25 20 5 123 42 200 300 400 100 0 0 0 0 0 0 0 0]" {
		t.Errorf("Items after Add has wrong value: %v", h.items)
	}

	h.Add(101)
	h.Add(124)
	h.Add(125)

	if fmt.Sprint(h.items) != "[0 10 1 25 20 5 123 42 200 300 400 100 101 124 125 0 0 0 0 0]" {
		t.Errorf("Items after Add has wrong value: %v", h.items)
	}

	h.Add(23)

	if fmt.Sprint(h.items) != "[0 10 1 23 20 5 123 25 200 300 400 100 101 124 125 42 0 0 0 0]" {
		t.Errorf("Items after Add has wrong value: %v", h.items)
	}
}
