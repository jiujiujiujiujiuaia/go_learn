package main

import "fmt"

/*
 * @lc app=leetcode.cn id=152 lang=golang
 *
 * [152] 乘积最大子序列
 */

// @lc code=start
func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	//这里引入的状态是当前最小值和当前最大值

	//maxDP[i + 1] = max(maxDP[i] * A[i + 1], A[i + 1],minDP[i] * A[i + 1])
	//minDP[i + 1] = min(minDP[i] * A[i + 1], A[i + 1],maxDP[i] * A[i + 1])
	//dp[i + 1] = max(dp[i], maxDP[i + 1])
	preMin,preMax := 1,1
	maxRes := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			preMax, preMin = preMin, preMax
		}
		preMax = max(preMax*nums[i], nums[i])
		preMin = min(preMin*nums[i], nums[i])

		if preMax > maxRes {
			maxRes = preMax
		}
		fmt.Println(preMax, preMin)
	}
	return maxRes
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	maxProduct([]int{1, 2, -3, -4, -5})
}

// @lc code=end
