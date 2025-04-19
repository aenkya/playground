package algo

func productExceptSelf(nums []int) []int {
	products := make([]int, len(nums))
	products[0] = 1
	for i := 1; i < len(nums); i++ {
		products[i] = products[i-1] * nums[i-1]
	}

	suffix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		products[i] = products[i] * suffix
		suffix *= nums[i]
	}

	return products
}
