package algo

func removeDuplicatesII(nums []int) int {
	/*
		Given an integer array nums sorted in non-decreasing order, remove some duplicates in-place such that each unique element appears at most twice. The relative order of the elements should be kept the same.

		Since it is impossible to change the length of the array in some languages, you must instead have the result be placed in the first part of the array nums. More formally, if there are k elements after removing the duplicates, then the first k elements of nums should hold the final result. It does not matter what you leave beyond the first k elements.

		Return k after placing the final result in the first k slots of nums.

		Do not allocate extra space for another array. You must do this by modifying the input array in-place with O(1) extra memory.
	*/
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return 1
	}

	uniqueID := 0
	prevDupID := uniqueID

	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[uniqueID] {
			uniqueID++
			nums[uniqueID] = nums[j]
			prevDupID = uniqueID
		} else if prevDupID == uniqueID {
			uniqueID++
			nums[uniqueID] = nums[j]
		}
	}

	return uniqueID + 1
}
