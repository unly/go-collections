package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet_Add(t *testing.T) {
	t.Run("single item", func(t *testing.T) {
		s := NewSet[int]()
		s.Add(42)

		assert.True(t, s.Contains(42))
		assert.Equal(t, 1, s.Size())
	})
	t.Run("add item twice", func(t *testing.T) {
		s := NewSet[int]()
		s.Add(42)
		s.Add(42)

		assert.True(t, s.Contains(42))
		assert.Equal(t, 1, s.Size())
	})
}

func TestSet_Contains(t *testing.T) {
	t.Run("contains item", func(t *testing.T) {
		s := NewSet[int]()
		s.Add(42)

		assert.True(t, s.Contains(42))
	})
	t.Run("does not contain item", func(t *testing.T) {
		s := NewSet[int]()
		s.Add(42)

		assert.False(t, s.Contains(43))
	})
}

func TestSet_ContainsAll(t *testing.T) {
	t.Run("contains all items", func(t *testing.T) {
		s := NewSet[int]()
		s.Add(42)
		s.Add(43)

		assert.True(t, s.ContainsAll(42, 43))
	})
	t.Run("does not contain all items", func(t *testing.T) {
		s := NewSet[int]()
		s.Add(42)
		s.Add(43)

		assert.False(t, s.ContainsAll(42, 43, 44))
	})
}

func TestSet_Delete(t *testing.T) {
	t.Run("delete single item", func(t *testing.T) {
		s := NewSet[int]()
		s.Add(42)

		assert.True(t, s.Contains(42))

		s.Delete(42)

		assert.False(t, s.Contains(42))
	})
}

func TestSet_Values(t *testing.T) {
	s := NewSet[int]()
	s.Add(42)
	s.Add(43)
	s.Add(44)

	var got []int
	for i := range s.Values() {
		got = append(got, i)
	}

	assert.ElementsMatch(t, got, []int{42, 43, 44})
}
