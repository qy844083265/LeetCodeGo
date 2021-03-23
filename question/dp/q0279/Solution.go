package main

/**
给定正整数 n，找到若干个完全平方数（比如 1, 4, 9, 16, ...）使得它们的和等于 n。你需要让组成和的完全平方数的个数最少。

给你一个整数 n ，返回和为 n 的完全平方数的 最少数量 。
*/
func numSquares(n int) int {
	var dp = make([]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = 1111
	}
	dp[0] = 0
	for i := 1; i*i <= n; i++ {
		for j := 1; j <= n; j++ {
			if j-i*i >= 0 {
				dp[j] = min(dp[j], dp[j-i*i]+1)
			}
		}
	}
	return dp[n]
}
func min(x, y int) int {
	if x >= y {
		return y
	}
	return x
}
func main() {

}
