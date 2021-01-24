package main

//给定一个只包含正整数的非空数组。是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
//
//注意:
//
//每个数组中的元素不会超过 100
//数组的大小不会超过 200

func canPartition(nums []int) bool {
	var sum = 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	var dp = make([]int, 20001)
	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
		}
	}
	if target == dp[target] {
		return true
	}
	return false
}
func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
func main() {

}
