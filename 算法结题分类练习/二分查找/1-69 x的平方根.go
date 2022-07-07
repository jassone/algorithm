package main

import "fmt"

//LeetCode 69. x 的平方根

//给你一个非负整数 x ，计算并返回 x 的 算术平方根 。
//由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。

//注意：不允许使用任何内置指数函数和算符，例如 pow(x, 0.5) 或者 x ** 0.5 。

//示例：
//输入：x = 4
//输出：2

// 方法1： 二分查找-官方
// 时间复杂度：O(logx)，即为二分查找需要的次数。
// 空间复杂度：O(1)。
func f1(x int) int {
	l, r := 0, x
	ans := -1
	for l <= r {
		mid := (l + r) / 2
		if mid*mid <= x {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return ans
}

func main() {
	num := 9
	fmt.Println(f1(num))
}
