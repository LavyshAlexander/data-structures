package Queue

import (
	"github.com/LavyshAlexander/data-structures/LinkedList"
)

// Implement without LinkedList usage
type (
	Queue[T comparable] struct {
		list *LinkedList.LinkedList[T]
	}
)

func New[T comparable]() *Queue[T] {
	return &Queue[T]{list: &LinkedList.LinkedList[T]{}}
}

func (q *Queue[T]) Enqueue(value T) {
	q.list.Append(value)
}

func (q *Queue[T]) Dequeue() (ret T) {
	popped := q.list.Head

	if popped != nil {
		q.list.Head = popped.Next
		ret = popped.Value
	}

	return
}

func (q *Queue[T]) IsEmpty() bool {
	return q.list.Head == nil
}

func (q *Queue[T]) Peek() T {
	if q.IsEmpty() {
		return *new(T)
	}

	return q.list.Head.Value
}

func (q *Queue[T]) String() string {
	return q.list.String()
}
