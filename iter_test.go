package collections

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCollect(t *testing.T) {
	t.Run("sample numbers", func(t *testing.T) {
		sample := []int{1, 5, 5, 42}
		vals := Collect(slices.Values(sample), len(sample))

		assert.Equal(t, vals, sample)
	})
	t.Run("empty", func(t *testing.T) {
		vals := Collect(slices.Values([]int{}), 0)

		assert.Len(t, vals, 0)
	})
	t.Run("nil", func(t *testing.T) {
		vals := Collect[int](nil, 0)

		assert.Len(t, vals, 0)
	})
	t.Run("sample numbers, incorrect hint", func(t *testing.T) {
		sample := []int{1, 5, 5, 42}
		vals := Collect(slices.Values(sample), 0)

		assert.Equal(t, vals, sample)
	})
	t.Run("empty, incorrect hint", func(t *testing.T) {
		vals := Collect(slices.Values([]int{}), 42)

		assert.Len(t, vals, 0)
	})
	t.Run("nil, incorrect hint", func(t *testing.T) {
		vals := Collect[int](nil, 42)

		assert.Len(t, vals, 0)
	})
}
