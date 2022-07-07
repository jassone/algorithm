package main

import "fmt"

//LeetCode 1049. 最后一块石头的重量 II

//有一堆石头，每块石头的重量都是正整数。
//每一回合，从中选出任意两块石头，然后将它们一起粉碎。假设石头的重量分别为 x 和 y，且 x <= y。那么粉碎的
// 可能结果如下：

//如果 x == y，那么两块石头都会被完全粉碎； 如果 x != y，那么重量为 x 的石头将会完全粉碎，而重量为 y
// 的石头新重量为 y-x。 最后，最多只会剩下一块石头。返回此石头最小的可能重量。如果没有石头剩下，就返回 0。

//示例： 输入：[2,7,4,1,8,1] 输出：1 解释： 组合 2 和 4，得到 2，所以数组转化为 [2,7,1,8,1]，
// 组合 7 和 8，得到 1，所以数组转化为 [2,1,1,1]， 组合 2 和 1，得到 1，所以数组转化为 [1,1,1]，

// 剩下0快，两边都是half
// 剩下1快，左边half+ ， 右边half-。
// 背包问题三种情况：>=, =, <=,本题是<=
// 目标是尽量装满half就行

//提示：
//1 <= stones.length <= 30
//1 <= stones[i] <= 1000

func lastStoneWeightII(stones []int) int {
	if len(stones) == 1 {
		return stones[0]
	}

	// 15001 = 30 * 1000 /2 +1
	dp := make([]int, 15001)
	// 求target
	sum := 0
	for _, v := range stones {
		sum += v
	}
	target := sum / 2

	// 遍历顺序
	for i := 0; i < len(stones); i++ {
		for j := target; j >= stones[i]; j-- {
			// 推导公式
			dp[j] = max(dp[j], dp[j-stones[i]]+stones[i])
		}
	}
	return sum - 2*dp[target]
}

func main() {
	stones := []int{2, 7, 4, 1, 8, 1}
	//fmt.Println(lastStoneWeightII(stones))
	fmt.Println(lastStoneWeightII22(stones))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 二维数组
func lastStoneWeightII22(stones []int) int {
	n := len(stones)
	if n == 1 {
		return stones[0]
	}

	sum := 0
	for _, v := range stones {
		sum += v
	}
	target := sum / 2
	rongliang := target + 1

	// 15001 = 30 * 1000 /2 +1
	// 定义dp
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		temp := make([]int, rongliang)
		dp[i] = temp
	}

	// 初始化，设置为第一个石头质量？？？
	for i := stones[0]; i < rongliang; i++ {
		dp[0][i] = stones[0]
	}

	for i := 1; i < n; i++ {
		for j := 0; j < rongliang; j++ {
			if stones[i] > j {
				dp[i][j] = dp[i-1][j]
			} else if stones[i] <= j {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-stones[i]]+stones[i])
			}
		}
	}

	// 结果，总-left右 = left左，  left左-left右=结果
	return (sum - dp[rongliang-1][rongliang-1]) - dp[rongliang-1][rongliang-1]
}
