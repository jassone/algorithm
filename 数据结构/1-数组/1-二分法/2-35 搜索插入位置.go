package main

import "fmt"

//LeetCode 35. 搜索插入位置
//给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，
//返回它将会被按顺序插入的位置。
//请必须使用时间复杂度为 O(log n) 的算法。

//提示:
// nums 为 无重复元素 的 升序 排列数组

// 方法1:暴力法，分为四种情况
//目标值在数组所有元素之前
//目标值等于数组中某一个元素
//目标值插入数组中的位置
//目标值在数组所有元素之后
//时间复杂度：O(n)  空间复杂度：O(1)
func f1(arr []int, target int) int {
	length := len(arr)

	// 先把简单的几种可能性判断下
	if target < arr[0] {
		return 0
	}
	if target > arr[length-1] {
		return length
	}

	for i := 0; i < length; i++ {
		// 目标值插入数组中的位置
		if arr[i] >= target { // 一旦发现大于或者等于target的num[i]，那么i就是我们要的结果
			return i
		}
	}

	return -1
}

// 方法2 ：二分法-卡尔
func f2(arr []int, target int) int {
	length := len(arr)

	// 先简单判断下在首尾之外的情况
	if target < arr[0] {
		return 0
	}
	if target > arr[length-1] {
		return length
	}

	left := 0
	right := length - 1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] < target {
			left = mid + 1
		} else if arr[mid] > target {
			right = mid - 1
		} else {
			return mid
		}
	}

	//fmt.Println(left,right) // 4,3
	// 这时候 left = right + 1, 说明right比target小，right后面的值比target大，
	// 所以要插入到right后面的位置，即left
	return left
}

// 方法3：二分法-官方-推荐
// 先给结果设置一个值，后面再不断调整，直至结果符合要求
//时间复杂度：O(logn)，其中 n 为数组的长度。二分查找所需的时间复杂度为 O(logn)。
//空间复杂度：O(1)。我们只需要常数空间存放若干变量。
func searchInsert(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	ans := n // 这里已经将ans设置为最大值了，所以后面的判断就是check能不能把ans减少
	for left <= right {
		mid := (right + left) / 2
		if target <= nums[mid] { // 如果在数组左边，则永远走到这里，ans=0
			ans = mid
			right = mid - 1
		} else { // 如果在最右边，则永远走到这里，ans=n； 这里只能让ans越来越大，所以不用处理
			left = mid + 1
		}
	}
	return ans
}

func main() {
	target := 9
	arr := []int{1, 2, 4, 6, 8}
	fmt.Println(f1(arr, target))
	fmt.Println(f2(arr, target))
	fmt.Println(searchInsert(arr, target))
}
