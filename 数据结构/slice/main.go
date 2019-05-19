package main

import (
	"fmt"
	"go_learn/util"
	"os"
	"strconv"
)

//1.建议多使用切片,仅仅复制三个字段(指向底层数组的指针,长度,容量),就可以实现函数间的传递,并且可以修改
//2.切片和数组应该都是引用类型，也就是在函数间传递过程中修改了值，这种修改是永久的（也就是使用指针修改的一样）
func main() {
	s1 := []int{1, 2, 3, 4, 5}
	for i := 1; i < len(os.Args); i++ {
		p, _ := strconv.Atoi(os.Args[i])
		switch p {
		case 1:
			slice()
		case 2:
			appends()
		case 3:
			iterable(s1)
		case 4:
			sliceAndStruct()
		case 5:
			initUsage()
		case 6:
			newAndMake()
		}
	}

}

func initUsage() {
	fmt.Printf(util.Gap, "切片初始化的方式")
	s2 := make([]int, 5)
	s1 := make([]int, 5)
	s1 = []int{1, 2, 3, 4, 5}
	s1 = s2[2:]
	util.PrintSliceMember(s1, "s1")
	//下面只执行了声明变量字面量，没有初始化，使用时会报错
	//var s3 []int
	//s3 = s1

}

//超出容量之后两个切片指向的就不是一个数组了
func slice() {
	util.PrintGap("超出容量之后两个切片指向的就不是一个数组了")
	s1 := []int{1, 2, 3}
	util.PrintInfo(s1, "s1")
	s2 := s1[:2]
	util.PrintInfo(s2, "s2")
	util.PrintGap("在不超出s2cap情况下,修改和添加s2元素,也会反应到s1上,也就是s1和s2指向同一个底层数组")
	s2 = append(s2, 10)
	util.PrintSliceMember(s1, "s1")
	util.PrintSliceMember(s2, "s2")
	util.PrintGap("如果超过了s2 cap,就会新建一个底层数组,把之前的复制过来,s1和s2没有指向同一个底层数组")
	s2 = append(s2, 100)
	util.PrintSliceMember(s1, "s1")
	util.PrintSliceMember(s2, "s2")
}

func appends() {
	util.PrintGap("append追加一个切片到另一个切片,这两指向的不是一个底层数组")
	s1 := []int{1, 2, 3}
	s2 := s1
	//s2 := make([]int, 0)
	s2 = append(s2, s1...)
	util.PrintSliceMember(s1, "s1")
	util.PrintSliceMember(s2, "s2")
	s2[0] = 100
	util.PrintSliceMember(s1, "s1")
	util.PrintSliceMember(s2, "s2")
}

func iterable(s1 []int) {
	util.PrintGap("迭代切片地址都是一样,说明是切片元素的复制,而不是引用")
	for _, v := range s1 {
		v = 100
		fmt.Printf("value is %v,address is %p\n", v, &v)
	}
	util.PrintGap("因此如果遍历过程中要修改的话，推荐索引访问")
	util.PrintSliceMember(s1, "s1")
	for i := 0; i < len(s1); i++ {
		s1[i] = 100
	}
	util.PrintSliceMember(s1, "s1")
}

func modifyStruct(studengs []util.Student) {
	for i := 0; i < len(studengs); i++ {
		studengs[i].Age = "100"
	}
}

func sliceAndStruct() {
	util.PrintGap("slice中保存结构体，函数间传递可以通过切片索引修改吗？")
	studengs := []util.Student{
		{"kobe", ""},
		{"ycw", ""},
	}
	fmt.Println(studengs)
	modifyStruct(studengs)
	fmt.Println(studengs)
}

func newAndMake() {
	util.PrintGap("使用new和make的不同")
	s1 := make([]int, 10)
	s2 := new([]int)
	util.PrintSliceMember(s1, "s1")
	util.PrintSliceMember(*s2, "s2")
	util.PrintGap("make使用于应用类型如slice,map,channel,分配内存还初始化")
	util.PrintGap("new 用于任何类型，但是只复制分配内存 -》 fmt.Println((*s2)[0]) 报错")
}
