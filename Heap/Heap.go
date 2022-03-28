package Heap

type (
	Ordered interface {
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
			~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
			~float32 | ~float64 |
			~string
	}

	Heap[T any] struct {
		size, capacity int
		compare        Comparer[T]
		items          []T
	}

	Comparer[T any] func(T, T) bool
)

func New[T any](compare Comparer[T]) *Heap[T] {
	capacity := 10

	return &Heap[T]{
		size:     0,
		capacity: capacity,
		compare:  compare,
		items:    make([]T, capacity),
	}
}

func (h *Heap[T]) Size() int {
	return h.size
}

func (h *Heap[T]) Capacity() int {
	return h.capacity
}

func (h *Heap[T]) Peek() (result T) {
	return h.items[0]
}

func (h *Heap[T]) Poll() (result T) {
	item := h.items[0]
	h.items[0] = h.items[h.size-1]
	h.heapifyDown()
	return item
}

func (h *Heap[T]) Add(value T) {
	h.ensureExtraCapacity()
	h.items[h.size] = value
	h.size++
	h.heapifyUp()
}

func (h *Heap[T]) heapifyDown() {

}

func (h *Heap[T]) heapifyUp() {

}

func (h *Heap[T]) getLeftChildIndex(index int) int  { return index*2 + 1 }
func (h *Heap[T]) getRightChildIndex(index int) int { return index*2 + 2 }
func (h *Heap[T]) getParentIndex(index int) int     { return (index - 1) / 2 }

func (h *Heap[T]) hasLeftChild(index int) bool  { return h.getLeftChildIndex(index) < h.size }
func (h *Heap[T]) hasRightChild(index int) bool { return h.getRightChildIndex(index) < h.size }
func (h *Heap[T]) hasParent(index int) bool     { return h.getParentIndex(index) >= 0 }

func (h *Heap[T]) leftChild(index int) T  { return h.items[h.getLeftChildIndex(index)] }
func (h *Heap[T]) rightChild(index int) T { return h.items[h.getRightChildIndex(index)] }
func (h *Heap[T]) parent(index int) T     { return h.items[h.getParentIndex(index)] }

func (h *Heap[T]) swap(leftIndex, rightIndex int) {
	h.items[leftIndex], h.items[rightIndex] = h.items[rightIndex], h.items[leftIndex]
}

func (h *Heap[T]) ensureExtraCapacity() {
	if h.size == h.capacity {
		h.items = append(h.items, make([]T, h.capacity)...)
		h.capacity *= 2
	}
}
