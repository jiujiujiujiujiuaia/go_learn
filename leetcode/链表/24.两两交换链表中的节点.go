/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func swapPairs(head *ListNode) *ListNode {
    dummyHead := new(ListNode)
    dummyHead.Next = head
    cur,pre := head,dummyHead
    for cur != nil {
        next := cur.Next
        if next == nil{
            break
        }

        pre.Next = next
        cur.Next = next.Next
        next.Next = cur

        pre = cur
        cur = cur.Next
    }
    return dummyHead.Next
}
// @lc code=end

