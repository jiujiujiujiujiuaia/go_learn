package main

/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
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

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

var result []int

func inorderTraversal(root *TreeNode) []int {
	result = make([]int, 0)
	solve(root)
	return result
}

func solve(node *TreeNode) {
	if node == nil {
		return
	}
	solve(node.Left)
	result = append(result, node.Val)
	solve(node.Right)
}

// @lc code=end
