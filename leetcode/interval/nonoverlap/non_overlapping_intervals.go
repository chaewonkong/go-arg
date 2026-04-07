package nonoverlap

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	end := intervals[0][1] // init
	cnt := 0

	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]

		// compare
		if cur[0] < end {
			cnt++
		} else {
			end = cur[1]
		}
	}

	return cnt
}
