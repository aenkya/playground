package algo

import "testing"

func TestMedian2SortedArraysV2(t *testing.T) {
	f := NewMedian2SortedArrays().Median2SortedArraysV2
	tests := []struct {
		name     string
		input    [][]int
		expected float64
	}{
		{
			name: "test1",
			input: [][]int{
				{1, 3},
				{2},
			},
			expected: 2.0,
		},
		{
			name: "test2",
			input: [][]int{
				{1, 2},
				{3, 4},
			},
			expected: 2.5,
		},
		{
			name: "test3",
			input: [][]int{
				{0, 0},
				{0, 0},
			},
			expected: 0.0,
		},
		{
			name: "test4",
			input: [][]int{
				{},
				{1},
			},
			expected: 1.0,
		},
		{
			name: "test5",
			input: [][]int{
				{2},
				{},
			},
			expected: 2.0,
		},
		{
			name: "test6",
			input: [][]int{
				{},
				{2, 3},
			},
			expected: 2.5,
		},
	}

	for _, test := range tests {
		actual := f(test.input[0], test.input[1])

		if actual != test.expected {
			t.Errorf("Expected %f, got %f", test.expected, actual)
		}
	}

}
