package main

import "fmt"

func total(day int) int {
	// 截止条件
	if  day == 1 {
		return 1
	}

	// 上一天
	preDay := total(day-1)

	// 当前天
	return (preDay+ 1) * 2
}

func main() {
	day := 10
	fmt.Println(total(day))
}
