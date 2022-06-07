package binarySearch

/*
 * @lc app=leetcode.cn id=704 lang=golang
 *
 * [704] 二分查找
 */

// @lc code=start
func search(nums []int, target int) int {
	leftBoundary := 0
	rightBoundary := len(nums) - 1

	//1.为什么是小于等于不是小于? 如果是left < right, 在没有搜索到target并且跳出循环的时候，边界条件是left=right，且这个cases没有被搜索到 要在循环外再判断一次nums[left] == target ? left : -1;
	//如果是l <= r, 就能很好的把l==r的cases涵盖进去
	for leftBoundary <= rightBoundary {
		//防止溢出的写法
		mid := leftBoundary + (rightBoundary-leftBoundary)/2

		//2.为什么搜索不到的情况是mid+1,mid-1? 因为mid已经被排除了，紧接着排除mid临近的两个点了
		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			leftBoundary = mid + 1
		} else if target < nums[mid] {
			rightBoundary = mid - 1
		}
	}
	return -1
}

// @lc code=end
