package main

/*
 * @lc app=leetcode.cn id=442 lang=golang
 *
 * [442] 数组中重复的数据
 */

// @lc code=start

//同448
func findDuplicates(nums []int) []int {
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		for nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	for i := 0; i < len(nums); i++ {
		if i+1 != nums[i] {
			res = append(res, nums[i])
		}
	}
	return res
}

// @lc code=end
