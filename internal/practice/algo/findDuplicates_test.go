package algo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindDuplicates(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		expect map[int]int
	}{
		{
			"should find all duplicate values and their count",
			[]int{1, 2, 2, 3, 4, 6, 6, 7, 8, 9, 5, 2, 6, 1, 8},
			map[int]int{1: 2, 2: 3, 6: 3, 8: 2},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := findDuplicates(test.input)
			assert.Equal(t, test.expect, actual)
		})
	}
}
