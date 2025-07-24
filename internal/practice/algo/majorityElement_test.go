package algo

import "testing"

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		given    []int
		expected int
	}{
		{
			[]int{2, 2, 1, 1, 1, 2, 2},
			2,
		},
		{
			[]int{3, 2, 3},
			3,
		},
		{
			[]int{1},
			1,
		},
		{
			[]int{6, 5, 5},
			5,
		},
		{
			[]int{4, 4, 4, 2, 2, 2, 4},
			4,
		},
		{
			[]int{9, 9, 9, 9, 1, 2, 3, 4},
			9,
		},
		{
			[]int{7, 7, 8, 8, 7},
			7,
		},
		{
			[]int{10, 10, 10, 10, 10},
			10,
		},
	}

	for _, tt := range tests {
		got := majorityElement(tt.given)
		if got != tt.expected {
			t.Errorf("got %v, expected %v", got, tt.expected)
		}
	}
}
