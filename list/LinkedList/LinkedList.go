package LinkedList

import (
	"fmt"
)

type LinkedListNode[T comparable] struct {
	Value T
	Next  *LinkedListNode[T]
}

type LinkedList[T comparable] struct {
	Head *LinkedListNode[T]
	Tail *LinkedListNode[T]
}

func (l *LinkedList[T]) Append(value T) {
	node := &LinkedListNode[T]{Value: value}

	if l.Head == nil {
		l.Head = node
		l.Tail = node

		return
	}

	l.Tail.Next = node
	l.Tail = node
}

func (l *LinkedList[T]) Prepend(value T) {
	node := &LinkedListNode[T]{Value: value}

	node.Next = l.Head
	l.Head = node

	if l.Tail == nil {
		l.Tail = node
	}
}

func (l *LinkedList[T]) Delete(value T) bool {
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
		}

		return true
	}

	prev := node
	node = node.Next
	for node != nil {
		if node.Value == value {
			prev.Next = node.Next
			return true
		}

		prev = node
		node = node.Next
	}

	return false
}

func (l *LinkedList[T]) Exist(value T) bool {
	node := l.Head

	for node != nil {
		if node.Value == value {
			return true
		}

		node = node.Next
	}

	return false
}

func (l *LinkedList[T]) Length() int {
	node := l.Head
	length := 0

	for node != nil {
		length++
		node = node.Next
	}

	return length
}

func (l *LinkedList[T]) String() string {
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

func (l *LinkedList[T]) PrintReverse() {
	current := l.Head
	for current != nil {
		defer fmt.Println(current.Value)
		current = current.Next
	}
}

func (l *LinkedList[T]) Reverse() *LinkedList[T] {
	reversed := &LinkedList[T]{}

	current := l.Head
	for current != nil {
		reversed.Prepend(current.Value)
		current = current.Next
	}

	return reversed
}
