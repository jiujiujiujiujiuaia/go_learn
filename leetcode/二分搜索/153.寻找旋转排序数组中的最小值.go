package main

/*
 * @lc app=leetcode.cn id=153 lang=golang
 *
 * [153] 寻找旋转排序数组中的最小值
 */

// @lc code=start
//边界情况：1.为什么说不能有重复数字？2.如果刚刚好原地不动怎么办?
//3.如果数组中只有一个元素怎么办，两个元素呢？从小规模的数据讨论

func findMin(nums []int) int {
	l, r := 0, len(nums)-1

	if nums[r] >= nums[l] {
		return nums[l]
	}
	for l <= r {
		mid := l + (r-l)/2

		//这里为什么不判断是否超出边界？
		//因为旋转后得数组最小值一定在数组得中间，而不是在两边0,和length
		//并且由于一定会有最小值，所以搜索区间不会到左右边界上
		if nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}

		if nums[mid-1] > nums[mid] {
			return nums[mid]
		}

		if nums[mid] > nums[l] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}

// @lc code=end
