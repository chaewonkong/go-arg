package sort

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	for i, tt := range []struct {
		nums1    []int
		nums2    []int
		m, n     int
		expected []int
	}{
		{nums1: []int{1, 2, 3, 0, 0, 0}, nums2: []int{2, 5, 6}, m: 3, n: 3, expected: []int{1, 2, 2, 3, 5, 6}},
		{nums1: []int{1}, nums2: []int{}, m: 1, n: 0, expected: []int{1}},
		{nums1: []int{0}, nums2: []int{1}, m: 0, n: 1, expected: []int{1}},
		{nums1: []int{}, nums2: []int{}, m: 0, n: 0, expected: []int{}},
		{nums1: []int{4, 5, 6, 0, 0, 0}, nums2: []int{1, 2, 3}, m: 3, n: 3, expected: []int{1, 2, 3, 4, 5, 6}},
		{nums1: []int{1, 2, 4, 5, 6, 0}, nums2: []int{3}, m: 5, n: 1, expected: []int{1, 2, 3, 4, 5, 6}},
		{nums1: []int{4, 0, 0, 0, 0, 0}, nums2: []int{1, 2, 3, 5, 6}, m: 1, n: 5, expected: []int{1, 2, 3, 4, 5, 6}},
	} {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			merge(tt.nums1, tt.m, tt.nums2, tt.n)
			assert.Equal(t, len(tt.expected), len(tt.nums1))
			for i, v := range tt.nums1 {
				assert.Equal(t, tt.expected[i], v)
			}
		})
	}
}
