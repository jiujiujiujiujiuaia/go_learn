package main

import "sort"

/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 多数元素
 */

// @lc code=start
//思路：这个题目求超过一半的数，把这堆数分成两堆，数a一定是某一堆的超过一半的数

//边界情况：（1）确定数组是否一定有多数？

//解法一：暴力解法，双重循环。 每遍历一个数，就统计这个数出现了多少次 o(n^2)

//解法二：哈希，利用额外的空间记录下来 然后统计最高频率的数 o(n)

//解法三：先排序，由于超过了数组的一半，因此length/2一定是该数 o(nlogn)
func majorityElement1(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

//解法四：分治法,因为是求大于一半的数，因此分成相等的两堆肯定是有一个该数 o(nlogn)
func majorityElement(nums []int) int {
	if len(nums) == 0{
		return 0
	}
	return helpMajorityElement(nums,0,len(nums)-1)
}

func helpMajorityElement(nums []int,l,r int)int{
	if l == r {
		return nums[l]
	}

	mid := l + (r - l ) >> 1
	left := helpMajorityElement(nums,l,mid)
	right := helpMajorityElement(nums,mid+1,r)

	if left == right{
		return left
	}
	leftCnt := countNum(nums,left)
	rightCnt := countNum(nums,right)
	
	if leftCnt > rightCnt{
		return left
	}else{
		return right
	}

}

func countNum(nums []int,target int)int{
	cnt := 0
	for _,num := range nums{
		if num == target{
			cnt++
		}
	}
	return cnt
}

//解法5 还有一种o(n)的投票器算法 很有意思

// @lc code=end
