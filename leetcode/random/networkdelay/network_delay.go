// Package networkdelay
//
//   - https://leetcode.com/problems/network-delay-time/submissions/1997418551/
//   - Dijkstra
//   - Time: O((V+E)logV)
//   - Space: O(V+E)
package networkdelay

import (
	"container/heap"
	"math"
)

/* Dijkstra
1. 시작점 k부터 노드 i까지의 최단거리를 보관하는 배열 dist[i]를 만든다 (INF로 초기화, dist[k] = 0)
2. heap에 (score, node)를 묶어 저장한다. score 기준 ASC 정렬
3. heap에서 가장 가까운 노드부터 꺼낸다
4. 꺼낸 d가 dist[u]보다 크면 낡은 항목이므로 스킵
5. u에 연결된 모든 노드 v를 순회하며, dist[u] + w < dist[v]이면 dist[v] 갱신 후 heap에 push
6. dist[1]~dist[n] 순회하며 INF인 노드가 있으면 -1 반환, 아니면 최댓값 갱신
7. 최댓값 반환
*/

func networkDelayTime(times [][]int, n int, k int) int {
	graph := make(map[int][][2]int)

	for _, ti := range times {
		u, v, w := ti[0], ti[1], ti[2]
		graph[u] = append(graph[u], [2]int{v, w})
	}

	// stores k -> i min distance
	dist := make([]int, n+1)

	// initialize with INF
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[k] = 0

	h := &minHeap{{0, k}}
	heap.Init(h)

	for h.Len() > 0 {
		cur := heap.Pop(h).([2]int)
		d, u := cur[0], cur[1]

		if d > dist[u] {
			continue // already handled shorter distance
		}

		for _, edge := range graph[u] {
			v, w := edge[0], edge[1]
			if dist[u]+w < dist[v] { // found shorter distance
				dist[v] = dist[u] + w
				heap.Push(h, [2]int{dist[v], v})
			}
		}
	}

	maxT := 0
	for i := 1; i <= n; i++ {
		if dist[i] == math.MaxInt32 {
			return -1
		}
		maxT = max(maxT, dist[i])
	}

	return maxT
}

var _ heap.Interface = (*minHeap)(nil)

type minHeap [][2]int

// Len implements heap.Interface.
func (m minHeap) Len() int {
	return len(m)
}

// Less implements heap.Interface.
func (m minHeap) Less(i int, j int) bool {
	return m[i][0] < m[j][0]
}

// Pop implements heap.Interface.
func (m *minHeap) Pop() any {
	old := *m
	n := len(old)
	v := old[n-1]
	*m = old[:n-1]

	return v
}

// Push implements heap.Interface.
func (m *minHeap) Push(x any) {
	*m = append(*m, x.([2]int))
}

// Swap implements heap.Interface.
func (m minHeap) Swap(i int, j int) {
	m[i], m[j] = m[j], m[i]
}
