package main

import (
	"fmt"
)

//LeetCode 78.子集

//给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。

//说明：解集不能包含重复的子集。
//示例: 输入: nums = [1,2,3] 输出: [ [3],   [1],   [2],   [1,2,3],   [1,3],   [2,3],   [1,2],   [] ]

// 注意给定的条件是数组中数字不重复

// 卡尔
// https://www.programmercarl.com/0078.%E5%AD%90%E9%9B%86.html

//如果把 子集问题、组合问题、分割问题都抽象为一棵树的话，那么组合问题和分割问题都是收集树的叶子节点，而子集问题是找树的所有节点！
//其实子集也是一种组合问题，因为它的集合是无序的，子集{1,2} 和 子集{2,1}是一样的。

//那么既然是无序，取过的元素不会重复取，写回溯算法的时候，for就要从startIndex开始，而不是从0开始！

//有同学问了，什么时候for可以从0开始呢？
//求排列问题的时候，就要从0开始，因为集合是有序的，{1, 2} 和{2, 1}是两个集合，排列问题我们后续的文章就会讲到的。

var res [][]int

func subset555(nums []int) [][]int {
	res = make([][]int, 0)
	Dfs([]int{}, nums, 0)
	return res
}
func Dfs(temp, nums []int, start int) {
	// 这里没有条件，因为所有节点都可收集
	tmp := make([]int, len(temp))
	copy(tmp, temp)
	res = append(res, tmp)

	for i := start; i < len(nums); i++ {
		temp = append(temp, nums[i])
		Dfs(temp, nums, i+1)
		temp = temp[:len(temp)-1]
	}
}

func main() {
	arr := []int{1, 2, 3}
	fmt.Println(subset555(arr))
}

//官方-太复杂
//https://leetcode.cn/problems/subsets/solution/zi-ji-by-leetcode-solution/
