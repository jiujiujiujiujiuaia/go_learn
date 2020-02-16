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
 func mergeKLists(lists []*ListNode) *ListNode {
    var headNode *ListNode
    for _,list := range lists{
        headNode = mergeTwoLists(headNode,list)
    }
    return headNode
}

func helpMergeTwoLists(listNode1,listNode2 *ListNode) *ListNode{
    newHead := new(ListNode)
    dummyHead := newHead
    for listNode1 != nil || listNode2 != nil{
        if listNode1 == nil{
            dummyHead.Next = listNode2
            break
        }

        if listNode2 == nil{
            dummyHead.Next = listNode1
            break
        }

        if listNode1.Val < listNode2.Val{
            dummyHead.Next = listNode1
            listNode1.Next = listNode1
        }else{
            dummyHead.Next = listNode2
            listNode2.Next = listNode2
        }

        dummyHead = dummyHead.Next
    }

    return newHead.Next
}

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

