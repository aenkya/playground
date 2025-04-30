package algo

import (
	"math"

	"enkya.org/playground/internal/utils"
)

func minSubArrayLen(target int, nums []int) int {
	l := 0
	r := 0
	sum := 0
	minLen := math.MaxInt

	for r < len(nums) {
		sum += nums[r]

		for sum >= target {
			minLen = utils.MinInt(minLen, r-l+1)
			sum -= nums[l]
			l++
		}

		r++
	}

	if minLen == math.MaxInt {
		return 0
	}

	return minLen
}
