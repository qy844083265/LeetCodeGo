package main

// 老师想给孩子们分发糖果，有 N 个孩子站成了一条直线，老师会根据每个孩子的表现，预先给他们评分。

// 你需要按照以下要求，帮助老师给这些孩子分发糖果：

//     每个孩子至少分配到 1 个糖果。
//     评分更高的孩子必须比他两侧的邻位孩子获得更多的糖果。

// 那么这样下来，老师至少需要准备多少颗糖果呢？

func candy(ratings []int) int {
	var length = len(ratings)
	var arr = make([]int, length, length)
	for i := 0; i < length; i++ {
		arr[i] = 1
	}
	for i := 1; i < length; i++ {
		if ratings[i] > ratings[i-1] {
			arr[i] = arr[i-1] + 1
		}
	}
	for i := length - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			arr[i] = max(arr[i], arr[i+1]+1)
		}
	}
	result := 0
	for i := 0; i < length; i++ {
		result += arr[i]
	}
	return result
}
func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
func main() {

}
