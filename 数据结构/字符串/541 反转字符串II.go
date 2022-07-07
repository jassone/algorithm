package main

// LeetCode 541. 反转字符串II
//给定一个字符串 s 和一个整数 k，你需要对从字符串开头算起的每隔 2k 个字符的前 k 个字符进行反转。

//如果剩余字符少于 k 个，则将剩余字符全部反转。

//如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样。

//示例:
//输入: s = "abcdefg", k = 2
//输出: "bacdfeg"

// 方法一：模拟-官方
//我们直接按题意进行模拟：反转每个下标从 2k 的倍数开始的，长度为 k 的子串。若该子串长度不足 k，
// 则反转整个子串。

//时间复杂度：O(n)，其中 n 是字符串 s 的长度。
//空间复杂度：O(1) 或 O(n)，取决于使用的语言中字符串类型的性质。如果字符串是可修改的，那么我们可以直接在
// 原字符串上修改，空间复杂度为 O(1)，否则需要使用 O(n) 的空间将字符串临时转换为可以修改的数据结构（例如数组），
// 空间复杂度为 O(n)。
func reverseStr1(s string, k int) string {
	t := []byte(s)
	for i := 0; i < len(s); i += 2 * k {
		sub := t[i:min(i+k, len(s))]
		for j, n := 0, len(sub); j < n/2; j++ {
			sub[j], sub[n-1-j] = sub[n-1-j], sub[j]
		}
	}
	return string(t)
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 卡尔解法-更清晰
func reverseStr2(s string, k int) string {
	ss := []byte(s)
	length := len(s)
	for i := 0; i < length; i += 2 * k {
		if i+k <= length {
			reverse(ss[i : i+k])
		} else {
			reverse(ss[i:length])
		}
	}
	return string(ss)
}
func reverse(b []byte) {
	left := 0
	right := len(b) - 1
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
}
func main() {

}
