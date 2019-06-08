package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func main() {
	stu1 := Student{
		"?", 11}
	stu2 := Student{
		"?", 11}

	passStructPoint(&stu1)
	fmt.Println(stu1)
	passStructValue(stu2)
	fmt.Println(stu2)
}

func passStructPoint(stu *Student) {
	stu.Name = "Point"
}

func passStructValue(stu Student) {
	stu.Name = "Value"
}
