package main

import (
	"fmt"
	"math"
)


func isZhiShu(num int) bool {
	for i:=2; float64(i) <= math.Sqrt(float64(num));i++ {
		if (num % i) == 0 {
			return false
		}
	}

	return true
}


func main() {
	var res []int
	for i:=1;i<100 ;i++  {
		if isZhiShu(i) {
			res = append(res,i)
		}
	}

	fmt.Println(res)
}
