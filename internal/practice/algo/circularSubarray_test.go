package algo

import "testing"

func TestMaxCircularSubarraySum(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, -2, 3, -2}, 3},
		{[]int{5, -3, 5}, 10},
		{[]int{-3, -2, -3}, -2},
		{[]int{8, -1, 3, 4}, 15},
		{[]int{2, -2, 2, 7, 8, 0}, 19},
		{[]int{0, -1, 0}, 0},
		{[]int{3}, 3},
		{[]int{-1}, -1},
		{[]int{1, 2, 3, 4, 5}, 15},
		{[]int{-2, 4, -5, 4, -5, 9, 4}, 15},
	}

	for _, tt := range tests {
		result := maxCircularSubarraySum(tt.nums)
		if result != tt.expected {
			t.Errorf("maxCircularSubarraySum(%v) = %d; want %d", tt.nums, result, tt.expected)
		}
	}
}
