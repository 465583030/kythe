// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kythe/proto/extraction_config.proto

package extraction_config_go_proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ExtractionConfiguration struct {
	RequiredImage        []*ExtractionConfiguration_Image      `protobuf:"bytes,1,rep,name=required_image,json=requiredImage" json:"required_image,omitempty"`
	RunCommand           []*ExtractionConfiguration_RunCommand `protobuf:"bytes,2,rep,name=run_command,json=runCommand" json:"run_command,omitempty"`
	EntryPoint           []string                              `protobuf:"bytes,3,rep,name=entry_point,json=entryPoint" json:"entry_point,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *ExtractionConfiguration) Reset()         { *m = ExtractionConfiguration{} }
func (m *ExtractionConfiguration) String() string { return proto.CompactTextString(m) }
func (*ExtractionConfiguration) ProtoMessage()    {}
func (*ExtractionConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_extraction_config_e419371b3fd43c0a, []int{0}
}
func (m *ExtractionConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtractionConfiguration.Unmarshal(m, b)
}
func (m *ExtractionConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtractionConfiguration.Marshal(b, m, deterministic)
}
func (dst *ExtractionConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtractionConfiguration.Merge(dst, src)
}
func (m *ExtractionConfiguration) XXX_Size() int {
	return xxx_messageInfo_ExtractionConfiguration.Size(m)
}
func (m *ExtractionConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtractionConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_ExtractionConfiguration proto.InternalMessageInfo

func (m *ExtractionConfiguration) GetRequiredImage() []*ExtractionConfiguration_Image {
	if m != nil {
		return m.RequiredImage
	}
	return nil
}

func (m *ExtractionConfiguration) GetRunCommand() []*ExtractionConfiguration_RunCommand {
	if m != nil {
		return m.RunCommand
	}
	return nil
}

func (m *ExtractionConfiguration) GetEntryPoint() []string {
	if m != nil {
		return m.EntryPoint
	}
	return nil
}

type ExtractionConfiguration_Image struct {
	Uri                  string                              `protobuf:"bytes,1,opt,name=uri" json:"uri,omitempty"`
	Name                 string                              `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	CopySpec             []*ExtractionConfiguration_CopySpec `protobuf:"bytes,3,rep,name=copy_spec,json=copySpec" json:"copy_spec,omitempty"`
	EnvVar               []*ExtractionConfiguration_EnvVar   `protobuf:"bytes,4,rep,name=env_var,json=envVar" json:"env_var,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *ExtractionConfiguration_Image) Reset()         { *m = ExtractionConfiguration_Image{} }
func (m *ExtractionConfiguration_Image) String() string { return proto.CompactTextString(m) }
func (*ExtractionConfiguration_Image) ProtoMessage()    {}
func (*ExtractionConfiguration_Image) Descriptor() ([]byte, []int) {
	return fileDescriptor_extraction_config_e419371b3fd43c0a, []int{0, 0}
}
func (m *ExtractionConfiguration_Image) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtractionConfiguration_Image.Unmarshal(m, b)
}
func (m *ExtractionConfiguration_Image) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtractionConfiguration_Image.Marshal(b, m, deterministic)
}
func (dst *ExtractionConfiguration_Image) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtractionConfiguration_Image.Merge(dst, src)
}
func (m *ExtractionConfiguration_Image) XXX_Size() int {
	return xxx_messageInfo_ExtractionConfiguration_Image.Size(m)
}
func (m *ExtractionConfiguration_Image) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtractionConfiguration_Image.DiscardUnknown(m)
}

var xxx_messageInfo_ExtractionConfiguration_Image proto.InternalMessageInfo

func (m *ExtractionConfiguration_Image) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *ExtractionConfiguration_Image) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ExtractionConfiguration_Image) GetCopySpec() []*ExtractionConfiguration_CopySpec {
	if m != nil {
		return m.CopySpec
	}
	return nil
}

func (m *ExtractionConfiguration_Image) GetEnvVar() []*ExtractionConfiguration_EnvVar {
	if m != nil {
		return m.EnvVar
	}
	return nil
}

type ExtractionConfiguration_CopySpec struct {
	Source               string   `protobuf:"bytes,1,opt,name=source" json:"source,omitempty"`
	Destination          string   `protobuf:"bytes,2,opt,name=destination" json:"destination,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExtractionConfiguration_CopySpec) Reset()         { *m = ExtractionConfiguration_CopySpec{} }
func (m *ExtractionConfiguration_CopySpec) String() string { return proto.CompactTextString(m) }
func (*ExtractionConfiguration_CopySpec) ProtoMessage()    {}
func (*ExtractionConfiguration_CopySpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_extraction_config_e419371b3fd43c0a, []int{0, 1}
}
func (m *ExtractionConfiguration_CopySpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtractionConfiguration_CopySpec.Unmarshal(m, b)
}
func (m *ExtractionConfiguration_CopySpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtractionConfiguration_CopySpec.Marshal(b, m, deterministic)
}
func (dst *ExtractionConfiguration_CopySpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtractionConfiguration_CopySpec.Merge(dst, src)
}
func (m *ExtractionConfiguration_CopySpec) XXX_Size() int {
	return xxx_messageInfo_ExtractionConfiguration_CopySpec.Size(m)
}
func (m *ExtractionConfiguration_CopySpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtractionConfiguration_CopySpec.DiscardUnknown(m)
}

var xxx_messageInfo_ExtractionConfiguration_CopySpec proto.InternalMessageInfo

func (m *ExtractionConfiguration_CopySpec) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *ExtractionConfiguration_CopySpec) GetDestination() string {
	if m != nil {
		return m.Destination
	}
	return ""
}

