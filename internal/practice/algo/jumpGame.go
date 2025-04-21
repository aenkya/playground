package algo

func canJump(nums []int) bool {
	memo := make(map[int]bool)
	return travel(memo, nums, 0) || canJumpGreedy(nums)
}

func travel(memo map[int]bool, nums []int, index int) bool {
	if val, ok := memo[index]; ok {
		return val
	}

	if index == len(nums)-1 {
		memo[index] = true
		return true
	}

	if index >= len(nums) || nums[index] == 0 {
		memo[index] = false
		return false
	}

	for i := nums[index]; i > 0; i-- {
		if travel(memo, nums, index+i) {
			return true
		}
	}

	return false
}

func canJumpGreedy(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	if len(nums) == 1 {
		return true
	}
	maxReach := 0
	for i := 0; i < len(nums); i++ {
		if i > maxReach {
			return false
		}
		maxReach = max(maxReach, i+nums[i])
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
