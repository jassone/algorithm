package main

import "fmt"

//LeetCode 474.一和零

//给你一个二进制字符串数组 strs 和两个整数 m 和 n 。

//请你找出并返回 strs 的最大子集的大小，该子集中 最多 有 m 个 0 和 n 个 1 。
// ----这里指子集中所有0最多有m个。

//如果 x 的所有元素也是 y 的元素，集合 x 是集合 y 的 子集 。

//示例 1：
//输入：strs = ["10", "0001", "111001", "1", "0"], m = 5, n = 3 输出：4
//解释：最多有 5 个 0 和 3 个 1 的最大子集是 {"10","0001","1","0"} ，因此答案是 4 。
// 其他满足题意但较小的子集包括 {"0001","1"} 和 {"10","1","0"} 。
// {"111001"} 不满足题意，因为它含 4 个 1 ，大于 n 的值 3 。

//示例 2： 输入：strs = ["10", "0", "1"], m = 1, n = 1 输出：2
// 解释：最大的子集是 {"0", "1"} ，所以答案是 2 。

//提示：
//1 <= strs.length <= 600
//1 <= strs[i].length <= 100
//strs[i] 仅由 '0' 和 '1' 组成
//1 <= m, n <= 100

// 这就是一个典型的01背包！,物品的重量有了两个维度而已。

// 实际上可以搞一个

// 卡尔-二维数组
func findMaxForm1(strs []string, m int, n int) int {
	// 定义数组
	dp := make([][]int, m+1)
	for i, _ := range dp {
		dp[i] = make([]int, n+1)
	}

	// 遍历
	for i := 0; i < len(strs); i++ {
		zeroNum, oneNum := 0, 0
		//计算0,1 个数
		//或者直接strings.Count(strs[i],"0")
		for _, v := range strs[i] {
			if v == '0' {
				zeroNum++
			}
		}
		oneNum = len(strs[i]) - zeroNum

		// 从后往前 遍历背包容量
		for j := m; j >= zeroNum; j-- {
			for k := n; k >= oneNum; k-- {
				// 推导公式
				dp[j][k] = max(dp[j][k], dp[j-zeroNum][k-oneNum]+1)
			}
		}
	}
	//fmt.Println(dp)
	//[[0 1 1 1]
	// [1 2 2 2]
	// [1 2 3 3]
	// [1 2 3 3]]
	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 传统背包-三维数组
func findMaxForm2(strs []string, m int, n int) int {
	//dp的第一个index代表项目的多少，第二个代表的是背包的容量
	//所以本处项目的多少是len（strs），容量为m和n
	dp := make([][][]int, len(strs)+1)
	for i := 0; i <= len(strs); i++ {
		//初始化背包容量
		strDp := make([][]int, m+1)
		for j := 0; j < m+1; j++ {
			tmp := make([]int, n+1)
			strDp[j] = tmp
		}
		dp[i] = strDp
	}
	for k, value := range strs {
		//统计每个字符串0和1的个数
		var zero, one int
		for _, v := range value {
			if v == '0' {
				zero++
			} else {
				one++
			}
		}
		k += 1
		//计算dp
		for i := 0; i <= m; i++ {
			for j := 0; j <= n; j++ {
				//如果装不下
				dp[k][i][j] = dp[k-1][i][j]
				//如果装的下
				if i >= zero && j >= one {
					dp[k][i][j] = getMax(dp[k-1][i][j], dp[k-1][i-zero][j-one]+1)
				}
			}
		}
	}
	return dp[len(strs)][m][n]
}
func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	arr := []string{"10", "0001", "111001", "1", "0"}
	m, n := 3, 3
	fmt.Println(findMaxForm1(arr, m, n))
	fmt.Println(findMaxForm2(arr, m, n))
}
