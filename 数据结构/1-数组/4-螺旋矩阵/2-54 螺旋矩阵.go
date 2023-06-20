package main

import (
	"fmt"
)

//LeetCode 54. 螺旋矩阵

// 给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。

//输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
//输出：[1,2,3,6,9,8,7,4,5]

//方法1：和题59一样的解法。边界问题。推荐
// 该算法的时间复杂度是O(m*n)，空间复杂度是O(1)。
func f1(arr [][]int) []int {
	n := len(arr)
	top, bottom := 0, n-1
	left, right := 0, n-1
	num := 1
	tar := n * n

	res := []int{} // 定义结果数组
	for num <= tar {
		for i := left; i <= right; i++ {
			res = append(res, arr[left][i])
			num++
		}
		top++
		for i := top; i <= bottom; i++ {
			res = append(res, arr[i][right])
			num++
		}
		right--
		for i := right; i >= left; i-- {
			res = append(res, arr[bottom][i])
			num++
		}
		bottom--
		for i := bottom; i >= top; i-- {
			res = append(res, arr[i][left])
			num++
		}
		left++
	}

	return res
}

// 官方比较复杂
//https://leetcode-cn.com/problems/spiral-matrix/solution/luo-xuan-ju-zhen-by-leetcode-solution/

func main() {
	arr := [][]int{
		{1, 2, 3, 4, 5},
		{16, 17, 18, 19, 6},
		{15, 24, 25, 20, 7},
		{14, 23, 22, 21, 8},
		{13, 12, 11, 10, 9},
	}

	fmt.Println(f1(arr))
}
