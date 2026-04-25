// Package sortfreq
//
//   - https://leetcode.com/problems/sort-characters-by-frequency/
//   - time: O(N)
//   - space: O(N)
package sortfreq

import (
	"container/heap"
	"strings"
)

func frequencySort(s string) string {
	h := &frequency{}
	heap.Init(h)

	// count frequency
	freqMap := make(map[rune]int)
	for _, r := range s {
		freqMap[r]++
	}

	// sort by frequency
	for r, f := range freqMap {
		heap.Push(h, character{value: r, freq: f})
	}

	sb := &strings.Builder{}
	for h.Len() > 0 {
		c := heap.Pop(h).(character)
		for range c.freq {
			sb.WriteRune(c.value)
		}
	}
	return sb.String()
}

type character struct {
	value rune
	freq  int
}

var _ heap.Interface = (*frequency)(nil)

type frequency []character

// Len implements heap.Interface.
func (f frequency) Len() int {
	return len(f)
}

// Less implements heap.Interface; max heap
func (f frequency) Less(i int, j int) bool {
	if f[i].freq == f[j].freq {
		return f[i].value < f[j].value
	}
	return f[i].freq > f[j].freq
}

// Pop implements heap.Interface.
func (f *frequency) Pop() any {
	old := *f
	n := len(old)
	val := old[n-1]
	*f = old[:n-1]

	return val
}

// Push implements heap.Interface.
func (f *frequency) Push(x any) {
	*f = append(*f, x.(character))
}

// Swap implements heap.Interface.
func (f frequency) Swap(i int, j int) {
	f[i], f[j] = f[j], f[i]
}

/*
* With Bucket Sort
 */

func frequencyBucketSort(s string) string {
	h := &frequency{}
	heap.Init(h)

	// count frequency
	freqMap := make(map[rune]int)
	for _, r := range s {
		freqMap[r]++
	}

	buckets := make([][]rune, len(s)+1)
	for r, cnt := range freqMap {
		buckets[cnt] = append(buckets[cnt], r)
	}

	sb := &strings.Builder{}
	for i := len(buckets) - 1; i >= 1; i-- {
		for _, r := range buckets[i] {
			for range i {
				sb.WriteRune(r)
			}
		}
	}

	return sb.String()
}
