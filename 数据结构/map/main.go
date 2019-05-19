package main

import (
	"fmt"
	"go_learn/util"
	"os"
	"sort"
	"strconv"
)

var slices []util.Student
var maps map[string]util.Student
var mapsInt map[int]string

func init() {
	slices = []util.Student{
		{"ycw", "100"}}
	maps = map[string]util.Student{
		"ycw": {
			"ycw", "100"},
		"kobe": {
			"kobe", "101"}}
	mapsInt = map[int]string{
		12: "b", 321: "a", 54: "c"}
}

func main() {
	for i := 1; i < len(os.Args); i++ {
		p, _ := strconv.Atoi(os.Args[i])
		switch p {
		case 1:
			passMapStruct(maps)
			passSliceStruct(slices)
		case 2:
			isExist()
		case 3:
			noSortMap(mapsInt)
		}
	}
}

func initUsage() {
	//下面代码会报错,下面代码是nil,未初始化
	//var a map[int]string
	//a[1] = "aaa"
	//:= 保证了声明变量加初始化
}

func noSortMap(mapsInt map[int]string) {
	util.PrintGap("golang中的map只有无序的，如果想有序的话，只有利用slice了")
	for key, value := range mapsInt {
		fmt.Printf("%v:%v\n", key, value)
	}
	util.PrintGap("使用slice排序key或者valu使得map有序输出或者使用结构体切片用sort包下的方法")
	sliceInt := make([]int, 0)
	for key, _ := range mapsInt {
		sliceInt = append(sliceInt, key)
	}
	sort.Ints(sliceInt)
	for _, key := range sliceInt {
		fmt.Printf("%v:%v\n", key, mapsInt[key])
	}
}

func isExist() {
	util.PrintGap("判断map某个值是否存在方法")
	hashMap := map[int]string{}
	hashMap[1] = "aa"
	if value, ok := hashMap[1]; ok {
		fmt.Println("存在:" + value)
	}
}

func passMapStruct(maps map[string]util.Student) {
	util.PrintGap("函数间传递map结构体")
	fmt.Println(maps)
	a := maps["ycw"]
	a.Age = "99"
	fmt.Println(maps)
}
func passSliceStruct(slices []util.Student) {
	util.PrintGap("函数间传递切片结构体")
	fmt.Println(slices)
	slices[0].Age = "99"
	fmt.Println(slices)
	util.PrintGap("总结：使用map结构体最好使用map*struct，方便函数间传递，还可以直接访问成员变量")
}
