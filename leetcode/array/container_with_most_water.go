package array

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1

	maxArea := 0

	for left < right {
		size := min(height[left], height[right]) * (right - left)
		maxArea = max(maxArea, size)

		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}
