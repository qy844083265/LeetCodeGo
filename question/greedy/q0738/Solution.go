package main

import (
	"fmt"
	"strconv"
)

/**
给定一个非负整数 N，找出小于或等于 N 的最大的整数，同时这个整数需要满足其各个位数上的数字是单调递增。
*/
func monotoneIncreasingDigits(N int) int {
	s := strconv.Itoa(N)
	ss := []byte(s)
	n := len(ss)
	var flag = n
	if n <= 1 {
		return N
	}
	for i := n - 2; i >= 0; i-- {
		if ss[i] > ss[i+1] {
			ss[i] -= 1
			flag = i + 1
		}
	}
	for j := flag; j < n; j++ {
		ss[j] = '9'
	}
	str := string(ss)
	res, _ := strconv.Atoi(str)
	return res
}
func main() {
	fmt.Println(monotoneIncreasingDigits(330))
}
