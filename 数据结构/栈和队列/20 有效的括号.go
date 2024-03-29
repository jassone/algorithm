package main

import "fmt"

//LeetCode 20. 有效的括号

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

//有效字符串需满足：
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 注意空字符串可被认为是有效字符串。

//示例 1:

//输入: "()"
//输出: true

//示例 4:
//
//输入: "([)]"
//输出: false

//题外话

//括号匹配是使用栈解决的经典问题。

//题意其实就像我们在写代码的过程中，要求括号的顺序是一样的，有左括号，相应的位置必须要有右括号。

//如果还记得编译原理的话，编译器在 词法分析的过程中处理括号、花括号等这个符号的逻辑，也是使用了栈这种数据结构。

//再举个例子，linux系统中，cd这个进入目录的命令我们应该再熟悉不过了。
//cd a/b/c/../../
//这个命令最后进入a目录，系统是如何知道进入了a目录呢 ，这就是栈的应用

//所以栈在计算机领域中应用是非常广泛的。

// *********由于栈结构的特殊性，非常适合做对称匹配类的题目。**************

// 思路
//先来分析一下 这里有三种不匹配的情况：

// 第一种情况，字符串里左方向的括号多余了 ，所以不匹配。
//  ({[]}()
// 第二种情况，括号没有多余，但是 括号的类型没有匹配上。 括号匹配
//  {{]}()
// 第三种情况，字符串里右方向的括号多余了，所以不匹配。
//  {[]}()}
// 动画： https://pic.imgdb.cn/item/6282726209475431290e8c2c.gif

// 方法1：栈，推荐1
//时间复杂度：O(n)，其中 n 是字符串 s 的长度。
//空间复杂度：O(n+∣Σ∣)，其中 Σ 表示字符集，本题中字符串只包含 6 种括号，∣Σ∣=6。
// 栈中的字符数量为 O(n)，而哈希表使用的空间为 O(∣Σ∣)，相加即可得到总空间复杂度。
func isValid(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []byte{}
	for i := 0; i < n; i++ {
		if pairs[s[i]] > 0 {
			// 比如 }[] 这时候就是 len(stack) == 0
			// 第二个判断:判断是否有左括号
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	// 这里判断是否都匹配完了
	return len(stack) == 0
}
// 方法2：前后循环对比法，推荐2
// 如果符合条件，则字符串长度是偶数，且整个字符串呈现一种对称性
func isValid2(s string) bool{
	len := len(s)
	if len == 0 || len %2 != 0 {
		return false
	}
	pairs := map[byte]byte{
		'{':'}',
		'[':']',
		'(':')',
	}
	for i:=0;i<len/2;i++{
		v1,ok1:=pairs[s[i]]
		if ok1 && v1 == s[len-1-i]{
			// todo
		} else{
			return false
		}
	}

	return true
}

func main() {
	s := "(({[]}))"
	fmt.Println(isValid(s))
	fmt.Println(isValid2(s))
}
