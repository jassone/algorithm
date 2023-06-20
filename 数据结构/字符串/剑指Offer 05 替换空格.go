package main

import "fmt"

//LeetCode 剑指Offer 05.替换空格

//请实现一个函数，把字符串 s 中的每个空格替换成"%20"。

//示例 1： 输入：s = "We are happy."
//输出："We%20are%20happy."

//方法一：字符数组-官方，推荐
// 由于每次替换从 1 个字符变成 3 个字符，使用字符数组可方便地进行替换。建立字符数组地长度为 s 的长度的 3 倍，
// 这样可保证字符数组可以容纳所有替换后的字符。

//时间复杂度：O(n)。遍历字符串 s 一遍。
//空间复杂度：O(n)。额外创建字符数组，长度为 s 的长度的 3 倍。
func replaceSpace1(s string) string {
	length := len(s)
	res := make([]byte, length)
	for _, v := range []rune(s) { // 这里换成[]byte(s)也可以,因为多字节字符不会出现空字符
		if v == ' ' {
			res = append(res, []byte("%20")...)
		} else {
			res = append(res, byte(v))
		}
	}

	return string(res)
}

// 方法2：双指针-卡尔版
// 思路
// 首先扩充数组到每个空格替换成"%20"之后的大小。

// 问题 ：有同学问了，为什么要从后向前填充，从前向后填充不行么？
// 从前向后填充就是O(n^2)的算法了，因为每次添加元素都要将添加元素之后的所有元素向后移动。

// 其实很多数组填充类的问题，都可以先预先给数组扩容带填充后的大小，然后在从后向前进行操作。
// 这么做有两个好处：
// 1 不用申请新数组。
// 2 从后向前填充元素，避免了从前先后填充元素要来的 每次添加元素都要将添加元素之后的所有元素向后移动。

//然后从后向前替换空格，也就是双指针法，
//时间复杂度：O(n)
//空间复杂度：O(1)-原地修改
func replaceSpace2(s string) string {
	b := []byte(s)
	length := len(b)
	spaceCount := 0
	// 计算空格数量
	for _, v := range b {
		if v == ' ' {
			spaceCount++
		}
	}
	// 扩展原有切片
	resizeCount := spaceCount * 2
	tmp := make([]byte, resizeCount)
	b = append(b, tmp...)
	i := length - 1 // 原字符串长度
	j := len(b) - 1 // 扩容后的字符串长度
	for i >= 0 {
		if b[i] != ' ' {
			b[j] = b[i]
			i--
			j--
		} else {
			b[j] = '0'
			b[j-1] = '2'
			b[j-2] = '%'
			i--
			j = j - 3
		}
	}
	return string(b)
}

func main() {
	str := "adf rtr tyy"
	fmt.Println(replaceSpace1(str))
	//fmt.Println(replaceSpace2(str))
}
