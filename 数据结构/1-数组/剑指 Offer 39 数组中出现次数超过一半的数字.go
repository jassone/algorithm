package main

import (
	"fmt"
	"sort"
)

//LeetCode 剑指 Offer 39. 数组中出现次数超过一半的数字

//给一个长度为 n 的数组，数组中有一个数字出现的次数超过数组长度的一半，
//请找出这个数字。例如输入一个长度为9的数组[1,2,3,2,2,2,5,4,2]。由于数字2在数组中出现了5次，
//超过数组长度的一半，因此输出2。
//示例：
//输入：[1,2,3,2,2,2,5,4,2]
//输出：2

// 方法1：哈希表法，该算法的时间复杂度是O(n)，空间复杂度也是O(n)。
func f1(arr []int) int {
	m := make(map[int]int)
	for _, v := range arr {
		m[v]++
	}

	num := 0
	target := 0
	for k, v := range m {
		if num < v {
			num = v
			target = k
		}
	}
	if num > len(arr)/2 {
		return target
	}

	return 0
}

//方法2 ： 排序算法
// 将数组进行排序，那排序后的数组的中点一定就是众数。
func f2(arr []int) int {
	sort.Sort(sort.IntSlice(arr))

	length := len(arr)
	return arr[length/2]
}

// 方法3：最经典的解法是Boyer-Moore 摩尔投票算法。
// Boyer-Moore 投票算法的核心思想是票数正负抵消，即遇到众数时，我们把票数+1，遇到非众数时，
// 我们把票数-1，则所有的票数和一定是大于0的。

//我们假设数组nums的众数是x，数组的长度为n。我们可以很容易的知道，若数组的前a个数字的票数和为0，
//那么剩余n-a个数字的票数和一定是大于0的，即后n-a个数字的众数仍然为x。
func f3(arr []int) int {
	count := 0
	res := 0
	for _, v := range arr {
		if count == 0 {
			res = v
		}

		if res == v {
			count++
		} else {
			count--
		}
	}

	return res
}

func main() {
	arr := []int{1, 2, 3, 2, 2, 2, 5, 5, 2}
	fmt.Println(f1(arr))
	fmt.Println(f2(arr))
	fmt.Println(f3(arr))
}
