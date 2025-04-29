package algo

func removeElement(nums []int, val int) int {
	/*
		constraints
		0 <= len(nums) <= 100
		0 <= val <= 100
		0 <= nums[i] <= 100
		- in place
		- return the number of elements after removing the duplicates
		interpretation
		- remove all occurrences of val from nums
		- return the number of elements after removing the duplicates
	*/
	k := 0

	for _, e := range nums {
		if e != val {
			nums[k] = e
			k++
		}
	}

	return k
}
