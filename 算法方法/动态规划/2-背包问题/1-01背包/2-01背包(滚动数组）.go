package main

import "fmt"

// https://programmercarl.com/背包理论基础01背包-2.html

// (滚动数组）-必须先背包，后物品，从上到下
func test_1_wei_bag_problem(weight, value []int, bagWeight int) int {
	// 定义 && 初始化
	// 这里都初始化为0
	dp := make([]int, bagWeight+1)
	// 递推顺序
	for i := 0; i < len(weight); i++ {
		// 这里必须倒序,区别二维,因为二维dp保存了i的状态
		// 倒序遍历是为了保证物品i只被放入一次！
		for j := bagWeight; j >= weight[i]; j-- {
			// 递推公式
			dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
		}
	}
	//fmt.Println(dp)
	// [0 15 15 20 35]
	return dp[bagWeight]
}

func main() {
	weight := []int{1, 3, 4}
	value := []int{15, 20, 30}
	res := test_1_wei_bag_problem(weight, value, 4)
	fmt.Println(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
