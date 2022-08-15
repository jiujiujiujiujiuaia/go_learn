package binarySearch

import "go_learn/leetcode/util"

/*
 * @lc app=leetcode.cn id=81 lang=golang
 *
 * [81] 搜索旋转排序数组 II
 */

// @lc code=start

//边界情况 ：[1,1,1,1,1] [1,2,3] [1,0,1,1,1] [1]
//（1）第一种情况就退化到了o(n) （2）第二种情况是原地旋转 （3）第三种是无法区分哪段有序
//（4）
//2022/08/15
//和33题相比，打破了mid在上半区或者下半区有序的设定，有的case无法判断上半区是否有序
func search4(nums []int, target int) bool {
	leftBoundary, rightBoundary := 0, len(nums)-1
	for leftBoundary <= rightBoundary {
		mid := leftBoundary + (rightBoundary-leftBoundary)/2
		if nums[mid] == target {
			return true
		}

		if nums[mid] == nums[leftBoundary] {
			util.PrintArrayItem(nums, leftBoundary, mid, true)
			leftBoundary++
			continue
		}

		if nums[leftBoundary] < nums[mid] {
			util.PrintArrayItem(nums, leftBoundary, mid, true)
			if target >= nums[leftBoundary] && target < nums[mid] {
				rightBoundary = mid - 1
			} else {
				leftBoundary = mid + 1
			}
		} else {
			util.PrintArrayItem(nums, mid, rightBoundary, false)
			if target > nums[mid] && target <= nums[rightBoundary] {
				leftBoundary = mid + 1
			} else {
				rightBoundary = mid - 1
			}
		}
	}
	return false
}

func search3(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return true
		}

		//这个题目可能有重复得元素，因此就不确定前半段有序还是后半段有序
		//[1,0,1,1,1]像这种情况
		//因此就l++，去掉这种特殊情况
		if nums[mid] == nums[l] {
			l++
			continue
		}

		//（1）num[mid]>num[l] l-mid这段有序 （2）否则 mid-r这段有序
		if nums[l] < nums[mid] {
			//(1)l-mid有序并且target在其中，则二分
			//否则这一段可以抛弃
			if target < nums[mid] && target >= nums[l] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return false
}

// @lc code=end
