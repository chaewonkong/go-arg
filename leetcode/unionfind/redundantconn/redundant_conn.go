// Package redundantconn
//
//   - https://leetcode.com/problems/redundant-connection/description/
//   - Time Complexity: O(n · α(n)) ≈ O(n)
//   - Space Complexity: O(n)
package redundantconn

/* Union Find
for each edge (u, v):
  if find(u) == find(v): // 이미 연결됨. 정답.
  else:
    union(u, v)
* */

func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	parent := make([]int, n+1)
	for i := range parent {
		parent[i] = i
	}

	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x]) // return find(parent[x]) -> path compression
		}
		return parent[x]
	}

	union := func(x, y int) bool {
		px, py := find(x), find(y)
		if px == py {
			// same root -> circular
			return false
		}

		parent[py] = px
		return true
	}

	for _, edge := range edges {
		if !union(edge[0], edge[1]) {
			return edge
		}
	}

	return nil
}
