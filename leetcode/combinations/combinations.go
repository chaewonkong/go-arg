package combinations

import "slices"

func combine(n int, k int) [][]int {

	results := [][]int{}

	var dfs func([]int, int)
	dfs = func(candidates []int, start int) {
		if len(candidates) == k {
			results = append(results, slices.Clone(candidates))
			return
		}

		for i := start; i <= n; i++ {
			candidates = append(candidates, i)
			dfs(candidates, i+1)
			candidates = candidates[:len(candidates)-1] // pop
		}

	}

	dfs([]int{}, 1)

	return results
}
