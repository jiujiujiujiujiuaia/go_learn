package main

/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 *
 * https://leetcode-cn.com/problems/search-in-rotated-sorted-array/description/
 *
 * algorithms
 * Medium (36.22%)
 * Likes:    434
 * Dislikes: 0
 * Total Accepted:    58.3K
 * Total Submissions: 161.1K
 * Testcase Example:  '[4,5,6,7,0,1,2]\n0'
 *
 * 假设按照升序排序的数组在预先未知的某个点上进行了旋转。
 *
 * ( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。
 *
 * 搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。
 *
 * 你可以假设数组中不存在重复的元素。
 *
 * 你的算法时间复杂度必须是 O(log n) 级别。
 *
 * 示例 1:
 *
 * 输入: nums = [4,5,6,7,0,1,2], target = 0
 * 输出: 4
 *
 *
 * 示例 2:
 *
 * 输入: nums = [4,5,6,7,0,1,2], target = 3
 * 输出: -1
 *
 */

// @lc code=start
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	var b int
	if nums[len(nums)-1] > nums[0] {
		b = 0
	} else {
		indexNew := findMax(nums)
		indexOld := len(nums) - 1
		//indexOld + b -len(nums) = indexNew
		//b代表偏移量
		b = indexNew + len(nums) - indexOld
		if b > len(nums) {
			b = b % len(nums)
		}
	}

	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		index := (mid + b) % len(nums)
		if nums[index] == target {
			return index
		} else if nums[index] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return -1
}

func findMax(nums []int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] > nums[mid+1] {
			return mid
		}

		if nums[mid-1] > nums[mid] {
			return mid - 1
		}

		if nums[mid] > nums[len(nums)-1] {
			l = mid
		} else {
			r = mid
		}
	}
	return -1
}

// @lc code=end
