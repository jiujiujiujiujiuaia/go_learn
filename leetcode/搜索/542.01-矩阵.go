package main

import "container/list"

import "math"

/*
 * @lc app=leetcode.cn id=542 lang=golang
 *
 * [542] 01 矩阵
 *
 * https://leetcode-cn.com/problems/01-matrix/description/
 */

// @lc code=start

//思路分析：暴力的办法就是每到一个值为1的点，就去搜索整个数组，来拿到到0的最小值
//广度搜索的思路就是是要先算临近的点，因为稍远的点是由临近的点得出的
//也就是每个点只有一次机会进入路径组

type point struct {
	x_axis   int
	y_axis   int
	distance int
}

var updateMatrixCol int
var updateMatrixRow int
var routeMap [][]bool

//解法一：使用bfs做
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

	//1. 预处理，把节点为0的元素丢进队列并map置为true
	queue := list.New()
	for i := 0; i < updateMatrixRow; i++ {
		for j := 0; j < updateMatrixCol; j++ {
			if matrix[i][j] == 0 {
				queue.PushBack(point{
					x_axis: i,
					y_axis: j,
				})
				routeMap[i][j] = true
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

//解法二：dp思想做，只不过分为两次，一次是左上到右下，一次是右下到左上。


func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// @lc code=end
