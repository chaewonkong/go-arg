package array

func containsDuplicate(nums []int) bool {
	checkMap := make(map[int]struct{})

	for _, num := range nums {
		if _, ok := checkMap[num]; ok {
			return true
		}
		checkMap[num] = struct{}{}
	}

	return false
}
