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

//2019/10/22重做
//假设没有接触过dp，那么就会觉得很神奇，比如有一段连续数组[1,-1,100]，
//我是怎么样可以在不暴力的情况下， 开天眼知道-1过后有一个100这么大的数字等着我
//假设现在接触了dp，我们就会知道，这就是动态规划的神奇，
//通过记录某一个我们约定好的方案（或者说策略）-> 状态转移方程式，然后就可以一遍得知应该连续-1
func solve(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	best := dp[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > dp[i-1]+nums[i] {
			dp[i] = nums[i]
		} else {
			dp[i] = nums[i] + dp[i]
		}

		if dp[i] > best {
			best = dp[i]
		}
	}
	return best
}
