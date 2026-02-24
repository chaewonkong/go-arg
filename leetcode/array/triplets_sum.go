package array

import "sort"

func threeSum(nums []int) [][]int {
	results := [][]int{}

	sort.Ints(nums) // asc order

	for i := range len(nums) - 2 {
		if i > 0 && nums[i] == nums[i-1] {
			continue // skip duplicate anchor
		}

		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum == 0 {
				results = append(results, []int{
					nums[i], nums[left], nums[right],
				})

				left++
				right--
				// dedup
				for left < right && nums[left] == nums[left-1] {
					left++ // skip
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum < 0 {
				// since nums is ordered in ascending, increment left by 1 will increase the sum
				left++
			} else {
				// decrement right by 1 will decrease the sum
				right--
			}
		}
	}

	return results
}
