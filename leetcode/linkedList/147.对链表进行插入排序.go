package linkedlist

import (
	"math"
)

/*
 * @lc app=leetcode.cn id=147 lang=golang
 *
 * [147] 对链表进行插入排序
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//题目要求是原地排序，并且有负数的存在
func insertionSortList(head *ListNode) *ListNode {
	dummyHead := &ListNode{Val: math.MinInt32, Next: nil}
	current := head
	length := 1
	//外层循环增加元素
	for current != nil {
		//内层循环找位置
		//相当于有两个链表，一个是已经排好序得，另一个是未排序得
		frontNode := dummyHead
		var next *ListNode
		for i := 0; i < length; i++ {
			//链表一定能找到自己的位置
			if (current.Val >= frontNode.Val && i == length-1) ||
				(current.Val >= frontNode.Val &&
					current.Val <= frontNode.Next.Val) {
				next = swapNode(current, frontNode)
				break
			}
			frontNode = frontNode.Next
		}
		current = next
		length++
	}
	return dummyHead.Next
}

func swapNode(insertNode, frontInsertNode *ListNode) *ListNode {
	next := frontInsertNode.Next
	frontInsertNode.Next = insertNode
	insertNodeNext := insertNode.Next
	insertNode.Next = next
	return insertNodeNext
}

// @lc code=end
