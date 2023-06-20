package main

import (
	"fmt"
	"math/big"
)

//LeetCode 62.不同路径

//一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。

//机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。

//问总共有多少条不同的路径？

//示例 2：
//输入：m = 2, n = 3
//输出：3
//解释： 从左上角开始，总共有 3 条路径可以到达右下角。

//向右 -> 向右 -> 向下
//向右 -> 向下 -> 向右
//向下 -> 向右 -> 向右

// 卡尔
// https://programmercarl.com/0062.%E4%B8%8D%E5%90%8C%E8%B7%AF%E5%BE%84.html#%E6%80%9D%E8%B7%AF

// 确定递推公式 -- 参考70 爬楼梯那题
//想要求dp[i][j]，只能有两个方向来推导出来，即dp[i - 1][j] 和 dp[i][j - 1]。

//此时在回顾一下 dp[i - 1][j] 表示啥，是从(0, 0)的位置到(i - 1, j)有几条路径，dp[i][j - 1]同理。

//那么很自然，dp[i][j] = dp[i - 1][j] + dp[i][j - 1]，因为dp[i][j]只有这两个方向过来。

//dp数组的初始化
//如何初始化呢，首先dp[i][0]一定都是1，因为从(0, 0)的位置到(i, 0)的路径只有一条，那么dp[0][j]也同理。

//所以初始化代码为：
//for (int i = 0; i < m; i++) dp[i][0] = 1;
//for (int j = 0; j < n; j++) dp[0][j] = 1;

// 方法1：深度搜索
// 注意题目中说机器人每次只能向下或者向右移动一步，那么其实机器人走过的路径可以抽象为一棵二叉树，而叶子节点就是终点！
// 但是可能会超时

// 方法2：动态规划
//时间复杂度：O(mn)。
//空间复杂度：O(mn)，即为存储所有状态需要的空间。注意到 f(i,j) 仅与第 i 行和第 i−1
// 行的状态有关，因此我们可以使用滚动数组代替代码中的二维数组，使空间复杂度降低为 O(n)。此外，
// 由于我们交换行列的值并不会对答案产生影响，因此我们总可以通过交换 m 和 n 使得 m≤n，
// 这样空间复杂度降低至 O(min(m,n))。
func uniquePaths222(m int, n int) int {
	dp := make([][]int, m)
	// 先把最外层的行列填上数
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

// 方法2-优化空间复杂度-todo
//dp[i][j] 表示节点 (i,j) 的路径数。
//每一个节点都可以由 上方 或者 左方 的节点达到，所以 dp[i][j] = dp[i-1][j] + dp[i][j-1]。
//由于每一行节点都只依赖上一行当前列以及当前行上一列的值，所以可以用一维滚动数组来处理：

//currRow[j] = currRow[j-1](当前行上一列) + currRow[j](上一行当前列)
//简化下就是：
//currRow[j] += currRow[j-1]`(当前行上一列)``

//时间复杂度：O(m × n)
//空间复杂度：O(n)
func uniquePaths2221(m int, n int) int {
	dp := make([]int, n)

	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] = dp[j-1] + dp[j]
		}
	}

	return dp[n-1]
}

// 方法3：组合的思想
//在这个图中，可以看出一共m，n的话，无论怎么走，走到终点都需要 m + n - 2 步。
// 在这m + n - 2 步中，一定有 n - 1 步是要向下走的，不用管什么时候向下走。
//那么有几种走法呢？ 可以转化为，给你m + n - 2个不同的数，随便取n - 1个数，有几种取法。
//那么这就是一个组合问题了。
func uniquePaths333(m, n int) int {
	return int(new(big.Int).Binomial(int64(m+n-2), int64(n-1)).Int64())
}

func main() {
	a, b := 3, 7
	fmt.Println(uniquePaths222(a, b))
	fmt.Println(uniquePaths2221(a, b))
	fmt.Println(uniquePaths333(a, b))
}
