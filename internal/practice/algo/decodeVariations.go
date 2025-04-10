package algo

import (
	"strconv"
)

func decodeVariations(s string) int {
	if s == "" {
		return 0
	}

	n := len(s)

	dp := make([]int, n+1)
	dp[n] = 1

	for i := n - 1; i >= 0; i-- {
		if s[i] == '0' {
			continue
		}

		dp[i] = dp[i+1]

		if i+1 < n {
			val, _ := strconv.Atoi(s[i : i+2])
			if val <= 26 {
				dp[i] += dp[i+2]
			}
		}
	}

	return dp[0]
}
