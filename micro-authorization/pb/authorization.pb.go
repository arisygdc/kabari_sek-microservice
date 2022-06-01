// Code generated by protoc-gen-go. DO NOT EDIT.
// source: authorization.proto

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

type CheckPermissionRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	PermissionName       string   `protobuf:"bytes,2,opt,name=permission_name,json=permissionName,proto3" json:"permission_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckPermissionRequest) Reset()         { *m = CheckPermissionRequest{} }
func (m *CheckPermissionRequest) String() string { return proto.CompactTextString(m) }
func (*CheckPermissionRequest) ProtoMessage()    {}
func (*CheckPermissionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1dbbe58d1e51a797, []int{0}
}

func (m *CheckPermissionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckPermissionRequest.Unmarshal(m, b)
}
func (m *CheckPermissionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckPermissionRequest.Marshal(b, m, deterministic)
}
func (m *CheckPermissionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckPermissionRequest.Merge(m, src)
}
func (m *CheckPermissionRequest) XXX_Size() int {
	return xxx_messageInfo_CheckPermissionRequest.Size(m)
}
func (m *CheckPermissionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckPermissionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckPermissionRequest proto.InternalMessageInfo

func (m *CheckPermissionRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CheckPermissionRequest) GetPermissionName() string {
	if m != nil {
		return m.PermissionName
	}
	return ""
}

type CheckPermissionResponse struct {
	ResponseCode         int32    `protobuf:"varint,1,opt,name=responseCode,proto3" json:"responseCode,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckPermissionResponse) Reset()         { *m = CheckPermissionResponse{} }
func (m *CheckPermissionResponse) String() string { return proto.CompactTextString(m) }
func (*CheckPermissionResponse) ProtoMessage()    {}
func (*CheckPermissionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1dbbe58d1e51a797, []int{1}
}

