package main

import "fmt"

//LeetCode 844. 比较含退格的字符串
//给定 s 和 t 两个字符串，当它们分别被输入到空白的文本编辑器后，如果两者相等，返回 true 。
//# 代表退格字符。

//注意：如果对空文本输入退格字符，文本继续为空。

//示例 1：
//输入：s = "ab#c", t = "ad#c"
//输出：true
//解释：s 和 t 都会变成 "ac"。

//示例 2：
//输入：s = "ab##", t = "c#d#"
//输出：true
//解释：s 和 t 都会变成 ""。

//方法一：重构字符串-栈-官方
//思路及算法
//最容易想到的方法是将给定的字符串中的退格符和应当被删除的字符都去除，还原给定字符串的一般形式。
//然后直接比较两字符串是否相等即可。

//具体地，我们用栈处理遍历过程，每次我们遍历到一个字符：
//  如果它是退格符，那么我们将栈顶弹出；
//  如果它是普通字符，那么我们将其压入栈中。
//时间复杂度：O(N+M)，其中 NN 和 MM 分别为字符串 SS 和 TT 的长度。我们需要遍历两字符串各一次。
//空间复杂度：O(N+M)，其中 NN 和 MM 分别为字符串 SS 和 TT 的长度。主要为还原出的字符串的开销。
func f1(s, t string) bool {
	return build(s) == build(t)
}
func build(str string) string {
	s := []byte{}
	for i := range str {
		if str[i] != '#' {
			s = append(s, str[i])
		} else if len(s) > 0 {
			s = s[:len(s)-1]
		}
	}
	return string(s)
}

// 方法2：双指针法-官方，背下来
//思路及算法

//一个字符是否会被删掉，只取决于该字符后面的退格符，而与该字符前面的退格符无关。
//因此当我们逆序地遍历字符串，就可以立即确定当前字符是否会被删掉。

//具体地，我们定义 skip 表示当前待删除的字符的数量。每次我们遍历到一个字符：
// 1 若该字符为退格符，则我们需要多删除一个普通字符，我们让 skip 加 1；

// 2 若该字符为普通字符：
//   若 skip 为 00，则说明当前字符不需要删去；
//   若 skip 不为 00，则说明当前字符需要删去，我们让 skip 减 1。

//这样，我们定义两个指针，分别指向两字符串的末尾。每次我们让两指针逆序地遍历两字符串，
//直到两字符串能够各自确定一个字符，然后将这两个字符进行比较。重复这一过程直到找到的两个字符不相等，
//或遍历完字符串为止。

//时间复杂度：O(N+M)，其中 NN 和 MM 分别为字符串 SS 和 TT 的长度。我们需要遍历两字符串各一次。
//空间复杂度：O(1)。对于每个字符串，我们只需要定义一个指针和一个计数器即可。
func f2(s, t string) bool {
	skipS, skipT := 0, 0
	i, j := len(s)-1, len(t)-1
	for i >= 0 || j >= 0 {
		for i >= 0 {
			if s[i] == '#' {
				skipS++
				i--
			} else if skipS > 0 {
				skipS--
				i--
			} else {
				break
			}
		}
		for j >= 0 {
			if t[j] == '#' {
				skipT++
				j--
			} else if skipT > 0 {
				skipT--
				j--
			} else {
				break
			}
		}
		if i >= 0 && j >= 0 {
			if s[i] != t[j] {
				return false
			}
		} else if i >= 0 || j >= 0 {
			return false
		}
		i--
		j--
	}

	return true
}

func main() {
	s1 := "ab#c"
	s2 := "ad#c"
	fmt.Println(f1(s1, s2))
	fmt.Println(f2(s1, s2))
}
