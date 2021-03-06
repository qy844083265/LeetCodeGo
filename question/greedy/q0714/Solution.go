package main

/**
给定一个整数数组 prices，其中第 i 个元素代表了第 i 天的股票价格 ；非负整数 fee 代表了交易股票的手续费用。

你可以无限次地完成交易，但是你每笔交易都需要付手续费。如果你已经购买了一个股票，在卖出它之前你就不能再继续购买股票了。

返回获得利润的最大值。

注意：这里的一笔交易指买入持有并卖出股票的整个过程，每笔交易你只需要为支付一次手续费。
*/
func maxProfit(prices []int, fee int) int {
	var result = 0
	var min = prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		}
		if prices[i] >= min && prices[i] <= min+fee {
			continue
		}
		if prices[i] > min+fee {
			result += prices[i] - min - fee
			min = prices[i] - fee
		}
	}
	return result
}
func main() {

}
