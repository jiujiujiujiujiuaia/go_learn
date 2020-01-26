package main

/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 *
 * https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/description/
 *
 * algorithms
 * Medium (38.23%)
 * Likes:    259
 * Dislikes: 0
 * Total Accepted:    48K
 * Total Submissions: 125.5K
 * Testcase Example:  '[5,7,7,8,8,10]\n8'
 *
 * 给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
 *
 * 你的算法时间复杂度必须是 O(log n) 级别。
 *
 * 如果数组中不存在目标值，返回 [-1, -1]。
 *
 * 示例 1:
 *
 * 输入: nums = [5,7,7,8,8,10], target = 8
 * 输出: [3,4]
 *
 * 示例 2:
 *
 * 输入: nums = [5,7,7,8,8,10], target = 6
 * 输出: [-1,-1]
 *
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	l,r := 0,len(nums)-1
	for l<=r{
		mid := l + (r-l)/2
		if nums[mid] == target{
			return handleEqualCondition(nums,target,mid)
		}else if nums[mid] > target{
			r = mid - 1
		}else{
			l = mid + 1
		}
	}

	return []int{-1,-1}
}

func handleEqualCondition(nums []int, target,index int)[]int{
	l,h := index,index
	for h < len(nums) - 1 && nums[h+1] == target{
		h++
	}

	for l>0 && nums[l-1] == target{
		l--
	}
	return []int{l,h}
}

// @lc code=end
