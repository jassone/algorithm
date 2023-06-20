package main

import "fmt"

//LeetCode 704. 二分查找

//算法描述：在一组有序数组中，将数组一分为二，将要查询的元素和分割点进行比较，分为三种情况
// 相等直接返回
// 元素大于分割点，在分割点右侧继续查找
// 元素小于分割点，在分割点左侧继续查找

//时间复杂： O(logn).

// 分析
// 前提要求 ： 必须是有序的数组，并能支持随机访问
// 搜索过程：****会搜索所有可能的情况，只要判断条件正确，搜索不到则真的搜索不到****

// 大家写二分法经常写乱，主要是因为对区间的定义没有想清楚，**区间的定义就是不变量。**
// 要在二分查找的过程中，**保持不变量**，就是在while寻找中每一次边界的处理都要坚持根据区间的定义来操作，
// 这就是循环不变量规则。

// 疑问点
//low 是否要等于 heigh
//mid 是否要加一或者减一

// 关键定义
//min := -1
//low := 0
//heigh := len(arr) - 1 // **
//for low <= heigh {    // ** 这里是左闭右闭的情况
//    mid := (low + heigh) / 2   // **
//    if arr[mid] < findData {
//        low = mid + 1  // 别忘了加1，因为大循环是左闭右闭，而这时肯定不会包含mid这个值，所以要+1
//    } else if arr[mid] > findData {
//        heigh = mid - 1  // 别忘了减1，理由同上
//    } else {
//        heigh = mid
//		  break;
//    }
//}

/***************版本1：左闭右闭区间*****************/
//high := len(arr) - 1 // 区间[low,high]
//for low <= high { // low==high是有意义的
//    mid := (low + high) / 2
//}

//变形：查找第一个值等于给定的
// 在相等的时候做处理，向前查
func main2() {
	arr := []int{1, 3, 3, 5, 6, 7, 8, 9, 10, 11, 13}
	id := bin_search2(arr, 3)
	fmt.Println(id)
}
func bin_search2(arr []int, findData int) int {
	min := -1
	low := 0
	heigh := len(arr) - 1
	for low <= heigh {
		mid := (low + heigh) / 2
		if arr[mid] < findData {
			low = mid + 1
		} else if arr[mid] > findData {
			heigh = mid - 1
		} else {
			heigh = mid - 1
			min = mid
		}
	}

	return min
}

//变形：查找最后一个值等于给定的值
// 在相等的时候做处理，向后查
func main3() {
	arr := []int{1, 3, 3, 5, 6, 7, 8, 9, 10, 11}
	id := bin_search3(arr, 3)
	fmt.Println(id)
}
func bin_search3(arr []int, findData int) int {
	max := -1
	low := 0
	heigh := len(arr) - 1
	for low <= heigh {
		mid := (low + heigh) / 2
		if arr[mid] < findData {
			low = mid + 1
		} else if arr[mid] > findData {
			heigh = mid - 1
		} else {
			low = mid + 1
			max = mid
		}
	}

	return max
}

//变形：查找第一个大于等于给定的值
func bin_search4(arr []int, findData int) int {
	res := -1
	low := 0
	heigh := len(arr) - 1
	for low <= heigh {
		mid := (low + heigh) / 2
		if arr[mid] >= findData {
			heigh = mid - 1
			res = mid
		} else {
			low = mid + 1
		}
	}

	return res
}

//变形：查找最后一个小于等于给定的值
func bin_search5(arr []int, findData int) int {
	res := -1
	low := 0
	heigh := len(arr) - 1
	for low <= heigh {
		mid := (low + heigh) / 2
		if arr[mid] <= findData {
			low = low + 1
			res = mid
		} else {
			heigh = mid - 1
		}
	}

	return res
}

// 变形：找到既返回-递归
func bin_search6(arr []int, target int) int {
	low := 0
	heigh := len(arr) - 1
	var mid int
	res := -1

	fmt.Println("中间数", (low+heigh)/2)
	fmt.Println(arr)

	if low <= heigh {
		mid = (low + heigh) / 2
		if arr[mid] < target {
			res = bin_search6(arr[mid+1:], target)
			if res >= 0 { // 跳转下标
				res = res + mid + 1
			}
		} else if arr[mid] > target {
			res = bin_search6(arr[:mid], target) // 特别要注意这里不能是 arr[:mid+1]
		} else {
			return mid
		}
	}

	return res
}

//实际应用
// 用户ip区间段查询
// 用于相似度查询

// 找到即返回：正常解法
func bin_search(arr []int, findData int) int {
	low := 0
	heigh := len(arr) - 1
	for low <= heigh {
		mid := (low + heigh) / 2
		if arr[mid] < findData {
			low = mid + 1
		} else if arr[mid] > findData {
			heigh = mid - 1
		} else {
			return mid
		}
	}

	return -1
}

/***************版本2：左闭右开区间*****************/
func bin_search7(nums []int, target int) int {
	high := len(nums) // 和版本1的不同点，这里不要减一
	low := 0
	for low < high { // 和版本1的不同点，low==high是没有意义的，比如[1,1)是没有意义的
		mid := (low + high) / 2
		if nums[mid] > target {
			high = mid // 因为右边是不包含的，既target 在左区间，在[left, mid)中，所以这里等于mid即可，
		} else if nums[mid] < target {
			low = mid + 1 // 因为左边是包含的，既target 在左区间，在[mid+1,right)中，所以这里要加1
		} else {
			return mid
		}
	}

	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}

	id := bin_search7(arr, 7)
	if id != -1 {
		fmt.Println(id, arr[id])
	} else {
		fmt.Println("没有找到数据")
	}
}
