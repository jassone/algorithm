package main

//LeetCode 202 快乐数
//编写一个算法来判断一个数 n 是不是快乐数。

//「快乐数」定义为：对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和，然后重复这个过程直到这个数变为 1，
// 也可能是 无限循环 但始终变不到 1。如果 可以变为 1，那么这个数就是快乐数。
// 如果 n 是快乐数就返回 True ；不是，则返回 False 。

//输入：19
//输出：true
//解释：
//1^2 + 9^2 = 82
//8^2 + 2^2 = 68
//6^2 + 8^2 = 100
//1^2 + 0^2 + 0^2 = 1

//题目中说了会 无限循环，那么也就是说求和的过程中，sum会重复出现，这对解题很重要！
//当我们遇到了要快速判断一个元素是否出现集合里的时候，就要考虑哈希法了。
//所以这道题目使用哈希法，来判断这个sum是否重复出现，如果重复了就是return false， 否则一直找到sum为1为止。

//方法一：用哈希集合检测循环-官方， 简单好理解
// 详细解读，参考官方：https://leetcode-cn.com/problems/happy-number/solution/kuai-le-shu-by-leetcode-solution/
func isHappy1(n int) bool {
	m := map[int]bool{}
	// !m[n]防止进入死循环
	for ; n != 1 && !m[n]; n, m[n] = step(n), true {
	}
	return n == 1
}
func step(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	return sum
}

// 方法二：快慢指针法-官方
// 如果不是快乐数，则肯定会进入循环，这样快慢指针就能在循环中碰到
// 详细解读，参考官方：https://leetcode-cn.com/problems/happy-number/solution/kuai-le-shu-by-leetcode-solution/
func isHappy2(n int) bool {
	slow, fast := n, step(n)
	for fast != 1 && slow != fast {
		slow = step(slow)
		fast = step(step(fast))
	}
	return fast == 1
}

func main() {

}
