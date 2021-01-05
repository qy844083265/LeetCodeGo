package main

import "sort"

func findContentChildren(g []int, s []int) int {
	var result = 0
	sort.Slice(g, func(x, y int) bool {
		return x < y
	})
	sort.Slice(s, func(x, y int) bool {
		return x < y
	})
	var index = len(s) - 1
	for i := len(g) - 1; i >= 0; i-- {
		if index >= 0 && s[index] >= g[i] {
			result++
			index--
		}
	}
	return result
}

func main() {

}
