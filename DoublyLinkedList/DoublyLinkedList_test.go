package DoublyLinkedList

import (
	"testing"
)

func TestAppend(t *testing.T) {
	list := DoublyLinkedList[int]{}

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
		t.Errorf("in tail got %v, wanted %v", tail, 30)
	}
}

func TestPrepend(t *testing.T) {
	list := DoublyLinkedList[int]{}

	list.Prepend(10)
	list.Prepend(20)
	list.Prepend(30)

	first := list.Head.Value
	second := list.Head.Next.Value
	third := list.Head.Next.Next.Value
	end := list.Head.Next.Next.Next
	tail := list.Tail.Value

	firstExpected := 30
	if first != firstExpected {
		t.Errorf("in first node got %v, wanted %v", first, firstExpected)
	}

	secondExpected := 20
	if second != secondExpected {
		t.Errorf("in second node got %v, wanted %v", second, secondExpected)
	}

	thirdExpected := 10
	if third != thirdExpected {
		t.Errorf("in third node got %v, wanted %v", third, thirdExpected)
	}

	if end != nil {
		t.Errorf("in end got %v, wanted %v", end, nil)
	}

	if tail != thirdExpected {
		t.Errorf("in tail got %v, wanted %v", tail, thirdExpected)
	}
}

func TestLength(t *testing.T) {
	list := DoublyLinkedList[int]{}

	list.Append(10)
	list.Append(20)
	list.Append(30)

	length := list.Length()

	if length != 3 {
		t.Errorf("got length %v, wanted %v", length, 3)
	}

	list.Delete(20)
	length = list.Length()

	if length != 2 {
		t.Errorf("got length %v, wanted %v", length, 3)
	}
}

func TestDelete(t *testing.T) {
	list := DoublyLinkedList[int]{}

	list.Append(10)
	list.Append(20)
	list.Append(30)
	list.Delete(20)

	first := list.Head.Value
	second := list.Head.Next.Value
	end := list.Head.Next.Next
	tail := list.Tail.Value

	firstExpected := 10
	if first != firstExpected {
		t.Errorf("in first node got %v, wanted %v", first, firstExpected)
	}

	secondExpected := 30
	if second != secondExpected {
		t.Errorf("in second node got %v, wanted %v", second, secondExpected)
	}

	if end != nil {
		t.Errorf("in end got %v, wanted %v", end, nil)
	}

	if tail != secondExpected {
		t.Errorf("in tail got %v, wanted %v", tail, secondExpected)
	}
}

func TestDeleteSingleHead(t *testing.T) {
	list := DoublyLinkedList[int]{}

	list.Append(10)
	list.Delete(10)

	head := list.Head
	if head != nil {
		t.Errorf("in head node got %v, wanted %v", head, nil)
	}

	tail := list.Tail
	if list.Tail != nil {
		t.Errorf("in tail node got %v, wanted %v", tail, nil)
	}
}

func TestExist(t *testing.T) {
	list := DoublyLinkedList[int]{}

	list.Append(10)
	list.Append(20)
	list.Append(30)

	firstExpected := 10
	isFirstExist := list.Exist(firstExpected)
	if !isFirstExist {
		t.Errorf("got %v, wanted %v", isFirstExist, true)
	}

	tailExpected := 30
	isTailExist := list.Exist(tailExpected)
	if !isTailExist {
		t.Errorf("got %v, wanted %v", isTailExist, true)
	}

	notExistedValue := 42
	nonExistedValueCheck := list.Exist(notExistedValue)
	if nonExistedValueCheck {
		t.Errorf("got %v, wanted %v", nonExistedValueCheck, false)
	}
}

func TestString(t *testing.T) {
	list := DoublyLinkedList[int]{}

	list.Append(10)
	list.Append(20)
	list.Append(30)

	stringValue := list.String()
	expectedString := "[ 10 20 30 ]"

	if stringValue != expectedString {
		t.Errorf("got %v, wanted %v", stringValue, expectedString)
	}

	list.Delete(20)
	stringValue = list.String()
	expectedString = "[ 10 30 ]"

	if stringValue != expectedString {
		t.Errorf("got %v, wanted %v", stringValue, expectedString)
	}

	emptyString := DoublyLinkedList[int]{}

	stringValue = emptyString.String()
	expectedString = "[]"

	if stringValue != expectedString {
		t.Errorf("got %v, wanted %v", stringValue, expectedString)
	}
}
