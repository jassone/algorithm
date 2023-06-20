package main

import (
	"fmt"
	"math"
)

// LeetCode 76. 最小覆盖子串

//给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t
//所有字符的子串，则返回空字符串 "" 。

//注意：
//对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
//如果 s 中存在这样的子串，我们保证它是唯一的答案。

//示例 1：
//输入：s = "ADOBECODEBANC", t = "ABC"
//输出："BANC"

// todo 待整理
// 方法一：滑动窗口
func minWindow(s string, t string) string {
	// 滑动窗口, 遍历s要更新统计，检查当前窗口是否都含有，是则更新左边界，否则继续遍历
	src, cnt := map[byte]int{}, map[byte]int{}
	// 统计t中的元素及次数
	for i := range t {
		src[t[i]]++
	}
	sLen := len(s)
	lenWin := math.MaxInt32
	// 结果的左右边界，左闭右开
	ansL, ansR := -1, -1
	// handler, check函数，功能检查当前窗口中是否全有t中元素
	check := func() bool {
		for k, v := range src {
			// s中不仅要有t中元素，也要考虑次数
			if cnt[k] < v {
				return false
			}
		}
		return true
	}
	// 滑动窗口遍历
	for l, r := 0, 0; r < sLen; r++ {
		// 如果字符在t中出现，通过次数判定，计入cnt哈希表
		if src[s[r]] > 0 {
			cnt[s[r]]++
		}
		// 检查当前窗口是否全含t中元素，同时不断减小窗口左边界
		for check() && l <= r {
			// 取最小长度
			if lenWin > r-l+1 {
				lenWin = r - l + 1
				// 更新左右边界
				ansL, ansR = l, l+lenWin
			}
			// 考虑进一步缩减窗口
			// 左边界的s[l]字符在src中含有，则窗口hash表中减去count
			// 否则，略过，即该字符不是src字符，直接更新左边界
			if _, ok := src[s[l]]; ok {
				cnt[s[l]]--
			}
			// 窗口左边界右移
			l++
		}
	}
	// 上述循环结束，如果找到合适子串，ansL/ansR一定更新了
	if ansL == -1 {
		return ""
	}
	// 左闭右开
	return s[ansL:ansR]
}

func main() {
	s1 := "ADOBECODEBANC"
	s2 := "ABC"
	fmt.Println(minWindow(s1, s2))
}
