package main

import "sort"

// 给定一个区间的集合，找到需要移除区间的最小数量，使剩余区间互不重叠。

// 注意:

//     可以认为区间的终点总是大于它的起点。
//     区间 [1,2] 和 [2,3] 的边界相互“接触”，但没有相互重叠。
func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var pre = 0
	var count = 1
	for i := 0; i < len(intervals); i++ {
		if intervals[i][0] >= intervals[pre][1] {
			pre = i
			count++
		} else if intervals[i][1] < intervals[pre][1] {
			pre = i
		}
	}
	return len(intervals) - count
}

func main() {

}
