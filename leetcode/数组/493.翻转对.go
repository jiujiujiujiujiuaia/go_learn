package main

import "fmt"

/*
 * @lc app=leetcode.cn id=493 lang=golang
 *
 * [493] 翻转对
 */

// @lc code=start
func reversePairs(nums []int) int {
	return sort(nums, 0, len(nums)-1)
}

func sort(nums []int, start, end int) int {
	//注意这里是>= 因为start,end相等的情况会导致无限循环
	if start >= end {
		return 0
	}
	mid := start + (end-start)>>1
	//(1)分别计算 两个半组内 翻转对
	left := sort(nums, start, mid)
	right := sort(nums, mid+1, end)
	//(2)计算 两个半组 比较的翻转对
	cnt := calculateTwoArrayNumber(nums, start, mid, end)
	//(3)进行原地归并
	merge(nums, start, mid, end)
	return cnt + left + right
}

func merge(nums []int, start, mid, end int) {
	fmt.Println(start, end)
	temp := make([]int, len(nums))
	for i := start; i <= end; i++ {
		temp[i] = nums[i]
	}

	for i, j, k := start, mid+1, start; k <= end; k++ {
		//(1)前半个数组耗尽，后半个数组直接插入
		if i > mid {
			nums[k] = temp[j]
			j++
		} else if j > end {
			nums[k] = temp[i]
			i++
		} else if temp[i] > temp[j] {
			nums[k] = temp[j]
			j++
		} else {
			nums[k] = temp[i]
			i++
		}
	}
}

func calculateTwoArrayNumber(nums []int, start, mid, end int) int {
	i := start
	j := mid + 1
	cnt := 0
	for i <= mid && j <= end {
		//由于是排序好的，如果i>2*j,那么i后面的数也比j大
		if nums[i] > 2*nums[j] {
			cnt += mid - i + 1
			j++
		} else {
			i++
		}
	}
	return cnt
}

// @lc code=end
