package collections

import (
	"iter"
	"slices"
)

type Queue[T any] struct {
	q []T
}

func (q *Queue[T]) Push(items ...T) {
	q.q = append(q.q, items...)
}

func (q *Queue[T]) Pop() T {
	if len(q.q) == 0 {
		var zero T
		return zero
	}

	item := q.q[0]
	var zero T
	q.q[0] = zero
	q.q = q.q[1:]
	if len(q.q) == 0 {
		q.q = nil
	}
	return item
}

func (q *Queue[T]) Peek() T {
	if len(q.q) == 0 {
		var zero T
		return zero
	}

	return q.q[0]
}

func (q *Queue[T]) Size() int {
	return len(q.q)
}

// Values returns an iterator over the items from the front to the back of the
// queue. The queue must not be modified while iterating.
func (q *Queue[T]) Values() iter.Seq[T] {
	return slices.Values(q.q)
}

// Ordered returns an iterator over the items in pop order. It is equivalent to
// Values. The queue must not be modified while iterating.
func (q *Queue[T]) Ordered() iter.Seq[T] {
	return q.Values()
}
