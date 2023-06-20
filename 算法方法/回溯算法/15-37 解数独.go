package main

import "fmt"

//LeetCode 37. 解数独

//编写一个程序，通过填充空格来解决数独问题。

//一个数独的解法需遵循如下规则： 数字 1-9 在每一行只能出现一次。 数字 1-9 在每一列只能出现一次。
//数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。 空白格用 '.' 表示。

//提示：
// 给定的数独序列只包含数字 1-9 和字符 '.' 。
// 你可以假设给定的数独只有唯一解。
// 给定数独永远是 9x9 形式的。

// 卡尔 todo

//在树形图中可以看出我们需要的是一个二维的递归（也就是两个for循环嵌套着递归）
//一个for循环遍历棋盘的行，一个for循环遍历棋盘的列，一行一列确定下来之后，递归遍历这个位置放9个数字的可能性！

//判断棋盘是否合法有如下三个维度：
// 同行是否重复
// 同列是否重复
// 9宫格里是否重复

func solveSudoku(board [][]byte) {
	backtracking(board)
}
func backtracking(board [][]byte) bool {
	for i := 0; i < 9; i++ { // 遍历行
		for j := 0; j < 9; j++ { // 遍历列
			//判断此位置是否适合填数字
			if board[i][j] != '.' {
				continue
			}
			//尝试填1-9
			for k := '1'; k <= '9'; k++ { // (i, j) 这个位置放k是否合适
				if isvalid(i, j, byte(k), board) == true { //如果满足要求就填
					board[i][j] = byte(k)            // 放置k
					if backtracking(board) == true { // 如果找到合适一组立刻返回
						return true
					}
					board[i][j] = '.' // 撤销k
				}
			}

			return false // 9个数都试完了，都不行，那么就返回false
		}
	}

	return true // 遍历完没有返回false，说明找到了合适棋盘位置了
}

//判断填入数字是否满足要求
func isvalid(row, col int, k byte, board [][]byte) bool {
	for i := 0; i < 9; i++ { //行
		if board[row][i] == k {
			return false
		}
	}
	for i := 0; i < 9; i++ { //列
		if board[i][col] == k {
			return false
		}
	}
	//方格
	startrow := (row / 3) * 3
	startcol := (col / 3) * 3
	for i := startrow; i < startrow+3; i++ {
		for j := startcol; j < startcol+3; j++ {
			if board[i][j] == k {
				return false
			}
		}
	}
	return true
}

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	solveSudoku(board)
	fmt.Println(board)
}
