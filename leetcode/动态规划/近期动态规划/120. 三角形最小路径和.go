package main

import "fmt"

//https://leetcode-cn.com/problems/triangle/submissions/
//这个题目很经典的展示了局部最优解不等于全局最优解的概念，这就区分了动态规划和贪心。
func main() {
	fmt.Println(minimumTotal2([][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}))
}

//这个空间复杂度为o（n^2）
func minimumTotal3(triangle [][]int) int {
	dp := make([][]int, len(triangle))
	for i := 0; i < len(triangle); i++ {
		dp[i] = make([]int, len(triangle[i]))
	}
	for i := len(triangle) - 1; i > -1; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			if i+1 < len(triangle) {
				dp[i][j] = triangle[i][j] + min(dp[i+1][j], dp[i+1][j+1])
			} else {
				dp[i][j] = triangle[i][j]
			}
		}
	}
	return dp[0][0]
}

//同时也可以使用滚动数组的办法，空间复杂度为o（n）
//根据上面自底向上的方法得出，算dp第三层时，只需要第四层的数据，dp第二层只要第三层，不需要第四层。
//因此采取覆盖的策略，可以不用二维数组的空间
func minimumTotal2(triangle [][]int) int {
	dp := make([]int, len(triangle))
	for row := len(triangle) - 1; row > -1; row-- {
		for j := 0; j <= row; j++ {
			if row == len(triangle)-1 {
				dp[j] = triangle[row][j]
			} else {
				dp[j] = triangle[row][j] + min(dp[j], dp[j+1])
			}
		}
	}
	return dp[0]
}

//这个是记忆化搜索的方法，原地修改，没有利用额外空间
func minimumTotal(triangle [][]int) int {
	for i := len(triangle) - 1; i > -1; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			if i+1 < len(triangle) {
				triangle[i][j] = triangle[i][j] + min(triangle[i+1][j], triangle[i+1][j+1])
			} else {
				triangle[i][j] = triangle[i][j]
			}
		}
	}
	return triangle[0][0]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
