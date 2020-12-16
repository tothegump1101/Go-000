// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package user_api_v1

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

type GetUserArgs struct {
	UserId               uint64   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserArgs) Reset()         { *m = GetUserArgs{} }
func (m *GetUserArgs) String() string { return proto.CompactTextString(m) }
func (*GetUserArgs) ProtoMessage()    {}
func (*GetUserArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *GetUserArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserArgs.Unmarshal(m, b)
}
func (m *GetUserArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserArgs.Marshal(b, m, deterministic)
}
func (m *GetUserArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserArgs.Merge(m, src)
}
func (m *GetUserArgs) XXX_Size() int {
	return xxx_messageInfo_GetUserArgs.Size(m)
}
func (m *GetUserArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserArgs.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserArgs proto.InternalMessageInfo

func (m *GetUserArgs) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type GetUserResp struct {
	UserId               uint64   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserName             string   `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserResp) Reset()         { *m = GetUserResp{} }
func (m *GetUserResp) String() string { return proto.CompactTextString(m) }
func (*GetUserResp) ProtoMessage()    {}
func (*GetUserResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *GetUserResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserResp.Unmarshal(m, b)
}
func (m *GetUserResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserResp.Marshal(b, m, deterministic)
}
func (m *GetUserResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserResp.Merge(m, src)
}
func (m *GetUserResp) XXX_Size() int {
	return xxx_messageInfo_GetUserResp.Size(m)
}
func (m *GetUserResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserResp proto.InternalMessageInfo

func (m *GetUserResp) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *GetUserResp) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetUserArgs)(nil), "user.api.v1.GetUserArgs")
	proto.RegisterType((*GetUserResp)(nil), "user.api.v1.GetUserResp")
}

func init() {
	proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf)
}

var fileDescriptor_116e343673f7ffaf = []byte{
	// 174 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x2d, 0x4e, 0x2d,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x06, 0xb3, 0x13, 0x0b, 0x32, 0xf5, 0xca, 0x0c,
	0x95, 0xd4, 0xb8, 0xb8, 0xdd, 0x53, 0x4b, 0x42, 0x8b, 0x53, 0x8b, 0x1c, 0x8b, 0xd2, 0x8b, 0x85,
	0xc4, 0xb9, 0xd8, 0x41, 0xb2, 0xf1, 0x99, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x2c, 0x41, 0x6c,
	0x20, 0xae, 0x67, 0x8a, 0x92, 0x33, 0x5c, 0x5d, 0x50, 0x6a, 0x71, 0x01, 0x4e, 0x75, 0x42, 0xd2,
	0x5c, 0x9c, 0x60, 0x89, 0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x0e,
	0x90, 0x80, 0x5f, 0x62, 0x6e, 0xaa, 0x91, 0x2b, 0x17, 0x0b, 0xc8, 0x04, 0x21, 0x5b, 0x2e, 0x76,
	0xa8, 0x61, 0x42, 0x12, 0x7a, 0x48, 0xae, 0xd1, 0x43, 0x72, 0x8a, 0x14, 0x56, 0x19, 0x90, 0xe5,
	0x4e, 0x9a, 0x5c, 0x22, 0xc9, 0xf9, 0xb9, 0x20, 0xdf, 0x64, 0xa5, 0x26, 0x97, 0xc0, 0x95, 0x05,
	0x30, 0x2e, 0x62, 0xe2, 0x0b, 0x80, 0x08, 0x82, 0x4d, 0x29, 0xc8, 0x4c, 0x62, 0x03, 0x7b, 0xd9,
	0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x17, 0x8b, 0x57, 0x55, 0x00, 0x01, 0x00, 0x00,
}