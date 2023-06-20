package main

import "fmt"

//LeetCode 46.全排列

//给定一个 没有重复 数字的序列，返回其所有可能的全排列。

//示例:
//输入: [1,2,3]
//输出: [ [1,2,3], [1,3,2], [2,1,3], [2,3,1], [3,1,2], [3,2,1] ]

// 卡尔
// 首先排列是有序的，也就是说 [1,2] 和 [2,1] 是两个集合，这和之前分析的子集以及组合所不同的地方。

//可以看出元素1在[1,2]中已经使用过了，但是在[2,1]中还要在使用一次1，所以处理排列问题就不用使用startIndex了。

//但排列问题需要一个used数组，标记已经选择的元素，其实就是记录此时path里都有哪些元素使用了，
// ****一个排列里一个元素只能使用一次。*****

//排列问题的不同：
// 每层都是从0开始搜索而不是startIndex
// 需要used数组记录path里都放了哪些元素了

var res [][]int

func permute(nums []int) [][]int {
	res = [][]int{}
	backTrack(nums, len(nums), []int{})
	return res
}
func backTrack(nums []int, numsLen int, path []int) {
	if len(nums) == 0 {
		p := make([]int, len(path))
		copy(p, path)
		res = append(res, p)
		return // 这里不写也可以因为下面也不会再循环了
	}

	for i := 0; i < numsLen; i++ {
		cur := nums[i]

		path = append(path, cur)
		nums = append(nums[:i], nums[i+1:]...)
		//直接使用切片，去掉当前层的这个元素，所以就不需要用used数组了

		// 入参元素：排除掉当前元素的数组
		backTrack(nums, len(nums), path)

		nums = append(nums[:i], append([]int{cur}, nums[i:]...)...)
		//回溯的时候切片也要复原(把cur补穿插进去)，元素位置不能变

		path = path[:len(path)-1]
	}
}

func main() {
	arr := []int{1, 2, 3}
	fmt.Println(permute(arr))
}

// 官方
// https://leetcode.cn/problems/permutations/solution/quan-pai-lie-by-leetcode-solution-2/
