package algo

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
		length   int
	}{
		{
			name:     "No duplicates",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
			length:   5,
		},
		{
			name:     "All duplicates",
			input:    []int{1, 1, 1, 1},
			expected: []int{1},
			length:   1,
		},
		{
			name:     "Mixed duplicates",
			input:    []int{1, 1, 2, 2, 3, 4, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
			length:   5,
		},
		{
			name:     "Single element",
			input:    []int{1},
			expected: []int{1},
			length:   1,
		},
		{
			name:     "Empty array",
			input:    []int{},
			expected: []int{},
			length:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := append([]int{}, tt.input...) // Create a copy to avoid modifying the original input
			length := removeDuplicates(nums)
			if length != tt.length {
				t.Errorf("expected length %d, got %d", tt.length, length)
			}
			if !reflect.DeepEqual(nums[:length], tt.expected) {
				t.Errorf("expected array %v, got %v", tt.expected, nums[:length])
			}
		})
	}
}
