package algo

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicatesII(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
		k        int
	}{
		{
			name:     "No duplicates",
			input:    []int{1, 2, 3, 4},
			expected: []int{1, 2, 3, 4},
			k:        4,
		},
		{
			name:     "Duplicates appear at most twice",
			input:    []int{1, 1, 1, 2, 2, 3},
			expected: []int{1, 1, 2, 2, 3},
			k:        5,
		},
		{
			name:     "All elements are the same",
			input:    []int{1, 1, 1, 1},
			expected: []int{1, 1},
			k:        2,
		},
		{
			name:     "Empty array",
			input:    []int{},
			expected: []int{},
			k:        0,
		},
		{
			name:     "Single element array",
			input:    []int{1},
			expected: []int{1},
			k:        1,
		},
		{
			name:     "Mixed duplicates",
			input:    []int{0, 0, 1, 1, 1, 2, 3, 3},
			expected: []int{0, 0, 1, 1, 2, 3, 3},
			k:        7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := append([]int{}, tt.input...) // Create a copy of the input
			k := removeDuplicatesII(nums)

			if k != tt.k {
				t.Errorf("expected k = %d, got %d", tt.k, k)
			}

			if !reflect.DeepEqual(nums[:k], tt.expected) {
				t.Errorf("expected nums = %v, got %v", tt.expected, nums[:k])
			}
		})
	}
}
