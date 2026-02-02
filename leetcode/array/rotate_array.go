package array

func _rotate(nums []int, k int) {
	n := len(nums)
	k = k % n
	tmp := make([]int, 0, n)

	for i := range k {
		tmp = append(tmp, nums[n-k+i])
	}
	for i := range n - k {
		tmp = append(tmp, nums[i])
	}

	copy(nums, tmp)
}

func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n // modulo

	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func reverse(nums []int, i, j int) {
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
