package Stack

import "github.com/LavyshAlexander/data-structures/LinkedList"

type Stack[T comparable] struct {
	list *LinkedList.LinkedList[T]
}

func (s *Stack[T]) Push(value T) {
	s.list.Prepend(value)
}

func (s *Stack[T]) Pop() (ret T) {
	popped := s.list.Head
	if popped != nil {
		s.list.Head = popped.Next
		ret = popped.Value
	}

	return
}

func (s *Stack[T]) IsEmpty() bool {
	return s.list.Head == nil
}

func (s *Stack[T]) Peek() T {
	if s.IsEmpty() {
		return *new(T)
	}

	return s.list.Head.Value
}

func (s *Stack[T]) String() string {
	return s.list.String()
}
