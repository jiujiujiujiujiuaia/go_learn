package main

import "fmt"

//go中没有继承的概念,相应的代码复用这一块,只支持组合,也就是说,只支持嵌入类型
type User struct {
	Name string
	Age int
}
type admin struct {
	User
	Level string
}

func(u User) SayHello(a string){
	fmt.Println("hello l am a user")
}

func(a admin) SayHello(b string){
	fmt.Println("hello l am a admin")
}


//接口也可以实现嵌入类型

type Getter interface {
	Get()
}

type Setter interface {
	Set(a int)
}

type Common interface {
	Getter
	Setter
}

//
func (u *User) Set(a int){
	u.Age = a
}

func main() {
	ad := admin{User{
			"liam",
			22,
		}, "超级管理员",
	}
	//无论是使用内部类型访问变量还是外部类型访问变量 都是可以的
	//虽然嵌入类型是组合的方式,但是很像是继承,拥有了这个变量
	fmt.Println(ad.Name)
	fmt.Println(ad.User.Name)
	
	//重名的方法会发生重写函数的情况
	ad.SayHello("")
	ad.User.SayHello("")

	//外部类型直接拥有内部类型实现的方法
	ad.Set(100)
	fmt.Println(ad.Age)
}
