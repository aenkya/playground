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
	}

	for _, tt := range tests {
		got := majorityElement(tt.given)
		if got != tt.expected {
			t.Errorf("got %v, expected %v", got, tt.expected)
		}
	}
}
