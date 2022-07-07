package main

import "fmt"

// 直接计算并打印
func func1(){
	mmap := make(map[int]int)
	for i := 0; i <= 20; i++ {
		if i > 1 {
			mmap[i] = i * mmap[i-1]
		} else {
			mmap[i] = i
		}

		fmt.Println(i, "的阶乘是", mmap[i])
	}
}

// 先计算，再打印
func func2() {
	mmap := make(map[int]int)
	mArr := []int{}
	for i := 1; i <= 20; i++ {
		mArr = append(mArr,i)
		if i>1 {
			mmap[i] = i * mmap[i-1]
		} else {
			mmap[i] = i
		}
	}

	for _,v := range mArr  {
		fmt.Println(v, "的阶乘是", mmap[v])
	}
}

func main() {
	//func1()
	func2()
}
