package main

import (
	"fmt"
	"math"
)

//LeetCode 494. 目标和

//给定一个非负整数数组，a1, a2, ..., an, 和一个目标数，S。现在你有两个符号 + 和 -。对于数组中的任意一个整数，
// 你都可以从 + 或 -中选择一个符号添加在前面。

//返回可以使最终数组和为目标数 S 的所有添加符号的方法数。

//示例：
//输入：nums: [1, 1, 1, 1, 1], S: 3
//输出：5

//解释：
//-1+1+1+1+1 = 3
//+1-1+1+1+1 = 3
//+1+1-1+1+1 = 3
//+1+1+1-1+1 = 3
//+1+1+1+1-1 = 3
//一共有5种方法让最终目标和为3。

//提示：
//
//数组非空，且长度不会超过 20 。
//初始的数组的和不会超过 1000 。
//保证返回的最终结果能被 32 位整数存下。

// 卡尔
// https://programmercarl.com/0494.目标和.html
// 动态规划
//本题要如何使表达式结果为target，

//既然为target，那么就一定有 left组合 - right组合 = target。
//left + right等于sum，而sum是固定的。
//公式来了， left - (sum - left) = target -> left = (target + sum)/2 。
//target是固定的，sum是固定的，left就可以求出来。

//此时问题就是在集合nums中找出和为left的组合。

// 方法1：回溯-官方
func findTargetSumWays1(nums []int, target int) (count int) {
	var backtrack func(int, int)
	backtrack = func(index, sum int) {
		if index == len(nums) {
			if sum == target {
				count++
			}
			return
		}
		// 妙呀，这里每次都处理两个逻辑，一个加，一个减少
		backtrack(index+1, sum+nums[index])
		backtrack(index+1, sum-nums[index])
	}
	backtrack(0, 0)
	return
}

// 动态规划-官方(再微调下)
func findTargetSumWays2(nums []int, target int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	// 一些初步过滤
	if float64(sum) < math.Abs(float64(target)) {
		return 0
	}

	if (sum+target)%2 == 1 {
		return 0
	}

	// 初始化
	wuSize, bagSize := len(nums), (sum+target)/2
	dp := make([][]int, wuSize+1)
	for i := range dp {
		dp[i] = make([]int, bagSize+1)
	}
	dp[0][0] = 1

	for i := 0; i < wuSize; i++ {
		for j := 0; j <= bagSize; j++ {
			dp[i+1][j] = dp[i][j]
			if nums[i] <= j {
				dp[i+1][j] += dp[i][j-nums[i]]
			}
		}
	}
	//fmt.Println(dp)
	//   0 1 2
	// [[1 0 0]
	//1 [1 1 0]
	//1 [1 2 1]
	//1 [1 3 3]]

	return dp[wuSize][bagSize]
}

// 卡尔-一维数组
func findTargetSumWays3(nums []int, target int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	// 一些初步过滤
	if float64(sum) < math.Abs(float64(target)) {
		return 0
	}

	if (sum+target)%2 == 1 {
		return 0
	}

	// 计算背包大小
	bagSize := (sum + target) / 2
	// 定义dp数组
	dp := make([]int, bagSize+1)
	// 初始化
	dp[0] = 1
	// 遍历顺序
	for i := 0; i < len(nums); i++ {
		for j := bagSize; j >= nums[i]; j-- {
			//推导公式
			dp[j] += dp[j-nums[i]]
			//fmt.Println(dp)
		}
	}
	return dp[bagSize]
}

func main() {
	nums := []int{1, 1, 1}
	target := 1
	//fmt.Println(findTargetSumWays1(nums, target))
	//fmt.Println(findTargetSumWays2(nums, target))
	fmt.Println(findTargetSumWays3(nums, target))
}
