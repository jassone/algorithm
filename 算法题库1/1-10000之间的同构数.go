package main

import (
	"fmt"
	"strconv"
)

//正整数n若是它平方数的尾部，则称n为同构数。
//例如：5的平方数是25，且5出现在25的右侧，那么5就是一个同构数。

func tonggou(num int) bool{
	newNum := num*num
	newNumLen := len(strconv.Itoa(newNum))

	len := len(strconv.Itoa(num))
	if newNumLen == len && newNum==num {
		return true
	} else {
		dishu := 1
		for i:=1;i<=len;i++ {
			dishu *= 10
		}
		if newNum % dishu == num {
			return true
		}
	}

	return false
}

func main () {
	num := 625
	if tonggou(num) {
		fmt.Printf("是")
	} else {
		fmt.Printf("否")
	}

}
