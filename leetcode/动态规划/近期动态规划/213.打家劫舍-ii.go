package main

/*
 * @lc app=leetcode.cn id=213 lang=golang
 *
 * [213] 打家劫舍 II
 */

// @lc code=start

//打家劫舍2只是在1的基础上稍稍变化了第一栋和最后一栋不可以相邻
//那就有了两种情况，一种是第一栋到倒数第二栋，另一种
//是第二栋到最后一栋
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	length := len(nums) - 1
	return max(helpRob(nums[:length]), helpRob(nums[1:]))
}

func helpRob(nums []int) int {
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
