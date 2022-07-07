package main

import "fmt"


func cal(str string){
	num , zimu, blank ,other := 0,0,0,0
	bytes := []rune(str)
	for _,v := range bytes {
		if  v >= '0' && v <= '9'{
			num++
		} else if v >= 'A' && v <= 'z'{
			zimu++
		} else if v == ' '{
			blank++
		} else {
			other++
		}
	}

	fmt.Println("数字",num)
	fmt.Println("字母",zimu)
	fmt.Println("空格",blank)
	fmt.Println("其他",other)
}

// 要求统计出里面的字母、数字、空格以及其他字符的个数。
func main() {
	str := "s34 df  8在7s"
	cal(str)
}
