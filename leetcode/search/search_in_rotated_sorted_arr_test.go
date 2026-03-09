package search

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	for idx, tt := range []struct {
		nums   []int
		target int
		want   int
	}{
		{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 0, want: 4},
		{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 3, want: -1},
		{nums: []int{1, 2, 3, 4, 5, 6}, target: 3, want: 2},
		{nums: []int{1}, target: 0, want: -1},
		{nums: []int{1}, target: 1, want: 0},
	} {
		t.Run(fmt.Sprint(idx), func(t *testing.T) {
			if got := binSearchWithPivot(tt.nums, tt.target); got != tt.want {
				t.Errorf("[CASE %d] expected: %d, got: %d", idx, tt.want, got)
			}
		})
	}
}
