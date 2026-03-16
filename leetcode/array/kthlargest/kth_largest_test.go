package kthlargest

import (
	"fmt"
	"testing"
)

func TestFindKthLargest(t *testing.T) {
	for idx, tt := range []struct {
		arr  []int
		k    int
		want int
	}{
		{arr: []int{3, 2, 3, 1, 2, 4, 5, 5, 6}, k: 4, want: 4},
		{arr: []int{3, 2, 1, 5, 6, 4}, k: 2, want: 5},
	} {
		t.Run(fmt.Sprint(idx), func(t *testing.T) {
			if got := findKthLargest(tt.arr, tt.k); got != tt.want {
				t.Errorf("%d, want: %d, got: %d", idx, tt.want, got)
			}
		})
	}
}
