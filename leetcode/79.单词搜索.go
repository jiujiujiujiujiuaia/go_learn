package main

/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 *
 * https://leetcode-cn.com/problems/word-search/description/
 *
 * algorithms
 * Medium (39.45%)
 * Likes:    263
 * Dislikes: 0
 * Total Accepted:    29.2K
 * Total Submissions: 73.6K
 * Testcase Example:  '[["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]\n"ABCCED"'
 *
 * 给定一个二维网格和一个单词，找出该单词是否存在于网格中。
 *
 * 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
 *
 * 示例:
 *
 * board =
 * [
 * ⁠ ['A','B','C','E'],
 * ⁠ ['S','F','C','S'],
 * ⁠ ['A','D','E','E']
 * ]
 *
 * 给定 word = "ABCCED", 返回 true.
 * 给定 word = "SEE", 返回 true.
 * 给定 word = "ABCB", 返回 false.
 *
 */

// @lc code=start

var route [][]bool
var col int
var row int

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
		} else {
			route[i][j] = false
			return false
		}
	} else {
		return false
	}
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
