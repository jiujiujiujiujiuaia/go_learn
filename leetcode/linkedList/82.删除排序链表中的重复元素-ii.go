package main

/*
 * @lc app=leetcode.cn id=82 lang=golang
 *
 * [82] 删除排序链表中的重复元素 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//func deleteDuplicates(head *ListNode) *ListNode {
//	dummyHead := new(ListNode)
//	dummyHead.Next = head
//	dummyHead.Val = math.MaxInt32
//	current, pre := dummyHead, dummyHead
//	for current != nil && current.Next != nil {
//		if current.Val == current.Next.Val {
//			val := current.Val
//
//			//只要用了某指针的下一个或值，都要先判断该指针是不是nil
//			for current != nil && current.Val == val {
//				pre.Next = current.Next
//				current = pre.Next
//			}
//		} else {
//			pre = current
//			current = current.Next
//		}
//	}
//	return dummyHead.Next
//}

// @lc code=end
