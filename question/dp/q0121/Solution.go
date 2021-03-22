package main

/**
 * 给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
 * <p>
 * 你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
 * <p>
 * 返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。
 */

func maxProfit(prices []int) int {
	var length = len(prices)
	var dp = make([][]int, length)
	for k, _ := range dp {
		dp[k] = make([]int, 2)
	}
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	for i := 1; i < length; i++ {
		dp[i][0] = max(dp[i-1][0], -prices[i])
		dp[i][1] = max(dp[i-1][1], prices[i]+dp[i-1][0])
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
