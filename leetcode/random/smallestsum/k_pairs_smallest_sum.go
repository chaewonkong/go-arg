// Package smallestsrm
//
//   - https://leetcode.com/problems/find-k-pairs-with-smallest-sums/description/
//   - time: O(Nlog(min(N, k)))
//   - space: O(min(N, k))
//     정렬된 배열들에서 k개를 순서대로 꺼내야 한다 => Heap
//     heap은 여러 후보 중 지금 당장 최선을 O(logN)에 뽑는 구조
package smallestsrm

import "container/heap"

/*
		초기 후보들을 힙에 push
	  while 결과 k개 미만:
	    pop 최솟값
	    결과에 추가
	    그 다음 후보(next state)를 push

*
*/
type item struct {
	sum, i, j int
}

var _ heap.Interface = (*minHeap)(nil)

type minHeap []item

// Len implements heap.Interface.
func (m minHeap) Len() int {
	return len(m)
}

// Less implements heap.Interface.
func (m minHeap) Less(i int, j int) bool {
	return m[i].sum < m[j].sum
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
	*m = append(*m, x.(item))
}

// Swap implements heap.Interface.
func (m minHeap) Swap(i int, j int) {
	m[i], m[j] = m[j], m[i]
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	h := &minHeap{}
	heap.Init(h)

	for i := 0; i < len(nums1) && i < k; i++ {
		heap.Push(h, item{nums1[i] + nums2[0], i, 0})
	}

	result := make([][]int, 0, k)
	for h.Len() > 0 && len(result) < k {
		cur := heap.Pop(h).(item)
		result = append(result, []int{nums1[cur.i], nums2[cur.j]})
		if cur.j+1 < len(nums2) {
			heap.Push(h, item{nums1[cur.i] + nums2[cur.j+1], cur.i, cur.j + 1})
		}
	}

	return result
}
