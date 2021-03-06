// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kythe/proto/java.proto

package java_go_proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import analysis_go_proto "kythe.io/kythe/proto/analysis_go_proto"
import storage_go_proto "kythe.io/kythe/proto/storage_go_proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type JarDetails struct {
	Jar                  []*JarDetails_Jar `protobuf:"bytes,1,rep,name=jar" json:"jar,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *JarDetails) Reset()         { *m = JarDetails{} }
func (m *JarDetails) String() string { return proto.CompactTextString(m) }
func (*JarDetails) ProtoMessage()    {}
func (*JarDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_java_1eb7f50e2cd49c5d, []int{0}
}
func (m *JarDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JarDetails.Unmarshal(m, b)
}
func (m *JarDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JarDetails.Marshal(b, m, deterministic)
}
func (dst *JarDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JarDetails.Merge(dst, src)
}
func (m *JarDetails) XXX_Size() int {
	return xxx_messageInfo_JarDetails.Size(m)
}
func (m *JarDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_JarDetails.DiscardUnknown(m)
}

var xxx_messageInfo_JarDetails proto.InternalMessageInfo

func (m *JarDetails) GetJar() []*JarDetails_Jar {
	if m != nil {
		return m.Jar
	}
	return nil
}

type JarDetails_Jar struct {
	VName                *storage_go_proto.VName       `protobuf:"bytes,1,opt,name=v_name,json=vName" json:"v_name,omitempty"`
	Entry                []*analysis_go_proto.FileInfo `protobuf:"bytes,2,rep,name=entry" json:"entry,omitempty"` // Deprecated: Do not use.
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *JarDetails_Jar) Reset()         { *m = JarDetails_Jar{} }
func (m *JarDetails_Jar) String() string { return proto.CompactTextString(m) }
func (*JarDetails_Jar) ProtoMessage()    {}
func (*JarDetails_Jar) Descriptor() ([]byte, []int) {
	return fileDescriptor_java_1eb7f50e2cd49c5d, []int{0, 0}
}
func (m *JarDetails_Jar) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JarDetails_Jar.Unmarshal(m, b)
}
func (m *JarDetails_Jar) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JarDetails_Jar.Marshal(b, m, deterministic)
}
func (dst *JarDetails_Jar) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JarDetails_Jar.Merge(dst, src)
}
func (m *JarDetails_Jar) XXX_Size() int {
	return xxx_messageInfo_JarDetails_Jar.Size(m)
}
func (m *JarDetails_Jar) XXX_DiscardUnknown() {
	xxx_messageInfo_JarDetails_Jar.DiscardUnknown(m)
}

var xxx_messageInfo_JarDetails_Jar proto.InternalMessageInfo

func (m *JarDetails_Jar) GetVName() *storage_go_proto.VName {
	if m != nil {
		return m.VName
	}
	return nil
}

// Deprecated: Do not use.
func (m *JarDetails_Jar) GetEntry() []*analysis_go_proto.FileInfo {
	if m != nil {
		return m.Entry
	}
	return nil
}

type JarEntryDetails struct {
	JarContainer         int32    `protobuf:"varint,1,opt,name=jar_container,json=jarContainer" json:"jar_container,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JarEntryDetails) Reset()         { *m = JarEntryDetails{} }
func (m *JarEntryDetails) String() string { return proto.CompactTextString(m) }
func (*JarEntryDetails) ProtoMessage()    {}
func (*JarEntryDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_java_1eb7f50e2cd49c5d, []int{1}
}
func (m *JarEntryDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JarEntryDetails.Unmarshal(m, b)
}
func (m *JarEntryDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JarEntryDetails.Marshal(b, m, deterministic)
}
func (dst *JarEntryDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JarEntryDetails.Merge(dst, src)
}
func (m *JarEntryDetails) XXX_Size() int {
	return xxx_messageInfo_JarEntryDetails.Size(m)
}
func (m *JarEntryDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_JarEntryDetails.DiscardUnknown(m)
}

var xxx_messageInfo_JarEntryDetails proto.InternalMessageInfo

func (m *JarEntryDetails) GetJarContainer() int32 {
	if m != nil {
		return m.JarContainer
	}
	return 0
}

