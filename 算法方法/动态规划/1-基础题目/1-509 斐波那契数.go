package main

import "fmt"

//LeetCode 509. 斐波那契数

//斐波那契数，通常用 F(n) 表示，形成的序列称为 斐波那契数列 。该数列由 0 和 1 开始，
// 后面的每一项数字都是前面两项数字的和。也就是： F(0) = 0，F(1) = 1 F(n) = F(n - 1) + F(n - 2)，
// 其中 n > 1 给你n ，请计算 F(n) 。

// 卡尔

// 方法1：递归
//时间复杂度：O(2^n)
//空间复杂度：O(n)，算上了编程语言中实现递归的系统栈所占空间
func isFibonacciSequence(num int) int {
	if num == 1 {
		return 1
	} else if 0 == num {
		return 0
	}

	return isFibonacciSequence(num-1) + isFibonacciSequence(num-2)
}

// 方法2：动态规划-严格意义的
//时间复杂度：O(n)。
//空间复杂度：O(n)。
func fib21(n int) int {
	if n <= 1 {
		return n
	}

	// 确定dp数组（dp table）以及下标的含义+初始化
	var i int
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1

	for i = 2; i <= n; i++ {
		//确定递推公式
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// 方法2：动态规划-状态压缩,
// 只需要维护两个数值就可以了，不需要记录整个序列。
//时间复杂度：O(n)。
//空间复杂度：O(1)。 // 这里节省空间复杂度
func fib22(n int) int {
	if n <= 1 {
		return n
	}

	sum := 0
	// 确定dp数组（dp table）以及下标的含义+初始化
	dp := make([]int, 2)
	dp[0] = 0
	dp[1] = 1

	for i := 2; i <= n; i++ {
		//确定递推公式
		sum = dp[0] + dp[1]
		dp[0] = dp[1]
		dp[1] = sum
	}
	return dp[1]
}

func main() {
	for i := 0; i < 10; i++ {
		//fmt.Println("第", i+1, "项的值为", isFibonacciSequence(i))
		fmt.Println("第", i+1, "项的值为", fib21(i))
		//fmt.Println("第", i+1, "项的值为", fib22(i))
		//fmt.Println("第", i+1, "项的值为", fib2(i))
	}
}

// 方法3：动态规划-官方
// https://leetcode.cn/problems/fibonacci-number/solution/fei-bo-na-qi-shu-by-leetcode-solution-o4ze/
func fib2(n int) int {
	if n < 2 {
		return n
	}
	p, q, r := 0, 0, 1
	for i := 2; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}
