package main

import "container/list"

/*
 * @lc app=leetcode.cn id=662 lang=golang
 *
 * [662] 二叉树最大宽度
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

//已经ac

type TreeNodeWithDepth struct {
	Node     *TreeNode
	Depth    int
	Position int
}

func widthOfBinaryTree(root *TreeNode) int {
	queue := list.New()
	queue.Push()
}

// @lc code=end
