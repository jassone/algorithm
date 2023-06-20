package main

import "fmt"

//LeetCode 343. 整数拆分

//给定一个正整数 n，将其拆分为至少两个正整数的和，并使这些整数的乘积最大化。 返回你可以获得的最大乘积。

//示例 1:
//输入: 2
//输出: 1
//解释: 2 = 1 + 1, 1 × 1 = 1。

//提示：2 <= n <= 58

// 方法1：动态规划-卡尔
func integerBreak2(n int) int {
	/**
	  动态五部曲
	  1.确定dp下标及其含义
	  2.确定递推公式
	  3.确定dp初始化
	  4.确定遍历顺序
	  5.打印dp
	  **/
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 1
	for i := 3; i < n+1; i++ {
		//fmt.Println("i=", i)
		curMax := 0
		for j := 1; j < i-1; j++ {
			//fmt.Println("j=", j)
			// i可以差分为i-j和j。由于需要最大值，故需要通过j遍历所有存在的值，
			// 取其中最大的值作为当前i的最大值，在求最大值的时候，一个是j与i-j相乘，
			// 一个是j与dp[i-j].
			//fmt.Println("dp[i]", dp[i])// 初始化dp[i] = 0
			curMax = max2(curMax, max2(j*(i-j), j*dp[i-j]))
			//fmt.Println("dp[i]", dp[i])
		}
		dp[i] = curMax
	}
	return dp[n]
}
func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	num := 4
	fmt.Println(integerBreak2(num))
}

// 方法1：动态规划-官方
//时间复杂度：O(n^2)
//空间复杂度：O(n)
func integerBreak1(n int) int {
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		curMax := 0
		for j := 1; j < i; j++ {
			curMax = max1(curMax, max1(j*(i-j), j*dp[i-j]))
		}
		dp[i] = curMax
	}
	return dp[n]
}
func max1(x, y int) int {
	if x > y {
		return x
	}
	return y
}
