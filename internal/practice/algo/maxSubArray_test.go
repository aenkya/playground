package algo

import "testing"

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			want: 6,
		},
		{
			name: "Example 2",
			nums: []int{1},
			want: 1,
		},
		{
			name: "Example 3",
			nums: []int{5, 4, -1, 7, 8},
			want: 23,
		},
		{
			name: "All Negative Numbers",
			nums: []int{-3, -2, -1},
			want: -1,
		},
		{
			name: "Mixed Positive and Negative Numbers",
			nums: []int{2, -1, 2, 3, 4, -5},
			want: 10,
		},
		{
			name: "Single Element",
			nums: []int{10},
			want: 10,
		},
		{
			name: "Empty Array",
			nums: []int{},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArrayBruteForce(tt.nums); got != tt.want {
				t.Errorf("maxSubArrayBruteForce() = %v, want %v", got, tt.want)
			}

			if got := maxSubArrayKadane(tt.nums); got != tt.want {
				t.Errorf("maxSubArrayKadane() = %v, want %v", got, tt.want)
			}
		})
	}
}
