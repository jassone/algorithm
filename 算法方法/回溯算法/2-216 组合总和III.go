package main

import "fmt"

// 同类型 77 216
//LeetCode 216.组合总和III

//找出所有相加之和为 n 的 k 个数的组合。组合中只允许含有 1 - 9 的正整数，并且每种组合中不存在重复的数字。

//说明：
// 所有数字都是正整数。
// 解集不能包含重复的组合。

//示例 1: 输入: k = 3, n = 7 输出: [[1,2,4]]
//示例 2: 输入: k = 3, n = 9 输出: [[1,2,6], [1,3,5], [2,3,4]]

// 卡尔
//https://www.programmercarl.com/0216.%E7%BB%84%E5%90%88%E6%80%BB%E5%92%8CIII.html

var res [][]int
var track []int

// 方法1：在77题目上微调
func combinationSum3(k int, n int) [][]int {
	res = [][]int{}
	track = []int{}
	if n <= 0 || k <= 0 || k < n { // 对异常数据进行初始判断
		return res
	}

	backtrack(k, n, 1) // 递归函数的返回值以及参数，开始真正回溯
	return res
}
func backtrack(k, n, start int) {
	if len(track) == n { // 终止条件
		sum := 0
		for _, v := range track {
			sum += v
		}
		if sum == k {
			temp := make([]int, n) // 这里必须用临时变量转换下，因为track是全局变量，后面会被覆盖
			copy(temp, track)
			res = append(res, temp)
		}
		return
	}
	//剪枝：track 长度加上区间 [start, n] 的长度小于 n，不可能构造出长度为 n 的 track
	if len(track)+(n-start+1) < n {
		return
	}

	// 或者放在for里面剪枝
	// for i:=startIndex;i<=9-(n-len(*track))+1;i++{//减枝（n-len(*track)表示还剩多少个可填充的元素）

	for i := start; i <= 9; i++ { // 单层搜索的过程
		track = append(track, i)
		backtrack(k, n, i+1)
		track = track[:len(track)-1]
	}
}

// 方法2：多个参数参与回溯 --在77题目上微调
var sum int

func combinationSum33(k int, n int) [][]int {
	res = [][]int{}
	track = []int{}
	sum = 0
	if n <= 0 || k <= 0 || k < n { // 对异常数据进行初始判断
		return res
	}

	backtrack33(k, n, 1) // 递归函数的返回值以及参数，开始真正回溯
	return res
}
func backtrack33(k, n, start int) {
	if len(track) == n { // 终止条件
		if sum == k {
			temp := make([]int, n) // 这里必须用临时变量转换下，因为track是全局变量，后面会被覆盖
			copy(temp, track)
			res = append(res, temp)
		}
		return
	}
	//剪枝：track 长度加上区间 [start, n] 的长度小于 n，不可能构造出长度为 n 的 track
	if len(track)+(n-start+1) < n {
		return
	}

	for i := start; i <= 9; i++ { // 单层搜索的过程
		sum += i // 也参与回溯
		track = append(track, i)
		backtrack33(k, n, i+1) // 注意i+1调整startIndex
		sum -= i               // 也参与回溯-再减掉
		track = track[:len(track)-1]
	}
}

func main() {
	fmt.Println(combinationSum3(7, 3))
	fmt.Println(combinationSum33(7, 3))
}

//https://leetcode.cn/problems/combination-sum-iii/solution/zu-he-zong-he-iii-by-leetcode-solution/
//方法一：二进制（子集）枚举-官方
func combinationSum34(k int, n int) (ans [][]int) {
	var temp []int
	check := func(mask int) bool {
		temp = nil
		sum := 0
		for i := 0; i < 9; i++ {
			if 1<<i&mask > 0 {
				temp = append(temp, i+1)
				sum += i + 1
			}
		}
		return len(temp) == k && sum == n
	}

	for mask := 0; mask < 1<<9; mask++ {
		if check(mask) {
			ans = append(ans, append([]int(nil), temp...))
		}
	}
	return
}

//方法二：组合枚举-官方
func combinationSum35(k int, n int) (ans [][]int) {
	var temp []int
	var dfs func(cur, rest int)
	dfs = func(cur, rest int) {
		// 找到一个答案
		if len(temp) == k && rest == 0 {
			ans = append(ans, append([]int(nil), temp...))
			return
		}
		// 剪枝：跳过的数字过多，后面已经无法选到 k 个数字
		if len(temp)+10-cur < k || rest < 0 {
			return
		}
		// 跳过当前数字
		dfs(cur+1, rest)
		// 选当前数字
		temp = append(temp, cur)
		dfs(cur+1, rest-cur)
		temp = temp[:len(temp)-1]
	}
	dfs(1, n)
	return
}
