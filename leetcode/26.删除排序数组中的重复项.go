package main

/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除排序数组中的重复项
 */

// @lc code=start

//一些case
//[1,1,1,1]
//[]
//[1,1]
//[1,2,3,4]

//这里的思想是双指针，一个指针指向待写入的位置，一个指针遍历数组
//可以写入的条件是当前位置和遍历位置的值不相同，才可以写入
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	j := 0

	for i := 1; i < len(nums); {
		//如果发现了有一个数和当前不同，那么就把下一个坑给这个不同的数
		//并且下一次以他作为锚
		if nums[i] != nums[j] {
			j++
			nums[i], nums[j] = nums[j], nums[i]
		}
		i++
	}
	return j + 1
}

// @lc code=end
