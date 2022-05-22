package linkedlist

/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 2/25

//这个题目的前置题就是翻转链表
func reverseKGroup(head *ListNode, k int) *ListNode {
	//（1）哑节点的作用就是为了方便头节点翻转
	dummyHead := new(ListNode)
	dummyHead.Next = head
	cur, pre := head, dummyHead

	//（2）先算出总长，方便计算哪一段不需要翻转
	length := 0
	for cur != nil {
		cur = cur.Next
		length++
	}

	cnt := length / k

	//（3）需要翻转的最后一个节点变成了第一个，pre.Next = last 然后替换pre和cur
	cur = head
	for cnt > 0 {
		lastReverse, lastReverseNext := helpReverseList(cur, k)
		pre.Next = lastReverse
		//cur是未翻转的头节点
		pre = cur
		cur = lastReverseNext
		cnt--
	}
	pre.Next = cur
	return dummyHead.Next
}

//这个函数返回的是翻转后链表的头节点和未翻转前的下一个
func helpReverseList(head *ListNode, cnt int) (*ListNode, *ListNode) {
	var next, pre *ListNode
	cur := head
	for cnt > 0 {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
		cnt--
	}
	return pre, cur
}

// @lc code=end
