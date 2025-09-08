package main

// maxProfit returns the maximum profit achievable with unlimited transactions.
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	total := 0
	for i := 0; i+1 < len(prices); i++ {
		if prices[i+1] > prices[i] {
			total += prices[i+1] - prices[i]
		}
	}
	return total
}
