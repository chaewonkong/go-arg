package array

func maxProfit(prices []int) int {
	profit := 0

	for i := range len(prices) - 1 {
		if prices[i+1] > prices[i] {
			profit += (prices[i+1] - prices[i])
		}
	}

	return profit
}
