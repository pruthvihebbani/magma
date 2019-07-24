// Code generated by protoc-gen-go. DO NOT EDIT.
// source: context.proto

package protos

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

type Context struct {
	SessionId            string   `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Imsi                 string   `protobuf:"bytes,2,opt,name=imsi,proto3" json:"imsi,omitempty"`
	Msk                  []byte   `protobuf:"bytes,3,opt,name=msk,proto3" json:"msk,omitempty"`
	Identity             string   `protobuf:"bytes,4,opt,name=identity,proto3" json:"identity,omitempty"`
	Msisdn               string   `protobuf:"bytes,5,opt,name=msisdn,proto3" json:"msisdn,omitempty"`
	Apn                  string   `protobuf:"bytes,6,opt,name=apn,proto3" json:"apn,omitempty"`
	MacAddr              string   `protobuf:"bytes,7,opt,name=mac_addr,json=macAddr,proto3" json:"mac_addr,omitempty"`
	IpAddr               string   `protobuf:"bytes,8,opt,name=ip_addr,json=ipAddr,proto3" json:"ip_addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Context) Reset()         { *m = Context{} }
func (m *Context) String() string { return proto.CompactTextString(m) }
func (*Context) ProtoMessage()    {}
func (*Context) Descriptor() ([]byte, []int) {
	return fileDescriptor_b64063be2fc89884, []int{0}
}

func (m *Context) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Context.Unmarshal(m, b)
}
func (m *Context) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Context.Marshal(b, m, deterministic)
}
func (m *Context) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Context.Merge(m, src)
}
func (m *Context) XXX_Size() int {
	return xxx_messageInfo_Context.Size(m)
}
func (m *Context) XXX_DiscardUnknown() {
	xxx_messageInfo_Context.DiscardUnknown(m)
}

var xxx_messageInfo_Context proto.InternalMessageInfo

func (m *Context) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *Context) GetImsi() string {
	if m != nil {
		return m.Imsi
	}
	return ""
}

func (m *Context) GetMsk() []byte {
	if m != nil {
		return m.Msk
	}
	return nil
}

func (m *Context) GetIdentity() string {
	if m != nil {
		return m.Identity
	}
	return ""
}

func (m *Context) GetMsisdn() string {
	if m != nil {
		return m.Msisdn
	}
	return ""
}

func (m *Context) GetApn() string {
	if m != nil {
		return m.Apn
	}
	return ""
}

func (m *Context) GetMacAddr() string {
	if m != nil {
		return m.MacAddr
	}
	return ""
}

func (m *Context) GetIpAddr() string {
	if m != nil {
		return m.IpAddr
	}
	return ""
}

type Void struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Void) Reset()         { *m = Void{} }
func (m *Void) String() string { return proto.CompactTextString(m) }
func (*Void) ProtoMessage()    {}
func (*Void) Descriptor() ([]byte, []int) {
	return fileDescriptor_b64063be2fc89884, []int{1}
}

func (m *Void) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Void.Unmarshal(m, b)
}
func (m *Void) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Void.Marshal(b, m, deterministic)
}
func (m *Void) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Void.Merge(m, src)
}
func (m *Void) XXX_Size() int {
	return xxx_messageInfo_Void.Size(m)
}
func (m *Void) XXX_DiscardUnknown() {
	xxx_messageInfo_Void.DiscardUnknown(m)
}

var xxx_messageInfo_Void proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Context)(nil), "aaa.protos.context")
	proto.RegisterType((*Void)(nil), "aaa.protos.Void")
}

func init() { proto.RegisterFile("context.proto", fileDescriptor_b64063be2fc89884) }

var fileDescriptor_b64063be2fc89884 = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8f, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0xa9, 0x5b, 0xdb, 0xee, 0xa0, 0x20, 0x73, 0xd0, 0x28, 0x08, 0xcb, 0x82, 0xb8, 0x27,
	0x73, 0xf0, 0x09, 0xf4, 0xe6, 0x75, 0x0f, 0x1e, 0xbc, 0x2c, 0x63, 0x27, 0x96, 0x41, 0x92, 0x94,
	0x4e, 0x50, 0xf7, 0x0d, 0x7d, 0x2c, 0x69, 0x52, 0x3c, 0xe5, 0xfb, 0xff, 0x2f, 0xfc, 0x30, 0x70,
	0xde, 0xc7, 0x90, 0xdc, 0x4f, 0x7a, 0x18, 0xa7, 0x98, 0x22, 0x02, 0x11, 0x15, 0xd4, 0xed, 0x6f,
	0x05, 0xed, 0x62, 0xf1, 0x16, 0x40, 0x9d, 0xaa, 0xc4, 0x70, 0x10, 0x36, 0xd5, 0xa6, 0xda, 0xad,
	0xf7, 0xeb, 0xa5, 0x79, 0x61, 0x44, 0xa8, 0xc5, 0xab, 0x98, 0x93, 0x2c, 0x32, 0xe3, 0x05, 0xac,
	0xbc, 0x7e, 0x9a, 0xd5, 0xa6, 0xda, 0x9d, 0xed, 0x67, 0xc4, 0x1b, 0xe8, 0x84, 0x5d, 0x48, 0x92,
	0x8e, 0xa6, 0xce, 0x3f, 0xff, 0x33, 0x5e, 0x42, 0xe3, 0x55, 0x94, 0x83, 0x39, 0xcd, 0x66, 0x49,
	0xf3, 0x0a, 0x8d, 0xc1, 0x34, 0xb9, 0x9c, 0x11, 0xaf, 0xa1, 0xf3, 0xd4, 0x1f, 0x88, 0x79, 0x32,
	0x6d, 0xae, 0x5b, 0x4f, 0xfd, 0x13, 0xf3, 0x84, 0x57, 0xd0, 0xca, 0x58, 0x4c, 0x57, 0x56, 0x64,
	0x9c, 0xc5, 0xb6, 0x81, 0xfa, 0x35, 0x0a, 0x3f, 0xdf, 0xbf, 0xdd, 0x79, 0x1a, 0x3c, 0xd9, 0x0f,
	0x37, 0xd8, 0x81, 0x92, 0xfb, 0xa6, 0xa3, 0x55, 0x37, 0x7d, 0x49, 0xef, 0xd4, 0x12, 0x91, 0x2d,
	0xb7, 0xbf, 0x37, 0xf9, 0x7d, 0xfc, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xbc, 0x10, 0xdc, 0x4a, 0x1f,
	0x01, 0x00, 0x00,
}
