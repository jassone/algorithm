package main

import (
	"fmt"
	"sort"
)

//LeetCode 455.分发饼干

//假设你是一位很棒的家长，想要给你的孩子们一些小饼干。但是，每个孩子最多只能给一块饼干。

//对每个孩子 i，都有一个胃口值 g[i]，这是能让孩子们满足胃口的饼干的最小尺寸；并且每块饼干 j，
// 都有一个尺寸 s[j] 。如果 s[j] >= g[i]，我们可以将这个饼干 j 分配给孩子 i ，这个孩子会得到满足。
// 你的目标是尽可能满足越多数量的孩子，并输出这个最大数值。

//示例 1:
//输入: g = [1,2,3], s = [1,1]
//输出: 1 解释:你有三个孩子和两块小饼干，3个孩子的胃口值分别是：1,2,3。虽然你有两块小饼干，
// 由于他们的尺寸都是1，你只能让胃口值是1的孩子满足。所以你应该输出1。

// 卡尔
//思路
//为了满足更多的小孩，就不要造成饼干尺寸的浪费。

//***大尺寸的饼干既可以满足胃口大的孩子也可以满足胃口小的孩子，那么就应该优先满足胃口大的。***
//这里的局部最优就是大饼干喂给胃口大的，充分利用饼干尺寸喂饱一个，全局最优就是喂饱尽可能多的小孩。

//可以尝试使用贪心策略，先将饼干数组和小孩数组排序。
//然后从后向前遍历小孩数组，用大饼干优先满足胃口大的，并统计满足小孩数量。

// 也可以换一个思路，小饼干先喂饱小胃口

//排序后，局部最优
//时间复杂度：O(nlogn)，就是快排O(nlogn)，遍历O(n)，加一起就是还是O(nlogn)。
//空间复杂度：O(1)
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	// 从小到大
	child := 0
	for sIdx := 0; child < len(g) && sIdx < len(s); sIdx++ {
		if s[sIdx] >= g[child] {
			//如果饼干的大小大于或等于孩子的为空则给与，否则不给予，继续寻找选一个饼干是否符合
			child++
		}
	}
	return child

	// 从大到小
	//count := 0
	//child := len(g) - 1
	//for sIdx := len(s) - 1; child >= 0 && sIdx >= 0; sIdx-- {
	//	if s[sIdx] >= g[child] {
	//		//如果饼干的大小大于或等于孩子的为空则给与，否则不给予，继续寻找选一个饼干是否符合
	//		child--
	//		count++
	//	}
	//}
	//
	//return count
}

func main() {
	g := []int{1, 2}
	s := []int{1, 2, 3}
	fmt.Println(findContentChildren(g, s))
}

// 官方
// https://leetcode.cn/problems/assign-cookies/solution/fen-fa-bing-gan-by-leetcode-solution-50se/
func findContentChildren22(greed, size []int) (ans int) {
	sort.Ints(greed)
	sort.Ints(size)
	n, m := len(greed), len(size)
	for i, j := 0, 0; i < n && j < m; i++ {
		for j < m && greed[i] > size[j] {
			j++
		}
		if j < m {
			ans++
			j++
		}
	}
	return
}
