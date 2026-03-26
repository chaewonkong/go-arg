package taskscheduler

import (
	"fmt"
	"testing"
)

func TestLeastInterval(t *testing.T) {
	for idx, tt := range []struct {
		tasks []byte
		n     int
		want  int
	}{
		{tasks: []byte{'A', 'A', 'A', 'B', 'B', 'B'}, n: 2, want: 8},
		{tasks: []byte{'A', 'C', 'A', 'B', 'D', 'B'}, n: 1, want: 6},
	} {
		t.Run(fmt.Sprint(idx), func(t *testing.T) {
			if got := leastInterval(tt.tasks, tt.n); got != tt.want {
				t.Errorf("%d, expected: %d, got: %d", idx, tt.want, got)
			}
		})
	}
}
