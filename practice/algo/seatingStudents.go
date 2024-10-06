package algo

/*
Seating Students

Have the function SeatingStudents(arr) read the array of integers stored in arr which will be in the following format:
[K, r1, r2, r3, ...] where K represents the number of desks in a classroom, and the rest of the integers in the array will be in sorted order and will represent the desks that are already occupied. All of the desks will be arranged in 2 columns, where desk #1 is at the top left, desk #2 is at the top right, desk #3 is below #1, desk #4 is below #2, etc. Your program should return the number of ways 2 students can be seated next to each other. This means 1 student is on the left and 1 student on the right, or 1 student is directly above or below the other student.

For example: if arr is [12, 2, 6, 7, 11] then this classrooms looks like the following picture:
[Picture]

Based on above arrangement of occupied desks, there are a total of 6 ways to seat 2 new students next to each other. The combinations are: 1 & 3, 1 & 9, 1 & 10, 4 & 6, 4 & 9, 5 & 7. So for this input your program should return 6. K will range from 2 to 24 and will always be an even number. After K, the number of occupied desks in the array can range from 0 to K.

Examples
Input: []int{6, 4}
Output: 4
Input: []int{8, 1, 8}
Output: 6

https://www.coderbyte.com/results/bhanson:Seating%20Students:Go
*/

func SeatingStudents(arr []int) int {
	K := arr[0]
	occupied := arr[1:]

	count := 0
	for i := 1; i <= K; i++ {
		if contains(occupied, i) {
			continue
		}

		if i%2 == 0 {
			if !contains(occupied, i-1) {
				count++
			}

			if i+2 <= K && !contains(occupied, i+2) {
				count++
			}

		} else {
			if i+2 <= K && !contains(occupied, i+2) {
				count++
			}
		}
	}

	return count
}

func contains(arr []int, num int) bool {
	for _, n := range arr {
		if n == num {
			return true
		}
	}
	return false
}
