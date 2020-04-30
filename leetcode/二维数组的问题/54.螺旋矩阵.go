package main

/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 *
 * https://leetcode-cn.com/problems/spiral-matrix/description/
 *
 * algorithms
 * Medium (37.79%)
 * Likes:    249
 * Dislikes: 0
 * Total Accepted:    33K
 * Total Submissions: 87.4K
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'
 *
 * 给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。
 *
 * 示例 1:
 *
 * 输入:
 * [
 * ⁠[ 1, 2, 3 ],
 * ⁠[ 4, 5, 6 ],
 * ⁠[ 7, 8, 9 ]
 * ]
 * 输出: [1,2,3,6,9,8,7,4,5]
 *
 *
 * 示例 2:
 *
 * 输入:
 * [
 * ⁠ [1, 2, 3, 4],
 * ⁠ [5, 6, 7, 8],
 * ⁠ [9,10,11,12]
 * ]
 * 输出: [1,2,3,4,8,12,11,10,9,5,6,7]
 *
 *
 */

// @lc code=start

//解法一：设置方向，每到边界或者已经走过的路就调转方向
func spiralOrder1(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}

	direction := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	res := make([]int, 0)
	route := make([][]bool, len(matrix))
	for i := 0; i < len(matrix); i++ {
		route[i] = make([]bool, len(matrix[0]))
	}

	i, j := 0, 0
	d := 0
	row, col := len(matrix), len(matrix[0])
	var iNext, jNext int
	//终止条件就是调转方向后还是走到已经走过的格子中去了
	for !checkBorder(i, j, row, col) && !route[i][j] {
		//提前算下一次的坐标，就是先判断是否有走偏，走偏了立马转换方向纠正过来
		if iNext, jNext = i+direction[d][0], j+direction[d][1]; checkBorder(iNext, jNext, row, col) || route[iNext][jNext] {
			if d+1 == 4 {
				d = 0
			} else {
				d++
			}
			//如果碰到了边界或者已经走过的地方，需要转换方向了
			iNext, jNext = i+direction[d][0], j+direction[d][1]
		}
		res = append(res, matrix[i][j])
		route[i][j] = true
		i, j = iNext, jNext
	}
	return res
}

//解法一的代码纯净版本
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}

	direction := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	res := make([]int, 0)
	route := make([][]bool, len(matrix))
	for i := 0; i < len(matrix); i++ {
		route[i] = make([]bool, len(matrix[0]))
	}

	i, j := 0, 0
	d := 0
	row, col := len(matrix), len(matrix[0])
	var iNext, jNext int
	//终止条件就是调转方向后还是走到已经走过的格子中去了
	for cnt := 0; cnt < len(matrix)*len(matrix[0]); cnt++ {
		res = append(res, matrix[i][j])
		route[i][j] = true
		if iNext, jNext = i+direction[d][0], j+direction[d][1]; checkBorder(iNext, jNext, row, col) || route[iNext][jNext] {
			d = (d + 1) % 4
			//如果碰到了边界或者已经走过的地方，需要转换方向了
			iNext, jNext = i+direction[d][0], j+direction[d][1]
		}
		i, j = iNext, jNext
	}
	return res
}

func checkBorder(i, j, row, col int) bool {
	if i == row || j == col || i == -1 || j == -1 {
		return true
	}
	return false
}

// @lc code=end
