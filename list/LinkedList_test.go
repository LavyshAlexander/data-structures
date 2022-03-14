package list

import (
	"testing"
)

func TestAppend(t *testing.T) {
	list := LinkedList{}

	list.Append(10)
	list.Append(20)
	list.Append(30)

	first := list.Head.Value
	second := list.Head.Next.Value
	third := list.Head.Next.Next.Value
	end := list.Head.Next.Next.Next
	tail := list.Tail.Value

	if first != 10 {
		t.Errorf("in first node got %v, wanted %v", first, 10)
	}

	if second != 20 {
		t.Errorf("in second node got %v, wanted %v", second, 20)
	}

	if third != 30 {
		t.Errorf("in third node got %v, wanted %v", third, 30)
	}

	if end != nil {
		t.Errorf("in end got %v, wanted %v", end, nil)
	}

	if tail != 30 {
		t.Errorf("in tail got %v, wanted %v", first, 30)
	}
}

func TestLength(t *testing.T) {
	list := LinkedList{}

	list.Append(10)
	list.Append(20)
	list.Append(30)

	length := list.Length()

	if length != 3 {
		t.Errorf("got length %v, wanted %v", length, 3)
	}
}
