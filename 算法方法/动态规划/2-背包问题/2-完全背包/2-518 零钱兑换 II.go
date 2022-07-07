package main

import "fmt"

//LeetCode 518. 零钱兑换 II

//给定不同面额的硬币和一个总金额。写出函数来计算可以凑成总金额的硬币组合数。假设每一种面额的硬币有无限个。

//示例 1:
//输入: amount = 5, coins = [1, 2, 5] 输出: 4
// 解释: 有四种方式可以凑成总金额: 5=5 5=2+2+1 5=2+1+1+1 5=1+1+1+1+1

// 方法1：动态规划
func change1(amount int, coins []int) int {
	// 定义dp数组
	dp := make([]int, amount+1)

	// 初始化,0大小的背包, 当然是不装任何东西了, 就是1种方法
	dp[0] = 1

	// 遍历顺序
	// 遍历物品
	for i := 0; i < len(coins); i++ {
		// 遍历背包
		for j := coins[i]; j <= amount; j++ {
			// 推导公式
			dp[j] += dp[j-coins[i]]
		}
	}
	return dp[amount]
}

// 二维数组-其他人
func change2(amount int, coins []int) int {
	m := len(coins)

	// 定义
	dp := make([][]int, m+1) //添加第i个硬币，总结额为j的硬币组合数
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, amount+1)
	}

	// 初始化
	for i := 0; i <= m; i++ { //当总金额为0时，组合数只能有1种，就是都不添加
		dp[i][0] = 1
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= amount; j++ {
			if j >= coins[i-1] { //总金额大于硬币值，组合数等于 添加i硬币总金额为j-coins[i-1时的组合数 + 未添加硬币总金额为j的组合数
				dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]] //这里注意，dp[i][j-cons[i-1]]是和01背包区别之处，因为完全背包数量没有限制，所以我仍然可以用第i的硬币
			} else { //总金额小于硬币值，组合数等于未添加i硬币总金额为j的组合数
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[m][amount]
}

func main() {
	amount := 5
	coins := []int{1, 2, 5}
	fmt.Println(change1(amount, coins))
	fmt.Println(change2(amount, coins))
}
