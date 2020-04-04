package main

/*
 * @lc app=leetcode.cn id=48 lang=golang
 *
 * [48] 旋转图像
 */

// @lc code=start

//解法一：题目要求是的顺时针90
//可以先顺时针180（对折），再逆时针90（每一行翻转）
func rotate1(matrix [][]int) {

	//（1）沿着对角线把矩形分为上半和下半，上半和下半交换
	for i := 0; i < len(matrix); i++ {
		for j := i; j < len(matrix); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	//(2)翻转每一行
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix)/2; j++ {
			matrix[i][j], matrix[i][len(matrix)-1-j] = matrix[i][len(matrix)-1-j], matrix[i][j]
		}
	}
}

//解法二：把矩阵平均切成四块，那么只需要按照块来旋转，即可完成
func rotate(matrix [][]int) {
	n := len(matrix)

	//这里n+1是为了应付n为奇数的情况
	//这里的i和j的范围是矩形左上四分之一的大小
	for i := 0; i < (n+1)/2; i++ {
		for j := 0; j < n/2; j++ {
			//(1)关于这里的坐标如何推的：
			//   i,j -> j,n-1-i
			//   j,n-1-i->n-1-i,n-1-j
			//上一个的列是下一个的行
			//上一个的行是下一个n-1-上一个行
			temp := matrix[i][j]
			matrix[i][j] = matrix[n-1-j][i]
			matrix[n-1-j][i] = matrix[n-1-i][n-1-j]
			matrix[n-1-i][n-1-j] = matrix[j][n-1-i]
			matrix[j][n-1-i] = temp
		}
	}
}

// @lc code=end
