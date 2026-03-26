package maxrect

func maximalRectangle(matrix [][]byte) int {
	cols := len(matrix[0])
	heights := make([]int, cols)
	maxArea := 0

	// make histogram
	for _, row := range matrix {
		for j := range cols {
			if row[j] == '1' {
				heights[j]++
			} else {
				heights[j] = 0
			}
		}

		maxArea = max(maxArea, largestRectArea(heights))
	}

	return maxArea
}

func largestRectArea(heights []int) int {
	stack := []int{}
	maxArea := 0

	// append 0 to flush stack
	heights = append(heights, 0)
	for i := range len(heights) {
		for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]] {
			// pop then calc area
			h := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]

			w := i // all items except i
			if len(stack) > 0 {
				w = i - stack[len(stack)-1] - 1
			}

			maxArea = max(maxArea, w*h)
		}

		stack = append(stack, i)
	}
	return maxArea
}
