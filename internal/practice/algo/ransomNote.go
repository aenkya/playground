package algo

func canConstruct(ransomNote, magazine string) bool {
	letterIndex := make(map[byte]int)

	for i := range magazine {
		letterIndex[magazine[i]]++
	}

	for i := range ransomNote {
		if count, ok := letterIndex[ransomNote[i]]; !ok || count <= 0 {
			return false
		}

		letterIndex[ransomNote[i]]--
	}

	return true
}
