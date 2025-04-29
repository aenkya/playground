package algo

import "testing"

func TestJump(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{2, 3, 1, 1, 4},
			expected: 2,
		},
		{
			name:     "Example 2",
			nums:     []int{2, 3, 0, 1, 4},
			expected: 2,
		},
		{
			name:     "Single element",
			nums:     []int{0},
			expected: 0,
		},
		{
			name:     "Two elements",
			nums:     []int{1, 2},
			expected: 1,
		},
		{
			name:     "Large jump at start",
			nums:     []int{10, 1, 1, 1, 1},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := jump(tt.nums)
			if result != tt.expected {
				t.Errorf("jump(%v) = %d; want %d", tt.nums, result, tt.expected)
			}
		})
	}
}
