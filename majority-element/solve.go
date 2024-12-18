package max_profit

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/?envType=study-plan-v2&envId=top-interview-150

func maxProfit(prices []int) int {
	minPrice := prices[0]
	profit := 0

	for _, price := range prices {
		minPrice = min(minPrice, price)

		profit = max(profit, price-minPrice)
	}

	return profit
}
