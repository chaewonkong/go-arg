package graph

import (
	"fmt"
	"testing"
)

func TestNumIslands(t *testing.T) {
	for idx, tt := range []struct {
		grid [][]byte
		want int
	}{
		{grid: [][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}, want: 1},
	} {
		t.Run(fmt.Sprint(idx), func(t *testing.T) {
			if got := numIslands(tt.grid); got != tt.want {
				t.Errorf("case: %d, want :%d, got: %d", idx, tt.want, got)
			}
		})
	}
}
