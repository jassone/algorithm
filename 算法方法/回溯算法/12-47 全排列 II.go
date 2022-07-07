package main

import (
	"fmt"
	"sort"
)

//LeetCode 47.全排列 II

//给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。

//示例 1：
//输入：nums = [1,1,2]
//输出： [[1,1,2], [1,2,1], [2,1,1]]

//示例 2：
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
//提示：

//1 <= nums.length <= 8
//-10 <= nums[i] <= 10

// 卡尔

//这道题目和46.全排列 的区别在与给定一个可包含重复数字的序列，要返回所有不重复的全排列。
//这里又涉及到去重了。

//在40.组合总和II 、90.子集II 我们分别详细讲解了组合问题和子集问题如何去重。

//那么排列问题其实也是一样的套路。
//****还要强调的是去重一定要对元素进行排序，这样我们才方便通过相邻的节点来判断是否重复使用了。****

var res [][]int

// 方法1：使用map，used不加入到回溯里
func permute(nums []int) [][]int {
	res = [][]int{}
	// 这里不需要对原始数组排序
	backTrack(nums, len(nums), []int{})
	return res
}
func backTrack(nums []int, numsLen int, path []int) {
	// 原始原始为空，则返回
	if len(nums) == 0 {
		p := make([]int, len(path))
		copy(p, path)
		res = append(res, p)
	}

	// 使用map
	used := make(map[int]bool)
	for i := 0; i < numsLen; i++ {
		//同一层不使用重复的数。
		if used[nums[i]] == true {
			continue
		}

		cur := nums[i]
		path = append(path, cur)

		used[nums[i]] = true // 因为只需要同层判断，所以used不需要回溯
		nums = append(nums[:i], nums[i+1:]...)

		backTrack(nums, len(nums), path)

		nums = append(nums[:i], append([]int{cur}, nums[i:]...)...)
		path = path[:len(path)-1]
	}
}

// 方法2：使用切片，used加入到回溯里
func permute2(nums []int) [][]int {
	res = [][]int{}
	// 这里需要对原始数组排序
	sort.Ints(nums)

	// 使用数组
	used := make([]bool, len(nums))

	backTrack2(nums, len(nums), []int{}, used)
	return res
}
func backTrack2(nums []int, numsLen int, path []int, used []bool) {
	// 原始元素为空，则返回
	if len(nums) == len(path) {
		p := make([]int, len(path))
		copy(p, path)
		res = append(res, p)
		return
	}

	for i := 0; i < numsLen; i++ {
		//同一树枝不使用重复的数。
		// 树层去重，效率更高
		//if i > 0 && nums[i-1] == nums[i] && used[i-1] == true {
		//	continue
		//}

		//同一层不使用重复的数。
		if i > 0 && nums[i-1] == nums[i] && used[i-1] == false {
			continue
		}

		// 我们使用一个 used 数组记录使用过的数字，使用过了就不再使用：
		// 因为每次传递到递归函数里面的都是完整的原始数组
		if used[i] == true {
			continue
		}

		path = append(path, nums[i])
		used[i] = true
		backTrack2(nums, len(nums), path, used)
		path = path[:len(path)-1]
		used[i] = false
	}
}

func main() {
	arr := []int{1, 1, 2}
	//fmt.Println(permute(arr))
	fmt.Println(permute2(arr))
}

// 官方
