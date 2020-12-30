package main

/**
字符串 S 由小写字母组成。我们要把这个字符串划分为尽可能多的片段，同一字母最多出现在一个片段中。返回一个表示每个字符串片段的长度的列表。
*/
func partitionLabels(S string) []int {
	var result = []int{}
	var arr = make([]int, 26, 26)
	// 获取每个字母的最后一个位置
	for i := 0; i < len(S); i++ {
		arr[S[i]-'a'] = i
	}
	left := 0
	right := 0
	for i := 0; i < len(S); i++ {
		right = max(right, arr[S[i]-'a'])
		if right == i {
			result = append(result, right-left+1)
			left = i + 1
		}
	}
	return result
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func main() {

}
