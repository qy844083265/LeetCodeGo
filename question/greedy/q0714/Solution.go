package main

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
