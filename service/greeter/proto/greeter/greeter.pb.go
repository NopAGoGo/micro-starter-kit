// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service/greeter/proto/greeter/greeter.proto

package greeter

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type HelloRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_50f378191cc5037f, []int{0}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type HelloResponse struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloResponse) Reset()         { *m = HelloResponse{} }
func (m *HelloResponse) String() string { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()    {}
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_50f378191cc5037f, []int{1}
}

func (m *HelloResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloResponse.Unmarshal(m, b)
}
func (m *HelloResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloResponse.Marshal(b, m, deterministic)
}
func (m *HelloResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloResponse.Merge(m, src)
}
func (m *HelloResponse) XXX_Size() int {
	return xxx_messageInfo_HelloResponse.Size(m)
}
func (m *HelloResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HelloResponse proto.InternalMessageInfo

func (m *HelloResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "mkit.service.greeter.v1.HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "mkit.service.greeter.v1.HelloResponse")
}

func init() {
	proto.RegisterFile("service/greeter/proto/greeter/greeter.proto", fileDescriptor_50f378191cc5037f)
}

var fileDescriptor_50f378191cc5037f = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2e, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x4f, 0x2f, 0x4a, 0x4d, 0x2d, 0x49, 0x2d, 0xd2, 0x2f, 0x28, 0xca, 0x2f,
	0xc9, 0x87, 0xf3, 0xa0, 0xb4, 0x1e, 0x58, 0x54, 0x48, 0x3c, 0x37, 0x3b, 0xb3, 0x44, 0x0f, 0xaa,
	0x43, 0x0f, 0x26, 0x57, 0x66, 0xa8, 0xa4, 0xc4, 0xc5, 0xe3, 0x91, 0x9a, 0x93, 0x93, 0x1f, 0x94,
	0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x22, 0x24, 0xc4, 0xc5, 0x92, 0x97, 0x98, 0x9b, 0x2a, 0xc1, 0xa8,
	0xc0, 0xa8, 0xc1, 0x19, 0x04, 0x66, 0x2b, 0x29, 0x72, 0xf1, 0x42, 0xd5, 0x14, 0x17, 0xe4, 0xe7,
	0x15, 0xa7, 0x0a, 0x09, 0x70, 0x31, 0xe7, 0x16, 0xa7, 0x43, 0xd5, 0x80, 0x98, 0x46, 0x19, 0x5c,
	0x7c, 0xee, 0x10, 0x43, 0x83, 0x21, 0x76, 0x08, 0x85, 0x71, 0xb1, 0x82, 0x35, 0x09, 0xa9, 0xea,
	0xe1, 0xb0, 0x5b, 0x0f, 0xd9, 0x62, 0x29, 0x35, 0x42, 0xca, 0x20, 0x76, 0x3b, 0xc5, 0x73, 0xe1,
	0xf2, 0x4b, 0x00, 0x63, 0x94, 0x53, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae,
	0x7e, 0x45, 0x6e, 0x4e, 0x76, 0x66, 0x5e, 0xba, 0x7e, 0x6e, 0x66, 0x72, 0x51, 0xbe, 0x6e, 0x71,
	0x49, 0x62, 0x51, 0x49, 0x6a, 0x91, 0x6e, 0x76, 0x66, 0x89, 0x3e, 0xde, 0x80, 0x4b, 0x62, 0x03,
	0x73, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x1a, 0x53, 0xd4, 0x8b, 0x60, 0x01, 0x00, 0x00,
}