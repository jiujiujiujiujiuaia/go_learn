package util

import "fmt"

const Gap = "------------%s------------\n"

func PrintInfo(s []int, name string) {
	fmt.Printf("name := %s,len := %d,cap :=%d \n", name, len(s), cap(s))
}

func PrintSliceMember(s []int, name string) {
	fmt.Printf("name=[%s],member=[%v]\n", name, s)
}

func PrintGap(str string) {
	fmt.Printf(Gap, str)
}

type Student struct {
	Name string
	Age  string
}
