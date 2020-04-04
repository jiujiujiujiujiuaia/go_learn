package main

/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] 搜索二维矩阵
 *
 * https://leetcode-cn.com/problems/search-a-2d-matrix/description/
 *
 * algorithms
 * Medium (36.45%)
 * Likes:    113
 * Dislikes: 0
 * Total Accepted:    25.5K
 * Total Submissions: 69.6K
 * Testcase Example:  '[[1,3,5,7],[10,11,16,20],[23,30,34,50]]\n3'
 *
 * 编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：
 *
 *
 * 每行中的整数从左到右按升序排列。
 * 每行的第一个整数大于前一行的最后一个整数。
 *
 *
 * 示例 1:
 *
 * 输入:
 * matrix = [
 * ⁠ [1,   3,  5,  7],
 * ⁠ [10, 11, 16, 20],
 * ⁠ [23, 30, 34, 50]
 * ]
 * target = 3
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入:
 * matrix = [
 * ⁠ [1,   3,  5,  7],
 * ⁠ [10, 11, 16, 20],
 * ⁠ [23, 30, 34, 50]
 * ]
 * target = 13
 * 输出: false
 *
 */

// @lc code=start

//2020/3/9
func searchMatrix(matrix [][]int, target int) bool {

	if len(matrix) == 0 {
		return false
	}

	//1.行和列
	row := len(matrix)
	col := len(matrix[0])
	l, h := 0, row*col-1

	//2.按照一维数组二分的思想，0,1,2,3,4,5,6..n，只要把这个编号能够和二维数组
	//的坐标对应上，就可以迎刃而解了。
	for l <= h {
		mid := l + (h-l)/2
		var ok bool
		//3.这是重点，把二维当作一维来做
		i, j := mid/col, mid%col
		l, h, ok = help(matrix[i][j], target, mid, l, h)
		if ok {
			return true
		}

	}
	return false
}

func help(number, target, mid, l, h int) (int, int, bool) {
	if number > target {
		return l, mid - 1, false
	} else if number < target {
		return mid + 1, h, false
	} else {
		return 0, 0, true
	}
}

// @lc code=end
