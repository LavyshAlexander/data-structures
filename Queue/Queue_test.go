package Queue

import "testing"

func TestLength(t *testing.T) {
	q := New[int]()

	if q.Length() != 0 {
		t.Error("Length of empty queue should be 0.")
	}

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	actualLength := q.Length()
	expectedLength := 3

	if actualLength != expectedLength {
		t.Errorf("Equal length %v is not expected actual %v", actualLength, expectedLength)
	}

	q.Dequeue()

	actualLength = q.Length()
	expectedLength = 2

	if actualLength != expectedLength {
		t.Errorf("Equal length %v is not expected actual %v", actualLength, expectedLength)
	}
}

func TestEnqueue(t *testing.T) {
	q := New[int]()

	firstEnqueue := 1
	q.Enqueue(firstEnqueue)
	topValue := q.first.value

	if topValue != firstEnqueue {
		t.Errorf("Top value %v is not equal %v", topValue, firstEnqueue)
	}

	secondEnqueue := 2
	q.Enqueue(secondEnqueue)
	topValue = q.first.value

	if topValue != firstEnqueue {
		t.Errorf("Top value %v is not equal %v", topValue, secondEnqueue)
	}
}

func TestDenqueue(t *testing.T) {
	q := New[int]()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	last := q.Dequeue()
	if last != 1 {
		t.Error("Dequeued item on queue should be 1")
	}

	last = q.Dequeue()
	if last != 2 {
		t.Error("Dequeued item on queue should be 2")
	}

	last = q.Dequeue()
	if last != 3 {
		t.Error("Dequeued item on queue should be 3")
	}

	if q.first != nil || q.last != nil {
		t.Errorf("Fully dequeued queue has not nilled pointers first - %v and last - %v", q.first, q.last)
	}
}

func TestPeek(t *testing.T) {
	q := New[int]()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	last := q.Peek()
	if last != 1 {
		t.Error("Peeked item on queue should be 1")
	}
	q.Dequeue()

	last = q.Peek()
	if last != 2 {
		t.Error("Peeked item on queue should be 2")
	}
	q.Dequeue()

	last = q.Peek()
	if last != 3 {
		t.Error("Peeked item on queue should be 3")
	}
}

func TestIsEmpty(t *testing.T) {
	q := New[int]()

	if !q.IsEmpty() {
		t.Error("Empty queue should return true on IsEmpty.")
	}

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	if q.IsEmpty() {
		t.Error("Not empty queue should return false on IsEmpty.")
	}
}

func TestString(t *testing.T) {
	q := New[int]()

	str := q.String()
	if str != "[]" {
		t.Errorf("String should return correct value on empty queue. Returned value %v", str)
	}

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	str = q.String()
	if str != "[ 1 2 3 ]" {
		t.Errorf("String should return correct value on not empty queue. Returned value %v", str)
	}
}
