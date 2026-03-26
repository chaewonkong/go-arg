// Package mergeksorted
// https://leetcode.com/problems/merge-k-sorted-lists/description/
package mergeksorted

import "container/heap"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	minHeap := &MinHeap{}
	for _, l := range lists {
		if l != nil {
			heap.Push(minHeap, l)
		}
	}

	dummy := &ListNode{}
	cur := dummy
	for minHeap.Len() > 0 {
		node := heap.Pop(minHeap).(*ListNode)
		cur.Next = node
		cur = cur.Next

		if node.Next != nil {
			heap.Push(minHeap, node.Next)
		}
	}

	return dummy.Next
}

var _ (heap.Interface) = (*MinHeap)(nil)

type MinHeap []*ListNode

// Len implements heap.Interface.
func (m MinHeap) Len() int {
	return len(m)
}

// Less implements heap.Interface.
func (m MinHeap) Less(i int, j int) bool {
	return m[i].Val < m[j].Val
}

// Pop implements heap.Interface.
func (m *MinHeap) Pop() any {
	old := *m
	n := len(old)
	val := old[n-1]
	*m = old[:n-1]
	return val
}

// Push implements heap.Interface.
func (m *MinHeap) Push(x any) {
	*m = append(*m, x.(*ListNode))
}

// Swap implements heap.Interface.
func (m MinHeap) Swap(i int, j int) {
	m[i], m[j] = m[j], m[i]
}
