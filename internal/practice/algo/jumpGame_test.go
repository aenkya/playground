package algo

import "testing"

func TestCanJump(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{
			name:     "Single element",
			nums:     []int{0},
			expected: true,
		},
		{
			name:     "Can jump to the end",
			nums:     []int{2, 3, 1, 1, 4},
			expected: true,
		},
		{
			name:     "Cannot jump to the end",
			nums:     []int{3, 2, 1, 0, 4},
			expected: false,
		},
		{
			name:     "Empty array",
			nums:     []int{},
			expected: false,
		},
		{
			name:     "All zeros",
			nums:     []int{0, 0, 0, 0},
			expected: false,
		},
		{
			name:     "Large jump at start",
			nums:     []int{5, 0, 0, 0, 0},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := canJump(tt.nums)
			if result != tt.expected {
				t.Errorf("canJump(%v) = %v; want %v", tt.nums, result, tt.expected)
			}
		})
	}
}
