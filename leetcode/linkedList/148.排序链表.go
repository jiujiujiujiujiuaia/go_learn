package main

/*
 * @lc app=leetcode.cn id=148 lang=golang
 *
 * [148] 排序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//归并排序的算法最重要的思想：先分治，再合并，值得注意的是，合并的时候
//合并的是两个有序的链表or数组

//注意归并排序的时间复杂度计算：n个节点能够被分为logn层
//每一层的时间是n，因此是 logn * n
func sortList(head *ListNode) *ListNode {
	return divide(head)
}

func divide(head *ListNode) *ListNode {
	//（1）当只剩下一个节点的时候返回
	if head == nil || head.Next == nil {
		return head
	}
	dummyHead := new(ListNode)
	dummyHead.Next = head
	//（2）快慢指针寻找分界点
	slowPoint, fastPoint := dummyHead, dummyHead
	for fastPoint != nil && fastPoint.Next != nil {
		slowPoint = slowPoint.Next
		fastPoint = fastPoint.Next.Next
	}
	//（3）最重要的是：找到分界点后要把链表切开
	copySlowPoint := slowPoint.Next
	slowPoint.Next = nil
	frontHead := divide(dummyHead.Next)
	backHead := divide(copySlowPoint)

	//（4）合并两个有序链表
	newHead := mergeTwoLists(frontHead, backHead)
	return newHead
}

//func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
//	newHeadNode := new(ListNode)
//	dummyNode := newHeadNode
//	for l1 != nil || l2 != nil {
//		if l1 == nil {
//			newHeadNode.Next = l2
//			break
//		}
//
//		if l2 == nil {
//			newHeadNode.Next = l1
//			break
//		}
//
//		if l1.Val > l2.Val {
//			newHeadNode.Next = l2
//			l2 = l2.Next
//		} else {
//			newHeadNode.Next = l1
//			l1 = l1.Next
//		}
//
//		newHeadNode = newHeadNode.Next
//	}
//	return dummyNode.Next
//}

// @lc code=end
