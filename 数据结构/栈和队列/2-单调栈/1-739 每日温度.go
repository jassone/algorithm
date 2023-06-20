package main

import "fmt"

// todo
//LeetCode 739. 每日温度

//请根据每日 气温 列表，重新生成一个列表。对应位置的输出为：要想观测到更高的气温，至少需要等待的天数。
// 如果气温在这之后都不会升高，请在该位置用 0 来代替。

//例如，给定一个列表 temperatures = [73, 74, 75, 71, 69, 72, 76, 73]，你的输出应该是 [1, 1, 4, 2, 1, 1, 0, 0]。

//提示：气温 列表长度的范围是 [1, 30000]。每个气温的值的均为华氏度，都是在 [30, 100] 范围内的整数。

// 卡尔

// 方法1：暴力
func dailyTemperatures1(t []int) []int {
	var res []int
	for i := 0; i < len(t)-1; i++ {
		j := i + 1
		for ; j < len(t); j++ {
			// 如果之后出现更高，说明找到了
			if t[j] > t[i] {
				res = append(res, j-i)
				break
			}
		}
		// 找完了都没找到
		if j == len(t) {
			res = append(res, 0)
		}
	}
	// 最后一个肯定是 0
	return append(res, 0)
}

// 单调递减栈
func dailyTemperatures2(num []int) []int {
	ans := make([]int, len(num))
	stack := []int{}
	for i, v := range num {
		// 栈不空，且当前遍历元素 v 破坏了栈的单调性
		for len(stack) != 0 && v > num[stack[len(stack)-1]] {
			// pop
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			ans[top] = i - top
		}
		stack = append(stack, i)
	}
	return ans
}

func main() {
	arr := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures1(arr))
	//fmt.Println(dailyTemperatures2(arr))
}
