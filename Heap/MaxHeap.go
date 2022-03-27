package Heap

func NewMax[T Ordered]() *Heap[T] {
	comparer := func(a, b T) bool { return a >= b }
	heap := New(comparer)

	return heap
}
