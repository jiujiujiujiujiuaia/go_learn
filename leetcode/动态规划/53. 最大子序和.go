package main

func MaxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	best := nums[0]
	dp[0] = best
	for i := 1; i < len(nums); i++ {
		dp[i] = max(nums[i], nums[i]+dp[i-1])
		best = max(best, dp[i])
	}
	return best
}
