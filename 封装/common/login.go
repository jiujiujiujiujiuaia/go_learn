package common

import "fmt"

//对外供调用的函数
//返回一个可以使用接口的调用者
//实现解耦 调用方和实现方没有强耦合
func NewLogin() Loginer{
	return defaultLogin(0)
}


type Loginer interface {
	Login()
}

type defaultLogin int

func (d defaultLogin) Login()  {
	fmt.Println("login in ...")
}