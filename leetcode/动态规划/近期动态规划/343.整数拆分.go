package main

import "fmt"

/*
 * @lc app=leetcode.cn id=343 lang=golang
 *
 * [343] 整数拆分
 */

// @lc code=start

//评论区说可以贪心的找出尽量多的2和3 没看明白
func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(dp[i], max((i-j)*j, dp[i-j]*j))
			fmt.Println(i, j, dp[i])
		}
	}
	return dp[n]
}

// func main() {
// 	a := 8
// 	fmt.Println(integerBreak(a))
// }

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end
