package main

import (
	"fmt"
	"sort"
)

//LeetCode 473.火柴拼正方形

// https://leetcode.cn/problems/matchsticks-to-square/solution/huo-chai-pin-zheng-fang-xing-by-leetcode-szdp/

// 方法1：回溯法-官方
//时间复杂度：O(4^n)，其中 n 是火柴的数目。每根火柴都可以选择放在 4 条边上，因此时间复杂度为 O(4^n)。
//空间复杂度：O(n)。递归栈需要占用 O(n) 的空间。
func makesquare1(matchsticks []int) bool {
	totalLen := 0
	for _, l := range matchsticks {
		totalLen += l
	}
	if totalLen%4 != 0 {
		return false
	}
	sort.Sort(sort.Reverse(sort.IntSlice(matchsticks))) // 减少搜索量

	return backtrace(matchsticks, 0, totalLen)
}
func backtrace(matchsticks []int, idx, totalLen int) bool {
	edges := [4]int{}

	if idx == len(matchsticks) {
		return true //
	}

	for i := range edges {
		edges[i] += matchsticks[idx]
		if edges[i] <= totalLen/4 && backtrace(matchsticks, idx+1, totalLen) {
			return true
		}
		edges[i] -= matchsticks[idx]
	}

	return false
}

func main() {
	arr := []int{1, 1, 2, 2, 2}
	fmt.Println(makesquare1(arr))
}
