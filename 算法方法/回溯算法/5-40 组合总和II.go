package main

import (
	"fmt"
	"sort"
)

//LeetCode 40.组合总和II

//给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

//candidates 中的每个数字在每个组合中只能使用一次。

//说明： 所有数字（包括目标数）都是正整数。 解集不能包含重复的组合。

//示例 1: 输入: candidates = [10,1,2,7,6,1,5], target = 8, 所求解集为: [ [1, 7], [1, 2, 5], [2, 6], [1, 1, 6] ]
//示例 2: 输入: candidates = [2,5,2,1,2], target = 5, 所求解集为: [   [1,2,2],   [5] ]

// 卡尔
// https://www.programmercarl.com/0040.%E7%BB%84%E5%90%88%E6%80%BB%E5%92%8CII.html

//最后本题和39.组合总和 要求一样，解集不能包含重复的组合。
//本题的难点在于区别2中：集合（数组candidates）有重复元素，但还不能有重复的组合。

//一些同学可能想了：我把所有组合求出来，再用set或者map去重，这么做很容易超时！

//由于可能超时，所以要在搜索的过程中就去掉重复组合。

//注意：我们要去重的是同一树层上的“使用过”，同一树枝上的都是一个组合里的元素，不用去重。

var trcak = []int{}
var res = [][]int{}
var used map[int]bool

// 方法1：使用used map
func combinationSum2(candidates []int, target int) [][]int {
	used = make(map[int]bool)
	sort.Ints(candidates)
	backtracking(0, 0, target, candidates, used)
	return res
}
func backtracking(startIndex, sum, target int, candidates []int, used map[int]bool) {
	//终止条件
	if sum == target {
		tmp := make([]int, len(trcak))
		copy(tmp, trcak)       //拷贝
		res = append(res, tmp) //放入结果集
		return
	}
	if sum > target {
		return
	}
	// 这个判断条件挪到下面for里面也可以

	//回溯
	// 因为candidates已经排序过了，且used也加入到回溯里
	// 当candidates[i] == candidates[i-1]
	// used[i - 1] == true，说明同一树枝candidates[i - 1]使用过，即还未回溯，递归进入了下一个了
	// used[i - 1] == false，说明同一树层candidates[i - 1]使用过，即回溯了，这时候进入了下一个循环
	for i := startIndex; i < len(candidates); i++ {
		if i > 0 && candidates[i] == candidates[i-1] && used[i-1] == false {
			continue
		}

		//更新路径集合和sum
		trcak = append(trcak, candidates[i])
		sum += candidates[i]
		used[i] = true

		// 上面这几行代码变量，在同一个枝条上是可以公用的，所以判断used[i-1] == false表示不在同一个枝条上
		// 即处于相同层

		//递归
		backtracking(i+1, sum, target, candidates, used)

		//回溯
		trcak = trcak[:len(trcak)-1]
		sum -= candidates[i]
		used[i] = false
	}
}

// 不使用used
func combinationSum2666(candidates []int, target int) [][]int {
	var trcak []int
	var res [][]int
	sort.Ints(candidates)
	backtracking666(0, 0, target, candidates, trcak, &res)
	return res
}
func backtracking666(startIndex, sum, target int, candidates, trcak []int, res *[][]int) {
	//终止条件
	if sum == target {
		tmp := make([]int, len(trcak))
		//拷贝
		copy(tmp, trcak)
		//放入结果集
		*res = append(*res, tmp)
		return
	}

	//回溯
	for i := startIndex; i < len(candidates) && sum+candidates[i] <= target; i++ {
		// 若当前树层有使用过相同的元素，则跳过
		// 注意这里的判断是同一层 for 循环
		if i > startIndex && candidates[i] == candidates[i-1] {
			continue
		}

		//更新路径集合和sum
		trcak = append(trcak, candidates[i])
		sum += candidates[i]
		backtracking666(i+1, sum, target, candidates, trcak, res)
		//回溯
		trcak = trcak[:len(trcak)-1]
		sum -= candidates[i]
	}
}

func main() {
	arr := []int{1, 1, 2}
	target := 3
	//fmt.Println(combinationSum2(arr, target))
	fmt.Println(combinationSum2666(arr, target))
}

// 方法1：回溯-官方
func combinationSum2333(candidates []int, target int) (ans [][]int) {
	sort.Ints(candidates)
	var freq [][2]int
	for _, num := range candidates {
		if freq == nil || num != freq[len(freq)-1][0] {
			freq = append(freq, [2]int{num, 1})
		} else {
			freq[len(freq)-1][1]++
		}
	}

	var sequence []int
	var dfs func(pos, rest int)
	dfs = func(pos, rest int) {
		if rest == 0 {
			ans = append(ans, append([]int(nil), sequence...))
			return
		}
		if pos == len(freq) || rest < freq[pos][0] {
			return
		}

		dfs(pos+1, rest)

		most := min(rest/freq[pos][0], freq[pos][1])
		for i := 1; i <= most; i++ {
			sequence = append(sequence, freq[pos][0])
			dfs(pos+1, rest-i*freq[pos][0])
		}
		sequence = sequence[:len(sequence)-most]
	}
	dfs(0, target)
	return
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
