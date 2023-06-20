package main

import (
	"fmt"
	"sort"
)

//LeetCode 349. 两个数组的交集
//题意：给定两个数组，编写一个函数来计算它们的交集。
// 说明： 输出结果中的每个元素一定是唯一的。 我们可以不考虑输出结果的顺序。

// 方法1：两层循环暴力遍历
//总时间复杂度是 O(mn)。

// 方法2：两个集合-官方
// 首先使用两个集合分别存储两个数组中的元素，然后遍历较小的集合，判断其中的每个元素是否在另一个集合中，
// 如果元素也在另一个集合中，则将该元素添加到返回值。该方法的时间复杂度可以降低到 O(m+n)。

//时间复杂度：O(m+n)，其中 m 和 n 分别是两个数组的长度。使用两个集合分别存储两个数组中的元素需要 O(m+n)
// 的时间，遍历较小的集合并判断元素是否在另一个集合中需要 O(min(m,n)) 的时间，因此总时间复杂度是 O(m+n)。
//空间复杂度：O(m+n)，其中 m 和 n 分别是两个数组的长度。空间复杂度主要取决于两个集合。
func intersection2(nums1 []int, nums2 []int) (intersection []int) {
	set1 := map[int]struct{}{}
	for _, v := range nums1 {
		set1[v] = struct{}{}
	}
	set2 := map[int]struct{}{}
	for _, v := range nums2 {
		set2[v] = struct{}{}
	}
	if len(set1) > len(set2) {
		set1, set2 = set2, set1
	}
	//fmt.Println(set1, set2)
	//map[4:{} 5:{} 9:{}] map[4:{} 8:{} 9:{}] 使用hash做去重处理
	for v := range set1 {
		if _, has := set2[v]; has {
			intersection = append(intersection, v)
		}
	}
	return
}

// 方法三：排序 + 双指针--官方
//时间复杂度：O(mlogm+nlogn)，其中 m 和 n 分别是两个数组的长度。对两个数组排序的时间复杂度分别是 O(mlogm)
// 和 O(nlogn)，双指针寻找交集元素的时间复杂度是 O(m+n)，因此总时间复杂度是 O(mlogm+nlogn)。
//空间复杂度：O(logm+logn)，其中 m 和 n 分别是两个数组的长度。空间复杂度主要取决于排序使用的额外空间。
func intersection3(nums1 []int, nums2 []int) (res []int) {
	sort.Ints(nums1)
	sort.Ints(nums2)
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		x, y := nums1[i], nums2[j]
		if x == y {
			// 如果结果数组为nil则可以直接追加，或者是当前的值比结果数组最后一个值都大则可以追加
			if res == nil || x > res[len(res)-1] {
				res = append(res, x)
			}
			i++
			j++
		} else if x < y {
			i++
		} else {
			j++
		}
	}
	return
}

// 方法4：其他人的-简版
func intersection4(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	for _, v := range nums1 {
		m[v] = 1
	}
	var res []int
	// 利用count>0，实现重复值只拿一次放入返回结果中
	for _, v := range nums2 {
		if count, ok := m[v]; ok && count > 0 {
			res = append(res, v)
			m[v]-- // 放入一个到结果里，就删除一个hash,保证不重复
		}
	}
	return res
}

//优化版，利用set，减少count统计
func intersection5(nums1 []int, nums2 []int) []int {
	set := make(map[int]struct{}, 0)
	res := make([]int, 0)
	for _, v := range nums1 {
		if _, ok := set[v]; !ok {
			set[v] = struct{}{}
		}
	}
	for _, v := range nums2 {
		//如果存在于上一个数组中，则加入结果集，并清空该set值
		if _, ok := set[v]; ok {
			res = append(res, v)
			delete(set, v) // 放入一个到结果里，就删除一个hash,保证不重复
		}
	}
	return res
}

func main() {
	nums1 := []int{4, 9, 5, 9, 9}
	nums2 := []int{9, 4, 9, 8, 4, 9}
	fmt.Println(intersection2(nums1, nums2))
	fmt.Println(intersection3(nums1, nums2))
	fmt.Println(intersection4(nums1, nums2))
	fmt.Println(intersection5(nums1, nums2))
}
