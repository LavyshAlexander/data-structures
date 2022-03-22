package Stack

import "fmt"

type (
	Stack[T any] struct {
		top    *Node[T]
		length int
	}
	Node[T any] struct {
		value    T
		previous *Node[T]
	}
)

func New[T comparable]() *Stack[T] {
	return &Stack[T]{nil, 0}
}

func (s *Stack[T]) Length() int {
	return s.length
}

func (s *Stack[T]) IsEmpty() bool {
	return s.top == nil
}

func (s *Stack[T]) Peek() T {
	if s.IsEmpty() {
		return *new(T)
	}

	return s.top.value
}

func (s *Stack[T]) Push(value T) {
	top := &Node[T]{value, s.top}
	s.top = top
}

func (s *Stack[T]) Pop() (ret T) {
	popped := s.top

	if popped != nil {
		s.top = popped.previous
		return popped.value
	}

	return
}

func (s *Stack[T]) String() string {
	if s.top == nil {
		return "[]"
	}

	node := s.top

	str := "[ "
	for node != nil {
		str += fmt.Sprintf("%v ", node.value)
		node = node.previous
	}
	str += "]"

	return str
}
