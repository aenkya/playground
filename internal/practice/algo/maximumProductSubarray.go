package algo

func maxProduct(nums []int) int {
	maxProduct := 1
	n := len(nums)
	for i := 0; i < n; i++ {
		maxProduct *= nums[i]
	}

	if maxProduct > 0 {
		return maxProduct
	}

	product := 1
	for l := 1; l < n; l++ {
		product *= nums[l-1]
		maxProduct = max(maxProduct, product)

		if product == 0 {
			product = 1
		}
	}

	product = 1
	for r := n - 1 - 1; r >= 0; r-- {
		product *= nums[r+1]
		maxProduct = max(maxProduct, product)

		if product == 0 {
			product = 1
		}
	}

	return maxProduct
}
