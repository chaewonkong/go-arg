package longestrepl

import "testing"

func TestCharacterReplacement(t *testing.T) {
	for _, tt := range []struct {
		s        string
		k        int
		expected int
	}{
		{"ABAB", 2, 4},
		{"AABABBA", 1, 4},
		{"A", 1, 1},
		{"ABBB", 2, 4},
		{"AAABBBB", 2, 6},
	} {
		t.Run(tt.s, func(t *testing.T) {
			if got := characterReplacement(tt.s, tt.k); got != tt.expected {
				t.Errorf("expected: %d, got: %d", tt.expected, got)
			}
		})
	}
}
