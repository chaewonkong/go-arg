package strs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstUniqChar(t *testing.T) {
	for _, tt := range []struct {
		input    string
		expected int
	}{
		{input: "leetcode", expected: 0},
		{input: "loveleetcode", expected: 2},
		{input: "aabb", expected: -1},
		{input: "a", expected: 0},
	} {
		t.Run(tt.input, func(t *testing.T) {
			actual := firstUniqChar(tt.input)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
