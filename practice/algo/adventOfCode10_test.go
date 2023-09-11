package algo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdventOfCode10(t *testing.T) {
	tests := []struct {
		name      string
		inputFile string
		expected  int
	}{
		{"TestAdventOfCode10", "input.txt", 16800},
	}

	// output
	// ░░████████░░██░░░░░░░░██████░░░░██░░░░░░░░████████░░░░████░░░░████████░░██░░░░░░
	// ░░██░░░░░░░░██░░░░░░░░██░░░░██░░██░░░░░░░░░░░░░░██░░██░░░░██░░░░░░░░██░░██░░░░░░
	// ░░██████░░░░██░░░░░░░░██░░░░██░░██░░░░░░░░░░░░██░░░░██░░░░░░░░░░░░██░░░░██░░░░░░
	// ░░██░░░░░░░░██░░░░░░░░██████░░░░██░░░░░░░░░░██░░░░░░██░░████░░░░██░░░░░░██░░░░░░
	// ░░██░░░░░░░░██░░░░░░░░██░░░░░░░░██░░░░░░░░██░░░░░░░░██░░░░██░░██░░░░░░░░██░░░░░░
	// ░░████████░░████████░░██░░░░░░░░████████░░████████░░░░████░░░░██████░░░░██████░░

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, adventOfCode10(tt.inputFile))
		})
	}
}
