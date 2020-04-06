package main

/*
 * @lc app=leetcode.cn id=448 lang=golang
 *
 * [448] 找到所有数组中消失的数字
 */

// @lc code=start

//同理287 寻找重复的数字,这个和那道题的不同时，有多个消失/重复的数字的时候，可以返回
//多个结果
func findDisappearedNumbers(nums []int) []int {
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		for nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	for i := 0; i < len(nums); i++ {
		if i+1 != nums[i] {
			res = append(res, i+1)
		}
	}
	return res
}

// @lc code=end
