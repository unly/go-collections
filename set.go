package collections

import (
	"iter"
	"maps"
)

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](items ...T) Set[T] {
	s := Set[T](make(map[T]struct{}, len(items)))
	s.Add(items...)

	return s
}

func (s Set[T]) Add(items ...T) {
	for _, i := range items {
		s[i] = struct{}{}
	}
}

func (s Set[T]) Delete(item T) {
	delete(s, item)
}

func (s Set[T]) Contains(item T) bool {
	_, ok := s[item]
	return ok
}

func (s Set[T]) ContainsAll(items ...T) bool {
	for _, i := range items {
		if !s.Contains(i) {
			return false
		}
	}
	return true
}

func (s Set[T]) Values() iter.Seq[T] {
	return maps.Keys(s)
}

func (s Set[T]) Size() int {
	return len(s)
}

func (s Set[T]) Eq(other Set[T]) bool {
	return maps.Equal(s, other)
}
