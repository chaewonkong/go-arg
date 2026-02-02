package array

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRotateArray(t *testing.T) {
	for i, tt := range []struct {
		nums     []int
		k        int
		expected []int
	}{
		{nums: []int{1, 2, 3, 4, 5, 6, 7}, k: 3, expected: []int{5, 6, 7, 1, 2, 3, 4}},
		{nums: []int{-1, -100, 3, 99}, k: 2, expected: []int{3, 99, -1, -100}},
		{nums: []int{1, 2, 3, 4, 5, 6, 7, 8}, k: 3, expected: []int{6, 7, 8, 1, 2, 3, 4, 5}},
		{nums: []int{99}, k: 1, expected: []int{99}},
		{nums: []int{99}, k: 2, expected: []int{99}},
		{nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53}, k: 82, expected: []int{25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24}},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			rotate(tt.nums, tt.k)
			require.Equal(t, len(tt.expected), len(tt.nums))
			for i := range len(tt.expected) {
				assert.Equal(t, tt.expected[i], tt.nums[i])
			}
		})
	}
}

func TestReverse(t *testing.T) {
	for i, tt := range []struct {
		nums, expected []int
		start, end     int
	}{
		{nums: []int{1, 2, 3, 4, 5, 6, 7}, start: 0, end: 6, expected: []int{7, 6, 5, 4, 3, 2, 1}},
		{nums: []int{7, 6, 5, 4, 3, 2, 1}, start: 0, end: 2, expected: []int{5, 6, 7, 4, 3, 2, 1}},
		{nums: []int{5, 6, 7, 4, 3, 2, 1}, start: 3, end: 6, expected: []int{5, 6, 7, 1, 2, 3, 4}},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			reverse(tt.nums, tt.start, tt.end)
			for i := range len(tt.expected) {
				assert.Equal(t, tt.expected[i], tt.nums[i])
			}
		})
	}
}
