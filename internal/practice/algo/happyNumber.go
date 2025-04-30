package algo

func isHappy(n int) bool {
	memo := make(map[int]bool)

	for n != 1 {
		if _, ok := memo[n]; ok {
			return false
		}

		memo[n] = true
		n = sumOfSquares(n)
	}

	return true
}

func sumOfSquares(n int) int {
	sum := 0

	for n > 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}

	return sum
}
