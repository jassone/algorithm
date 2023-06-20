package main

//LeetCode 77 组合
//给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。

//示例:
//输入: n = 4, k = 2
//输出:
//[
//[2,4],
//[3,4],
//[2,3],
//[1,2],
//[1,3],
//[1,4],
//]

//思路 卡尔
//https://www.programmercarl.com/0077.%E7%BB%84%E5%90%88.html

//本题这是回溯法的经典题目。

//直接的解法当然是使用for循环，例如示例中k为2，很容易想到 用两个for循环，这样就可以输出 和示例中一样的结果。

// ****如果是一个集合来求组合的话，就需要startIndex****

// 详细wiki: https://www.programmercarl.com/0077.%E7%BB%84%E5%90%88.html#%E5%9B%9E%E6%BA%AF%E6%B3%95%E4%B8%89%E9%83%A8%E6%9B%B2

var res [][]int
var track []int

// 同类型 77 216, 推荐
// +剪枝
func combine33(n int, k int) [][]int {
	res = [][]int{}
	track = []int{}
	if n <= 0 || k <= 0 || k > n { // 对异常数据进行初始判断
		return res
	}
	backtrack(n, k, 1) // 递归函数的返回值以及参数，开始真正回溯
	return res
}
func backtrack(n, k, start int) {
	if len(track) == k { // 终止条件
		temp := make([]int, k) // 这里必须用临时变量转换下，因为track是全局变量，后面会被覆盖
		copy(temp, track)
		res = append(res, temp)
		return
	}
	//剪枝：track 长度加上区间 [start, n] 的长度小于 k，不可能构造出长度为 k 的 track
	if len(track)+(n-start+1) < k {
		return
	}

	for i := start; i <= n; i++ { // 单层搜索的过程
		track = append(track, i)     // 放进去
		backtrack(n, k, i+1)         // 注意i+1调整startIndex
		track = track[:len(track)-1] // 再放出来
	}
}

func main() {
	//fmt.Println(combine33(4, 2))
	//fmt.Println(combine2(4, 2))
}

//方法一：递归实现组合型枚举-官方-不推荐

//思路与算法
//https://leetcode.cn/problems/combinations/solution/zu-he-by-leetcode-solution/
func combine1(n int, k int) (ans [][]int) {
	temp := []int{}
	var dfs func(int)
	dfs = func(cur int) {
		// 剪枝：temp 长度加上区间 [cur, n] 的长度小于 k，不可能构造出长度为 k 的 temp
		if len(temp)+(n-cur+1) < k {
			return
		}
		// 记录合法的答案
		if len(temp) == k {
			comb := make([]int, k)
			copy(comb, temp)
			ans = append(ans, comb)
			return
		}
		// 考虑选择当前位置
		temp = append(temp, cur)
		dfs(cur + 1)
		temp = temp[:len(temp)-1]
		// 考虑不选择当前位置
		dfs(cur + 1)
	}
	dfs(1)
	return
}

//方法二：非递归（字典序法）实现组合型枚举-官方-不推荐
func combine2(n int, k int) (ans [][]int) {
	// 初始化
	// 将 temp 中 [0, k - 1] 每个位置 i 设置为 i + 1，即 [0, k - 1] 存 [1, k]
	// 末尾加一位 n + 1 作为哨兵
	temp := []int{}
	for i := 1; i <= k; i++ {
		temp = append(temp, i)
	}
	temp = append(temp, n+1)
	//fmt.Println(temp)

	for j := 0; j < k; {
		comb := make([]int, k)
		copy(comb, temp[:k])
		ans = append(ans, comb)
		// 寻找第一个 temp[j] + 1 != temp[j + 1] 的位置 t
		// 我们需要把 [0, t - 1] 区间内的每个位置重置成 [1, t]
		for j = 0; j < k && temp[j]+1 == temp[j+1]; j++ {
			temp[j] = j + 1
		}
		// j 是第一个 temp[j] + 1 != temp[j + 1] 的位置
		temp[j]++
	}
	return
}
