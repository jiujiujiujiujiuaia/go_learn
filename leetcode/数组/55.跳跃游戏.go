package main

/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */

// @lc code=start

//解法一：动态规划的解法 o(n^2)
func canJump1(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	//这里的状态表示从起始位置0是否可以到达n
	dp := make([]bool, len(nums))
	dp[0] = true

	for i := 0; i < len(nums); i++ {
		//1.从内层循环可以看到，每出现一个新的目标i，就是遍历i之前的起跳点，
		//看能不能达到i，那么就可以看到，其实只要最远的起跳点可以到达，
		//就不需要把之前的都遍历一遍了。
		//2.这里贪心的关键点是是选择能够达到距离最远（j+nums[j]）
		//而不是能跳的最远（nums[j]）
		for j := 0; j < i; j++ {
			if dp[j] && j+nums[j] >= i {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(nums)-1]
}

//解法二：从动态规划中看到贪心（见上文分析）
func canJump(nums []int) bool {
	furtherStep := 0
	for i := 0; i < len(nums); i++ {
		//furtherStep代表当前可以跳到的最远距离，如果当前最远比索引i小，那么就无法完成
		if i > furtherStep {
			return false
		}
		//如果可以到达i，那么看从i出发最多可以跳多远
		furtherStep = max(furtherStep, nums[i]+i)
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end
