package search

func search(nums []int, target int) int {
	for idx, num := range nums {
		if num == target {
			return idx
		}
	}

	return -1
}

func binSearch(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	mid := (left + right) / 2

	for left <= right {
		mid = (left + right) / 2

		if nums[mid] == target {
			return mid
		}

		if nums[left] <= nums[mid] {
			// 왼쪽정렬
			if nums[left] <= target && nums[mid] > target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// 오른쪽 정렬
			if nums[mid] < target && nums[right] >= target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}

func binSearchWithPivot(nums []int, target int) int {
	// find pivot
	n := len(nums)
	left := 0
	right := n - 1

	for left < right {
		mid := (left + right) / 2

		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	pivot := left

	// chose sub-array
	left = 0
	right = n - 1
	if target >= nums[pivot] && target <= nums[n-1] {
		// right side of pivot
		left = pivot
	} else {
		right = pivot - 1
	}

	// binary search
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}
