package main

import (
	"fmt"
	"path/filepath"
)

type GetAllCourseReq struct {
	CourseId    int64
	CourseName  string
	TeacherName string
}

func mains() {
	fmt.Println("# Start new stream: ip:129.204.8.88:44454,125.220.29.210:55689")
	fmt.Println("2019-09-21 16:37:26| cli(129.204.8.88:44454)   -> ser(125.220.29.210:55689) | 应用层协议内容")
	fmt.Println("2019-09-21 16:37:26| ser(125.220.29.210:55689) -> cli(129.204.8.88:44454)   | 应用层协议内容")
	fmt.Println("2019-09-21 16:37:26| cli(129.204.8.88:44454)   -> ser(125.220.29.210:55689) | 应用层协议内容")
	fmt.Println("2019-09-21 16:37:26| cli(129.204.8.88:44454)   -> ser(125.220.29.210:55689) | 应用层协议内容")
	fmt.Println()
	fmt.Println("# Start new stream: ip:129.204.8.88:44455,125.220.29.210:55689")
	fmt.Println("2019-09-21 16:37:26| cli(129.204.8.88:44454)   -> ser(125.220.29.210:55689) | 应用层协议内容")
	fmt.Println("2019-09-21 16:37:26| ser(125.220.29.210:55689) -> cli(129.204.8.88:44454)   | 应用层协议内容")
	fmt.Println("2019-09-21 16:37:26| cli(129.204.8.88:44454)   -> ser(125.220.29.210:55689) | 应用层协议内容")
	fmt.Println("2019-09-21 16:37:26| cli(129.204.8.88:44454)   -> ser(125.220.29.210:55689) | 应用层协议内容")
	fmt.Println()
	fmt.Println("# Start new stream: ip:129.204.8.88:44455,125.220.29.210:55689")
	fmt.Println("2019-09-21 16:37:26| cli(129.204.8.88:44454)->ser(125.220.29.210:55689) |\nApplication Layer (200 bytes) = { \nhead = { AppName=go_demo Cmd=GetAllCourse TraceId=1111 Uid=870572451 Seq=111111 AuthKey=111 AuthType=111 AuthIp=127.0.0.1 AuthAppid=11}\nbody = { message=GetAllCourseReq CourseId=121778 CourseName=秋季班课程 TeacherName=超人气老师}\n}")

	dir, _ := filepath.Abs("./plug/")
	fmt.Println(dir)

	a := []int{1, 2, 3, 4, 5, 6}
	temp1 := a[:1]
	fmt.Println(cap(temp1), len(temp1), temp1)
	b := append(a[:1], 3, 4, 5, 7)
	fmt.Println(a)
	fmt.Println(b)

	slice := []int{10, 20, 30, 40}
	newSlice := append(slice, 50)
	fmt.Printf("Before slice = %v, Pointer = %p, len = %d, cap = %d\n", slice, &slice, len(slice), cap(slice))
	fmt.Printf("Before newSlice = %v, Pointer = %p, len = %d, cap = %d\n", newSlice, &newSlice, len(newSlice), cap(newSlice))
	newSlice[1] += 10
	fmt.Printf("After slice = %v, Pointer = %p, len = %d, cap = %d\n", slice, &slice, len(slice), cap(slice))
	fmt.Printf("After newSlice = %v, Pointer = %p, len = %d, cap = %d\n", newSlice, &newSlice, len(newSlice), cap(newSlice))

	tempSolution := make([][]int, 2)

	adsd := [][]int{{1, 2}, {3, 4}}
	copy(tempSolution, adsd)
	fmt.Println(tempSolution)
}
