package main

import "fmt"

//获取最大公约数
// 辗转相减法是一种简便的求出两数最大公约数的方法。
//（更相减损术）辗转相减法（求最大公约数），即尼考曼彻斯法，
// 其特色是做一系列减法，从而求得最大公约数。
func getMaximumCommonDivisor(a, b int) int {
	for a !=b {
		if a > b {
			a = a-b
		} else if a < b {
			b = b-a
		}
	}
	return a
}

func main() {
	var a, b = 2, 6
	num := getMaximumCommonDivisor(a, b)
	fmt.Println("a,b的最大公约数是：", num)

	//求最小公倍数相对来说就比较简单了。只需要先求出最大公约数。
	//用两个数的乘积除以最大公约数即可
	fmt.Println("a,b的最小公倍数是：", a*b/num)
}
