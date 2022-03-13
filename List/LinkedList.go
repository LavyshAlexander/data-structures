package list

import (
	"fmt"
)

type LinkedListNode struct {
	Value int
	Next  *LinkedListNode
}

type LinkedList struct {
	Head *LinkedListNode
	Tail *LinkedListNode
}

func (l *LinkedList) Append(value int) {
	node := &LinkedListNode{Value: value}

	if l.Head == nil {
		l.Head = node
		l.Tail = node

		return
	}

	l.Tail.Next = node
	l.Tail = node
}

func (l *LinkedList) Prepend(value int) {
	node := &LinkedListNode{Value: value}

	node.Next = l.Head
	l.Head = node

	if l.Tail == nil {
		l.Tail = node
	}
}

func (l *LinkedList) Delete(value int) bool {
	node := l.Head

	if node == nil {
		return false
	}

	if node.Value == value {
		l.Head = node.Next
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

func (l *LinkedList) String() string {
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
