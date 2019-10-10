package main

import (
	"fmt"
	echo_service "go_learn/bingo的一些测试/demo_pb/pb"
	"reflect"
	"strings"

	"github.com/lexkong/log"

	"github.com/golang/protobuf/proto"
)

func main() {
	proto.RegisterType((*echo_service.EchoBackReq)(nil), "s")
	reqBody, ok := reflect.New(proto.MessageType("s").Elem()).Interface().(proto.Message)
	a := echo_service.EchoBackReq{proto.String("hello World"), nil}
	bytes, _ := proto.Marshal(&a)
	fmt.Println(Tool_DecimalByteSlice2HexString(bytes))
	if !ok {
		log.Infof("handler not support type")
	}
	fmt.Println(reqBody)
	printPb()
}
func Tool_DecimalByteSlice2HexString(DecimalSlice []byte) string {
	var sa = make([]string, 0)
	for _, v := range DecimalSlice {
		sa = append(sa, fmt.Sprintf("%02X", v))
	}
	ss := strings.Join(sa, "")
	return ss
}

func printPb() {
	fmt.Println("proto文件：")
	fmt.Println("message ApplyCourseInfo{")
	fmt.Println("optional int64 uid_type = 1;")
	fmt.Println("optional int64 uid = 2;")
	fmt.Println("optional int64 cid = 3;")
	fmt.Println("optional string operator = 4;")
	fmt.Println("optional string reason = 5;")
	fmt.Println("}")
	a := echo_service.ApplyCourseInfo{proto.Int64(1), proto.Int64(870572451), proto.Int64(122345), proto.String("liam"), proto.String("111"), nil}
	bytes, _ := proto.Marshal(&a)
	fmt.Println()
	fmt.Println("二进制流：", bytes)
	fmt.Println()
	fmt.Println("pb标识号和值对应关系" + "1:1 2:870572451 3:122345 4:'liam' 5:'111'")
}
