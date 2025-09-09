package algo

import (
	"reflect"
	"sort"
	"testing"
)

// Helper function to sort slices of slices for comparison
func sort2DSlice(s [][]int) {
	sort.Slice(s, func(i, j int) bool {
		for x := range s[i] {
			if x >= len(s[j]) {
				return false
			}
			if s[i][x] < s[j][x] {
				return true
			}
			if s[i][x] > s[j][x] {
				return false
			}
		}
		return len(s[i]) < len(s[j])
	})
}

func TestPermute(t *testing.T) {
	tests := []struct {
		input    []int
		expected [][]int
	}{
		{
			input:    []int{},
			expected: [][]int{{}},
		},
		{
			input:    []int{1},
			expected: [][]int{{1}},
		},
		{
			input:    []int{1, 2},
			expected: [][]int{{1, 2}, {2, 1}},
		},
		{
			input: []int{1, 2, 3},
			expected: [][]int{
				{1, 2, 3},
				{1, 3, 2},
				{2, 1, 3},
				{2, 3, 1},
				{3, 1, 2},
				{3, 2, 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run("permute v1 - backtracking:", func(t *testing.T) {
			got := permute(tt.input)
			sort2DSlice(got)
			sort2DSlice(tt.expected)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("permute(%v) = %v, want %v", tt.input, got, tt.expected)
			}
		})

		t.Run("permute v2 - backtracking optimised:", func(t *testing.T) {
			got := permutev2(tt.input)
			sort2DSlice(got)
			sort2DSlice(tt.expected)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("permute(%v) = %v, want %v", tt.input, got, tt.expected)
			}
		})

		t.Run("permute v2 - iterative:", func(t *testing.T) {
			got := permutev3(tt.input)
			sort2DSlice(got)
			sort2DSlice(tt.expected)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("permute(%v) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}
