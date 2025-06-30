package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack_Push(t *testing.T) {
	t.Run("single item", func(t *testing.T) {
		var s Stack[int]

		s.Push(42)
		assert.Equal(t, 42, s.Pop())
	})
	t.Run("multiple items", func(t *testing.T) {
		var s Stack[int]

		s.Push(42)
		s.Push(45)
		assert.Equal(t, 45, s.Pop())
		assert.Equal(t, 42, s.Pop())
	})
}

func TestStack_Pop(t *testing.T) {
	t.Run("empty stack", func(t *testing.T) {
		var s Stack[int]

		assert.Equal(t, 0, s.Pop())
	})
	t.Run("single item", func(t *testing.T) {
		var s Stack[int]

		s.Push(42)
		assert.Equal(t, 42, s.Pop())
		assert.Equal(t, 0, s.Pop())
	})
}

func TestStack_Peek(t *testing.T) {
	t.Run("empty stack", func(t *testing.T) {
		var s Stack[int]

		assert.Equal(t, 0, s.Peek())
	})
	t.Run("single item", func(t *testing.T) {
		var s Stack[int]

		s.Push(42)
		assert.Equal(t, 42, s.Peek())
	})
}

func TestStack_Size(t *testing.T) {
	t.Run("empty stack", func(t *testing.T) {
		var s Stack[int]

		assert.Equal(t, 0, s.Size())
	})
	t.Run("sample items", func(t *testing.T) {
		var s Stack[int]
		s.Push(42)
		s.Push(42)
		s.Push(42)

		assert.Equal(t, 3, s.Size())
	})
}

func TestStack_Values(t *testing.T) {
	t.Run("empty stack", func(t *testing.T) {
		var s Stack[int]

		for range s.Values() {
			t.Fail()
		}
	})
	t.Run("sample items", func(t *testing.T) {
		var s Stack[int]
		s.Push(42)
		s.Push(43)
		s.Push(44)

		var got []int
		for i := range s.Values() {
			got = append(got, i)
		}

		assert.ElementsMatch(t, []int{44, 43, 42}, got)
	})
}

func TestStack_Ordered(t *testing.T) {
	t.Run("empty stack", func(t *testing.T) {
		var s Stack[int]

		for range s.Ordered() {
			t.Fail()
		}
	})
	t.Run("sample items", func(t *testing.T) {
		var s Stack[int]
		s.Push(42)
		s.Push(43)
		s.Push(44)

		var got []int
		for i := range s.Ordered() {
			got = append(got, i)
		}

		assert.Equal(t, []int{44, 43, 42}, got)
	})
}
