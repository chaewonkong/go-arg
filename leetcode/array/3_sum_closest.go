package array

import (
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	result := nums[n-1] + nums[n-2] + nums[n-3] // biggest

	for i := range n - 2 {
		left := i + 1
		right := n - 1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum == target {
				return sum
			}

			if diff(result, target) > diff(sum, target) {
				result = sum
			}

			if sum < target {
				left++
			} else {
				right--
			}
		}
	}

	return result
}

func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
