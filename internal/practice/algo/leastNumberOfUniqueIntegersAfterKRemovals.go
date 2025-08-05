package algo

import "sort"

// Given an array of integers and an integer k, remove exactly k elements from the array so that the number of unique integers in the array is minimized.

// Return the minimum number of unique integers left in the array after k removals.

/*
loop through and generate frequency map
then generate sorted list of map
then remove lowest frequency values until you use up k
the result should be the ints with the highest frequency and therefore minimise uniqueness
*/

func FindLeastNumOfUniqueInts(arr []int, k int) int {
	freqMap := make(map[int]int)
	for _, v := range arr { // 2,4,1,3,3,3,2,2,2
		if _, exists := freqMap[v]; !exists {
			freqMap[v] = 0
		}
		freqMap[v] += 1
	}
	// 2:4, 4:1, 1:1 3:3

	frequencies := make([]int, 0)
	for _, v := range freqMap {
		frequencies = append(frequencies, v)
	}

	sort.Ints(frequencies) // 1,1,3,4
	i := k                 // 6
	for i > 0 {
		i -= frequencies[0]
		if i >= 0 {
			frequencies = frequencies[1:]
		}
	}

	return len(frequencies) // 1
}

/*
using a min heap
get frequencies in map
insert frequencies into minheap
if fre
*/
