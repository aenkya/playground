package algo

import "testing"

func TestSeatingStudents(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "test1",
			input:    []int{6, 4},
			expected: 4,
		},
		{
			name:     "test2",
			input:    []int{8, 1, 8},
			expected: 6,
		},
		{
			name:     "test3",
			input:    []int{12, 2, 6, 7, 11},
			expected: 6,
		},
		{
			name:     "test4",
			input:    []int{12, 2, 6, 7, 11, 12},
			expected: 5,
		},
		{
			name:     "edge case 1",
			input:    []int{2, 1},
			expected: 0,
		},
		{
			name:     "edge case 2",
			input:    []int{2, 2},
			expected: 0,
		},
		{
			name:     "edge case 3",
			input:    []int{2, 1, 2},
			expected: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := SeatingStudents(test.input)
			if actual != test.expected {
				t.Errorf("For %s, Expected %d, got %d", test.name, test.expected, actual)
			}
		})
	}
}
