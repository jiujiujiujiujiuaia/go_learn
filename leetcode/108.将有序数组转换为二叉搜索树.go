package main

/*
 * @lc app=leetcode.cn id=108 lang=golang
 *
 * [108] 将有序数组转换为二叉搜索树
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//这个可以和面试官交流下，一个数组可能有多种二叉搜索的可能
func sortedArrayToBST(nums []int) *TreeNode {
	return helpSortedArrayToBST(nums, 0, len(nums)-1)
}

func helpSortedArrayToBST(nums []int, l, r int) *TreeNode {
	if l > r {
		return nil
	}

	mid := l + (r-l)/2
	root := new(TreeNode)
	root.Val = nums[mid]
	root.Left = helpSortedArrayToBST(nums, l, mid-1)
	root.Right = helpSortedArrayToBST(nums, mid+1, r)
	return root
}

// @lc code=end
