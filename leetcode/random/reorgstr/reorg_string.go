// Package reorgstr
//
//   - https://leetcode.com/problems/reorganize-string/description/
//   - time: O(N)
//   - space: O(1)
package reorgstr

import (
	"container/heap"
)

func reorganizeString(s string) string {
	maxH := &maxHeap{}
	heap.Init(maxH)

	freq := [26]int{}
	for _, c := range s {
		freq[c-'a']++
	}

	for i, f := range freq {
		if f == 0 {
			continue
		}
		ch := character{
			code:      i,
			frequency: f,
		}
		heap.Push(maxH, ch)
	}

	c := heap.Pop(maxH).(character)
	n := len(s)
	if c.frequency > (n+1)/2 {
		return ""
	}
	heap.Push(maxH, c)

	result := []byte{}
	prev := character{code: -1}
	for maxH.Len() > 0 {
		cur := heap.Pop(maxH).(character)
		result = append(result, byte('a'+cur.code))
		cur.frequency--

		if prev.frequency > 0 {
			heap.Push(maxH, prev)
		}
		prev = cur
	}

	return string(result)
}

type character struct {
	code      int
	frequency int
}

type maxHeap []character

// Len implements heap.Interface.
func (m maxHeap) Len() int {
	return len(m)
}

// Less implements heap.Interface.
func (m maxHeap) Less(i int, j int) bool {
	return m[i].frequency > m[j].frequency
}

// Pop implements heap.Interface.
func (m *maxHeap) Pop() any {
	old := *m
	n := len(old)
	v := old[n-1]
	*m = old[:n-1]

	return v
}

// Push implements heap.Interface.
func (m *maxHeap) Push(x any) {
	*m = append(*m, x.(character))
}

// Swap implements heap.Interface.
func (m maxHeap) Swap(i int, j int) {
	m[i], m[j] = m[j], m[i]
}
