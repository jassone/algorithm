package main

import "fmt"

// 约瑟夫环问题 todo
// 编号为 1-N 的 N 个士兵围坐在一起形成一个圆圈，从编号为 1 的士兵开始依次
// 报数（1，2，3…这样依次报），数到 k 的 士兵会被杀掉出列，之后的士兵再从 1 开始报数。
// 直到最后剩下一士兵，求这个士兵的编号。

func josephus(n int, k int) int {
	// 截止条件
	if n == 1 {
		return n
	}

	// 下一轮


	//


	return (josephus(n-1, k)+ k-1)%n + 1
}

func main() {
	n := 3
	k := 2
	res := josephus(n, k)
	fmt.Println(res)
}


