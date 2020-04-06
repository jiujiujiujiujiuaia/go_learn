package main

/*
 * @lc app=leetcode.cn id=230 lang=golang
 *
 * [230] 二叉搜索树中第K小的元素
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
var res, cnt int

//二叉搜索树的经典中序遍历解题
func kthSmallest(root *TreeNode, k int) int {
	res, cnt = 0, 0
	helpKthSmallest(root, k)
	return res
}

func helpKthSmallest(root *TreeNode, k int) bool {
	if root == nil {
		return false
	}

	if helpKthSmallest(root.Left, k) {
		return true
	}
	cnt++
	if cnt == k {
		res = root.Val
		return true
	}
	if helpKthSmallest(root.Right, k) {
		return true
	}
	return false
}

// @lc code=end
