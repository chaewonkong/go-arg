package substr

import (
	"fmt"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	for i, tt := range []struct {
		s        string
		expected int
	}{
		{s: "abcabcbb", expected: 3},
		{s: "pwwkew", expected: 3},
		{s: "bbbbb", expected: 1},
		{s: "#!123_-&*  ssabc", expected: 10},
		{s: "", expected: 0},
		{s: "a", expected: 1},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			if actual := lengthOfLongestSubstringWithSlidingWindow(tt.s); actual != tt.expected {
				t.Errorf("expected: %d, actual: %d", tt.expected, actual)
			}
		})
	}
}
