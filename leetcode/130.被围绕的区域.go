package main

/*
 * @lc app=leetcode.cn id=130 lang=golang
 *
 * [130] 被围绕的区域
 */

// @lc code=start

//解法一：很经典的利用dfs搜索的题目，题目中的变量visit direction 以及函数
//initRoute checkBoard 都可以称之为模板了。

//解法二：并查集的思路解决：遍历二维数组，如果该点在边界的话，则和dummy节点
//进行union操作，然后其他节点就和它周围的o节点union
//然后再遍历一遍数组，对每个o节点判断和dummy节点的联通性，如果联通，则不为x
//否则就为x
var col int
var row int
var visit [][]bool
var direction = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}

	initRoute(board)
	//首行和尾行dfs
	for i := 0; i < col; i++ {
		if board[i][0] == 'O' {
			dfs(board, i, 0)
		}

		if board[i][row-1] == 'O' {
			dfs(board, i, row-1)
		}
	}
	//首列和尾列dfs
	for j := 0; j < row; j++ {
		if board[0][j] == 'O' {
			dfs(board, 0, j)
		}

		if board[col-1][j] == 'O' {
			dfs(board, col-1, j)
		}
	}

	for i := 0; i < col; i++ {
		for j := 0; j < row; j++ {
			if board[i][j] == '#' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}

		}
	}
}

func dfs(board [][]byte, i, j int) {
	//1.判断是否访问过该点以及该点有没有超出边界
	if !checkBoard(i, j) {
		return
	}

	visit[i][j] = true
	//2.如果该点为x，直接返回，如果该点为o，进一步探究
	if board[i][j] == 'O' {
		board[i][j] = '#'
		for _, d := range direction {
			dfs(board, i+d[0], j+d[1])
		}
	}
}

func checkBoard(i, j int) bool {
	if i < 0 || i == col || j < 0 || j == row || visit[i][j] {
		return false
	}
	return true
}

func initRoute(board [][]byte) {
	col = len(board)
	row = len(board[0])
	visit = make([][]bool, col)
	for i := 0; i < len(visit); i++ {
		visit[i] = make([]bool, row)
	}
}

// @lc code=end
