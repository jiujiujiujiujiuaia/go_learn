package binarySearch

/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 */

// @lc code=start

//2022/7/26
func searchInsert(nums []int, target int) int {
	leftBoundary, rightBoundary := 0, len(nums)-1
	for leftBoundary <= rightBoundary {
		mid := leftBoundary + (rightBoundary-leftBoundary)/2
		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			leftBoundary = mid + 1
		} else {
			rightBoundary = mid - 1
		}
	}

	//三种情况
	//1.target > 所有元素，expected 返回len(nums)
	//2.target < 所有元素，expected 返回0
	//3.target 在nums[l],nums[r]之间，如果存在，返回mid
	//如果不存在那么target肯定是在某两个元素之间,比如target=2.5 nums=[1,2,3,4]
	//然后最后某次循环l=2,r=3，然后target > nums[l],l + 1 = r
	//然后target < nums[r], r - 1,循环退出，所以可以看到应该返回l
	return leftBoundary
}

//2020/3/10
func searchInsert1(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	//循环的终止条件是l = r + 1因此
	//1.当target比num所有都小时，r=l-1=-1 l = 0
	//2.当target比num所有都大时  l=r+1=len(num)
	//3.当targget在num范围中但是不在num时，也是l，可以进行举例
	return l
}

// @lc code=end
