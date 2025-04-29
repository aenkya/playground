package algo

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return 1
	}

	uniqueID := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[uniqueID] {
			uniqueID++
			nums[uniqueID] = nums[i]
		}
	}

	return uniqueID + 1
}
