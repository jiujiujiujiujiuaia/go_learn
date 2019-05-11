package main

import "fmt"

func main() {
	//下面代码会报错,下面代码是nil,未初始化
	//var a map[int]string
	//a[1] = "aaa"

	//


	//:= 保证了声明变量加初始化
	a := map[int]string{}
	a[1] = "aa"

	//判断map某个值是否存在
	isExist(a)

}

func isExist(hashMap map[int]string){
	if value,ok := hashMap[1];ok{
		fmt.Println("存在:"+value)
	}
}


