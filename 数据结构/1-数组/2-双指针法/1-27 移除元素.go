package main

import "fmt"

//LeetCode 27. 移除元素
//给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，
//并返回移除后数组的新长度。

//不要使用额外的**数组空间**，你必须仅使用 O(1) 额外空间并原地修改输入数组。
//元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。

//示例 1: 给定 nums = [3,2,2,3], val = 3, 函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。

//示例 2: 给定 nums = [0,1,2,2,3,0,4,2], val = 2, 函数应该返回新的长度 5, 并且 nums
//中的前五个元素为 0, 1, 3, 0, 4。

// 方法1：暴力解法-redo
//这个题目暴力的解法就是两层for循环，一个for循环遍历数组元素 ，第二个for循环更新数组。
// 时间复杂度：O(n^2)
// 空间复杂度：O(1)
func f1(arr []int, target int) int {
	end := len(arr)
	for i := 0; i < end; {
		if arr[i] == target { // 注意注意：如果找到了目标数后，i不能++，因为替换后的数可能还是目标数
			for j := i + 1; j < end; j++ {
				arr[j-1] = arr[j]
			}
			//fmt.Println(arr)
			end--
		} else {
			i++
		}
	}

	return end
}

// 优化后的方法：双指针法-官方-推荐
//时间复杂度：O(n)
//空间复杂度：O(1)
func f2(nums []int, val int) int {
	left := 0
	for _, v := range nums { // v 即 nums[right],第二个指针
		if v != val {
			nums[left] = v
			left++
		}
	}

	return left
}

// 双指针优化-官方
//如果要移除的元素恰好在数组的开头，例如序列 [1,2,3,4,5][1,2,3,4,5]，当 val 为 1 时，我们需要把每一个元素都左移一位。
//注意到题目中说：「元素的顺序可以改变」。实际上我们可以直接将最后一个元素 55 移动到序列开头，
//取代元素 1，得到序列 [5,2,3,4][5,2,3,4]，同样满足题目要求。这个优化在序列中 val 元素的数量较少时非常有效。

//实现方面，我们依然使用双指针，两个指针初始时分别位于数组的首尾，向中间移动遍历该序列。
//时间复杂度： O(n)，其中 n 为序列的长度。我们只需要遍历该序列至多一次。
//空间复杂度： O(1)。我们只需要常数的空间保存若干变量。
func f3(nums []int, val int) int {
	left, right := 0, len(nums)
	for left < right {
		if nums[left] == val {
			nums[left] = nums[right-1]
			right-- // 替换后只把尾部往前挪，但是首部不动，因为挪过去的数可能也是要删除的
		} else {
			left++
		}
	}

	return left
}

func main() {
	target := 3
	arr := []int{3, 2, 3}
	fmt.Println(f1(arr, target))
	fmt.Println(f2(arr, target))
	fmt.Println(f3(arr, target))
}
