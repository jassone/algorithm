package main

import "fmt"

//LeetCode 41. 缺失的第一个正数

//给你一个无重复元素未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
//请你实现时间复杂度为O(n)并且只使用常数级别额外空间的解决方案。

//示例：
//输入：nums = [1，2，0]
//输出：3

//方法1 ：哈希表，推荐
//对于一个无重复元素、长度为N的数组，其中没有出现的最小整数只能在[1,N+1]中，这是因为如果[1,N]
//在数组中都出现了，说明这N个数已经把数组填满了，那么答案是N+1，否则就是[1，N]中没有出现的最小整数。
//所以，我们可以申请一个辅助数组temp，大小为N，我们通过遍历原数组，将属于[1,N]范围内的数，
//放入辅助数组中相应的位置，使得temp[i-1] = i 成立。在遍历完成后，temp中第一个不满足temp[i-1] = i
//条件的就是最小的正整数，如果都满足，那么最小正整数就是N+1。

//我们可以知道该算法的时间复杂度和空间复杂度都是O(n)，
func f1(arr []int) int {
	var res int
	length := len(arr)
	newArr := make([]int, length)

	// 取出所有正整数，并进行自动排序
	for _, v := range arr {
		if v >= 1 && v <= length {
			newArr[v-1] = v
		}
	}

	for k, _ := range newArr {
		if newArr[k] != k+1 {
			res = k + 1
			break
		}
	}
	// 如果没找到，说明前面都是连续的，则要取最后一个数字+1
	if res == 0 {
		res = arr[length-1] + 1
	}

	return res
}

//官方 todo
func firstMissingPositive1(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] <= 0 {
			nums[i] = n + 1
		}
	}
	for i := 0; i < n; i++ {
		num := abs(nums[i])
		if num <= n {
			fmt.Println(num - 1)
			nums[num-1] = -abs(nums[num-1])
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return n + 1
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

//方法2：置换法-官方
// 通过观察，我们可以发现辅助数组和原数组大小一样，那么我们能否复用原数组nums呢？答案显然是可以的。
//我们在遍历数组的过程中，假设遍历到的元素值为x，如果x属于[1,N]，我们将元素x和nums[x-1]的元素进行互换，
//使得x出现在正确的位置上，否则不做处理。当遍历完成后，nums中第一个不满足nums[i-1] = i 条件的就是最小的
//正整数，如果都满足，那么最小正整数就是N+1。
// 该算法的时间复杂度是O(n)，空间复杂度是O(1)。
func firstMissingPositive2(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}

func main() {
	arr := []int{1, 2, 3, 0, 8}
	fmt.Println(f1(arr))
	fmt.Println(firstMissingPositive2(arr))
}
