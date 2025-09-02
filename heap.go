package collections

import (
	"cmp"
	"iter"
	"slices"
)

// NewOrderedHeap creates a new Heap using NewHeap with the default less function for
// the given cmp.Ordered type.
func NewOrderedHeap[T cmp.Ordered]() *Heap[T] {
	return NewHeap(cmp.Less[T])
}

// NewHeap creates a new Heap using the given less function for ordering. To create a
// min or a max heap, adjust the less function accordingly.
// For example,
//
//	h := NewHeap(func(a, b int) bool { return a < b })
//
// to create a new min heap using the int type. And logically
//
//	h := NewHeap(func(a, b int) bool { return a > b })
//
// for a max heap.
func NewHeap[T any](less func(a, b T) bool) *Heap[T] {
	return &Heap[T]{
		less: less,
	}
}

type Heap[T any] struct {
	h    []T
	less func(x, y T) bool
}

func (h *Heap[T]) Push(item T) {
	h.h = append(h.h, item)
	j := len(h.h) - 1
	i := (j - 1) / 2
	for i != j && h.less(h.h[j], h.h[i]) {
		h.swap(i, j)
		j = i
		i = (j - 1) / 2
	}
}

func (h *Heap[T]) Pop() T {
	n := len(h.h)
	if n == 0 {
		var zero T
		return zero
	}

	h.swap(0, n-1)
	h.down()
	x := h.h[n-1]
	h.h = h.h[:n-1]
	return x
}

func (h *Heap[T]) Peek() T {
	if len(h.h) == 0 {
		var zero T
		return zero
	}

	return h.h[0]
}

func (h *Heap[T]) Size() int {
	return len(h.h)
}

func (h *Heap[T]) Values() iter.Seq[T] {
	return slices.Values(h.h)
}

func (h *Heap[T]) Ordered() iter.Seq[T] {
	return func(yield func(T) bool) {
		c := Heap[T]{
			h:    append([]T{}, h.h...),
			less: h.less,
		}
		for c.Size() > 0 {
			if !yield(c.Pop()) {
				return
			}
		}
	}
}

func (h *Heap[T]) swap(i, j int) {
	h.h[i], h.h[j] = h.h[j], h.h[i]
}

func (h *Heap[T]) down() {
	n := len(h.h) - 1
	i := 0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			return
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && h.less(h.h[j2], h.h[j1]) {
			j = j2 // = 2*i + 2 // right child
		}
		if !h.less(h.h[j], h.h[i]) {
			return
		}
		h.swap(i, j)
		i = j
	}
}
