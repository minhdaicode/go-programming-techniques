package array

// Time complexity: O(n)
// Space complexity: O(1)
func MaxProfitII(prices []int) int {
	profit := 0
	prev := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] > prev {
			profit += prices[i] - prev
		}
		prev = prices[i]
	}
	return profit
}
