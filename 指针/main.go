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
	fmt.Println("p 自身内存地址是 ",&p)
	fmt.Println("p 存放的值是 ",p)
	fmt.Println("i 本身的地址是 ",&i)
	point_pass(p)
	fmt.Println("i is ",i)
	value_pass(i)

	fmt.Println(prex)
	person := make(map[string]int)
	person["liam"] = 22
	fmt.Printf("map 地址是%p\n",&person)
	makeUsage(person)
	fmt.Println("liam 年龄被修改为了 ", person["liam"])
	fmt.Println(prex)
	//下面代码会报错,因为没有给地址初始值,new和make就不会有这样的问题
	var a *int
	a = new(int)  //不去掉注释会报错
	*a = 5
	fmt.Println(a)
}

//注意,如果函数间传指针,会修改指针内存地址,但是其值(指针存放别人的地址不会改变)
//因此go中传递无论是指针还是基本类型,变量存放的值都不变,这也是go值传递的核心(拷贝值的副本)
func point_pass(p *int){
	fmt.Println(prex)
	fmt.Println("传进point_pass函数的p本身地址",&p)
	*p = 100
	//没有变化
	fmt.Println("p存的值是",&(*p))
}

func value_pass(i int){
	fmt.Println(prex)
	fmt.Println("传进value_pass2函数中i的地址 ",&i)
}

//go语言中只有指针类型,引用类型,基本类型,值传递,没有引用传递这个概念
//也就是我想修改离开原本的作用域修改一个变量的值,得传递指针类型


//可能会奇怪没有传指针却可以修改map,那是因为使用make相当返回的是map的一个指针类型
func makeUsage(p map[string]int){
	fmt.Printf("makeUsage函数内地址是%p\n",&p)
	p["liam"] = 18
}