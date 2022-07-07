package main

import "fmt"

//LeetCode 367. 有效的完全平方数
//给定一个 正整数 num ，编写一个函数，如果 num 是一个完全平方数，则返回 true ，
//否则返回 false 。

//进阶：不要 使用任何内置的库函数，如  sqrt 。

// 方法1：二分法-官方
func f1(x int) int {
	l, r := 0, x
	ans := -1
	for l <= r {
		mid := (l + r) / 2
		sqrt := mid * mid
		if sqrt < x {
			l = mid + 1
		} else if sqrt > x {
			r = mid - 1
		} else {
			return mid
		}
	}

	return ans
}

func main() {
	num := 9
	fmt.Println(f1(num))
}
