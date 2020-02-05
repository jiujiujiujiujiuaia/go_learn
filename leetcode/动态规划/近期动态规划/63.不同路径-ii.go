package main

/*
 * @lc app=leetcode.cn id=63 lang=golang
 *
 * [63] 不同路径 II
 */

// @lc code=start

//解法的不同路径1相似，只不过加了一点点限制条件而已
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 {
		return 0
	}
	m, n := len(obstacleGrid[0]), len(obstacleGrid)

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}

	dp[0][0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				if i-1 >= 0 {
					dp[i][j] += dp[i-1][j]
				}
				if j-1 >= 0 {
					dp[i][j] += dp[i][j-1]
				}
			}
		}
	}
	return dp[n-1][m-1]
}

// @lc code=end
