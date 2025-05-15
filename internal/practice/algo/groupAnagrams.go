package algo

import "sort"

func groupAnagrams(strs []string) [][]string {
	memo := make(map[string][]string)

	for _, s := range strs {
		runes := []rune(s)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})

		key := string(runes)
		memo[key] = append(memo[key], s)
	}

	result := make([][]string, 0, len(memo))
	for _, v := range memo {
		result = append(result, v)
	}

	return result
}
