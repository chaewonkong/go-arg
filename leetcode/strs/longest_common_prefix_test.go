package strs

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	for _, tt := range []struct {
		input    []string
		expected string
	}{
		{input: []string{"flower", "flow", "flight"}, expected: "fl"},
		{input: []string{"dog", "racecar", "car"}, expected: ""},
		{input: []string{"flower", "flow", "flight", ""}, expected: ""},
	} {
		t.Run(tt.expected, func(t *testing.T) {
			if longestCommonPrefix(tt.input) != tt.expected {
				t.Errorf("err test failed; expected: %q, actual: %q", tt.expected, tt.input)
			}
		})
	}
}
