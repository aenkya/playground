package algo

func majorityElement(nums []int) int {
	count := 0
	candidate := 0

	for _, elem := range nums {
		if count == 0 {
			candidate = elem
		}

		if elem == candidate {
			count++
		} else {
			count--
		}
	}

	return candidate
}
