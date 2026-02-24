package array

import "testing"

func TestThreeSum(t *testing.T) {
	res := threeSum([]int{-1, 0, 1, 2, -1, 4})
	for _, item := range res {
		t.Log(item)
	}
}
