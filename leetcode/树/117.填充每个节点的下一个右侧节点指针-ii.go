package main

/*
 * @lc app=leetcode.cn id=117 lang=golang
 *
 * [117] 填充每个节点的下一个右侧节点指针 II
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

//题目要求的是o(1)空间,要不然用队列层序遍历更简单
//这道题和116不同的点在于是，这道题是非完美二叉树，也就是有很多异常情况
func connect(root *Node) *Node {

}

// @lc code=end
