package algo

import "testing"

func TestCandy(t *testing.T) {
	tests := []struct {
		name     string
		ratings  []int
		expected int
	}{
		{
			name:     "Single child",
			ratings:  []int{1},
			expected: 1,
		},
		{
			name:     "All equal ratings",
			ratings:  []int{1, 1, 1},
			expected: 3,
		},
		{
			name:     "Increasing ratings",
			ratings:  []int{1, 2, 3},
			expected: 6,
		},
		{
			name:     "Decreasing ratings",
			ratings:  []int{3, 2, 1},
			expected: 6,
		},
		{
			name:     "Mixed ratings",
			ratings:  []int{1, 0, 2},
			expected: 5,
		},
		{
			name:     "Complex case",
			ratings:  []int{1, 2, 2},
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := candy(tt.ratings)
			if result != tt.expected {
				t.Errorf("candy(%v) = %d; want %d", tt.ratings, result, tt.expected)
			}
		})
	}
}
