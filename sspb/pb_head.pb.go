// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb_head.proto

package sspb

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

type PBHead struct {
	OpCode               int32    `protobuf:"varint,1,opt,name=opCode,proto3" json:"opCode,omitempty"`
	ErrCode              int32    `protobuf:"varint,2,opt,name=errCode,proto3" json:"errCode,omitempty"`
	ErrMessage           string   `protobuf:"bytes,3,opt,name=errMessage,proto3" json:"errMessage,omitempty"`
	Rpcid                uint64   `protobuf:"varint,6,opt,name=rpcid,proto3" json:"rpcid,omitempty"`
	Flags                uint64   `protobuf:"varint,7,opt,name=flags,proto3" json:"flags,omitempty"`
	Ext                  string   `protobuf:"bytes,11,opt,name=ext,proto3" json:"ext,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PBHead) Reset()         { *m = PBHead{} }
func (m *PBHead) String() string { return proto.CompactTextString(m) }
func (*PBHead) ProtoMessage()    {}
func (*PBHead) Descriptor() ([]byte, []int) {
	return fileDescriptor_863d2b855dd8538c, []int{0}
}

func (m *PBHead) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PBHead.Unmarshal(m, b)
}
func (m *PBHead) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PBHead.Marshal(b, m, deterministic)
}
func (m *PBHead) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PBHead.Merge(m, src)
}
func (m *PBHead) XXX_Size() int {
	return xxx_messageInfo_PBHead.Size(m)
}
func (m *PBHead) XXX_DiscardUnknown() {
	xxx_messageInfo_PBHead.DiscardUnknown(m)
}

var xxx_messageInfo_PBHead proto.InternalMessageInfo

func (m *PBHead) GetOpCode() int32 {
	if m != nil {
		return m.OpCode
	}
	return 0
}

func (m *PBHead) GetErrCode() int32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *PBHead) GetErrMessage() string {
	if m != nil {
		return m.ErrMessage
	}
	return ""
}

func (m *PBHead) GetRpcid() uint64 {
	if m != nil {
		return m.Rpcid
	}
	return 0
}

func (m *PBHead) GetFlags() uint64 {
	if m != nil {
		return m.Flags
	}
	return 0
}

func (m *PBHead) GetExt() string {
	if m != nil {
		return m.Ext
	}
	return ""
}

func init() {
	proto.RegisterType((*PBHead)(nil), "sspb.PBHead")
}

func init() { proto.RegisterFile("pb_head.proto", fileDescriptor_863d2b855dd8538c) }

var fileDescriptor_863d2b855dd8538c = []byte{
	// 160 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x48, 0x8a, 0xcf,
	0x48, 0x4d, 0x4c, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x29, 0x2e, 0x2e, 0x48, 0x52,
	0x9a, 0xc1, 0xc8, 0xc5, 0x16, 0xe0, 0xe4, 0x91, 0x9a, 0x98, 0x22, 0x24, 0xc6, 0xc5, 0x96, 0x5f,
	0xe0, 0x9c, 0x9f, 0x92, 0x2a, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x1a, 0x04, 0xe5, 0x09, 0x49, 0x70,
	0xb1, 0xa7, 0x16, 0x15, 0x81, 0x25, 0x98, 0xc0, 0x12, 0x30, 0xae, 0x90, 0x1c, 0x17, 0x57, 0x6a,
	0x51, 0x91, 0x6f, 0x6a, 0x71, 0x71, 0x62, 0x7a, 0xaa, 0x04, 0xb3, 0x02, 0xa3, 0x06, 0x67, 0x10,
	0x92, 0x88, 0x90, 0x08, 0x17, 0x6b, 0x51, 0x41, 0x72, 0x66, 0x8a, 0x04, 0x9b, 0x02, 0xa3, 0x06,
	0x4b, 0x10, 0x84, 0x03, 0x12, 0x4d, 0xcb, 0x49, 0x4c, 0x2f, 0x96, 0x60, 0x87, 0x88, 0x82, 0x39,
	0x42, 0x02, 0x5c, 0xcc, 0xa9, 0x15, 0x25, 0x12, 0xdc, 0x60, 0x43, 0x40, 0xcc, 0x24, 0x36, 0xb0,
	0x3b, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc9, 0x69, 0xe5, 0xc2, 0xb8, 0x00, 0x00, 0x00,
}