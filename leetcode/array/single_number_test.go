package array

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleNumber(t *testing.T) {
	for i, tt := range []struct {
		input  []int
		output int
	}{
		{input: []int{2, 2, 1}, output: 1},
		{input: []int{4, 1, 2, 1, 2}, output: 4},
		{input: []int{1}, output: 1},
	} {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			assert.Equal(t, tt.output, singleNumber(tt.input))
		})
	}
}
