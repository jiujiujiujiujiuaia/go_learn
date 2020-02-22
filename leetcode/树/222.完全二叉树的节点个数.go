package main

import (
	"fmt"
	"math"
)

/*
 * @lc app=leetcode.cn id=222 lang=golang
 *
 * [222] 完全二叉树的节点个数
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

//算法思想：完全二叉树是最后一层不满的树，即算出n-1层总数+第n层的节点数即可
//难点在于最后一层节点的计算，通过二分的方式，如果落在上半区就是右子树，如果下半区就是左子树
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root != nil && (root.Left == nil && root.Right == nil) {
		return 1
	}
	//(1)计算树的总高n
	depth := calculateTreeLength(root)
	exceptLastLevelTotalNode := math.Exp2(float64(depth-1)) - 1

	//(2)然后判断第n层有多少个节点，加上n-1层之前的节点
	lastLevelNodeTotal := math.Exp2(float64(depth - 1))
	for i := 0; i < int(lastLevelNodeTotal); i++ {
		if existNode(root, i, int(lastLevelNodeTotal), depth-1) {
			fmt.Println(i)
			exceptLastLevelTotalNode++
		}
	}
	return int(exceptLastLevelTotalNode)
}

func calculateTreeLength(root *TreeNode) int {
	depth := 0
	for root != nil {
		depth++
		root = root.Left
	}

	return depth
}

//这里加了一个count，就是第三层的节点只需要遍历两次可以找到，避免空节点的情况
func existNode(root *TreeNode, index, total, count int) bool {
	cnt := 0
	l, r := 0, total-1
	for l <= r && cnt < count {
		mid := l + (r-l)/2
		if index > mid {
			l = mid + 1
			root = root.Right
		} else if index <= mid {
			r = mid - 1
			root = root.Left
		}
		cnt++
	}
	return root != nil
}

// @lc code=end
