// Code generated by protoc-gen-go. DO NOT EDIT.
// source: echo_service.proto

/*
Package echo_service is a generated protocol buffer package.

It is generated from these files:
	echo_service.proto

It has these top-level messages:
	EchoBackReq
	EchoBackRsp
	EchoNoneReq
	EchoNoneRsp
*/
package echo_service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type EchoBackReq struct {
	Text             *string `protobuf:"bytes,1,req,name=text" json:"text,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *EchoBackReq) Reset()                    { *m = EchoBackReq{} }
func (m *EchoBackReq) String() string            { return proto.CompactTextString(m) }
func (*EchoBackReq) ProtoMessage()               {}
func (*EchoBackReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *EchoBackReq) GetText() string {
	if m != nil && m.Text != nil {
		return *m.Text
	}
	return ""
}

type EchoBackRsp struct {
	Result           *int32  `protobuf:"varint,1,req,name=result" json:"result,omitempty"`
	Text             *string `protobuf:"bytes,2,req,name=text" json:"text,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *EchoBackRsp) Reset()                    { *m = EchoBackRsp{} }
func (m *EchoBackRsp) String() string            { return proto.CompactTextString(m) }
func (*EchoBackRsp) ProtoMessage()               {}
func (*EchoBackRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *EchoBackRsp) GetResult() int32 {
	if m != nil && m.Result != nil {
		return *m.Result
	}
	return 0
}

func (m *EchoBackRsp) GetText() string {
	if m != nil && m.Text != nil {
		return *m.Text
	}
	return ""
}

type EchoNoneReq struct {
	Text             *string `protobuf:"bytes,1,req,name=text" json:"text,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *EchoNoneReq) Reset()                    { *m = EchoNoneReq{} }
func (m *EchoNoneReq) String() string            { return proto.CompactTextString(m) }
func (*EchoNoneReq) ProtoMessage()               {}
func (*EchoNoneReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *EchoNoneReq) GetText() string {
	if m != nil && m.Text != nil {
		return *m.Text
	}
	return ""
}

type EchoNoneRsp struct {
	Result           *int32 `protobuf:"varint,1,req,name=result" json:"result,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *EchoNoneRsp) Reset()                    { *m = EchoNoneRsp{} }
func (m *EchoNoneRsp) String() string            { return proto.CompactTextString(m) }
func (*EchoNoneRsp) ProtoMessage()               {}
func (*EchoNoneRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *EchoNoneRsp) GetResult() int32 {
	if m != nil && m.Result != nil {
		return *m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*EchoBackReq)(nil), "echo_service.EchoBackReq")
	proto.RegisterType((*EchoBackRsp)(nil), "echo_service.EchoBackRsp")
	proto.RegisterType((*EchoNoneReq)(nil), "echo_service.EchoNoneReq")
	proto.RegisterType((*EchoNoneRsp)(nil), "echo_service.EchoNoneRsp")
}

func init() { proto.RegisterFile("echo_service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 155 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4a, 0x4d, 0xce, 0xc8,
	0x8f, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2,
	0x41, 0x16, 0x53, 0x92, 0xe6, 0xe2, 0x76, 0x4d, 0xce, 0xc8, 0x77, 0x4a, 0x4c, 0xce, 0x0e, 0x4a,
	0x2d, 0x14, 0xe2, 0xe1, 0x62, 0x29, 0x49, 0xad, 0x28, 0x91, 0x60, 0x54, 0x60, 0xd2, 0xe0, 0x54,
	0xd2, 0x46, 0x92, 0x2c, 0x2e, 0x10, 0xe2, 0xe3, 0x62, 0x2b, 0x4a, 0x2d, 0x2e, 0xcd, 0x81, 0x48,
	0xb3, 0xc2, 0x15, 0x33, 0x81, 0x15, 0x43, 0x4d, 0xf2, 0xcb, 0xcf, 0x4b, 0xc5, 0x34, 0x49, 0x16,
	0x49, 0x12, 0xd3, 0x24, 0xa3, 0x89, 0x8c, 0x10, 0xf9, 0x60, 0x88, 0xab, 0x84, 0x1c, 0xb8, 0x38,
	0x60, 0x16, 0x0b, 0x49, 0xea, 0xa1, 0x78, 0x02, 0xc9, 0xb5, 0x52, 0xb8, 0xa4, 0x8a, 0x0b, 0x60,
	0x26, 0x80, 0x2c, 0xc4, 0x66, 0x02, 0xd4, 0x95, 0x52, 0xb8, 0xa4, 0x8a, 0x0b, 0x00, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x11, 0xc9, 0xf4, 0xcb, 0x3c, 0x01, 0x00, 0x00,
}
