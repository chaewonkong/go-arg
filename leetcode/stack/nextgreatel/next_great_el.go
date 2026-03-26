package nextgreatel

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	result := make([]int, 0, len(nums1))
	elMap := make(map[int]int)

	stack := []int{}

	for i := range len(nums2) {
		for len(stack) > 0 && stack[len(stack)-1] < nums2[i] {
			prev := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			elMap[prev] = nums2[i]
		}

		elMap[nums2[i]] = -1
		stack = append(stack, nums2[i])
	}

	for _, n1 := range nums1 {
		result = append(result, elMap[n1])
	}

	return result
}
