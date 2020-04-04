/*
 * @lc app=leetcode.cn id=80 lang=golang
 *
 * [80] 删除排序数组中的重复项 II
 */

// @lc code=start

//这里的思想是双指针，一个指针指向待写入的位置，一个指针遍历数组
//这道题和前面那道题的区别在于，写入的条件变成了：重复小于2次
//1.也就是说如果i和j相等，且j==j-1，那么不可以
//2.如果j != j-1 i为任何值都可以插入
//才能写入
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1{
		return 1
	}

	j := 1
	for i := 2; i < len(nums);i++ {
		if nums[j] != nums[j-1] {
			j++
			nums[i], nums[j] = nums[j], nums[i]
		}else if nums[j] == nums[j-1] && nums[j] != nums[i]{
			j++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	return j + 1
}

// @lc code=end

