package main

/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] 打家劫舍
 */

// @lc code=start

//解法一：空间复杂度较o(n)
func rob1(nums []int) int {
	//1.为了满足递推公式，dp增加两个空间
	dp := make([]int, len(nums)+2)

	dp[0] = 0
	dp[1] = 0

	//面临一个新房子的新时候，是选择偷这个，还是不偷这个偷前面一个
	//dp(n) = max(dp(n-2)+a(n),dp(n-1))
	for i := 2; i < len(dp); i++ {
		//2.dp数组的索引和nums数组的索引下标意义不同
		dp[i] = max(dp[i-2]+nums[i-2], dp[i-1])
	}
	return dp[len(nums)+1]
}

//解法二：空间复杂度常数级别
func rob(nums []int) int {
	//1.根据递推公式可以看到这里只和前一个dp状态和前前一个dp状态有关
	//因此就用两个变量来存储
	preMaxMoney, prePreMaxMoney := 0, 0
	for i := 0; i < len(nums); i++ {
		curMaxMoney := max(nums[i]+prePreMaxMoney, preMaxMoney)
		prePreMaxMoney = preMaxMoney
		preMaxMoney = curMaxMoney
	}
	return preMaxMoney
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end
