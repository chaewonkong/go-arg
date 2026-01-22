package array

func missingNumber(nums []int) int {
	length := len(nums)
	desiredSum := (length + 1) * length / 2
	actualSum := 0
	for _, num := range nums {
		actualSum += num
	}

	return desiredSum - actualSum
}
