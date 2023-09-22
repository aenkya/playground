package algo

func findDuplicates(nums []int) map[int]int {
	occurrences := make(map[int]int)

	for _, v := range nums {
		occurrences[v] += 1
	}

	for k, v := range occurrences {
		if v == 1 {
			delete(occurrences, k)
		}
	}
	return occurrences
}
