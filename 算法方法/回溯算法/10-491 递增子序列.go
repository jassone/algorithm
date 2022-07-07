package main

import "fmt"

//LeetCode 491.递增子序列

//给你一个整数数组 nums ，找出并返回所有该数组中不同的递增子序列，递增子序列中 至少有两个元素 。你可以按 任意顺序 返回答案。
//
//数组中可能含有重复元素，如出现两个整数相等，也可以视作递增序列的一种特殊情况。

//示例:
//输入: [4, 7, 6, 7]
//输出: [[4, 6], [4, 7], [4, 6, 7], [4, 6, 7, 7], [6, 7], [6, 7, 7], [7,7], [4,7,7]]
//说明:

//给定数组的长度不会超过15。
//数组中的整数范围是 [-100,100]。
//给定数组中可能包含重复数字，相等的数字应该被视为递增的一种情况。

// 卡尔
// 注意只要考虑同一层不能有重复使用的数字即可，枝条上可以

func findSubsequences(nums []int) [][]int {
	var subRes []int
	var res [][]int
	backTring(0, nums, subRes, &res)
	return res
}
func backTring(startIndex int, nums, subRes []int, res *[][]int) {
	// 每个节点上都收集
	if len(subRes) > 1 {
		tmp := make([]int, len(subRes))
		copy(tmp, subRes)
		*res = append(*res, tmp)
	}

	used := [201]int{} //记录本层元素使用记录，因为作用域只在当前函数内
	for i := startIndex; i < len(nums); i++ {
		//分两种情况判断：
		// 一，当前取的元素小于子集的最后一个元素，则继续寻找下一个适合的元素
		// 二，当前取的元素在本层已经出现过了，所以跳过该元素，继续寻找
		if len(subRes) > 0 && nums[i] < subRes[len(subRes)-1] || used[nums[i]+100] == 1 {
			continue
		}

		used[nums[i]+100] = 1 //表示本层该元素使用过了
		subRes = append(subRes, nums[i])
		backTring(i+1, nums, subRes, res)
		subRes = subRes[:len(subRes)-1]
	}
}

func main() {
	arr := []int{4, 7, 6, 7}
	fmt.Println(findSubsequences(arr))
}

// 官方-太复杂
//https://leetcode.cn/problems/increasing-subsequences/solution/di-zeng-zi-xu-lie-by-leetcode-solution/
