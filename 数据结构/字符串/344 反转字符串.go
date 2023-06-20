package main

//LeetCode 344.反转字符串
//编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 char[] 的形式给出。

//不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。

//你可以假设数组中的所有字符都是 ASCII 码表中的可打印字符。

//示例 1：
//输入：["h","e","l","l","o"]
//输出：["o","l","l","e","h"]

//方法一：双指针-官方
//时间复杂度：O(N)，其中 NN 为字符数组的长度。一共执行了 N/2N/2 次的交换。
//空间复杂度：O(1)。只使用了常数空间来存放若干变量。
func reverseString(s []byte) {
	for left, right := 0, len(s)-1; left < right; left++ {
		s[left], s[right] = s[right], s[left]
		right--
	}
}

// 卡尔版-更清晰,推荐
func reverseString2(s []byte) {
	left := 0
	right := len(s) - 1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

func main() {

}