func (m *CheckPermissionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckPermissionResponse.Unmarshal(m, b)
}
func (m *CheckPermissionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckPermissionResponse.Marshal(b, m, deterministic)
}
func (m *CheckPermissionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckPermissionResponse.Merge(m, src)
}
func (m *CheckPermissionResponse) XXX_Size() int {
	return xxx_messageInfo_CheckPermissionResponse.Size(m)
}
func (m *CheckPermissionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckPermissionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckPermissionResponse proto.InternalMessageInfo

func (m *CheckPermissionResponse) GetResponseCode() int32 {
	if m != nil {
		return m.ResponseCode
	}
	return 0
}

func (m *CheckPermissionResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type BannStatusRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	PermissionName       string   `protobuf:"bytes,2,opt,name=permission_name,json=permissionName,proto3" json:"permission_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BannStatusRequest) Reset()         { *m = BannStatusRequest{} }
func (m *BannStatusRequest) String() string { return proto.CompactTextString(m) }
func (*BannStatusRequest) ProtoMessage()    {}
func (*BannStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1dbbe58d1e51a797, []int{2}
}

func (m *BannStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BannStatusRequest.Unmarshal(m, b)
}
func (m *BannStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BannStatusRequest.Marshal(b, m, deterministic)
}
func (m *BannStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BannStatusRequest.Merge(m, src)
}
func (m *BannStatusRequest) XXX_Size() int {
	return xxx_messageInfo_BannStatusRequest.Size(m)
}
func (m *BannStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BannStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BannStatusRequest proto.InternalMessageInfo

func (m *BannStatusRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *BannStatusRequest) GetPermissionName() string {
	if m != nil {
		return m.PermissionName
	}
	return ""
}

type BannStatusResponse struct {
	ResponseCode         int32    `protobuf:"varint,1,opt,name=responseCode,proto3" json:"responseCode,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	PermissionName       string   `protobuf:"bytes,3,opt,name=permission_name,json=permissionName,proto3" json:"permission_name,omitempty"`
	BannAt               int64    `protobuf:"varint,4,opt,name=bannAt,proto3" json:"bannAt,omitempty"`
	BannedExp            int64    `protobuf:"varint,5,opt,name=bannedExp,proto3" json:"bannedExp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BannStatusResponse) Reset()         { *m = BannStatusResponse{} }
func (m *BannStatusResponse) String() string { return proto.CompactTextString(m) }
func (*BannStatusResponse) ProtoMessage()    {}
func (*BannStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1dbbe58d1e51a797, []int{3}
}

func (m *BannStatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BannStatusResponse.Unmarshal(m, b)
}
func (m *BannStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BannStatusResponse.Marshal(b, m, deterministic)
}
func (m *BannStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BannStatusResponse.Merge(m, src)
}
func (m *BannStatusResponse) XXX_Size() int {
	return xxx_messageInfo_BannStatusResponse.Size(m)
}
func (m *BannStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BannStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BannStatusResponse proto.InternalMessageInfo

func (m *BannStatusResponse) GetResponseCode() int32 {
	if m != nil {
		return m.ResponseCode
	}
	return 0
}

func (m *BannStatusResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *BannStatusResponse) GetPermissionName() string {
	if m != nil {
		return m.PermissionName
	}
	return ""
}

func (m *BannStatusResponse) GetBannAt() int64 {
	if m != nil {
		return m.BannAt
	}
	return 0
}

func (m *BannStatusResponse) GetBannedExp() int64 {
	if m != nil {
		return m.BannedExp
	}
	return 0
}

func init() {
	proto.RegisterType((*CheckPermissionRequest)(nil), "pb.CheckPermissionRequest")
	proto.RegisterType((*CheckPermissionResponse)(nil), "pb.CheckPermissionResponse")
	proto.RegisterType((*BannStatusRequest)(nil), "pb.BannStatusRequest")
	proto.RegisterType((*BannStatusResponse)(nil), "pb.BannStatusResponse")
}

func init() {
	proto.RegisterFile("authorization.proto", fileDescriptor_1dbbe58d1e51a797)
}

var fileDescriptor_1dbbe58d1e51a797 = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x95, 0x53, 0x5a, 0xd4, 0x13, 0xb4, 0xe2, 0x10, 0xc1, 0x0a, 0x0c, 0x55, 0x16, 0x3a, 0x65,
	0x80, 0x91, 0xa9, 0x54, 0x2c, 0x08, 0x21, 0x08, 0x03, 0x23, 0x72, 0xc8, 0x89, 0x5a, 0x28, 0xb6,
	0x89, 0x1d, 0x09, 0xf1, 0x25, 0xfc, 0x06, 0x7f, 0x88, 0x1a, 0x52, 0x52, 0x1a, 0x6f, 0xb0, 0xdd,
	0xbd, 0x77, 0x7e, 0x77, 0x7e, 0x77, 0xb0, 0x2f, 0x2a, 0xb7, 0xd0, 0xa5, 0x7c, 0x17, 0x4e, 0x6a,
	0x95, 0x98, 0x52, 0x3b, 0x8d, 0x81, 0xc9, 0xe2, 0x3b, 0x08, 0xe7, 0x0b, 0x7a, 0x7a, 0xb9, 0xa5,
	0xb2, 0x90, 0xd6, 0x4a, 0xad, 0x52, 0x7a, 0xad, 0xc8, 0x3a, 0x1c, 0x41, 0x20, 0x73, 0xce, 0x26,
	0x6c, 0x3a, 0x4c, 0x03, 0x99, 0xe3, 0x09, 0x8c, 0xcd, 0x4f, 0xd1, 0xa3, 0x12, 0x05, 0xf1, 0xa0,
	0x26, 0x47, 0x2d, 0x7c, 0x23, 0x0a, 0x8a, 0x1f, 0xe0, 0xb0, 0x23, 0x69, 0x8d, 0x56, 0x96, 0x30,
	0x86, 0x9d, 0xb2, 0x89, 0xe7, 0x3a, 0xa7, 0x5a, 0xbd, 0x9f, 0xfe, 0xc2, 0x90, 0xc3, 0x76, 0x41,
	0xd6, 0x8a, 0xe7, 0x95, 0xfe, 0x2a, 0x8d, 0xaf, 0x61, 0xef, 0x42, 0x28, 0x75, 0xef, 0x84, 0xab,
	0xec, 0x9f, 0xc7, 0xfc, 0x64, 0x80, 0xeb, 0x72, 0xff, 0x31, 0xa2, 0xaf, 0x7b, 0xcf, 0xd7, 0x1d,
	0x43, 0x18, 0x64, 0x42, 0xa9, 0x99, 0xe3, 0x5b, 0x13, 0x36, 0xed, 0xa5, 0x4d, 0x86, 0xc7, 0x30,
	0x5c, 0x46, 0x94, 0x5f, 0xbe, 0x19, 0xde, 0xaf, 0xa9, 0x16, 0x38, 0xfd, 0x60, 0xb0, 0x3b, 0x5b,
	0xdf, 0x24, 0x5e, 0xc1, 0x78, 0xc3, 0x6c, 0x8c, 0x12, 0x93, 0x25, 0xfe, 0xa5, 0x46, 0x47, 0x5e,
	0xae, 0xf9, 0xfa, 0x39, 0x40, 0x6b, 0x08, 0x1e, 0x2c, 0x4b, 0x3b, 0x7e, 0x47, 0xe1, 0x26, 0xfc,
	0xfd, 0x38, 0x1b, 0xd4, 0x37, 0x75, 0xf6, 0x15, 0x00, 0x00, 0xff, 0xff, 0x0f, 0xe4, 0x18, 0xe1,
	0x6a, 0x02, 0x00, 0x00,
}
