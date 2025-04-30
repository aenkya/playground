package algo

import "enkya.org/playground/internal/utils"

func lengthOfLongestSubstring(s string) int {
	memo := make(map[byte]int)
	l := 0
	maxLen := 0

	for r := range s {
		if idx, ok := memo[s[r]]; ok && idx >= l {
			l = idx + 1
		}

		memo[s[r]] = r
		maxLen = utils.MaxInt(maxLen, r-l+1)
	}

	return maxLen
}
