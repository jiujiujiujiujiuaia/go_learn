package main

import "fmt"

/*
 * @lc app=leetcode.cn id=343 lang=golang
 *
 * [343] 整数拆分
 */

// @lc code=start

//评论区说可以贪心的找出尽量多的2和3 没看明白

//这个题目其实有点启发的，动态规划的问题由于有一个dp数组的存在
//如果n很大，会有比较大的空间开销，并且这个题目返回乘积的结果可能会出现溢出的情况

//因此：无论是小数据规模还是大数据规模，贪心是更好的解法
//同时由于这道题目可能存在n很大的情况，快速幂取余和快速幂要掌握呀

//贪心的思路：当n<=4 都是一些固定答案
//当n>4 优先分解成3 原因：(n-3) * 3 > (n-2)*2 -> n >= 5
func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1

	//这里比较重要的是dp[i-j]不一定比i-j大
	//正面例子dp[6]=9 dp[10]=4*dp[6]
	//反面例子dp[4]=2*2 而不是dp[4]=2*dp[2]
	for i := 1; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(dp[i], max((i-j)*j, dp[i-j]*j))
			fmt.Println(i, j, dp[i])
		}
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end
