package anagram

import (
	"fmt"
	"slices"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	for _, tt := range []struct {
		s, p     string
		expected []int
	}{
		{"cbaebabacd", "abc", []int{0, 6}},
		{"abab", "ab", []int{0, 1, 2}},
	} {
		t.Run(fmt.Sprintf("s: %s, p: %s", tt.s, tt.p), func(t *testing.T) {
			got := findAnagrams(tt.s, tt.p)

			slices.Sort(tt.expected)
			slices.Sort(got)
			if !slices.Equal(tt.expected, got) {
				t.Errorf("expected: %v, got: %v", tt.expected, got)
			}
		})
	}
}
