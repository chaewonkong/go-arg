package nextgreatel

import (
	"fmt"
	"testing"
)

func TestNextGreatElement(t *testing.T) {
	for idx, tt := range []struct {
		nums1, nums2, want []int
	}{
		{nums1: []int{4, 1, 2}, nums2: []int{1, 3, 4, 2}, want: []int{-1, 3, -1}},
		{nums1: []int{2, 4}, nums2: []int{1, 2, 3, 4}, want: []int{3, -1}},
	} {
		t.Run(fmt.Sprint(idx), func(t *testing.T) {
			got := nextGreaterElement(tt.nums1, tt.nums2)
			if len(tt.want) != len(got) {
				t.Errorf("expected len:%d, got: %d", len(tt.want), len(got))
			}

			for i := range len(tt.want) {
				if tt.want[i] != got[i] {
					t.Errorf("expected v: %d, got: %d", tt.want[i], got[i])
				}
			}
		})
	}
}
