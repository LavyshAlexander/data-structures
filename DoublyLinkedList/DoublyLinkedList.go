package DoublyLinkedList

import "fmt"

type DoublyLinkedListNode[T comparable] struct {
	Value    T
	Previous *DoublyLinkedListNode[T]
	Next     *DoublyLinkedListNode[T]
}

type DoublyLinkedList[T comparable] struct {
	Head *DoublyLinkedListNode[T]
	Tail *DoublyLinkedListNode[T]
}

func (l *DoublyLinkedList[T]) Append(value T) {
	node := &DoublyLinkedListNode[T]{Value: value}

	if l.Head == nil {
		l.Head = node
		l.Tail = node
	} else {
		node.Previous = l.Tail
		l.Tail.Next = node
		l.Tail = node
	}
}

func (l *DoublyLinkedList[T]) Prepend(value T) {
	node := &DoublyLinkedListNode[T]{Value: value}

	if l.Head == nil {
		l.Head = node
		l.Tail = node
	} else {
		node.Next = l.Head
		l.Head.Previous = node
		l.Head = node
	}
}

func (l *DoublyLinkedList[T]) Delete(value T) bool {
	node := l.Head

	if node == nil {
		return false
	}

	if node.Value == value {
		if l.Head == l.Tail {
			l.Head = nil
			l.Tail = nil
		} else {
			l.Head = node.Next
			node.Previous = nil
		}

		return true
	}

	node = node.Next

	for node != nil {
		if node.Value == value {
			node.Previous.Next = node.Next
			node.Next.Previous = node.Previous
			return true
		}

		node = node.Next
	}

	return false
}

func (l *DoublyLinkedList[T]) Exist(value T) bool {
	node := l.Head

	for node != nil {
		if node.Value == value {
			return true
		}

		node = node.Next
	}

	return false
}

func (l *DoublyLinkedList[T]) Length() int {
	length := 0
	node := l.Head

	for node != nil {
		length++
		node = node.Next
	}

	return length
}

func (l *DoublyLinkedList[T]) String() string {
	if l.Head == nil {
		return "[]"
	}

	node := l.Head

	str := "[ "
	for node != nil {
		str += fmt.Sprintf("%v ", node.Value)
		node = node.Next
	}
	str += "]"

	return str
}
