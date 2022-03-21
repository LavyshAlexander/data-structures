package main

type List[T any] interface {
	Append(value T)
	Prepend(value T)
	Delete(value T) bool
	Exist(value T) bool
	Length() int
}

type Stack[T any] interface {
	Push(value T)
	Pop() T
	IsEmpty() bool
	Peek() T
}

type Queue[T any] interface {
	Enqueue(value T)
	Dequeue() T
	IsEmpty() bool
	Peek() T
}
