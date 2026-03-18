package dailytemp

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDailyTemperajures(t *testing.T) {
	for idx, tt := range []struct {
		temps []int
		want  []int
	}{
		{temps: []int{73, 74, 75, 71, 69, 72, 76, 73}, want: []int{1, 1, 4, 2, 1, 1, 0, 0}},
	} {
		t.Run(fmt.Sprint(idx), func(t *testing.T) {
			got := dailyTemperaturesWithStack(tt.temps)

			assert.Equal(t, len(tt.want), len(got), "len should be equal")
			for i := range len(tt.want) {
				assert.Equal(t, tt.want[i], got[i], "expected: %d, got: %d", tt.want[i], got[i])
			}
		})
	}
}
