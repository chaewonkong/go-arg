package findmedian

import "container/heap"

type MedianFinder struct {
	left  *MaxHeap
	right *MinHeap
}

func Constructor() MedianFinder {
	return MedianFinder{
		left:  &MaxHeap{},
		right: &MinHeap{},
	}
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(this.left, num)

	heap.Push(this.right, heap.Pop(this.left))

	if this.left.Len() < this.right.Len() {
		heap.Push(this.left, heap.Pop(this.right))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.left.Len() > this.right.Len() {
		return float64((*this.left)[0])
	}

	return float64((*this.left)[0]+(*this.right)[0]) / 2.0
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

var _ heap.Interface = (*MinHeap)(nil)

type MinHeap []int

// Len implements heap.Interface.
func (m MinHeap) Len() int {
	return len(m)
}

// Less implements heap.Interface.
func (m MinHeap) Less(i int, j int) bool {
	return m[i] < m[j]
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
	*m = append(*m, x.(int))
}

// Swap implements heap.Interface.
func (m MinHeap) Swap(i int, j int) {
	m[i], m[j] = m[j], m[i]
}

var _ heap.Interface = (*MaxHeap)(nil)

type MaxHeap []int

// Len implements heap.Interface.
func (m MaxHeap) Len() int {
	return len(m)
}

// Less implements heap.Interface.
func (m MaxHeap) Less(i int, j int) bool {
	return m[i] > m[j]
}

// Pop implements heap.Interface.
func (m *MaxHeap) Pop() any {
	old := *m
	n := len(old)
	val := old[n-1]
	*m = old[:n-1]

	return val
}

// Push implements heap.Interface.
func (m *MaxHeap) Push(x any) {
	*m = append(*m, x.(int))
}

// Swap implements heap.Interface.
func (m MaxHeap) Swap(i int, j int) {
	m[i], m[j] = m[j], m[i]
}
