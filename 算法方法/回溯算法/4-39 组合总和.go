package main

import "fmt"

//LeetCode 39. 组合总和

//给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

//candidates 中的数字可以无限制重复被选取。

//说明：
//所有数字（包括 target）都是正整数。
//解集不能包含重复的组合。

//示例 1： 输入：candidates = [2,3,6,7], target = 7, 所求解集为： [ [7], [2,2,3] ]
//示例 2： 输入：candidates = [2,3,5], target = 8, 所求解集为： [   [2,2,2,2],   [2,3,3],   [3,5] ]

var trcak []int
var res [][]int

// 卡尔
// https://www.programmercarl.com/0039.%E7%BB%84%E5%90%88%E6%80%BB%E5%92%8C.html

func combinationSum55(candidates []int, target int) [][]int {
	backtracking55(0, 0, candidates, target)
	return res
}
func backtracking55(startIndex, sum int, candidates []int, target int) {
	//终止条件
	if sum == target {
		tmp := make([]int, len(trcak))
		copy(tmp, trcak)       //拷贝
		res = append(res, tmp) //放入结果集
		return
	}
	if sum > target { // 剪枝
		return
	}

	//回溯
	for i := startIndex; i < len(candidates); i++ {
		//更新路径集合和sum
		trcak = append(trcak, candidates[i])
		sum += candidates[i]
		//递归
		backtracking55(i, sum, candidates, target) // 注意这i不用++
		//回溯
		trcak = trcak[:len(trcak)-1]
		sum -= candidates[i]
	}
}

func main() {
	arr := []int{2, 3, 5}
	target := 8
	fmt.Println(combinationSum55(arr, target))
}

//方法1：回溯-官方
func combinationSum(candidates []int, target int) (ans [][]int) {
	comb := []int{}
	var dfs func(target, idx int)
	dfs = func(target, idx int) {
		if idx == len(candidates) {
			return
		}
		if target == 0 {
			ans = append(ans, append([]int(nil), comb...))
			return
		}
		// 直接跳过
		dfs(target, idx+1)
		// 选择当前数
		if target-candidates[idx] >= 0 {
			comb = append(comb, candidates[idx])
			dfs(target-candidates[idx], idx)
			comb = comb[:len(comb)-1]
		}
	}
	dfs(target, 0)
	return
}
