package algo

import (
	"testing"
)

func TestAtoi(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"-123", -123},
		{"  -", 0},
		{" 1231231231311133", 2147483647},
		{"-999999999999", -2147483648},
		{"  -0012gfg4", -12},
		{"0012gfg4", 12},
		{"4444g4444", 4444},
		{"453  2345", 453},
		{"   4 453", 4},
		{"", 0},
		{"+", 0},
		{"-", 0},
		{"000000", 0},
		{"2147483647", 2147483647},
		{"-2147483648", -2147483648},
		{"2147483648", 2147483647},
		{"-2147483649", -2147483648},
	}

	for _, tt := range tests {
		got := atoi(tt.input)
		if got != tt.expected {
			t.Errorf("atoi(%q) = %d; want %d", tt.input, got, tt.expected)
		}
	}
}
