package algo

import (
	"fmt"

	"enkya.org/playground/practice/io"
	"enkya.org/playground/utils"
)

type LongestSubstring struct {
	description string
	examples    []io.IO
	testData    []io.IO
	versions    []func(s string) int
}

func (ls *LongestSubstring) Describe() {
	fmt.Printf("\nDescription: %s\n", ls.description)
	fmt.Println("Examples:")

	for _, e := range ls.examples {
		fmt.Printf("\tInput: %v\n\tOutput: %v\n", e.Input, e.Output)
	}
}

func (ls *LongestSubstring) RunAlgo() {
	for _, version := range ls.versions {
		for _, test := range ls.testData {
			version(test.Input.(string))
		}
	}
}

func (ls *LongestSubstring) LoadTestData() {
	ls.testData = []io.IO{
		{Input: "abcabcbb", Output: 3},
		{Input: "bbbbb", Output: 1},
		{Input: "pwwkew", Output: 3},
		{Input: "", Output: 0},
	}
}

// brute force: O(n^2)
func (ls *LongestSubstring) LengthOfLongestSubstringV1(s string) int {
	longestSubstr := 0

	for i := range s {
		seen := make(map[rune]int)

		for j, v := range s[i:] {
			if _, found := seen[v]; found {
				longestSubstr = utils.MaxInt(longestSubstr, len(s[i:i+j]))
				break
			}

			seen[v] = 1
			longestSubstr = utils.MaxInt(longestSubstr, len(s[i:i+j+1]))
		}
	}

	return longestSubstr
}

// Sliding Window: O(n)
func (ls *LongestSubstring) LengthOfLongestSubstringV2(s string) int {
	result := 0

	for i := 0; i < len(s); i++ {
		set := make(map[byte]bool)
		length := 0

		for j := i; j < len(s); j++ {
			if set[s[j]] {
				break
			}

			set[s[j]] = true
			length++
		}

		result = utils.MaxInt(result, length)
	}

	return result
}

func NewLongestSubstring() *LongestSubstring {
	ls := &LongestSubstring{
		description: "Given a string s, find the length of the longest substring without repeating characters.",
		examples: []io.IO{
			{Input: "abcabcbb", Output: 3},
			{Input: "bbbbb", Output: 1},
			{Input: "pwwkew", Output: 3},
			{Input: "", Output: 0},
		},
	}

	ls.versions = []func(s string) int{
		ls.LengthOfLongestSubstringV1,
		ls.LengthOfLongestSubstringV2,
	}

	return ls
}
