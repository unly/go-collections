package collections

import (
	"cmp"
	"iter"
	"slices"
)

// NewOrderedHeap creates a new Heap using NewHeap with the default less function for
// the given cmp.Ordered type. There can optionally be passed in items to initialize
// the heap with.
func NewOrderedHeap[T cmp.Ordered](items ...T) *Heap[T] {
	return NewHeap(cmp.Less[T], items...)
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
// for a max heap. There can optionally be passed in items to initialize
// the heap with.
func NewHeap[T any](less func(a, b T) bool, items ...T) *Heap[T] {
	h := &Heap[T]{
		less: less,
		h:    items,
	}
	h.init()
	return h
}

type Heap[T any] struct {
	less func(x, y T) bool
	h    []T
}

func (h *Heap[T]) Push(items ...T) {
	for _, item := range items {
		h.h = append(h.h, item)
		h.up(len(h.h) - 1)
	}
}

func (h *Heap[T]) Pop() T {
	n := len(h.h)
	if n == 0 {
		var zero T
		return zero
	}

	h.swap(0, n-1)
	h.down(0, n-1)
	x := h.h[n-1]
	var zero T
	h.h[n-1] = zero
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

func (h *Heap[T]) init() {
	n := len(h.h)
	for i := n/2 - 1; i >= 0; i-- {
		h.down(i, n)
	}
}

func (h *Heap[T]) up(j int) {
	for {
		i := (j - 1) / 2
		if i == j || !h.less(h.h[j], h.h[i]) {
			return
		}
		h.swap(i, j)
		j = i
	}
}

func (h *Heap[T]) down(i, n int) {
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
