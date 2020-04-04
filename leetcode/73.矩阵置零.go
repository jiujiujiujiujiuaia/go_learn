package main

/*
 * @lc app=leetcode.cn id=73 lang=golang
 *
 * [73] 矩阵置零
 */

// @lc code=start
var IsRow bool
var IsCol bool

//解法一：用两个数组，一个表示行，一个表示列，如果某个存在，则标记
//时间复杂度相同，空间n+m

//解法二：对于不是第一行或者第一列的数，如果为0，就把它的首列或者首行标记=0
//标记完后再遍历解决
//时间复杂度是nXm 空间是原地的
func setZeroes(matrix [][]int) {
	IsCol = false
	IsRow = false
	//1.判断第一列
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			IsCol = true
			break
		}
	}

	//2.判断第一行
	for i := 0; i < len(matrix[0]); i++ {
		if matrix[0][i] == 0 {
			IsRow = true
			break
		}
	}

	//（1）先标记非第一行第一列的情况（这里有顺序关系的）
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	//（2）再处理之前的标记
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	//（3）再处理首行首列
	if IsCol {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}

	if IsRow {
		for i := 0; i < len(matrix[0]); i++ {
			matrix[0][i] = 0
		}
	}
}

// @lc code=end
