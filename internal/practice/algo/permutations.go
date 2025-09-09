package algo

func permute(nums []int) [][]int {
	perms := make([][]int, 0)
	var backtrack func()
	perm := make([]int, 0, len(nums))
	used := make([]bool, len(nums))

	backtrack = func() {
		if len(perm) == len(nums) {
			perms = append(perms, append([]int{}, perm...))
			return
		}

		for i := range nums {
			if used[i] {
				continue
			}

			used[i] = true
			perm = append(perm, nums[i])

			backtrack()

			perm = perm[:len(perm)-1]
			used[i] = false
		}
	}

	backtrack()
	return perms
}
