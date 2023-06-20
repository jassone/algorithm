package main

import "fmt"

//LeetCode 746. 使用最小花费爬楼梯

//数组的每个下标作为一个阶梯，第 i 个阶梯对应着一个非负数的体力花费值 cost[i]（下标从 0 开始）。

//每当你爬上一个阶梯你都要花费对应的体力值，一旦支付了相应的体力值，你就可以选择向上爬一个阶梯或者爬两个阶梯。

//请你找出达到楼层顶部的最低花费。在开始时，你可以选择从下标为 0 或 1 的元素作为初始阶梯。

//示例 1：
//输入：cost = [10, 15, 20] 输出：15 解释：最低花费是从 cost[1] 开始，然后走两步即可到阶梯顶，一共花费 15 。

//示例 2：
//输入：cost = [1, 100, 1, 1, 1, 100, 1, 1, 100, 1] 输出：6 解释：最低花费方式是从 cost[0] 开始，
// 逐个经过那些 1 ，跳过 cost[3] ，一共花费 6 。

//提示：
//cost 的长度范围是 [2, 1000]。
//cost[i] 将会是一个整型数据，范围为 [0, 999] 。

// 注意点：
// 起步可以从下标0或者1开始
// 顶部不是走到了最后一步，而是最后一步再踏一步
//
//

// 卡尔

// 方法1：动态规划-官方
//时间复杂度：O(n)，其中 n 是数组 cost 的长度。需要依次计算每个 dp 值，每个值的计算需要常数时间，
// 因此总时间复杂度是 O(n)。
//空间复杂度：O(n)。
func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost))
	dp[0], dp[1] = cost[0], cost[1]
	for i := 2; i < len(cost); i++ {
		dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
	}
	// 注意最后一步可以理解为不用花费，所以取倒数第一步，第二步的最少值
	return min(dp[len(cost)-1], dp[len(cost)-2])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 官方-更好理解
func minCostClimbingStairs33(cost []int) int {
	n := len(cost)
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[n]
}

func min33(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 优化-空间复杂度-官方 todo
//时间复杂度：O(n)，其中 n 是数组 cost 的长度。需要依次计算每个 dp 值，每个值的计算需要常数时间，
// 因此总时间复杂度是 O(n)。
//空间复杂度：O(1)。使用滚动数组的思想，只需要使用有限的额外空间。
func minCostClimbingStairs44(cost []int) int {
	n := len(cost)
	pre, cur := 0, 0
	for i := 2; i <= n; i++ {
		pre, cur = cur, min44(cur+cost[i-1], pre+cost[i-2])
	}
	return cur
}
func min44(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	arr := []int{10, 15, 20}
	//fmt.Println(minCostClimbingStairs(arr))
	fmt.Println(minCostClimbingStairs33(arr))
	fmt.Println(minCostClimbingStairs44(arr))
}
