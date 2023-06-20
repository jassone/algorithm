package main

import (
	"fmt"
	"sort"
)

//LeetCode 1. 两数之和

//给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
//你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
//你可以按任意顺序返回答案。

// 第一种 ，枚举法-官方。O(n^2) 时间复杂度的本质。
func twoSum(nums []int, target int) []int {
	for i, x := range nums {
		for j := i + 1; j < len(nums); j++ {
			if x+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// 第二种做法，hash表法-官方
// 如果需要在数组中查找某个元素是否存在时，我们可以使用哈希表法，
// 可以将查找元素的时间复杂度从O(n)减低到O(1)。

// 由于使用了外部内存，内存/空间复杂度也是O(n)。，时间复杂度则降至 O(n)；

func twoSum2(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil
}

// 第三种方法，双指针法，有个缺点是返回的下标不是原先的数组下标
// 我们可以通过一个快指针和慢指针在一个for循环下完成两个for循环的工作。例如，
// 当我们需要枚举数组中的两个元素时，如果我们发现随着第一个元素的递增，第二个元素是递减的，
// 那么就可以使用双指针的方法，将枚举的时间复杂度从O(N^2)减低到O(N)。

// 优化思路
//当我们需要枚举数组中的两个元素时，如果我们发现随着第一个元素的递增，第二个元素是递减的，
//那么就可以使用双指针的方法，将枚举的时间复杂度从O(N^2)减低到O(N)。所以我们可以采用双指针法来求解。
//首先，我们先对数组进行排序，然后用left和right指针分别指向数组的左边和右边，此时sum=nums[left]+nums[right]，
//根据sum和target的大小关系，我们来移动指针。

// 如果sum>target，右指针左移，减小sum的值，即right=right-1。
// 如果sum<target，左指针右移，增大sum的值，即left=left+1。
// 如果sum=target，直接返回。
func doublePointer(arr []int, num int) []int {
	result := []int{}
	sort.Sort(sort.IntSlice(arr))
	left := 0
	right := len(arr) - 1
	for left < right {
		if arr[left]+arr[right] > num {
			right--
		} else if arr[left]+arr[right] < num {
			left++
		} else {
			result = append(result, left, right)
			break
		}
	}
	//sort.Sort(arr)
	//fmt.Println(arr)

	return result
}

func main() {
	nums := []int{1, 5, 8, 7, 6, 44, 9}
	//res := twoSum(nums,11)
	res := doublePointer(nums, 11)
	fmt.Println(res)
}
