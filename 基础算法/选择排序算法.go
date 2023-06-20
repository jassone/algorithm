package main

//算法描述：从未排序数据中选择最大或者最小的值和当前值交换 O(n^2).

//算法步骤
//选择一个数当最小值或者最大值，进行比较然后交换
//循环向后查进行

import "fmt"

//获取切片里面的最大值
func SelectMax(arr []int) int{
	len := len(arr)
	max := arr[0]

	for i := 1;i < len ;i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}

	return max
}

//切片排序
func SelectSort(arr []int) []int {
	length := len(arr)
	for i:=0;i<length;i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}

		if i != min {
			arr[min],arr[i]  = arr[i],arr[min]
		}
	}

	return arr
}

//选择排序
func main() {
	arr := []int{3, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	max := SelectMax(arr)
	selectsort := SelectSort(arr)
	fmt.Println(max)
	fmt.Println(selectsort)
}