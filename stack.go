package collections

import (
	"iter"
	"slices"
)

type Stack[T any] struct {
	s []T
}

func (s *Stack[T]) Push(items ...T) {
	s.s = append(s.s, items...)
}

func (s *Stack[T]) Pop() T {
	if len(s.s) == 0 {
		var zero T
		return zero
	}

	item := s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]
	return item
}

func (s *Stack[T]) Peek() T {
	if len(s.s) == 0 {
		var zero T
		return zero
	}

	return s.s[len(s.s)-1]
}

func (s *Stack[T]) Size() int {
	return len(s.s)
}

func (s *Stack[T]) Values() iter.Seq[T] {
	return slices.Values(s.s)
}

func (s *Stack[T]) Ordered() iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := len(s.s) - 1; i >= 0; i-- {
			if !yield(s.s[i]) {
				return
			}
		}
	}
}
