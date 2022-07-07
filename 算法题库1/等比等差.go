package main

import "fmt"

func dengbi(start, num, step int) int {
	// 找出截止条件
	if num == 1  {
		return start
	}

	// 找出下一个循环的返回值
	next := dengbi(start,num - 1 ,step)

	// 当前的返回值
	return  next * step
}

func dengcha(start, num, step int) int {
	if num == 1  {
		return start
	}

	return  dengcha(start,num - 1 ,step) + step
}

func main() {
	fmt.Println(dengbi(1,4,2))
	fmt.Println(dengcha(1,4,2))

}
