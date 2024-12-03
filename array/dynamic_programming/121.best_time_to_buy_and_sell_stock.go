package array

// Time complexity: O(n)
// Space complexity: O(1)
func MaxProfit(prices []int) int {
	profit := 0
	minPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if newProfit := prices[i] - minPrice; newProfit > profit {
			profit = newProfit
		}
	}
	return profit
}
