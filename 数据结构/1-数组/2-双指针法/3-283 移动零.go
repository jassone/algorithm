package main

import "fmt"

//LeetCode 283 移动零
//给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

//请注意 ，必须在不复制数组的情况下原地对数组进行操作。
// 方法1：双指针-卡尔
func f1(arr []int) []int {
	slow, fast, length := 0, 0, len(arr)
	for ; fast < length; fast++ {
		if arr[fast] != 0 {
			arr[slow] = arr[fast]
			slow++
		}
	}

	return arr
}

// 方法2：双指针-官方
// left:满指针，right:快指针
//时间复杂度：O(n)，其中 n 为序列长度。每个位置至多被遍历两次。
//空间复杂度：O(1)。只需要常数的空间存放若干变量。
func moveZeroes(nums []int) []int {
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			// 其实这里只要替换一个即可
			//nums[left] = nums[right]
			left++
		}
		right++
	}

	return nums
}

func main() {
	arr := []int{1, 2, 0, 0, 4, 0, 5}
	fmt.Println(f1(arr))
	fmt.Println(moveZeroes(arr))
}
