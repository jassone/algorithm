package main

import "fmt"

// todo
//LeetCode 28. 实现 strStr()

//实现 strStr() 函数。即从一个字符串中找到另外一个字符串的起始位置

//给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置
// (从0开始)。如果不存在，则返回  -1。

//示例 1: 输入: haystack = "hello", needle = "ll" 输出: 2
//示例 2: 输入: haystack = "aaaaa", needle = "bba" 输出: -1

//说明: 当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。 对于本题而言，
// 当 needle 是空字符串时我们应当返回 0 。这与C语言的 strstr() 以及 Java的 indexOf() 定义相符。

// 方法1：暴力-官方
//时间复杂度：O(n×m)，其中 n 是字符串 haystack 的长度，m 是字符串 needle 的长度。
// 最坏情况下我们需要将字符串 needle 与字符串 haystack 的所有长度为 m 的子串均匹配一次。
//空间复杂度：O(1)。我们只需要常数的空间保存若干变量。
func strStr1(haystack, needle string) int {
	n, m := len(haystack), len(needle)
outer:
	for i := 0; i+m <= n; i++ {
		for j := range needle {
			if haystack[i+j] != needle[j] {
				continue outer
			}
		}
		return i
	}
	return -1
}

// 方法2：KMP-官方
//时间复杂度：O(n+m)，其中 n 是字符串 haystack 的长度，mm 是字符串 needle 的长度。我们至多需要遍历两字符串一次。
//空间复杂度：O(m)，其中 m 是字符串 needle 的长度。我们只需要保存字符串 needle 的前缀函数。

// 1 前缀表无减一或者右移
func strStr3(haystack string, needle string) int {
	n := len(needle)
	if n == 0 {
		return 0
	}
	j := 0
	next := make([]int, n)
	getNext3(next, needle)
	for i := 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1] // 回退到j的前一位
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == n {
			return i - n + 1
		}
	}
	return -1
}

// 方法二: 前缀表无减一或者右移
//构造next数组其实就是计算模式串s，前缀表的过程。 主要有如下三步：
// 1 初始化
// 2 处理前后缀不相同的情况
// 3 处理前后缀相同的情况
// 4 更新next数组的值
// getNext 构造前缀表next
// params:
//		  next 前缀表数组
//		  s 模式串
func getNext3(next []int, s string) {
	j := 0 // j表示 最长相等前后缀长度，且指向前缀末尾位置
	next[0] = j

	// i指向后缀末尾位置
	for i := 1; i < len(s); i++ {
		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		next[i] = j
	}
	fmt.Println(next)
	// [0 1 0 1 2 0]
}

// 另外一种：前缀表使用减1实现-卡尔版
func strStr2(haystack string, needle string) int {
	n := len(needle)
	if n == 0 {
		return 0
	}
	next := make([]int, n)
	getNext2(next, needle)
	j := -1 // 模式串的起始位置 next为-1 因此也为-1
	for i := 0; i < len(haystack); i++ {
		for j >= 0 && haystack[i] != needle[j+1] {
			j = next[j] // 寻找下一个匹配点
		}
		if haystack[i] == needle[j+1] {
			j++
		}
		if j == n-1 { // j指向了模式串的末尾
			return i - n + 1
		}
	}
	return -1
}

// getNext 构造前缀表next
// params:
//		  next 前缀表数组
//		  s 模式串
func getNext2(next []int, s string) {
	j := -1 // j表示 最长相等前后缀长度，且指向前缀末尾位置
	next[0] = j

	// i指向后缀末尾位置
	for i := 1; i < len(s); i++ {
		for j >= 0 && s[i] != s[j+1] {
			j = next[j] // 回退前一位
		}
		if s[i] == s[j+1] {
			j++
		}
		next[i] = j // next[i]是i（包括i）之前的最长相等前后缀长度
	}
	fmt.Println(next)
	//[-1 0 -1 0 1 -1]
}

func main() {
	str := "aabaabaafa"
	needle := "aabaaf"
	fmt.Println(strStr1(str, needle))
	fmt.Println(strStr3(str, needle))
	fmt.Println(strStr2(str, needle))
}
