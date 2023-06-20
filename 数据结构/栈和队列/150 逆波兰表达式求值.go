package main

import (
	"fmt"
	"strconv"
)

//LeetCode 150. 逆波兰表达式求值

//根据 逆波兰表示法，求表达式的值。

//有效的运算符包括 + ,  - ,  * ,  / 。每个运算对象可以是整数，也可以是另一个逆波兰表达式。

//说明：
//整数除法只保留整数部分。 给定逆波兰表达式总是有效的。换句话说，表达式总会得出有效数值且不存在除数为 0 的情况。

//示例 1：
//输入: ["2", "1", "+", "3", " * "]
//输出: 9
//解释: 该算式转化为常见的中缀算术表达式为：((2 + 1) * 3) = 9

//说明：
//逆波兰表达式：是一种后缀表达式，所谓后缀就是指算符写在后面。

//平常使用的算式则是一种中缀表达式，如 ( 1 + 2 ) * ( 3 + 4 ) 。

//该算式的逆波兰表达式写法为 ( ( 1 2 + ) ( 3 4 + ) * ) 。

//逆波兰表达式主要有以下两个优点：
// 去掉括号后表达式无歧义，上式即便写成 1 2 + 3 4 + * 也可以依据次序计算出正确结果。
// 适合用栈操作运算：遇到数字则入栈；遇到算符则取出栈顶两个数字进行计算，并将结果压入栈中。

//题外话
//我们习惯看到的表达式都是中缀表达式，因为符合我们的习惯，但是中缀表达式对于计算机来说就不是很友好了。
//例如：4 + 13 / 5，这就是中缀表达式，计算机从左到右去扫描的话，扫到13，还要判断13后面是什么运算法，
//还要比较一下优先级，然后13还和后面的5做运算，做完运算之后，还要向前回退到 4 的位置，继续做加法，你说麻不麻烦！
//那么将中缀表达式，转化为后缀表达式之后：["4", "13", "5", "/", "+"] ，就不一样了，计算机可以利用栈里顺序处理，
//不需要考虑优先级了。也不用回退了， 所以后缀表达式对计算机来说是非常友好的。
//****可以说本题不仅仅是一道好题，也展现出计算机的思考方式。*****
//在1970年代和1980年代，惠普在其所有台式和手持式计算器中都使用了RPN（后缀表达式），直到2020年代仍在某些模型中使用了RPN。

// *****分析：其实逆波兰表达式相当于是二叉树中的后序遍历。******

// 方法1：栈-官方,推荐
//时间复杂度：O(n)，其中 n 是数组 tokens 的长度。需要遍历数组 tokens 一次，计算逆波兰表达式的值。
//空间复杂度：O(n)，其中 n 是数组 tokens 的长度。使用栈存储计算过程中的数，栈内元素个数不会超过逆波兰表达式的长度。
func evalRPN1(tokens []string) int {
	stack := []int{}
	for _, token := range tokens {
		val, err := strconv.Atoi(token)
		if err == nil { // 如果正确将字符转换为数字，则压入栈
			stack = append(stack, val)
		} else {
			num1, num2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			switch token {
			case "+":
				stack = append(stack, num1+num2)
			case "-":
				stack = append(stack, num1-num2)
			case "*":
				stack = append(stack, num1*num2)
			default:
				stack = append(stack, num1/num2)
			}
		}
	}
	return stack[0]
}

// 方法2：-数组-官方,推荐2
// 其实是使用一个数组模拟栈操作。

// 对于一个有效的逆波兰表达式，其长度 n 一定是奇数，且操作数的个数一定比运算符的个数多 1 个，即包含 (n+1) / 2 个操作数。

//时间复杂度：O(n)，其中 n 是数组 tokens 的长度。需要遍历数组 tokens 一次，计算逆波兰表达式的值。
//空间复杂度：O(n)，其中 n 是数组 tokens 的长度。需要创建长度为 (n+1) / 2 个的数组模拟栈操作。
func evalRPN2(tokens []string) int {
	stack := make([]int, (len(tokens)+1)/2)
	index := -1
	for _, token := range tokens {
		val, err := strconv.Atoi(token)
		if err == nil {
			index++
			stack[index] = val
		} else {
			index--
			switch token {
			case "+":
				stack[index] += stack[index+1]
			case "-":
				stack[index] -= stack[index+1]
			case "*":
				stack[index] *= stack[index+1]
			default:
				stack[index] /= stack[index+1]
			}
		}
	}
	return stack[0]
}

func main() {
	slice := []string{ "2", "1", "+", "3", "*"}

	//fmt.Println(evalRPN1(slice))
	fmt.Println(evalRPN2(slice))
}
