package linkedlist

/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	newHeadNode := new(ListNode)
	dummyNode := newHeadNode
	for l1 != nil || l2 != nil {
		if l1 == nil {
			newHeadNode.Next = l2
			break
		}

		if l2 == nil {
			newHeadNode.Next = l1
			break
		}

		if l1.Val > l2.Val {
			newHeadNode.Next = l2
			l2 = l2.Next
		} else {
			newHeadNode.Next = l1
			l1 = l1.Next
		}

		newHeadNode = newHeadNode.Next
	}
	return dummyNode.Next
}

// @lc code=end
