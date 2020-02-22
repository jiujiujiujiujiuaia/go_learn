package main

/*
 * @lc app=leetcode.cn id=99 lang=golang
 *
 * [99] 恢复二叉搜索树
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
var treeNodeArray []int

func recoverTree(root *TreeNode) {
	treeNodeArray := make([]int, 0)
	transformTreeToArray(root)
	wrongIndex := make([]int, 0)

	//如果把一个有序数组中的两个交换位置的坐标找出来？？
	for i := 1; i < len(treeNodeArray); i++ {
		if treeNodeArray[i] <= treeNodeArray[i-1] {

		}
	}

}

func transformTreeToArray(root *TreeNode) {
	if root == nil {
		return
	}

	transformTreeToArray(root.Left)
	treeNodeArray = append(treeNodeArray, root.Val)
	transformTreeToArray(root.Right)
}

// @lc code=end
