package algo

import (
	"reflect"
	"sort"
	"testing"
)

func TestGenerateParentheses(t *testing.T) {
	tests := []struct {
		n        int
		expected []string
	}{
		{
			n:        0,
			expected: []string{""},
		},
		{
			n:        1,
			expected: []string{"()"},
		},
		{
			n:        2,
			expected: []string{"(())", "()()"},
		},
		{
			n:        3,
			expected: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
	}

	for _, tt := range tests {
		got := generateParentheses(tt.n)
		// Sort both slices since order is not guaranteed
		sort.Strings(got)
		sort.Strings(tt.expected)

		if !reflect.DeepEqual(got, tt.expected) {
			t.Errorf("generateParentheses(%d) = %v, want %v", tt.n, got, tt.expected)
		}
	}
}
