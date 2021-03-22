package main

/**
给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。

设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。

注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
*/
func maxProfit(prices []int) int {
	length := len(prices)
	var dp = make([][]int, length)
	for k, _ := range dp {
		dp[k] = make([]int, 2)
	}
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	for i := 1; i < length; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}
	return dp[length-1][1]
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
func main() {

}
