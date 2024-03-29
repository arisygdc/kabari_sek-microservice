// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package pb

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

type LoginRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginResponse struct {
	UserToken            string   `protobuf:"bytes,2,opt,name=userToken,proto3" json:"userToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetUserToken() string {
	if m != nil {
		return m.UserToken
	}
	return ""
}

type RegisterRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Firstname            string   `protobuf:"bytes,3,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname             string   `protobuf:"bytes,4,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Birth                string   `protobuf:"bytes,5,opt,name=birth,proto3" json:"birth,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (m *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(m, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

func (m *RegisterRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *RegisterRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *RegisterRequest) GetFirstname() string {
	if m != nil {
		return m.Firstname
	}
	return ""
}

func (m *RegisterRequest) GetLastname() string {
	if m != nil {
		return m.Lastname
	}
	return ""
}

func (m *RegisterRequest) GetBirth() string {
	if m != nil {
		return m.Birth
	}
	return ""
}

type RegisterResponse struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterResponse) Reset()         { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()    {}
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *RegisterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResponse.Unmarshal(m, b)
}
func (m *RegisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResponse.Marshal(b, m, deterministic)
}
func (m *RegisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResponse.Merge(m, src)
}
func (m *RegisterResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterResponse.Size(m)
}
func (m *RegisterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResponse proto.InternalMessageInfo

func (m *RegisterResponse) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "pb.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "pb.LoginResponse")
	proto.RegisterType((*RegisterRequest)(nil), "pb.RegisterRequest")
	proto.RegisterType((*RegisterResponse)(nil), "pb.RegisterResponse")
}

func init() {
	proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf)
}

var fileDescriptor_116e343673f7ffaf = []byte{
	// 234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x51, 0xbb, 0x4e, 0xc4, 0x30,
	0x10, 0x54, 0x8e, 0x0b, 0xba, 0x5b, 0x81, 0x38, 0xcc, 0x15, 0x51, 0x44, 0x81, 0x5c, 0x51, 0x80,
	0x0b, 0x10, 0xbf, 0x40, 0x45, 0x15, 0xc1, 0x07, 0xc4, 0x62, 0x09, 0x56, 0xc0, 0x36, 0x5e, 0x47,
	0x7c, 0x0b, 0x7f, 0x8b, 0xfc, 0xc8, 0x03, 0x0a, 0x1a, 0xca, 0x99, 0xd9, 0xd9, 0x59, 0x8f, 0x01,
	0x06, 0x42, 0x27, 0xac, 0x33, 0xde, 0xb0, 0x95, 0x95, 0xfc, 0x1e, 0x8e, 0x1e, 0x4c, 0xa7, 0x74,
	0x83, 0x1f, 0x03, 0x92, 0x67, 0x35, 0x6c, 0xc2, 0x84, 0x6e, 0xdf, 0xb1, 0x2a, 0x2e, 0x8a, 0xcb,
	0x6d, 0x33, 0xe1, 0xa0, 0xd9, 0x96, 0xe8, 0xd3, 0xb8, 0xe7, 0x6a, 0x95, 0xb4, 0x11, 0xf3, 0x6b,
	0x38, 0xce, 0x7b, 0xc8, 0x1a, 0x4d, 0xc8, 0xce, 0x61, 0x1b, 0x8c, 0x8f, 0xa6, 0x47, 0x9d, 0xa7,
	0x67, 0x82, 0x7f, 0x15, 0x70, 0xd2, 0x60, 0xa7, 0xc8, 0xa3, 0xfb, 0x67, 0x74, 0x48, 0x7a, 0x51,
	0x8e, 0x7c, 0x34, 0x1e, 0xa4, 0xa4, 0x89, 0x08, 0xce, 0xb7, 0x36, 0x8b, 0xeb, 0xe4, 0x1c, 0x31,
	0xdb, 0x43, 0x29, 0x95, 0xf3, 0xaf, 0x55, 0x19, 0x85, 0x04, 0xb8, 0x80, 0xdd, 0x7c, 0x5a, 0x7e,
	0xcd, 0x1f, 0xb7, 0xdd, 0xf4, 0xb0, 0x7e, 0x22, 0x74, 0xec, 0x0a, 0xca, 0x58, 0x01, 0xdb, 0x09,
	0x2b, 0xc5, 0xb2, 0xd5, 0xfa, 0x74, 0xc1, 0xe4, 0x8d, 0x77, 0xb0, 0x19, 0x53, 0xd8, 0x59, 0x90,
	0x7f, 0xd5, 0x51, 0xef, 0x7f, 0x92, 0xc9, 0x26, 0x0f, 0xe3, 0xd7, 0xdd, 0x7e, 0x07, 0x00, 0x00,
	0xff, 0xff, 0x3b, 0x66, 0x00, 0x25, 0xc8, 0x01, 0x00, 0x00,
}
