package main

import (
	"strings"
)

type Stack struct {
	arr  []string
	tail int
}

func (stack Stack) push(res string) {
	stack.tail++
	stack.arr = append(stack.arr, res)
}

func (stack Stack) pop() string {
	res := stack.arr[stack.tail]
	stack.arr = stack.arr[:stack.tail]
	stack.tail--
	return res
}

func (stack Stack) isEmpty() bool {
	if len(stack.arr) == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	stack := Stack{[]string{}, 0}
	a := "/home//foo/.."
	str := strings.Split(a, "/")
	for _, v := range str {
		if v != " " {

		}
	}
}
