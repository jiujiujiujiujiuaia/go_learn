package main

/*
 * @lc app=leetcode.cn id=1143 lang=golang
 *
 * [1143] 最长公共子序列
 */

// @lc code=start
func longestCommonSubsequence(text1 string, text2 string) int {
	col := len(text1)
	row := len(text2)

	//加一层是为了方便初步处理，也就是i-1,j-1的时候不需要作是否超出界限的判断
	dp := make([][]int, row+1)
	for i := 0; i < row+1; i++ {
		dp[i] = make([]int, col+1)
	}

	for i := 1; i <= row; i++ {
		for j := 1; j <= col; j++ {
			if text2[i-1] == text1[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[row][col]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end
