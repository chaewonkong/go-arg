package math

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {
	for _, tt := range []struct {
		n      int
		output []string
	}{
		{
			n:      3,
			output: []string{"1", "2", "Fizz"},
		},
		{
			n:      5,
			output: []string{"1", "2", "Fizz", "4", "Buzz"},
		},
		{
			n:      15,
			output: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
		},
		{
			n:      1,
			output: []string{"1"},
		},
	} {
		t.Run(strconv.Itoa(tt.n), func(t *testing.T) {
			result := fizzBuzz(tt.n)
			assert.Equal(t, tt.output, result)
		})
	}
}
