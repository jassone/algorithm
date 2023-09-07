package main

import (
	"fmt"
)

// 推荐
func InsertionSort(arr []int) {
	for i, v := range arr {
		preIndex := i - 1
		for preIndex >= 0 && arr[preIndex] > v {
			arr[preIndex+1] = arr[preIndex]
			preIndex -= 1
		}
		arr[preIndex+1] = v
	}
}

func main() {
	arr := make([]int, 0)
	arr = append(arr, 8, 5, 7, 2)
	//fmt.Println("初始数组:", arr)
	//InsertSort(arr)

	InsertionSort(arr)
	fmt.Println(arr)
}

func InsertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		insertVal := arr[i]
		insertIndex := i - 1 //下标
		for insertIndex >= 0 && arr[insertIndex] < insertVal {
			arr[insertIndex+1] = arr[insertIndex]
			insertIndex--
		}
		//插入
		if insertIndex+1 != i {
			arr[insertIndex+1] = insertVal
		}
		fmt.Printf("第 %d 次 %v \n", i, arr)
	}
}
