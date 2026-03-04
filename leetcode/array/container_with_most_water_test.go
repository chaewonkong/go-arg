package array

import (
	"fmt"
	"testing"
)

func TestMaxArea(t *testing.T) {
	for i, tt := range []struct {
		height  []int
		maxArea int
	}{
		{height: []int{0, 0, 0, 0, 0}, maxArea: 0},
		{height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, maxArea: 49},
		{height: []int{2, 2}, maxArea: 2},
		{height: []int{1, 1}, maxArea: 1},
	} {
		t.Run(fmt.Sprintf("case: %d", i), func(t *testing.T) {
			if actual := maxArea(tt.height); actual != tt.maxArea {
				t.Errorf("expected: %d, got: %d", tt.maxArea, actual)
			}
		})
	}
}
