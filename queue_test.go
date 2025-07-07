package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue_Push(t *testing.T) {
	t.Run("push item", func(t *testing.T) {
		var q Queue[int]

		q.Push(42)

		assert.Equal(t, 42, q.Peek())
	})
	t.Run("push second item", func(t *testing.T) {
		var q Queue[int]

		q.Push(42, 43)

		assert.Equal(t, 42, q.Peek())
	})
}

func TestQueue_Pop(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		var q Queue[int]

		assert.Equal(t, 0, q.Pop())
	})

	t.Run("sample items", func(t *testing.T) {
		var q Queue[int]
		q.Push(42, 43, 44)

		assert.Equal(t, 42, q.Pop())
		assert.Equal(t, 43, q.Pop())
		assert.Equal(t, 44, q.Pop())
	})
}

func TestQueue_Peek(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		var q Queue[int]

		assert.Equal(t, 0, q.Peek())
	})
	t.Run("sample items", func(t *testing.T) {
		var q Queue[int]
		q.Push(42, 43, 44)

		assert.Equal(t, 42, q.Peek())
	})
}

func TestQueue_Size(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		var q Queue[int]

		assert.Equal(t, 0, q.Size())
	})
	t.Run("sample items", func(t *testing.T) {
		var q Queue[int]
		q.Push(42, 43, 44)

		assert.Equal(t, 3, q.Size())
	})
}

func TestQueue_Values(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		var q Queue[int]

		for range q.Values() {
			t.Fail()
		}
	})

	t.Run("sample items", func(t *testing.T) {
		var q Queue[int]
		q.Push(42, 43, 44)

		var got []int
		for i := range q.Values() {
			got = append(got, i)
		}

		assert.Equal(t, []int{42, 43, 44}, got)
	})
}

func TestQueue_Ordered(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		var q Queue[int]

		for range q.Ordered() {
			t.Fail()
		}
	})

	t.Run("sample items", func(t *testing.T) {
		var q Queue[int]
		q.Push(42, 43, 44)

		var got []int
		for i := range q.Ordered() {
			got = append(got, i)
		}

		assert.Equal(t, []int{42, 43, 44}, got)
	})
}
