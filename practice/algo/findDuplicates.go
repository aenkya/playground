package algo

func findDuplicates(nums []int) map[int]int {
	return findDups(nums)
}

func findDups(nums []int) map[int]int {
	occur := make(map[int]int)

	for _, v := range nums {
		occur[v] += 1
	}

	for k, v := range occur {
		if v == 1 {
			delete(occur, k)
		}
	}

	return occur
}
