package main

import "fmt"

// 水仙花数是指一个 3 位数，它的每个位上的数字的
// 3次幂之和等于它本身（例如：1^3 + 5^3+ 3^3 = 153）
func sxh(num int) bool {
	n1 := num / 100
	n2 := (num % 100) / 10
	n3 := num % 10
	if (n1*n1*n1 + n2*n2*n2 + n3*n3*n3) == num {
		return true
	}
	return false
}

func main() {
	len := 1000
	for i:=100;i<len;i++ {
		res := sxh(i)
		if res {
			fmt.Println(i)
		}
	}
}
