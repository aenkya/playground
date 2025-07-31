package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinHeap(t *testing.T) {
	t.Run("Push and Pop", func(t *testing.T) {
		h := NewMinHeap()
		h.Push(10)
		h.Push(4)
		h.Push(15)
		h.Push(20)
		h.Push(2)
		h.Push(8)

		expectedOrder := []int{2, 4, 8, 10, 15, 20}
		for _, expected := range expectedOrder {
			val, err := h.Pop()
			assert.NoError(t, err)
			assert.Equal(t, expected, val)
		}

		// Check if heap is empty
		_, err := h.Pop()
		assert.Error(t, err)
		assert.Equal(t, "heap is empty", err.Error())
	})

	t.Run("Pop from empty heap", func(t *testing.T) {
		h := NewMinHeap()
		_, err := h.Pop()
		assert.Error(t, err, "should return an error when popping from an empty heap")
	})

	t.Run("DecreaseKey", func(t *testing.T) {
		h := NewMinHeap()
		h.Push(10)
		h.Push(4)
		h.Push(15)
		h.Push(20)
		h.Push(2)
		h.Push(8)

		err := h.DecreaseKey(15, 1)
		assert.NoError(t, err)

		// After decreasing 15 to 1, 1 should be the new minimum.
		val, err := h.Pop()
		assert.NoError(t, err)
		assert.Equal(t, 1, val)

		// The rest should follow in order
		expectedOrder := []int{2, 4, 8, 10, 20}
		for _, expected := range expectedOrder {
			val, err := h.Pop()
			assert.NoError(t, err)
			assert.Equal(t, expected, val)
		}
	})

	t.Run("DecreaseKey value not found", func(t *testing.T) {
		h := NewMinHeap()
		h.Push(10)
		err := h.DecreaseKey(5, 2)
		assert.Error(t, err)
		assert.Equal(t, "value not found", err.Error())
	})

	t.Run("DecreaseKey new value is greater", func(t *testing.T) {
		h := NewMinHeap()
		h.Push(10)
		h.Push(5)
		err := h.DecreaseKey(5, 12)
		assert.Error(t, err)
		assert.Equal(t, "new value is greater than current value", err.Error())
	})

	t.Run("Heap with one element", func(t *testing.T) {
		h := NewMinHeap()
		h.Push(42)
		val, err := h.Pop()
		assert.NoError(t, err)
		assert.Equal(t, 42, val)

		_, err = h.Pop()
		assert.Error(t, err)
	})
}
