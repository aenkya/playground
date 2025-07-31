package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxHeap(t *testing.T) {
	t.Run("Push and Pop", func(t *testing.T) {
		h := NewMaxHeap()
		h.Push(10)
		h.Push(4)
		h.Push(15)
		h.Push(20)
		h.Push(2)
		h.Push(8)

		expectedOrder := []int{20, 15, 10, 8, 4, 2}
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
		h := NewMaxHeap()
		_, err := h.Pop()
		assert.Error(t, err, "should return an error when popping from an empty heap")
	})

	t.Run("Max", func(t *testing.T) {
		h := NewMaxHeap()

		// Test on empty heap
		_, err := h.Max()
		assert.Error(t, err)
		assert.Equal(t, "heap is empty", err.Error())

		h.Push(10)
		h.Push(4)
		h.Push(15)

		max, err := h.Max()
		assert.NoError(t, err)
		assert.Equal(t, 15, max)

		// Ensure Max() doesn't remove the element
		val, err := h.Pop()
		assert.NoError(t, err)
		assert.Equal(t, 15, val)

		// Check new max
		max, err = h.Max()
		assert.NoError(t, err)
		assert.Equal(t, 10, max)
	})

	t.Run("IncreaseKey", func(t *testing.T) {
		h := NewMaxHeap()
		h.Push(10)
		h.Push(4)
		h.Push(15)
		h.Push(20)
		h.Push(2)
		h.Push(8)

		err := h.IncreaseKey(8, 22)
		assert.NoError(t, err)

		// After increasing 8 to 22, 22 should be the new maximum.
		val, err := h.Pop()
		assert.NoError(t, err)
		assert.Equal(t, 22, val)

		// The rest should follow in order
		expectedOrder := []int{20, 15, 10, 4, 2}
		for _, expected := range expectedOrder {
			val, err := h.Pop()
			assert.NoError(t, err)
			assert.Equal(t, expected, val)
		}
	})

	t.Run("IncreaseKey value not found", func(t *testing.T) {
		h := NewMaxHeap()
		h.Push(10)
		err := h.IncreaseKey(5, 12)
		assert.Error(t, err)
		assert.Equal(t, "value not found in heap", err.Error())
	})

	t.Run("IncreaseKey new value is smaller", func(t *testing.T) {
		h := NewMaxHeap()
		h.Push(10)
		h.Push(5)
		err := h.IncreaseKey(10, 2)
		assert.Error(t, err)
		assert.Equal(t, "new value is less than current value", err.Error())
	})

	t.Run("Heap with one element", func(t *testing.T) {
		h := NewMaxHeap()
		h.Push(42)
		val, err := h.Pop()
		assert.NoError(t, err)
		assert.Equal(t, 42, val)
		_, err = h.Pop()
		assert.Error(t, err)
	})
}
