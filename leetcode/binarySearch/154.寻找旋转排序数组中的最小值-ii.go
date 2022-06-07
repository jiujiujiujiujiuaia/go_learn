package binarySearch

/*
 * @lc app=leetcode.cn id=154 lang=golang
 *
 * [154] 寻找旋转排序数组中的最小值 II
 */

// @lc code=start

//这个和153相比，有重复的情况数字的情况

//边界情况 ：[1,1,1,1,1] [1,2,3] [1,0,1,1,1] [1]
//（1）第一种情况就退化到了o(n) （2）第二种情况是原地旋转 （3）第三种是无法区分哪段有序
//（4）

func findMin2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	//153 >= 可以处理掉nums=[1]得情况，但是154有重复元素，可能有[3,1,3]得情况
	if nums[len(nums)-1] > nums[0] {
		return nums[0]
	}

	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + ((r - l) >> 1)

		//这里要判断边界得情况，由于元素可能重复得原因
		if mid+1 < len(nums) && nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}

		if mid-1 > -1 && nums[mid-1] > nums[mid] {
			return nums[mid]
		}

		if nums[mid] == nums[l] {
			l++
			continue
		}

		if nums[mid] > nums[l] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	//注意哦这里不能返回-1，因为有[1,1,1,1]这个情况
	return nums[l]
}
