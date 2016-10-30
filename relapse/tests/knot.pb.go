// Code generated by protoc-gen-gogo.
// source: knot.proto
// DO NOT EDIT!

/*
Package tests is a generated protocol buffer package.

It is generated from these files:
	knot.proto

It has these top-level messages:
	Knot
	BightKnot
*/
package tests

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

type Knot struct {
	Bight            []*BightKnot `protobuf:"bytes,1,rep,name=Bight" json:"Bight,omitempty"`
	Elbow            *bool        `protobuf:"varint,2,opt,name=Elbow" json:"Elbow,omitempty"`
	BitterEnd        *Knot        `protobuf:"bytes,3,opt,name=BitterEnd" json:"BitterEnd,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Knot) Reset()                    { *m = Knot{} }
func (m *Knot) String() string            { return proto.CompactTextString(m) }
func (*Knot) ProtoMessage()               {}
func (*Knot) Descriptor() ([]byte, []int) { return fileDescriptorKnot, []int{0} }

func (m *Knot) GetBight() []*BightKnot {
	if m != nil {
		return m.Bight
	}
	return nil
}

func (m *Knot) GetElbow() bool {
	if m != nil && m.Elbow != nil {
		return *m.Elbow
	}
	return false
}

func (m *Knot) GetBitterEnd() *Knot {
	if m != nil {
		return m.BitterEnd
	}
	return nil
}

type BightKnot struct {
	Loop             *Knot  `protobuf:"bytes,1,opt,name=Loop" json:"Loop,omitempty"`
	Turn             *bool  `protobuf:"varint,2,opt,name=Turn" json:"Turn,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *BightKnot) Reset()                    { *m = BightKnot{} }
func (m *BightKnot) String() string            { return proto.CompactTextString(m) }
func (*BightKnot) ProtoMessage()               {}
func (*BightKnot) Descriptor() ([]byte, []int) { return fileDescriptorKnot, []int{1} }

func (m *BightKnot) GetLoop() *Knot {
	if m != nil {
		return m.Loop
	}
	return nil
}

func (m *BightKnot) GetTurn() bool {
	if m != nil && m.Turn != nil {
		return *m.Turn
	}
	return false
}

func init() {
	proto.RegisterType((*Knot)(nil), "tests.Knot")
	proto.RegisterType((*BightKnot)(nil), "tests.BightKnot")
}
func (this *Knot) Description() (desc *github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet) {
	return KnotDescription()
}
func (this *BightKnot) Description() (desc *github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet) {
	return KnotDescription()
}
func KnotDescription() (desc *github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet) {
	d := &github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet{}
	var gzipped = []byte{
		// 3449 bytes of a gzipped FileDescriptorSet
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xcc, 0x5a, 0x4b, 0x70, 0x1c, 0xd5,
		0xd5, 0x66, 0x5e, 0xd2, 0xcc, 0x99, 0xd1, 0xa8, 0x75, 0x25, 0xe4, 0xb1, 0x00, 0x5b, 0x16, 0x06,
		0x64, 0xf8, 0x91, 0x29, 0x63, 0x1b, 0x7b, 0xfc, 0x83, 0x33, 0x92, 0xc6, 0x42, 0x8e, 0x5e, 0xe9,
		0x91, 0xc0, 0xc0, 0xa2, 0xeb, 0xaa, 0xfb, 0xce, 0xa8, 0xed, 0x9e, 0xee, 0x49, 0x77, 0x8f, 0x6d,
		0x79, 0x65, 0x8a, 0x3c, 0x8a, 0xa2, 0xf2, 0x4e, 0x55, 0x78, 0x43, 0xa8, 0x4a, 0x48, 0xc8, 0x0b,
		0xf2, 0x5a, 0x64, 0x95, 0x0d, 0xc9, 0x2e, 0x55, 0xec, 0xb3, 0x49, 0x8a, 0xaa, 0xbc, 0x48, 0x42,
		0x12, 0x57, 0x65, 0xe1, 0x45, 0x52, 0xf7, 0xd5, 0xd3, 0xf3, 0x90, 0x7b, 0x44, 0x15, 0x90, 0x95,
		0xe6, 0x9e, 0x7b, 0xbe, 0xaf, 0xef, 0x3d, 0xe7, 0xdc, 0x73, 0x4e, 0xdf, 0x16, 0x3c, 0x79, 0x14,
		0x26, 0x6b, 0x8e, 0x53, 0xb3, 0xc8, 0xe1, 0x86, 0xeb, 0xf8, 0xce, 0x66, 0xb3, 0x7a, 0xd8, 0x20,
		0x9e, 0xee, 0x9a, 0x0d, 0xdf, 0x71, 0x67, 0x98, 0x0c, 0x0d, 0x73, 0x8d, 0x19, 0xa9, 0x31, 0xb5,
		0x0c, 0x23, 0x67, 0x4c, 0x8b, 0xcc, 0x07, 0x8a, 0x15, 0xe2, 0xa3, 0x13, 0x90, 0xac, 0x9a, 0x16,
		0x29, 0xc4, 0x26, 0x13, 0xd3, 0xd9, 0x23, 0x07, 0x67, 0x3a, 0x40, 0x33, 0xed, 0x88, 0x35, 0x2a,
		0x56, 0x19, 0x62, 0xea, 0xdd, 0x24, 0x8c, 0xf6, 0x98, 0x45, 0x08, 0x92, 0x36, 0xae, 0x53, 0xc6,
		0xd8, 0x74, 0x46, 0x65, 0xbf, 0x51, 0x01, 0x06, 0x1b, 0x58, 0xbf, 0x80, 0x6b, 0xa4, 0x10, 0x67,
		0x62, 0x39, 0x44, 0xfb, 0x00, 0x0c, 0xd2, 0x20, 0xb6, 0x41, 0x6c, 0x7d, 0xbb, 0x90, 0x98, 0x4c,
		0x4c, 0x67, 0xd4, 0x90, 0x04, 0xdd, 0x03, 0x23, 0x8d, 0xe6, 0xa6, 0x65, 0xea, 0x5a, 0x48, 0x0d,
		0x26, 0x13, 0xd3, 0x29, 0x55, 0xe1, 0x13, 0xf3, 0x2d, 0xe5, 0xbb, 0x60, 0xf8, 0x12, 0xc1, 0x17,
		0xc2, 0xaa, 0x59, 0xa6, 0x9a, 0xa7, 0xe2, 0x90, 0xe2, 0x1c, 0xe4, 0xea, 0xc4, 0xf3, 0x70, 0x8d,
		0x68, 0xfe, 0x76, 0x83, 0x14, 0x92, 0x6c, 0xf7, 0x93, 0x5d, 0xbb, 0xef, 0xdc, 0x79, 0x56, 0xa0,
		0xd6, 0xb7, 0x1b, 0x04, 0x95, 0x20, 0x43, 0xec, 0x66, 0x9d, 0x33, 0xa4, 0x76, 0xb0, 0x5f, 0xd9,
		0x6e, 0xd6, 0x3b, 0x59, 0xd2, 0x14, 0x26, 0x28, 0x06, 0x3d, 0xe2, 0x5e, 0x34, 0x75, 0x52, 0x18,
		0x60, 0x04, 0x77, 0x75, 0x11, 0x54, 0xf8, 0x7c, 0x27, 0x87, 0xc4, 0xa1, 0x39, 0xc8, 0x90, 0xcb,
		0x3e, 0xb1, 0x3d, 0xd3, 0xb1, 0x0b, 0x83, 0x8c, 0xe4, 0x8e, 0x1e, 0x5e, 0x24, 0x96, 0xd1, 0x49,
		0xd1, 0xc2, 0xa1, 0xe3, 0x30, 0xe8, 0x34, 0x7c, 0xd3, 0xb1, 0xbd, 0x42, 0x7a, 0x32, 0x36, 0x9d,
		0x3d, 0x72, 0x6b, 0xcf, 0x40, 0x58, 0xe5, 0x3a, 0xaa, 0x54, 0x46, 0x8b, 0xa0, 0x78, 0x4e, 0xd3,
		0xd5, 0x89, 0xa6, 0x3b, 0x06, 0xd1, 0x4c, 0xbb, 0xea, 0x14, 0x32, 0x8c, 0x60, 0x7f, 0xf7, 0x46,
		0x98, 0xe2, 0x9c, 0x63, 0x90, 0x45, 0xbb, 0xea, 0xa8, 0x79, 0xaf, 0x6d, 0x8c, 0xc6, 0x61, 0xc0,
		0xdb, 0xb6, 0x7d, 0x7c, 0xb9, 0x90, 0x63, 0x11, 0x22, 0x46, 0x53, 0xff, 0x4e, 0xc1, 0x70, 0x3f,
		0x21, 0x76, 0x0a, 0x52, 0x55, 0xba, 0xcb, 0x42, 0x7c, 0x37, 0x36, 0xe0, 0x98, 0x76, 0x23, 0x0e,
		0x7c, 0x40, 0x23, 0x96, 0x20, 0x6b, 0x13, 0xcf, 0x27, 0x06, 0x8f, 0x88, 0x44, 0x9f, 0x31, 0x05,
		0x1c, 0xd4, 0x1d, 0x52, 0xc9, 0x0f, 0x14, 0x52, 0xe7, 0x60, 0x38, 0x58, 0x92, 0xe6, 0x62, 0xbb,
		0x26, 0x63, 0xf3, 0x70, 0xd4, 0x4a, 0x66, 0xca, 0x12, 0xa7, 0x52, 0x98, 0x9a, 0x27, 0x6d, 0x63,
		0x34, 0x0f, 0xe0, 0xd8, 0xc4, 0xa9, 0x6a, 0x06, 0xd1, 0xad, 0x42, 0x7a, 0x07, 0x2b, 0xad, 0x52,
		0x95, 0x2e, 0x2b, 0x39, 0x5c, 0xaa, 0x5b, 0xe8, 0x64, 0x2b, 0xd4, 0x06, 0x77, 0x88, 0x94, 0x65,
		0x7e, 0xc8, 0xba, 0xa2, 0x6d, 0x03, 0xf2, 0x2e, 0xa1, 0x71, 0x4f, 0x0c, 0xb1, 0xb3, 0x0c, 0x5b,
		0xc4, 0x4c, 0xe4, 0xce, 0x54, 0x01, 0xe3, 0x1b, 0x1b, 0x72, 0xc3, 0x43, 0x74, 0x3b, 0x04, 0x02,
		0x8d, 0x85, 0x15, 0xb0, 0x2c, 0x94, 0x93, 0xc2, 0x15, 0x5c, 0x27, 0x13, 0x27, 0x20, 0xdf, 0x6e,
		0x1e, 0x34, 0x06, 0x29, 0xcf, 0xc7, 0xae, 0xcf, 0xa2, 0x30, 0xa5, 0xf2, 0x01, 0x52, 0x20, 0x41,
		0x6c, 0x83, 0x65, 0xb9, 0x94, 0x4a, 0x7f, 0x4e, 0x3c, 0x00, 0x43, 0x6d, 0x8f, 0xef, 0x17, 0x38,
		0xf5, 0xec, 0x00, 0x8c, 0xf5, 0x8a, 0xb9, 0x9e, 0xe1, 0x3f, 0x0e, 0x03, 0x76, 0xb3, 0xbe, 0x49,
		0xdc, 0x42, 0x82, 0x31, 0x88, 0x11, 0x2a, 0x41, 0xca, 0xc2, 0x9b, 0xc4, 0x2a, 0x24, 0x27, 0x63,
		0xd3, 0xf9, 0x23, 0xf7, 0xf4, 0x15, 0xd5, 0x33, 0x4b, 0x14, 0xa2, 0x72, 0x24, 0x7a, 0x08, 0x92,
		0x22, 0xc5, 0x51, 0x86, 0xbb, 0xfb, 0x63, 0xa0, 0xb1, 0xa8, 0x32, 0x1c, 0xba, 0x05, 0x32, 0xf4,
		0x2f, 0xb7, 0xed, 0x00, 0x5b, 0x73, 0x9a, 0x0a, 0xa8, 0x5d, 0xd1, 0x04, 0xa4, 0x59, 0x98, 0x19,
		0x44, 0x96, 0x86, 0x60, 0x4c, 0x1d, 0x63, 0x90, 0x2a, 0x6e, 0x5a, 0xbe, 0x76, 0x11, 0x5b, 0x4d,
		0xc2, 0x02, 0x26, 0xa3, 0xe6, 0x84, 0xf0, 0x11, 0x2a, 0x43, 0xfb, 0x21, 0xcb, 0xa3, 0xd2, 0xb4,
		0x0d, 0x72, 0x99, 0x65, 0x9f, 0x94, 0xca, 0x03, 0x75, 0x91, 0x4a, 0xe8, 0xe3, 0xcf, 0x7b, 0x8e,
		0x2d, 0x5d, 0xcb, 0x1e, 0x41, 0x05, 0xec, 0xf1, 0x0f, 0x74, 0x26, 0xbe, 0xdb, 0x7a, 0x6f, 0xaf,
		0x33, 0x16, 0xa7, 0x7e, 0x1e, 0x87, 0x24, 0x3b, 0x6f, 0xc3, 0x90, 0x5d, 0x7f, 0x6c, 0xad, 0xac,
		0xcd, 0xaf, 0x6e, 0xcc, 0x2e, 0x95, 0x95, 0x18, 0xca, 0x03, 0x30, 0xc1, 0x99, 0xa5, 0xd5, 0xd2,
		0xba, 0x12, 0x0f, 0xc6, 0x8b, 0x2b, 0xeb, 0xc7, 0x8f, 0x2a, 0x89, 0x00, 0xb0, 0xc1, 0x05, 0xc9,
		0xb0, 0xc2, 0xfd, 0x47, 0x94, 0x14, 0x52, 0x20, 0xc7, 0x09, 0x16, 0xcf, 0x95, 0xe7, 0x8f, 0x1f,
		0x55, 0x06, 0xda, 0x25, 0xf7, 0x1f, 0x51, 0x06, 0xd1, 0x10, 0x64, 0x98, 0x64, 0x76, 0x75, 0x75,
		0x49, 0x49, 0x07, 0x9c, 0x95, 0x75, 0x75, 0x71, 0x65, 0x41, 0xc9, 0x04, 0x9c, 0x0b, 0xea, 0xea,
		0xc6, 0x9a, 0x02, 0x01, 0xc3, 0x72, 0xb9, 0x52, 0x29, 0x2d, 0x94, 0x95, 0x6c, 0xa0, 0x31, 0xfb,
		0xd8, 0x7a, 0xb9, 0xa2, 0xe4, 0xda, 0x96, 0x75, 0xff, 0x11, 0x65, 0x28, 0x78, 0x44, 0x79, 0x65,
		0x63, 0x59, 0xc9, 0xa3, 0x11, 0x18, 0xe2, 0x8f, 0x90, 0x8b, 0x18, 0xee, 0x10, 0x1d, 0x3f, 0xaa,
		0x28, 0xad, 0x85, 0x70, 0x96, 0x91, 0x36, 0xc1, 0xf1, 0xa3, 0x0a, 0x9a, 0x9a, 0x83, 0x14, 0x8b,
		0x2e, 0x84, 0x20, 0xbf, 0x54, 0x9a, 0x2d, 0x2f, 0x69, 0xab, 0x6b, 0xeb, 0x8b, 0xab, 0x2b, 0xa5,
		0x25, 0x25, 0xd6, 0x92, 0xa9, 0xe5, 0x4f, 0x6d, 0x2c, 0xaa, 0xe5, 0x79, 0x25, 0x1e, 0x96, 0xad,
		0x95, 0x4b, 0xeb, 0xe5, 0x79, 0x25, 0x31, 0xa5, 0xc3, 0x58, 0xaf, 0x3c, 0xd3, 0xf3, 0x64, 0x84,
		0x5c, 0x1c, 0xdf, 0xc1, 0xc5, 0x8c, 0xab, 0xcb, 0xc5, 0xaf, 0xc5, 0x60, 0xb4, 0x47, 0xae, 0xed,
		0xf9, 0x90, 0xd3, 0x90, 0xe2, 0x21, 0xca, 0xab, 0xcf, 0xa1, 0x9e, 0x49, 0x9b, 0x05, 0x6c, 0x57,
		0x05, 0x62, 0xb8, 0x70, 0x05, 0x4e, 0xec, 0x50, 0x81, 0x29, 0x45, 0xd7, 0x22, 0x9f, 0x8a, 0x41,
		0x61, 0x27, 0xee, 0x88, 0x44, 0x11, 0x6f, 0x4b, 0x14, 0xa7, 0x3a, 0x17, 0x70, 0x60, 0xe7, 0x3d,
		0x74, 0xad, 0xe2, 0xf5, 0x18, 0x8c, 0xf7, 0x6e, 0x54, 0x7a, 0xae, 0xe1, 0x21, 0x18, 0xa8, 0x13,
		0x7f, 0xcb, 0x91, 0xc5, 0xfa, 0xce, 0x1e, 0x25, 0x80, 0x4e, 0x77, 0xda, 0x4a, 0xa0, 0xc2, 0x35,
		0x24, 0xb1, 0x53, 0xb7, 0xc1, 0x57, 0xd3, 0xb5, 0xd2, 0xa7, 0xe3, 0x70, 0x73, 0x4f, 0xf2, 0x9e,
		0x0b, 0xbd, 0x0d, 0xc0, 0xb4, 0x1b, 0x4d, 0x9f, 0x17, 0x64, 0x9e, 0x9f, 0x32, 0x4c, 0xc2, 0xce,
		0x3e, 0xcd, 0x3d, 0x4d, 0x3f, 0x98, 0x4f, 0xb0, 0x79, 0xe0, 0x22, 0xa6, 0x70, 0xa2, 0xb5, 0xd0,
		0x24, 0x5b, 0xe8, 0xbe, 0x1d, 0x76, 0xda, 0x55, 0xeb, 0xee, 0x03, 0x45, 0xb7, 0x4c, 0x62, 0xfb,
		0x9a, 0xe7, 0xbb, 0x04, 0xd7, 0x4d, 0xbb, 0xc6, 0x12, 0x70, 0xba, 0x98, 0xaa, 0x62, 0xcb, 0x23,
		0xea, 0x30, 0x9f, 0xae, 0xc8, 0x59, 0x8a, 0x60, 0x55, 0xc6, 0x0d, 0x21, 0x06, 0xda, 0x10, 0x7c,
		0x3a, 0x40, 0x4c, 0x3d, 0x33, 0x08, 0xd9, 0x50, 0x5b, 0x87, 0x0e, 0x40, 0xee, 0x3c, 0xbe, 0x88,
		0x35, 0xd9, 0xaa, 0x73, 0x4b, 0x64, 0xa9, 0x6c, 0x4d, 0xb4, 0xeb, 0xf7, 0xc1, 0x18, 0x53, 0x71,
		0x9a, 0x3e, 0x71, 0x35, 0xdd, 0xc2, 0x9e, 0xc7, 0x8c, 0x96, 0x66, 0xaa, 0x88, 0xce, 0xad, 0xd2,
		0xa9, 0x39, 0x39, 0x83, 0x8e, 0xc1, 0x28, 0x43, 0xd4, 0x9b, 0x96, 0x6f, 0x36, 0x2c, 0xa2, 0xd1,
		0x97, 0x07, 0x8f, 0x25, 0xe2, 0x60, 0x65, 0x23, 0x54, 0x63, 0x59, 0x28, 0xd0, 0x15, 0x79, 0x68,
		0x01, 0x6e, 0x63, 0xb0, 0x1a, 0xb1, 0x89, 0x8b, 0x7d, 0xa2, 0x91, 0x4f, 0x37, 0xb1, 0xe5, 0x69,
		0xd8, 0x36, 0xb4, 0x2d, 0xec, 0x6d, 0x15, 0xc6, 0xc2, 0x04, 0x7b, 0xa9, 0xee, 0x82, 0x50, 0x2d,
		0x33, 0xcd, 0x92, 0x6d, 0x3c, 0x8c, 0xbd, 0x2d, 0x54, 0x84, 0x71, 0x46, 0xe4, 0xf9, 0xae, 0x69,
		0xd7, 0x34, 0x7d, 0x8b, 0xe8, 0x17, 0xb4, 0xa6, 0x5f, 0x3d, 0x51, 0xb8, 0x25, 0xcc, 0xc0, 0x16,
		0x59, 0x61, 0x3a, 0x73, 0x54, 0x65, 0xc3, 0xaf, 0x9e, 0x40, 0x15, 0xc8, 0x51, 0x7f, 0xd4, 0xcd,
		0x2b, 0x44, 0xab, 0x3a, 0x2e, 0x2b, 0x2e, 0xf9, 0x1e, 0x87, 0x3b, 0x64, 0xc4, 0x99, 0x55, 0x01,
		0x58, 0x76, 0x0c, 0x52, 0x4c, 0x55, 0xd6, 0xca, 0xe5, 0x79, 0x35, 0x2b, 0x59, 0xce, 0x38, 0x2e,
		0x8d, 0xa9, 0x9a, 0x13, 0xd8, 0x38, 0xcb, 0x63, 0xaa, 0xe6, 0x48, 0x0b, 0x1f, 0x83, 0x51, 0x5d,
		0xe7, 0xdb, 0x36, 0x75, 0x4d, 0x74, 0xf9, 0x5e, 0x41, 0x69, 0xb3, 0x97, 0xae, 0x2f, 0x70, 0x05,
		0x11, 0xe6, 0x1e, 0x3a, 0x09, 0x37, 0xb7, 0xec, 0x15, 0x06, 0x8e, 0x74, 0xed, 0xb2, 0x13, 0x7a,
		0x0c, 0x46, 0x1b, 0xdb, 0xdd, 0x40, 0xd4, 0xf6, 0xc4, 0xc6, 0x76, 0x27, 0xec, 0x0e, 0xf6, 0xe6,
		0xe6, 0x12, 0x1d, 0xfb, 0xc4, 0x28, 0xec, 0x09, 0x6b, 0x87, 0x26, 0xd0, 0x61, 0x50, 0x74, 0x5d,
		0x23, 0x36, 0xde, 0xb4, 0x88, 0x86, 0x5d, 0x62, 0x63, 0xaf, 0xb0, 0x3f, 0xac, 0x9c, 0xd7, 0xf5,
		0x32, 0x9b, 0x2d, 0xb1, 0x49, 0x74, 0x37, 0x8c, 0x38, 0x9b, 0xe7, 0x75, 0x1e, 0x5c, 0x5a, 0xc3,
		0x25, 0x55, 0xf3, 0x72, 0xe1, 0x20, 0x33, 0xd3, 0x30, 0x9d, 0x60, 0xa1, 0xb5, 0xc6, 0xc4, 0xe8,
		0x10, 0x28, 0xba, 0xb7, 0x85, 0xdd, 0x06, 0xab, 0xee, 0x5e, 0x03, 0xeb, 0xa4, 0x70, 0x07, 0x57,
		0xe5, 0xf2, 0x15, 0x29, 0x46, 0xe7, 0x60, 0xac, 0x69, 0x9b, 0xb6, 0x4f, 0xdc, 0x86, 0x4b, 0x68,
		0x93, 0xce, 0x4f, 0x5a, 0xe1, 0x0f, 0x83, 0x3b, 0xb4, 0xd9, 0x1b, 0x61, 0x6d, 0xee, 0x5d, 0x75,
		0xb4, 0xd9, 0x2d, 0x9c, 0x2a, 0x42, 0x2e, 0xec, 0x74, 0x94, 0x01, 0xee, 0x76, 0x25, 0x46, 0x6b,
		0xe8, 0xdc, 0xea, 0x3c, 0xad, 0x7e, 0x8f, 0x97, 0x95, 0x38, 0xad, 0xc2, 0x4b, 0x8b, 0xeb, 0x65,
		0x4d, 0xdd, 0x58, 0x59, 0x5f, 0x5c, 0x2e, 0x2b, 0x89, 0xbb, 0x33, 0xe9, 0x3f, 0x0e, 0x2a, 0x57,
		0xaf, 0x5e, 0xbd, 0x1a, 0x9f, 0x7a, 0x3b, 0x0e, 0xf9, 0xf6, 0xce, 0x17, 0xfd, 0x3f, 0xec, 0x91,
		0xaf, 0xa9, 0x1e, 0xf1, 0xb5, 0x4b, 0xa6, 0xcb, 0xe2, 0xb0, 0x8e, 0x79, 0xef, 0x18, 0x98, 0x70,
		0x4c, 0x68, 0x55, 0x88, 0xff, 0xa8, 0xe9, 0xd2, 0x28, 0xab, 0x63, 0x1f, 0x2d, 0xc1, 0x7e, 0xdb,
		0xd1, 0x3c, 0x1f, 0xdb, 0x06, 0x76, 0x0d, 0xad, 0x75, 0x41, 0xa0, 0x61, 0x5d, 0x27, 0x9e, 0xe7,
		0xf0, 0x12, 0x10, 0xb0, 0xdc, 0x6a, 0x3b, 0x15, 0xa1, 0xdc, 0xca, 0x8d, 0x25, 0xa1, 0xda, 0xe1,
		0xee, 0xc4, 0x4e, 0xee, 0xbe, 0x05, 0x32, 0x75, 0xdc, 0xd0, 0x88, 0xed, 0xbb, 0xdb, 0xac, 0x5f,
		0x4b, 0xab, 0xe9, 0x3a, 0x6e, 0x94, 0xe9, 0xf8, 0xc3, 0xf3, 0x41, 0xd8, 0x8e, 0xbf, 0x4d, 0x40,
		0x2e, 0xdc, 0xb3, 0xd1, 0x16, 0x58, 0x67, 0xf9, 0x39, 0xc6, 0x8e, 0xef, 0xed, 0x37, 0xec, 0xf0,
		0x66, 0xe6, 0x68, 0xe2, 0x2e, 0x0e, 0xf0, 0x4e, 0x4a, 0xe5, 0x48, 0x5a, 0x34, 0xe9, 0x81, 0x25,
		0xbc, 0x3f, 0x4f, 0xab, 0x62, 0x84, 0x16, 0x60, 0xe0, 0xbc, 0xc7, 0xb8, 0x07, 0x18, 0xf7, 0xc1,
		0x1b, 0x73, 0x9f, 0xad, 0x30, 0xf2, 0xcc, 0xd9, 0x8a, 0xb6, 0xb2, 0xaa, 0x2e, 0x97, 0x96, 0x54,
		0x01, 0x47, 0x7b, 0x21, 0x69, 0xe1, 0x2b, 0xdb, 0xed, 0x29, 0x9e, 0x89, 0xfa, 0x35, 0xfc, 0x5e,
		0x48, 0x5e, 0x22, 0xf8, 0x42, 0x7b, 0x62, 0x65, 0xa2, 0x0f, 0x31, 0xf4, 0x0f, 0x43, 0x8a, 0xd9,
		0x0b, 0x01, 0x08, 0x8b, 0x29, 0x37, 0xa1, 0x34, 0x24, 0xe7, 0x56, 0x55, 0x1a, 0xfe, 0x0a, 0xe4,
		0xb8, 0x54, 0x5b, 0x5b, 0x2c, 0xcf, 0x95, 0x95, 0xf8, 0xd4, 0x31, 0x18, 0xe0, 0x46, 0xa0, 0x47,
		0x23, 0x30, 0x83, 0x72, 0x93, 0x18, 0x0a, 0x8e, 0x98, 0x9c, 0xdd, 0x58, 0x9e, 0x2d, 0xab, 0x4a,
		0x3c, 0xec, 0x5e, 0x0f, 0x72, 0xe1, 0x76, 0xed, 0xa3, 0x89, 0xa9, 0x5f, 0xc4, 0x20, 0x1b, 0x6a,
		0xbf, 0x68, 0xe1, 0xc7, 0x96, 0xe5, 0x5c, 0xd2, 0xb0, 0x65, 0x62, 0x4f, 0x04, 0x05, 0x30, 0x51,
		0x89, 0x4a, 0xfa, 0x75, 0xda, 0x47, 0xb2, 0xf8, 0x97, 0x63, 0xa0, 0x74, 0xb6, 0x6e, 0x1d, 0x0b,
		0x8c, 0x7d, 0xac, 0x0b, 0x7c, 0x31, 0x06, 0xf9, 0xf6, 0x7e, 0xad, 0x63, 0x79, 0x07, 0x3e, 0xd6,
		0xe5, 0xbd, 0x10, 0x83, 0xa1, 0xb6, 0x2e, 0xed, 0x7f, 0x6a, 0x75, 0xcf, 0x27, 0x60, 0xb4, 0x07,
		0x0e, 0x95, 0x44, 0x3b, 0xcb, 0x3b, 0xec, 0x7b, 0xfb, 0x79, 0xd6, 0x0c, 0xad, 0x96, 0x6b, 0xd8,
		0xf5, 0x45, 0xf7, 0x7b, 0x08, 0x14, 0xd3, 0x20, 0xb6, 0x6f, 0x56, 0x4d, 0xe2, 0x8a, 0x57, 0x70,
		0xde, 0xe3, 0x0e, 0xb7, 0xe4, 0xfc, 0x2d, 0xfc, 0xff, 0x00, 0x35, 0x1c, 0xcf, 0xf4, 0xcd, 0x8b,
		0x44, 0x33, 0x6d, 0xf9, 0xbe, 0x4e, 0x7b, 0xde, 0xa4, 0xaa, 0xc8, 0x99, 0x45, 0xdb, 0x0f, 0xb4,
		0x6d, 0x52, 0xc3, 0x1d, 0xda, 0x34, 0xf7, 0x25, 0x54, 0x45, 0xce, 0x04, 0xda, 0x07, 0x20, 0x67,
		0x38, 0x4d, 0xda, 0x3e, 0x70, 0x3d, 0x9a, 0x6a, 0x63, 0x6a, 0x96, 0xcb, 0x02, 0x15, 0xd1, 0xdf,
		0xb5, 0x2e, 0x0a, 0x72, 0x6a, 0x96, 0xcb, 0xb8, 0xca, 0x5d, 0x30, 0x8c, 0x6b, 0x35, 0x97, 0x92,
		0x4b, 0x22, 0xde, 0xb4, 0xe6, 0x03, 0x31, 0x53, 0x9c, 0x38, 0x0b, 0x69, 0x69, 0x07, 0x5a, 0xcd,
		0xa8, 0x25, 0xb4, 0x06, 0xbf, 0xae, 0x89, 0x4f, 0x67, 0xd4, 0xb4, 0x2d, 0x27, 0x0f, 0x40, 0xce,
		0xf4, 0xb4, 0xd6, 0xbd, 0x61, 0x7c, 0x32, 0x3e, 0x9d, 0x56, 0xb3, 0xa6, 0x17, 0x5c, 0x14, 0x4d,
		0xbd, 0x1e, 0x87, 0x7c, 0xfb, 0xbd, 0x27, 0x9a, 0x87, 0xb4, 0xe5, 0xe8, 0x98, 0x05, 0x02, 0xbf,
		0x74, 0x9f, 0x8e, 0xb8, 0x2a, 0x9d, 0x59, 0x12, 0xfa, 0x6a, 0x80, 0x9c, 0xf8, 0x4d, 0x0c, 0xd2,
		0x52, 0x8c, 0xc6, 0x21, 0xd9, 0xc0, 0xfe, 0x16, 0xa3, 0x4b, 0xcd, 0xc6, 0x95, 0x98, 0xca, 0xc6,
		0x54, 0xee, 0x35, 0xb0, 0xcd, 0x42, 0x40, 0xc8, 0xe9, 0x98, 0xfa, 0xd5, 0x22, 0xd8, 0x60, 0xed,
		0xb0, 0x53, 0xaf, 0x13, 0xdb, 0xf7, 0xa4, 0x5f, 0x85, 0x7c, 0x4e, 0x88, 0xd1, 0x3d, 0x30, 0xe2,
		0xbb, 0xd8, 0xb4, 0xda, 0x74, 0x93, 0x4c, 0x57, 0x91, 0x13, 0x81, 0x72, 0x11, 0xf6, 0x4a, 0x5e,
		0x83, 0xf8, 0x58, 0xdf, 0x22, 0x46, 0x0b, 0x34, 0xc0, 0x2e, 0xd5, 0xf6, 0x08, 0x85, 0x79, 0x31,
		0x2f, 0xb1, 0x53, 0xef, 0xc4, 0x60, 0x44, 0x36, 0xf0, 0x46, 0x60, 0xac, 0x65, 0x00, 0x6c, 0xdb,
		0x8e, 0x1f, 0x36, 0x57, 0x77, 0x28, 0x77, 0xe1, 0x66, 0x4a, 0x01, 0x48, 0x0d, 0x11, 0x4c, 0xd4,
		0x01, 0x5a, 0x33, 0x3b, 0x9a, 0x6d, 0x3f, 0x64, 0xc5, 0xa5, 0x36, 0xfb, 0x32, 0xc2, 0xdf, 0xfa,
		0x80, 0x8b, 0x68, 0xa7, 0x8f, 0xc6, 0x20, 0xb5, 0x49, 0x6a, 0xa6, 0x2d, 0xae, 0xda, 0xf8, 0x40,
		0x5e, 0xe0, 0x25, 0x83, 0x0b, 0xbc, 0xd9, 0x27, 0x60, 0x54, 0x77, 0xea, 0x9d, 0xcb, 0x9d, 0x55,
		0x3a, 0xde, 0x3c, 0xbd, 0x87, 0x63, 0x8f, 0x43, 0xab, 0x3b, 0x7b, 0x35, 0x16, 0x7b, 0x2d, 0x9e,
		0x58, 0x58, 0x9b, 0x7d, 0x23, 0x3e, 0xb1, 0xc0, 0xa1, 0x6b, 0x72, 0xa7, 0x2a, 0xa9, 0x5a, 0x44,
		0xa7, 0xab, 0x87, 0x57, 0x0e, 0xc2, 0xbd, 0x35, 0xd3, 0xdf, 0x6a, 0x6e, 0xce, 0xe8, 0x4e, 0xfd,
		0x70, 0xcd, 0xa9, 0x39, 0xad, 0x8f, 0x41, 0x74, 0xc4, 0x06, 0xec, 0x97, 0xf8, 0x20, 0x94, 0x09,
		0xa4, 0x13, 0x91, 0x5f, 0x8f, 0x8a, 0x2b, 0x30, 0x2a, 0x94, 0x35, 0x76, 0x23, 0xcd, 0xfb, 0x70,
		0x74, 0xc3, 0x5b, 0x89, 0xc2, 0x5b, 0xef, 0xb2, 0x4a, 0xa7, 0x8e, 0x08, 0x28, 0x9d, 0xe3, 0x9d,
		0x7a, 0x51, 0x85, 0x9b, 0xdb, 0xf8, 0xf8, 0xd1, 0x24, 0x6e, 0x04, 0xe3, 0xdb, 0x82, 0x71, 0x34,
		0xc4, 0x58, 0x11, 0xd0, 0xe2, 0x1c, 0x0c, 0xed, 0x86, 0xeb, 0x57, 0x82, 0x2b, 0x47, 0xc2, 0x24,
		0x0b, 0x30, 0xcc, 0x48, 0xf4, 0xa6, 0xe7, 0x3b, 0x75, 0x96, 0xf7, 0x6e, 0x4c, 0xf3, 0xeb, 0x77,
		0xf9, 0x59, 0xc9, 0x53, 0xd8, 0x5c, 0x80, 0x2a, 0x3e, 0x02, 0x63, 0x54, 0xc2, 0x52, 0x4b, 0x98,
		0x2d, 0xfa, 0x1e, 0xa5, 0xf0, 0xce, 0x53, 0xfc, 0x48, 0x8d, 0x06, 0x04, 0x21, 0xde, 0x90, 0x27,
		0x6a, 0xc4, 0xf7, 0x89, 0xeb, 0x69, 0xd8, 0xb2, 0xd0, 0x0d, 0xbf, 0xd0, 0x14, 0x9e, 0x7b, 0xaf,
		0xdd, 0x13, 0x0b, 0x1c, 0x59, 0xb2, 0xac, 0xe2, 0x06, 0xec, 0xe9, 0xe1, 0xd9, 0x3e, 0x38, 0x9f,
		0x17, 0x9c, 0x63, 0x5d, 0xde, 0xa5, 0xb4, 0x6b, 0x20, 0xe5, 0x81, 0x3f, 0xfa, 0xe0, 0x7c, 0x41,
		0x70, 0x22, 0x81, 0x95, 0x6e, 0xa1, 0x8c, 0x67, 0x61, 0xe4, 0x22, 0x71, 0x37, 0x1d, 0x4f, 0xbc,
		0xfc, 0xf7, 0x41, 0xf7, 0xa2, 0xa0, 0x1b, 0x16, 0x40, 0x76, 0x15, 0x40, 0xb9, 0x4e, 0x42, 0xba,
		0x8a, 0x75, 0xd2, 0x07, 0xc5, 0x4b, 0x82, 0x62, 0x90, 0xea, 0x53, 0x68, 0x09, 0x72, 0x35, 0x47,
		0x54, 0x97, 0x68, 0xf8, 0xcb, 0x02, 0x9e, 0x95, 0x18, 0x41, 0xd1, 0x70, 0x1a, 0x4d, 0x8b, 0x96,
		0x9e, 0x68, 0x8a, 0x57, 0x24, 0x85, 0xc4, 0x08, 0x8a, 0x5d, 0x98, 0xf5, 0x55, 0x49, 0xe1, 0x85,
		0xec, 0x79, 0x1a, 0xb2, 0x8e, 0x6d, 0x6d, 0x3b, 0x76, 0x3f, 0x8b, 0xf8, 0xa6, 0x60, 0x00, 0x01,
		0xa1, 0x04, 0xa7, 0x20, 0xd3, 0xaf, 0x23, 0xbe, 0x25, 0xe0, 0x69, 0x22, 0x3d, 0xb0, 0x00, 0xc3,
		0x32, 0xc9, 0x98, 0x8e, 0xdd, 0x07, 0xc5, 0xb7, 0x05, 0x45, 0x3e, 0x04, 0x13, 0xdb, 0xf0, 0x89,
		0xe7, 0xd7, 0x48, 0x3f, 0x24, 0xaf, 0xcb, 0x6d, 0x08, 0x88, 0x30, 0xe5, 0x26, 0xb1, 0xf5, 0xad,
		0xfe, 0x18, 0xbe, 0x23, 0x4d, 0x29, 0x31, 0x94, 0x62, 0x0e, 0x86, 0xea, 0xd8, 0xf5, 0xb6, 0xb0,
		0xd5, 0x97, 0x3b, 0xbe, 0x2b, 0x38, 0x72, 0x01, 0x48, 0x58, 0xa4, 0x69, 0xef, 0x86, 0xe6, 0x0d,
		0x69, 0x91, 0x10, 0x4c, 0x1c, 0x3d, 0xcf, 0x67, 0xf7, 0x2b, 0xbb, 0x61, 0xfb, 0x9e, 0x3c, 0x7a,
		0x1c, 0xbb, 0x1c, 0x66, 0x3c, 0x05, 0x19, 0xcf, 0xbc, 0xd2, 0x17, 0xcd, 0xf7, 0xa5, 0xa7, 0x19,
		0x80, 0x82, 0x1f, 0x83, 0xbd, 0x3d, 0x53, 0x7d, 0x1f, 0x64, 0x3f, 0x10, 0x64, 0xe3, 0x3d, 0xd2,
		0xbd, 0x48, 0x09, 0xbb, 0xa5, 0xfc, 0xa1, 0x4c, 0x09, 0xa4, 0x83, 0x6b, 0x8d, 0x76, 0xe7, 0x1e,
		0xae, 0xee, 0xce, 0x6a, 0x3f, 0x92, 0x56, 0xe3, 0xd8, 0x36, 0xab, 0xad, 0xc3, 0xb8, 0x60, 0xdc,
		0x9d, 0x5f, 0xdf, 0x94, 0x89, 0x95, 0xa3, 0x37, 0xda, 0xbd, 0xfb, 0x04, 0x4c, 0x04, 0xe6, 0x94,
		0x8d, 0xa5, 0xa7, 0xd5, 0x71, 0xa3, 0x0f, 0xe6, 0xb7, 0x04, 0xb3, 0xcc, 0xf8, 0x41, 0x67, 0xea,
		0x2d, 0xe3, 0x06, 0x25, 0x3f, 0x07, 0x05, 0x49, 0xde, 0xb4, 0x5d, 0xa2, 0x3b, 0x35, 0xdb, 0xbc,
		0x42, 0x8c, 0x3e, 0xa8, 0x7f, 0xdc, 0xe1, 0xaa, 0x8d, 0x10, 0x9c, 0x32, 0x2f, 0x82, 0x12, 0xf4,
		0x1b, 0x9a, 0x59, 0x6f, 0x38, 0xae, 0x1f, 0xc1, 0xf8, 0x13, 0xe9, 0xa9, 0x00, 0xb7, 0xc8, 0x60,
		0xc5, 0x32, 0xe4, 0xd9, 0xb0, 0xdf, 0x90, 0xfc, 0xa9, 0x20, 0x1a, 0x6a, 0xa1, 0x44, 0xe2, 0xd0,
		0x9d, 0x7a, 0x03, 0xbb, 0xfd, 0xe4, 0xbf, 0x9f, 0xc9, 0xc4, 0x21, 0x20, 0x3c, 0xfa, 0x86, 0x3b,
		0x2a, 0x31, 0x8a, 0xfa, 0x78, 0x5d, 0x78, 0xf2, 0x9a, 0x38, 0xb3, 0xed, 0x85, 0xb8, 0xb8, 0x44,
		0xcd, 0xd3, 0x5e, 0x2e, 0xa3, 0xc9, 0x9e, 0xba, 0x16, 0x58, 0xa8, 0xad, 0x5a, 0x16, 0xcf, 0xc0,
		0x50, 0x5b, 0xa9, 0x8c, 0xa6, 0xfa, 0x8c, 0xa0, 0xca, 0x85, 0x2b, 0x65, 0xf1, 0x18, 0x24, 0x69,
		0xd9, 0x8b, 0x86, 0x7f, 0x56, 0xc0, 0x99, 0x7a, 0xf1, 0x41, 0x48, 0xcb, 0x72, 0x17, 0x0d, 0xfd,
		0x9c, 0x80, 0x06, 0x10, 0x0a, 0x97, 0xa5, 0x2e, 0x1a, 0xfe, 0x79, 0x09, 0x97, 0x10, 0x0a, 0xef,
		0xdf, 0x84, 0xbf, 0x7c, 0x26, 0x29, 0xd2, 0x95, 0xb4, 0xdd, 0x29, 0x18, 0x14, 0x35, 0x2e, 0x1a,
		0xfd, 0xb4, 0x78, 0xb8, 0x44, 0x14, 0x1f, 0x80, 0x54, 0x9f, 0x06, 0xff, 0x82, 0x80, 0x72, 0xfd,
		0xe2, 0x1c, 0x64, 0x43, 0x75, 0x2d, 0x1a, 0xfe, 0x45, 0x01, 0x0f, 0xa3, 0xe8, 0xd2, 0x45, 0x5d,
		0x8b, 0x26, 0xf8, 0x92, 0x5c, 0xba, 0x40, 0x50, 0xb3, 0xc9, 0x92, 0x16, 0x8d, 0xfe, 0xb2, 0xb4,
		0xba, 0x84, 0x14, 0x4f, 0x43, 0x26, 0x48, 0x53, 0xd1, 0xf8, 0xaf, 0x08, 0x7c, 0x0b, 0x43, 0x2d,
		0x10, 0x4a, 0x93, 0xd1, 0x14, 0x5f, 0x95, 0x16, 0x08, 0xa1, 0xe8, 0x31, 0xea, 0x2c, 0x7d, 0xd1,
		0x4c, 0x5f, 0x93, 0xc7, 0xa8, 0xa3, 0xf2, 0x51, 0x6f, 0xb2, 0x6c, 0x11, 0x4d, 0xf1, 0x75, 0xe9,
		0x4d, 0xa6, 0x4f, 0x97, 0xd1, 0x59, 0x4b, 0xa2, 0x39, 0xbe, 0x21, 0x97, 0xd1, 0x51, 0x4a, 0x8a,
		0x6b, 0x80, 0xba, 0xeb, 0x48, 0x34, 0xdf, 0xb3, 0x82, 0x6f, 0xa4, 0xab, 0x8c, 0x14, 0x1f, 0x85,
		0xf1, 0xde, 0x35, 0x24, 0x9a, 0xf5, 0xb9, 0x6b, 0x1d, 0x5d, 0x7f, 0xb8, 0x84, 0x14, 0xd7, 0x5b,
		0x5d, 0x7f, 0xb8, 0x7e, 0x44, 0xd3, 0x3e, 0x7f, 0xad, 0xfd, 0xc5, 0x2e, 0x5c, 0x3e, 0x8a, 0x25,
		0x80, 0x56, 0xea, 0x8e, 0xe6, 0x7a, 0x51, 0x70, 0x85, 0x40, 0xf4, 0x68, 0x88, 0xcc, 0x1d, 0x8d,
		0x7f, 0x49, 0x1e, 0x0d, 0x81, 0x28, 0x9e, 0x82, 0xb4, 0xdd, 0xb4, 0x2c, 0x1a, 0x1c, 0xe8, 0xc6,
		0xff, 0x10, 0x52, 0xf8, 0xd3, 0x75, 0x71, 0x30, 0x24, 0xa0, 0x78, 0x0c, 0x52, 0xa4, 0xbe, 0x49,
		0x8c, 0x28, 0xe4, 0x9f, 0xaf, 0xcb, 0x84, 0x40, 0xb5, 0x8b, 0xa7, 0x01, 0xf8, 0x4b, 0x23, 0xfb,
		0x1e, 0x10, 0x81, 0xfd, 0xcb, 0x75, 0xf1, 0xad, 0xb9, 0x05, 0x69, 0x11, 0xf0, 0x2f, 0xd7, 0x37,
		0x26, 0x78, 0xaf, 0x9d, 0x80, 0xbd, 0x68, 0x9e, 0x84, 0xc1, 0xf3, 0x9e, 0x63, 0xfb, 0xb8, 0x16,
		0x85, 0xfe, 0xab, 0x40, 0x4b, 0x7d, 0x6a, 0xb0, 0xba, 0xe3, 0x12, 0x1f, 0xd7, 0xbc, 0x28, 0xec,
		0xdf, 0x04, 0x36, 0x00, 0x50, 0xb0, 0x8e, 0x3d, 0xbf, 0x9f, 0x7d, 0xff, 0x5d, 0x82, 0x25, 0x80,
		0x2e, 0x9a, 0xfe, 0xbe, 0x40, 0xb6, 0xa3, 0xb0, 0xef, 0xcb, 0x45, 0x0b, 0xfd, 0xe2, 0x83, 0x90,
		0xa1, 0x3f, 0xf9, 0xff, 0x5f, 0x44, 0x80, 0xff, 0x21, 0xc0, 0x2d, 0x04, 0x7d, 0xb2, 0xe7, 0x1b,
		0xbe, 0x19, 0x6d, 0xec, 0x7f, 0x0a, 0x4f, 0x4b, 0xfd, 0x62, 0x09, 0xb2, 0x9e, 0x6f, 0x18, 0x4d,
		0x97, 0x5f, 0x44, 0x45, 0xc0, 0xff, 0x75, 0x3d, 0x78, 0x99, 0x0b, 0x30, 0xb3, 0x07, 0x7a, 0xdf,
		0x2d, 0xc1, 0x82, 0xb3, 0xe0, 0xf0, 0x5b, 0x25, 0xf8, 0x4f, 0x0c, 0xe0, 0x82, 0xed, 0xf8, 0xe2,
		0xfa, 0x27, 0x45, 0x13, 0xbf, 0x37, 0xb1, 0xbb, 0x4b, 0xa3, 0x29, 0x07, 0x92, 0x9f, 0xb4, 0x1d,
		0x1f, 0xdd, 0x09, 0xa9, 0x59, 0xb3, 0xb6, 0xe5, 0x8b, 0x4b, 0x38, 0x65, 0x86, 0xb1, 0xcd, 0x30,
		0x19, 0x55, 0x50, 0xf9, 0x34, 0x1a, 0x83, 0x54, 0xd9, 0xda, 0x74, 0x2e, 0x89, 0x6f, 0x22, 0x7c,
		0x80, 0x0e, 0x41, 0x66, 0xd6, 0xa4, 0x8d, 0x4f, 0xd9, 0x36, 0xc4, 0xbf, 0x6c, 0x64, 0x05, 0x03,
		0x03, 0xb7, 0x66, 0xa7, 0x3e, 0x41, 0x55, 0x05, 0x29, 0xda, 0x0f, 0xc9, 0x25, 0xc7, 0x69, 0xb0,
		0xef, 0x13, 0x1d, 0x10, 0x36, 0x81, 0x10, 0x24, 0xd7, 0x9b, 0xae, 0x2d, 0x9e, 0xc6, 0x7e, 0xcf,
		0xa6, 0xdf, 0xff, 0xdd, 0xbe, 0xd8, 0x9b, 0xbf, 0xdf, 0x17, 0xfb, 0x6f, 0x00, 0x00, 0x00, 0xff,
		0xff, 0x6b, 0x98, 0x43, 0x97, 0x28, 0x2d, 0x00, 0x00,
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
func (this *Knot) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&tests.Knot{")
	if this.Bight != nil {
		s = append(s, "Bight: "+fmt.Sprintf("%#v", this.Bight)+",\n")
	}
	if this.Elbow != nil {
		s = append(s, "Elbow: "+valueToGoStringKnot(this.Elbow, "bool")+",\n")
	}
	if this.BitterEnd != nil {
		s = append(s, "BitterEnd: "+fmt.Sprintf("%#v", this.BitterEnd)+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *BightKnot) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&tests.BightKnot{")
	if this.Loop != nil {
		s = append(s, "Loop: "+fmt.Sprintf("%#v", this.Loop)+",\n")
	}
	if this.Turn != nil {
		s = append(s, "Turn: "+valueToGoStringKnot(this.Turn, "bool")+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringKnot(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func extensionToGoStringKnot(m github_com_gogo_protobuf_proto.Message) string {
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

func init() { proto.RegisterFile("knot.proto", fileDescriptorKnot) }

var fileDescriptorKnot = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xca, 0xce, 0xcb, 0x2f,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2d, 0x49, 0x2d, 0x2e, 0x29, 0x96, 0xd2, 0x4d,
	0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0x4f, 0xcf, 0xd7, 0x07,
	0xcb, 0x26, 0x95, 0xa6, 0x81, 0x79, 0x60, 0x0e, 0x98, 0x05, 0xd1, 0xa5, 0x94, 0xcf, 0xc5, 0xe2,
	0x9d, 0x97, 0x5f, 0x22, 0xa4, 0xc6, 0xc5, 0xea, 0x94, 0x99, 0x9e, 0x51, 0x22, 0xc1, 0xa8, 0xc0,
	0xac, 0xc1, 0x6d, 0x24, 0xa0, 0x07, 0x36, 0x4d, 0x0f, 0x2c, 0x06, 0x52, 0x10, 0x04, 0x91, 0x16,
	0x12, 0xe1, 0x62, 0x75, 0xcd, 0x49, 0xca, 0x2f, 0x97, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x08, 0x82,
	0x70, 0x84, 0x34, 0xb9, 0x38, 0x9d, 0x32, 0x4b, 0x4a, 0x52, 0x8b, 0x5c, 0xf3, 0x52, 0x24, 0x98,
	0x15, 0x18, 0x35, 0xb8, 0x8d, 0xb8, 0xa1, 0x26, 0x80, 0x35, 0x23, 0x64, 0x95, 0x1c, 0x40, 0x4a,
	0xa1, 0x86, 0x0a, 0xc9, 0x73, 0xb1, 0xf8, 0xe4, 0xe7, 0x17, 0x48, 0x30, 0x62, 0x6a, 0x01, 0x4b,
	0x08, 0x09, 0x71, 0xb1, 0x84, 0x94, 0x16, 0xe5, 0x41, 0x6d, 0x03, 0xb3, 0x9d, 0x38, 0x3e, 0x3c,
	0x94, 0x63, 0xdc, 0xf0, 0x48, 0x8e, 0x11, 0x10, 0x00, 0x00, 0xff, 0xff, 0x56, 0x97, 0x4a, 0xe2,
	0xff, 0x00, 0x00, 0x00,
}
