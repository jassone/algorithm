package main

import "fmt"

// 卡尔
//https://programmercarl.com/背包理论基础01背包-1.html

func test_2_wei_bag_problem1(weight, value []int, bagweight int) int {
	// 定义dp数组
	dp := make([][]int, len(weight))
	for i, _ := range dp {
		dp[i] = make([]int, bagweight+1)
	}

	// 初始化
	for j := bagweight; j >= weight[0]; j-- {
		dp[0][j] = dp[0][j-weight[0]] + value[0]
	}
	// [[0 15 15 15 15] //这行直接手动赋值也行,但是如果放不下时，只能初始化为0
	// [0 0 0 0 0]
	// [0 0 0 0 0]]

	// 递推公式-先物品，后背包,weight数组的大小 就是物品个数
	// 这里不需要额外的if判断，因为j的开头已经把不合理的下标过滤了
	// 不放：dp[i-1][j]
	// 放：dp[i-1][j-weight[i]]+value[i]
	for i := 1; i < len(weight); i++ {
		//正序
		for j := weight[i]; j <= bagweight; j++ {
			dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+value[i])
		}

		// 也可以倒序
		//for j := bagweight; j >= weight[i]; j-- {
		//	dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+value[i])
		//}
	}
	// [[0 15 15 15 15]
	// [0 0 0 20 35]
	// [0 0 0 0 35]]

	// 递推公式-先背包，后物品，这里因为j从0开始，所以可能放不下某个物品，所有要加if筛选下
	//for j := 0; j <= bagweight; j++ {
	//	for i := 1; i < len(weight); i++ {
	//		if j < weight[i] {
	//			dp[i][j] = dp[i-1][j] // 等于上一个
	//		} else {
	//			dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+value[i])
	//		}
	//	}
	//}
	// [[0 15 15 15 15]
	// [0 15 15 20 35]
	// [0 15 15 20 35]]

	return dp[len(weight)-1][bagweight]
}

func main() {
	weight := []int{1, 3, 4}
	value := []int{15, 20, 30}
	result := test_2_wei_bag_problem1(weight, value, 4)
	fmt.Println(result)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
