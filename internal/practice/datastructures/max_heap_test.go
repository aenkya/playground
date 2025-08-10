package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxHeap(t *testing.T) {
	t.Run("Pop", func(t *testing.T) {
		testCases := []struct {
			name          string
			initialValues []int
			expectedOrder []int
			expectErr     bool
			errMsg        string
		}{
			{
				name:          "multiple elements",
				initialValues: []int{10, 4, 15, 20, 2, 8},
				expectedOrder: []int{20, 15, 10, 8, 4, 2},
			},
			{
				name:          "single element",
				initialValues: []int{42},
				expectedOrder: []int{42},
			},
			{
				name:      "from empty heap",
				expectErr: true,
				errMsg:    "heap is empty",
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				h := NewMaxHeap()
				for _, v := range tc.initialValues {
					h.Push(v)
				}

				if tc.expectErr {
					_, err := h.Pop()
					assert.Error(t, err)
					assert.Equal(t, tc.errMsg, err.Error())

					return
				}

				for _, expected := range tc.expectedOrder {
					val, err := h.Pop()
					assert.NoError(t, err)
					assert.Equal(t, expected, val)
				}

				// After all expected elements are popped, the heap should be empty.
				_, err := h.Pop()
				assert.Error(t, err)
				assert.Equal(t, "heap is empty", err.Error())
			})
		}
	})

	t.Run("Max", func(t *testing.T) {
		t.Run("on empty heap", func(t *testing.T) {
			h := NewMaxHeap()
			_, err := h.Max()
			assert.Error(t, err)
			assert.Equal(t, "heap is empty", err.Error())
		})

		t.Run("on populated heap", func(t *testing.T) {
			h := NewMaxHeap()
			h.Push(10)
			h.Push(4)
			h.Push(15)

			maxn, err := h.Max()
			assert.NoError(t, err)
			assert.Equal(t, 15, maxn)

			// Ensure Max() doesn't remove the element
			val, err := h.Pop()
			assert.NoError(t, err)
			assert.Equal(t, 15, val)

			// Check new max
			maxn, err = h.Max()
			assert.NoError(t, err)
			assert.Equal(t, 10, maxn)
		})
	})

	t.Run("IncreaseKey", func(t *testing.T) {
		testCases := []struct {
			name          string
			initialValues []int
			oldVal        int
			newVal        int
			expectErr     bool
			errMsg        string
			expectedOrder []int
		}{
			{
				name:          "successful increase makes it the new max",
				initialValues: []int{10, 4, 15, 20, 2, 8},
				oldVal:        8,
				newVal:        22,
				expectErr:     false,
				expectedOrder: []int{22, 20, 15, 10, 4, 2},
			},
			{
				name:          "value not found",
				initialValues: []int{10},
				oldVal:        5,
				newVal:        12,
				expectErr:     true,
				errMsg:        "value not found in heap",
			},
			{
				name:          "new value is smaller",
				initialValues: []int{10, 5},
				oldVal:        10,
				newVal:        2,
				expectErr:     true,
				errMsg:        "new value is less than current value",
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				h := NewMaxHeap()
				for _, v := range tc.initialValues {
					h.Push(v)
				}

				err := h.IncreaseKey(tc.oldVal, tc.newVal)

				if tc.expectErr {
					assert.Error(t, err)
					assert.Equal(t, tc.errMsg, err.Error())
				} else {
					assert.NoError(t, err)
					// Verify the heap property by popping all elements
					for _, expected := range tc.expectedOrder {
						val, popErr := h.Pop()
						assert.NoError(t, popErr)
						assert.Equal(t, expected, val)
					}
				}
			})
		}
	})
}
