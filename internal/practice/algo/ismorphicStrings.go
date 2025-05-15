package algo

func isIsomorphic(s, t string) bool {
	charIndexS := make(map[byte]int)
	charIndexT := make(map[byte]int)

	for i := range s {
		if charIndexS[s[i]] != charIndexT[t[i]] {
			return false
		}

		charIndexS[s[i]] = i
		charIndexT[t[i]] = i
	}

	return true
}
