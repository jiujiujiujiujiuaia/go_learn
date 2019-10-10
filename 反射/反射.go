package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

type User struct {
	Name string
	Age  int
}

func (u User) UserFucntion(args string) {
	fmt.Println("this is UserFucntion execute args=" + args)
}

func (u User) userFunction() {
	fmt.Println("userFunction")
}

func (u User) String() string {
	return ""
}

func main() {
	u := User{"zx", 18}
	t := reflect.TypeOf(u)
	v := reflect.ValueOf(u)

	i := 1
	var w io.Writer = os.Stdout

	fmt.Println()
	printTypeAndValueKind(t, v)
	//fmt.Println(v.Elem())
	printTypeAndValue(w)
	printTypeAndValue(i)
	printTypeAndValue(u)

	u1 := v.Interface().(User)
	fmt.Printf("v.Interface().(User) - > user=[%+v]\n", u1)

	t1 := v.Type()
	fmt.Printf("v.Type()=[%+v]\n", t1)

	//这里只能找到可以导出的，也就是大写的filed或者fuction
	for i := 0; i < t.NumField(); i++ {
		fmt.Println(t.Field(i).Name)
	}

	for i := 0; i < t.NumMethod(); i++ {
		fmt.Printf("method name=%s type=%s,pkgPath=%s\n", t.Method(i).Name,
			t.Method(i).Type, t.Method(i).PkgPath)
	}

	//反射调用结构体的方法
	method := v.MethodByName("UserFucntion")
	args := []reflect.Value{reflect.ValueOf("参数")}
	method.Call(args)
}

func printTypeAndValue(ins interface{}) {
	fmt.Printf("TypeOf=[%+v] ValueOf=[%+v] TypeName=[%+v] \n", reflect.TypeOf(ins), reflect.ValueOf(ins), reflect.TypeOf(ins).Name())
}

func printTypeAndValueKind(t reflect.Type, v reflect.Value) {
	fmt.Printf("type.kind=[%+v] value.kind=[%+v]\n", t.Kind(), v.Kind())
}
