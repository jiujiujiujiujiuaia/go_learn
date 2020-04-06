package main

import (
	"container/list"
)

/*
 * @lc app=leetcode.cn id=103 lang=golang
 *
 * [103] 二叉树的锯齿形层次遍历
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

//2020/3/10

//解法一 借用了map的额外空间
func zigzagLevelOrder1(root *TreeNode) [][]int {
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
		if key%2 != 0 {
			val = reverse(val)
		}
		res[key] = append(res[key], val...)
	}
	return res
}

func reverse(array []int) []int {
	for i, j := 0, len(array)-1; i < j; {
		array[i], array[j] = array[j], array[i]
		i++
		j--
	}
	return array
}

//解法二 省去了那部分额外的空间
func zigzagLevelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	queue := list.New()
	queue.PushBack(root)
	//1.从第一层level开始
	queue.PushBack(1)
	for queue.Len() != 0 {
		node := queue.Front().Value.(*TreeNode)
		level := queue.Front().Next().Value.(int)
		queue.Remove(queue.Front())
		queue.Remove(queue.Front())
		if node != nil {
			res = insertArray(res, node.Val, level)
			level++
			queue.PushBack(node.Left)
			queue.PushBack(level)
			queue.PushBack(node.Right)
			queue.PushBack(level)
		}
	}
	return res
}

func insertArray(res [][]int, val, level int) [][]int {
	if level > len(res) {
		newLevel := []int{val}
		res = append(res, newLevel)
	} else if level%2 != 0 {
		res[level-1] = append(res[level-1], val)
	} else {
		res[level-1] = insertFontHeadVal(res[level-1], val)
	}
	return res
}

//这里的操作就是从头节点插入，因此可以使用双端队列减少时间消耗
func insertFontHeadVal(res []int, val int) []int {
	newRes := make([]int, len(res)+1)
	for i := 0; i < len(res); i++ {
		newRes[i+1] = res[i]
	}
	newRes[0] = val
	return newRes
}

// @lc code=end
