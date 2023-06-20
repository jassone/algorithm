package main

import (
	"fmt"
)

//LeetCode 131.分割回文串

// 给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。

//返回 s 所有可能的分割方案。

//示例: 输入: "aab" 输出: [ ["aa","b"], ["a","a","b"] ]

// 卡尔
// https://www.programmercarl.com/0131.%E5%88%86%E5%89%B2%E5%9B%9E%E6%96%87%E4%B8%B2.html
//本题这涉及到的难点：
// 切割问题可以抽象为组合问题
// 如何模拟那些切割线
// 切割问题中递归如何终止
// 在递归循环中如何截取子串
// 如何判断回文

//一些同学可能想不清楚 回溯究竟是如何切割字符串呢？
//我们来分析一下切割，****其实切割问题类似组合问题。*****

func partition(s string) [][]string {
	var tmpString []string //切割字符串集合
	var res [][]string     //结果集合
	backTracking(s, tmpString, 0, &res)

	return res
}
func backTracking(s string, tmpString []string, startIndex int, res *[][]string) {
	if startIndex == len(s) { //到达字符串末尾了
		//进行一次切片拷贝，怕之后的操作影响tmpString切片内的值
		t := make([]string, len(tmpString))
		copy(t, tmpString)
		*res = append(*res, t)
	}

	for i := startIndex; i < len(s); i++ {
		//处理（首先通过startIndex和i判断切割的区间，进而判断该区间的字符串是否为回文，
		// 若为回文，则加入到tmpString，否则继续后移，找到回文区间）（这里为一层处理）
		if isPartition(s, startIndex, i) {
			tmpString = append(tmpString, s[startIndex:i+1])
		} else {
			continue
		}
		//递归
		backTracking(s, tmpString, i+1, res)
		//回溯
		tmpString = tmpString[:len(tmpString)-1]
	}
}

//判断是否为回文
func isPartition(s string, startIndex, end int) bool {
	left := startIndex
	right := end
	for left < right {
		if s[left] != s[right] {
			return false
		}
		//移动左右指针
		left++
		right--
	}
	return true
}

func main() {
	str := "aab"
	fmt.Println(partition(str))
}

// 官方太复杂-略
//https://leetcode.cn/problems/palindrome-partitioning/solution/fen-ge-hui-wen-chuan-by-leetcode-solutio-6jkv/
