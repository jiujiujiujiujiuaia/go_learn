package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"go_learn/bingo的一些测试"

	"github.com/golang/protobuf/proto"
)

type ProtoErrorCode int

const (
	EC_PROTO_OK                   ProtoErrorCode = 0
	EC_PROTO_BAD_PACKAGE          ProtoErrorCode = 1
	EC_PROTO_ERR_STX              ProtoErrorCode = 2
	EC_PROTO_ERR_ETX              ProtoErrorCode = 3
	EC_PROTO_ERR_PKG_LEN          ProtoErrorCode = 4
	EC_PROTO_ERR_BAD_PACKAGE_TYPE ProtoErrorCode = 5
)

const (
	//表示QAPP协议的开始符号
	QAPP_STX_C = 0x2
	//表示QAPP协议的终止符号
	QAPP_ETX_C = 0x3

	QUN_STX_C = 0x2
	QUN_ETX_C = 0x3
)

//表示qapp协议的包,qapp协议的核心,是一个包头,加package实现
type QappPkg struct {
	qapp_pkg proto.Message
}

func main() {
	var pkg QAppReqPkg
	a := "hello world"
	pkg.qapp_pkg = &bingo的一些测试.QAppResponse{proto.Uint64(100), proto.Uint32(256), proto.String("error"), []byte(a), &bingo的一些测试.ExtInfo{}, []byte{}}
	bytes, _ := pkg.Encode()
	pkg.Decode(bytes)

}

//实现了Encode操作,所有协议相关的逻辑,只需要实现Encode操作即可。
func (pkg *QappPkg) Encode() ([]byte, error) {

	buff, err := proto.Marshal(pkg.qapp_pkg)
	fmt.Println(string(buff))
	if err != nil {
		return nil, err
	}

	var pkg_len uint32 = 2 + 4 + uint32(len(buff))

	buffer := bytes.NewBuffer(nil)
	buffer.WriteByte(QAPP_STX_C)
	binary.Write(buffer, binary.BigEndian, &pkg_len)
	buffer.Write(buff)
	buffer.WriteByte(QAPP_ETX_C)

	return buffer.Bytes(), nil
}

//表示协议的解码逻辑
func (pkg *QappPkg) Decode(buff []byte) error {

	if len(buff) < 2+4 {
		return EC_PROTO_ERR_PKG_LEN
	}

	buffer := bytes.NewBuffer(buff)
	if qapp_stx, err := buffer.ReadByte(); err != nil || qapp_stx != QAPP_STX_C {
		return EC_PROTO_ERR_STX
	}

	var pkg_len uint32
	if err := binary.Read(buffer, binary.BigEndian, &pkg_len); err != nil {

		return err
	}

	if uint32(len(buff)) < pkg_len {
		return EC_PROTO_ERR_PKG_LEN
	}

	data_buffer := buffer.Next(int(pkg_len) - 1 - 4 - 1)

	if qapp_etx, err := buffer.ReadByte(); err != nil || qapp_etx != QAPP_ETX_C {
		return EC_PROTO_ERR_ETX
	}

	if err := proto.Unmarshal(data_buffer, pkg.qapp_pkg); err != nil {
		return err
	}

	//表示没有错误
	return nil
}

//表示Qapp协议的请求包
type QAppReqPkg struct {
	QappPkg
}

//表示Qapp协议的响应包
type QAppRspPkg struct {
	QappPkg
}

func (code ProtoErrorCode) Error() string {

	var str string
	switch code {
	case EC_PROTO_OK:
		str = ""
	case EC_PROTO_BAD_PACKAGE:
		str = "错误数据包"
	case EC_PROTO_ERR_STX:
		str = "错误的起始符"
	case EC_PROTO_ERR_ETX:
		str = "错误的终止符"
	case EC_PROTO_ERR_PKG_LEN:
		str = "错误的数据包长度"
	case EC_PROTO_ERR_BAD_PACKAGE_TYPE:
		str = "异常的数据包类型"
	default:
		str = "未知错误"
	}

	return str
}
