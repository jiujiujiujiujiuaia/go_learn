package main

import "container/list"

/*
 * @lc app=leetcode.cn id=107 lang=golang
 *
 * [107] 二叉树的层次遍历 II
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

//同样看有没有更好的办法

//更好的办法详情见102
func levelOrderBottom(root *TreeNode) [][]int {
	valToLevelMap := make(map[int][]int, 0)
	queue := list.New()
	queue.PushBack(root)
	queue.PushBack(0)
	for queue.Len() != 0 {
		node := queue.Front().Value.(*TreeNode)
		index := queue.Front().Next().Value.(int)
		queue.Remove(queue.Front())
		queue.Remove(queue.Front())
		if node != nil {
			valToLevelMap[index] = append(valToLevelMap[index], node.Val)
			index++
			queue.PushBack(node.Left)
			queue.PushBack(index)
			queue.PushBack(node.Right)
			queue.PushBack(index)
		}
	}

	res := make([][]int, len(valToLevelMap))
	length := len(valToLevelMap) - 1
	for key, val := range valToLevelMap {
		res[length-key] = append(res[length-key], val...)
	}
	return res
}

// @lc code=end
