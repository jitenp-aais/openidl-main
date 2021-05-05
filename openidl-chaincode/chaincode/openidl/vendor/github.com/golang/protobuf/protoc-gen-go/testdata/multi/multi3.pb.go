// SPDX-License-Identifier: Apache-2.0
// Code generated by protoc-gen-go. DO NOT EDIT.
// source: multi/multi3.proto

package multitest

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

type Multi3_HatType int32

const (
	Multi3_FEDORA Multi3_HatType = 1
	Multi3_FEZ    Multi3_HatType = 2
)

var Multi3_HatType_name = map[int32]string{
	1: "FEDORA",
	2: "FEZ",
}

var Multi3_HatType_value = map[string]int32{
	"FEDORA": 1,
	"FEZ":    2,
}

func (x Multi3_HatType) Enum() *Multi3_HatType {
	p := new(Multi3_HatType)
	*p = x
	return p
}

func (x Multi3_HatType) String() string {
	return proto.EnumName(Multi3_HatType_name, int32(x))
}

func (x *Multi3_HatType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Multi3_HatType_value, data, "Multi3_HatType")
	if err != nil {
		return err
	}
	*x = Multi3_HatType(value)
	return nil
}

func (Multi3_HatType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_580398fc0bbeeaa7, []int{0, 0}
}

type Multi3 struct {
	HatType              *Multi3_HatType `protobuf:"varint,1,opt,name=hat_type,json=hatType,enum=multitest.Multi3_HatType" json:"hat_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Multi3) Reset()         { *m = Multi3{} }
func (m *Multi3) String() string { return proto.CompactTextString(m) }
func (*Multi3) ProtoMessage()    {}
func (*Multi3) Descriptor() ([]byte, []int) {
	return fileDescriptor_580398fc0bbeeaa7, []int{0}
}

func (m *Multi3) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Multi3.Unmarshal(m, b)
}
func (m *Multi3) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Multi3.Marshal(b, m, deterministic)
}
func (m *Multi3) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Multi3.Merge(m, src)
}
func (m *Multi3) XXX_Size() int {
	return xxx_messageInfo_Multi3.Size(m)
}
func (m *Multi3) XXX_DiscardUnknown() {
	xxx_messageInfo_Multi3.DiscardUnknown(m)
}

var xxx_messageInfo_Multi3 proto.InternalMessageInfo

func (m *Multi3) GetHatType() Multi3_HatType {
	if m != nil && m.HatType != nil {
		return *m.HatType
	}
	return Multi3_FEDORA
}

func init() {
	proto.RegisterEnum("multitest.Multi3_HatType", Multi3_HatType_name, Multi3_HatType_value)
	proto.RegisterType((*Multi3)(nil), "multitest.Multi3")
}

func init() {
	proto.RegisterFile("multi/multi3.proto", fileDescriptor_580398fc0bbeeaa7)
}

var fileDescriptor_580398fc0bbeeaa7 = []byte{
	// 170 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xca, 0x2d, 0xcd, 0x29,
	0xc9, 0xd4, 0x07, 0x93, 0xc6, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x9c, 0x60, 0x5e, 0x49,
	0x6a, 0x71, 0x89, 0x52, 0x1c, 0x17, 0x9b, 0x2f, 0x58, 0x4a, 0xc8, 0x84, 0x8b, 0x23, 0x23, 0xb1,
	0x24, 0xbe, 0xa4, 0xb2, 0x20, 0x55, 0x82, 0x51, 0x81, 0x51, 0x83, 0xcf, 0x48, 0x52, 0x0f, 0xae,
	0x4e, 0x0f, 0xa2, 0x48, 0xcf, 0x23, 0xb1, 0x24, 0xa4, 0xb2, 0x20, 0x35, 0x88, 0x3d, 0x03, 0xc2,
	0x50, 0x92, 0xe3, 0x62, 0x87, 0x8a, 0x09, 0x71, 0x71, 0xb1, 0xb9, 0xb9, 0xba, 0xf8, 0x07, 0x39,
	0x0a, 0x30, 0x0a, 0xb1, 0x73, 0x31, 0xbb, 0xb9, 0x46, 0x09, 0x30, 0x39, 0x39, 0x47, 0x39, 0xa6,
	0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0xa7, 0xe7, 0xe7, 0x24, 0xe6, 0xa5,
	0xeb, 0x83, 0x5d, 0x91, 0x54, 0x9a, 0x06, 0x61, 0x24, 0xeb, 0xa6, 0xa7, 0xe6, 0xe9, 0xa6, 0xe7,
	0xeb, 0x83, 0x2c, 0x4a, 0x49, 0x2c, 0x49, 0x84, 0x38, 0xd6, 0x1a, 0x6e, 0x39, 0x20, 0x00, 0x00,
	0xff, 0xff, 0xd5, 0xa4, 0x1a, 0x0e, 0xc4, 0x00, 0x00, 0x00,
}
