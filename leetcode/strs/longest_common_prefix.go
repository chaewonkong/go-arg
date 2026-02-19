package strs

import (
	"math"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	b := &strings.Builder{}

	minLen := math.MaxInt
	for _, s := range strs {
		minLen = min(minLen, len(s))
	}

	if minLen == 0 {
		return ""
	}

	for i := range minLen {
		target := strs[0][i]
		for _, str := range strs {
			if str[i] != target {
				return b.String()
			}
		}
		b.WriteByte(target)
	}

	return b.String()
}
