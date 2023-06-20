package main

import (
	"fmt"
	"sort"
)

//LeetCode 90.子集II

//给定一个可能包含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
//说明：解集不能包含重复的子集。

//示例:
//输入: [1,2,2]
//输出: [ [2], [1], [1,2,2], [2,2], [1,2], [] ]

// 卡尔
// 同78题，这里多了个重复判断
// 注意只要考虑同一层不能有重复使用的数字即可，枝条上可以
//https://www.programmercarl.com/0090.%E5%AD%90%E9%9B%86II.html#_90-%E5%AD%90%E9%9B%86ii

var res [][]int

func subsetsWithDup(nums []int) [][]int {
	res = make([][]int, 0)
	sort.Ints(nums) // 需要先排序，为后面重复判断做准备
	dfs([]int{}, nums, 0)
	return res
}
func dfs(temp, num []int, start int) {
	tmp := make([]int, len(temp))
	copy(tmp, temp)
	res = append(res, tmp)

	for i := start; i < len(num); i++ {
		if i > start && num[i] == num[i-1] {
			continue
		}
		temp = append(temp, num[i])
		dfs(temp, num, i+1)
		temp = temp[:len(temp)-1]
	}
}

func main() {
	arr := []int{1, 2, 2}
	fmt.Println(subsetsWithDup(arr))
}

//
