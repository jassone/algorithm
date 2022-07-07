package main

import (
	"fmt"
)

//LeetCode 128. 最长连续序列

//给定无序数组arr，返回其中最长的连续子序列的长度(要求值连续，位置可以不连续，
// 例如 3,4,5,6为连续的自然数）。请你设计并实现时间复杂度为 O(n) 的算法解决此问题。

//示例：
//输入：[100,4,200,1,3,2]
//输出：4
//说明：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。

// 方法1：暴力法(配合hash法)
func f1(arr []int) int {
	hashmap := make(map[int]bool)
	for _, i := range arr {
		hashmap[i] = true
	}
	// O(n)

	result := 0
	for current := range hashmap { // O(n)
		cnt := 1
		for hashmap[current+1] { // O(n)
			current++
			cnt++
		}
		if result < cnt {
			result = cnt
		}
	}
	// 综合
	//O(n*n)
	return result
}

//方法2：优化后的-官方
//首先把所有数字存在map里面，然后遍历map，此时要对current进行判断，如果hashmap[current-1]存在，
//那就跳过这次循环，因为我们要去找最小的一个数，如果不存在，这么那个数字必定是某个连续序列中首位数字，
//这个时候去计数hashmap[current+1]有几个，即可。外层循环需要 O(n)的时间复杂度，只有当一个数是连续
//序列的第一个数的情况下才会进入内层循环，然后在内层循环中匹配连续序列中的数，因此数组中的每个数只会
//进入内层循环一次。
//时间复杂度：O(n)，其中 n 为数组的长度。
//空间复杂度：O(n)。哈希表存储数组中所有的数需要 O(n) 的空间。
func f2(nums []int) int {
	hashmap := make(map[int]bool)
	for _, i := range nums {
		hashmap[i] = true
	}
	// 时间复杂度O(n)

	result := 0
	for current := range hashmap { // 时间复杂度O(n)
		if !hashmap[current-1] { //不存在比current小的数
			cnt := 1
			for hashmap[current+1] { // 时间复杂度最大的一个是O(n-1)
				current++
				cnt++
			}
			if result < cnt {
				result = cnt
			}
		}
	}
	// O(n+n+n-1) = O(n)
	return result
}

func main() {
	arr := []int{8, 1, 4, 3, 2, 5}
	fmt.Println(f1(arr))
	fmt.Println(f2(arr))
}
