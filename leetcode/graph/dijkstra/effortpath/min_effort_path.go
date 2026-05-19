// Package effortpath
//
//   - https://leetcode.com/problems/path-with-minimum-effort/
//   - time: O(M*N*log(M*N)); for M*N times, push/pop log(M*N)
//   - space: O(M*N);
package effortpath

import (
	"container/heap"
	"math"
)

/*
 Dijkstra Time Complexity:
  - for all nodes, pop at least 1 times: V * log(V)
  - for all edges, if newEffort < dist[nr][nc]; then push: E * log(V)
  - O((V+E)*log(V))
  - V = M*N
  - E = 4 * M * N (up, down, left, right)
  - O((M×N + 4×M×N) × log(M×N)) = O((M*N) * log(M*N))
*/

func minimumEffortPath(heights [][]int) int {
	rows, cols := len(heights), len(heights[0])

	dist := make([][]int, rows)
	for i := range dist {
		dist[i] = make([]int, cols)
		for j := range dist[i] {
			dist[i][j] = math.MaxInt32
		}
	}
	// start
	dist[0][0] = 0

	h := &minHeap{{0, 0, 0}}
	heap.Init(h)

	directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	for h.Len() > 0 {
		cur := heap.Pop(h).([3]int)
		effort, r, c := cur[0], cur[1], cur[2]

		// arrived in dest
		if r == rows-1 && c == cols-1 {
			return effort
		}

		if effort > dist[r][c] {
			continue
		}

		for _, d := range directions {
			nr, nc := r+d[0], c+d[1]
			if nr < 0 || nc < 0 || nr >= rows || nc >= cols {
				continue //out of bound
			}

			newEffort := max(effort, abs(heights[nr][nc]-heights[r][c]))
			if newEffort < dist[nr][nc] {
				dist[nr][nc] = newEffort
				heap.Push(h, [3]int{newEffort, nr, nc})
			}
		}
	}

	return dist[rows-1][cols-1]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

var _ heap.Interface = (*minHeap)(nil)

type minHeap [][3]int

// Len implements [heap.Interface].
func (m minHeap) Len() int {
	return len(m)
}

// Less implements [heap.Interface].
func (m minHeap) Less(i int, j int) bool {
	return m[i][0] < m[j][0]
}

// Pop implements [heap.Interface].
func (m *minHeap) Pop() any {
	old := *m
	n := len(old)
	x := old[n-1]
	old = old[:n-1]
	*m = old
	return x
}

// Push implements [heap.Interface].
func (m *minHeap) Push(x any) {
	*m = append(*m, x.([3]int))
}

// Swap implements [heap.Interface].
func (m minHeap) Swap(i int, j int) {
	m[i], m[j] = m[j], m[i]
}
