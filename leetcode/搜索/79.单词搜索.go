package main

/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 *
 */

// @lc code=start

var route [][]bool
var col int
var row int

//典型的搜索类题目+回溯，有的题目是访问过这个节点后，不能再访问了。

//解法一，这个应该算是暴力了，遍历所有的grid，然后每一个grid作为起点
//深度优先遍历。
//同时有一个map记录走过的路程，因为不能同一个字母用两次
//但是有一种特殊情况就是坏情况把好情况的字母用掉了，因此还要作回溯把走过改为未走过
//注：可以写一些测试用例，这样好回顾上面说的那种特殊情况
func exist(board [][]byte, word string) bool {
	if len(board) == 0 {
		return false
	}

	col = len(board)
	row = len(board[0])

	for i := 0; i < col; i++ {
		for j := 0; j < row; j++ {
			initRoute(col, row)
			if judgeWordIsExist(board, i, j, word, 0) {
				return true
			}
		}
	}
	return false
}

func judgeWordIsExist(board [][]byte, i, j int, word string, strIndex int) bool {
	if strIndex == len(word) {
		return true
	}

	if !checkBoard(i, j) {
		return false
	}

	if word[strIndex] == board[i][j] {
		route[i][j] = true
		strIndex++
		if judgeWordIsExist(board, i+1, j, word, strIndex) ||
			judgeWordIsExist(board, i, j+1, word, strIndex) ||
			judgeWordIsExist(board, i-1, j, word, strIndex) ||
			judgeWordIsExist(board, i, j-1, word, strIndex) {
			return true
		}
		//1.这一步回溯非常重要，因为有可能这还会被后续用到，这必须回溯。
		route[i][j] = false
	}
	return false
}

func checkBoard(i, j int) bool {
	if i < 0 || i == col || j < 0 || j == row || route[i][j] {
		return false
	}
	return true
}

func initRoute(col, row int) {
	route = make([][]bool, col)
	for i := 0; i < len(route); i++ {
		route[i] = make([]bool, row)
	}
}

// @lc code=end
