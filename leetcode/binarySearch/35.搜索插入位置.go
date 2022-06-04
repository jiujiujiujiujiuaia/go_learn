package binarySearch

/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 */

// @lc code=start

//2020/3/10
func searchInsert(nums []int, target int) int {
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
