package main

import (
	"fmt"
	"sort"
)

//LeetCode 242.有效的字母异位词

//字母异位词:字符的个数和字符串长度一样

//给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

//示例 1: 输入: s = "anagram", t = "nagaram" 输出: true
//示例 2: 输入: s = "rat", t = "car" 输出: false
//说明: 你可以假设字符串只包含小写字母。
// 方法1：暴力解法
// 两层for循环，同时还要记录字符是否重复出现，很明显时间复杂度是 O(n^2)。

// 方法二：排序-官方,推荐1
//t 是 s 的异位词等价于「两个字符串排序后相等」。因此我们可以对字符串 s 和 t 分别排序，看排序后的字符串
// 是否相等即可判断。此外，如果 s 和 t 的长度不同，t 必然不是 s 的异位词。

// 时间复杂度：O(nlogn)，其中 n 为 s 的长度。排序的时间复杂度为 O(nlogn)，比较两个字符串是否相等时间复杂度为
// O(n)，因此总体时间复杂度为 O(nlogn+n)=O(nlogn)。
// 空间复杂度：O(logn)。排序需要 O(logn) 的空间复杂度。注意，在某些语言（比如 Java & JavaScript）中字符串是不可变的，
// 因此我们需要额外的 O(n) 的空间来拷贝字符串。但是我们忽略这一复杂度分析，因为：
// 这依赖于语言的细节；
// 这取决于函数的设计方式，例如，可以将函数参数类型更改为 char[]。
func isAnagram2(s, t string) bool {
	s1, s2 := []byte(s), []byte(t)
	sort.Slice(s1, func(i, j int) bool { return s1[i] < s1[j] })
	sort.Slice(s2, func(i, j int) bool { return s2[i] < s2[j] })
	return string(s1) == string(s2)
}

// 方法3：hash法-官方,推荐2
// 时间复杂度：O(n)，其中 n 为 s 的长度。
// 空间复杂度：O(S)，其中 S 为字符集大小，此处 S=26。
func isAnagram3(s, t string) bool {
	var c1, c2 [26]int
	for _, ch := range s {
		c1[ch-'a']++ // 字符相减等于ascill码值相减
	}
	for _, ch := range t {
		c2[ch-'a']++
	}
	//fmt.Println(c1, c2)
	//[3 0 0 0 0 0 1 0 0 0 0 0 1 1 0 0 0 1 0 0 0 0 0 0 0 0]
	return c1 == c2
}

// 进阶-解决字符多字节问题
func isAnagram33(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	cnt := map[rune]int{}
	for _, ch := range s {
		cnt[ch]++
	}
	for _, ch := range t {
		cnt[ch]--
		if cnt[ch] < 0 {
			return false
		}
	}
	return true
}

func main() {
	s := "anagram"
	t := "nagaram"
	fmt.Println(isAnagram3(s, t))

}
