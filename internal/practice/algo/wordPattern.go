package algo

import "strings"

func wordPattern(pattern string, s string) bool {
	split := strings.Fields(s)

	if len(pattern) != len(split) {
		return false
	}

	letterIndex := make(map[byte]int)
	wordIndex := make(map[string]int)

	for i := 0; i < len(pattern); i++ {
		if _, ok := letterIndex[pattern[i]]; !ok {
			letterIndex[pattern[i]] = i
		}

		if _, ok := wordIndex[split[i]]; !ok {
			wordIndex[split[i]] = i
		}

		if letterIndex[pattern[i]] != wordIndex[split[i]] {
			return false
		}
	}

	return true
}
