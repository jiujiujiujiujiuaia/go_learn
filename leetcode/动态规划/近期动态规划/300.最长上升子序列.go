package main

import (
	"fmt"
	"strconv"
)

/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长上升子序列
 */

// @lc code=start

//面试真题，面试官不仅要求我输出最长的值，还要把这个子串输出出来，怎么做？
//再加一个dp数组的去记录这个状态，表示0-i这个范围内最长的
func advance(nums []int) string {
	strDp := make([]string, len(nums)+1)
	dp := make([]int, len(nums)+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = 1
	}

	var maxNumber int = 0
	var maxLongestString string = ""
	for i := 1; i <= len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i-1] > nums[j] {
				if dp[i] < dp[j+1]+1 {
					dp[i] = dp[j+1] + 1
					strDp[i] = strDp[j+1] + strconv.Itoa(nums[i-1])
				}
			}
		}
		if dp[i] > maxNumber {
			maxLongestString = strDp[i]
			maxNumber = dp[i]
		}
	}
	return maxLongestString
}

//解法一：时间复杂度为o（n^2）的解法
//dp[i]状态表示为0-i的最长子序列的长度
func lengthOfLIS1(nums []int) int {
	dp := make([]int, len(nums)+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = 1
	}

	var maxNumber int = 0
	//1.dp[i]状态表示为前i个数的最长子序列的个数
	//当表示前多少个的时候，前0就是没有，前1个就是数组0索引的那个元素
	//因此dp数组的下标索引和num索引差1

	//2.做了一下午题目，我发现这一类双层循环的问题
	//就是那种在dp(n-1)的情况下，你添加了一个新的状态or元素，那么这个元素需要和前面的n-1
	//元素可能会产生多个状态，因此才会有双层循环的解法
	for i := 1; i <= len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i-1] > nums[j] {
				dp[i] = max(dp[i], dp[j+1]+1)
			}
		}
		if dp[i] > maxNumber {
			maxNumber = dp[i]
		}
	}
	return maxNumber
}

//解法二 复杂度o(nlogn) 改变了状态的定义
//这是一种贪心的思路，假设有两个长度一样的最长序列，我们肯定
//想要增长的慢的序列，例如1，2，3 1，100 ，200 我们肯定选择前者
//因此维护这样一个数组，他是递增的并且尾巴尽量的小
//如果大于尾巴 直接插入
//如果小于尾巴，则找到他在这个数组的位置
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	//tails[i]表示前i为最长子序列中最小的尾数
	tails := make([]int, 1)
	tails[0] = nums[0]

	//如果当前元素大于tails最后一个元素，则直接插入
	//如果小于，二分查找大于该元素中最小的插入进入
	for i := 1; i <= len(nums); i++ {
		if nums[i-1] > tails[len(tails)-1] {
			tails = append(tails, nums[i-1])
			continue
		}

		l, r := 0, len(tails)-1

		//这个插入的位置可以同35题
		for r >= l {
			mid := l + (r-l)/2
			if nums[i-1] > tails[mid] {
				l = mid + 1
			} else if nums[i-1] < tails[mid] {
				r = mid - 1
			} else {
				l = mid
				break
			}
		}

		//为什么是l，因为该元素最后的状态是小于等于某个数，那么就是r=mid-1,离该
		//元素最近且大于他的就是l下标
		tails[l] = nums[i-1]
	}
	fmt.Println(tails)
	return len(tails)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end
