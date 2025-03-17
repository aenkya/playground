package algo

import (
	ds "enkya.org/playground/internal/practice/datastructures"
	"enkya.org/playground/internal/utils"
)

func longestSubarray(nums []int, limit int) int {
	mins, maxs := ds.NewQueue(), ds.NewQueue()

	var l, r, ans int

	for r < len(nums) {
		cur := nums[r]

		minsint, ok := mins.Peek().(int)
		if !ok {
			continue
		}

		for !mins.IsEmpty() && cur < minsint {
			mins.Dequeue()
		}

		maxint, ok := maxs.Peek().(int)
		if !ok {
			continue
		}

		for !maxs.IsEmpty() && cur > maxint {
			maxs.Dequeue()
		}

		mins.Enqueue(cur)
		maxs.Enqueue(cur)

		for !mins.IsEmpty() && !maxs.IsEmpty() && (maxint-minsint) > limit {
			if nums[l] == minsint {
				mins.Dequeue()
			}

			if nums[l] == maxint {
				maxs.Dequeue()
			}

			l++
		}

		ans = utils.MaxInt(ans, r-l+1)
		r++
	}

	return ans
}
