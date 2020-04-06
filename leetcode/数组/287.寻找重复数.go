package main

/*
 * @lc app=leetcode.cn id=287 lang=golang
 *
 * [287] 寻找重复数
 *
 */

// @lc code=start

//2020/4/2 突然发现这道题目有三个要求：
//1.空间复杂度为1
//2.不允许更改原数组
//3.时间复杂度小于n^2

//解法一 这道题如果对时间复杂度没有要求，可以使用先排序再找
//优点：ologn 没有使用额外空间 缺点 数组原地修改了

//解法二 使用额外空间的集合数据结构进行修改

//解法三 把这个题目转换成环状数组（由于即不允许额外空间，还要求原地数组）
//解法三才是最终的答案

//解法四
//对于一个长度为n 范围是[1,n-1]，如果这个数组中没有重复元素，最终所有的元素都会归位
//如果有重复元素，最终重复元素会指向已经归为的元素发现重复了
func findDuplicate4(nums []int) int {
	for i := 0; i < len(nums); i++ {
		//1.如果索引i指向的元素不是i，那么就让nums[i]归为
		//eg[1,0] 索引0指向的是1 因此和索引1的值交换，让nums[i]回到了原有的位置
		for nums[i] != i {
			//2.如果出现了某个元素指向的那个位置已经是归为了，说明这个元素重复了
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}
	return -1
}

//解法三
func findDuplicate(nums []int) int {
	slow, fast := 0, 0
	//1.预处理满指针走一格，快指针走两格
	slow = nums[slow]
	fast = nums[nums[fast]]

	//2.可以看notability上的笔记，这个数组一定是环状的
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	//3.找到相交点后，开始像142题一样，
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return fast
}

// @lc code=end
