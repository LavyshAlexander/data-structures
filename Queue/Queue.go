package Queue

import "fmt"

type (
	Queue[T any] struct {
		first  *Node[T]
		last   *Node[T]
		length int
	}

	Node[T any] struct {
		value T
		next  *Node[T]
	}
)

func New[T comparable]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Length() int {
	return q.length
}

func (q *Queue[T]) Enqueue(value T) {
	node := &Node[T]{value, nil}

	if q.first == nil {
		q.first = node
		q.last = node
	} else {
		q.last.next = node
		q.last = node
	}

	q.length++
}

func (q *Queue[T]) Dequeue() (ret T) {
	if q.first == nil {
		return
	}

	dequeued := q.first

	if q.length == 1 {
		q.first = nil
		q.last = nil
	} else {
		q.first = dequeued.next
	}

	q.length--
	ret = dequeued.value
	return
}

func (q *Queue[T]) IsEmpty() bool {
	return q.length == 0
}

func (q *Queue[T]) Peek() T {
	if q.IsEmpty() {
		return *new(T)
	}

	return q.first.value
}

func (q *Queue[T]) String() string {
	if q.IsEmpty() {
		return "[]"
	}

	str := "[ "
	current := q.first
	for current != nil {
		str += fmt.Sprintf("%v ", current.value)
		current = current.next
	}
	str += "]"

	return str
}
