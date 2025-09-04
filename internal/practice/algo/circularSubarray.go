package algo

func maxCircularSubarraySum(nums []int) int {
	n := len(nums)
	maxSum, minSum, currentMaxSum, currentMinSum, totalSum := nums[0], nums[0], nums[0], nums[0], nums[0]

	for i := 1; i < n; i++ {
		currentMaxSum = max(nums[i], nums[i]+currentMaxSum)
		maxSum = max(maxSum, currentMaxSum)

		currentMinSum = min(nums[i], nums[i]+currentMinSum)
		minSum = min(minSum, currentMinSum)

		totalSum += nums[i]
	}

	if maxSum < 0 {
		return maxSum
	}

	return max(maxSum, totalSum-minSum)
}
