package main

import (
	"fmt"
)

type stackStruct struct {
	con  []string
	tail int
}

func (s *stackStruct) push(item string) {
	s.con = append(s.con, item)
	s.tail++
}

func (s *stackStruct) pop() string {
	item := s.con[s.tail]
	s.con = s.con[:s.tail]
	s.tail--
	return item
}

func (s *stackStruct) isEmpty() bool {
	if len(s.con) == 0 {
		return true
	} else {
		return false
	}
}

//这个题目，其实是有效的括号的子括号也是有效的，因此从内部开始往外消除
//而栈的特性是先进的后出，pop出来的是栈顶的元素刚刚好是最内部的元素。
func isValid(s string) bool {
	stack := new(stackStruct)
	stack.con = make([]string, 0)
	stack.tail = -1
	maps := map[string]string{
		"{": "}", "(": ")", "[": "]"}
	for i := 0; i < len(s); i++ {
		if !stack.isEmpty() {
			item := stack.pop()
			if maps[item] != string(s[i]) {
				stack.push(item)
				stack.push(string(s[i]))
			}
		} else {
			stack.push(string(s[i]))
		}
	}
	if stack.isEmpty() {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(isValid("[{[]{}}]{}[][]"))

}
