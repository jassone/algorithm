package main

import "fmt"

//算法描述：是对插入算法的一种优化，利用对问题的二分化，实现递归完成快速排序 ，
//在所有算法中二分法是最常用的方式，将问题尽量的分成两种情况加以分析，
//最终以形成类似树的方式加以利用，因为在比较模型中的算法中，最快的排序时间 负载度为 O(nlgn).

//算法步骤
// 将数据根据一个值按照大小分成左右两边，左边小于此值，右边大于此值
// 将两边数据进行递归调用步骤1
// 将所有数据合并

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	firstValue := arr[0]
	current := make([]int, 0)
	low := make([]int, 0)
	heigh := make([]int, 0)
	current = append(current, firstValue)
	for i := 1; i < len(arr); i++ {
		if arr[i] < firstValue {
			low = append(low, arr[i])
		} else if arr[i] > firstValue {
			heigh = append(heigh, arr[i])
		} else {
			current = append(current, arr[i])
		}
	}

	low, heigh = QuickSort(low), QuickSort(heigh)
	result := append(append(low, current...), heigh...)
	return result
}

// 快读排序算法
func main() {
	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	fmt.Println(QuickSort(arr))
}
