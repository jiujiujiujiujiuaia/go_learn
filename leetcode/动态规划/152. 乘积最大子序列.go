package main

//做法有问题，待解决
func MaxProduct(nums []int) int {
	dp := make([]int, len(nums))
	best := nums[0]
	dp[0] = best
	for i := 1; i < len(nums); i++ {
		dp[i] = If(abs(nums[i]) > abs(nums[i]*dp[i-1]), nums[i], nums[i]*dp[i-1])
		best = max(best, dp[i])
	}
	return best
}
func abs(a int) int {
	if a >= 0 {
		return a
	} else {
		return -a
	}
}

func If(flag bool, a, b int) int {
	if flag {
		return a
	} else {
		return b
	}
}
