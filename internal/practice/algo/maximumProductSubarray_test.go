package algo

import (
	"testing"
)

func TestMaxProduct(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{2, 3, -2, 4}, 6},
		{[]int{-2, 0, -1}, 0},
		{[]int{-2, 3, -4}, 24},
		{[]int{0, 2}, 2},
		{[]int{-2}, -2},
		{[]int{1, -2, 3, -4, 5, -6}, 360},
		{[]int{2, -5, -2, -4, 3}, 24},
		{[]int{0, 0, 0}, 0},
		{[]int{1, 2, 3, 4}, 24},
		{[]int{-1, -3, -10, 0, 60}, 60},
		{[]int{-2, -3, 0, -2, -40}, 80},
	}

	for _, tt := range tests {
		result := maxProduct(tt.nums)
		if result != tt.expected {
			t.Errorf("maxProduct(%v) = %d; want %d", tt.nums, result, tt.expected)
		}
	}
}
