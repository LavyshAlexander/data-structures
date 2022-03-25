package Stack

import "testing"

func TestLength(t *testing.T) {
	s := New[int]()

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
