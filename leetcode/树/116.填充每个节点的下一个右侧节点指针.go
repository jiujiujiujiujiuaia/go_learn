package main

/*
 * @lc app=leetcode.cn id=116 lang=golang
 *
 * [116] 填充每个节点的下一个右侧节点指针
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

//解法一：递归
func connect1(root *Node) *Node {
	if root == nil {
		return nil
	}

	//这里是重点，除了要管自己的左右节点相连接外
	//还要管自己右节点连接到平级节点的左节点
	//换句话说上一层管下一层
	if root.Left != nil {
		root.Left.Next = root.Right
		if root.Next != nil {
			root.Right.Next = root.Next.Left
		}
	}

	connect(root.Left)
	connect(root.Right)
	return root
}

//解法二：迭代，简单清晰些
//相当于说用o(1)空间去实现层序遍历（有next指针才可以）
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	pre, cur := root, root
	for pre.Left != nil {
		cur = pre
		for cur != nil {
			if cur.Left != nil {
				cur.Left.Next = cur.Right
			}
			if cur.Next != nil {
				cur.Right.Next = cur.Next.Left
			}
			cur = cur.Next
		}
		pre = pre.Left
	}
	return root
}

//解法三：迭代，没有上一种简单清晰
//相当于说用o(1)空间去实现层序遍历（有next指针才可以）
func connect2(root *Node) *Node {
	if root == nil {
		return root
	}

	var cur *Node
	start, pre := root, root
	for pre.Left != nil {
		//当遍历到一行末尾时，需要转到下一行
		if cur == nil {
			pre.Left.Next = pre.Right

			pre = start.Left
			cur = start.Right
			start = start.Left
			//将下一层的所有next链接起来
		} else {
			//上一层管下一层
			pre.Left.Next = pre.Right
			pre.Right.Next = cur.Left

			pre = pre.Next
			cur = cur.Next
		}
	}
	return root
}

// @lc code=end
