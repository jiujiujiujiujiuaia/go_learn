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
func spiralOrder(matrix [][]int) []int {
	direction := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	res := make([]int, 0)
	route := make([][]bool, len(matrix))
	for i := 0; i < len(matrix[0]); i++ {
		route[i] = make([]bool, len(matrix[0]))
	}

	i, j := 0, 0
	d := 0
	row, col := len(matrix), len(matrix[0])
	var iNext, jNext int
	for !checkBorder(i, j, row, col) && !route[i][j] {
		if iNext, jNext = i+direction[d][0], j+direction[d][1]; checkBorder(iNext, jNext, row, col) || route[iNext][jNext] {
			if d+1 == 4 {
				d = 0
			} else {
				d++
			}
			iNext, jNext = i+direction[d][0], j+direction[d][1]
		}
		res = append(res, matrix[i][j])
		route[i][j] = true
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

func main() {
	a := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	spiralOrder(a)
}

// @lc code=end
