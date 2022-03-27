package Heap

type (
	Ordered interface {
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
			~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
			~float32 | ~float64 |
			~string
	}

	Heap[T any] struct {
		size, capacity uint
		compare        Comparer[T]
		items          []T
	}

	Comparer[T any] func(T, T) bool
)

func New[T any](compare Comparer[T]) *Heap[T] {
	capacity := uint(10)

	return &Heap[T]{
		size:     0,
		capacity: capacity,
		compare:  compare,
		items:    make([]T, capacity),
	}
}

func (h *Heap[T]) Size() uint {
	return h.size
}

func (h *Heap[T]) Capacity() uint {
	return h.capacity
}

func (h *Heap[T]) Peak() (result T) {
	return
}

func (h *Heap[T]) Poll() (result T) {
	return
}

func (h *Heap[T]) Add(value T) {

}
