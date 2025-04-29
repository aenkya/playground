package sort

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{5, 2, 0, 1, 3, 1},
			expected: []int{0, 1, 1, 2, 3, 5},
		},
	}

	for _, tt := range tests {
		t.Run("should sort", func(t *testing.T) {
			bubbleSort(tt.input)

			if !reflect.DeepEqual(tt.input, tt.expected) {
				t.Errorf("got %v, wanted %v", tt.input, tt.expected)
			}
		})
	}
}
