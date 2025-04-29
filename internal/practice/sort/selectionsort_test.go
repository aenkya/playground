package sort

import (
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	tests := []struct {
		given    []int
		expected []int
	}{
		{
			[]int{5, 2, 0, 1, 3, 1},
			[]int{0, 1, 1, 2, 3, 5},
		},
		{
			[]int{0, -1, -2},
			[]int{-2, -1, 0},
		},
		{
			[]int{},
			[]int{},
		},
		{
			[]int{1},
			[]int{1},
		},
	}

	for _, tt := range tests {
		t.Run("should sort", func(t *testing.T) {
			selectionsort(tt.given)

			if !reflect.DeepEqual(tt.given, tt.expected) {
				t.Errorf("received %v, expected %v", tt.given, tt.expected)
			}
		})
	}
}
