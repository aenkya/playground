package algo

import (
	ds "enkya.org/playground/practice/datastructures"
	"enkya.org/playground/utils"
)

func longestSubarray(nums []int, limit int) int {
	mins, max := ds.NewQueue(), ds.NewQueue()

	var ans int
	l, r := 0, 0

	for r < len(nums) {

		cur := nums[r]

		for !mins.IsEmpty() && cur < mins.Peek().(int) {
			mins.Dequeue()
		}

		for !max.IsEmpty() && cur > max.Peek().(int) {
			max.Dequeue()
		}

		mins.Enqueue(cur)
		max.Enqueue(cur)

		for !mins.IsEmpty() && !max.IsEmpty() && max.Peek().(int)-mins.Peek().(int) > limit {
			if nums[l] == mins.Peek().(int) {
				mins.Dequeue()
			}

			if nums[l] == max.Peek().(int) {
				max.Dequeue()
			}

			l++
		}

		ans = utils.MaxInt(ans, r-l+1)
		r++
	}

	return ans
}