type ExtractionConfiguration_EnvVar struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExtractionConfiguration_EnvVar) Reset()         { *m = ExtractionConfiguration_EnvVar{} }
func (m *ExtractionConfiguration_EnvVar) String() string { return proto.CompactTextString(m) }
func (*ExtractionConfiguration_EnvVar) ProtoMessage()    {}
func (*ExtractionConfiguration_EnvVar) Descriptor() ([]byte, []int) {
	return fileDescriptor_extraction_config_e419371b3fd43c0a, []int{0, 2}
}
func (m *ExtractionConfiguration_EnvVar) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtractionConfiguration_EnvVar.Unmarshal(m, b)
}
func (m *ExtractionConfiguration_EnvVar) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtractionConfiguration_EnvVar.Marshal(b, m, deterministic)
}
func (dst *ExtractionConfiguration_EnvVar) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtractionConfiguration_EnvVar.Merge(dst, src)
}
func (m *ExtractionConfiguration_EnvVar) XXX_Size() int {
	return xxx_messageInfo_ExtractionConfiguration_EnvVar.Size(m)
}
func (m *ExtractionConfiguration_EnvVar) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtractionConfiguration_EnvVar.DiscardUnknown(m)
}

var xxx_messageInfo_ExtractionConfiguration_EnvVar proto.InternalMessageInfo

