package main

/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		if val == nums[i] && val == nums[j] {
			j--
		} else if val == nums[i] && val != nums[j] {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		} else {
			i++
		}
	}
	return i
}

// @lc code=end
