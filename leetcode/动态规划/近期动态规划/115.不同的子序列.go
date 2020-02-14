package main

/*
 * @lc app=leetcode.cn id=115 lang=golang
 *
 * [115] 不同的子序列
 */

// @lc code=start
func numDistinct(s string, t string) int {
	dp := make([][]int, len(s)+1)
	for i := 0; i < len(s)+1; i++ {
		dp[i] = make([]int, len(t)+1)
		dp[i][0] = 1
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			if s[i-1] == t[j-1] {
				//如果s(i)与t(j)相同，那么用s(i)这个字符和不用s(i)这个字符的情况
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				//两个不相同，只有不用这个字符了
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(s)][len(t)]
}

// @lc code=end
