package main

import (
	"fmt"
)

var prex string = "------------------"

//指针要点
//1.函数间传递值想要修改,只能通过指针
//2.指针类型如果声明不给初始内存地址会报错
//3.对于切片make创造出来的map,slice数据类型,可以直接在函数间传递(因为是一种封装的数据结构,封装好的结构体,内部含有指针)
func main() {
	i := 10
	p := &i
	fmt.Println("1.任意一个变量，都有两个属性，一个是自己的内存地址，一个是自己的值")
	fmt.Println("i=",i)
	fmt.Println("&i=",&i)
	fmt.Println(prex)
	fmt.Println("2.而对于指针而言，指针的值是别人的地址，指针的地址是自己本身的内存地址")
	fmt.Println("&p=",&p)
	fmt.Printf("p=%p == &i=%p\n",p,&i)
	fmt.Println("*p=",*p)
	fmt.Println(prex)
	fmt.Println("3.[指针的运算符] &是取地址符号，任何一个变量都有自己的地址。*是指针变量的运算符，代表找到地址指向的那个值")
	
	fmt.Println(prex)
	point_pass(p,&i)
	fmt.Println("退出函数作用域后，发现i的值被修改了")
	fmt.Println("i=",i)

	fmt.Println(prex)
	value_pass(i,&i)
	fmt.Println("退出函数作用域后，发现i的值并没有被修改")
	fmt.Println("i=",i)

	fmt.Println(prex)
	fmt.Println("6.[golang数据结构] go语言中只有指针类型,引用类型,基本类型。而使用map array这一类数据结构，修改其property是不需要传递指针的")
	person := make(map[string]int)
	person["liam"] = 22
	fmt.Printf("&map=%p\n",&person)
	fmt.Println("liam.age=",person["liam"])
	makeUsage(person)
	fmt.Println("liam.age= ", person["liam"])

	fmt.Println(prex)
	fmt.Println("7.[初始化指针] 初始化指针一定要赋值，无论是使用new,make还是直接=赋值")
	var a *int
	fmt.Println("如果去掉a = new(int)进行初始化，就会抛panic")
	a = new(int)
	*a = 5
	fmt.Println(a)
}

//注意,如果函数间传指针,会修改指针内存地址,但是其值(指针存放别人的地址不会改变)
//因此go中传递无论是指针还是基本类型,变量存放的值都不变,这也是go值传递的核心(拷贝值的副本)
func point_pass(p *int, i *int){
	fmt.Println("4.[函数间传递指针] go的函数间参数传递是浅拷贝，也就是只是拷贝副本 -> 变量的地址会变，但是值不会变 -> 拷贝变量的值，然后在内存中创造另外一个变量，赋予给其")
	fmt.Println("&p=",&p)
	fmt.Printf("p=%p == &i=%p\n",p,i)
	fmt.Println("在函数中修改*p=100")
	*p = 100
}

func value_pass(i int, pointI *int){
	fmt.Println("5.[函数间传递值] 同理，由于是值拷贝（浅拷贝），因此i=100的值会被复制给另外一个新的变量，所以i的地址发生了变化")
	fmt.Printf("&i=%p != pointI=%p\n",&i,pointI)
	i=1000
	fmt.Println("i=",1000)
}

//go语言中只有指针类型,引用类型,基本类型,值传递,没有引用传递这个概念
//也就是我想修改离开原本的作用域修改一个变量的值,得传递指针类型
//可能会奇怪没有传指针却可以修改map,那是因为使用make相当返回的是map的一个指针类型
func makeUsage(p map[string]int){
	fmt.Printf("函数中, &map=%p\n",&p)
	p["liam"] = 18
}