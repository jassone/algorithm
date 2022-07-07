package main

import (
	"fmt"
	"math"
)

//LeetCode 53. 最大子序和

//给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

//示例: 输入: [-2,1,-3,4,-1,2,1,-5,4] 输出: 6 解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。

// 卡尔
//方法1：暴力解法
//暴力解法的思路，第一层for 就是设置起始位置，第二层for循环遍历数组寻找最大值

//时间复杂度：O(n^2)
//空间复杂度：O(1)
func maxSubArray11(nums []int) int {
	MaxSum := nums[0]
	temMax := 0
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			temMax += nums[j]
			if temMax > MaxSum {
				MaxSum = temMax
			}
		}
		temMax = 0
	}
	return MaxSum
}

//贪心解法
//贪心贪的是哪里呢？

//如果 -2 1 在一起，计算起点的时候，一定是从1开始计算，因为负数只会拉低总和，这就是贪心贪的地方！

//局部最优：当前“连续和”为负数的时候立刻放弃，从下一个元素重新计算“连续和”，因为负数加上下一个元素 “连续和”只会越来越小。

//全局最优：选取最大“连续和”

//****局部最优的情况下，并记录最大的“连续和”，可以推出全局最优。***

//贪心的思路为局部最优：当前“连续和”为负数的时候立刻放弃，从下一个元素重新计算“连续和”，因为负数加上下一个元素
// “连续和”只会越来越小。从而推出全局最优：选取最大“连续和”

//****这相当于是暴力解法中的不断调整最大子序和区间的起始位置。***

//那有同学问了，区间终止位置不用调整么？ 如何才能得到最大“连续和”呢？

//区间的终止位置，其实就是如果count取到最大值了，及时记录下来了。例如如下代码：
//if (count > result) result = count;

//****这样相当于是用result记录最大子序和区间和（变相的算是调整了终止位置）****

//数组都为负数，maxSum记录的就是最小的负数，如果数组里有int最小值，那么最终result就是int最小值。

//时间复杂度：O(n)，其中 n 为 nums 数组的长度。我们只需要遍历一遍数组即可求得答案。
//空间复杂度：O(1)。我们只需要常数空间存放若干变量。
// 更好理解
func maxSubArray22(nums []int) int {
	curSum, maxSum := 0, math.MinInt
	for i := 0; i < len(nums); i++ {
		curSum += nums[i]
		if curSum > maxSum {
			maxSum = curSum
		}
		if curSum < 0 {
			curSum = 0
		}
	}
	return maxSum
}

// 少用一个变量,更精简
func maxSubArray33(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

func main() {
	arr := []int{-2, -1, -3, 4, -1, 2, 1, -5, 4}
	//fmt.Println(maxSubArray11(arr))
	fmt.Println(maxSubArray22(arr))
	//fmt.Println(maxSubArray33(arr))
}
