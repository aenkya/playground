package algo

import (
	"fmt"
	"math"
)

func longestSubarrayA1(nums []int, limit int) int {
	maxS := 0
	minS := 0
	s := make([]int, 0, len(nums))

	for i, v := range nums {

		if i < len(nums)-1 && v == nums[i+1] {
			s = append(s, v)
			continue
		}

		if v >= maxS {
			if checkIfADLessThanLimit(minS, v, limit) {
				s = append(s, v)
			}
			maxS = v
		} else if v <= minS {
			if checkIfADLessThanLimit(v, maxS, limit) {
				s = append(s, v)
			}
			minS = v
		} else {
			if checkIfADLessThanLimit(minS, v, limit) && checkIfADLessThanLimit(v, maxS, limit) {
				s = append(s, v)
			}
		}
	}

	fmt.Println(s)
	return len(s)
}

func checkIfADLessThanLimit(min, max, limit int) bool {
	return math.Abs(float64(max-min)) <= float64(limit)
}

/*

Function: longestsubarray
Input: nums []int, limit int
Output: len(s) int

BEGIN
    sort(nums) // o(nlogn)
    s <- []
    DECLARE maxS <- 0
    DECLARE minS <- 0

    loop through nums with i: // + O(n)
        currVal <- nums[i]
        if currVal == nums[i+1]:
            s <- append(s, currVal)
            continue
        MAX = MAX(MAXS, CURRVAL)
        MIN = MIN(MINS, CURRVAL)
        IF checkIfADLessThanLimit(MIN, MAX, limit)
            S <- APPEND(S, CURRVAL)

    RETURN len(s)

    // IF SORTED
    L := 0, R := LEN(NUMS)-1
    WHILE L<R && currDIF > :


END

FUNCTION checkIfADLessThanLimit(min, max, limit)
    DECLARE AD <- math.Abs(max-min)

    IF AD <= limit:
        RETURN true
    RETURN FALSE
END FUNCTION


*/

func longestSubarrayA2(nums []int, limit int) int {
	maxL := 0
	for i := 0; i < len(nums); i++ {
		l := 1
		for j := i + 1; j < len(nums); j++ {
			if math.Abs(float64(nums[i]-nums[j])) < float64(limit) {
				l++
			}
		}

		if l > maxL {
			maxL = l
		}
	}

	return maxL
}

func longestSubarrayA3(nums []int, limit int) int {
	maxL := 0
	for i := 0; i < len(nums); i++ {
		l := 1
		max := nums[i]
		min := nums[i]

		for j := i + 1; j < len(nums); j++ {
			if nums[j] > max {
				max = nums[j]
			}
			if nums[j] < min {
				min = nums[j]
			}

			diff := float64(max - min)

			if math.Abs(diff) > float64(limit) {
				break
			}

			l++

		}

		if l > maxL {
			maxL = l
		}
	}

	return maxL
}
