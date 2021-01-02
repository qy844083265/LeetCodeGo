package main

// 在一条环路上有 N 个加油站，其中第 i 个加油站有汽油 gas[i] 升。
// 你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1 个加油站需要消耗汽油 cost[i] 升。你从其中的一个加油站出发，开始时油箱为空。
// 如果你可以绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1。

func canCompleteCircuit(gas []int, cost []int) int {
	var curSum = 0
	var totalSum = 0
	var start = 0
	for i := 0; i < len(gas); i++ {
		ret := gas[i] - cost[i]
		curSum += ret
		totalSum += ret
		if curSum < 0 {
			start = i + 1
			curSum = 0
		}
	}
	if totalSum < 0 {
		return -1
	}
	return start
}
func main() {

}
