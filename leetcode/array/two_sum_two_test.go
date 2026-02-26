package array

import (
	"slices"
	"strconv"
	"testing"
)

func TestTwoSumTwo(t *testing.T) {
	for i, tt := range []struct {
		numbers, expected []int
		target            int
	}{
		{numbers: []int{2, 7, 11, 15}, expected: []int{1, 2}, target: 9},
		{numbers: []int{2, 3, 4}, expected: []int{1, 3}, target: 6},
		{numbers: []int{-1, 0}, expected: []int{1, 2}, target: -1},
		{numbers: []int{-100, 38, 39, 40, 100, 105, 123, 144}, expected: []int{1, 5}, target: 0},
	} {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if actual := twoSumtwo(tt.numbers, tt.target); !slices.Equal(tt.expected, actual) {
				t.Errorf("expected: %v, actual:%v ", tt.expected, actual)
			}
		})
	}
}
