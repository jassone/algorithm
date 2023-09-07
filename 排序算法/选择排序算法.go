package main

//算法描述：从未排序数据中选择最大或者最小的值和当前值交换 O(n^2).

//算法步骤
//选择一个数当最小值或者最大值，进行比较然后交换
//循环向后查进行

import "fmt"

// 切片排序
func SelectSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length-1; i++ {
		// 假设最小值是无序区的第一个位置
		min_loc := i
		// 找到最小的下标值， 自己不用跟自己比较
		for j := i + 1; j < len(arr); j++ {
			if arr[min_loc] > arr[j] {
				min_loc = j
			}
		}
		if i != min_loc {
			arr[min_loc], arr[i] = arr[i], arr[min_loc]
		}
	}

	return arr
}

// 选择排序
func main() {
	arr := []int{3, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	selectsort := SelectSort(arr)
	fmt.Println(selectsort)
}
