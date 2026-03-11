package permutations

import "slices"

func permute(nums []int) [][]int {
	n := len(nums)
	results := [][]int{}
	path := make([]int, 0, n)
	used := make([]bool, n)

	var dfs func()

	dfs = func() {
		if len(path) == n {
			results = append(results, slices.Clone(path))
			return
		}

		for i := range n {
			if used[i] {
				continue
			}

			used[i] = true
			path = append(path, nums[i])
			dfs()
			path = path[:len(path)-1]
			used[i] = false
		}
	}

	dfs()

	return results
}
