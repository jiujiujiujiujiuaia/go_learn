package linkedlist

import (
	"go_learn/leetcode/util"
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	cases := []struct {
		name           string
		input1, input2 []int
		expected       []int
	}{
		{"Normal cases without carry one", []int{1, 2, 3}, []int{4, 5, 6}, []int{5, 7, 9}},
		{"Carry one bit", []int{1, 2, 3}, []int{9, 5, 6}, []int{0, 8, 9}},
		{"Carry more than one bit", []int{1, 4, 3}, []int{9, 5, 6}, []int{0, 0, 0, 1}},
		{"Carry more than one bit", []int{9, 9, 9}, []int{9, 9, 9}, []int{8, 9, 9, 1}},
		{"Zero as input", []int{0}, []int{1, 2, 3}, []int{1, 2, 3}},
		{"Zero as input", []int{0}, []int{0}, []int{0}},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			l1 := util.CreateLinkedList(tt.input1)
			l2 := util.CreateLinkedList(tt.input2)
			got := addTwoNumbers(l1, l2)
			gotArray := util.TransformLinkedListToArray(got)
			if !reflect.DeepEqual(gotArray, tt.expected) {
				t.Errorf("addTwoNumbers() = %v, want %v", gotArray, tt.expected)
			}
		})
	}
}
