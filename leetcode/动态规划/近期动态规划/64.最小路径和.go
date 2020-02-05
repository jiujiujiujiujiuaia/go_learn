package main

/*
 * @lc app=leetcode.cn id=64 lang=golang
 *
 * [64] 最小路径和
 */

// @lc code=start
func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		dp[i] = make([]int, len(grid[0]))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i-1 >= 0 && j-1 >= 0 {
				dp[i][j] += min(dp[i-1][j], dp[i][j-1])
			} else if i-1 < 0 && j-1 >= 0 {
				dp[i][j] += dp[i][j-1]
			} else if j-1 < 0 && i-1 >= 0 {
				dp[i][j] += dp[i-1][j]
			}
			dp[i][j] += grid[i][j]
		}
	}
	return dp[len(grid)-1][len(grid[0])-1]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// @lc code=end
