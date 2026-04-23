// Package numprovinces
//
//   - https://leetcode.com/problems/number-of-provinces/description/
//   - time: O(N·α(N))
//   - space: O(N)
package numprovinces

func findCircleNum(isConnected [][]int) int {
	parent := make([]int, len(isConnected))
	for i := range parent {
		parent[i] = i
	}

	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	union := func(x, y int) {
		px, py := find(x), find(y)
		if px != py {
			parent[px] = py
		}
	}

	for i := range len(isConnected) {
		for j := range i { // while j < i; [j][i] == [i][j]
			if isConnected[i][j] == 1 {
				union(i, j)
			}
		}
	}

	cnt := 0
	for i := range parent {
		if find(i) == i { // root만 parent[x] == x인 점을 이용.
			cnt++
		}
	}

	return cnt
}
