package algo

import "testing"

func TestLongestContiguousSubArrayWithADLTLimit(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		limit int
		want  int
	}{
		{
			name:  "test 1",
			input: []int{8, 2, 4, 7},
			limit: 4,
			want:  2,
		},
		{
			name:  "test 2",
			input: []int{10, 1, 2, 4, 7, 2},
			limit: 5,
			want:  4,
		},
		{
			name:  "test 3",
			input: []int{4, 2, 2, 2, 4, 4, 2, 2},
			limit: 0,
			want:  3,
		},
		{
			name:  "test 4",
			input: []int{8},
			limit: 10,
			want:  1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := longestSubarrayA1(test.input, test.limit)
			if got != test.want {
				t.Errorf("got %d, want %d", got, test.want)
			}
		})
	}
}
