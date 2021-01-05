package main

func fib(n int) int {
	if n <= 1 {
		return n
	}
	var dp = make([]int, 2)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		sum := dp[0] + dp[1]
		dp[0] = dp[1]
		dp[1] = sum
	}
	return dp[1]
}
func main() {

}
