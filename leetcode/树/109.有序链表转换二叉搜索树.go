package main

/*
 * @lc app=leetcode.cn id=109 lang=golang
 *
 * [109] 有序链表转换二叉搜索树
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//解法一：先把链表转成数组，用数组的方法做 时间复杂度 o(n)，但是使用了额外的数组空间

//解法二：递归链表 时间复杂度是 o(nlogn)
func sortedListToBST1(head *ListNode) *TreeNode {
	return helpSortedListToBST(head)
}

func helpSortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	var pre *ListNode
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	root := new(TreeNode)
	root.Val = slow.Val

	//(1)如果只有一个点
	if pre == nil {
		return root
	}

	//(2)如果大于一个点，则需要把链表切一次，丢弃中间点
	next := slow.Next
	pre.Next = nil

	root.Left = helpSortedListToBST(head)
	root.Right = helpSortedListToBST(next)

	return root
}

//解法三：模拟中序遍历的过程 时间复杂度位0(n)
var next *ListNode

func sortedListToBST(head *ListNode) *TreeNode {
	next = head
	length := getLength(head)
	return help(0, length-1)
}

func help(start, end int) *TreeNode {
	if start > end {
		return nil
	}

	mid := start + (end-start)>>1
	left := help(start, mid-1)

	root := &TreeNode{Val: next.Val}
	root.Left = left

	next = next.Next

	right := help(mid+1, end)
	root.Right = right

	return root
}

func getLength(head *ListNode) int {
	res := 0
	for head != nil {
		head = head.Next
		res++
	}
	return res
}

// @lc code=end
