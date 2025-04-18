package algo

import "testing"

func TestHIndex(t *testing.T) {
	tests := []struct {
		name      string
		citations []int
		expected  int
	}{
		{
			name:      "No citations",
			citations: []int{},
			expected:  0,
		},
		{
			name:      "Single citation",
			citations: []int{5},
			expected:  1,
		},
		{
			name:      "All citations are the same",
			citations: []int{3, 3, 3, 3},
			expected:  3,
		},
		{
			name:      "Descending citations",
			citations: []int{6, 5, 3, 1, 0},
			expected:  3,
		},
		{
			name:      "Mixed citations",
			citations: []int{1, 4, 1, 4, 2, 1, 3},
			expected:  3,
		},
		{
			name:      "All zeros",
			citations: []int{0, 0, 0, 0},
			expected:  0,
		},
		{
			name:      "High citations",
			citations: []int{10, 8, 5, 4, 3},
			expected:  4,
		},
		{
			name:      "Single high citation",
			citations: []int{100},
			expected:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := hIndex(tt.citations)
			if result != tt.expected {
				t.Errorf("hIndex(%v) = %d; want %d", tt.citations, result, tt.expected)
			}
		})
	}
}
