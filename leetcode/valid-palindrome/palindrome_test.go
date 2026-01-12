package validpalindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	for _, tt := range []struct {
		s        string
		expected bool
	}{
		{
			s:        "A man, a plan, a canal: Panama",
			expected: true,
		},
		{
			s:        ",",
			expected: true,
		},
		{
			s:        "abc010cba",
			expected: true,
		},
		{
			s:        " ",
			expected: true,
		},
		{
			s:        "race a car",
			expected: false,
		},
		{
			s:        "palindrome",
			expected: false,
		},
	} {
		t.Run(tt.s, func(t *testing.T) {
			assert.Equal(t, tt.expected, isPalindrome(tt.s))
		})

	}
}
