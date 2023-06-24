package algo

import (
	"fmt"
	"testing"

	"enkya.org/playground/utils"
)

func TestLongestSubstringV1(t *testing.T) {
	f := NewLongestSubstring().LengthOfLongestSubstringV1
	fmt.Println("Function name:", utils.FunctionName(f))

	tests := []struct {
		input    string
		expected int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
		{" ", 1},
		{"au", 2},
		{"dvdf", 3},
		{"abba", 2},
		{"aab", 2},
		{"tmmzuxt", 5},
	}

	for _, test := range tests {
		actual := f(test.input)

		if actual != test.expected {
			t.Errorf("Expected %d, got %d", test.expected, actual)
		}
	}
}

func TestLongestSubstringV2(t *testing.T) {
	f := NewLongestSubstring().LengthOfLongestSubstringV2
	fmt.Println("Function name:", utils.FunctionName(f))

	tests := []struct {
		input    string
		expected int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
		{" ", 1},
		{"au", 2},
		{"dvdf", 3},
		{"abba", 2},
		{"aab", 2},
		{"tmmzuxt", 5},
	}

	for _, test := range tests {
		actual := f(test.input)

		if actual != test.expected {
			t.Errorf("Expected %d, got %d", test.expected, actual)
		}
	}
}
