package largesthist

import (
	"fmt"
	"testing"
)

func TestLargestRectangleArea(t *testing.T) {
	for idx, tc := range []struct {
		heights []int
		want    int
	}{
		{heights: []int{2, 1, 5, 6, 2, 3}, want: 10},
		{heights: []int{2, 4}, want: 4},
		{heights: []int{3, 4, 5, 2, 3, 2, 1}, want: 12},
	} {
		t.Run(fmt.Sprint(idx), func(t *testing.T) {
			if got := largestRectangleArea(tc.heights); got != tc.want {
				t.Errorf("expected: %d, got: %d", tc.want, got)
			}
		})
	}
}
