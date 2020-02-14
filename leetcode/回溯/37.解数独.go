package main

import "strconv"

/*
 * @lc app=leetcode.cn id=37 lang=golang
 *
 * [37] 解数独
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

//主函数
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
