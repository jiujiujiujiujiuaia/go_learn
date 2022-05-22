package linkedlist

import (
	"go_learn/leetcode/util"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	cases := []struct {
		Name            string
		input, expected []int
		indexOfDeleted  int
	}{
		{"Delete the last one", []int{1, 2, 3, 4}, []int{1, 2, 3}, 1},
		{"Delete the middle one", []int{1, 2, 3, 4}, []int{1, 3, 4}, 3},
		{"Delete first one", []int{1, 2, 3, 4}, []int{2, 3, 4}, 4},
		{"Delete first one when there is only one node", []int{1}, []int{}, 1},
		{"abnormal index", []int{1, 2, 3, 4}, []int{1, 2, 3, 4}, -1},
		{"the deleted index exceed linked list length  ", []int{1, 2, 3, 4}, []int{1, 2, 3, 4}, 100},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			head := util.CreateLinkedList(c.input)
			expected := util.CreateLinkedList(c.expected)
			if result := removeNthFromEnd(head, c.indexOfDeleted); !util.CompareTwoLinkedList(result, expected) {
				actual := util.TransformLinkedListToArray(result)
				t.Errorf("input=%v,indexOfDeleted=%v,expected=%v,actual=%v", c.input, c.indexOfDeleted, c.expected, actual)
			}
		})
	}
}
