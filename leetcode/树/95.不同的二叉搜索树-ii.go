package main

/*
 * @lc app=leetcode.cn id=95 lang=golang
 *
 * [95] 不同的二叉搜索树 II
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


//前置题目是95，这里是递归方法做的，还可以用动态规划，有空补上
func generateTrees(n int) []*TreeNode {
	if n <= 0 {
		return []*TreeNode{}
	}
	res := generateTree(1, n)
	return res
}

//算法思想：以某个点为根，左右两边依次递归去获得能够构成二叉搜索树。
//然后回退保存下来

//典型的使用递归方法解题，只考虑当下！
func generateTree(start, end int) []*TreeNode {
	res := make([]*TreeNode, 0)
	if start > end {
		res = append(res, nil)
		return res
	}

	for i := start; i <= end; i++ {
		for _, leftTreeNode := range generateTree(start, i-1) {
			for _, rightTreeNode := range generateTree(i+1, end) {
				root := new(TreeNode)
				root.Val = i
				root.Left = leftTreeNode
				root.Right = rightTreeNode
				res = append(res, root)
			}
		}
	}
	return res
}

// @lc code=end
