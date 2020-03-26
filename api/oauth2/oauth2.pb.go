// Code generated by protoc-gen-go. DO NOT EDIT.
// source: oauth2.proto

package oauth2

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

type AccessTokenReq struct {
	ClientId             string   `protobuf:"bytes,1,opt,name=ClientId,proto3" json:"ClientId,omitempty"`
	ClientSecret         string   `protobuf:"bytes,2,opt,name=ClientSecret,proto3" json:"ClientSecret,omitempty"`
	Code                 string   `protobuf:"bytes,3,opt,name=Code,proto3" json:"Code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccessTokenReq) Reset()         { *m = AccessTokenReq{} }
func (m *AccessTokenReq) String() string { return proto.CompactTextString(m) }
func (*AccessTokenReq) ProtoMessage()    {}
func (*AccessTokenReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d329d12d5a0c320, []int{0}
}

func (m *AccessTokenReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessTokenReq.Unmarshal(m, b)
}
func (m *AccessTokenReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessTokenReq.Marshal(b, m, deterministic)
}
func (m *AccessTokenReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessTokenReq.Merge(m, src)
}
func (m *AccessTokenReq) XXX_Size() int {
	return xxx_messageInfo_AccessTokenReq.Size(m)
}
func (m *AccessTokenReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessTokenReq.DiscardUnknown(m)
}

var xxx_messageInfo_AccessTokenReq proto.InternalMessageInfo

func (m *AccessTokenReq) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *AccessTokenReq) GetClientSecret() string {
	if m != nil {
		return m.ClientSecret
	}
	return ""
}

func (m *AccessTokenReq) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type AccessTokenReply struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=AccessToken,proto3" json:"AccessToken,omitempty"`
	ExpiresIn            string   `protobuf:"bytes,2,opt,name=ExpiresIn,proto3" json:"ExpiresIn,omitempty"`
	RefreshToken         string   `protobuf:"bytes,3,opt,name=RefreshToken,proto3" json:"RefreshToken,omitempty"`
	Openid               string   `protobuf:"bytes,4,opt,name=Openid,proto3" json:"Openid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccessTokenReply) Reset()         { *m = AccessTokenReply{} }
func (m *AccessTokenReply) String() string { return proto.CompactTextString(m) }
func (*AccessTokenReply) ProtoMessage()    {}
func (*AccessTokenReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d329d12d5a0c320, []int{1}
}

func (m *AccessTokenReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessTokenReply.Unmarshal(m, b)
}
func (m *AccessTokenReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessTokenReply.Marshal(b, m, deterministic)
}
func (m *AccessTokenReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessTokenReply.Merge(m, src)
}
func (m *AccessTokenReply) XXX_Size() int {
	return xxx_messageInfo_AccessTokenReply.Size(m)
}
func (m *AccessTokenReply) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessTokenReply.DiscardUnknown(m)
}

var xxx_messageInfo_AccessTokenReply proto.InternalMessageInfo

func (m *AccessTokenReply) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *AccessTokenReply) GetExpiresIn() string {
	if m != nil {
		return m.ExpiresIn
	}
	return ""
}

func (m *AccessTokenReply) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

func (m *AccessTokenReply) GetOpenid() string {
	if m != nil {
		return m.Openid
	}
	return ""
}

type VerifyTokenReq struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=AccessToken,proto3" json:"AccessToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyTokenReq) Reset()         { *m = VerifyTokenReq{} }
func (m *VerifyTokenReq) String() string { return proto.CompactTextString(m) }
func (*VerifyTokenReq) ProtoMessage()    {}
func (*VerifyTokenReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d329d12d5a0c320, []int{2}
}

func (m *VerifyTokenReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyTokenReq.Unmarshal(m, b)
}
func (m *VerifyTokenReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyTokenReq.Marshal(b, m, deterministic)
}
func (m *VerifyTokenReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyTokenReq.Merge(m, src)
}
func (m *VerifyTokenReq) XXX_Size() int {
	return xxx_messageInfo_VerifyTokenReq.Size(m)
}
func (m *VerifyTokenReq) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyTokenReq.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyTokenReq proto.InternalMessageInfo

func (m *VerifyTokenReq) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

type VerifyTokenReply struct {
	Openid               string   `protobuf:"bytes,1,opt,name=Openid,proto3" json:"Openid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyTokenReply) Reset()         { *m = VerifyTokenReply{} }
func (m *VerifyTokenReply) String() string { return proto.CompactTextString(m) }
func (*VerifyTokenReply) ProtoMessage()    {}
func (*VerifyTokenReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d329d12d5a0c320, []int{3}
}

func (m *VerifyTokenReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyTokenReply.Unmarshal(m, b)
}
func (m *VerifyTokenReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyTokenReply.Marshal(b, m, deterministic)
}
func (m *VerifyTokenReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyTokenReply.Merge(m, src)
}
func (m *VerifyTokenReply) XXX_Size() int {
	return xxx_messageInfo_VerifyTokenReply.Size(m)
}
func (m *VerifyTokenReply) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyTokenReply.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyTokenReply proto.InternalMessageInfo

func (m *VerifyTokenReply) GetOpenid() string {
	if m != nil {
		return m.Openid
	}
	return ""
}

type RefreshTokenReq struct {
	ClientId             string   `protobuf:"bytes,1,opt,name=ClientId,proto3" json:"ClientId,omitempty"`
	ClientSecret         string   `protobuf:"bytes,2,opt,name=ClientSecret,proto3" json:"ClientSecret,omitempty"`
	RefreshToken         string   `protobuf:"bytes,3,opt,name=RefreshToken,proto3" json:"RefreshToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RefreshTokenReq) Reset()         { *m = RefreshTokenReq{} }
func (m *RefreshTokenReq) String() string { return proto.CompactTextString(m) }
func (*RefreshTokenReq) ProtoMessage()    {}
func (*RefreshTokenReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d329d12d5a0c320, []int{4}
}

func (m *RefreshTokenReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RefreshTokenReq.Unmarshal(m, b)
}
func (m *RefreshTokenReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RefreshTokenReq.Marshal(b, m, deterministic)
}
func (m *RefreshTokenReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RefreshTokenReq.Merge(m, src)
}
func (m *RefreshTokenReq) XXX_Size() int {
	return xxx_messageInfo_RefreshTokenReq.Size(m)
}
func (m *RefreshTokenReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RefreshTokenReq.DiscardUnknown(m)
}

var xxx_messageInfo_RefreshTokenReq proto.InternalMessageInfo

func (m *RefreshTokenReq) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *RefreshTokenReq) GetClientSecret() string {
	if m != nil {
		return m.ClientSecret
	}
	return ""
}

func (m *RefreshTokenReq) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

type RefreshTokenReply struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=AccessToken,proto3" json:"AccessToken,omitempty"`
	ExpiresIn            string   `protobuf:"bytes,2,opt,name=ExpiresIn,proto3" json:"ExpiresIn,omitempty"`
	RefreshToken         string   `protobuf:"bytes,3,opt,name=RefreshToken,proto3" json:"RefreshToken,omitempty"`
	Openid               string   `protobuf:"bytes,4,opt,name=Openid,proto3" json:"Openid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RefreshTokenReply) Reset()         { *m = RefreshTokenReply{} }
func (m *RefreshTokenReply) String() string { return proto.CompactTextString(m) }
func (*RefreshTokenReply) ProtoMessage()    {}
func (*RefreshTokenReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d329d12d5a0c320, []int{5}
}

func (m *RefreshTokenReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RefreshTokenReply.Unmarshal(m, b)
}
func (m *RefreshTokenReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RefreshTokenReply.Marshal(b, m, deterministic)
}
func (m *RefreshTokenReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RefreshTokenReply.Merge(m, src)
}
func (m *RefreshTokenReply) XXX_Size() int {
	return xxx_messageInfo_RefreshTokenReply.Size(m)
}
func (m *RefreshTokenReply) XXX_DiscardUnknown() {
	xxx_messageInfo_RefreshTokenReply.DiscardUnknown(m)
}

var xxx_messageInfo_RefreshTokenReply proto.InternalMessageInfo

func (m *RefreshTokenReply) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *RefreshTokenReply) GetExpiresIn() string {
	if m != nil {
		return m.ExpiresIn
	}
	return ""
}

func (m *RefreshTokenReply) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

func (m *RefreshTokenReply) GetOpenid() string {
	if m != nil {
		return m.Openid
	}
	return ""
}

type RemoveTokenReq struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=AccessToken,proto3" json:"AccessToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveTokenReq) Reset()         { *m = RemoveTokenReq{} }
func (m *RemoveTokenReq) String() string { return proto.CompactTextString(m) }
func (*RemoveTokenReq) ProtoMessage()    {}
func (*RemoveTokenReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d329d12d5a0c320, []int{6}
}

func (m *RemoveTokenReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveTokenReq.Unmarshal(m, b)
}
func (m *RemoveTokenReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveTokenReq.Marshal(b, m, deterministic)
}
func (m *RemoveTokenReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveTokenReq.Merge(m, src)
}
func (m *RemoveTokenReq) XXX_Size() int {
	return xxx_messageInfo_RemoveTokenReq.Size(m)
}
func (m *RemoveTokenReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveTokenReq.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveTokenReq proto.InternalMessageInfo

func (m *RemoveTokenReq) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

type RemoveTokenReply struct {
	IsOk                 bool     `protobuf:"varint,1,opt,name=IsOk,proto3" json:"IsOk,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveTokenReply) Reset()         { *m = RemoveTokenReply{} }
func (m *RemoveTokenReply) String() string { return proto.CompactTextString(m) }
func (*RemoveTokenReply) ProtoMessage()    {}
func (*RemoveTokenReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d329d12d5a0c320, []int{7}
}

func (m *RemoveTokenReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveTokenReply.Unmarshal(m, b)
}
func (m *RemoveTokenReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveTokenReply.Marshal(b, m, deterministic)
}
func (m *RemoveTokenReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveTokenReply.Merge(m, src)
}
func (m *RemoveTokenReply) XXX_Size() int {
	return xxx_messageInfo_RemoveTokenReply.Size(m)
}
func (m *RemoveTokenReply) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveTokenReply.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveTokenReply proto.InternalMessageInfo

func (m *RemoveTokenReply) GetIsOk() bool {
	if m != nil {
		return m.IsOk
	}
	return false
}

func init() {
	proto.RegisterType((*AccessTokenReq)(nil), "oauth2.AccessTokenReq")
	proto.RegisterType((*AccessTokenReply)(nil), "oauth2.AccessTokenReply")
	proto.RegisterType((*VerifyTokenReq)(nil), "oauth2.VerifyTokenReq")
	proto.RegisterType((*VerifyTokenReply)(nil), "oauth2.VerifyTokenReply")
	proto.RegisterType((*RefreshTokenReq)(nil), "oauth2.RefreshTokenReq")
	proto.RegisterType((*RefreshTokenReply)(nil), "oauth2.RefreshTokenReply")
	proto.RegisterType((*RemoveTokenReq)(nil), "oauth2.RemoveTokenReq")
	proto.RegisterType((*RemoveTokenReply)(nil), "oauth2.RemoveTokenReply")
}

func init() {
	proto.RegisterFile("oauth2.proto", fileDescriptor_8d329d12d5a0c320)
}

var fileDescriptor_8d329d12d5a0c320 = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x53, 0xc1, 0x4e, 0x83, 0x40,
	0x10, 0x2d, 0xb5, 0x21, 0xed, 0xd8, 0xd4, 0x3a, 0x87, 0x8a, 0x8d, 0x87, 0x66, 0x0f, 0xc6, 0x78,
	0xe8, 0xa1, 0x7e, 0x41, 0x53, 0x3d, 0x70, 0x22, 0x41, 0xe3, 0x5d, 0x61, 0x9a, 0x92, 0x22, 0x8b,
	0x40, 0x8d, 0xfc, 0x84, 0x1e, 0xfd, 0x5d, 0xc3, 0x82, 0x2c, 0x43, 0x9a, 0xa8, 0x89, 0x07, 0x6f,
	0x3b, 0x6f, 0x76, 0xdf, 0xbc, 0x79, 0x3c, 0x60, 0x28, 0x1f, 0x76, 0xd9, 0x66, 0x31, 0x8f, 0x13,
	0x99, 0x49, 0x34, 0xcb, 0x4a, 0xf8, 0x30, 0x5a, 0x7a, 0x1e, 0xa5, 0xe9, 0x9d, 0xdc, 0x52, 0xe4,
	0xd2, 0x33, 0x4e, 0xa1, 0xbf, 0x0a, 0x03, 0x8a, 0x32, 0xdb, 0xb7, 0x8c, 0x99, 0x71, 0x31, 0x70,
	0xeb, 0x1a, 0x05, 0x0c, 0xcb, 0xf3, 0x2d, 0x79, 0x09, 0x65, 0x56, 0x57, 0xf5, 0x19, 0x86, 0x08,
	0xbd, 0x95, 0xf4, 0xc9, 0x3a, 0x50, 0x3d, 0x75, 0x16, 0x6f, 0x06, 0x8c, 0xd9, 0x98, 0x38, 0xcc,
	0x71, 0x06, 0x87, 0x0d, 0xac, 0x9a, 0xd5, 0x84, 0xf0, 0x0c, 0x06, 0x37, 0xaf, 0x71, 0x90, 0x50,
	0x6a, 0x47, 0xd5, 0x2c, 0x0d, 0x14, 0x62, 0x5c, 0x5a, 0x27, 0x94, 0x6e, 0x4a, 0x82, 0x72, 0x20,
	0xc3, 0x70, 0x02, 0xa6, 0x13, 0x53, 0x14, 0xf8, 0x56, 0x4f, 0x75, 0xab, 0x4a, 0x2c, 0x60, 0x74,
	0x4f, 0x49, 0xb0, 0xce, 0xeb, 0xb5, 0xbf, 0x55, 0x23, 0x2e, 0x61, 0xcc, 0xde, 0x14, 0x3b, 0x68,
	0x7e, 0x83, 0xf1, 0xef, 0xe0, 0xa8, 0xa9, 0xe3, 0x2f, 0x7c, 0xfd, 0xc1, 0xba, 0xe2, 0xdd, 0x80,
	0x63, 0x3e, 0xf7, 0x1f, 0x18, 0xed, 0xd2, 0x93, 0x7c, 0xa1, 0x5f, 0x18, 0x7d, 0x0e, 0x63, 0xf6,
	0xa6, 0xd8, 0x01, 0xa1, 0x67, 0xa7, 0xce, 0x56, 0x5d, 0xef, 0xbb, 0xea, 0xbc, 0xf8, 0xe8, 0x82,
	0xe9, 0xa8, 0x18, 0xe3, 0x92, 0x91, 0xe2, 0x64, 0x5e, 0x85, 0x9d, 0x67, 0x7b, 0x6a, 0xed, 0xc5,
	0xe3, 0x30, 0x17, 0x9d, 0x82, 0xa2, 0xf1, 0x79, 0x35, 0x05, 0xcf, 0x89, 0xa6, 0x68, 0x67, 0x41,
	0x74, 0xf0, 0x9a, 0x1b, 0x85, 0x27, 0x5f, 0x77, 0x5b, 0x59, 0x98, 0x9e, 0xee, 0x6f, 0xd4, 0x42,
	0x1a, 0xeb, 0x6b, 0x21, 0xdc, 0x47, 0x2d, 0xa4, 0xed, 0x95, 0xe8, 0x3c, 0x9a, 0xea, 0x27, 0xbf,
	0xfa, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x8f, 0x89, 0x5d, 0xe2, 0xf4, 0x03, 0x00, 0x00,
}