package main

//LeetCode 96.不同的二叉搜索树

// 给定一个整数 n，求以 1 ... n 为节点组成的二叉搜索树有多少种？

// 方法1：动态规划-难度大 todo
func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}
