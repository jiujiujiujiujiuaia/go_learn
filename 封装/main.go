package main

import (
	"fmt"
	"go_learn/封装/common"
	"go_learn/封装/structs"

	"reflect"
)

func main() {
	//隐藏了实现
	v := common.New(100)
	fmt.Println(reflect.TypeOf(v))

	//暴露调用者给调用方
	l := common.NewLogin()
	l.Login()

	//因为结构体user部分没有公开给
	var user structs.User
	user.Name = "100"
	fmt.Println(user.Name)

	//下面代码会报错,因为email成员变量没有使用权
	a := structs.User{"100", "100"}
	fmt.Println(a.Name)

}
