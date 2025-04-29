package algo

import "testing"

func TestCanCompleteCircuit(t *testing.T) {
	tests := []struct {
		name     string
		gas      []int
		cost     []int
		expected int
	}{
		{
			name:     "Valid circuit exists",
			gas:      []int{1, 2, 3, 4, 5},
			cost:     []int{3, 4, 5, 1, 2},
			expected: 3,
		},
		{
			name:     "No valid circuit",
			gas:      []int{2, 3, 4},
			cost:     []int{3, 4, 3},
			expected: -1,
		},
		{
			name:     "Single station, valid circuit",
			gas:      []int{5},
			cost:     []int{4},
			expected: 0,
		},
		{
			name:     "Single station, no valid circuit",
			gas:      []int{3},
			cost:     []int{5},
			expected: -1,
		},
		{
			name:     "All stations have equal gas and cost",
			gas:      []int{1, 1, 1, 1},
			cost:     []int{1, 1, 1, 1},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := canCompleteCircuit(tt.gas, tt.cost)
			if result != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, result)
			}
		})
	}
}
