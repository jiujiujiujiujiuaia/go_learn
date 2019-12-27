package main

import "strconv"

/*
 * @lc app=leetcode.cn id=37 lang=golang
 *
 * [37] 解数独
 *
 * https://leetcode-cn.com/problems/sudoku-solver/description/
 *
 * algorithms
 * Hard (58.22%)
 * Likes:    270
 * Dislikes: 0
 * Total Accepted:    15.3K
 * Total Submissions: 26.1K
 * Testcase Example:  '[["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]'
 *
 * 编写一个程序，通过已填充的空格来解决数独问题。
 *
 * 一个数独的解法需遵循如下规则：
 *
 *
 * 数字 1-9 在每一行只能出现一次。
 * 数字 1-9 在每一列只能出现一次。
 * 数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
 *
 *
 * 空白格用 '.' 表示。
 *
 *
 *
 * 一个数独。
 *
 *
 *
 * 答案被标成红色。
 *
 * Note:
 *
 *
 * 给定的数独序列只包含数字 1-9 和字符 '.' 。
 * 你可以假设给定的数独只有唯一解。
 * 给定数独永远是 9x9 形式的。
 *
 *
 */

// @lc code=start
var row []map[byte]bool
var col []map[byte]bool
var box []map[byte]bool
var boardCopy [][]byte
var solveFlag bool

func solveSudoku(board [][]byte) {
	if len(board) == 0 {
		return
	}
	initGrid(board)
	backTrace(0, 0)

}

func placeNextNumver(i, j int) {
	if i == 8 && j == 8 {
		solveFlag = true
	} else {
		if j == 8 {
			backTrace(i+1, 0)
		} else {
			backTrace(i, j+1)
		}
	}
}

func backTrace(i, j int) {
	if boardCopy[i][j] == '.' {
		for number := 1; number < 10; number++ {
			b := intToByte(number)
			idx := (i/3)*3 + j/3
			if row[i][b] || col[j][b] || box[idx][b] {
				continue
			} else {
				placeNumber(i, j, number)
				placeNextNumver(i, j)
				if solveFlag {
					return
				} else {
					removeNumber(i, j, number)
				}
			}
		}
	} else {
		placeNextNumver(i, j)
	}
}

func placeNumber(i, j int, number int) {
	b := intToByte(number)
	boardCopy[i][j] = b
	idx := (i/3)*3 + j/3
	row[i][b] = true
	col[j][b] = true
	box[idx][b] = true
}

func removeNumber(i, j int, number int) {
	b := intToByte(number)
	boardCopy[i][j] = '.'
	idx := (i/3)*3 + j/3
	row[i][b] = false
	col[j][b] = false
	box[idx][b] = false
}

func intToByte(i int) byte {
	str := strconv.FormatInt(int64(i), 10)
	return []byte(str)[0]
}

func initGrid(board [][]byte) {
	row = make([]map[byte]bool, 9)
	col = make([]map[byte]bool, 9)
	box = make([]map[byte]bool, 9)
	for i := 0; i < 9; i++ {
		row[i] = make(map[byte]bool)
		col[i] = make(map[byte]bool)
		box[i] = make(map[byte]bool)
	}
	boardCopy = board
	solveFlag = false

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if number := board[i][j]; number != '.' {
				idx := (i/3)*3 + j/3
				row[i][number] = true
				col[j][number] = true
				box[idx][number] = true
			}
		}
	}

}

// @lc code=end
