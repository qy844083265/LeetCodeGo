package main

/**
给你一个二进制字符串数组 strs 和两个整数 m 和 n 。

请你找出并返回 strs 的最大子集的大小，该子集中 最多 有 m 个 0 和 n 个 1 。

如果 x 的所有元素也是 y 的元素，集合 x 是集合 y 的 子集 。
*/
func findMaxForm(strs []string, m int, n int) int {
	var dp = make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}
	for _, str := range strs {
		zeroNum := 0
		oneNum := 0
		for _, i2 := range str {
			if i2 == '0' {
				zeroNum++
			} else {
				oneNum++
			}
		}
		for i := m; i >= zeroNum; i-- {
			for j := n; j >= oneNum; j-- {
				dp[i][j] = max(dp[i][j], dp[i-zeroNum][j-oneNum]+1)
			}
		}
	}
	return dp[m][n]
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
func main() {

}
