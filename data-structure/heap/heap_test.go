package heap

import (
	"fmt"
	"testing"
)

func TestHeapRoundtrip(t *testing.T) {
	for idx, tt := range []struct {
		inserts []int
	}{
		{inserts: []int{}},
		{inserts: []int{5, 4, 3, 2, 1}},
		{inserts: []int{1, 2, 3, 4, 5}},
		{inserts: []int{-10, 100, 20, 35}},
		{inserts: []int{1, 9, 10, 3, 4, 7, 2, 5, 8, 6}},
	} {
		t.Run(fmt.Sprint(idx), func(t *testing.T) {
			h := New()
			for _, v := range tt.inserts {
				h.Insert(v)
			}

			minV, ok := h.RemoveMax()
			if !ok {
				return
			}
			for {
				v, ok := h.RemoveMax()
				if !ok {
					break
				}

				if v > minV {
					t.Errorf("element not in descending order, case : %d", idx)
				}
				minV = v
			}
		})
	}
}
