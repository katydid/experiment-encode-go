// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: protonum.proto

/*
Package protonum is a generated protocol buffer package.

It is generated from these files:
	protonum.proto

It has these top-level messages:
	ProtoNum
	KeyValue
	TopsyTurvy
	Topsy
	Turvy
	Knot
	BightKnot
*/
package protonum

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

// ProtoNum is used for testing of the protonum package.
type ProtoNum struct {
	KeyValue         []*KeyValue `protobuf:"bytes,2,rep,name=KeyValue" json:"KeyValue,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *ProtoNum) Reset()                    { *m = ProtoNum{} }
func (m *ProtoNum) String() string            { return proto.CompactTextString(m) }
func (*ProtoNum) ProtoMessage()               {}
func (*ProtoNum) Descriptor() ([]byte, []int) { return fileDescriptorProtonum, []int{0} }

func (m *ProtoNum) GetKeyValue() []*KeyValue {
	if m != nil {
		return m.KeyValue
	}
	return nil
}

// KeyValue is used for testing of the protonum package.
type KeyValue struct {
	Key              *string `protobuf:"bytes,1,opt,name=Key" json:"Key,omitempty"`
	Value            *string `protobuf:"bytes,2,opt,name=Value" json:"Value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *KeyValue) Reset()                    { *m = KeyValue{} }
func (m *KeyValue) String() string            { return proto.CompactTextString(m) }
func (*KeyValue) ProtoMessage()               {}
func (*KeyValue) Descriptor() ([]byte, []int) { return fileDescriptorProtonum, []int{1} }

func (m *KeyValue) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *KeyValue) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

