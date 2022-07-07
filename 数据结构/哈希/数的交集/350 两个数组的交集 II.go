package main

import "sort"

//LeetCode 350.两个数组的交集 II
//给你两个整数数组 nums1 和 nums2 ，请你以数组形式返回两数组的交集。返回结果中每个元素出现的次数，应与元素在两个数组
// 中都出现的次数一致（如果出现次数不一致，则考虑取较小值）。可以不考虑输出结果的顺序。

//输入：nums1 = [1,2,2,1], nums2 = [2,2]
//输出：[2,2]

//方法一：哈希表-官方
//时间复杂度：O(m+n)，其中 m 和 n 分别是两个数组的长度。需要遍历两个数组并对哈希表进行操作，哈希表操作的时间复杂度是
// O(1)，因此总时间复杂度与两个数组的长度和呈线性关系。
//空间复杂度：O(min(m,n))，其中 m 和 n 分别是两个数组的长度。对较短的数组进行哈希表的操作，哈希表的大小不会超过较短
// 的数组的长度。为返回值创建一个数组 intersection，其长度为较短的数组的长度。
func intersect1(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return intersect1(nums2, nums1)
	}
	m := map[int]int{}
	for _, num := range nums1 {
		m[num]++
	}

	intersection := []int{}
	for _, num := range nums2 {
		if m[num] > 0 {
			intersection = append(intersection, num)
			m[num]-- // 找到了相同的，放入结果集中，并删除一个
		}
	}
	return intersection
}

//方法二：排序 + 双指针-官方
//时间复杂度：O(mlogm+nlogn)，其中 m 和 n 分别是两个数组的长度。对两个数组进行排序的时间复杂度是 O(mlogm+nlogn)，
// 遍历两个数组的时间复杂度是 O(m+n)，因此总时间复杂度是 O(mlogm+nlogn)。
//空间复杂度：O(min(m,n))，其中 m 和 n 分别是两个数组的长度。为返回值创建一个数组 intersection，其长度为较短的
// 数组的长度。不过在 C++ 中，我们可以直接创建一个 vector，不需要把答案临时存放在一个额外的数组中，所以这种实现的
// 空间复杂度为 O(1)。
func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	length1, length2 := len(nums1), len(nums2)
	index1, index2 := 0, 0

	intersection := []int{}
	for index1 < length1 && index2 < length2 {
		if nums1[index1] < nums2[index2] {
			index1++
		} else if nums1[index1] > nums2[index2] {
			index2++
		} else {
			intersection = append(intersection, nums1[index1])
			index1++
			index2++
		}
	}
	return intersection
}

// 总结
//如果 nums2的元素存储在磁盘上，磁盘内存是有限的，并且你不能一次加载所有的元素到内存中。那么就无法高效地对 nums2
//进行排序，因此推荐使用方法一而不是方法二。在方法一中，nums2 只关系到查询操作，因此每次读取 nums2中的一部分数据，
//并进行处理即可。

func main() {

}
