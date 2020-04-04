package main

import (
	"math"
)

/*
 * @lc app=leetcode.cn id=188 lang=golang
 *
 * [188] 买卖股票的最佳时机 IV
 */

// @lc code=start

//当k>1小于len(prices)/2的情况：
//对于第i天，最多能交易k次。都有这样的状态，要么买入，要么卖出
//一个二维数组的用i，j表示，i表示第i天，j表示当天是买入还是卖出的case{0,1}x
func maxProfit(k int, prices []int) int {
	if k<1{
		return 0
	}

	//1.如果允许交易次数大于数组一半时
	if k>=len(prices)/2{
		return MakeMostProfitByGreedy(prices)
	}

	//2.如果次数在1-n/2之间

	//(1)建立一个三维数组的状态表示第i天进行了最多可以进行k次交易持有和不持有的最大利润
	dp := preHandle(k, len(prices))
	dp[0][0][0] = 0
	dp[0][0][1] = math.MinInt32
	for i := 1; i <= len(prices); i++ {
		for j := 1; j <= k; j++ {
			//(2)当我写好状态转移后，就要去思考，如何推动这个方程从最初开始计算，
			//就是这个方程的初试条件是什么
			if i == 1 {
				//(2.1)由于第一天第k次交易的状态依赖第0天第k次交易,因此这里需要给出初始条件
				dp[0][j][0] = 0
				dp[0][j][1] = math.MinInt32
			}

			//(3.1)今天的不持有股票最大利润:1.昨天也不持有2.昨天持有股票，今天卖掉了
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i-1])
			//(3.2)今天的持有股票最大利润是：1.昨天也持有股票2.昨天没有股票，今天买了股票，交易状态发生了变化
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i-1])
		}
	}

	for i:=0;i<=len(prices);i++{
		fmt.Println(i,dp[i])
}
	fmt.Println(dp)
	return dp[len(prices)][k][0]
}

//当允许交易次数大于数组一半时，说明相当于可以无限交易，因此贪心解决
func MakeMostProfitByGreedy(prices []int) int {
	maxProfit := 0 
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxProfit += prices[i] - prices[i-1]
		}
	}
	return maxProfit
}

func preHandle(k, days int) [][][2]int {
	k++
	days++
	res := make([][][2]int, days)
	for i := 0; i < days; i++ {
		res[i] = make([][2]int, k)
		for j := 0; j < k; j++ {
			res[i][j] = [2]int{}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end
