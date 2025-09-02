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
	q.q = q.q[1:]
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

func (q *Queue[T]) Values() iter.Seq[T] {
	return slices.Values(q.q)
}

func (q *Queue[T]) Ordered() iter.Seq[T] {
	return q.Values()
}
