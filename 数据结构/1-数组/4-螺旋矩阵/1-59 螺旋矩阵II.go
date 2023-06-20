package main

import "fmt"

//LeetCode 59.螺旋矩阵II

//给定一个正整数 n，生成一个包含 1 到 n^2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。

//示例:
//输入: 3 输出: [ [ 1, 2, 3 ], [ 8, 9, 4 ], [ 7, 6, 5 ] ]

//本题依然是要坚持循环不变量原则。

//方法1：切一下，推荐，好理解
//模拟顺时针画矩阵的过程:
// 填充上行从左到右
// 填充右列从上到下
// 填充下行从右到左
// 填充左列从下到上
//由外向内一圈一圈这么画下去。
//1 1 2
//4 5 2
//4 3 3
func generateMatrix(n int) [][]int {
	top, bottom := 0, n-1
	left, right := 0, n-1
	num := 1
	tar := n * n
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	for num <= tar {
		for i := left; i <= right; i++ {
			matrix[top][i] = num
			num++
		}
		top++
		for i := top; i <= bottom; i++ {
			matrix[i][right] = num
			num++
		}
		right--
		for i := right; i >= left; i-- {
			matrix[bottom][i] = num
			num++
		}
		bottom--
		for i := bottom; i >= top; i-- {
			matrix[i][left] = num
			num++
		}
		left++
	}

	return matrix
}

//方法二：按层模拟-官方 todo
//可以将矩阵看成若干层，首先填入矩阵最外层的元素，其次填入矩阵次外层的元素，直到填入矩阵最内层的元素。
//时间复杂度：O(n^2 )，其中 n 是给定的正整数。矩阵的大小是 n×n，需要填入矩阵中的每个元素。
//空间复杂度：O(1)。除了返回的矩阵以外，空间复杂度是常数。
func generateMatrix2(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	num := 1
	left, right, top, bottom := 0, n-1, 0, n-1
	for left <= right && top <= bottom {
		for column := left; column <= right; column++ {
			matrix[top][column] = num
			num++
		}
		for row := top + 1; row <= bottom; row++ {
			matrix[row][right] = num
			num++
		}
		if left < right && top < bottom {
			for column := right - 1; column > left; column-- {
				matrix[bottom][column] = num
				num++
			}
			for row := bottom; row > top; row-- {
				matrix[row][left] = num
				num++
			}
		}
		left++
		right--
		top++
		bottom--
	}
	return matrix
}

func main() {
	num := 3
	fmt.Println(generateMatrix(num))
}
