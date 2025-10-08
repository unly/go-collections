package collections

import (
	"iter"
)

// Collect collects all the values from the iterator and returns
// a slice holding them. The sizeHint parameter can be specified to
// optimize the memory allocation of the slice and avoid unnecessary
// copies.
func Collect[T any](iter iter.Seq[T], sizeHint int) []T {
	if iter == nil {
		return nil
	}

	collected := make([]T, 0, sizeHint)
	for val := range iter {
		collected = append(collected, val)
	}

	return collected
}
