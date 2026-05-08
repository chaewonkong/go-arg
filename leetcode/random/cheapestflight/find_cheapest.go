// Package cheapestflight
//
//   - https://leetcode.com/problems/cheapest-flights-within-k-stops/
//   - time: O(k * E)
//   - space: O(n * k + E)
package cheapestflight

import (
	"math"
)

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	graph := make(map[int][][2]int)
	for _, fl := range flights {
		s, d, p := fl[0], fl[1], fl[2]
		graph[s] = append(graph[s], [2]int{d, p})
	}

	// memo[node][stopsLeft] = 이 상태에서 도달했을 때의 최소 total cost
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, k+2)
		for j := range memo[i] {
			memo[i][j] = math.MaxInt32
		}
	}

	minP := math.MaxInt32
	var dfs func(node, total, left int)
	dfs = func(node, total, left int) {
		if total >= minP {
			return
		}

		// 같은 (node, left) 상태를 더 높은 cost로 재방문하면 스킵
		if total >= memo[node][left] {
			return
		}
		memo[node][left] = total

		if node == dst {
			minP = total
			return
		}
		if left == 0 {
			return
		}
		for _, edge := range graph[node] {
			v, p := edge[0], edge[1]
			dfs(v, total+p, left-1)
		}
	}
	dfs(src, 0, k+1)
	if minP == math.MaxInt32 {
		return -1
	}
	return minP
}

func bellman_ford(n int, flights [][]int, src int, dst int, k int) int {
	prices := make([]int, n)
	for i := range prices {
		prices[i] = math.MaxInt32
	}
	prices[src] = 0 // start pos

	for range k + 1 {
		tmp := append([]int{}, prices...)

		for _, fl := range flights {
			s, d, p := fl[0], fl[1], fl[2]

			if prices[s] == math.MaxInt32 {
				continue
			}

			if prices[s]+p < tmp[d] {
				tmp[d] = prices[s] + p
			}
		}

		prices = tmp
	}

	if prices[dst] == math.MaxInt32 {
		return -1
	}
	return prices[dst]
}
