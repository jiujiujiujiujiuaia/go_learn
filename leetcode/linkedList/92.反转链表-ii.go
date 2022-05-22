package linkedlist

/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	dummyHead := new(ListNode)
	dummyHead.Next = head
	cur := dummyHead
	firstReverse := new(ListNode)
	cnt := 1
	for cur != nil {
		//（1）因为多加了一个哑节点，cur是翻转序列的前一个节点
		if cnt == m {
			//first是没有翻转的第一个节点，firstNotReverse是翻转后的最后一个节点
			firstReverse = cur.Next
			lastReverse, lastReverseNext := reverseList(cur.Next, n-m)
			firstReverse.Next = lastReverseNext
			cur.Next = lastReverse
			break
		}

		//(3)没有就往后遍历
		cnt++
		cur = cur.Next
	}
	return dummyHead.Next
}

//翻转m至n区间的链表,重点是返回的是：
//（1）第n个节点，也就是翻转后的头（2）第n个节点的后面节点
func reverseList(head *ListNode, cnt int) (*ListNode, *ListNode) {
	var next, pre *ListNode
	cur := head
	for cnt >= 0 {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
		cnt--
	}
	// fmt.Println("new head is ", pre.Val)
	// fmt.Println("next is ", cur.Val)
	return pre, cur
}

// @lc code=end
