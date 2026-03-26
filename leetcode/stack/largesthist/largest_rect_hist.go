// Package largesthist
// https://leetcode.com/problems/largest-rectangle-in-histogram/
package largesthist

func largestRectangleArea(heights []int) int {
	stack := []int{} // idx
	area := 0

	for i := range len(heights) + 1 {
		h := 0 // to flush
		if i < len(heights) {
			h = heights[i]
		}

		for len(stack) > 0 && heights[stack[len(stack)-1]] > h {
			// pop
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			height := heights[top]

			// calculate area
			width := i
			if len(stack) > 0 {
				width = i - stack[len(stack)-1] - 1
			}

			area = max(area, width*height)
		}

		stack = append(stack, i)
	}

	return area
}
