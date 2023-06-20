package main

import "fmt"

// LeetCode 904. 水果成篮

//你正在探访一家农场，农场从左到右种植了一排果树。这些树用一个整数数组 fruits 表示，其中 fruits[i]
//是第 i 棵树上的水果 种类 。

//你想要尽可能多地收集水果。然而，农场的主人设定了一些严格的规矩，你必须按照要求采摘水果：

//你只有 两个 篮子，并且每个篮子只能装 单一类型 的水果。每个篮子能够装的水果总量没有限制。
//你可以选择任意一棵树开始采摘，你必须从 每棵 树（包括开始采摘的树）上 恰好摘一个水果 。采摘的水果应当符合
//篮子中的水果类型。每采摘一次，你将会向右移动到下一棵树，并继续采摘。
//一旦你走到某棵树前，但水果不符合篮子的水果类型，那么就必须停止采摘。
//给你一个整数数组 fruits ，返回你可以收集的水果的 最大 数目。

//示例 1：
//输入：fruits = [1,2,1]
//输出：3
//解释：可以采摘全部 3 棵树。

// todo 待整理
//问题等价于，找到最长的子序列，最多含有两种“类型”（tree[i] 的值）。
//时间复杂度：O(N)，其中 N 是 tree 的长度。
//空间复杂度：O(N)。
func totalFruit2(fruits []int) int {
	i := 0
	counter := make(map[int]int)
	res := 0

	for j := 0; j < len(fruits); j++ {
		counter[fruits[j]] += 1 // 放入窗口
		// 超出窗口的限制，调整窗口
		for len(counter) > 2 {
			counter[fruits[i]] -= 1
			if counter[fruits[i]] == 0 {
				delete(counter, fruits[i])
			}
			// 窗口左边界移动
			i++
		}
		res = max(res, j-i+1)
	}
	//fmt.Println(counter)
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	arr := []int{1, 2, 1}
	fmt.Println(totalFruit2(arr))
}
