package binarySearch

/*
 * @lc app=leetcode.cn id=81 lang=golang
 *
 * [81] 搜索旋转排序数组 II
 */

// @lc code=start

//边界情况 ：[1,1,1,1,1] [1,2,3] [1,0,1,1,1] [1]
//（1）第一种情况就退化到了o(n) （2）第二种情况是原地旋转 （3）第三种是无法区分哪段有序
//（4）

func search(nums []int, target int) bool {
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