type JavaDetails struct {
	Classpath            []string `protobuf:"bytes,1,rep,name=classpath" json:"classpath,omitempty"`
	Sourcepath           []string `protobuf:"bytes,2,rep,name=sourcepath" json:"sourcepath,omitempty"`
	Bootclasspath        []string `protobuf:"bytes,3,rep,name=bootclasspath" json:"bootclasspath,omitempty"`
	ExtraJavacopts       []string `protobuf:"bytes,10,rep,name=extra_javacopts,json=extraJavacopts" json:"extra_javacopts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JavaDetails) Reset()         { *m = JavaDetails{} }
func (m *JavaDetails) String() string { return proto.CompactTextString(m) }
func (*JavaDetails) ProtoMessage()    {}
func (*JavaDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_java_1eb7f50e2cd49c5d, []int{2}
}
func (m *JavaDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JavaDetails.Unmarshal(m, b)
}
func (m *JavaDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JavaDetails.Marshal(b, m, deterministic)
}
func (dst *JavaDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JavaDetails.Merge(dst, src)
}
func (m *JavaDetails) XXX_Size() int {
	return xxx_messageInfo_JavaDetails.Size(m)
}
func (m *JavaDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_JavaDetails.DiscardUnknown(m)
}

var xxx_messageInfo_JavaDetails proto.InternalMessageInfo

func (m *JavaDetails) GetClasspath() []string {
	if m != nil {
		return m.Classpath
	}
	return nil
}

func (m *JavaDetails) GetSourcepath() []string {
	if m != nil {
		return m.Sourcepath
	}
	return nil
}

func (m *JavaDetails) GetBootclasspath() []string {
	if m != nil {
		return m.Bootclasspath
	}
	return nil
}

func (m *JavaDetails) GetExtraJavacopts() []string {
	if m != nil {
		return m.ExtraJavacopts
	}
	return nil
}

func init() {
	proto.RegisterType((*JarDetails)(nil), "kythe.proto.JarDetails")
	proto.RegisterType((*JarDetails_Jar)(nil), "kythe.proto.JarDetails.Jar")
	proto.RegisterType((*JarEntryDetails)(nil), "kythe.proto.JarEntryDetails")
	proto.RegisterType((*JavaDetails)(nil), "kythe.proto.JavaDetails")
}

func init() { proto.RegisterFile("kythe/proto/java.proto", fileDescriptor_java_1eb7f50e2cd49c5d) }

var fileDescriptor_java_1eb7f50e2cd49c5d = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xcf, 0x4e, 0xc2, 0x40,
	0x10, 0x87, 0x53, 0x1a, 0x48, 0x98, 0x8a, 0x24, 0x9b, 0x68, 0x2a, 0x1a, 0x45, 0x34, 0x11, 0x0f,
	0x96, 0x04, 0x13, 0x1f, 0x00, 0xff, 0x24, 0xf6, 0xe0, 0xa1, 0x07, 0xaf, 0x64, 0xa8, 0x2b, 0xb4,
	0x2e, 0x3b, 0x64, 0x76, 0x6d, 0xe4, 0x55, 0x3c, 0xf8, 0xac, 0xa6, 0x0b, 0x0d, 0xe5, 0x34, 0xd3,
	0xef, 0xf7, 0x4d, 0xa7, 0xdd, 0x85, 0xe3, 0xaf, 0xb5, 0x5d, 0xc8, 0xd1, 0x8a, 0xc9, 0xd2, 0x28,
	0xc7, 0x02, 0x23, 0xd7, 0x8a, 0xc0, 0xf1, 0xcd, 0x43, 0xaf, 0x57, 0x97, 0x50, 0xa3, 0x5a, 0x9b,
	0xcc, 0x6c, 0xb3, 0x93, 0x7a, 0x66, 0x2c, 0x31, 0xce, 0xb7, 0x63, 0x83, 0x3f, 0x0f, 0x20, 0x46,
	0x7e, 0x92, 0x16, 0x33, 0x65, 0xc4, 0x1d, 0xf8, 0x39, 0x72, 0xe8, 0xf5, 0xfd, 0x61, 0x30, 0x3e,
	0x8d, 0x6a, 0x0b, 0xa2, 0x9d, 0x55, 0xb6, 0x49, 0xe9, 0xf5, 0x10, 0xfc, 0x18, 0x59, 0xdc, 0x42,
	0xab, 0x98, 0x6a, 0x5c, 0xca, 0xd0, 0xeb, 0x7b, 0xc3, 0x60, 0x2c, 0xf6, 0x06, 0xdf, 0xdf, 0x70,
	0x29, 0x93, 0x66, 0x51, 0x16, 0x31, 0x82, 0xa6, 0xd4, 0x96, 0xd7, 0x61, 0xc3, 0xad, 0x38, 0xda,
	0x33, 0x5f, 0x32, 0x25, 0x5f, 0xf5, 0x27, 0x4d, 0x1a, 0xa1, 0x97, 0x6c, 0xbc, 0xc1, 0x03, 0x74,
	0x63, 0xe4, 0xe7, 0xb2, 0xaf, 0x3e, 0xf2, 0x0a, 0x3a, 0x39, 0xf2, 0x34, 0x25, 0x6d, 0x31, 0xd3,
	0x92, 0xdd, 0xd6, 0x66, 0x72, 0x90, 0x23, 0x3f, 0x56, 0x6c, 0xf0, 0xeb, 0x41, 0x10, 0x63, 0x81,
	0xd5, 0xd0, 0x19, 0xb4, 0x53, 0x85, 0xc6, 0xac, 0xd0, 0x2e, 0xdc, 0xff, 0xb5, 0x93, 0x1d, 0x10,
	0xe7, 0x00, 0x86, 0xbe, 0x39, 0x95, 0x2e, 0x6e, 0xb8, 0xb8, 0x46, 0xc4, 0x35, 0x74, 0x66, 0x44,
	0x76, 0xf7, 0x06, 0xdf, 0x29, 0xfb, 0x50, 0xdc, 0x40, 0x57, 0xfe, 0x58, 0xc6, 0x69, 0x79, 0x49,
	0x29, 0xad, 0xac, 0x09, 0xc1, 0x79, 0x87, 0x0e, 0xc7, 0x15, 0x9d, 0x5c, 0xc2, 0x45, 0x4a, 0xcb,
	0x68, 0x4e, 0x34, 0x57, 0x32, 0xfa, 0x90, 0x85, 0x25, 0x52, 0xa6, 0x7e, 0x16, 0xb3, 0x96, 0x2b,
	0xf7, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc4, 0x3c, 0xd5, 0xe6, 0xfd, 0x01, 0x00, 0x00,
}
