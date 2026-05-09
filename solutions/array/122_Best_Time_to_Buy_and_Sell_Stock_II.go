package array

/*
https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii?envType=study-plan-v2&envId=top-interview-150
*/
func MaxProfitII(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	buyPrice := prices[0]
	profit := 0
	totalProfit := 0

	/*
	   speed -> o(n)
	   space -> o(3) == o(1)
	   beats 100%
	*/
	for i := 1; i < len(prices); i++ {
		profit = prices[i] - buyPrice
		if profit > 0 { // Booking profit
			totalProfit += profit
		}
		/*
		   1) Sell and Buy back after the profit booking or
		   2) Buy the stock when price is lower than previous
		*/
		buyPrice = prices[i]
	}
	return totalProfit
}
