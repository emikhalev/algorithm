// https://leetcode.com/problems/best-time-to-buy-and-sell-stock
package best_time_to_buy_and_sell_stock

func maxProfit(prices []int) int {
	minPrice := 1<<31 - 1
	maxProfit := 0

	for i := 0; i < len(prices); i++ {
		if minPrice > prices[i] {
			minPrice = prices[i]
		} else {
			profit := prices[i] - minPrice
			if maxProfit < profit {
				maxProfit = profit
			}
		}
	}

	return maxProfit
}
