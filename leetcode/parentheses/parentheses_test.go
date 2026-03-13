package parentheses

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateParentheses(t *testing.T) {
	for idx, tt := range []struct {
		input int
		want  []string
	}{
		{input: 1, want: []string{"()"}},
		{input: 3, want: []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
	} {
		t.Run(fmt.Sprint(idx), func(t *testing.T) {
			got := generateParenthesis(tt.input)
			assert.ElementsMatch(t, tt.want, got)
		})
	}
}
