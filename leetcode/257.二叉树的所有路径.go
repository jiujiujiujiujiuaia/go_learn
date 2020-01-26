package main

import "strconv"

/*
 * @lc app=leetcode.cn id=257 lang=golang
 *
 * [257] 二叉树的所有路径
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
var res []string

func binaryTreePaths(root *TreeNode) []string {
	preBinaryTreePathsHandle()
	help(root, "")
	return res
}

func help(root *TreeNode, path string) {
	if root == nil {
		return
	}

	valString := strconv.FormatInt(int64(root.Val), 10)
	path += valString

	if root.Left == nil && root.Right == nil {
		res = append(res, path)
	} else {
		//放在这个地方，妙！这样就不用处理最后一个节点问题
		path += "->"
		help(root.Left, path)
		help(root.Right, path)
	}
}

func preBinaryTreePathsHandle() {
	res = make([]string, 0)
}

// @lc code=end
