package main

/*
 * @lc app=leetcode.cn id=450 lang=golang
 *
 * [450] 删除二叉搜索树中的节点
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
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == key {
		//包含了左右子节点均为空，左空右不空，右空左不空的情况
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}

		//左右都不空
		//从左边找到最小的节点
		minNode := get(root.Right)
		root.Val = minNode.Val
		//然后把这个最小节点删除，这个最小节点一定是一个叶子节点
		root.Right = deleteNode(root.Right, minNode.Val)
	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else {
		root.Right = deleteNode(root.Right, key)
	}
	return root
}

func get(root *TreeNode) *TreeNode {
	for root.Left != nil {
		root = root.Left
	}
	return root
}

// @lc code=end
