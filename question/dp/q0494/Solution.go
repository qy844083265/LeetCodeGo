package main

/**
给定一个非负整数数组，a1, a2, ..., an, 和一个目标数，S。现在你有两个符号 + 和 -。对于数组中的任意一个整数，你都可以从 + 或 -中选择一个符号添加在前面。

返回可以使最终数组和为目标数 S 的所有添加符号的方法数。
*/
func findTargetSumWays(nums []int, S int) int {
	var sum = 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum < S {
		return 0
	}
	if (S+sum)%2 == 1 {
		return 0
	}
	target := (S + sum) / 2
	var dp = make([]int, target+1)
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}
	return dp[target]
}
func main() {

}
