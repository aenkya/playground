package algo

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	memo := make(map[rune]int)

	for _, v := range s {
		memo[v]++
	}

	for _, v := range t {
		memo[v]--
	}

	for _, v := range memo {
		if v != 0 {
			return false
		}
	}

	return true
}
