package linkedlist

import (
	"go_learn/leetcode/util"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	tests := []struct {
		name           string
		input1, input2 []int
		expected       []int
	}{
		{"normal cases with same length", []int{1, 2, 3}, []int{1, 2, 3}, []int{1, 1, 2, 2, 3, 3}},
		{"normal cases with different length", []int{1, 2, 3, 4}, []int{1, 2, 3}, []int{1, 1, 2, 2, 3, 3, 4}},
		{"normal cases with different length", []int{1, 2, 3}, []int{1, 2, 3, 4}, []int{1, 1, 2, 2, 3, 3, 4}},
		{"nil node", []int{}, []int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
		{"nil node", []int{}, []int{}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := util.CreateLinkedList(tt.input1)
			l2 := util.CreateLinkedList(tt.input2)
			expected := util.CreateLinkedList(tt.expected)

			got := mergeTwoLists(l1, l2)
			gotArray := util.TransformLinkedListToArray(got)
			if !util.CompareTwoLinkedList(expected, got) {
				t.Errorf("name=%v failed, expected=%v, got=%v", tt.name, tt.expected, gotArray)
			}
		})
	}
}
