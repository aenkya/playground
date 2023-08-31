package algo

import "testing"

func TestLatestTime(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected string
	}{
		{"1", []int{1, 2, 3, 4}, "23:41"},
		{"2", []int{5, 5, 5, 5}, ""},
		{"3", []int{0, 0, 0, 0}, "00:00"},
		{"4", []int{0, 0, 1, 0}, "10:00"},
		{"5", []int{0, 0, 0, 1}, "10:00"},
		{"6", []int{0, 0, 1, 1}, "11:00"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := largestTimeFromDigits(test.input)

			if actual != test.expected {
				t.Errorf("Expected %s, got %s", test.expected, actual)
			}
		})
	}
}
