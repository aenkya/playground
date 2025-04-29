package algo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		nums         []int
		expectedNums []int
		val          int
		want         int
	}{
		{
			nums:         []int{3, 2, 2, 3},
			val:          3,
			want:         2,
			expectedNums: []int{2, 2},
		},
		{
			nums:         []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:          2,
			want:         5,
			expectedNums: []int{0, 1, 3, 0, 4},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("for input: %v, val: %d", tt.nums, tt.val), func(t *testing.T) {
			got := removeElement(tt.nums, tt.val)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(tt.nums[:got], tt.expectedNums) {
				t.Errorf("got %v, want %v", tt.nums[:got], tt.expectedNums)
			}
		})
	}
}
