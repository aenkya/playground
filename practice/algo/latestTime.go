package algo

import (
	"fmt"
	"math"
	"strconv"
)

func largestTimeFromDigits(arr []int) string {
	lh := -1
	curr := -1
	second := -1
	lhh := 0.0
	lmh := 0.0

	for i, v := range arr {
		if v < 3 {
			if lh < v {
				lh = v
				curr = i
			}
		}
	}

	if lh == -1 {
		return ""
	}

	for i, v := range arr {
		if i == curr {
			continue
		}

		num, _ := strconv.ParseFloat(fmt.Sprintf("%d%d", lh, v), 64)
		diff := math.Abs(24.0 - num)

		if diff < math.Abs(24.0-lhh) {
			lhh = num
			second = i
		}
	}

	minDigit1, minDigit2 := -1, -1

	for i, v := range arr {
		if i == second || i == curr {
			continue
		}

		if minDigit1 < 0 {
			minDigit1 = v
		} else if minDigit2 < 0 {
			minDigit2 = v
		}
	}

	min1, _ := strconv.ParseFloat(fmt.Sprintf("%d%d", minDigit1, minDigit2), 64)
	min1diff := math.Abs(60.0 - min1)
	min2, _ := strconv.ParseFloat(fmt.Sprintf("%d%d", minDigit2, minDigit1), 64)
	min2diff := math.Abs(60.0 - min2)

	if min1diff < min2diff {
		lmh = min1
	} else {
		lmh = min2
	}

	return fmt.Sprintf("%f:%f", lhh, lmh)
}
