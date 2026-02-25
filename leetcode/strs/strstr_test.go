package strs

import (
	"fmt"
	"testing"
)

func TestStrStr(t *testing.T) {
	for _, tt := range []struct {
		needle, haystack string
		expected         int
	}{
		{haystack: "sadbutsad", needle: "sad", expected: 0},
		{haystack: "leetcode", needle: "leeto", expected: -1},
		{haystack: "a", needle: "a", expected: 0},
		{haystack: "aaa", needle: "a", expected: 0},
		{haystack: "ace", needle: "ce", expected: 1},
		{haystack: "sabactuallysad", needle: "sad", expected: 11},
		{haystack: "mississippi", needle: "issip", expected: 4},
		{haystack: "aaa", needle: "aaaa", expected: -1},
	} {
		t.Run(fmt.Sprintf("needle: %s, haystack: %s", tt.needle, tt.haystack), func(t *testing.T) {
			if actual := strStr(tt.haystack, tt.needle); actual != tt.expected {
				t.Errorf("expected: %d, got: %d", tt.expected, actual)
			}
		})
	}
}
