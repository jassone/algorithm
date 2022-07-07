package main

import (
	"fmt"
	"sort"
)

// LeetCode 88. 合并两个有序数组

//给你两个按非递减顺序排列的整数数组nums1和nums2，另有两个整数m和n，
// 分别表示nums1和nums2中的元素数目。请你合并nums2到nums1中，
// 使合并后的数组同样按非递减顺序排列。
//注意：最终，合并后数组不应由函数返回，而是存储在数组nums1中。为了应对这种情况，
// nums1的初始长度为m + n，其中前 m 个元素表示应合并的元素，后 n 个元素为 0 ，应忽略。nums2 的长度为 n 。
//示例：
//输入：nums1 = [1,2,3,0,0,0]， m = 3，nums2 = [2,5,6]，n = 3
//输出：[1,2,2,3,5,6]
//解释：需要合并 [1,2,3] 和 [2,5,6] 。合并结果是 [1,2,2,3,5,6] 。

// 方案1，暴力-官方
// 最简单暴力的方法就是直接把nums2放入nums1的后n个位置，
// 然后直接对nums1进行排序就好了。
//时间复杂度： O((m+n)log(m+n))。
//排序序列长度为  m+n，套用快速排序的时间复杂度即可，平均情况为  O((m+n)log(m+n))。
//空间复杂度： O(log(m+n))。
//排序序列长度为  m+n，套用快速排序的空间复杂度即可，平均情况为  O(log(m+n))。
func f1(arr1 []int, arr1Len int, arr2 []int, arr2Len int) []int {
	//arr1 = append(arr1[:arr1Len], arr2...) 如果是append,会多出来三个0
	copy(arr1[arr1Len:], arr2)
	//sort.Sort(sort.IntSlice(arr1))
	sort.Ints(arr1)
	return arr1
}

// 方案2，双指针法-官方，推荐
// 那么既然给定的两个数组是有序的，那我们何不把这个条件利用起来，来优化代码。
// 所以，我们可以使用两个指针p1和p2分别指向两个数组的起始位置，然后比较大小，
// 将较小的放入结果中，然后指针后移，直到将所有元素都排好序。
//时间复杂度：O(m+n)。
// 指针移动单调递增，最多移动 m+n 次，因此时间复杂度为 O(m+n)。
//空间复杂度：O(m+n)。
// 需要建立长度为m+n 的中间数组 sorted。
func merge(nums1 []int, m int, nums2 []int, n int) []int {
	sorted := make([]int, 0, m+n)
	p1, p2 := 0, 0
	for {
		if p1 == m { // 第一个数组处理完了
			sorted = append(sorted, nums2[p2:]...)
			break
		}
		if p2 == n { // 第二个数组处理完了
			sorted = append(sorted, nums1[p1:]...)
			break
		}
		if nums1[p1] < nums2[p2] {
			sorted = append(sorted, nums1[p1])
			p1++
		} else {
			sorted = append(sorted, nums2[p2])
			p2++
		}
	}
	copy(nums1, sorted)

	return nums1
}

// 自己写的
func f2(arr1 []int, arr1Len int, arr2 []int, arr2Len int) []int {
	result := []int{}
	p1, p2 := 0, 0
	for p1 < arr1Len && p2 < arr2Len {
		if arr1[p1] <= arr2[p2] {
			result = append(result, arr1[p1])
			p1++
		} else {
			result = append(result, arr2[p2])
			p2++
		}
	}

	if p1 == arr1Len { // 说明arr1读取完了,处理arr2
		for _, v := range arr2[p2:arr2Len] {
			result = append(result, v)
		}
	} else { // 说明arr2读取完了,处理arr1
		for _, v := range arr1[p1:arr1Len] {
			result = append(result, v)
		}
	}

	return result
}

// 方法三：逆向双指针-官方
//观察可知，nums1的后半部分是空的，可以直接覆盖而不会影响结果。因此可以指针设置为从后向前遍历，
//每次取两者之中的较大者放进 nums1 的最后面。
//时间复杂度：O(m+n)。
// 指针移动单调递增，最多移动 m+n 次，因此时间复杂度为 O(m+n)。
//空间复杂度：O(1)。
func merge3(nums1 []int, m int, nums2 []int, n int) []int {
	for p1, p2, tail := m-1, n-1, m+n-1; p1 >= 0 || p2 >= 0; tail-- {
		var cur int
		if p1 == -1 { // 第一个数组已经处理完
			cur = nums2[p2]
			p2--
		} else if p2 == -1 { // 第二个数组已经处理完
			cur = nums1[p1]
			p1--
		} else if nums1[p1] > nums2[p2] {
			cur = nums1[p1]
			p1--
		} else {
			cur = nums2[p2]
			p2--
		}
		nums1[tail] = cur
	}

	return nums1
}

func main() {
	arr1 := []int{1, 2, 3, 0, 0, 0}
	arr1Len := 3
	arr2 := []int{2, 5, 6, 0, 0, 0}
	arr2Len := 3
	fmt.Println(f1(arr1, arr1Len, arr2, arr2Len))
	fmt.Println(merge(arr1, arr1Len, arr2, arr2Len))
	fmt.Println(merge3(arr1, arr1Len, arr2, arr2Len))
}
