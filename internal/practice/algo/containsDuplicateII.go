package algo

func containsDuplicateII(nums []int, k int) bool {
	memo := make(map[int]int)

	for i, v := range nums {
		if _, ok := memo[v]; ok {
			if i-memo[v] <= k {
				return true
			}
		}

		memo[v] = i
	}

	return false
}
