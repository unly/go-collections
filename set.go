package collections

import (
	"iter"
	"maps"
)

type Set[T comparable] struct {
	s map[T]struct{}
}

func (s *Set[T]) Add(items ...T) {
	if s.s == nil {
		s.s = make(map[T]struct{}, len(items))
	}

	for _, i := range items {
		s.s[i] = struct{}{}
	}
}

func (s *Set[T]) Delete(items ...T) {
	for _, i := range items {
		delete(s.s, i)
	}
}

func (s *Set[T]) Contains(item T) bool {
	if s.s == nil {
		return false
	}

	_, ok := s.s[item]
	return ok
}

func (s *Set[T]) ContainsAll(items ...T) bool {
	for _, i := range items {
		if !s.Contains(i) {
			return false
		}
	}
	return true
}

// Values returns an iterator over the items in no particular order. The set
// must not be modified while iterating.
func (s *Set[T]) Values() iter.Seq[T] {
	return maps.Keys(s.s)
}

func (s *Set[T]) Size() int {
	return len(s.s)
}

// Eq reports whether both sets contain the same items. A nil other is treated
// as an empty set.
func (s *Set[T]) Eq(other *Set[T]) bool {
	if other == nil {
		return s.Size() == 0
	}

	return maps.Equal(s.s, other.s)
}
