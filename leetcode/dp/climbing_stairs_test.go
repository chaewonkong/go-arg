package dp

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClimbingStairs(t *testing.T) {
	for _, tt := range []struct {
		input    int
		expected int
	}{
		{input: 1, expected: 1},
		{input: 2, expected: 2},
		{input: 3, expected: 3},
		{input: 4, expected: 5},
		{input: 5, expected: 8},
		{input: 6, expected: 13},
	} {
		t.Run(fmt.Sprintf("n = %d", tt.input), func(t *testing.T) {
			assert.Equal(t, tt.expected, climbStairs(tt.input))
		})
	}
}
