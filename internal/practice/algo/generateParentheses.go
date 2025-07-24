package algo

func generateParentheses(n int) []string {
	res := make([]string, 0)

	dfs(&res, n, 0, 0, "")

	return res
}

func dfs(res *[]string, n, l, r int, s string) {
	if len(s) == n*2 {
		*res = append(*res, s)
	}

	if l < n {
		dfs(res, n, l+1, r, s+"(")
	}

	if r < l {
		dfs(res, n, l, r+1, s+")")
	}
}
