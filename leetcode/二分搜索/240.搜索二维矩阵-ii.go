package main

/*
 * @lc app=leetcode.cn id=240 lang=golang
 *
 * [240] 搜索二维矩阵 II
 */

// @lc code=start
//2020/3/9

//边界情况考虑：1.数组中包含了target，2.数组中不包含target（这里面又有三种情况，大于所有，小于所有
//介于中间但是没有）3.数组为空的异常情况

//解法一：每一行都是有序的，因此按照一行一行的暴力二分查找

//解法二：从右上角开始，这个解法最坏的情况是o(m+n)，因为遍历所有的情况
func searchMatrix(matrix [][]int, target int) bool {
	//1.单行单列为空的异常情况不讨论
	if len(matrix) == 0 {
		return false
	}

	//2.起点是右上角
	x, y := 0, len(matrix[0])-1
	for x <= len(matrix)-1 && y >= 0 {
		if target > matrix[x][y] {
			x++
		} else if target < matrix[x][y] {
			y--
		} else {
			return true
		}
	}
	return false
}

//解法三 经过观察这个矩阵左上角是最小的，右下角是最大的
//把矩阵以中心切割成4块，如果大于中心，那么左上块矩阵可以排除，如果小于中心，可以排序右下
//然后中心这个点的坐标是由左上和右下两个坐标点相加除二的出来的。

// @lc code=end
