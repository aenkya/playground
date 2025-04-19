package algo

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Example 1",
			input:    []int{1, 2, 3, 4},
			expected: []int{24, 12, 8, 6},
		},
		{
			name:     "Example 2 with zero",
			input:    []int{0, 1, 2, 3},
			expected: []int{6, 0, 0, 0},
		},
		{
			name:     "Single element",
			input:    []int{5},
			expected: []int{1},
		},
		{
			name:     "Two elements",
			input:    []int{3, 4},
			expected: []int{4, 3},
		},
		{
			name:     "All zeros",
			input:    []int{0, 0, 0},
			expected: []int{0, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := productExceptSelf(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("productExceptSelf(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}
