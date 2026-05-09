package array

func MaxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	minPrice := prices[0] // Store the min. price along the way
	profit := 0           // Compute the profit for current iteration
	maxProfit := 0        // Store the maxProfit accross iterations

	/*
	   eg: [7,1,5,3,6,4]
	   0 -> minPrice = 7, profit = 0, maxProfit = 0
	   1 -> minPrice = 1, profit = -6 so 0, maxProfit = 0
	   2 -> minPrice = 1, profit = 4, maxProfit = 4
	   3 -> minPrice = 1, profit = 2, maxProfit = 4
	   4 -> minPrice = 1, profit = 5, maxProfit = 5
	   5 -> minPrice = 1, profit = 3, maxProfit = 5
	*/

	/*
	   speed -> o(n)
	   space -> o(4) == o(1)
	   beats 100%
	*/
	for i := 1; i < len(prices); i++ {
		profit = max(profit, prices[i]-minPrice)
		maxProfit = max(maxProfit, profit)
		minPrice = min(minPrice, prices[i])
	}
	return maxProfit
}
