package algo

import (
	psort "enkya.org/playground/internal/practice/sort"
)

func hIndex(citations []int) int {
	n := len(citations)
	if n == 0 {
		return 0
	}

	sorted := psort.RecursiveMergeSort(citations)

	phIndex := n
	hIndex := 0

	for i := 0; i < n; i++ {
		hIndex = max(hIndex, min(sorted[i], phIndex))
		phIndex--
	}

	return hIndex
}
