package main

import "fmt"

//LeetCode 416. 分割等和子集

//给定一个只包含正整数的非空数组。是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。

//注意: 每个数组中的元素不会超过 100 数组的大小不会超过 200

//示例 1: 输入: [1, 5, 11, 5] 输出: true 解释: 数组可以分割成 [1, 5, 5] 和 [11].

//示例 2: 输入: [1, 2, 3, 5] 输出: false 解释: 数组不能分割成两个元素和相等的子集.

//提示：
//1 <= nums.length <= 200
//1 <= nums[i] <= 100

//本题要求集合里能否出现总和为 sum / 2 的子集。
//那么来一一对应一下本题，看看背包问题如果来解决。

//只有确定了如下四点，才能把01背包问题套到本题上来。
// 背包的体积为sum / 2
// 背包要放入的商品（集合里的元素）重量为 元素的数值，价值也为元素的数值
// 背包如果正好装满，说明找到了总和为 sum / 2 的子集。
// 背包中每一个元素是不可重复放入。
//以上分析完，我们就可以套用01背包，来解决这个问题了。

//01背包相对于本题，主要要理解，题目中物品是nums[i]，重量是nums[i]，价值也是nums[i]，背包体积是sum/2。
//看代码的话，就可以发现，基本就是按照01背包的写法来的。

//时间复杂度：O(n^2)
//空间复杂度：O(n)，虽然dp数组大小为一个常数，但是大常数，子集长度。

// 方法1：动态规划-卡尔
func canPartition1(nums []int) bool {
	/**
	  动态五部曲：
	      1.确定dp数组和下标含义
	      2.确定递推公式
	      3.dp数组初始化
	      4.dp遍历顺序
	      5.打印
	  **/
	//确定和
	var sum int
	for _, v := range nums {
		sum += v
	}
	if sum%2 != 0 { //如果和为奇数，则不可能分成两个相等的数组
		return false
	}
	sum /= 2

	//确定dp数组和下标含义
	//dp[i][j], i是下标[0,i]区间里的所有整数， 在他们中间选一些数使得这些数之和为j
	var dp [][]bool

	//初始化数组
	//****因为其他元素是从上一行或者上一行左边元素得来的，所以需要多赋值一行和一列****
	dp = make([][]bool, len(nums)+1)
	for i, _ := range dp {
		dp[i] = make([]bool, sum+1)
		dp[i][0] = true
	}
	for i := 1; i <= len(nums); i++ { // 变量4个数字
		for j := 1; j <= sum; j++ { //j是固定总量
			if j >= nums[i-1] { //如果容量够用则可放入背包
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
			} else { //如果容量不够用则不拿，维持前一个状态
				dp[i][j] = dp[i-1][j] // 不选择
			}
		}
	}

	//   [true false false false false false false false false false false false]
	// 1 [true true  false false false false false false false false false false]
	// 5 [true true  false false false true  true  false false false false false]
	//11 [true true  false false false true  true  false false false false true]
	// 5 [true true  false false false true  true  false false false true  true]
	return dp[len(nums)][sum]
}

// 分割等和子集 动态规划----用滚动数组
// 时间复杂度O(n^2) 空间复杂度O(n)
func canPartition2(nums []int) bool {
	n := len(nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}
	// 如果 nums 的总和为奇数则不可能平分成两个子集
	if sum%2 == 1 {
		return false
	}

	target := sum / 2
	dp := make([]bool, target+1)
	dp[0] = true
	for i := 0; i < n; i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] = dp[j] || dp[j-nums[i]]
		}
	}
	return dp[target]
}

func main() {
	arr := []int{1, 5, 11, 5}
	//fmt.Println(canPartition1(arr))
	fmt.Println(canPartition2(arr))
}

// 方法2：回溯算法，但是会超时
//时间复杂度：O(n*2^n)
//空间复杂度：O(n)，子集长度。
