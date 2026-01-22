package array

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMissingNumber(t *testing.T) {
	for _, tt := range []struct {
		nums     []int
		expected int
	}{
		{nums: []int{3, 0, 1}, expected: 2},
		{nums: []int{0, 1}, expected: 2},
		{nums: []int{9, 6, 4, 2, 3, 5, 7, 0, 1}, expected: 8},
		{nums: []int{0}, expected: 1},
		{nums: []int{1}, expected: 0},
		{nums: []int{1, 2}, expected: 0},
	} {
		t.Run(fmt.Sprintf("array: %v", tt.nums), func(t *testing.T) {
			assert.Equal(t, tt.expected, missingNumber(tt.nums))
		})
	}
}
