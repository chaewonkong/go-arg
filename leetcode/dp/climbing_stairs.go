package dp

func climbStairs(n int) int {
	// f(n) = f(n-1) + f(n-2)
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}
	a := 1
	b := 2
	for i := range n {
		if i < 2 {
			continue
		}
		sum := a + b
		a = b
		b = sum
	}

	return b
}
