package Heap

import (
	"fmt"
	"testing"
)

func TestMaxAdd(t *testing.T) {
	h := NewMax[int]()

	h.Add(10)
	h.Add(1)
	h.Add(5)

	if fmt.Sprint(h.items) != "[10 1 5 0 0 0 0 0 0 0]" {
		t.Errorf("Items after Add has wrong value: %v", h.items)
	}

	h.Add(25)
	h.Add(20)
	h.Add(100)
	h.Add(123)

	if fmt.Sprint(h.items) != "[123 20 100 1 10 5 25 0 0 0]" {
		t.Errorf("Items after Add has wrong value: %v", h.items)
	}

	h.Add(42)
	h.Add(200)
	h.Add(300)

	if fmt.Sprint(h.items) != "[300 200 100 42 123 5 25 1 20 10]" {
		t.Errorf("Items after Add has wrong value: %v", h.items)
	}

	h.Add(400)

	if fmt.Sprint(h.items) != "[400 300 100 42 200 5 25 1 20 10 123 0 0 0 0 0 0 0 0 0]" {
		t.Errorf("Items after Add has wrong value: %v", h.items)
	}

	h.Add(0)

	if fmt.Sprint(h.items) != "[400 300 100 42 200 5 25 1 20 10 123 0 0 0 0 0 0 0 0 0]" {
		t.Errorf("Items after Add has wrong value: %v", h.items)
	}

	h.Add(101)
	h.Add(124)
	h.Add(125)

	if fmt.Sprint(h.items) != "[400 300 125 42 200 100 124 1 20 10 123 0 5 25 101 0 0 0 0 0]" {
		t.Errorf("Items after Add has wrong value: %v", h.items)
	}

	h.Add(23)

	if fmt.Sprint(h.items) != "[400 300 125 42 200 100 124 23 20 10 123 0 5 25 101 1 0 0 0 0]" {
		t.Errorf("Items after Add has wrong value: %v", h.items)
	}
}
