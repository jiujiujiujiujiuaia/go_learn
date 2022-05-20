package util

import "testing"

func TestCreateMultipleLinkedList(t *testing.T) {
	cases := []struct {
		Name  string
		input []int
	}{
		{"normal case", []int{1, 2, 3}},
		{"empty input", []int{}},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			headNode := CreateLinkedList(c.input)
			for _, val := range c.input {
				if headNode == nil {
					t.Errorf("Insufficient linkedList input length=%v, the missing node value=%v", len(c.input), val)
				} else {
					if headNode.Val == val {
						headNode = headNode.Next
						continue
					} else {
						t.Errorf("node val is not equal expected value, node value=%v,expeceted value=%v", headNode.Val, val)
					}
				}
			}

			if headNode != nil {
				t.Errorf("too many nodes, the input length=%v", len(c.input))
			}
		})
	}

}

func TestCompareTwoLinkedList(t *testing.T) {
	cases := []struct {
		Name           string
		input1, input2 []int
		expected       bool
	}{
		{"Compare two same linked list", []int{1, 2, 3}, []int{1, 2, 3}, true},
		{"Empty input", []int{}, []int{}, true},
		{"Compare two linked lists of different lengths", []int{1, 2, 3}, []int{1, 2, 3, 4}, false},
		{"Compare two linked lists of different value with same length", []int{1, 2, 3}, []int{1, 2, 4}, false},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			head1 := CreateLinkedList(c.input1)
			head2 := CreateLinkedList(c.input2)
			actual := CompareTwoLinkedList(head1, head2)
			if actual != c.expected {
				t.Errorf("input1=%v,input2=%v,expected value=%v,actual value=%v", c.input1, c.input2, c.expected, actual)
			}
		})
	}

}
