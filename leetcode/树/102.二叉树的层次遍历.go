package main

import (
	"container/list"
)

/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层次遍历
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

//解法一 使用了map额外空间来储存
func levelOrder1(root *TreeNode) [][]int {
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
	for key, val := range valToLevelMap {
		res[key] = append(res[key], val...)
	}
	return res
}

//解法二 去掉了map的额外空间，遍历的过程中就可以实现数据的存储
func levelOrder(root *TreeNode) [][]int {
	res := make([][]int,0)
	queue := list.New()
	queue.PushBack(root)
	queue.PushBack(1)
	for queue.Len() != 0 {
		node := queue.Front().Value.(*TreeNode)
		level := queue.Front().Next().Value.(int)
		queue.Remove(queue.Front())
		queue.Remove(queue.Front())
		if node != nil {
			res = insertArray(res,node.Val,level)
			level++
			queue.PushBack(node.Left)
			queue.PushBack(level)
			queue.PushBack(node.Right)
			queue.PushBack(level)
		}
	}
	return res
}

func insertArray(res [][]int,val,level int)[][]int{
	if level > len(res){
		newLevel := []int{val}
		res = append(res,newLevel)
	}else{
		res[level-1]  = append(res[level-1],val)
	}
	return res
}
// @lc code=end
