package main

import (
	"fmt"
)

// LeetCode 3. 无重复字符的最长子串

// 给我们一个字符串，让我们求最长的无重复字符的子串，
// 注意这里是子串，不是子序列，所以必须是连续的。

// 方法1：暴力算法，两个for循环

// 方法2，滑动窗口法-推荐
//例如给定数组[2,2,3,4,8,99,3]，窗口大小为3，求出每个窗口的元素和就是
//固定大小窗口的问题，如果求数组[2,2,3,4,8,99,3] 的最长连续子数组就是窗口可变的问题。
//使用滑动窗口，我们可以有效减低算法的时间复杂度。

// 使用滑动窗口求解问题时，主要需要注意：
// 1 什么条件下移动窗口的起始位置。
// 2 何时动态的扩展窗口。
func lengthOfLongestSubstring(s string) int {
	var leftIndex, maxCount, currentMaxCount int
	usedMap := make(map[rune]int)
	for index, value := range []rune(s) {
		if usedIndex, ok := usedMap[value]; ok {
			leftIndex = usedIndex + 1
			// 一旦发现重复，则窗口从后面开始，前面的都抛弃掉
			//fmt.Println(index,string(value), usedIndex,leftIndex)
		} else {
			currentMaxCount = index - leftIndex + 1
			maxCount = max(maxCount, currentMaxCount)
		}

		usedMap[value] = index // 注意这里会覆盖前面的key，即起到了删除之前老数据的作用
	}

	return maxCount
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// 滑动窗口-官方 todo
// 时间复杂度：O(N)，其中 NN 是字符串的长度。左指针和右指针分别会遍历整个字符串一次。
// 空间复杂度：O(∣Σ∣)，其中 Σ 表示字符集（即字符串中可以出现的字符），∣Σ∣ 表示字符集的大小。
// 在本题中没有明确说明字符集，因此可以默认为所有 ASCII 码在 [0, 128)[0,128) 内的字符，
// 即 ∣Σ∣=128。我们需要用到哈希集合来存储出现过的字符，而字符最多有  ∣Σ∣ 个，因此空间复杂度为
// O(∣Σ∣)。
func lengthOfLongestSubstring2(s string) int {
	// 哈希集合，记录每个字符是否出现过
	m := map[byte]int{}
	n := len(s)
	// 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			// 左指针向右移动一格，移除一个字符
			delete(m, s[i-1])
		}
		for rk+1 < n && m[s[rk+1]] == 0 {
			// 不断地移动右指针
			m[s[rk+1]]++
			rk++
		}
		// 第 i 到 rk 个字符是一个极长的无重复字符子串
		ans = max(ans, rk-i+1)
	}
	return ans
}

func main() {
	var t string = "12312"
	fmt.Println(lengthOfLongestSubstring(t))
	fmt.Println(lengthOfLongestSubstring2(t))
}
