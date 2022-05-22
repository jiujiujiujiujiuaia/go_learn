package linkedlist

/*
 * @lc app=leetcode.cn id=160 lang=golang
 *
 * [160] 相交链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//多和面试官进行交流，（1）链表是否有环？（2）可以用额外的空间还是常数空间？（3）可以遍历链表几次

//2020/3/15感想 简单在纸上画下把链表A接到链表B后面 链表B接到链表A上面
//这样如果有交点，那么两段路程终点是一样得！

//思想：如果存在相交的点那么意味着这个点距离最后一个点的距离是相同的，
//因此把两段路程合并，也就是a和b走的路程一样，相交的点同一位置，即可判断
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pointA, pointB := headA, headB
	for pointA != pointB {
		if pointA == nil {
			pointA = headB
		} else {
			pointA = pointA.Next
		}

		if pointB == nil {
			pointB = headA
		} else {
			pointB = pointB.Next
		}
	}
	return pointA
}

// @lc code=end
