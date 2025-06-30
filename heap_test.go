package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeap_Push(t *testing.T) {
	t.Run("", func(t *testing.T) {
		h := NewOrderedHeap[int]()

		h.Push(1)
		h.Push(2)
		h.Push(3)

		assert.Equal(t, 3, h.Size())
	})
}

func TestHeap_Pop(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		h := NewOrderedHeap[int]()

		assert.Equal(t, 0, h.Pop())
	})
	t.Run("sample pop", func(t *testing.T) {
		h := NewOrderedHeap[int]()

		h.Push(1)
		h.Push(2)
		h.Push(3)
		h.Push(4)
		h.Push(5)
		h.Push(6)

		assert.Equal(t, 1, h.Pop())
		assert.Equal(t, 2, h.Pop())
		assert.Equal(t, 3, h.Pop())
		assert.Equal(t, 4, h.Pop())
		assert.Equal(t, 5, h.Pop())
		assert.Equal(t, 6, h.Pop())
	})
	t.Run("same values", func(t *testing.T) {
		h := NewOrderedHeap[int]()

		h.Push(1)
		h.Push(1)
		h.Push(1)
		h.Push(1)
		h.Push(1)
		h.Push(1)

		assert.Equal(t, 1, h.Pop())
		assert.Equal(t, 1, h.Pop())
		assert.Equal(t, 1, h.Pop())
		assert.Equal(t, 1, h.Pop())
		assert.Equal(t, 1, h.Pop())
		assert.Equal(t, 1, h.Pop())
	})
	t.Run("max heap", func(t *testing.T) {
		h := NewHeap(func(a, b int) bool {
			return a > b
		})

		h.Push(1)
		h.Push(2)
		h.Push(3)
		h.Push(4)
		h.Push(5)
		h.Push(6)

		assert.Equal(t, 6, h.Pop())
		assert.Equal(t, 5, h.Pop())
		assert.Equal(t, 4, h.Pop())
		assert.Equal(t, 3, h.Pop())
		assert.Equal(t, 2, h.Pop())
		assert.Equal(t, 1, h.Pop())
	})
}

func TestHeap_Peek(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		h := NewOrderedHeap[int]()

		assert.Equal(t, 0, h.Peek())
	})
	t.Run("sample peek", func(t *testing.T) {
		h := NewOrderedHeap[int]()
		h.Push(1)
		h.Push(2)

		assert.Equal(t, 1, h.Peek())
	})
}

func TestHeap_Size(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		h := NewOrderedHeap[int]()

		assert.Equal(t, 0, h.Size())
	})
	t.Run("sample size", func(t *testing.T) {
		h := NewOrderedHeap[int]()
		h.Push(1)
		h.Push(2)
		h.Push(3)
		h.Push(4)

		assert.Equal(t, 4, h.Size())
	})
}

func TestHeap_Values(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		h := NewOrderedHeap[int]()

		for range h.Values() {
			t.Fail()
		}
	})
	t.Run("sample values", func(t *testing.T) {
		h := NewOrderedHeap[int]()
		h.Push(4)
		h.Push(3)
		h.Push(2)
		h.Push(1)

		var got []int
		for i := range h.Values() {
			got = append(got, i)
		}

		assert.ElementsMatch(t, []int{3, 4, 2, 1}, got)
	})
}

func TestHeap_Ordered(t *testing.T) {
	t.Run("sample values", func(t *testing.T) {
		h := NewOrderedHeap[int]()
		h.Push(4)
		h.Push(3)
		h.Push(2)
		h.Push(1)
		h.Push(0)
		h.Push(-1)

		var got []int
		for i := range h.Ordered() {
			got = append(got, i)
		}

		assert.Equal(t, []int{-1, 0, 1, 2, 3, 4}, got)
	})
}
