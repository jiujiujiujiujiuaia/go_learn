package main

import "container/list"

/*
 * @lc app=leetcode.cn id=542 lang=golang
 *
 * [542] 01 矩阵
 *
 * https://leetcode-cn.com/problems/01-matrix/description/
 *
 * algorithms
 * Medium (36.92%)
 * Likes:    128
 * Dislikes: 0
 * Total Accepted:    8.5K
 * Total Submissions: 23K
 * Testcase Example:  '[[0,0,0],[0,1,0],[0,0,0]]'
 *
 * 给定一个由 0 和 1 组成的矩阵，找出每个元素到最近的 0 的距离。
 *
 * 两个相邻元素间的距离为 1 。
 *
 * 示例 1:
 * 输入:
 *
 *
 * 0 0 0
 * 0 1 0
 * 0 0 0
 *
 *
 * 输出:
 *
 *
 * 0 0 0
 * 0 1 0
 * 0 0 0
 *
 *
 * 示例 2:
 * 输入:
 *
 *
 * 0 0 0
 * 0 1 0
 * 1 1 1
 *
 *
 * 输出:
 *
 *
 * 0 0 0
 * 0 1 0
 * 1 2 1
 *
 *
 * 注意:
 *
 *
 * 给定矩阵的元素个数不超过 10000。
 * 给定矩阵中至少有一个元素是 0。
 * 矩阵中的元素只在四个方向上相邻: 上、下、左、右。
 *
 *
 */

// @lc code=start
type point struct {
	x_axis   int
	y_axis   int
	distance int
}

var col int
var row int
var matrixRoute [][]bool

func updateMatrix(matrix [][]int) [][]int {
	if len(matrix) == 0 {
		return nil
	}
	col = len(matrix)
	row = len(matrix[0])
	res := make([][]int, col)
	for i := 0; i < col; i++ {
		res[i] = make([]int, row)
	}
	for i := 0; i < col; i++ {
		for j := 0; j < row; j++ {
			initMatrixRoute(col, row)
			d := solveUpdateMatrix(matrix, i, j)
			res[i][j] = d
		}
	}
	return res
}

func solveUpdateMatrix(matrix [][]int, i, j int) int {
	queue := list.New()
	p := point{
		x_axis:   i,
		y_axis:   j,
		distance: 0,
	}
	queue.PushBack(p)
	for queue.Len() != 0 {
		head := queue.Front()
		queue.Remove(head)
		x := head.Value.(point).x_axis
		y := head.Value.(point).y_axis
		matrixRoute[x][y] = true
		if matrix[x][y] == 1 {
			push(queue, head.Value.(point))
		} else {
			return head.Value.(point).distance
		}
	}
	return 0
}

func initMatrixRoute(col, row int) {
	matrixRoute = make([][]bool, col)
	for i := 0; i < len(matrixRoute); i++ {
		matrixRoute[i] = make([]bool, row)
	}
}

func push(queue *list.List, p point) {
	x, y, d := p.x_axis, p.y_axis, p.distance
	if x+1 < col && !matrixRoute[x+1][y] {
		queue.PushBack(point{x_axis: x + 1, y_axis: y, distance: d + 1})
	}

	if y+1 < row && !matrixRoute[x][y+1] {
		queue.PushBack(point{x_axis: x, y_axis: y + 1, distance: d + 1})
	}

	if x-1 > -1 && !matrixRoute[x-1][y] {
		queue.PushBack(point{x_axis: x - 1, y_axis: y, distance: d + 1})
	}

	if y-1 > -1 && !matrixRoute[x][y-1] {
		queue.PushBack(point{x_axis: x, y_axis: y - 1, distance: d + 1})
	}
}

// @lc code=end
