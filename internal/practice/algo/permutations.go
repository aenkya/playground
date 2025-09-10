package algo

import "math/big"

// classic backtracking
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
				//nolint:gocritic // this is intentional to not use the same permutation
				perm2 := append(perm, num)
				backtrack(perm2)
			}
		}
	}

	backtrack([]int{})

	return perms
}

// optimised for better memory management
func permutev2(nums []int) [][]int {
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

// iterative approach
func permutev3(nums []int) [][]int {
	n64 := factorial(len(nums))
	if n64 > int64(^uint(0)>>1) {
		panic("factorial(len(nums)) exceeds int capacity")
	}

	n := int(n64)
	perms := make([][]int, 1, n)
	perms[0] = make([]int, 0, len(nums))

	for _, num := range nums {
		newPerms := make([][]int, 0, n)

		for _, perm := range perms {
			for i := 0; i <= len(perm); i++ {
				newPerm := make([]int, 0, len(perm)+1)
				newPerm = append(newPerm, perm[:i]...)
				newPerm = append(newPerm, num)
				newPerm = append(newPerm, perm[i:]...)

				newPerms = append(newPerms, newPerm)
			}
		}

		perms = append([][]int{}, newPerms...)
	}

	return perms
}

func factorial(x int) int64 {
	n := new(big.Int)
	n.MulRange(1, int64(x))

	return n.Int64()
}
