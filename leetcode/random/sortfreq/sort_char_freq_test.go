package sortfreq

import (
	"fmt"
	"testing"
)

func TestFrequencySort(t *testing.T) {
	for idx, tt := range []struct {
		s, expected string
	}{
		{"tree", "eert"},
		{"cccaaa", "aaaccc"},
		{"Aabb", "bbAa"},
	} {
		t.Run(fmt.Sprintf("case: %d", idx), func(t *testing.T) {
			if got := frequencySort(tt.s); got != tt.expected {
				t.Errorf("expected: %s, got: %s", tt.expected, got)
			}
		})
	}
}
