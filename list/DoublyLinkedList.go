package list

import "fmt"

type DoublyLinkedListNode struct {
	Value    int
	Previous *DoublyLinkedListNode
	Next     *DoublyLinkedListNode
}

type DoublyLinkedList struct {
	Head *DoublyLinkedListNode
	Tail *DoublyLinkedListNode
}

func (l *DoublyLinkedList) Append(value int) {

}

func (l *DoublyLinkedList) Prepend(value int) {

}

func (l *DoublyLinkedList) Delete(value int) bool {

	return false
}

func (l *DoublyLinkedList) Exist(value int) bool {

	return false
}

func (l *DoublyLinkedList) Length() int {
	length := 0

	return length
}

func (l *DoublyLinkedList) String() string {
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
