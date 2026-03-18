package dailytemp

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	results := make([]int, 0, n)
	for i := range n {
		j := i + 1

		for j < n {
			if temperatures[i] < temperatures[j] {
				results = append(results, j-i)
				break
			}
			j++
		}

		if j == n {
			results = append(results, 0)
		}
	}

	return results
}

// using monotonic decreasing stack: always bigger numbers in the bottom.
func dailyTemperaturesWithStack(temperatures []int) []int {
	n := len(temperatures)
	result := make([]int, n) // zero for default
	stack := []int{}

	for i := range n {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temperatures[i] {
			// found
			prev := stack[len(stack)-1]
			stack = stack[:len(stack)-1] // pop

			result[prev] = i - prev
		}

		stack = append(stack, i)
	}

	return result
}
