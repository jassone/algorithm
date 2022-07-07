package main

import (
	"fmt"
	"sort"
)

//LeetCode 34 在排序数组中查找元素的第一个和最后一个位置

//给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置
// 和结束位置。
//如果数组中不存在目标值 target，返回 [-1, -1]。

//进阶：你可以设计并实现时间复杂度为 $O(logn)$ 的算法解决此问题吗？

//示例 1：
//输入：nums = [5,7,7,8,8,10], target = 8
//输出：[3,4]

//示例 2：
//输入：nums = [5,7,7,8,8,10], target = 6
//输出：[-1,-1]

//分三种情况考虑
//情况一：target 在数组范围的右边或者左边，例如数组{3, 4, 5}，target为2或者数组{3, 4, 5},
// target为6，此时应该返回{-1, -1}
//情况二：target 在数组范围中，且数组中不存在target，例如数组{3,6,7},target为5，此时应该返回{-1, -1}
//情况三：target 在数组范围中，且数组中存在target，例如数组{3,6,7},target为6，此时应该返回{1, 1}

// 思路
// 两个二分方法，一个找最左，一个找最右

// 方法1：二分法
func f1(arr []int, target int) []int {
	leftMix := getLeft(arr, target)
	rightMax := getRight(arr, target)
	//fmt.Println(leftMix,rightMax)
	if leftMix != -1 && rightMax != -1 { // 在数组中被找到
		return []int{leftMix, rightMax}
	}

	return []int{-1, -1}

}
func getLeft(arr []int, target int) int {
	length := len(arr)
	left := 0
	right := length - 1
	res := -1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] > target {
			right = mid - 1
		} else if arr[mid] < target {
			left = mid + 1
		} else if arr[mid] == target {
			right = mid - 1
			res = mid
		}
	}

	return res
}
func getRight(arr []int, target int) int {
	length := len(arr)
	left := 0
	right := length - 1
	res := -1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] > target {
			right = mid - 1
		} else if arr[mid] < target {
			left = mid + 1
		} else if arr[mid] == target {
			left = mid + 1
			res = mid
		}
	}

	return res
}

// 方法2：二分法-官方
//时间复杂度： O(logn) ，其中 n 为数组的长度。二分查找的时间复杂度为 O(logn)，一共会执行两次，
// 因此总时间复杂度为 O(logn)。
//空间复杂度：O(1) 。只需要常数空间存放若干变量。
func searchRange(nums []int, target int) []int {
	leftmost := sort.SearchInts(nums, target)
	if leftmost == len(nums) || nums[leftmost] != target {
		return []int{-1, -1}
	}
	// 走到这里，说明目标数字在数组中
	rightmost := sort.SearchInts(nums, target+1) - 1
	return []int{leftmost, rightmost}
}

func main() {
	target := 7
	arr := []int{1, 3, 4, 5, 7, 7, 7, 9}
	fmt.Println(f1(arr, target))
	fmt.Println(searchRange(arr, target))
}
