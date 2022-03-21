package Queue

import (
	"github.com/LavyshAlexander/data-structures/DoublyLinkedList"
)

type Queue[T comparable] struct {
	list *DoublyLinkedList.DoublyLinkedList[T]
}

func (q *Queue[T]) Push(value T) {
	q.list.Append(value)
}

func (q *Queue[T]) Pop() (ret T) {
	popped := q.list.Tail
	if popped != nil {
		q.list.Tail = popped.Previous
		ret = popped.Value
	}

	return
}

func (s *Queue[T]) IsEmpty() bool {
	return s.list.Head == nil
}

func (s *Queue[T]) Peek() T {
	if s.IsEmpty() {
		return *new(T)
	}

	return s.list.Tail.Value
}

func (s *Queue[T]) String() string {
	return s.list.String()
}
