package algo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotateArray(t *testing.T) {
	tests := []struct {
		given     []int
		rotations int
		want      []int
	}{
		{
			[]int{1, 2, 3, 4, 5, 6, 7},
			3,
			[]int{5, 6, 7, 1, 2, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			rotate(tt.given, tt.rotations)
			assert.Equal(t, tt.want, tt.given)
		})
	}
}
