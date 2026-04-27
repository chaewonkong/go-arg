// Package mineatingspeed
//
//   - https://leetcode.com/problems/koko-eating-bananas/description/
//   - time: O(NlogM); M = max(piles)
//   - space: O(1)
//   - 패턴: 이진탐색으로 최소/최대 찾기. 찾으면 return하지 않고 계속 찾음.
package mineatingspeed

import "slices"

func minEatingSpeed(piles []int, h int) int {
	lo, hi := 1, slices.Max(piles)

	canFinish := func(piles []int, k, h int) bool {
		total := 0
		for _, p := range piles {
			v := p / k
			if p%k > 0 {
				v++
			}
			total += v
		}

		return total <= h
	}

	for lo < hi {
		mid := (lo + hi) / 2
		if canFinish(piles, mid, h) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}
