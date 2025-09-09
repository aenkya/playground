package algo

func permute(nums []int) [][]int {
	perms := make([][]int, 0)
	var backtrack func([]int)
	backtrack = func(perm []int) {
		if len(perm) == len(nums) {
			perms = append(perms, perm)
			return
		}

		for _, num := range nums {
			if !contains(perm, num) {
				perm2 := append(perm, num)
				backtrack(perm2)
			}
		}
	}

	backtrack([]int{})
	return perms
}
