package main

import "fmt"

// todo
//LeetCode 4. 寻找两个正序数组的中位数(上中位数)

//给定两个递增数组arr1和arr2，已知两个数组的长度都为N，求两个数组中所有数的上中位数。
//上中位数：假设递增序列长度为n，为第n/2个数。
//要求：时间复杂度 O (n)，空间复杂度 O(1)。
//进阶：时间复杂度为O(logN)，空间复杂度为O(1)。

//示例：
//输入：[1，2，3，4]，[3，4，5，6]
//返回值：3
//说明：总共有8个数，上中位数是第4小的数，所以返回3。

// 方法1：归并方法-合并数组并排序
// 这道题最直观的想法就是，将两个有序数组合并成一个大的有序数组。
// 大的有序数组的上中位数为第n/2个数。我们可以知道该算法的时间复杂度是O(N)，空间复杂度也是O(N)。

// 方法2：二分查找,推荐
//我们也不需要合并两个有序数组，我们只需要找到上中位数的位置即可。对于给定两个长度为N的数组，
//我们可以知道其中位数的位置为N，所以我们维护两个指针，初始时分别指向两个数组的下标0的位置，
//每次将指向较小值的指针后移一位（如果一个指针已经到达数组的末尾，则只需要移动另一个数组的指针），
//直到到达中位数的位置。该算法的时间复杂度是O(n)，空间复杂度是O(1)。
func f1(arr1, arr2 []int) int {
	var target int
	targetPointer := len(arr1)
	arr1CurPointer := 0
	arr2CurPointer := 0

	for ; targetPointer > 0; targetPointer-- {
		if arr1[arr1CurPointer] <= arr2[arr2CurPointer] {
			target = arr1[arr1CurPointer]
			arr1CurPointer++
		} else if arr1[arr1CurPointer] > arr2[arr2CurPointer] {
			target = arr2[arr2CurPointer]
			arr2CurPointer++
		}
	}

	return target
}

// 方法2：二分查找 todo
// 对于长度为N的数组arr1和arr2来说，它的上中位数是两个有序数组的第N个元素。
//所以，我们把这道题目可以转换成寻找两个有序数组的第k小的元素，其中k=N。

//要找到第N个元素，我们可以比较arr1[N/2-1]和arr2[N/2-1]，其中“/”代表整数除法。
//由于arr1[N/2-1]和arr2[N/2-1]的前面分别有arr1[0...N/2-2]和arr2[0...N/2-2]，
//即N/2-1个元素。对于arr1[N/2-1]和arr2[N/2-1]的较小值，最多只会有N/2-1+N/2-1=N-2个元素比它小，
//所以它不是第N小的元素。

//因此，我们可以归纳出以下三种情况。
//  1 如果arr1[N/2-1] < arr2[N/2-1]，则比arr1[N/2-1] 小的数最多只有 arr1的前N/2-1个数和arr2的
//    前N/2-1个数，即比arr1[N/2-1] 小的数最多只有N-2个，因此arr1[N/2-1]不可能是第N个数，
//     arr1[0]到arr1[N/2-1]也都不可能是第N个数，所以可以删除。
//  2 如果arr1[N/2-1] > arr2[N/2-1]，则可以排除arr2[0]到arr2[N/2-1]。
//  3 如果arr1[N/2-1]==arr2[N/2-1]，可以归为第一种情况进行处理。
//可以看到，经过一轮比较后，我们可以排查N/2个不可能是第N小的数，查找范围缩小了一半。
//同时，我们将在排除后的新数组上继续进行二分查找，并且根据我们排除数的个数，减少 N 的值，
//这是因为我们排除的数都不大于第 N 小的数。
func f2(arr1, arr2 []int) int {
	length := len(arr1)
	//如果N=1,直接返回两个数组的首元素的最小值即可
	if length == 1 {
		return min(arr1[0], arr2[0])
	}

	index1, index2 := 0, 0
	//中位数位置为length,不超过分区数组的下标
	k := length

	for k > 1 {
		new_index1 := index1 + k/2 - 1
		new_index2 := index2 + k/2 - 1
		data1, data2 := arr1[new_index1], arr2[new_index2]
		//选择较小值，同时将前k/2个元素删除
		if data1 <= data2 {
			k = k - k/2
			// 删除前k     //2个元素
			index1 = new_index1 + 1
		} else {

			k = k - k/2
			// 删除前k     //2个元素
			index2 = new_index2 + 1
		}
	}

	return min(arr1[index1], arr2[index2])
}
func min(n1, n2 int) int {
	if n1 <= n2 {
		return n1
	} else {
		return n2
	}
}

func main() {
	arr1 := []int{1, 2, 2, 5}
	arr2 := []int{3, 4, 5, 6}

	fmt.Println(f1(arr1, arr2))
	fmt.Println(f2(arr1, arr2))

	arr3 := []int{1, 2, 2, 5}
	arr4 := []int{3, 4, 5, 6, 7}
	fmt.Println(f3(arr3, arr4))
	fmt.Println(f4(arr3, arr4))
}

// https://leetcode-cn.com/problems/median-of-two-sorted-arrays/solution/xun-zhao-liang-ge-you-xu-shu-zu-de-zhong-wei-s-114/

// 加餐
// 如果给定的两个有序数组大小不同，即给定两个大小分别为 m 和 n 的正序（从小到大）数组
//nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。

//根据中位数的定义，当m+n为奇数时，中位数是两个有序数组的第(m+n+1)/2个元素。
//当m+n为偶数时，中位数是两个有序数组的第(m+n)/2和(m+n)/2+1个元素的平均值。
//因此这道题我们可以转化成寻求两个有序数组的第k小的数，其中k为(m+n)/2或者(m+n)/2+1。

//所以该题的解题思路和上一题的解法类似，不过这里有一些情况需要特殊处理。
// 1 如果nums1[k/2-1]或者nums[k/2-1]越界，那么我们需要选择对应数组的最后一个元素，
//   即min(k/2-1,m-1)或者min(k/2-1,n-1)。
// 2 如果一个数组为空，我们可以直接返回另一个数组中第 k 小的元素。
// 3 如果k=1，我们只需要返回两个数组首元素的最小值即可。
func f3(arr1, arr2 []int) int {

	return 0
}

// 双指针法
func f4(arr1, arr2 []int) int {
	len1 := len(arr1)
	len2 := len(arr2)
	length := len1 + len2

	maxKey := length/2 + 1
	curKey1, curKey2 := 0, 0
	cur1, cur2, cur := 0, 0, 0

	var arr []int

	for ; maxKey > 0; maxKey-- {
		if curKey1 >= len1 {
			cur2 = arr2[curKey2]
			cur = cur2
			curKey2++
		} else if curKey2 >= len2 {
			cur1 = arr1[curKey1]
			cur = cur1
			curKey1++
		} else {
			cur1 = arr1[curKey1]
			cur2 = arr2[curKey2]
			if cur1 <= cur2 {
				curKey1++
				cur = cur1
			} else {
				curKey2++
				cur = cur2
			}
		}

		if maxKey <= 2 {
			arr = append(arr, cur)
		}
	}

	fmt.Println(arr)
	if length%2 == 0 { // 偶数
		return (arr[0] + arr[1]) / 2
	} else {
		return arr[1]
	}
}
