package util

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateLinkedList(data []int) *ListNode {
	head := new(ListNode)
	dummyHead := head
	for i := 0; i < len(data); i++ {
		node := new(ListNode)
		node.Val = data[i]
		head.Next = node
		head = node
	}
	return dummyHead.Next
}

func CompareTwoLinkedList(head1 *ListNode, head2 *ListNode) bool {
	for head1 != nil && head2 != nil {
		if head1.Val == head2.Val {
			head1 = head1.Next
			head2 = head2.Next
			continue
		} else {
			return false
		}
	}

	if head1 == nil && head2 == nil {
		return true
	} else {
		return false
	}
}

func TransformLinkedListToArray(head *ListNode) []int {
	result := []int{}
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}
