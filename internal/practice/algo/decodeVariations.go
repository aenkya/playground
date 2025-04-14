package algo

func decodeVariationsBruteforce(s string) int {
	return decodeVariationsRecursive(s, 0)
}

func decodeVariationsRecursive(s string, index int) int {
	n := len(s)

	if index == n {
		return 1
	}

	if s[index] == '0' {
		return 0
	}

	count := decodeVariationsRecursive(s, index+1)

	if index+1 < n {
		val := (int(s[index]-'0') * 10) + int(s[index+1]-'0')
		if 10 < val && val <= 26 {
			count += decodeVariationsRecursive(s, index+2)
		}
	}

	return count
}

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
			val := (int(s[i]-'0') * 10) + int(s[i+1]-'0')
			if val <= 26 {
				dp[i] += dp[i+2]
			}
		}
	}

	return dp[0]
}
