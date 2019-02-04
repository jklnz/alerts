// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/conversion_action_counting_type.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v0/enums"

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

// Indicates how conversions for this action will be counted. For more
// information, see https://support.google.com/google-ads/answer/3438531.
type ConversionActionCountingTypeEnum_ConversionActionCountingType int32

const (
	// Not specified.
	ConversionActionCountingTypeEnum_UNSPECIFIED ConversionActionCountingTypeEnum_ConversionActionCountingType = 0
	// Used for return value only. Represents value unknown in this version.
	ConversionActionCountingTypeEnum_UNKNOWN ConversionActionCountingTypeEnum_ConversionActionCountingType = 1
	// Count only one conversion per click.
	ConversionActionCountingTypeEnum_ONE_PER_CLICK ConversionActionCountingTypeEnum_ConversionActionCountingType = 2
	// Count all conversions per click.
	ConversionActionCountingTypeEnum_MANY_PER_CLICK ConversionActionCountingTypeEnum_ConversionActionCountingType = 3
)

var ConversionActionCountingTypeEnum_ConversionActionCountingType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ONE_PER_CLICK",
	3: "MANY_PER_CLICK",
}
var ConversionActionCountingTypeEnum_ConversionActionCountingType_value = map[string]int32{
	"UNSPECIFIED":    0,
	"UNKNOWN":        1,
	"ONE_PER_CLICK":  2,
	"MANY_PER_CLICK": 3,
}

func (x ConversionActionCountingTypeEnum_ConversionActionCountingType) String() string {
	return proto.EnumName(ConversionActionCountingTypeEnum_ConversionActionCountingType_name, int32(x))
}
func (ConversionActionCountingTypeEnum_ConversionActionCountingType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_conversion_action_counting_type_0429eb2d6c22eb59, []int{0, 0}
}

// Container for enum describing the conversion deduplication mode for
// conversion optimizer.
type ConversionActionCountingTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConversionActionCountingTypeEnum) Reset()         { *m = ConversionActionCountingTypeEnum{} }
func (m *ConversionActionCountingTypeEnum) String() string { return proto.CompactTextString(m) }
func (*ConversionActionCountingTypeEnum) ProtoMessage()    {}
func (*ConversionActionCountingTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_conversion_action_counting_type_0429eb2d6c22eb59, []int{0}
}
func (m *ConversionActionCountingTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConversionActionCountingTypeEnum.Unmarshal(m, b)
}
func (m *ConversionActionCountingTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConversionActionCountingTypeEnum.Marshal(b, m, deterministic)
}
func (dst *ConversionActionCountingTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConversionActionCountingTypeEnum.Merge(dst, src)
}
func (m *ConversionActionCountingTypeEnum) XXX_Size() int {
	return xxx_messageInfo_ConversionActionCountingTypeEnum.Size(m)
}
func (m *ConversionActionCountingTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ConversionActionCountingTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ConversionActionCountingTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ConversionActionCountingTypeEnum)(nil), "google.ads.googleads.v0.enums.ConversionActionCountingTypeEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.ConversionActionCountingTypeEnum_ConversionActionCountingType", ConversionActionCountingTypeEnum_ConversionActionCountingType_name, ConversionActionCountingTypeEnum_ConversionActionCountingType_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/conversion_action_counting_type.proto", fileDescriptor_conversion_action_counting_type_0429eb2d6c22eb59)
}

var fileDescriptor_conversion_action_counting_type_0429eb2d6c22eb59 = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xd1, 0x4a, 0xf3, 0x30,
	0x1c, 0xc5, 0xbf, 0x76, 0xf0, 0x09, 0x19, 0xea, 0xcc, 0xb5, 0x03, 0xb7, 0x07, 0x48, 0x0b, 0xde,
	0xc5, 0xab, 0x2c, 0xd6, 0x31, 0xa6, 0x59, 0x51, 0x37, 0x51, 0x0a, 0xa5, 0xb6, 0x21, 0x14, 0xd6,
	0xa4, 0x34, 0x6d, 0x61, 0x4f, 0xe0, 0x7b, 0x78, 0xe9, 0xa3, 0xf8, 0x28, 0x3e, 0x80, 0xd7, 0xd2,
	0x74, 0x2b, 0xde, 0xd8, 0x9b, 0xe4, 0xf0, 0xff, 0x9f, 0xfc, 0x38, 0x39, 0x80, 0x0a, 0xa5, 0xc4,
	0x96, 0x3b, 0x51, 0xa2, 0x9d, 0x56, 0x36, 0xaa, 0x76, 0x1d, 0x2e, 0xab, 0x4c, 0x3b, 0xb1, 0x92,
	0x35, 0x2f, 0x74, 0xaa, 0x64, 0x18, 0xc5, 0x65, 0x73, 0xc5, 0xaa, 0x92, 0x65, 0x2a, 0x45, 0x58,
	0xee, 0x72, 0x8e, 0xf2, 0x42, 0x95, 0x0a, 0x8e, 0xdb, 0x97, 0x28, 0x4a, 0x34, 0xea, 0x20, 0xa8,
	0x76, 0x91, 0x81, 0x4c, 0xdf, 0x2c, 0x70, 0x41, 0x3b, 0x10, 0x31, 0x1c, 0xba, 0xc7, 0x3c, 0xee,
	0x72, 0xee, 0xc9, 0x2a, 0x9b, 0xc6, 0xe0, 0xbc, 0xcf, 0x03, 0x4f, 0xc1, 0x70, 0xcd, 0x1e, 0x7c,
	0x8f, 0x2e, 0x6e, 0x16, 0xde, 0xf5, 0xe8, 0x1f, 0x1c, 0x82, 0xa3, 0x35, 0x5b, 0xb2, 0xd5, 0x13,
	0x1b, 0x59, 0xf0, 0x0c, 0x1c, 0xaf, 0x98, 0x17, 0xfa, 0xde, 0x7d, 0x48, 0x6f, 0x17, 0x74, 0x39,
	0xb2, 0x21, 0x04, 0x27, 0x77, 0x84, 0x3d, 0xff, 0x9a, 0x0d, 0x66, 0xdf, 0x16, 0x98, 0xc4, 0x2a,
	0x43, 0xbd, 0x79, 0x67, 0x93, 0xbe, 0x20, 0x7e, 0xf3, 0x63, 0xdf, 0x7a, 0x99, 0xed, 0x19, 0x42,
	0x6d, 0x23, 0x29, 0x90, 0x2a, 0x84, 0x23, 0xb8, 0x34, 0x7d, 0x1c, 0x8a, 0xcc, 0x53, 0xfd, 0x47,
	0xaf, 0x57, 0xe6, 0x7c, 0xb7, 0x07, 0x73, 0x42, 0x3e, 0xec, 0xf1, 0xbc, 0x45, 0x91, 0x44, 0xa3,
	0x56, 0x36, 0x6a, 0xe3, 0xa2, 0xa6, 0x18, 0xfd, 0x79, 0xd8, 0x07, 0x24, 0xd1, 0x41, 0xb7, 0x0f,
	0x36, 0x6e, 0x60, 0xf6, 0x5f, 0xf6, 0xa4, 0x1d, 0x62, 0x4c, 0x12, 0x8d, 0x71, 0xe7, 0xc0, 0x78,
	0xe3, 0x62, 0x6c, 0x3c, 0xaf, 0xff, 0x4d, 0xb0, 0xcb, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe1,
	0x28, 0xb6, 0xde, 0xef, 0x01, 0x00, 0x00,
}
