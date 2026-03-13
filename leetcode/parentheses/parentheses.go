package parentheses

func generateParenthesis(n int) []string {
	results := []string{}

	var dfs func(path string, open, close int)
	dfs = func(path string, open, close int) {
		if len(path) == 2*n { // complete
			results = append(results, path)
			return
		}

		if open < n { // can open
			dfs(path+"(", open+1, close)
		}

		if close < open { // can close
			dfs(path+")", open, close+1)
		}
	}

	dfs("", 0, 0)

	return results
}
