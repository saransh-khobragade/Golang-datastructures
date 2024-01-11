// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/
package leetcode

func maxProfit(prices []int) int {
	buy := prices[0]
	profit := 0

	for _, x := range prices {
		if x < buy {
			buy = x //if current is less than buy,getting lowest buy point
		} else {
			if profit < x-buy {
				profit = x - buy //if curr is more than buy, if sell at this point should given max profit
			}
		}
	}
	return profit
}
