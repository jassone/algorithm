package main

import "fmt"

//LeetCode 剑指Offer58-II.左旋转字符串
//字符串的左旋转操作是把字符串前面的若干个字符转移到字符串的尾部。请定义一个函数实现字符串左旋转操作的功能。
// 比如，输入字符串"abcdefg"和数字2，该函数将返回左旋转两位得到的结果"cdefgab"。

//示例 1：
//输入: s = "abcdefg", k = 2
//输出: "cdefgab"

//示例 2：
//输入: s = "lrloseumgh", k = 6
//输出: "umghlrlose"

//限制：
//1 <= k < s.length <= 10000

// 方法1：截取
func reverseLeftWords1(s string, n int) string {
	for i := 0; i < n; i++ {
		s += string(s[i])
	}
	s = s[n:]
	return s
}

// 升级：不能申请额外空间，只能在本串上操作。
//具体步骤为：
// 1 反转区间为前n的子串
// 2 反转区间为n到末尾的子串
// 3 反转整个字符串

// 关键点：本题是先局部反转再 整体反转
func reverseLeftWords2(s string, n int) string {
	b := []byte(s)
	// 1. 反转前n个字符
	// 2. 反转第n到end字符
	// 3. 反转整个字符
	reverse1(b, 0, n-1)
	reverse1(b, n, len(b)-1)
	reverse1(b, 0, len(b)-1)
	return string(b)
}

// 切片是引用传递
func reverse1(b []byte, left, right int) {
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
}

func main() {
	n := 2
	str := "aberc"
	fmt.Println(reverseLeftWords1(str, n))
	fmt.Println(reverseLeftWords2(str, n))
}
