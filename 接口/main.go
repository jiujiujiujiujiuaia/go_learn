package main

import (
	"fmt"
)

// go接口要点
// 1.接口和实现接口者必须在同一个packge下
// 2.必须实现接口下所有方法
// 3 实现接口时定义的接受者不同,就会引发方法集的概念,接收者是指针时,只能用指针对象调用,如果接收者是值,那对象两者都可以
func main() {
	var animal Animal
	var cat Cat
	var dog Dog
	var bird Bird
	var OtherFile Chicken


	//多态
	animal = cat
	animal.print()

	animal = dog
	animal.print()

	animal = bird
	animal.print()

	animal = OtherFile
	animal.print()

	Test()
}
type Animal interface {
	print()

}

//go中可以实现接口的除了结构体,还有任意的type打头用户自定义的类型
type Bird int

type Cat struct {

}

type Dog struct {

}

func (cat Cat) print(){
	fmt.Println("i am a cat")
}

func (dog Dog) print(){
	fmt.Println("i am a dog")
}

func (bird Bird ) print(){
	fmt.Println("i am a bird")
}

// 所有的type 类型都实现了empty interface 也就是所有类型都是interface
func Test(a ...interface{}){
	fmt.Println("a","b","c")
}