// TopsyTurvy is used for testing of the protonum package.
type TopsyTurvy struct {
	Topsy            *Topsy `protobuf:"bytes,1,opt,name=Topsy" json:"Topsy,omitempty"`
	Turvy            *Turvy `protobuf:"bytes,2,opt,name=Turvy" json:"Turvy,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *TopsyTurvy) Reset()                    { *m = TopsyTurvy{} }
func (m *TopsyTurvy) String() string            { return proto.CompactTextString(m) }
func (*TopsyTurvy) ProtoMessage()               {}
func (*TopsyTurvy) Descriptor() ([]byte, []int) { return fileDescriptorProtonum, []int{2} }

func (m *TopsyTurvy) GetTopsy() *Topsy {
	if m != nil {
		return m.Topsy
	}
	return nil
}

func (m *TopsyTurvy) GetTurvy() *Turvy {
	if m != nil {
		return m.Turvy
	}
	return nil
}

// Topsy is used for testing of the protonum package.
type Topsy struct {
	A                *int64 `protobuf:"varint,1,opt,name=A" json:"A,omitempty"`
	B                *int64 `protobuf:"varint,2,opt,name=B" json:"B,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Topsy) Reset()                    { *m = Topsy{} }
func (m *Topsy) String() string            { return proto.CompactTextString(m) }
func (*Topsy) ProtoMessage()               {}
func (*Topsy) Descriptor() ([]byte, []int) { return fileDescriptorProtonum, []int{3} }

func (m *Topsy) GetA() int64 {
	if m != nil && m.A != nil {
		return *m.A
	}
	return 0
}

func (m *Topsy) GetB() int64 {
	if m != nil && m.B != nil {
		return *m.B
	}
	return 0
}

// Turvy is used for testing of the protonum package.
type Turvy struct {
	B                *int64 `protobuf:"varint,1,opt,name=B" json:"B,omitempty"`
	A                *int64 `protobuf:"varint,2,opt,name=A" json:"A,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Turvy) Reset()                    { *m = Turvy{} }
func (m *Turvy) String() string            { return proto.CompactTextString(m) }
func (*Turvy) ProtoMessage()               {}
func (*Turvy) Descriptor() ([]byte, []int) { return fileDescriptorProtonum, []int{4} }

func (m *Turvy) GetB() int64 {
	if m != nil && m.B != nil {
		return *m.B
	}
	return 0
}

func (m *Turvy) GetA() int64 {
	if m != nil && m.A != nil {
		return *m.A
	}
	return 0
}

// Knot is used for testing of the protonum package.
type Knot struct {
	Bight            []*BightKnot `protobuf:"bytes,1,rep,name=Bight" json:"Bight,omitempty"`
	Elbow            *bool        `protobuf:"varint,2,opt,name=Elbow" json:"Elbow,omitempty"`
	BitterEnd        *Knot        `protobuf:"bytes,3,opt,name=BitterEnd" json:"BitterEnd,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Knot) Reset()                    { *m = Knot{} }
func (m *Knot) String() string            { return proto.CompactTextString(m) }
func (*Knot) ProtoMessage()               {}
func (*Knot) Descriptor() ([]byte, []int) { return fileDescriptorProtonum, []int{5} }

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

// BightKnot is used for testing of the protonum package.
type BightKnot struct {
	Loop             *Knot  `protobuf:"bytes,1,opt,name=Loop" json:"Loop,omitempty"`
	Turn             *bool  `protobuf:"varint,2,opt,name=Turn" json:"Turn,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *BightKnot) Reset()                    { *m = BightKnot{} }
func (m *BightKnot) String() string            { return proto.CompactTextString(m) }
func (*BightKnot) ProtoMessage()               {}
func (*BightKnot) Descriptor() ([]byte, []int) { return fileDescriptorProtonum, []int{6} }

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
	proto.RegisterType((*ProtoNum)(nil), "protonum.ProtoNum")
	proto.RegisterType((*KeyValue)(nil), "protonum.KeyValue")
	proto.RegisterType((*TopsyTurvy)(nil), "protonum.TopsyTurvy")
	proto.RegisterType((*Topsy)(nil), "protonum.Topsy")
	proto.RegisterType((*Turvy)(nil), "protonum.Turvy")
	proto.RegisterType((*Knot)(nil), "protonum.Knot")
	proto.RegisterType((*BightKnot)(nil), "protonum.BightKnot")
}
func (this *ProtoNum) Description() (desc *github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet) {
	return ProtonumDescription()
}
func (this *KeyValue) Description() (desc *github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet) {
	return ProtonumDescription()
}
func (this *TopsyTurvy) Description() (desc *github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet) {
	return ProtonumDescription()
}
func (this *Topsy) Description() (desc *github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet) {
	return ProtonumDescription()
}
func (this *Turvy) Description() (desc *github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet) {
	return ProtonumDescription()
}
func (this *Knot) Description() (desc *github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet) {
	return ProtonumDescription()
}
func (this *BightKnot) Description() (desc *github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet) {
	return ProtonumDescription()
}
func ProtonumDescription() (desc *github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet) {
	d := &github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet{}
	var gzipped = []byte{
		// 3832 bytes of a gzipped FileDescriptorSet
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x5a, 0x5b, 0x70, 0x1b, 0xd7,
		0x79, 0x36, 0x6e, 0x24, 0xf0, 0x03, 0x04, 0x97, 0x87, 0xb4, 0x04, 0xd1, 0xb1, 0x25, 0xc1, 0x37,
		0xca, 0x4e, 0xa8, 0x8c, 0x6c, 0xc9, 0xf2, 0xaa, 0x89, 0x0b, 0x80, 0x2b, 0x06, 0x32, 0x49, 0x20,
		0x0b, 0xd0, 0x96, 0xdd, 0x87, 0x9d, 0xe5, 0xe2, 0x00, 0x5c, 0x69, 0xb1, 0xbb, 0xd9, 0x5d, 0x48,
		0xa6, 0xa6, 0x0f, 0xee, 0xb8, 0x97, 0xc9, 0x74, 0x7a, 0xef, 0x4c, 0x13, 0xd7, 0x71, 0xdb, 0xcc,
		0xa4, 0x6e, 0x93, 0x5e, 0x92, 0xa6, 0x4d, 0x2f, 0x4f, 0x7d, 0x49, 0xdb, 0xa7, 0xce, 0xe4, 0xbd,
		0x0f, 0x6d, 0xc7, 0x33, 0xbd, 0xb9, 0x6d, 0xda, 0xf8, 0x4d, 0x2f, 0x99, 0x73, 0x5b, 0xec, 0x02,
		0xa0, 0x16, 0xcc, 0x8c, 0x9d, 0x27, 0xf2, 0xfc, 0xe7, 0xff, 0xbe, 0xf3, 0x9f, 0xff, 0xfc, 0xe7,
		0xfc, 0xff, 0x39, 0x0b, 0xf8, 0xc6, 0x55, 0x38, 0x37, 0x70, 0x9c, 0x81, 0x85, 0x2f, 0xba, 0x9e,
		0x13, 0x38, 0x07, 0xa3, 0xfe, 0xc5, 0x1e, 0xf6, 0x0d, 0xcf, 0x74, 0x03, 0xc7, 0xdb, 0xa4, 0x32,
		0xb4, 0xcc, 0x34, 0x36, 0x85, 0x46, 0x75, 0x17, 0x56, 0xae, 0x9b, 0x16, 0xde, 0x0a, 0x15, 0x3b,
		0x38, 0x40, 0x57, 0x21, 0xdb, 0x37, 0x2d, 0x5c, 0x49, 0x9d, 0xcb, 0x6c, 0x14, 0x2f, 0x3d, 0xb1,
		0x39, 0x01, 0xda, 0x8c, 0x23, 0xda, 0x44, 0xac, 0x52, 0x44, 0xf5, 0xfd, 0x2c, 0xac, 0xce, 0xe8,
		0x45, 0x08, 0xb2, 0xb6, 0x3e, 0x24, 0x8c, 0xa9, 0x8d, 0x82, 0x4a, 0xff, 0x47, 0x15, 0x58, 0x74,
		0x75, 0xe3, 0xb6, 0x3e, 0xc0, 0x95, 0x34, 0x15, 0x8b, 0x26, 0x7a, 0x0c, 0xa0, 0x87, 0x5d, 0x6c,
		0xf7, 0xb0, 0x6d, 0x1c, 0x55, 0x32, 0xe7, 0x32, 0x1b, 0x05, 0x35, 0x22, 0x41, 0xcf, 0xc2, 0x8a,
		0x3b, 0x3a, 0xb0, 0x4c, 0x43, 0x8b, 0xa8, 0xc1, 0xb9, 0xcc, 0x46, 0x4e, 0x95, 0x58, 0xc7, 0xd6,
		0x58, 0xf9, 0x69, 0x58, 0xbe, 0x8b, 0xf5, 0xdb, 0x51, 0xd5, 0x22, 0x55, 0x2d, 0x13, 0x71, 0x44,
		0xb1, 0x01, 0xa5, 0x21, 0xf6, 0x7d, 0x7d, 0x80, 0xb5, 0xe0, 0xc8, 0xc5, 0x95, 0x2c, 0x9d, 0xfd,
		0xb9, 0xa9, 0xd9, 0x4f, 0xce, 0xbc, 0xc8, 0x51, 0xdd, 0x23, 0x17, 0xa3, 0x1a, 0x14, 0xb0, 0x3d,
		0x1a, 0x32, 0x86, 0xdc, 0x31, 0xfe, 0x53, 0xec, 0xd1, 0x70, 0x92, 0x25, 0x4f, 0x60, 0x9c, 0x62,
		0xd1, 0xc7, 0xde, 0x1d, 0xd3, 0xc0, 0x95, 0x05, 0x4a, 0xf0, 0xf4, 0x14, 0x41, 0x87, 0xf5, 0x4f,
		0x72, 0x08, 0x1c, 0x6a, 0x40, 0x01, 0xbf, 0x11, 0x60, 0xdb, 0x37, 0x1d, 0xbb, 0xb2, 0x48, 0x49,
		0x9e, 0x9c, 0xb1, 0x8a, 0xd8, 0xea, 0x4d, 0x52, 0x8c, 0x71, 0xe8, 0x0a, 0x2c, 0x3a, 0x6e, 0x60,
		0x3a, 0xb6, 0x5f, 0xc9, 0x9f, 0x4b, 0x6d, 0x14, 0x2f, 0x7d, 0x62, 0x66, 0x20, 0xb4, 0x98, 0x8e,
		0x2a, 0x94, 0x51, 0x13, 0x24, 0xdf, 0x19, 0x79, 0x06, 0xd6, 0x0c, 0xa7, 0x87, 0x35, 0xd3, 0xee,
		0x3b, 0x95, 0x02, 0x25, 0x38, 0x3b, 0x3d, 0x11, 0xaa, 0xd8, 0x70, 0x7a, 0xb8, 0x69, 0xf7, 0x1d,
		0xb5, 0xec, 0xc7, 0xda, 0xe8, 0x14, 0x2c, 0xf8, 0x47, 0x76, 0xa0, 0xbf, 0x51, 0x29, 0xd1, 0x08,
		0xe1, 0xad, 0xea, 0x5f, 0x2f, 0xc0, 0xf2, 0x3c, 0x21, 0x76, 0x0d, 0x72, 0x7d, 0x32, 0xcb, 0x4a,
		0xfa, 0x24, 0x3e, 0x60, 0x98, 0xb8, 0x13, 0x17, 0x7e, 0x44, 0x27, 0xd6, 0xa0, 0x68, 0x63, 0x3f,
		0xc0, 0x3d, 0x16, 0x11, 0x99, 0x39, 0x63, 0x0a, 0x18, 0x68, 0x3a, 0xa4, 0xb2, 0x3f, 0x52, 0x48,
		0xdd, 0x84, 0xe5, 0xd0, 0x24, 0xcd, 0xd3, 0xed, 0x81, 0x88, 0xcd, 0x8b, 0x49, 0x96, 0x6c, 0x2a,
		0x02, 0xa7, 0x12, 0x98, 0x5a, 0xc6, 0xb1, 0x36, 0xda, 0x02, 0x70, 0x6c, 0xec, 0xf4, 0xb5, 0x1e,
		0x36, 0xac, 0x4a, 0xfe, 0x18, 0x2f, 0xb5, 0x88, 0xca, 0x94, 0x97, 0x1c, 0x26, 0x35, 0x2c, 0xf4,
		0xe2, 0x38, 0xd4, 0x16, 0x8f, 0x89, 0x94, 0x5d, 0xb6, 0xc9, 0xa6, 0xa2, 0x6d, 0x1f, 0xca, 0x1e,
		0x26, 0x71, 0x8f, 0x7b, 0x7c, 0x66, 0x05, 0x6a, 0xc4, 0x66, 0xe2, 0xcc, 0x54, 0x0e, 0x63, 0x13,
		0x5b, 0xf2, 0xa2, 0x4d, 0xf4, 0x38, 0x84, 0x02, 0x8d, 0x86, 0x15, 0xd0, 0x53, 0xa8, 0x24, 0x84,
		0x7b, 0xfa, 0x10, 0xaf, 0xdf, 0x83, 0x72, 0xdc, 0x3d, 0x68, 0x0d, 0x72, 0x7e, 0xa0, 0x7b, 0x01,
		0x8d, 0xc2, 0x9c, 0xca, 0x1a, 0x48, 0x82, 0x0c, 0xb6, 0x7b, 0xf4, 0x94, 0xcb, 0xa9, 0xe4, 0x5f,
		0xf4, 0x93, 0xe3, 0x09, 0x67, 0xe8, 0x84, 0x9f, 0x9a, 0x5e, 0xd1, 0x18, 0xf3, 0xe4, 0xbc, 0xd7,
		0x5f, 0x80, 0xa5, 0xd8, 0x04, 0xe6, 0x1d, 0xba, 0xfa, 0xd3, 0xf0, 0xf0, 0x4c, 0x6a, 0x74, 0x13,
		0xd6, 0x46, 0xb6, 0x69, 0x07, 0xd8, 0x73, 0x3d, 0x4c, 0x22, 0x96, 0x0d, 0x55, 0xf9, 0xb7, 0xc5,
		0x63, 0x62, 0x6e, 0x3f, 0xaa, 0xcd, 0x58, 0xd4, 0xd5, 0xd1, 0xb4, 0xf0, 0x99, 0x42, 0xfe, 0xdf,
		0x17, 0xa5, 0x37, 0xdf, 0x7c, 0xf3, 0xcd, 0x74, 0xf5, 0x4b, 0x0b, 0xb0, 0x36, 0x6b, 0xcf, 0xcc,
		0xdc, 0xbe, 0xa7, 0x60, 0xc1, 0x1e, 0x0d, 0x0f, 0xb0, 0x47, 0x9d, 0x94, 0x53, 0x79, 0x0b, 0xd5,
		0x20, 0x67, 0xe9, 0x07, 0xd8, 0xaa, 0x64, 0xcf, 0xa5, 0x36, 0xca, 0x97, 0x9e, 0x9d, 0x6b, 0x57,
		0x6e, 0xee, 0x10, 0x88, 0xca, 0x90, 0xe8, 0xb3, 0x90, 0xe5, 0x47, 0x34, 0x61, 0x78, 0x66, 0x3e,
		0x06, 0xb2, 0x97, 0x54, 0x8a, 0x43, 0x8f, 0x40, 0x81, 0xfc, 0x65, 0xb1, 0xb1, 0x40, 0x6d, 0xce,
		0x13, 0x01, 0x89, 0x0b, 0xb4, 0x0e, 0x79, 0xba, 0x4d, 0x7a, 0x58, 0xa4, 0xb6, 0xb0, 0x4d, 0x02,
		0xab, 0x87, 0xfb, 0xfa, 0xc8, 0x0a, 0xb4, 0x3b, 0xba, 0x35, 0xc2, 0x34, 0xe0, 0x0b, 0x6a, 0x89,
		0x0b, 0x5f, 0x21, 0x32, 0x74, 0x16, 0x8a, 0x6c, 0x57, 0x99, 0x76, 0x0f, 0xbf, 0x41, 0x4f, 0xcf,
		0x9c, 0xca, 0x36, 0x5a, 0x93, 0x48, 0xc8, 0xf0, 0xb7, 0x7c, 0xc7, 0x16, 0xa1, 0x49, 0x87, 0x20,
		0x02, 0x3a, 0xfc, 0x0b, 0x93, 0x07, 0xf7, 0xa3, 0xb3, 0xa7, 0x37, 0x19, 0x53, 0xd5, 0xef, 0xa4,
		0x21, 0x4b, 0xcf, 0x8b, 0x65, 0x28, 0x76, 0x5f, 0x6b, 0x2b, 0xda, 0x56, 0x6b, 0xbf, 0xbe, 0xa3,
		0x48, 0x29, 0x54, 0x06, 0xa0, 0x82, 0xeb, 0x3b, 0xad, 0x5a, 0x57, 0x4a, 0x87, 0xed, 0xe6, 0x5e,
		0xf7, 0xca, 0xf3, 0x52, 0x26, 0x04, 0xec, 0x33, 0x41, 0x36, 0xaa, 0xf0, 0xdc, 0x25, 0x29, 0x87,
		0x24, 0x28, 0x31, 0x82, 0xe6, 0x4d, 0x65, 0xeb, 0xca, 0xf3, 0xd2, 0x42, 0x5c, 0xf2, 0xdc, 0x25,
		0x69, 0x11, 0x2d, 0x41, 0x81, 0x4a, 0xea, 0xad, 0xd6, 0x8e, 0x94, 0x0f, 0x39, 0x3b, 0x5d, 0xb5,
		0xb9, 0xb7, 0x2d, 0x15, 0x42, 0xce, 0x6d, 0xb5, 0xb5, 0xdf, 0x96, 0x20, 0x64, 0xd8, 0x55, 0x3a,
		0x9d, 0xda, 0xb6, 0x22, 0x15, 0x43, 0x8d, 0xfa, 0x6b, 0x5d, 0xa5, 0x23, 0x95, 0x62, 0x66, 0x3d,
		0x77, 0x49, 0x5a, 0x0a, 0x87, 0x50, 0xf6, 0xf6, 0x77, 0xa5, 0x32, 0x5a, 0x81, 0x25, 0x36, 0x84,
		0x30, 0x62, 0x79, 0x42, 0x74, 0xe5, 0x79, 0x49, 0x1a, 0x1b, 0xc2, 0x58, 0x56, 0x62, 0x82, 0x2b,
		0xcf, 0x4b, 0xa8, 0xda, 0x80, 0x1c, 0x8d, 0x2e, 0x84, 0xa0, 0xbc, 0x53, 0xab, 0x2b, 0x3b, 0x5a,
		0xab, 0xdd, 0x6d, 0xb6, 0xf6, 0x6a, 0x3b, 0x52, 0x6a, 0x2c, 0x53, 0x95, 0xcf, 0xef, 0x37, 0x55,
		0x65, 0x4b, 0x4a, 0x47, 0x65, 0x6d, 0xa5, 0xd6, 0x55, 0xb6, 0xa4, 0x4c, 0xd5, 0x80, 0xb5, 0x59,
		0xe7, 0xe4, 0xcc, 0x9d, 0x11, 0x59, 0xe2, 0xf4, 0x31, 0x4b, 0x4c, 0xb9, 0xa6, 0x96, 0xf8, 0xab,
		0x29, 0x58, 0x9d, 0x91, 0x2b, 0x66, 0x0e, 0xf2, 0x12, 0xe4, 0x58, 0x88, 0xb2, 0xec, 0x79, 0x61,
		0x66, 0xd2, 0xa1, 0x01, 0x3b, 0x95, 0x41, 0x29, 0x2e, 0x5a, 0x41, 0x64, 0x8e, 0xa9, 0x20, 0x08,
		0xc5, 0x94, 0x91, 0x6f, 0xa5, 0xa0, 0x72, 0x1c, 0x77, 0xc2, 0x41, 0x91, 0x8e, 0x1d, 0x14, 0xd7,
		0x26, 0x0d, 0x38, 0x7f, 0xfc, 0x1c, 0xa6, 0xac, 0x78, 0x2f, 0x05, 0xa7, 0x66, 0x17, 0x5a, 0x33,
		0x6d, 0xf8, 0x2c, 0x2c, 0x0c, 0x71, 0x70, 0xe8, 0x88, 0x62, 0xe3, 0xa9, 0x19, 0x29, 0x8c, 0x74,
		0x4f, 0xfa, 0x8a, 0xa3, 0xa2, 0x39, 0x30, 0x73, 0x5c, 0xb5, 0xc4, 0xac, 0x99, 0xb2, 0xf4, 0x8b,
		0x69, 0x78, 0x78, 0x26, 0xf9, 0x4c, 0x43, 0x1f, 0x05, 0x30, 0x6d, 0x77, 0x14, 0xb0, 0x82, 0x82,
		0x9d, 0x4f, 0x05, 0x2a, 0xa1, 0x7b, 0x9f, 0x9c, 0x3d, 0xa3, 0x20, 0xec, 0xcf, 0xd0, 0x7e, 0x60,
		0x22, 0xaa, 0x70, 0x75, 0x6c, 0x68, 0x96, 0x1a, 0xfa, 0xd8, 0x31, 0x33, 0x9d, 0xca, 0xd5, 0x9f,
		0x06, 0xc9, 0xb0, 0x4c, 0x6c, 0x07, 0x9a, 0x1f, 0x78, 0x58, 0x1f, 0x9a, 0xf6, 0x80, 0x1e, 0xc0,
		0x79, 0x39, 0xd7, 0xd7, 0x2d, 0x1f, 0xab, 0xcb, 0xac, 0xbb, 0x23, 0x7a, 0x09, 0x82, 0xe6, 0x38,
		0x2f, 0x82, 0x58, 0x88, 0x21, 0x58, 0x77, 0x88, 0xa8, 0x7e, 0x3b, 0x0f, 0xc5, 0x48, 0x59, 0x8a,
		0xce, 0x43, 0xe9, 0x96, 0x7e, 0x47, 0xd7, 0xc4, 0x55, 0x83, 0x79, 0xa2, 0x48, 0x64, 0x6d, 0x7e,
		0xdd, 0xf8, 0x34, 0xac, 0x51, 0x15, 0x67, 0x14, 0x60, 0x4f, 0x33, 0x2c, 0xdd, 0xf7, 0xa9, 0xd3,
		0xf2, 0x54, 0x15, 0x91, 0xbe, 0x16, 0xe9, 0x6a, 0x88, 0x1e, 0x74, 0x19, 0x56, 0x29, 0x62, 0x38,
		0xb2, 0x02, 0xd3, 0xb5, 0xb0, 0x46, 0x2e, 0x3f, 0x3e, 0x3d, 0x88, 0x43, 0xcb, 0x56, 0x88, 0xc6,
		0x2e, 0x57, 0x20, 0x16, 0xf9, 0x68, 0x0b, 0x1e, 0xa5, 0xb0, 0x01, 0xb6, 0xb1, 0xa7, 0x07, 0x58,
		0xc3, 0x5f, 0x18, 0xe9, 0x96, 0xaf, 0xe9, 0x76, 0x4f, 0x3b, 0xd4, 0xfd, 0xc3, 0xca, 0x1a, 0x21,
		0xa8, 0xa7, 0x2b, 0x29, 0xf5, 0x0c, 0x51, 0xdc, 0xe6, 0x7a, 0x0a, 0x55, 0xab, 0xd9, 0xbd, 0xcf,
		0xe9, 0xfe, 0x21, 0x92, 0xe1, 0x14, 0x65, 0xf1, 0x03, 0xcf, 0xb4, 0x07, 0x9a, 0x71, 0x88, 0x8d,
		0xdb, 0xda, 0x28, 0xe8, 0x5f, 0xad, 0x3c, 0x12, 0x1d, 0x9f, 0x5a, 0xd8, 0xa1, 0x3a, 0x0d, 0xa2,
		0xb2, 0x1f, 0xf4, 0xaf, 0xa2, 0x0e, 0x94, 0xc8, 0x62, 0x0c, 0xcd, 0x7b, 0x58, 0xeb, 0x3b, 0x1e,
		0xcd, 0x2c, 0xe5, 0x19, 0x3b, 0x3b, 0xe2, 0xc1, 0xcd, 0x16, 0x07, 0xec, 0x3a, 0x3d, 0x2c, 0xe7,
		0x3a, 0x6d, 0x45, 0xd9, 0x52, 0x8b, 0x82, 0xe5, 0xba, 0xe3, 0x91, 0x80, 0x1a, 0x38, 0xa1, 0x83,
		0x8b, 0x2c, 0xa0, 0x06, 0x8e, 0x70, 0xef, 0x65, 0x58, 0x35, 0x0c, 0x36, 0x67, 0xd3, 0xd0, 0xf8,
		0x15, 0xc5, 0xaf, 0x48, 0x31, 0x67, 0x19, 0xc6, 0x36, 0x53, 0xe0, 0x31, 0xee, 0xa3, 0x17, 0xe1,
		0xe1, 0xb1, 0xb3, 0xa2, 0xc0, 0x95, 0xa9, 0x59, 0x4e, 0x42, 0x2f, 0xc3, 0xaa, 0x7b, 0x34, 0x0d,
		0x44, 0xb1, 0x11, 0xdd, 0xa3, 0x49, 0xd8, 0x0b, 0xb0, 0xe6, 0x1e, 0xba, 0xd3, 0xb8, 0xd5, 0x28,
		0x0e, 0xb9, 0x87, 0xee, 0x24, 0xf0, 0x49, 0x7a, 0x5f, 0xf5, 0xb0, 0xa1, 0x07, 0xb8, 0x57, 0x39,
		0x1d, 0x55, 0x8f, 0x74, 0xa0, 0x8b, 0x20, 0x19, 0x86, 0x86, 0x6d, 0xfd, 0xc0, 0xc2, 0x9a, 0xee,
		0x61, 0x5b, 0xf7, 0x2b, 0x67, 0xa3, 0xca, 0x65, 0xc3, 0x50, 0x68, 0x6f, 0x8d, 0x76, 0xa2, 0x67,
		0x60, 0xc5, 0x39, 0xb8, 0x65, 0xb0, 0x90, 0xd4, 0x5c, 0x0f, 0xf7, 0xcd, 0x37, 0x2a, 0x4f, 0x50,
		0xff, 0x2e, 0x93, 0x0e, 0x1a, 0x90, 0x6d, 0x2a, 0x46, 0x17, 0x40, 0x32, 0xfc, 0x43, 0xdd, 0x73,
		0x69, 0x4d, 0xe0, 0xbb, 0xba, 0x81, 0x2b, 0x4f, 0x32, 0x55, 0x26, 0xdf, 0x13, 0x62, 0xb2, 0x25,
		0xfc, 0xbb, 0x66, 0x3f, 0x10, 0x8c, 0x4f, 0xb3, 0x2d, 0x41, 0x65, 0x9c, 0x6d, 0x03, 0x24, 0xe2,
		0x8a, 0xd8, 0xc0, 0x1b, 0x54, 0xad, 0xec, 0x1e, 0xba, 0xd1, 0x71, 0x1f, 0x87, 0x25, 0xa2, 0x39,
		0x1e, 0xf4, 0x02, 0xab, 0x67, 0xdc, 0xc3, 0xc8, 0x88, 0x1f, 0x59, 0x69, 0x59, 0x95, 0xa1, 0x14,
		0x8d, 0x4f, 0x54, 0x00, 0x16, 0xa1, 0x52, 0x8a, 0xe4, 0xfa, 0x46, 0x6b, 0x8b, 0x64, 0xe9, 0xd7,
		0x15, 0x29, 0x4d, 0xaa, 0x85, 0x9d, 0x66, 0x57, 0xd1, 0xd4, 0xfd, 0xbd, 0x6e, 0x73, 0x57, 0x91,
		0x32, 0xd1, 0xb2, 0xf4, 0xbb, 0x69, 0x28, 0xc7, 0x6f, 0x18, 0xe8, 0x27, 0xe0, 0xb4, 0x78, 0x0e,
		0xf0, 0x71, 0xa0, 0xdd, 0x35, 0x3d, 0xba, 0x65, 0x86, 0x3a, 0xab, 0xb0, 0xc3, 0x45, 0x5b, 0xe3,
		0x5a, 0x1d, 0x1c, 0xbc, 0x6a, 0x7a, 0x64, 0x43, 0x0c, 0xf5, 0x00, 0xed, 0xc0, 0x59, 0xdb, 0xd1,
		0xfc, 0x40, 0xb7, 0x7b, 0xba, 0xd7, 0xd3, 0xc6, 0x0f, 0x31, 0x9a, 0x6e, 0x18, 0xd8, 0xf7, 0x1d,
		0x96, 0xaa, 0x42, 0x96, 0x4f, 0xd8, 0x4e, 0x87, 0x2b, 0x8f, 0xcf, 0xf0, 0x1a, 0x57, 0x9d, 0x08,
		0xb0, 0xcc, 0x71, 0x01, 0xf6, 0x08, 0x14, 0x86, 0xba, 0xab, 0x61, 0x3b, 0xf0, 0x8e, 0x68, 0x5d,
		0x99, 0x57, 0xf3, 0x43, 0xdd, 0x55, 0x48, 0xfb, 0xe3, 0x29, 0xef, 0xff, 0x29, 0x03, 0xa5, 0x68,
		0x6d, 0x49, 0x4a, 0x75, 0x83, 0xe6, 0x91, 0x14, 0x3d, 0x69, 0x1e, 0x7f, 0x60, 0x25, 0xba, 0xd9,
		0x20, 0x09, 0x46, 0x5e, 0x60, 0x15, 0x9f, 0xca, 0x90, 0x24, 0xb9, 0x93, 0xb3, 0x05, 0xb3, 0x5b,
		0x4c, 0x5e, 0xe5, 0x2d, 0xb4, 0x0d, 0x0b, 0xb7, 0x7c, 0xca, 0xbd, 0x40, 0xb9, 0x9f, 0x78, 0x30,
		0xf7, 0x8d, 0x0e, 0x25, 0x2f, 0xdc, 0xe8, 0x68, 0x7b, 0x2d, 0x75, 0xb7, 0xb6, 0xa3, 0x72, 0x38,
		0x3a, 0x03, 0x59, 0x4b, 0xbf, 0x77, 0x14, 0x4f, 0x45, 0x54, 0x34, 0xaf, 0xe3, 0xcf, 0x40, 0xf6,
		0x2e, 0xd6, 0x6f, 0xc7, 0x13, 0x00, 0x15, 0x7d, 0x84, 0xa1, 0x7f, 0x11, 0x72, 0xd4, 0x5f, 0x08,
		0x80, 0x7b, 0x4c, 0x7a, 0x08, 0xe5, 0x21, 0xdb, 0x68, 0xa9, 0x24, 0xfc, 0x25, 0x28, 0x31, 0xa9,
		0xd6, 0x6e, 0x2a, 0x0d, 0x45, 0x4a, 0x57, 0x2f, 0xc3, 0x02, 0x73, 0x02, 0xd9, 0x1a, 0xa1, 0x1b,
		0xa4, 0x87, 0x78, 0x93, 0x73, 0xa4, 0x44, 0xef, 0xfe, 0x6e, 0x5d, 0x51, 0xa5, 0x74, 0x74, 0x79,
		0x7d, 0x28, 0x45, 0xcb, 0xca, 0x8f, 0x27, 0xa6, 0xfe, 0x26, 0x05, 0xc5, 0x48, 0x99, 0x48, 0x0a,
		0x14, 0xdd, 0xb2, 0x9c, 0xbb, 0x9a, 0x6e, 0x99, 0xba, 0xcf, 0x83, 0x02, 0xa8, 0xa8, 0x46, 0x24,
		0xf3, 0x2e, 0xda, 0xc7, 0x62, 0xfc, 0xbb, 0x29, 0x90, 0x26, 0x4b, 0xcc, 0x09, 0x03, 0x53, 0x3f,
		0x56, 0x03, 0xdf, 0x49, 0x41, 0x39, 0x5e, 0x57, 0x4e, 0x98, 0x77, 0xfe, 0xc7, 0x6a, 0xde, 0x3f,
		0xa7, 0x61, 0x29, 0x56, 0x4d, 0xce, 0x6b, 0xdd, 0x17, 0x60, 0xc5, 0xec, 0xe1, 0xa1, 0xeb, 0x04,
		0xd8, 0x36, 0x8e, 0x34, 0x0b, 0xdf, 0xc1, 0x56, 0xa5, 0x4a, 0x0f, 0x8a, 0x8b, 0x0f, 0xae, 0x57,
		0x37, 0x9b, 0x63, 0xdc, 0x0e, 0x81, 0xc9, 0xab, 0xcd, 0x2d, 0x65, 0xb7, 0xdd, 0xea, 0x2a, 0x7b,
		0x8d, 0xd7, 0xb4, 0xfd, 0xbd, 0x97, 0xf7, 0x5a, 0xaf, 0xee, 0xa9, 0x92, 0x39, 0xa1, 0xf6, 0x11,
		0x6e, 0xf5, 0x36, 0x48, 0x93, 0x46, 0xa1, 0xd3, 0x30, 0xcb, 0x2c, 0xe9, 0x21, 0xb4, 0x0a, 0xcb,
		0x7b, 0x2d, 0xad, 0xd3, 0xdc, 0x52, 0x34, 0xe5, 0xfa, 0x75, 0xa5, 0xd1, 0xed, 0xb0, 0x0b, 0x7c,
		0xa8, 0xdd, 0x8d, 0x6f, 0xea, 0xb7, 0x33, 0xb0, 0x3a, 0xc3, 0x12, 0x54, 0xe3, 0x77, 0x07, 0x76,
		0x9d, 0xf9, 0xd4, 0x3c, 0xd6, 0x6f, 0x92, 0x94, 0xdf, 0xd6, 0xbd, 0x80, 0x5f, 0x35, 0x2e, 0x00,
		0xf1, 0x92, 0x1d, 0x98, 0x7d, 0x13, 0x7b, 0xfc, 0xbd, 0x83, 0x5d, 0x28, 0x96, 0xc7, 0x72, 0xf6,
		0xe4, 0xf1, 0x49, 0x40, 0xae, 0xe3, 0x9b, 0x81, 0x79, 0x07, 0x6b, 0xa6, 0x2d, 0x1e, 0x47, 0xc8,
		0x05, 0x23, 0xab, 0x4a, 0xa2, 0xa7, 0x69, 0x07, 0xa1, 0xb6, 0x8d, 0x07, 0xfa, 0x84, 0x36, 0x39,
		0xc0, 0x33, 0xaa, 0x24, 0x7a, 0x42, 0xed, 0xf3, 0x50, 0xea, 0x39, 0x23, 0x52, 0x75, 0x31, 0x3d,
		0x92, 0x2f, 0x52, 0x6a, 0x91, 0xc9, 0x42, 0x15, 0x5e, 0x4f, 0x8f, 0x5f, 0x65, 0x4a, 0x6a, 0x91,
		0xc9, 0x98, 0xca, 0xd3, 0xb0, 0xac, 0x0f, 0x06, 0x1e, 0x21, 0x17, 0x44, 0xec, 0x86, 0x50, 0x0e,
		0xc5, 0x54, 0x71, 0xfd, 0x06, 0xe4, 0x85, 0x1f, 0x48, 0x4a, 0x26, 0x9e, 0xd0, 0x5c, 0xf6, 0x32,
		0x97, 0xde, 0x28, 0xa8, 0x79, 0x5b, 0x74, 0x9e, 0x87, 0x92, 0xe9, 0x6b, 0xe3, 0x47, 0xe6, 0xf4,
		0xb9, 0xf4, 0x46, 0x5e, 0x2d, 0x9a, 0x7e, 0xf8, 0x40, 0x57, 0x7d, 0x2f, 0x0d, 0xe5, 0xf8, 0x23,
		0x39, 0xda, 0x82, 0xbc, 0xe5, 0x18, 0x3a, 0x0d, 0x2d, 0xf6, 0x85, 0x66, 0x23, 0xe1, 0x5d, 0x7d,
		0x73, 0x87, 0xeb, 0xab, 0x21, 0x72, 0xfd, 0x1f, 0x53, 0x90, 0x17, 0x62, 0x74, 0x0a, 0xb2, 0xae,
		0x1e, 0x1c, 0x52, 0xba, 0x5c, 0x3d, 0x2d, 0xa5, 0x54, 0xda, 0x26, 0x72, 0xdf, 0xd5, 0x6d, 0x1a,
		0x02, 0x5c, 0x4e, 0xda, 0x64, 0x5d, 0x2d, 0xac, 0xf7, 0xe8, 0xf5, 0xc3, 0x19, 0x0e, 0xb1, 0x1d,
		0xf8, 0x62, 0x5d, 0xb9, 0xbc, 0xc1, 0xc5, 0xe8, 0x59, 0x58, 0x09, 0x3c, 0xdd, 0xb4, 0x62, 0xba,
		0x59, 0xaa, 0x2b, 0x89, 0x8e, 0x50, 0x59, 0x86, 0x33, 0x82, 0xb7, 0x87, 0x03, 0xdd, 0x38, 0xc4,
		0xbd, 0x31, 0x68, 0x81, 0xbe, 0xc0, 0x9e, 0xe6, 0x0a, 0x5b, 0xbc, 0x5f, 0x60, 0xab, 0xdf, 0x4b,
		0xc1, 0x8a, 0xb8, 0x30, 0xf5, 0x42, 0x67, 0xed, 0x02, 0xe8, 0xb6, 0xed, 0x04, 0x51, 0x77, 0x4d,
		0x87, 0xf2, 0x14, 0x6e, 0xb3, 0x16, 0x82, 0xd4, 0x08, 0xc1, 0xfa, 0x10, 0x60, 0xdc, 0x73, 0xac,
		0xdb, 0xce, 0x42, 0x91, 0x7f, 0x01, 0xa1, 0x9f, 0xd1, 0xd8, 0x15, 0x1b, 0x98, 0x88, 0xdc, 0xac,
		0xd0, 0x1a, 0xe4, 0x0e, 0xf0, 0xc0, 0xb4, 0xf9, 0xbb, 0x26, 0x6b, 0x88, 0xb7, 0xda, 0x6c, 0xf8,
		0x56, 0x5b, 0xbf, 0x09, 0xab, 0x86, 0x33, 0x9c, 0x34, 0xb7, 0x2e, 0x4d, 0x5c, 0xf3, 0xfd, 0xcf,
		0xa5, 0x5e, 0x87, 0x71, 0x89, 0xf9, 0xd5, 0x74, 0x66, 0xbb, 0x5d, 0xff, 0x7a, 0x7a, 0x7d, 0x9b,
		0xe1, 0xda, 0x62, 0x9a, 0x2a, 0xee, 0x5b, 0xd8, 0x20, 0xa6, 0xc3, 0x0f, 0x9e, 0x82, 0x4f, 0x0d,
		0xcc, 0xe0, 0x70, 0x74, 0xb0, 0x69, 0x38, 0xc3, 0x8b, 0x03, 0x67, 0xe0, 0x8c, 0x3f, 0x1b, 0x92,
		0x16, 0x6d, 0xd0, 0xff, 0xf8, 0xa7, 0xc3, 0x42, 0x28, 0x5d, 0x4f, 0xfc, 0xce, 0x28, 0xef, 0xc1,
		0x2a, 0x57, 0xd6, 0xe8, 0xb7, 0x0b, 0x76, 0x85, 0x40, 0x0f, 0x7c, 0xff, 0xa9, 0x7c, 0xeb, 0x7d,
		0x9a, 0xab, 0xd5, 0x15, 0x0e, 0x25, 0x7d, 0xec, 0x96, 0x21, 0xab, 0xf0, 0x70, 0x8c, 0x8f, 0xed,
		0x4b, 0xec, 0x25, 0x30, 0x7e, 0x97, 0x33, 0xae, 0x46, 0x18, 0x3b, 0x1c, 0x2a, 0x37, 0x60, 0xe9,
		0x24, 0x5c, 0x7f, 0xc7, 0xb9, 0x4a, 0x38, 0x4a, 0xb2, 0x0d, 0xcb, 0x94, 0xc4, 0x18, 0xf9, 0x81,
		0x33, 0xa4, 0x87, 0xde, 0x83, 0x69, 0xfe, 0xfe, 0x7d, 0xb6, 0x51, 0xca, 0x04, 0xd6, 0x08, 0x51,
		0xb2, 0x0c, 0xf4, 0x73, 0x4d, 0x0f, 0x1b, 0x56, 0x02, 0xc3, 0x3f, 0x70, 0x43, 0x42, 0x7d, 0xf9,
		0x15, 0x58, 0x23, 0xff, 0xd3, 0x33, 0x29, 0x6a, 0x49, 0xf2, 0x6b, 0x57, 0xe5, 0x7b, 0x6f, 0xb1,
		0xbd, 0xb8, 0x1a, 0x12, 0x44, 0x6c, 0x8a, 0xac, 0xe2, 0x00, 0x07, 0x01, 0xf6, 0x7c, 0x4d, 0xb7,
		0x66, 0x99, 0x17, 0x79, 0x2e, 0xa8, 0x7c, 0xf9, 0x83, 0xf8, 0x2a, 0x6e, 0x33, 0x64, 0xcd, 0xb2,
		0xe4, 0x7d, 0x38, 0x3d, 0x23, 0x2a, 0xe6, 0xe0, 0x7c, 0x9b, 0x73, 0xae, 0x4d, 0x45, 0x06, 0xa1,
		0x6d, 0x83, 0x90, 0x87, 0x6b, 0x39, 0x07, 0xe7, 0x6f, 0x73, 0x4e, 0xc4, 0xb1, 0x62, 0x49, 0x09,
		0xe3, 0x0d, 0x58, 0xb9, 0x83, 0xbd, 0x03, 0xc7, 0xe7, 0x4f, 0x34, 0x73, 0xd0, 0xbd, 0xc3, 0xe9,
		0x96, 0x39, 0x90, 0xbe, 0xd9, 0x10, 0xae, 0x17, 0x21, 0xdf, 0xd7, 0x0d, 0x3c, 0x07, 0xc5, 0x57,
		0x38, 0xc5, 0x22, 0xd1, 0x27, 0xd0, 0x1a, 0x94, 0x06, 0x0e, 0x4f, 0x4b, 0xc9, 0xf0, 0x77, 0x39,
		0xbc, 0x28, 0x30, 0x9c, 0xc2, 0x75, 0xdc, 0x91, 0x45, 0x72, 0x56, 0x32, 0xc5, 0xef, 0x08, 0x0a,
		0x81, 0xe1, 0x14, 0x27, 0x70, 0xeb, 0xef, 0x0a, 0x0a, 0x3f, 0xe2, 0xcf, 0x97, 0xa0, 0xe8, 0xd8,
		0xd6, 0x91, 0x63, 0xcf, 0x63, 0xc4, 0xef, 0x71, 0x06, 0xe0, 0x10, 0x42, 0x70, 0x0d, 0x0a, 0xf3,
		0x2e, 0xc4, 0xd7, 0x3e, 0x10, 0xdb, 0x43, 0xac, 0xc0, 0x36, 0x2c, 0x8b, 0x03, 0xca, 0x74, 0xec,
		0x39, 0x28, 0x7e, 0x9f, 0x53, 0x94, 0x23, 0x30, 0x3e, 0x8d, 0x00, 0xfb, 0xc1, 0x00, 0xcf, 0x43,
		0xf2, 0x9e, 0x98, 0x06, 0x87, 0x70, 0x57, 0x1e, 0x60, 0xdb, 0x38, 0x9c, 0x8f, 0xe1, 0x0f, 0x84,
		0x2b, 0x05, 0x86, 0x50, 0x34, 0x60, 0x69, 0xa8, 0x7b, 0xfe, 0xa1, 0x6e, 0xcd, 0xb5, 0x1c, 0x7f,
		0xc8, 0x39, 0x4a, 0x21, 0x88, 0x7b, 0x64, 0x64, 0x9f, 0x84, 0xe6, 0xeb, 0xc2, 0x23, 0x11, 0x18,
		0xdf, 0x7a, 0x7e, 0x40, 0xdf, 0xb3, 0x4e, 0xc2, 0xf6, 0x0d, 0xb1, 0xf5, 0x18, 0x76, 0x37, 0xca,
		0x78, 0x0d, 0x0a, 0xbe, 0x79, 0x6f, 0x2e, 0x9a, 0x3f, 0x12, 0x2b, 0x4d, 0x01, 0x04, 0xfc, 0x1a,
		0x9c, 0x99, 0x99, 0x26, 0xe6, 0x20, 0xfb, 0x63, 0x4e, 0x76, 0x6a, 0x46, 0xaa, 0xe0, 0x47, 0xc2,
		0x49, 0x29, 0xff, 0x44, 0x1c, 0x09, 0x78, 0x82, 0xab, 0x4d, 0x2e, 0x0a, 0xbe, 0xde, 0x3f, 0x99,
		0xd7, 0xfe, 0x54, 0x78, 0x8d, 0x61, 0x63, 0x5e, 0xeb, 0xc2, 0x29, 0xce, 0x78, 0xb2, 0x75, 0xfd,
		0xa6, 0x38, 0x58, 0x19, 0x7a, 0x3f, 0xbe, 0xba, 0x3f, 0x05, 0xeb, 0xa1, 0x3b, 0x45, 0x45, 0xea,
		0x6b, 0x43, 0xdd, 0x9d, 0x83, 0xf9, 0x5b, 0x9c, 0x59, 0x9c, 0xf8, 0x61, 0x49, 0xeb, 0xef, 0xea,
		0x2e, 0x21, 0xbf, 0x09, 0x15, 0x41, 0x3e, 0xb2, 0x3d, 0x6c, 0x38, 0x03, 0xdb, 0xbc, 0x87, 0x7b,
		0x73, 0x50, 0xff, 0xd9, 0xc4, 0x52, 0xed, 0x47, 0xe0, 0x84, 0xb9, 0x09, 0x52, 0x58, 0xab, 0x68,
		0xe6, 0xd0, 0x75, 0xbc, 0x20, 0x81, 0xf1, 0xdb, 0x62, 0xa5, 0x42, 0x5c, 0x93, 0xc2, 0x64, 0x05,
		0xca, 0xb4, 0x39, 0x6f, 0x48, 0xfe, 0x39, 0x27, 0x5a, 0x1a, 0xa3, 0xf8, 0xc1, 0x61, 0x38, 0x43,
		0x57, 0xf7, 0xe6, 0x39, 0xff, 0xfe, 0x42, 0x1c, 0x1c, 0x1c, 0xc2, 0x0f, 0x8e, 0xe0, 0xc8, 0xc5,
		0x24, 0xdb, 0xcf, 0xc1, 0xf0, 0x1d, 0x71, 0x70, 0x08, 0x0c, 0xa7, 0x10, 0x05, 0xc3, 0x1c, 0x14,
		0x7f, 0x29, 0x28, 0x04, 0x86, 0x50, 0x7c, 0x7e, 0x9c, 0x68, 0x3d, 0x3c, 0x30, 0xfd, 0xc0, 0x63,
		0x75, 0xf0, 0x83, 0xa9, 0xfe, 0xea, 0x83, 0x78, 0x11, 0xa6, 0x46, 0xa0, 0xf2, 0x0d, 0x58, 0x9e,
		0x28, 0x31, 0x50, 0xd2, 0x6f, 0x3f, 0x2a, 0x3f, 0xf3, 0x21, 0x3f, 0x8c, 0xe2, 0x15, 0x86, 0xbc,
		0x43, 0xd6, 0x3d, 0x5e, 0x07, 0x24, 0x93, 0xbd, 0xf5, 0x61, 0xb8, 0xf4, 0xb1, 0x32, 0x40, 0xbe,
		0x0e, 0x4b, 0xb1, 0x1a, 0x20, 0x99, 0xea, 0x67, 0x39, 0x55, 0x29, 0x5a, 0x02, 0xc8, 0x97, 0x21,
		0x4b, 0xf2, 0x79, 0x32, 0xfc, 0xe7, 0x38, 0x9c, 0xaa, 0xcb, 0x9f, 0x81, 0xbc, 0xc8, 0xe3, 0xc9,
		0xd0, 0x9f, 0xe7, 0xd0, 0x10, 0x42, 0xe0, 0x22, 0x87, 0x27, 0xc3, 0x7f, 0x41, 0xc0, 0x05, 0x84,
		0xc0, 0xe7, 0x77, 0xe1, 0xdf, 0xfe, 0x62, 0x96, 0x9f, 0xc3, 0xc2, 0x77, 0xd7, 0x60, 0x91, 0x27,
		0xef, 0x64, 0xf4, 0x17, 0xf9, 0xe0, 0x02, 0x21, 0xbf, 0x00, 0xb9, 0x39, 0x1d, 0xfe, 0x4b, 0x1c,
		0xca, 0xf4, 0xe5, 0x06, 0x14, 0x23, 0x09, 0x3b, 0x19, 0xfe, 0xcb, 0x1c, 0x1e, 0x45, 0x11, 0xd3,
		0x79, 0xc2, 0x4e, 0x26, 0xf8, 0x15, 0x61, 0x3a, 0x47, 0x10, 0xb7, 0x89, 0x5c, 0x9d, 0x8c, 0xfe,
		0x55, 0xe1, 0x75, 0x01, 0x91, 0x5f, 0x82, 0x42, 0x78, 0xfe, 0x26, 0xe3, 0x7f, 0x8d, 0xe3, 0xc7,
		0x18, 0xe2, 0x81, 0xc8, 0xf9, 0x9f, 0x4c, 0xf1, 0xeb, 0xc2, 0x03, 0x11, 0x14, 0xd9, 0x46, 0x93,
		0x39, 0x3d, 0x99, 0xe9, 0x37, 0xc4, 0x36, 0x9a, 0x48, 0xe9, 0x64, 0x35, 0xe9, 0x31, 0x98, 0x4c,
		0xf1, 0x9b, 0x62, 0x35, 0xa9, 0x3e, 0x31, 0x63, 0x32, 0x49, 0x26, 0x73, 0xfc, 0x96, 0x30, 0x63,
		0x22, 0x47, 0xca, 0x6d, 0x40, 0xd3, 0x09, 0x32, 0x99, 0xef, 0x4b, 0x9c, 0x6f, 0x65, 0x2a, 0x3f,
		0xca, 0xaf, 0xc2, 0xa9, 0xd9, 0xc9, 0x31, 0x99, 0xf5, 0xcb, 0x1f, 0x4e, 0x5c, 0x67, 0xa2, 0xb9,
		0x51, 0xee, 0x8e, 0x4f, 0xd9, 0x68, 0x62, 0x4c, 0xa6, 0x7d, 0xfb, 0xc3, 0xf8, 0x41, 0x1b, 0xcd,
		0x8b, 0x72, 0x0d, 0x60, 0x9c, 0x93, 0x92, 0xb9, 0xde, 0xe1, 0x5c, 0x11, 0x10, 0xd9, 0x1a, 0x3c,
		0x25, 0x25, 0xe3, 0xbf, 0x22, 0xb6, 0x06, 0x47, 0x90, 0xad, 0x21, 0xb2, 0x51, 0x32, 0xfa, 0x5d,
		0xb1, 0x35, 0x04, 0x44, 0xbe, 0x06, 0x79, 0x7b, 0x64, 0x59, 0x24, 0xb6, 0xd0, 0x83, 0x7f, 0xce,
		0x54, 0xf9, 0x8f, 0xfb, 0x1c, 0x2c, 0x00, 0xf2, 0x65, 0xc8, 0xe1, 0xe1, 0x01, 0xee, 0x25, 0x21,
		0xff, 0xf3, 0xbe, 0x38, 0x4f, 0x88, 0xb6, 0xfc, 0x12, 0x00, 0xbb, 0x4c, 0xd3, 0xaf, 0x44, 0x09,
		0xd8, 0xff, 0xba, 0xcf, 0x7f, 0x29, 0x31, 0x86, 0x8c, 0x09, 0xd8, 0xef, 0x2e, 0x1e, 0x4c, 0xf0,
		0x41, 0x9c, 0x80, 0x5e, 0xc0, 0x5f, 0x84, 0xc5, 0x5b, 0xbe, 0x63, 0x07, 0xfa, 0x20, 0x09, 0xfd,
		0xdf, 0x1c, 0x2d, 0xf4, 0x89, 0xc3, 0x86, 0x8e, 0x87, 0x03, 0x7d, 0xe0, 0x27, 0x61, 0xff, 0x87,
		0x63, 0x43, 0x00, 0x01, 0x1b, 0xba, 0x1f, 0xcc, 0x33, 0xef, 0xff, 0x15, 0x60, 0x01, 0x20, 0x46,
		0x93, 0xff, 0x6f, 0xe3, 0xa3, 0x24, 0xec, 0xf7, 0x85, 0xd1, 0x5c, 0x5f, 0xfe, 0x0c, 0x14, 0xc8,
		0xbf, 0xec, 0xd7, 0x43, 0x09, 0xe0, 0xff, 0xe3, 0xe0, 0x31, 0x82, 0x8c, 0xec, 0x07, 0xbd, 0xc0,
		0x4c, 0x76, 0xf6, 0xff, 0xf3, 0x95, 0x16, 0xfa, 0x72, 0x0d, 0x8a, 0x7e, 0xd0, 0xeb, 0x8d, 0x78,
		0x45, 0x93, 0x00, 0xff, 0xc1, 0xfd, 0xf0, 0x92, 0x1b, 0x62, 0xea, 0xe7, 0x67, 0x3f, 0xd6, 0xc1,
		0xb6, 0xb3, 0xed, 0xb0, 0x67, 0x3a, 0xf8, 0x5a, 0x96, 0xd7, 0x93, 0xf6, 0x68, 0xc8, 0x9f, 0xd5,
		0xf2, 0xa2, 0xbd, 0x7e, 0xb2, 0xf7, 0xb8, 0xaa, 0x0c, 0x79, 0xca, 0xba, 0x37, 0x1a, 0xa2, 0x4d,
		0xc8, 0xbf, 0x8c, 0x8f, 0x5e, 0x89, 0xfc, 0x58, 0x0b, 0x6d, 0x86, 0xe3, 0x88, 0x1e, 0x35, 0xd4,
		0xa9, 0x5e, 0x1a, 0xeb, 0x23, 0x09, 0x32, 0x2f, 0xe3, 0x23, 0xfe, 0xbb, 0x18, 0xf2, 0x2f, 0x5a,
		0x83, 0x9c, 0xa0, 0x22, 0x32, 0xd6, 0xa8, 0xbe, 0x0e, 0xd0, 0x75, 0x5c, 0xff, 0xa8, 0x3b, 0xf2,
		0xee, 0x1c, 0xa1, 0x27, 0x21, 0x47, 0x5b, 0x14, 0x57, 0xbc, 0xb4, 0x3c, 0x1e, 0x8e, 0x8a, 0x55,
		0xd6, 0x4b, 0xd5, 0x88, 0x3e, 0xff, 0x95, 0x5a, 0x54, 0x8d, 0x88, 0x55, 0xd6, 0x5b, 0x7d, 0x9c,
		0xb3, 0xa1, 0x12, 0xa4, 0x6a, 0x94, 0x32, 0xa3, 0xa6, 0x6a, 0xa4, 0x55, 0xa7, 0xc8, 0x8c, 0x9a,
		0xaa, 0x53, 0x25, 0x3a, 0x36, 0x15, 0x73, 0xa5, 0x3a, 0x83, 0x70, 0xa5, 0x5a, 0x75, 0x04, 0xd9,
		0x97, 0x6d, 0x27, 0x40, 0x17, 0x20, 0x57, 0x37, 0x07, 0x87, 0x01, 0x7f, 0xf2, 0x5d, 0x1d, 0x0f,
		0x4c, 0xc5, 0x44, 0x47, 0x65, 0x1a, 0x64, 0xba, 0x8a, 0x75, 0xe0, 0xdc, 0xe5, 0x5f, 0x12, 0x59,
		0x03, 0x7d, 0x12, 0x0a, 0x75, 0x93, 0xd4, 0x95, 0x8a, 0xdd, 0xe3, 0x3f, 0xc8, 0x2a, 0x47, 0x7c,
		0x4a, 0xf0, 0x63, 0x85, 0x6a, 0x83, 0x68, 0x73, 0x5e, 0x54, 0x85, 0xec, 0x8e, 0xe3, 0xb8, 0xdc,
		0x35, 0x93, 0x28, 0xda, 0x87, 0x10, 0x64, 0xbb, 0x23, 0xcf, 0xe6, 0x63, 0xd2, 0xff, 0xeb, 0xf9,
		0xef, 0xff, 0xcb, 0x63, 0xa9, 0x6f, 0xfe, 0xeb, 0x63, 0xa9, 0x1f, 0x06, 0x00, 0x00, 0xff, 0xff,
		0x87, 0x5a, 0x9b, 0x62, 0xcc, 0x31, 0x00, 0x00,
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
func (this *ProtoNum) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&protonum.ProtoNum{")
	if this.KeyValue != nil {
		s = append(s, "KeyValue: "+fmt.Sprintf("%#v", this.KeyValue)+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *KeyValue) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&protonum.KeyValue{")
	if this.Key != nil {
		s = append(s, "Key: "+valueToGoStringProtonum(this.Key, "string")+",\n")
	}
	if this.Value != nil {
		s = append(s, "Value: "+valueToGoStringProtonum(this.Value, "string")+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *TopsyTurvy) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&protonum.TopsyTurvy{")
	if this.Topsy != nil {
		s = append(s, "Topsy: "+fmt.Sprintf("%#v", this.Topsy)+",\n")
	}
	if this.Turvy != nil {
		s = append(s, "Turvy: "+fmt.Sprintf("%#v", this.Turvy)+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Topsy) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&protonum.Topsy{")
	if this.A != nil {
		s = append(s, "A: "+valueToGoStringProtonum(this.A, "int64")+",\n")
	}
	if this.B != nil {
		s = append(s, "B: "+valueToGoStringProtonum(this.B, "int64")+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Turvy) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&protonum.Turvy{")
	if this.B != nil {
		s = append(s, "B: "+valueToGoStringProtonum(this.B, "int64")+",\n")
	}
	if this.A != nil {
		s = append(s, "A: "+valueToGoStringProtonum(this.A, "int64")+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Knot) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&protonum.Knot{")
	if this.Bight != nil {
		s = append(s, "Bight: "+fmt.Sprintf("%#v", this.Bight)+",\n")
	}
	if this.Elbow != nil {
		s = append(s, "Elbow: "+valueToGoStringProtonum(this.Elbow, "bool")+",\n")
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
	s = append(s, "&protonum.BightKnot{")
	if this.Loop != nil {
		s = append(s, "Loop: "+fmt.Sprintf("%#v", this.Loop)+",\n")
	}
	if this.Turn != nil {
		s = append(s, "Turn: "+valueToGoStringProtonum(this.Turn, "bool")+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringProtonum(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}

func init() { proto.RegisterFile("protonum.proto", fileDescriptorProtonum) }

var fileDescriptorProtonum = []byte{
	// 317 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x59, 0xd3, 0x42, 0x3a, 0x95, 0x2a, 0xa3, 0x87, 0xe0, 0xa1, 0x94, 0x15, 0xa1, 0x82,
	0xa6, 0xd0, 0xa3, 0xb7, 0xae, 0xf4, 0x54, 0x11, 0x59, 0x8a, 0x07, 0x6f, 0x46, 0x63, 0x1a, 0x68,
	0xb2, 0x21, 0xee, 0x2a, 0x79, 0x2b, 0x5f, 0x49, 0x9f, 0xc0, 0x47, 0x90, 0x9d, 0xdd, 0x34, 0xea,
	0x6d, 0xfe, 0x99, 0xef, 0xff, 0xff, 0x64, 0x61, 0x54, 0xd5, 0x4a, 0xab, 0xd2, 0x14, 0x31, 0x0d,
	0x18, 0xb6, 0xfa, 0xe4, 0x32, 0xcb, 0xf5, 0xc6, 0x24, 0xf1, 0x93, 0x2a, 0x66, 0x99, 0xca, 0xd4,
	0x8c, 0x2e, 0x89, 0x79, 0x21, 0x45, 0x82, 0x26, 0x67, 0xe4, 0x57, 0x10, 0xde, 0xd9, 0xe1, 0xd6,
	0x14, 0x18, 0x43, 0xb8, 0x4a, 0x9b, 0xfb, 0xc7, 0xad, 0x49, 0xa3, 0xbd, 0x49, 0x30, 0x1d, 0xce,
	0x31, 0xde, 0xf5, 0xb4, 0x17, 0xb9, 0x63, 0xf8, 0xbc, 0xe3, 0xf1, 0x10, 0x82, 0x55, 0xda, 0x44,
	0x6c, 0xc2, 0xa6, 0x03, 0x69, 0x47, 0x3c, 0x86, 0x7e, 0x1b, 0x65, 0x77, 0x4e, 0xf0, 0x07, 0x80,
	0xb5, 0xaa, 0x5e, 0x9b, 0xb5, 0xa9, 0xdf, 0x1a, 0x3c, 0x83, 0x3e, 0x29, 0xf2, 0x0d, 0xe7, 0x07,
	0x5d, 0x1d, 0xad, 0xa5, 0xbb, 0x12, 0x66, 0x79, 0x8a, 0xfa, 0x8b, 0xd9, 0xb5, 0x74, 0x57, 0x7e,
	0xea, 0xd3, 0x70, 0x1f, 0xd8, 0x82, 0x22, 0x03, 0xc9, 0x16, 0x56, 0x09, 0x72, 0x06, 0x92, 0x09,
	0x82, 0xa8, 0x9b, 0xd6, 0x1e, 0x12, 0xce, 0xe2, 0xa1, 0x05, 0x37, 0xd0, 0x5b, 0x95, 0x4a, 0xe3,
	0x39, 0xf4, 0x45, 0x9e, 0x6d, 0x74, 0xc4, 0xe8, 0x39, 0x8e, 0xba, 0x62, 0x5a, 0x5b, 0x46, 0x3a,
	0xc2, 0xfe, 0xee, 0x72, 0x9b, 0xa8, 0x77, 0x0a, 0x09, 0xa5, 0x13, 0x78, 0x01, 0x03, 0x91, 0x6b,
	0x9d, 0xd6, 0xcb, 0xf2, 0x39, 0x0a, 0xe8, 0xeb, 0x47, 0xbf, 0xde, 0xd4, 0xfa, 0x3b, 0x80, 0x5f,
	0x5b, 0xda, 0xe7, 0x22, 0x87, 0xde, 0x8d, 0x52, 0x95, 0x7f, 0x9a, 0xff, 0x2e, 0xba, 0x21, 0x42,
	0x6f, 0x6d, 0xea, 0xd2, 0x77, 0xd2, 0x2c, 0xc2, 0xef, 0xcf, 0x31, 0xfb, 0xf8, 0x1a, 0xb3, 0x9f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x5a, 0xec, 0x74, 0x43, 0x25, 0x02, 0x00, 0x00,
}
