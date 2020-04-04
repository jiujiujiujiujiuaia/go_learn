package main

/*
 * @lc app=leetcode.cn id=106 lang=golang
 *
 * [106] 从中序与后序遍历序列构造二叉树
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

//同105
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: postorder[len(postorder)-1]}
	rootIndex := find(postorder[len(postorder)-1], inorder)
	root.Left = buildTree(inorder[0:rootIndex], postorder[0:rootIndex])
	root.Right = buildTree(inorder[rootIndex+1:], postorder[rootIndex:len(postorder)-1])
	return root
}

func find(target int, inorder []int) int {
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == target {
			return i
		}
	}
	return -1
}

// @lc code=end
