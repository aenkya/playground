package algo

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return 1
	}

	unique_idx := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[unique_idx] {
			unique_idx++
			nums[unique_idx] = nums[i]
		}
	}

	return unique_idx + 1
}
