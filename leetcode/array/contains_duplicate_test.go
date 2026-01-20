package array

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsDuplicate(t *testing.T) {
	for _, tt := range []struct {
		arr      []int
		expected bool
	}{
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 2, 3, 3}, true},
		{[]int{1, 2, 3, 4}, false},
		{[]int{-100000000, -20000, 3, 1, 4, 3}, true},
		{[]int{1}, false},
	} {
		t.Run(fmt.Sprintf("arr: %v", tt.arr), func(t *testing.T) {
			assert.Equal(t, tt.expected, containsDuplicate(tt.arr))
		})
	}
}