func (m *ExtractionConfiguration_EnvVar) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ExtractionConfiguration_EnvVar) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type ExtractionConfiguration_RunCommand struct {
	Command              string   `protobuf:"bytes,1,opt,name=command" json:"command,omitempty"`
	Arg                  []string `protobuf:"bytes,2,rep,name=arg" json:"arg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExtractionConfiguration_RunCommand) Reset()         { *m = ExtractionConfiguration_RunCommand{} }
func (m *ExtractionConfiguration_RunCommand) String() string { return proto.CompactTextString(m) }
func (*ExtractionConfiguration_RunCommand) ProtoMessage()    {}
func (*ExtractionConfiguration_RunCommand) Descriptor() ([]byte, []int) {
	return fileDescriptor_extraction_config_e419371b3fd43c0a, []int{0, 3}
}
func (m *ExtractionConfiguration_RunCommand) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtractionConfiguration_RunCommand.Unmarshal(m, b)
}
func (m *ExtractionConfiguration_RunCommand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtractionConfiguration_RunCommand.Marshal(b, m, deterministic)
}
func (dst *ExtractionConfiguration_RunCommand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtractionConfiguration_RunCommand.Merge(dst, src)
}
func (m *ExtractionConfiguration_RunCommand) XXX_Size() int {
	return xxx_messageInfo_ExtractionConfiguration_RunCommand.Size(m)
}
func (m *ExtractionConfiguration_RunCommand) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtractionConfiguration_RunCommand.DiscardUnknown(m)
}

var xxx_messageInfo_ExtractionConfiguration_RunCommand proto.InternalMessageInfo

func (m *ExtractionConfiguration_RunCommand) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *ExtractionConfiguration_RunCommand) GetArg() []string {
	if m != nil {
		return m.Arg
	}
	return nil
}

func init() {
	proto.RegisterType((*ExtractionConfiguration)(nil), "kythe.proto.ExtractionConfiguration")
	proto.RegisterType((*ExtractionConfiguration_Image)(nil), "kythe.proto.ExtractionConfiguration.Image")
	proto.RegisterType((*ExtractionConfiguration_CopySpec)(nil), "kythe.proto.ExtractionConfiguration.CopySpec")
	proto.RegisterType((*ExtractionConfiguration_EnvVar)(nil), "kythe.proto.ExtractionConfiguration.EnvVar")
	proto.RegisterType((*ExtractionConfiguration_RunCommand)(nil), "kythe.proto.ExtractionConfiguration.RunCommand")
}

func init() {
	proto.RegisterFile("kythe/proto/extraction_config.proto", fileDescriptor_extraction_config_e419371b3fd43c0a)
}

var fileDescriptor_extraction_config_e419371b3fd43c0a = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x41, 0x4f, 0xe3, 0x30,
	0x10, 0x85, 0x95, 0xa6, 0x4d, 0x9b, 0x89, 0x76, 0xb5, 0xb2, 0x56, 0xbb, 0x51, 0x2e, 0x44, 0x70,
	0xa9, 0x40, 0xa4, 0x52, 0xb9, 0x70, 0x6f, 0x7b, 0x80, 0x53, 0x31, 0x12, 0xd7, 0xc8, 0xa4, 0xa6,
	0x58, 0x10, 0x3b, 0xb8, 0x76, 0x44, 0x7e, 0x1d, 0xbf, 0x82, 0xff, 0x83, 0x3c, 0x49, 0x4a, 0x2f,
	0x48, 0x3d, 0xe5, 0xbd, 0x27, 0xcf, 0x17, 0xbf, 0x31, 0x9c, 0xbd, 0x34, 0xe6, 0x99, 0xcf, 0x2a,
	0xad, 0x8c, 0x9a, 0xf1, 0x77, 0xa3, 0x59, 0x61, 0x84, 0x92, 0x79, 0xa1, 0xe4, 0x93, 0xd8, 0x66,
	0x98, 0x93, 0x08, 0x0f, 0xb5, 0xe6, 0xf4, 0x73, 0x08, 0xff, 0x57, 0xfb, 0x83, 0x0b, 0x3c, 0x67,
	0x35, 0x73, 0x86, 0xdc, 0xc1, 0x6f, 0xcd, 0xdf, 0xac, 0xd0, 0x7c, 0x93, 0x8b, 0x92, 0x6d, 0x79,
	0xec, 0xa5, 0xfe, 0x34, 0x9a, 0x9f, 0x67, 0x07, 0x84, 0xec, 0x87, 0xe9, 0xec, 0xc6, 0x4d, 0xd0,
	0x5f, 0x3d, 0x01, 0x2d, 0x59, 0x43, 0xa4, 0xad, 0xbb, 0x4f, 0x59, 0x32, 0xb9, 0x89, 0x07, 0xc8,
	0x9b, 0x1d, 0xc5, 0xa3, 0x56, 0x2e, 0xda, 0x31, 0x0a, 0x7a, 0xaf, 0xc9, 0x09, 0x44, 0x5c, 0x1a,
	0xdd, 0xe4, 0x95, 0x12, 0xd2, 0xc4, 0x7e, 0xea, 0x4f, 0x43, 0x0a, 0x18, 0xad, 0x5d, 0x92, 0x7c,
	0x78, 0x30, 0x6a, 0x7f, 0xfe, 0x07, 0x7c, 0xab, 0x45, 0xec, 0xa5, 0xde, 0x34, 0xa4, 0x4e, 0x12,
	0x02, 0x43, 0xc9, 0x4a, 0x1e, 0x0f, 0x30, 0x42, 0x4d, 0x6e, 0x21, 0x2c, 0x54, 0xd5, 0xe4, 0xbb,
	0x8a, 0x17, 0x88, 0x8b, 0xe6, 0x97, 0x47, 0x5d, 0x70, 0xa1, 0xaa, 0xe6, 0xbe, 0xe2, 0x05, 0x9d,
	0x14, 0x9d, 0x22, 0x4b, 0x18, 0x73, 0x59, 0xe7, 0x35, 0xd3, 0xf1, 0x10, 0x49, 0x17, 0x47, 0x91,
	0x56, 0xb2, 0x7e, 0x60, 0x9a, 0x06, 0x1c, 0xbf, 0xc9, 0x12, 0x26, 0x3d, 0x9b, 0xfc, 0x83, 0x60,
	0xa7, 0xac, 0x2e, 0x78, 0x57, 0xa3, 0x73, 0x24, 0x85, 0x68, 0xc3, 0x77, 0x46, 0x48, 0x24, 0x74,
	0x85, 0x0e, 0xa3, 0x64, 0x0e, 0x41, 0xcb, 0xdd, 0xb7, 0xf6, 0x0e, 0x5a, 0xff, 0x85, 0x51, 0xcd,
	0x5e, 0x6d, 0xbf, 0x8a, 0xd6, 0x24, 0xd7, 0x00, 0xdf, 0x6b, 0x27, 0x31, 0x8c, 0xfb, 0x87, 0x6b,
	0x47, 0x7b, 0xeb, 0x36, 0xcb, 0xf4, 0x16, 0x9f, 0x33, 0xa4, 0x4e, 0x3e, 0x06, 0xd8, 0xf0, 0xea,
	0x2b, 0x00, 0x00, 0xff, 0xff, 0xa6, 0x23, 0xac, 0xf5, 0x92, 0x02, 0x00, 0x00,
}
