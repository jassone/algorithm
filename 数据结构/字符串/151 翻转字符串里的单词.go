package main

import "fmt"

// todo
//LeetCode 151.翻转字符串里的单词

//给定一个字符串，逐个翻转字符串中的每个单词。

//示例 1：
//输入: "the sky is blue"
//输出: "blue is sky the"

//示例 2：
//输入: "  hello world!  "
//输出: "world! hello"
//解释: 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。

//示例 3：
//输入: "a good   example"
//输出: "example good a"
//解释: 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。

// 方法一：自行编写对应的函数

// 方法二：卡尔版-双指针
//思路如下：
// 移除多余空格
// 将整个字符串反转
// 将每个单词反转

// 关键点：本题是先整体反转 再 局部反转

//时间复杂度：O(n),其中 n 为输入字符串的长度。
//空间复杂度为O(1)。
func reverseWords2(s string) string {
	//1.使用双指针删除冗余的空格
	slowIndex, fastIndex := 0, 0
	b := []byte(s)
	//删除头部冗余空格
	for len(b) > 0 && fastIndex < len(b) && b[fastIndex] == ' ' {
		fastIndex++
	}

	//删除单词间冗余空格
	for ; fastIndex < len(b); fastIndex++ {
		if fastIndex-1 > 0 && b[fastIndex-1] == b[fastIndex] && b[fastIndex] == ' ' {
			continue
		}
		b[slowIndex] = b[fastIndex]
		slowIndex++
	}

	//删除尾部冗余空格
	if slowIndex-1 > 0 && b[slowIndex-1] == ' ' {
		b = b[:slowIndex-1]
	} else {
		b = b[:slowIndex] // 截止下，后面即使有字符也不要了
	}

	//2.反转整个字符串
	reverse(b, 0, len(b)-1)
	//3.反转单个单词  i单词开始位置，j单词结束位置
	i := 0
	for i < len(b) {
		j := i
		for ; j < len(b) && b[j] != ' '; j++ {
		}
		reverse(b, i, j-1)
		i = j
		i++
	}
	return string(b)
}
func reverse(b []byte, left, right int) {
	for left < right {
		(b)[left], (b)[right] = (b)[right], (b)[left]
		left++
		right--
	}
}

// 方法三：双端队列
//思路和算法
//由于双端队列支持从队列头部插入的方法，因此我们可以沿着字符串一个一个单词处理，然后将单词压入队列的头部，再将队列转成字符串即可。

func main() {
	str := "the  sky is  blue  "
	fmt.Println(reverseWords2(str))
	fmt.Println(reverseWords33(str))

}

// 方法：找空格，
//找空格，而不是找字符串，两空格中间的可能是单词。
func reverseWords33(s string) (res string) {
	s = " " + s + " "
	l, r := len(s)-1, len(s)-1
	for i := len(s) - 2; i >= 0; i-- {
		if s[i] == ' ' {
			l, r = i, l
			if r > l+1 {
				res = res + s[l+1:r] + " "
			}
		}
	}
	return res[:len(res)-1]
}
