// Package mergeintervals
// https://leetcode.com/problems/merge-intervals/description/
// time: O(n log n) / space: O(n)
package mergeintervals

import "sort"

func merge(intervals [][]int) [][]int {
	// sort
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := [][]int{intervals[0]}
	// find

	for _, cur := range intervals[1:] {
		last := merged[len(merged)-1]

		if cur[0] <= last[1] { // overlap
			last[1] = max(cur[1], last[1])
		} else {
			merged = append(merged, cur)
		}
	}

	return merged
}
