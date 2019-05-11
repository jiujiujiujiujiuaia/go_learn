package main

import "fmt"

const gap  = "------------------------"

//1.建议多使用切片,仅仅复制三个字段(指向底层数组的指针,长度,容量),就可以实现函数间的传递,并且可以修改
func main() {
	s1 := []int{1,2,3,4,5}
	slice()
	appends()
	iterable(s1)
	fmt.Println(s1)
	arrays()
}
//切片初始化的方式
func initUsage(){
	//s1 := make(int[],5)
	//s1 := []int{1,2,3,4,5}
	//s1 := s2[2:]
	//var s1 int[]
}


func printInfo(s []int,name string){
	fmt.Printf("name := %s,len := %d,cap :=%d \n",name,len(s),cap(s))
}
func slice(){
	fmt.Println("slice()")
	s1 := []int{1,2,3}
	printInfo(s1,"s1")
	s2 := s1[:2]
	printInfo(s2,"s2")
	//在不超出s2 cap情况下,修改和添加s2元素,也会反应到s1上,也就是s1和s2指向同一个底层数组
	s2 = append(s2,10)
	fmt.Println(s1)
	fmt.Println(s2)
	//如果超过了s2 cap,那么底层数据就没有了足够的容量,就会新建一个底层数组,把之前的复制过来
	//那么s1和s2就没有指向同一个底层数组了
	s2 = append(s2,100)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(gap)
}

func appends(){
	//append追加一个切片到另一个切片,这两指向的不是一个底层数组
	fmt.Println("appends()")
	s1 := []int{1,2,3}
	var s2 []int
	s2 = append(s2,s1...)
	printInfo(s1,"s1")
	printInfo(s2,"s2")
	s2[0] = 100
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(gap)
}

func iterable(s1 []int){
	fmt.Println("iterable()")

	//迭代切片地址都是一样,说明是切片元素的复制,而不是引用
	for _,v := range s1{
		v = 100
		fmt.Printf("value is %v,address is %p\n",v,&v)
	}
	fmt.Println(s1)
	//如果想修改的话
	for i:=0;i<len(s1);i++{
		s1[i] = 100
	}
	fmt.Println(s1)
}

//

type student struct {
	name string
	age string
}

func iterables(studengs []student){
	for i:=0;i<len(studengs);i++{
		studengs[i].age = "100"
	}
}


func arrays(){
	fmt.Println(gap)
	studengs := []student{
		{"kobe",""},
		{"ycw",""},
	}
	iterables(studengs)
	fmt.Println(studengs)

}
