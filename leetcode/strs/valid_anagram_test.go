package strs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidAnagram(t *testing.T) {
	for _, tt := range []struct {
		a, b     string
		expected bool
	}{
		{a: "anagram", b: "nagaram", expected: true},
		{a: "a", b: "a", expected: true},
		{a: "rat", b: "car", expected: false},
		{a: "rat", b: "art", expected: true},
	} {
		t.Run(tt.a, func(t *testing.T) {

			assert.Equal(t, tt.expected, isAnagram(tt.a, tt.b))
		})
	}
}
