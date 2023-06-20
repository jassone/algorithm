package main

import (
	"fmt"
	"strings"
)

//LeetCode 51. N皇后

//n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。

//给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。

//每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。

//首先来看一下皇后们的约束条件：
// 不能同行
// 不能同列
// 不能同斜线

var res [][]string

//棋盘的宽度就是for循环的长度，递归的深度就是棋盘的高度，这样就可以套进回溯法的模板里了。
func solveNQueens(n int) [][]string {
	// 先规划好二维数组
	chessboard := make([][]string, n)
	for i := 0; i < n; i++ {
		chessboard[i] = make([]string, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			chessboard[i][j] = "."
		}
	}

	backtrack(0, n, chessboard)
	return res
}
func backtrack(row, n int, chessboard [][]string) {
	if row == n { // 到了最后一个节点了
		temp := make([]string, n)
		for i, rowStr := range chessboard {
			temp[i] = strings.Join(rowStr, "")
		}
		res = append(res, temp)
		return
	}

	for i := 0; i < n; i++ {
		if isValid(n, row, i, chessboard) {
			chessboard[row][i] = "Q"
			backtrack(row+1, n, chessboard)
			chessboard[row][i] = "."
		}
	}
}

// 判断行列和斜线上是否有Q
func isValid(n, row, col int, chessboard [][]string) bool {
	for i := 0; i < row; i++ { // 当前列上没有Q
		if chessboard[i][col] == "Q" {
			return false
		}
	}
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if chessboard[i][j] == "Q" {
			return false
		}
	}
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if chessboard[i][j] == "Q" {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(solveNQueens(4))
}
