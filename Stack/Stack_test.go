package Stack

import "testing"

func TestLength(t *testing.T) {
	s := New[int]()

	if s.Length() != 0 {
		t.Error("Length of empty stack should be 0.")
	}

	s.Push(1)
	s.Push(2)
	s.Push(3)

	actualLength := s.Length()
	expectedLength := 3

	if actualLength != expectedLength {
		t.Errorf("Equal length %v is not expected actual %v", actualLength, expectedLength)
	}
}

func TestPush(t *testing.T) {
	s := New[int]()

	firstPush := 1
	s.Push(firstPush)
	topValue := s.top.value

	if topValue != firstPush {
		t.Errorf("Top value %v is not equal %v", topValue, firstPush)
	}

	secondPush := 2
	s.Push(secondPush)
	topValue = s.top.value

	if topValue != secondPush {
		t.Errorf("Top value %v is not equal %v", topValue, secondPush)
	}
}

func TestPop(t *testing.T) {
	s := New[int]()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	last := s.Pop()
	if last != 3 {
		t.Error("Last item on stack should be 3")
	}

	last = s.Pop()
	if last != 2 {
		t.Error("Last item on stack should be 2")
	}

	last = s.Pop()
	if last != 1 {
		t.Error("Last item on stack should be 1")
	}
}

func TestPeek(t *testing.T) {
	s := New[int]()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	last := s.Peek()
	if last != 3 {
		t.Error("Peeked item on stack should be 3")
	}
	s.Pop()

	last = s.Peek()
	if last != 2 {
		t.Error("Peeked item on stack should be 2")
	}
	s.Pop()

	last = s.Peek()
	if last != 1 {
		t.Error("Peeked item on stack should be 1")
	}
}

func TestIsEmpty(t *testing.T) {
	s := New[int]()

	if !s.IsEmpty() {
		t.Error("Empty stack should return true on IsEmpty.")
	}

	s.Push(1)
	s.Push(2)
	s.Push(3)

	if s.IsEmpty() {
		t.Error("Not empty stack should return false on IsEmpty.")
	}
}

func TestString(t *testing.T) {
	s := New[int]()

	str := s.String()
	if str != "[]" {
		t.Errorf("String should return correct value on empty stack. Returned value %v", str)
	}

	s.Push(1)
	s.Push(2)
	s.Push(3)

	str = s.String()
	if str != "[ 3 2 1 ]" {
		t.Errorf("String should return correct value on not empty stack. Returned value %v", str)
	}
}
