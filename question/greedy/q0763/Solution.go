package main

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
