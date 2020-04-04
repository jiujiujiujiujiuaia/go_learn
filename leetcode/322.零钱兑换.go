package main

/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */

// @lc code=start
func coinChange(coins []int, amount int) int {
	dp := make([]int,amount + 1)
	for i:=1;i<len(dp);i++{
		dp[i] = amount+1
	}
	dp[0] = 0
	for i := 1;i<= amount;i++{
		for j:=0;j<len(coins);j++{
			if i >= coins[j]{
				dp[i] = min(dp[i],dp[i-coins[j]] + 1)
			}
		}
	}
	fmt.Println(dp,amount)
	if dp[amount] == amount+1{
		return -1
	}else{
		return dp[amount]
	}
}

func min(a,b int)int{
	if a < b{
		return a
	}else{
		return b
	}
}

// @lc code=end
