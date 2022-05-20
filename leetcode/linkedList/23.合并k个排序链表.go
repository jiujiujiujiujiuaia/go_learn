package main

/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个排序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//比较暴力的解法：遍历所有的链表，放入一个数组当中
//然后对这个数组进行快排，排完之后再遍历放入到链表当中
//缺点：额外的空间，快排的时间复杂度是nlogn，最坏的情况还可能要降低到n^2

//解法一：题目转换成合并两个有序链表，然后返回新链表的头，再合并下一个链表。
//时间复杂度：每个节点都会被比较k次（k是链表头的数量），然后N代表是链表总个数
//所以时间复杂度是NK
func mergeKLists(lists []*ListNode) *ListNode {
	var headNode *ListNode
	for _, list := range lists {
		headNode = mergeTwoLists(headNode, list)
	}
	return headNode
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

//解法二：把链表头都放入优先级队列中，然后把各个链表头放入队列中，
//只要队列不为空，就不短的delMin，然后把Min的下一个节点加入队列中
//时间复杂度：每一个节点，都要被比较logk，那么时间复杂度就是nlogk
// @lc code=end
