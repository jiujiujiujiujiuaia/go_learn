package main

import "container/list"

import "math"

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

//思路分析：暴力的办法就是每到一个值为1的点，就去搜索整个数组，来拿到到0的最小值
//广度搜索的思路就是是要先算临近的点，因为稍远的点是由临近的点得出的

//也就是每个点只有一次机会进入路径组

//思路2：这个题的dp思路有意思，先只比较一个元素和他自己的左和上的元素最小值，
//然后再比较比较一个元素和他自己的右和下的元素最小值
//这样比的好处是，在比任何一个元素时，自己的邻居元素已经被算出来了，可以用记忆化搜索的方式理解。

type point struct {
	x_axis   int
	y_axis   int
	distance int
}

var updateMatrixCol int
var updateMatrixRow int
var routeMap [][]bool

func updateMatrix(matrix [][]int) [][]int {
	if len(matrix) == 0 {
		return [][]int{}
	}
	updateMatrixRow = len(matrix)
	updateMatrixCol = len(matrix[0])
	routeMap = make([][]bool, updateMatrixRow)
	for i := 0; i < updateMatrixRow; i++ {
		routeMap[i] = make([]bool, updateMatrixCol)
	}

	queue := list.New()
	for i := 0; i < updateMatrixRow; i++ {
		for j := 0; j < updateMatrixCol; j++ {
			if matrix[i][j] == 0 {
				queue.PushBack(point{
					x_axis: i,
					y_axis: j,
				})
				routeMap[i][j] = true
			} else {
				matrix[i][j] = math.MaxInt32
			}
		}
	}
	vectors := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for queue.Len() > 0 {
		headPoint := queue.Front()
		queue.Remove(queue.Front())
		x, y := headPoint.Value.(point).x_axis, headPoint.Value.(point).y_axis
		for _, vector := range vectors {
			nextX, nextY := x+vector[0], y+vector[1]
			if nextX > -1 && nextX < updateMatrixRow &&
				nextY > -1 && nextY < updateMatrixCol && !routeMap[nextX][nextY] {
				queue.PushBack(point{
					x_axis: nextX,
					y_axis: nextY,
				})
				routeMap[nextX][nextY] = true
				matrix[nextX][nextY] = matrix[x][y] + 1
			}
		}
	}

	return matrix
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// @lc code=end
