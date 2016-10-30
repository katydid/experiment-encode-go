// Code generated by protoc-gen-gogo.
// source: msg.proto
// DO NOT EDIT!

/*
Package prototests is a generated protocol buffer package.

It is generated from these files:
	msg.proto

It has these top-level messages:
	BigMsg
	SmallMsg
*/
package prototests

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import github_com_gogo_protobuf_protoc_gen_gogo_descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
import compress_gzip "compress/gzip"
import bytes "bytes"
import io_ioutil "io/ioutil"

import strings "strings"
import sort "sort"
import strconv "strconv"
import reflect "reflect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// BigMsg contains a field and a message field.
type BigMsg struct {
	Field            *int64    `protobuf:"varint,1,opt,name=Field" json:"Field,omitempty"`
	Msg              *SmallMsg `protobuf:"bytes,3,opt,name=Msg" json:"Msg,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *BigMsg) Reset()                    { *m = BigMsg{} }
func (m *BigMsg) String() string            { return proto.CompactTextString(m) }
func (*BigMsg) ProtoMessage()               {}
func (*BigMsg) Descriptor() ([]byte, []int) { return fileDescriptorMsg, []int{0} }

func (m *BigMsg) GetField() int64 {
	if m != nil && m.Field != nil {
		return *m.Field
	}
	return 0
}

func (m *BigMsg) GetMsg() *SmallMsg {
	if m != nil {
		return m.Msg
	}
	return nil
}

// SmallMsg only contains some native fields.
type SmallMsg struct {
	ScarBusStop      *string  `protobuf:"bytes,1,opt,name=ScarBusStop" json:"ScarBusStop,omitempty"`
	FlightParachute  []uint32 `protobuf:"fixed32,12,rep,name=FlightParachute" json:"FlightParachute,omitempty"`
	MapShark         *string  `protobuf:"bytes,18,opt,name=MapShark" json:"MapShark,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *SmallMsg) Reset()                    { *m = SmallMsg{} }
func (m *SmallMsg) String() string            { return proto.CompactTextString(m) }
func (*SmallMsg) ProtoMessage()               {}
func (*SmallMsg) Descriptor() ([]byte, []int) { return fileDescriptorMsg, []int{1} }

func (m *SmallMsg) GetScarBusStop() string {
	if m != nil && m.ScarBusStop != nil {
		return *m.ScarBusStop
	}
	return ""
}

func (m *SmallMsg) GetFlightParachute() []uint32 {
	if m != nil {
		return m.FlightParachute
	}
	return nil
}

func (m *SmallMsg) GetMapShark() string {
	if m != nil && m.MapShark != nil {
		return *m.MapShark
	}
	return ""
}

func init() {
	proto.RegisterType((*BigMsg)(nil), "prototests.BigMsg")
	proto.RegisterType((*SmallMsg)(nil), "prototests.SmallMsg")
}
func (this *BigMsg) Description() (desc *github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet) {
	return MsgDescription()
}
func (this *SmallMsg) Description() (desc *github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet) {
	return MsgDescription()
}
func MsgDescription() (desc *github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet) {
	d := &github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet{}
	var gzipped = []byte{
		// 3477 bytes of a gzipped FileDescriptorSet
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xcc, 0x5a, 0x5d, 0x6c, 0x1b, 0xd7,
		0x95, 0x0e, 0xff, 0x24, 0xf2, 0x90, 0xa2, 0x46, 0x57, 0x8a, 0x4c, 0x2b, 0x89, 0x2d, 0x2b, 0x4e,
		0xa2, 0x24, 0x1b, 0x39, 0x70, 0x6c, 0xc7, 0xa6, 0x37, 0x31, 0x28, 0x8a, 0x56, 0x64, 0x88, 0x12,
		0x77, 0x28, 0x25, 0x4e, 0xf2, 0x30, 0xb8, 0x1a, 0x5e, 0x52, 0x63, 0x0f, 0x67, 0xb8, 0x33, 0x43,
		0xdb, 0xf2, 0x93, 0x83, 0xec, 0x0f, 0x82, 0x60, 0x77, 0xb3, 0xbb, 0x05, 0x9a, 0xff, 0xa4, 0x01,
		0xda, 0xb4, 0xe9, 0x5f, 0xd2, 0xbf, 0x87, 0x3e, 0xf5, 0x25, 0xed, 0x5b, 0x81, 0xbc, 0xf7, 0x25,
		0x45, 0x80, 0xfe, 0xa5, 0x6d, 0xda, 0x1a, 0x68, 0x01, 0xbf, 0x14, 0xf7, 0x6f, 0x38, 0xfc, 0x91,
		0x87, 0x0a, 0x90, 0xa4, 0x4f, 0xe2, 0x3d, 0xf7, 0x7c, 0xdf, 0x9c, 0x7b, 0xee, 0xb9, 0xe7, 0x9c,
		0xb9, 0x23, 0x78, 0xe6, 0x18, 0xcc, 0x36, 0x6c, 0xbb, 0x61, 0x92, 0x23, 0x2d, 0xc7, 0xf6, 0xec,
		0xad, 0x76, 0xfd, 0x48, 0x8d, 0xb8, 0xba, 0x63, 0xb4, 0x3c, 0xdb, 0x59, 0x60, 0x32, 0x34, 0xce,
		0x35, 0x16, 0xa4, 0xc6, 0x5c, 0x19, 0x26, 0xce, 0x1a, 0x26, 0x59, 0xf2, 0x15, 0xab, 0xc4, 0x43,
		0x27, 0x21, 0x5e, 0x37, 0x4c, 0x92, 0x8b, 0xcc, 0xc6, 0xe6, 0xd3, 0x47, 0x0f, 0x2f, 0xf4, 0x80,
		0x16, 0xba, 0x11, 0x15, 0x2a, 0x56, 0x19, 0x62, 0xee, 0xa3, 0x38, 0x4c, 0x0e, 0x98, 0x45, 0x08,
		0xe2, 0x16, 0x6e, 0x52, 0xc6, 0xc8, 0x7c, 0x4a, 0x65, 0xbf, 0x51, 0x0e, 0x46, 0x5b, 0x58, 0xbf,
		0x88, 0x1b, 0x24, 0x17, 0x65, 0x62, 0x39, 0x44, 0x07, 0x00, 0x6a, 0xa4, 0x45, 0xac, 0x1a, 0xb1,
		0xf4, 0x9d, 0x5c, 0x6c, 0x36, 0x36, 0x9f, 0x52, 0x03, 0x12, 0x74, 0x3f, 0x4c, 0xb4, 0xda, 0x5b,
		0xa6, 0xa1, 0x6b, 0x01, 0x35, 0x98, 0x8d, 0xcd, 0x27, 0x54, 0x85, 0x4f, 0x2c, 0x75, 0x94, 0xef,
		0x81, 0xf1, 0xcb, 0x04, 0x5f, 0x0c, 0xaa, 0xa6, 0x99, 0x6a, 0x96, 0x8a, 0x03, 0x8a, 0x45, 0xc8,
		0x34, 0x89, 0xeb, 0xe2, 0x06, 0xd1, 0xbc, 0x9d, 0x16, 0xc9, 0xc5, 0xd9, 0xea, 0x67, 0xfb, 0x56,
		0xdf, 0xbb, 0xf2, 0xb4, 0x40, 0x6d, 0xec, 0xb4, 0x08, 0x2a, 0x40, 0x8a, 0x58, 0xed, 0x26, 0x67,
		0x48, 0xec, 0xe2, 0xbf, 0x92, 0xd5, 0x6e, 0xf6, 0xb2, 0x24, 0x29, 0x4c, 0x50, 0x8c, 0xba, 0xc4,
		0xb9, 0x64, 0xe8, 0x24, 0x37, 0xc2, 0x08, 0xee, 0xe9, 0x23, 0xa8, 0xf2, 0xf9, 0x5e, 0x0e, 0x89,
		0x43, 0x45, 0x48, 0x91, 0x2b, 0x1e, 0xb1, 0x5c, 0xc3, 0xb6, 0x72, 0xa3, 0x8c, 0xe4, 0xae, 0x01,
		0xbb, 0x48, 0xcc, 0x5a, 0x2f, 0x45, 0x07, 0x87, 0x4e, 0xc0, 0xa8, 0xdd, 0xf2, 0x0c, 0xdb, 0x72,
		0x73, 0xc9, 0xd9, 0xc8, 0x7c, 0xfa, 0xe8, 0xed, 0x03, 0x03, 0x61, 0x9d, 0xeb, 0xa8, 0x52, 0x19,
		0xad, 0x80, 0xe2, 0xda, 0x6d, 0x47, 0x27, 0x9a, 0x6e, 0xd7, 0x88, 0x66, 0x58, 0x75, 0x3b, 0x97,
		0x62, 0x04, 0x07, 0xfb, 0x17, 0xc2, 0x14, 0x8b, 0x76, 0x8d, 0xac, 0x58, 0x75, 0x5b, 0xcd, 0xba,
		0x5d, 0x63, 0x34, 0x0d, 0x23, 0xee, 0x8e, 0xe5, 0xe1, 0x2b, 0xb9, 0x0c, 0x8b, 0x10, 0x31, 0x9a,
		0xfb, 0x6b, 0x02, 0xc6, 0x87, 0x09, 0xb1, 0xd3, 0x90, 0xa8, 0xd3, 0x55, 0xe6, 0xa2, 0x7b, 0xf1,
		0x01, 0xc7, 0x74, 0x3b, 0x71, 0xe4, 0x53, 0x3a, 0xb1, 0x00, 0x69, 0x8b, 0xb8, 0x1e, 0xa9, 0xf1,
		0x88, 0x88, 0x0d, 0x19, 0x53, 0xc0, 0x41, 0xfd, 0x21, 0x15, 0xff, 0x54, 0x21, 0x75, 0x1e, 0xc6,
		0x7d, 0x93, 0x34, 0x07, 0x5b, 0x0d, 0x19, 0x9b, 0x47, 0xc2, 0x2c, 0x59, 0x28, 0x49, 0x9c, 0x4a,
		0x61, 0x6a, 0x96, 0x74, 0x8d, 0xd1, 0x12, 0x80, 0x6d, 0x11, 0xbb, 0xae, 0xd5, 0x88, 0x6e, 0xe6,
		0x92, 0xbb, 0x78, 0x69, 0x9d, 0xaa, 0xf4, 0x79, 0xc9, 0xe6, 0x52, 0xdd, 0x44, 0xa7, 0x3a, 0xa1,
		0x36, 0xba, 0x4b, 0xa4, 0x94, 0xf9, 0x21, 0xeb, 0x8b, 0xb6, 0x4d, 0xc8, 0x3a, 0x84, 0xc6, 0x3d,
		0xa9, 0x89, 0x95, 0xa5, 0x98, 0x11, 0x0b, 0xa1, 0x2b, 0x53, 0x05, 0x8c, 0x2f, 0x6c, 0xcc, 0x09,
		0x0e, 0xd1, 0x9d, 0xe0, 0x0b, 0x34, 0x16, 0x56, 0xc0, 0xb2, 0x50, 0x46, 0x0a, 0xd7, 0x70, 0x93,
		0xcc, 0x9c, 0x84, 0x6c, 0xb7, 0x7b, 0xd0, 0x14, 0x24, 0x5c, 0x0f, 0x3b, 0x1e, 0x8b, 0xc2, 0x84,
		0xca, 0x07, 0x48, 0x81, 0x18, 0xb1, 0x6a, 0x2c, 0xcb, 0x25, 0x54, 0xfa, 0x73, 0xe6, 0x61, 0x18,
		0xeb, 0x7a, 0xfc, 0xb0, 0xc0, 0xb9, 0x17, 0x47, 0x60, 0x6a, 0x50, 0xcc, 0x0d, 0x0c, 0xff, 0x69,
		0x18, 0xb1, 0xda, 0xcd, 0x2d, 0xe2, 0xe4, 0x62, 0x8c, 0x41, 0x8c, 0x50, 0x01, 0x12, 0x26, 0xde,
		0x22, 0x66, 0x2e, 0x3e, 0x1b, 0x99, 0xcf, 0x1e, 0xbd, 0x7f, 0xa8, 0xa8, 0x5e, 0x58, 0xa5, 0x10,
		0x95, 0x23, 0xd1, 0xa3, 0x10, 0x17, 0x29, 0x8e, 0x32, 0xdc, 0x37, 0x1c, 0x03, 0x8d, 0x45, 0x95,
		0xe1, 0xd0, 0x6d, 0x90, 0xa2, 0x7f, 0xb9, 0x6f, 0x47, 0x98, 0xcd, 0x49, 0x2a, 0xa0, 0x7e, 0x45,
		0x33, 0x90, 0x64, 0x61, 0x56, 0x23, 0xb2, 0x34, 0xf8, 0x63, 0xba, 0x31, 0x35, 0x52, 0xc7, 0x6d,
		0xd3, 0xd3, 0x2e, 0x61, 0xb3, 0x4d, 0x58, 0xc0, 0xa4, 0xd4, 0x8c, 0x10, 0x3e, 0x4e, 0x65, 0xe8,
		0x20, 0xa4, 0x79, 0x54, 0x1a, 0x56, 0x8d, 0x5c, 0x61, 0xd9, 0x27, 0xa1, 0xf2, 0x40, 0x5d, 0xa1,
		0x12, 0xfa, 0xf8, 0x0b, 0xae, 0x6d, 0xc9, 0xad, 0x65, 0x8f, 0xa0, 0x02, 0xf6, 0xf8, 0x87, 0x7b,
		0x13, 0xdf, 0x1d, 0x83, 0x97, 0xd7, 0x1b, 0x8b, 0x73, 0x3f, 0x8a, 0x42, 0x9c, 0x9d, 0xb7, 0x71,
		0x48, 0x6f, 0x3c, 0x59, 0x29, 0x69, 0x4b, 0xeb, 0x9b, 0x8b, 0xab, 0x25, 0x25, 0x82, 0xb2, 0x00,
		0x4c, 0x70, 0x76, 0x75, 0xbd, 0xb0, 0xa1, 0x44, 0xfd, 0xf1, 0xca, 0xda, 0xc6, 0x89, 0x63, 0x4a,
		0xcc, 0x07, 0x6c, 0x72, 0x41, 0x3c, 0xa8, 0xf0, 0xd0, 0x51, 0x25, 0x81, 0x14, 0xc8, 0x70, 0x82,
		0x95, 0xf3, 0xa5, 0xa5, 0x13, 0xc7, 0x94, 0x91, 0x6e, 0xc9, 0x43, 0x47, 0x95, 0x51, 0x34, 0x06,
		0x29, 0x26, 0x59, 0x5c, 0x5f, 0x5f, 0x55, 0x92, 0x3e, 0x67, 0x75, 0x43, 0x5d, 0x59, 0x5b, 0x56,
		0x52, 0x3e, 0xe7, 0xb2, 0xba, 0xbe, 0x59, 0x51, 0xc0, 0x67, 0x28, 0x97, 0xaa, 0xd5, 0xc2, 0x72,
		0x49, 0x49, 0xfb, 0x1a, 0x8b, 0x4f, 0x6e, 0x94, 0xaa, 0x4a, 0xa6, 0xcb, 0xac, 0x87, 0x8e, 0x2a,
		0x63, 0xfe, 0x23, 0x4a, 0x6b, 0x9b, 0x65, 0x25, 0x8b, 0x26, 0x60, 0x8c, 0x3f, 0x42, 0x1a, 0x31,
		0xde, 0x23, 0x3a, 0x71, 0x4c, 0x51, 0x3a, 0x86, 0x70, 0x96, 0x89, 0x2e, 0xc1, 0x89, 0x63, 0x0a,
		0x9a, 0x2b, 0x42, 0x82, 0x45, 0x17, 0x42, 0x90, 0x5d, 0x2d, 0x2c, 0x96, 0x56, 0xb5, 0xf5, 0xca,
		0xc6, 0xca, 0xfa, 0x5a, 0x61, 0x55, 0x89, 0x74, 0x64, 0x6a, 0xe9, 0x5f, 0x36, 0x57, 0xd4, 0xd2,
		0x92, 0x12, 0x0d, 0xca, 0x2a, 0xa5, 0xc2, 0x46, 0x69, 0x49, 0x89, 0xcd, 0xe9, 0x30, 0x35, 0x28,
		0xcf, 0x0c, 0x3c, 0x19, 0x81, 0x2d, 0x8e, 0xee, 0xb2, 0xc5, 0x8c, 0xab, 0x6f, 0x8b, 0xdf, 0x8a,
		0xc0, 0xe4, 0x80, 0x5c, 0x3b, 0xf0, 0x21, 0x67, 0x20, 0xc1, 0x43, 0x94, 0x57, 0x9f, 0x7b, 0x07,
		0x26, 0x6d, 0x16, 0xb0, 0x7d, 0x15, 0x88, 0xe1, 0x82, 0x15, 0x38, 0xb6, 0x4b, 0x05, 0xa6, 0x14,
		0x7d, 0x46, 0x3e, 0x1b, 0x81, 0xdc, 0x6e, 0xdc, 0x21, 0x89, 0x22, 0xda, 0x95, 0x28, 0x4e, 0xf7,
		0x1a, 0x70, 0x68, 0xf7, 0x35, 0xf4, 0x59, 0xf1, 0x76, 0x04, 0xa6, 0x07, 0x37, 0x2a, 0x03, 0x6d,
		0x78, 0x14, 0x46, 0x9a, 0xc4, 0xdb, 0xb6, 0x65, 0xb1, 0xbe, 0x7b, 0x40, 0x09, 0xa0, 0xd3, 0xbd,
		0xbe, 0x12, 0xa8, 0x60, 0x0d, 0x89, 0xed, 0xd6, 0x6d, 0x70, 0x6b, 0xfa, 0x2c, 0x7d, 0x2e, 0x0a,
		0xb7, 0x0e, 0x24, 0x1f, 0x68, 0xe8, 0x1d, 0x00, 0x86, 0xd5, 0x6a, 0x7b, 0xbc, 0x20, 0xf3, 0xfc,
		0x94, 0x62, 0x12, 0x76, 0xf6, 0x69, 0xee, 0x69, 0x7b, 0xfe, 0x7c, 0x8c, 0xcd, 0x03, 0x17, 0x31,
		0x85, 0x93, 0x1d, 0x43, 0xe3, 0xcc, 0xd0, 0x03, 0xbb, 0xac, 0xb4, 0xaf, 0xd6, 0x3d, 0x08, 0x8a,
		0x6e, 0x1a, 0xc4, 0xf2, 0x34, 0xd7, 0x73, 0x08, 0x6e, 0x1a, 0x56, 0x83, 0x25, 0xe0, 0x64, 0x3e,
		0x51, 0xc7, 0xa6, 0x4b, 0xd4, 0x71, 0x3e, 0x5d, 0x95, 0xb3, 0x14, 0xc1, 0xaa, 0x8c, 0x13, 0x40,
		0x8c, 0x74, 0x21, 0xf8, 0xb4, 0x8f, 0x98, 0x7b, 0x7e, 0x14, 0xd2, 0x81, 0xb6, 0x0e, 0x1d, 0x82,
		0xcc, 0x05, 0x7c, 0x09, 0x6b, 0xb2, 0x55, 0xe7, 0x9e, 0x48, 0x53, 0x59, 0x45, 0xb4, 0xeb, 0x0f,
		0xc2, 0x14, 0x53, 0xb1, 0xdb, 0x1e, 0x71, 0x34, 0xdd, 0xc4, 0xae, 0xcb, 0x9c, 0x96, 0x64, 0xaa,
		0x88, 0xce, 0xad, 0xd3, 0xa9, 0xa2, 0x9c, 0x41, 0xc7, 0x61, 0x92, 0x21, 0x9a, 0x6d, 0xd3, 0x33,
		0x5a, 0x26, 0xd1, 0xe8, 0xcb, 0x83, 0xcb, 0x12, 0xb1, 0x6f, 0xd9, 0x04, 0xd5, 0x28, 0x0b, 0x05,
		0x6a, 0x91, 0x8b, 0x96, 0xe1, 0x0e, 0x06, 0x6b, 0x10, 0x8b, 0x38, 0xd8, 0x23, 0x1a, 0xf9, 0xd7,
		0x36, 0x36, 0x5d, 0x0d, 0x5b, 0x35, 0x6d, 0x1b, 0xbb, 0xdb, 0xb9, 0xa9, 0x20, 0xc1, 0x7e, 0xaa,
		0xbb, 0x2c, 0x54, 0x4b, 0x4c, 0xb3, 0x60, 0xd5, 0x1e, 0xc3, 0xee, 0x36, 0xca, 0xc3, 0x34, 0x23,
		0x72, 0x3d, 0xc7, 0xb0, 0x1a, 0x9a, 0xbe, 0x4d, 0xf4, 0x8b, 0x5a, 0xdb, 0xab, 0x9f, 0xcc, 0xdd,
		0x16, 0x64, 0x60, 0x46, 0x56, 0x99, 0x4e, 0x91, 0xaa, 0x6c, 0x7a, 0xf5, 0x93, 0xa8, 0x0a, 0x19,
		0xba, 0x1f, 0x4d, 0xe3, 0x2a, 0xd1, 0xea, 0xb6, 0xc3, 0x8a, 0x4b, 0x76, 0xc0, 0xe1, 0x0e, 0x38,
		0x71, 0x61, 0x5d, 0x00, 0xca, 0x76, 0x8d, 0xe4, 0x13, 0xd5, 0x4a, 0xa9, 0xb4, 0xa4, 0xa6, 0x25,
		0xcb, 0x59, 0xdb, 0xa1, 0x31, 0xd5, 0xb0, 0x7d, 0x1f, 0xa7, 0x79, 0x4c, 0x35, 0x6c, 0xe9, 0xe1,
		0xe3, 0x30, 0xa9, 0xeb, 0x7c, 0xd9, 0x86, 0xae, 0x89, 0x2e, 0xdf, 0xcd, 0x29, 0x5d, 0xfe, 0xd2,
		0xf5, 0x65, 0xae, 0x20, 0xc2, 0xdc, 0x45, 0xa7, 0xe0, 0xd6, 0x8e, 0xbf, 0x82, 0xc0, 0x89, 0xbe,
		0x55, 0xf6, 0x42, 0x8f, 0xc3, 0x64, 0x6b, 0xa7, 0x1f, 0x88, 0xba, 0x9e, 0xd8, 0xda, 0xe9, 0x85,
		0xdd, 0xc5, 0xde, 0xdc, 0x1c, 0xa2, 0x63, 0x8f, 0xd4, 0x72, 0xfb, 0x82, 0xda, 0x81, 0x09, 0x74,
		0x04, 0x14, 0x5d, 0xd7, 0x88, 0x85, 0xb7, 0x4c, 0xa2, 0x61, 0x87, 0x58, 0xd8, 0xcd, 0x1d, 0x0c,
		0x2a, 0x67, 0x75, 0xbd, 0xc4, 0x66, 0x0b, 0x6c, 0x12, 0xdd, 0x07, 0x13, 0xf6, 0xd6, 0x05, 0x9d,
		0x07, 0x97, 0xd6, 0x72, 0x48, 0xdd, 0xb8, 0x92, 0x3b, 0xcc, 0xdc, 0x34, 0x4e, 0x27, 0x58, 0x68,
		0x55, 0x98, 0x18, 0xdd, 0x0b, 0x8a, 0xee, 0x6e, 0x63, 0xa7, 0xc5, 0xaa, 0xbb, 0xdb, 0xc2, 0x3a,
		0xc9, 0xdd, 0xc5, 0x55, 0xb9, 0x7c, 0x4d, 0x8a, 0xd1, 0x79, 0x98, 0x6a, 0x5b, 0x86, 0xe5, 0x11,
		0xa7, 0xe5, 0x10, 0xda, 0xa4, 0xf3, 0x93, 0x96, 0xfb, 0xd5, 0xe8, 0x2e, 0x6d, 0xf6, 0x66, 0x50,
		0x9b, 0xef, 0xae, 0x3a, 0xd9, 0xee, 0x17, 0xce, 0xe5, 0x21, 0x13, 0xdc, 0x74, 0x94, 0x02, 0xbe,
		0xed, 0x4a, 0x84, 0xd6, 0xd0, 0xe2, 0xfa, 0x12, 0xad, 0x7e, 0x4f, 0x95, 0x94, 0x28, 0xad, 0xc2,
		0xab, 0x2b, 0x1b, 0x25, 0x4d, 0xdd, 0x5c, 0xdb, 0x58, 0x29, 0x97, 0x94, 0xd8, 0x7d, 0xa9, 0xe4,
		0xaf, 0x47, 0x95, 0x6b, 0xd7, 0xae, 0x5d, 0x8b, 0xce, 0xbd, 0x1f, 0x85, 0x6c, 0x77, 0xe7, 0x8b,
		0xfe, 0x19, 0xf6, 0xc9, 0xd7, 0x54, 0x97, 0x78, 0xda, 0x65, 0xc3, 0x61, 0x71, 0xd8, 0xc4, 0xbc,
		0x77, 0xf4, 0x5d, 0x38, 0x25, 0xb4, 0xaa, 0xc4, 0x7b, 0xc2, 0x70, 0x68, 0x94, 0x35, 0xb1, 0x87,
		0x56, 0xe1, 0xa0, 0x65, 0x6b, 0xae, 0x87, 0xad, 0x1a, 0x76, 0x6a, 0x5a, 0xe7, 0x82, 0x40, 0xc3,
		0xba, 0x4e, 0x5c, 0xd7, 0xe6, 0x25, 0xc0, 0x67, 0xb9, 0xdd, 0xb2, 0xab, 0x42, 0xb9, 0x93, 0x1b,
		0x0b, 0x42, 0xb5, 0x67, 0xbb, 0x63, 0xbb, 0x6d, 0xf7, 0x6d, 0x90, 0x6a, 0xe2, 0x96, 0x46, 0x2c,
		0xcf, 0xd9, 0x61, 0xfd, 0x5a, 0x52, 0x4d, 0x36, 0x71, 0xab, 0x44, 0xc7, 0x9f, 0xdd, 0x1e, 0x04,
		0xfd, 0xf8, 0x8b, 0x18, 0x64, 0x82, 0x3d, 0x1b, 0x6d, 0x81, 0x75, 0x96, 0x9f, 0x23, 0xec, 0xf8,
		0xde, 0x79, 0xd3, 0x0e, 0x6f, 0xa1, 0x48, 0x13, 0x77, 0x7e, 0x84, 0x77, 0x52, 0x2a, 0x47, 0xd2,
		0xa2, 0x49, 0x0f, 0x2c, 0xe1, 0xfd, 0x79, 0x52, 0x15, 0x23, 0xb4, 0x0c, 0x23, 0x17, 0x5c, 0xc6,
		0x3d, 0xc2, 0xb8, 0x0f, 0xdf, 0x9c, 0xfb, 0x5c, 0x95, 0x91, 0xa7, 0xce, 0x55, 0xb5, 0xb5, 0x75,
		0xb5, 0x5c, 0x58, 0x55, 0x05, 0x1c, 0xed, 0x87, 0xb8, 0x89, 0xaf, 0xee, 0x74, 0xa7, 0x78, 0x26,
		0x1a, 0xd6, 0xf1, 0xfb, 0x21, 0x7e, 0x99, 0xe0, 0x8b, 0xdd, 0x89, 0x95, 0x89, 0x3e, 0xc3, 0xd0,
		0x3f, 0x02, 0x09, 0xe6, 0x2f, 0x04, 0x20, 0x3c, 0xa6, 0xdc, 0x82, 0x92, 0x10, 0x2f, 0xae, 0xab,
		0x34, 0xfc, 0x15, 0xc8, 0x70, 0xa9, 0x56, 0x59, 0x29, 0x15, 0x4b, 0x4a, 0x74, 0xee, 0x38, 0x8c,
		0x70, 0x27, 0xd0, 0xa3, 0xe1, 0xbb, 0x41, 0xb9, 0x45, 0x0c, 0x05, 0x47, 0x44, 0xce, 0x6e, 0x96,
		0x17, 0x4b, 0xaa, 0x12, 0x0d, 0x6e, 0xaf, 0x0b, 0x99, 0x60, 0xbb, 0xf6, 0xf9, 0xc4, 0xd4, 0x8f,
		0x23, 0x90, 0x0e, 0xb4, 0x5f, 0xb4, 0xf0, 0x63, 0xd3, 0xb4, 0x2f, 0x6b, 0xd8, 0x34, 0xb0, 0x2b,
		0x82, 0x02, 0x98, 0xa8, 0x40, 0x25, 0xc3, 0x6e, 0xda, 0xe7, 0x62, 0xfc, 0xeb, 0x11, 0x50, 0x7a,
		0x5b, 0xb7, 0x1e, 0x03, 0x23, 0x5f, 0xa8, 0x81, 0xaf, 0x46, 0x20, 0xdb, 0xdd, 0xaf, 0xf5, 0x98,
		0x77, 0xe8, 0x0b, 0x35, 0xef, 0x95, 0x08, 0x8c, 0x75, 0x75, 0x69, 0xff, 0x50, 0xd6, 0xbd, 0x1c,
		0x83, 0xc9, 0x01, 0x38, 0x54, 0x10, 0xed, 0x2c, 0xef, 0xb0, 0x1f, 0x18, 0xe6, 0x59, 0x0b, 0xb4,
		0x5a, 0x56, 0xb0, 0xe3, 0x89, 0xee, 0xf7, 0x5e, 0x50, 0x8c, 0x1a, 0xb1, 0x3c, 0xa3, 0x6e, 0x10,
		0x47, 0xbc, 0x82, 0xf3, 0x1e, 0x77, 0xbc, 0x23, 0xe7, 0x6f, 0xe1, 0xff, 0x04, 0xa8, 0x65, 0xbb,
		0x86, 0x67, 0x5c, 0x22, 0x9a, 0x61, 0xc9, 0xf7, 0x75, 0xda, 0xf3, 0xc6, 0x55, 0x45, 0xce, 0xac,
		0x58, 0x9e, 0xaf, 0x6d, 0x91, 0x06, 0xee, 0xd1, 0xa6, 0xb9, 0x2f, 0xa6, 0x2a, 0x72, 0xc6, 0xd7,
		0x3e, 0x04, 0x99, 0x9a, 0xdd, 0xa6, 0xed, 0x03, 0xd7, 0xa3, 0xa9, 0x36, 0xa2, 0xa6, 0xb9, 0xcc,
		0x57, 0x11, 0xfd, 0x5d, 0xe7, 0xa2, 0x20, 0xa3, 0xa6, 0xb9, 0x8c, 0xab, 0xdc, 0x03, 0xe3, 0xb8,
		0xd1, 0x70, 0x28, 0xb9, 0x24, 0xe2, 0x4d, 0x6b, 0xd6, 0x17, 0x33, 0xc5, 0x99, 0x73, 0x90, 0x94,
		0x7e, 0xa0, 0xd5, 0x8c, 0x7a, 0x42, 0x6b, 0xf1, 0xeb, 0x9a, 0xe8, 0x7c, 0x4a, 0x4d, 0x5a, 0x72,
		0xf2, 0x10, 0x64, 0x0c, 0x57, 0xeb, 0xdc, 0x1b, 0x46, 0x67, 0xa3, 0xf3, 0x49, 0x35, 0x6d, 0xb8,
		0xfe, 0x45, 0xd1, 0xdc, 0xdb, 0x51, 0xc8, 0x76, 0xdf, 0x7b, 0xa2, 0x25, 0x48, 0x9a, 0xb6, 0x8e,
		0x59, 0x20, 0xf0, 0x4b, 0xf7, 0xf9, 0x90, 0xab, 0xd2, 0x85, 0x55, 0xa1, 0xaf, 0xfa, 0xc8, 0x99,
		0x9f, 0x47, 0x20, 0x29, 0xc5, 0x68, 0x1a, 0xe2, 0x2d, 0xec, 0x6d, 0x33, 0xba, 0xc4, 0x62, 0x54,
		0x89, 0xa8, 0x6c, 0x4c, 0xe5, 0x6e, 0x0b, 0x5b, 0x2c, 0x04, 0x84, 0x9c, 0x8e, 0xe9, 0xbe, 0x9a,
		0x04, 0xd7, 0x58, 0x3b, 0x6c, 0x37, 0x9b, 0xc4, 0xf2, 0x5c, 0xb9, 0xaf, 0x42, 0x5e, 0x14, 0x62,
		0x74, 0x3f, 0x4c, 0x78, 0x0e, 0x36, 0xcc, 0x2e, 0xdd, 0x38, 0xd3, 0x55, 0xe4, 0x84, 0xaf, 0x9c,
		0x87, 0xfd, 0x92, 0xb7, 0x46, 0x3c, 0xac, 0x6f, 0x93, 0x5a, 0x07, 0x34, 0xc2, 0x2e, 0xd5, 0xf6,
		0x09, 0x85, 0x25, 0x31, 0x2f, 0xb1, 0x73, 0x1f, 0x44, 0x60, 0x42, 0x36, 0xf0, 0x35, 0xdf, 0x59,
		0x65, 0x00, 0x6c, 0x59, 0xb6, 0x17, 0x74, 0x57, 0x7f, 0x28, 0xf7, 0xe1, 0x16, 0x0a, 0x3e, 0x48,
		0x0d, 0x10, 0xcc, 0x34, 0x01, 0x3a, 0x33, 0xbb, 0xba, 0xed, 0x20, 0xa4, 0xc5, 0xa5, 0x36, 0xfb,
		0x32, 0xc2, 0xdf, 0xfa, 0x80, 0x8b, 0x68, 0xa7, 0x8f, 0xa6, 0x20, 0xb1, 0x45, 0x1a, 0x86, 0x25,
		0xae, 0xda, 0xf8, 0x40, 0x5e, 0xe0, 0xc5, 0xfd, 0x0b, 0xbc, 0xc5, 0xa7, 0x61, 0x52, 0xb7, 0x9b,
		0xbd, 0xe6, 0x2e, 0x2a, 0x3d, 0x6f, 0x9e, 0xee, 0x63, 0x91, 0xa7, 0xa0, 0xd3, 0x9d, 0xbd, 0x19,
		0x89, 0xbc, 0x15, 0x8d, 0x2d, 0x57, 0x16, 0xdf, 0x89, 0xce, 0x2c, 0x73, 0x68, 0x45, 0xae, 0x54,
		0x25, 0x75, 0x93, 0xe8, 0xd4, 0x7a, 0x78, 0xe3, 0x30, 0x3c, 0xd0, 0x30, 0xbc, 0xed, 0xf6, 0xd6,
		0x82, 0x6e, 0x37, 0x8f, 0x34, 0xec, 0x86, 0xdd, 0xf9, 0x18, 0x44, 0x47, 0x6c, 0xc0, 0x7e, 0x89,
		0x0f, 0x42, 0x29, 0x5f, 0x3a, 0x13, 0xfa, 0xf5, 0x28, 0xbf, 0x06, 0x93, 0x42, 0x59, 0x63, 0x37,
		0xd2, 0xbc, 0x0f, 0x47, 0x37, 0xbd, 0x95, 0xc8, 0xbd, 0xf7, 0x11, 0xab, 0x74, 0xea, 0x84, 0x80,
		0xd2, 0x39, 0xde, 0xa9, 0xe7, 0x55, 0xb8, 0xb5, 0x8b, 0x8f, 0x1f, 0x4d, 0xe2, 0x84, 0x30, 0xbe,
		0x2f, 0x18, 0x27, 0x03, 0x8c, 0x55, 0x01, 0xcd, 0x17, 0x61, 0x6c, 0x2f, 0x5c, 0x3f, 0x15, 0x5c,
		0x19, 0x12, 0x24, 0x59, 0x86, 0x71, 0x46, 0xa2, 0xb7, 0x5d, 0xcf, 0x6e, 0xb2, 0xbc, 0x77, 0x73,
		0x9a, 0x9f, 0x7d, 0xc4, 0xcf, 0x4a, 0x96, 0xc2, 0x8a, 0x3e, 0x2a, 0xff, 0x38, 0x4c, 0x51, 0x09,
		0x4b, 0x2d, 0x41, 0xb6, 0xf0, 0x7b, 0x94, 0xdc, 0x07, 0xcf, 0xf2, 0x23, 0x35, 0xe9, 0x13, 0x04,
		0x78, 0x03, 0x3b, 0xd1, 0x20, 0x9e, 0x47, 0x1c, 0x57, 0xc3, 0xa6, 0x89, 0x6e, 0xfa, 0x85, 0x26,
		0xf7, 0xd2, 0xc7, 0xdd, 0x3b, 0xb1, 0xcc, 0x91, 0x05, 0xd3, 0xcc, 0x6f, 0xc2, 0xbe, 0x01, 0x3b,
		0x3b, 0x04, 0xe7, 0xcb, 0x82, 0x73, 0xaa, 0x6f, 0x77, 0x29, 0x6d, 0x05, 0xa4, 0xdc, 0xdf, 0x8f,
		0x21, 0x38, 0x5f, 0x11, 0x9c, 0x48, 0x60, 0xe5, 0xb6, 0x50, 0xc6, 0x73, 0x30, 0x71, 0x89, 0x38,
		0x5b, 0xb6, 0x2b, 0x5e, 0xfe, 0x87, 0xa0, 0x7b, 0x55, 0xd0, 0x8d, 0x0b, 0x20, 0xbb, 0x0a, 0xa0,
		0x5c, 0xa7, 0x20, 0x59, 0xc7, 0x3a, 0x19, 0x82, 0xe2, 0x35, 0x41, 0x31, 0x4a, 0xf5, 0x29, 0xb4,
		0x00, 0x99, 0x86, 0x2d, 0xaa, 0x4b, 0x38, 0xfc, 0x75, 0x01, 0x4f, 0x4b, 0x8c, 0xa0, 0x68, 0xd9,
		0xad, 0xb6, 0x49, 0x4b, 0x4f, 0x38, 0xc5, 0x1b, 0x92, 0x42, 0x62, 0x04, 0xc5, 0x1e, 0xdc, 0xfa,
		0xa6, 0xa4, 0x70, 0x03, 0xfe, 0x3c, 0x03, 0x69, 0xdb, 0x32, 0x77, 0x6c, 0x6b, 0x18, 0x23, 0xbe,
		0x22, 0x18, 0x40, 0x40, 0x28, 0xc1, 0x69, 0x48, 0x0d, 0xbb, 0x11, 0x5f, 0x15, 0xf0, 0x24, 0x91,
		0x3b, 0xb0, 0x0c, 0xe3, 0x32, 0xc9, 0x18, 0xb6, 0x35, 0x04, 0xc5, 0xd7, 0x04, 0x45, 0x36, 0x00,
		0x13, 0xcb, 0xf0, 0x88, 0xeb, 0x35, 0xc8, 0x30, 0x24, 0x6f, 0xcb, 0x65, 0x08, 0x88, 0x70, 0xe5,
		0x16, 0xb1, 0xf4, 0xed, 0xe1, 0x18, 0xbe, 0x2e, 0x5d, 0x29, 0x31, 0x94, 0xa2, 0x08, 0x63, 0x4d,
		0xec, 0xb8, 0xdb, 0xd8, 0x1c, 0x6a, 0x3b, 0xbe, 0x21, 0x38, 0x32, 0x3e, 0x48, 0x78, 0xa4, 0x6d,
		0xed, 0x85, 0xe6, 0x1d, 0xe9, 0x91, 0x00, 0x4c, 0x1c, 0x3d, 0xd7, 0x63, 0xf7, 0x2b, 0x7b, 0x61,
		0xfb, 0xa6, 0x3c, 0x7a, 0x1c, 0x5b, 0x0e, 0x32, 0x9e, 0x86, 0x94, 0x6b, 0x5c, 0x1d, 0x8a, 0xe6,
		0x5b, 0x72, 0xa7, 0x19, 0x80, 0x82, 0x9f, 0x84, 0xfd, 0x03, 0x53, 0xfd, 0x10, 0x64, 0xdf, 0x16,
		0x64, 0xd3, 0x03, 0xd2, 0xbd, 0x48, 0x09, 0x7b, 0xa5, 0xfc, 0x8e, 0x4c, 0x09, 0xa4, 0x87, 0xab,
		0x42, 0xbb, 0x73, 0x17, 0xd7, 0xf7, 0xe6, 0xb5, 0xef, 0x4a, 0xaf, 0x71, 0x6c, 0x97, 0xd7, 0x36,
		0x60, 0x5a, 0x30, 0xee, 0x6d, 0x5f, 0xdf, 0x95, 0x89, 0x95, 0xa3, 0x37, 0xbb, 0x77, 0xf7, 0x69,
		0x98, 0xf1, 0xdd, 0x29, 0x1b, 0x4b, 0x57, 0x6b, 0xe2, 0xd6, 0x10, 0xcc, 0xef, 0x09, 0x66, 0x99,
		0xf1, 0xfd, 0xce, 0xd4, 0x2d, 0xe3, 0x16, 0x25, 0x3f, 0x0f, 0x39, 0x49, 0xde, 0xb6, 0x1c, 0xa2,
		0xdb, 0x0d, 0xcb, 0xb8, 0x4a, 0x6a, 0x43, 0x50, 0x7f, 0xaf, 0x67, 0xab, 0x36, 0x03, 0x70, 0xca,
		0xbc, 0x02, 0x8a, 0xdf, 0x6f, 0x68, 0x46, 0xb3, 0x65, 0x3b, 0x5e, 0x08, 0xe3, 0xf7, 0xe5, 0x4e,
		0xf9, 0xb8, 0x15, 0x06, 0xcb, 0x97, 0x20, 0xcb, 0x86, 0xc3, 0x86, 0xe4, 0x0f, 0x04, 0xd1, 0x58,
		0x07, 0x25, 0x12, 0x87, 0x6e, 0x37, 0x5b, 0xd8, 0x19, 0x26, 0xff, 0xfd, 0x50, 0x26, 0x0e, 0x01,
		0xe1, 0xd1, 0x37, 0xde, 0x53, 0x89, 0x51, 0xd8, 0xc7, 0xeb, 0xdc, 0x33, 0xd7, 0xc5, 0x99, 0xed,
		0x2e, 0xc4, 0xf9, 0x55, 0xea, 0x9e, 0xee, 0x72, 0x19, 0x4e, 0xf6, 0xec, 0x75, 0xdf, 0x43, 0x5d,
		0xd5, 0x32, 0x7f, 0x16, 0xc6, 0xba, 0x4a, 0x65, 0x38, 0xd5, 0xbf, 0x09, 0xaa, 0x4c, 0xb0, 0x52,
		0xe6, 0x8f, 0x43, 0x9c, 0x96, 0xbd, 0x70, 0xf8, 0xbf, 0x0b, 0x38, 0x53, 0xcf, 0x3f, 0x02, 0x49,
		0x59, 0xee, 0xc2, 0xa1, 0xff, 0x21, 0xa0, 0x3e, 0x84, 0xc2, 0x65, 0xa9, 0x0b, 0x87, 0xff, 0xa7,
		0x84, 0x4b, 0x08, 0x85, 0x0f, 0xef, 0xc2, 0x9f, 0x3c, 0x1f, 0x17, 0xe9, 0x4a, 0xfa, 0xee, 0x34,
		0x8c, 0x8a, 0x1a, 0x17, 0x8e, 0x7e, 0x4e, 0x3c, 0x5c, 0x22, 0xf2, 0x0f, 0x43, 0x62, 0x48, 0x87,
		0xff, 0x97, 0x80, 0x72, 0xfd, 0x7c, 0x11, 0xd2, 0x81, 0xba, 0x16, 0x0e, 0xff, 0x6f, 0x01, 0x0f,
		0xa2, 0xa8, 0xe9, 0xa2, 0xae, 0x85, 0x13, 0xfc, 0x8f, 0x34, 0x5d, 0x20, 0xa8, 0xdb, 0x64, 0x49,
		0x0b, 0x47, 0xbf, 0x20, 0xbd, 0x2e, 0x21, 0xf9, 0x33, 0x90, 0xf2, 0xd3, 0x54, 0x38, 0xfe, 0x7f,
		0x05, 0xbe, 0x83, 0xa1, 0x1e, 0x08, 0xa4, 0xc9, 0x70, 0x8a, 0xff, 0x93, 0x1e, 0x08, 0xa0, 0xe8,
		0x31, 0xea, 0x2d, 0x7d, 0xe1, 0x4c, 0xff, 0x2f, 0x8f, 0x51, 0x4f, 0xe5, 0xa3, 0xbb, 0xc9, 0xb2,
		0x45, 0x38, 0xc5, 0x97, 0xe4, 0x6e, 0x32, 0x7d, 0x6a, 0x46, 0x6f, 0x2d, 0x09, 0xe7, 0xf8, 0xb2,
		0x34, 0xa3, 0xa7, 0x94, 0xe4, 0x2b, 0x80, 0xfa, 0xeb, 0x48, 0x38, 0xdf, 0x8b, 0x82, 0x6f, 0xa2,
		0xaf, 0x8c, 0xe4, 0x9f, 0x80, 0xe9, 0xc1, 0x35, 0x24, 0x9c, 0xf5, 0xa5, 0xeb, 0x3d, 0x5d, 0x7f,
		0xb0, 0x84, 0xe4, 0x37, 0x3a, 0x5d, 0x7f, 0xb0, 0x7e, 0x84, 0xd3, 0xbe, 0x7c, 0xbd, 0xfb, 0xc5,
		0x2e, 0x58, 0x3e, 0xf2, 0x05, 0x80, 0x4e, 0xea, 0x0e, 0xe7, 0x7a, 0x55, 0x70, 0x05, 0x40, 0xf4,
		0x68, 0x88, 0xcc, 0x1d, 0x8e, 0x7f, 0x4d, 0x1e, 0x0d, 0x81, 0xc8, 0x9f, 0x86, 0xa4, 0xd5, 0x36,
		0x4d, 0x1a, 0x1c, 0xe8, 0xe6, 0xff, 0x10, 0x92, 0xfb, 0xcd, 0x0d, 0x71, 0x30, 0x24, 0x20, 0x7f,
		0x1c, 0x12, 0xa4, 0xb9, 0x45, 0x6a, 0x61, 0xc8, 0xdf, 0xde, 0x90, 0x09, 0x81, 0x6a, 0xe7, 0xcf,
		0x00, 0xf0, 0x97, 0x46, 0xf6, 0x3d, 0x20, 0x04, 0xfb, 0xbb, 0x1b, 0xe2, 0x5b, 0x73, 0x07, 0xd2,
		0x21, 0xe0, 0x5f, 0xae, 0x6f, 0x4e, 0xf0, 0x71, 0x37, 0x01, 0x7b, 0xd1, 0x3c, 0x05, 0xa3, 0x17,
		0x5c, 0xdb, 0xf2, 0x70, 0x23, 0x0c, 0xfd, 0x7b, 0x81, 0x96, 0xfa, 0xd4, 0x61, 0x4d, 0xdb, 0x21,
		0x1e, 0x6e, 0xb8, 0x61, 0xd8, 0x3f, 0x08, 0xac, 0x0f, 0xa0, 0x60, 0x1d, 0xbb, 0xde, 0x30, 0xeb,
		0xfe, 0xa3, 0x04, 0x4b, 0x00, 0x35, 0x9a, 0xfe, 0xbe, 0x48, 0x76, 0xc2, 0xb0, 0x9f, 0x48, 0xa3,
		0x85, 0x7e, 0xfe, 0x11, 0x48, 0xd1, 0x9f, 0xfc, 0xff, 0x2f, 0x42, 0xc0, 0x7f, 0x12, 0xe0, 0x0e,
		0x82, 0x3e, 0xd9, 0xf5, 0x6a, 0x9e, 0x11, 0xee, 0xec, 0x3f, 0x8b, 0x9d, 0x96, 0xfa, 0xf9, 0x02,
		0xa4, 0x5d, 0xaf, 0x56, 0x6b, 0x3b, 0xfc, 0x22, 0x2a, 0x04, 0xfe, 0x97, 0x1b, 0xfe, 0xcb, 0x9c,
		0x8f, 0x59, 0x3c, 0x34, 0xf8, 0x6e, 0x09, 0x96, 0xed, 0x65, 0x9b, 0xdf, 0x2a, 0xc1, 0x0b, 0x51,
		0x48, 0x35, 0xdd, 0x86, 0xb8, 0xfd, 0xe1, 0x87, 0x83, 0x26, 0x7f, 0x77, 0x66, 0x6f, 0x17, 0x47,
		0x73, 0x67, 0x61, 0x64, 0xd1, 0x68, 0x94, 0xdd, 0x06, 0x9a, 0x82, 0x04, 0x33, 0x8d, 0x7d, 0x30,
		0x88, 0xa9, 0x7c, 0x80, 0xee, 0x86, 0x58, 0xd9, 0x6d, 0x88, 0x7f, 0xc5, 0x98, 0x5a, 0xe8, 0x3c,
		0x68, 0xa1, 0xda, 0xc4, 0xa6, 0x59, 0x76, 0x1b, 0x2a, 0x55, 0x98, 0x73, 0x20, 0x29, 0x05, 0x68,
		0x16, 0xd2, 0x55, 0x1d, 0x3b, 0x8b, 0x6d, 0xb7, 0xea, 0xd9, 0x2d, 0xf9, 0xaf, 0x06, 0x01, 0x11,
		0x9a, 0x87, 0xf1, 0xb3, 0xa6, 0xd1, 0xd8, 0xf6, 0x2a, 0xd8, 0xc1, 0xfa, 0x76, 0xdb, 0x23, 0xb9,
		0xcc, 0x6c, 0x6c, 0x7e, 0x54, 0xed, 0x15, 0xa3, 0x19, 0x48, 0x96, 0x71, 0xab, 0xba, 0x8d, 0x9d,
		0x8b, 0xec, 0xab, 0x75, 0x4a, 0xf5, 0xc7, 0x8b, 0x99, 0x4f, 0x3e, 0x3c, 0x10, 0xf9, 0xdb, 0x87,
		0x07, 0x22, 0xef, 0xfe, 0xf2, 0x40, 0xe4, 0xef, 0x01, 0x00, 0x00, 0xff, 0xff, 0x03, 0x17, 0xc9,
		0xf0, 0x39, 0x2d, 0x00, 0x00,
	}
	r := bytes.NewReader(gzipped)
	gzipr, err := compress_gzip.NewReader(r)
	if err != nil {
		panic(err)
	}
	ungzipped, err := io_ioutil.ReadAll(gzipr)
	if err != nil {
		panic(err)
	}
	if err := github_com_gogo_protobuf_proto.Unmarshal(ungzipped, d); err != nil {
		panic(err)
	}
	return d
}
func (this *BigMsg) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&prototests.BigMsg{")
	if this.Field != nil {
		s = append(s, "Field: "+valueToGoStringMsg(this.Field, "int64")+",\n")
	}
	if this.Msg != nil {
		s = append(s, "Msg: "+fmt.Sprintf("%#v", this.Msg)+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *SmallMsg) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&prototests.SmallMsg{")
	if this.ScarBusStop != nil {
		s = append(s, "ScarBusStop: "+valueToGoStringMsg(this.ScarBusStop, "string")+",\n")
	}
	if this.FlightParachute != nil {
		s = append(s, "FlightParachute: "+fmt.Sprintf("%#v", this.FlightParachute)+",\n")
	}
	if this.MapShark != nil {
		s = append(s, "MapShark: "+valueToGoStringMsg(this.MapShark, "string")+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringMsg(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func extensionToGoStringMsg(m github_com_gogo_protobuf_proto.Message) string {
	e := github_com_gogo_protobuf_proto.GetUnsafeExtensionsMap(m)
	if e == nil {
		return "nil"
	}
	s := "proto.NewUnsafeXXX_InternalExtensions(map[int32]proto.Extension{"
	keys := make([]int, 0, len(e))
	for k := range e {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	ss := []string{}
	for _, k := range keys {
		ss = append(ss, strconv.Itoa(k)+": "+e[int32(k)].GoString())
	}
	s += strings.Join(ss, ",") + "})"
	return s
}
func NewPopulatedBigMsg(r randyMsg, easy bool) *BigMsg {
	this := &BigMsg{}
	if r.Intn(10) != 0 {
		v1 := int64(r.Int63())
		if r.Intn(2) == 0 {
			v1 *= -1
		}
		this.Field = &v1
	}
	if r.Intn(10) != 0 {
		this.Msg = NewPopulatedSmallMsg(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedMsg(r, 4)
	}
	return this
}

func NewPopulatedSmallMsg(r randyMsg, easy bool) *SmallMsg {
	this := &SmallMsg{}
	if r.Intn(10) != 0 {
		v2 := randStringMsg(r)
		this.ScarBusStop = &v2
	}
	if r.Intn(10) != 0 {
		v3 := r.Intn(10)
		this.FlightParachute = make([]uint32, v3)
		for i := 0; i < v3; i++ {
			this.FlightParachute[i] = uint32(r.Uint32())
		}
	}
	if r.Intn(10) != 0 {
		v4 := randStringMsg(r)
		this.MapShark = &v4
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedMsg(r, 19)
	}
	return this
}

type randyMsg interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneMsg(r randyMsg) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringMsg(r randyMsg) string {
	v5 := r.Intn(100)
	tmps := make([]rune, v5)
	for i := 0; i < v5; i++ {
		tmps[i] = randUTF8RuneMsg(r)
	}
	return string(tmps)
}
func randUnrecognizedMsg(r randyMsg, maxFieldNumber int) (data []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		data = randFieldMsg(data, r, fieldNumber, wire)
	}
	return data
}
func randFieldMsg(data []byte, r randyMsg, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		data = encodeVarintPopulateMsg(data, uint64(key))
		v6 := r.Int63()
		if r.Intn(2) == 0 {
			v6 *= -1
		}
		data = encodeVarintPopulateMsg(data, uint64(v6))
	case 1:
		data = encodeVarintPopulateMsg(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		data = encodeVarintPopulateMsg(data, uint64(key))
		ll := r.Intn(100)
		data = encodeVarintPopulateMsg(data, uint64(ll))
		for j := 0; j < ll; j++ {
			data = append(data, byte(r.Intn(256)))
		}
	default:
		data = encodeVarintPopulateMsg(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return data
}
func encodeVarintPopulateMsg(data []byte, v uint64) []byte {
	for v >= 1<<7 {
		data = append(data, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	data = append(data, uint8(v))
	return data
}

func init() { proto.RegisterFile("msg.proto", fileDescriptorMsg) }

var fileDescriptorMsg = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xcc, 0x2d, 0x4e, 0xd7,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x02, 0x53, 0x25, 0xa9, 0xc5, 0x25, 0xc5, 0x52, 0xba,
	0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0xe9, 0xf9, 0xe9, 0xf9, 0xfa,
	0x60, 0xb9, 0xa4, 0xd2, 0x34, 0x30, 0x0f, 0xcc, 0x01, 0xb3, 0x20, 0x5a, 0x95, 0xdc, 0xb8, 0xd8,
	0x9c, 0x32, 0xd3, 0x7d, 0x8b, 0xd3, 0x85, 0x44, 0xb8, 0x58, 0xdd, 0x32, 0x53, 0x73, 0x52, 0x24,
	0x18, 0x15, 0x18, 0x35, 0x98, 0x83, 0x20, 0x1c, 0x21, 0x35, 0x2e, 0x66, 0xdf, 0xe2, 0x74, 0x09,
	0x66, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x11, 0x3d, 0x84, 0x45, 0x7a, 0xc1, 0xb9, 0x89, 0x39, 0x39,
	0xbe, 0xc5, 0xe9, 0x41, 0x20, 0x05, 0x4a, 0x45, 0x5c, 0x1c, 0x30, 0x01, 0x21, 0x05, 0x2e, 0xee,
	0xe0, 0xe4, 0xc4, 0x22, 0xa7, 0xd2, 0xe2, 0xe0, 0x92, 0xfc, 0x02, 0xb0, 0x79, 0x9c, 0x41, 0xc8,
	0x42, 0x42, 0x1a, 0x5c, 0xfc, 0x6e, 0x39, 0x99, 0xe9, 0x19, 0x25, 0x01, 0x89, 0x45, 0x89, 0xc9,
	0x19, 0xa5, 0x25, 0xa9, 0x12, 0x3c, 0x0a, 0xcc, 0x1a, 0xec, 0x41, 0xe8, 0xc2, 0x42, 0x52, 0x5c,
	0x1c, 0xbe, 0x89, 0x05, 0xc1, 0x19, 0x89, 0x45, 0xd9, 0x12, 0x42, 0x60, 0x83, 0xe0, 0x7c, 0x27,
	0x9e, 0x0f, 0x0f, 0xe5, 0x18, 0x7f, 0x3c, 0x94, 0x63, 0xdc, 0xf0, 0x48, 0x8e, 0x11, 0x10, 0x00,
	0x00, 0xff, 0xff, 0x88, 0x10, 0x4a, 0x47, 0x10, 0x01, 0x00, 0x00,
}
