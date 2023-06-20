package main

// todo
//LeetCode 454 四数相加II
//给定四个包含整数的数组列表 A , B , C , D ,计算有多少个元组 (i, j, k, l) ，使得 A[i] + B[j] + C[k] + D[l] = 0。

//为了使问题简单化，所有的 A, B, C, D 具有相同的长度 N，且 0 ≤ N ≤ 500 。所有整数的范围在 -2^28 到 2^28 - 1 之间，最终结果不会超过 2^31 - 1 。

//例如:
//输入: A = [ 1, 2] B = [-2,-1] C = [-1, 2] D = [ 0, 2]
//输出: 2
//解释: 两个元组如下:
//(0, 0, 0, 1) -> A[0] + B[0] + C[0] + D[1] = 1 + (-2) + (-1) + 2 = 0
//(1, 1, 0, 0) -> A[1] + B[1] + C[0] + D[0] = 2 + (-1) + (-1) + 0 = 0

//方法一：分组 + 哈希表-官方
//我们可以将四个数组分成两部分，AA 和 BB 为一组，CC 和 DD 为另外一组。

//时间复杂度：O(n^2)。我们使用了两次二重循环，时间复杂度均为 O(n^2)。在循环中对哈希映射进行的修改以及查询操作的期望
// 时间复杂度均为 O(1)，因此总时间复杂度为 O(n^2)。
//空间复杂度：O(n^2)，即为哈希映射需要使用的空间。在最坏的情况下，A[i]+B[j] 的值均不相同，因此值的个数为 n^2
//，也就需要 O(n^2) 的空间。
func fourSumCount(a, b, c, d []int) (ans int) {
	countAB := map[int]int{}
	for _, v := range a {
		for _, w := range b {
			countAB[v+w]++
		}
	}
	for _, v := range c {
		for _, w := range d {
			ans += countAB[-v-w] // 查看map中是否有，有则不为0，无则为0
		}
	}
	return
}

func main() {

}
