// Code generated by protoc-gen-gogo.
// source: google/protobuf/type.proto
// DO NOT EDIT!

package google_protobuf

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// The syntax in which a protocol buffer element is defined.
type Syntax int32

const (
	// Syntax `proto2`.
	Syntax_SYNTAX_PROTO2 Syntax = 0
	// Syntax `proto3`.
	Syntax_SYNTAX_PROTO3 Syntax = 1
)

var Syntax_name = map[int32]string{
	0: "SYNTAX_PROTO2",
	1: "SYNTAX_PROTO3",
}
var Syntax_value = map[string]int32{
	"SYNTAX_PROTO2": 0,
	"SYNTAX_PROTO3": 1,
}

func (x Syntax) String() string {
	return proto.EnumName(Syntax_name, int32(x))
}
func (Syntax) EnumDescriptor() ([]byte, []int) { return fileDescriptorType, []int{0} }

// Basic field types.
type Field_Kind int32

const (
	// Field type unknown.
	Field_TYPE_UNKNOWN Field_Kind = 0
	// Field type double.
	Field_TYPE_DOUBLE Field_Kind = 1
	// Field type float.
	Field_TYPE_FLOAT Field_Kind = 2
	// Field type int64.
	Field_TYPE_INT64 Field_Kind = 3
	// Field type uint64.
	Field_TYPE_UINT64 Field_Kind = 4
	// Field type int32.
	Field_TYPE_INT32 Field_Kind = 5
	// Field type fixed64.
	Field_TYPE_FIXED64 Field_Kind = 6
	// Field type fixed32.
	Field_TYPE_FIXED32 Field_Kind = 7
	// Field type bool.
	Field_TYPE_BOOL Field_Kind = 8
	// Field type string.
	Field_TYPE_STRING Field_Kind = 9
	// Field type group. Proto2 syntax only, and deprecated.
	Field_TYPE_GROUP Field_Kind = 10
	// Field type message.
	Field_TYPE_MESSAGE Field_Kind = 11
	// Field type bytes.
	Field_TYPE_BYTES Field_Kind = 12
	// Field type uint32.
	Field_TYPE_UINT32 Field_Kind = 13
	// Field type enum.
	Field_TYPE_ENUM Field_Kind = 14
	// Field type sfixed32.
	Field_TYPE_SFIXED32 Field_Kind = 15
	// Field type sfixed64.
	Field_TYPE_SFIXED64 Field_Kind = 16
	// Field type sint32.
	Field_TYPE_SINT32 Field_Kind = 17
	// Field type sint64.
	Field_TYPE_SINT64 Field_Kind = 18
)

var Field_Kind_name = map[int32]string{
	0:  "TYPE_UNKNOWN",
	1:  "TYPE_DOUBLE",
	2:  "TYPE_FLOAT",
	3:  "TYPE_INT64",
	4:  "TYPE_UINT64",
	5:  "TYPE_INT32",
	6:  "TYPE_FIXED64",
	7:  "TYPE_FIXED32",
	8:  "TYPE_BOOL",
	9:  "TYPE_STRING",
	10: "TYPE_GROUP",
	11: "TYPE_MESSAGE",
	12: "TYPE_BYTES",
	13: "TYPE_UINT32",
	14: "TYPE_ENUM",
	15: "TYPE_SFIXED32",
	16: "TYPE_SFIXED64",
	17: "TYPE_SINT32",
	18: "TYPE_SINT64",
}
var Field_Kind_value = map[string]int32{
	"TYPE_UNKNOWN":  0,
	"TYPE_DOUBLE":   1,
	"TYPE_FLOAT":    2,
	"TYPE_INT64":    3,
	"TYPE_UINT64":   4,
	"TYPE_INT32":    5,
	"TYPE_FIXED64":  6,
	"TYPE_FIXED32":  7,
	"TYPE_BOOL":     8,
	"TYPE_STRING":   9,
	"TYPE_GROUP":    10,
	"TYPE_MESSAGE":  11,
	"TYPE_BYTES":    12,
	"TYPE_UINT32":   13,
	"TYPE_ENUM":     14,
	"TYPE_SFIXED32": 15,
	"TYPE_SFIXED64": 16,
	"TYPE_SINT32":   17,
	"TYPE_SINT64":   18,
}

func (x Field_Kind) String() string {
	return proto.EnumName(Field_Kind_name, int32(x))
}
func (Field_Kind) EnumDescriptor() ([]byte, []int) { return fileDescriptorType, []int{1, 0} }

// Whether a field is optional, required, or repeated.
type Field_Cardinality int32

const (
	// For fields with unknown cardinality.
	Field_CARDINALITY_UNKNOWN Field_Cardinality = 0
	// For optional fields.
	Field_CARDINALITY_OPTIONAL Field_Cardinality = 1
	// For required fields. Proto2 syntax only.
	Field_CARDINALITY_REQUIRED Field_Cardinality = 2
	// For repeated fields.
	Field_CARDINALITY_REPEATED Field_Cardinality = 3
)

var Field_Cardinality_name = map[int32]string{
	0: "CARDINALITY_UNKNOWN",
	1: "CARDINALITY_OPTIONAL",
	2: "CARDINALITY_REQUIRED",
	3: "CARDINALITY_REPEATED",
}
var Field_Cardinality_value = map[string]int32{
	"CARDINALITY_UNKNOWN":  0,
	"CARDINALITY_OPTIONAL": 1,
	"CARDINALITY_REQUIRED": 2,
	"CARDINALITY_REPEATED": 3,
}

func (x Field_Cardinality) String() string {
	return proto.EnumName(Field_Cardinality_name, int32(x))
}
func (Field_Cardinality) EnumDescriptor() ([]byte, []int) { return fileDescriptorType, []int{1, 1} }

// A protocol buffer message type.
type Type struct {
	// The fully qualified message name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The list of fields.
	Fields []*Field `protobuf:"bytes,2,rep,name=fields" json:"fields,omitempty"`
	// The list of types appearing in `oneof` definitions in this type.
	Oneofs []string `protobuf:"bytes,3,rep,name=oneofs" json:"oneofs,omitempty"`
	// The protocol buffer options.
	Options []*Option `protobuf:"bytes,4,rep,name=options" json:"options,omitempty"`
	// The source context.
	SourceContext *SourceContext `protobuf:"bytes,5,opt,name=source_context,json=sourceContext" json:"source_context,omitempty"`
	// The source syntax.
	Syntax Syntax `protobuf:"varint,6,opt,name=syntax,proto3,enum=google.protobuf.Syntax" json:"syntax,omitempty"`
}

func (m *Type) Reset()                    { *m = Type{} }
func (m *Type) String() string            { return proto.CompactTextString(m) }
func (*Type) ProtoMessage()               {}
func (*Type) Descriptor() ([]byte, []int) { return fileDescriptorType, []int{0} }

func (m *Type) GetFields() []*Field {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *Type) GetOptions() []*Option {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *Type) GetSourceContext() *SourceContext {
	if m != nil {
		return m.SourceContext
	}
	return nil
}

// A single field of a message type.
type Field struct {
	// The field type.
	Kind Field_Kind `protobuf:"varint,1,opt,name=kind,proto3,enum=google.protobuf.Field_Kind" json:"kind,omitempty"`
	// The field cardinality.
	Cardinality Field_Cardinality `protobuf:"varint,2,opt,name=cardinality,proto3,enum=google.protobuf.Field_Cardinality" json:"cardinality,omitempty"`
	// The field number.
	Number int32 `protobuf:"varint,3,opt,name=number,proto3" json:"number,omitempty"`
	// The field name.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// The field type URL, without the scheme, for message or enumeration
	// types. Example: `"type.googleapis.com/google.protobuf.Timestamp"`.
	TypeUrl string `protobuf:"bytes,6,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	// The index of the field type in `Type.oneofs`, for message or enumeration
	// types. The first type has index 1; zero means the type is not in the list.
	OneofIndex int32 `protobuf:"varint,7,opt,name=oneof_index,json=oneofIndex,proto3" json:"oneof_index,omitempty"`
	// Whether to use alternative packed wire representation.
	Packed bool `protobuf:"varint,8,opt,name=packed,proto3" json:"packed,omitempty"`
	// The protocol buffer options.
	Options []*Option `protobuf:"bytes,9,rep,name=options" json:"options,omitempty"`
	// The field JSON name.
	JsonName string `protobuf:"bytes,10,opt,name=json_name,json=jsonName,proto3" json:"json_name,omitempty"`
	// The string value of the default value of this field. Proto2 syntax only.
	DefaultValue string `protobuf:"bytes,11,opt,name=default_value,json=defaultValue,proto3" json:"default_value,omitempty"`
}

func (m *Field) Reset()                    { *m = Field{} }
func (m *Field) String() string            { return proto.CompactTextString(m) }
func (*Field) ProtoMessage()               {}
func (*Field) Descriptor() ([]byte, []int) { return fileDescriptorType, []int{1} }

func (m *Field) GetOptions() []*Option {
	if m != nil {
		return m.Options
	}
	return nil
}

// Enum type definition.
type Enum struct {
	// Enum type name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Enum value definitions.
	Enumvalue []*EnumValue `protobuf:"bytes,2,rep,name=enumvalue" json:"enumvalue,omitempty"`
	// Protocol buffer options.
	Options []*Option `protobuf:"bytes,3,rep,name=options" json:"options,omitempty"`
	// The source context.
	SourceContext *SourceContext `protobuf:"bytes,4,opt,name=source_context,json=sourceContext" json:"source_context,omitempty"`
	// The source syntax.
	Syntax Syntax `protobuf:"varint,5,opt,name=syntax,proto3,enum=google.protobuf.Syntax" json:"syntax,omitempty"`
}

func (m *Enum) Reset()                    { *m = Enum{} }
func (m *Enum) String() string            { return proto.CompactTextString(m) }
func (*Enum) ProtoMessage()               {}
func (*Enum) Descriptor() ([]byte, []int) { return fileDescriptorType, []int{2} }

func (m *Enum) GetEnumvalue() []*EnumValue {
	if m != nil {
		return m.Enumvalue
	}
	return nil
}

func (m *Enum) GetOptions() []*Option {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *Enum) GetSourceContext() *SourceContext {
	if m != nil {
		return m.SourceContext
	}
	return nil
}

// Enum value definition.
type EnumValue struct {
	// Enum value name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Enum value number.
	Number int32 `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
	// Protocol buffer options.
	Options []*Option `protobuf:"bytes,3,rep,name=options" json:"options,omitempty"`
}

func (m *EnumValue) Reset()                    { *m = EnumValue{} }
func (m *EnumValue) String() string            { return proto.CompactTextString(m) }
func (*EnumValue) ProtoMessage()               {}
func (*EnumValue) Descriptor() ([]byte, []int) { return fileDescriptorType, []int{3} }

func (m *EnumValue) GetOptions() []*Option {
	if m != nil {
		return m.Options
	}
	return nil
}

// A protocol buffer option, which can be attached to a message, field,
// enumeration, etc.
type Option struct {
	// The option's name. For example, `"java_package"`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The option's value. For example, `"com.google.protobuf"`.
	Value *Any `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *Option) Reset()                    { *m = Option{} }
func (m *Option) String() string            { return proto.CompactTextString(m) }
func (*Option) ProtoMessage()               {}
func (*Option) Descriptor() ([]byte, []int) { return fileDescriptorType, []int{4} }

func (m *Option) GetValue() *Any {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*Type)(nil), "google.protobuf.Type")
	proto.RegisterType((*Field)(nil), "google.protobuf.Field")
	proto.RegisterType((*Enum)(nil), "google.protobuf.Enum")
	proto.RegisterType((*EnumValue)(nil), "google.protobuf.EnumValue")
	proto.RegisterType((*Option)(nil), "google.protobuf.Option")
	proto.RegisterEnum("google.protobuf.Syntax", Syntax_name, Syntax_value)
	proto.RegisterEnum("google.protobuf.Field_Kind", Field_Kind_name, Field_Kind_value)
	proto.RegisterEnum("google.protobuf.Field_Cardinality", Field_Cardinality_name, Field_Cardinality_value)
}

var fileDescriptorType = []byte{
	// 768 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x55, 0xcd, 0x6e, 0xda, 0x58,
	0x14, 0x1e, 0x83, 0x31, 0xf8, 0x38, 0x10, 0xe7, 0x26, 0x4a, 0x1c, 0x22, 0x65, 0x22, 0x66, 0x16,
	0x51, 0x16, 0x44, 0x43, 0x46, 0xa3, 0xd9, 0x42, 0x70, 0x18, 0x2b, 0xc4, 0xf6, 0x5c, 0x4c, 0x13,
	0x56, 0xc8, 0x01, 0x13, 0x91, 0x38, 0x36, 0xc2, 0xa6, 0x0d, 0x0f, 0xd1, 0x77, 0xa8, 0xba, 0xec,
	0xba, 0x0f, 0xd1, 0xb7, 0x6a, 0xef, 0xbd, 0x06, 0x63, 0x7e, 0x2a, 0xa5, 0xed, 0x8e, 0xf3, 0x7d,
	0xdf, 0xf9, 0xbd, 0xc7, 0x07, 0x28, 0x3e, 0xf8, 0xfe, 0x83, 0xeb, 0x9c, 0x8f, 0xc6, 0x7e, 0xe8,
	0xdf, 0x4f, 0x06, 0xe7, 0xe1, 0x74, 0xe4, 0x94, 0x99, 0x85, 0xb6, 0x23, 0xae, 0x3c, 0xe7, 0x8a,
	0x87, 0xab, 0x62, 0xdb, 0x9b, 0x46, 0x6c, 0xf1, 0xcf, 0x55, 0x2a, 0xf0, 0x27, 0xe3, 0x9e, 0xd3,
	0xed, 0xf9, 0x5e, 0xe8, 0xbc, 0x84, 0x91, 0xaa, 0xf4, 0x3e, 0x05, 0xbc, 0x45, 0x12, 0x20, 0x04,
	0xbc, 0x67, 0x3f, 0x3b, 0x0a, 0x77, 0xc2, 0x9d, 0x8a, 0x98, 0xfd, 0x46, 0x65, 0x10, 0x06, 0x43,
	0xc7, 0xed, 0x07, 0x4a, 0xea, 0x24, 0x7d, 0x2a, 0x55, 0xf6, 0xcb, 0x2b, 0xf9, 0xcb, 0x57, 0x94,
	0xc6, 0x33, 0x15, 0xda, 0x07, 0xc1, 0xf7, 0x1c, 0x7f, 0x10, 0x28, 0x69, 0xa2, 0x17, 0xf1, 0xcc,
	0x42, 0x7f, 0x41, 0xd6, 0x1f, 0x85, 0x43, 0xdf, 0x0b, 0x14, 0x9e, 0x05, 0x3a, 0x58, 0x0b, 0x64,
	0x30, 0x1e, 0xcf, 0x75, 0x48, 0x85, 0xc2, 0x72, 0xbd, 0x4a, 0x86, 0x14, 0x26, 0x55, 0x8e, 0xd7,
	0x3c, 0x5b, 0x4c, 0x76, 0x19, 0xa9, 0x70, 0x3e, 0x48, 0x9a, 0xe8, 0x1c, 0x84, 0x60, 0xea, 0x85,
	0xf6, 0x8b, 0x22, 0x10, 0xf7, 0xc2, 0x86, 0xc4, 0x2d, 0x46, 0xe3, 0x99, 0xac, 0xf4, 0x59, 0x80,
	0x0c, 0x6b, 0x8a, 0xb8, 0xf2, 0x4f, 0x43, 0xaf, 0xcf, 0x06, 0x52, 0xa8, 0x1c, 0x6d, 0x6e, 0xbd,
	0x7c, 0x4d, 0x24, 0x98, 0x09, 0x51, 0x1d, 0xa4, 0x9e, 0x3d, 0xee, 0x0f, 0x3d, 0xdb, 0x1d, 0x86,
	0x53, 0x32, 0x32, 0xea, 0x57, 0xfa, 0x8e, 0xdf, 0xe5, 0x42, 0x89, 0x93, 0x6e, 0x74, 0x86, 0xde,
	0xe4, 0xf9, 0xde, 0x19, 0x93, 0x19, 0x72, 0xa7, 0x19, 0x3c, 0xb3, 0xe2, 0xf7, 0xe1, 0x13, 0xef,
	0x73, 0x08, 0x39, 0xba, 0x1c, 0xdd, 0xc9, 0xd8, 0x65, 0xfd, 0x89, 0x38, 0x4b, 0xed, 0xf6, 0xd8,
	0x45, 0xbf, 0x83, 0xc4, 0x86, 0xdf, 0x25, 0x95, 0x39, 0x2f, 0x4a, 0x96, 0xc5, 0x02, 0x06, 0x69,
	0x14, 0xa1, 0x79, 0x46, 0x76, 0xef, 0xc9, 0xe9, 0x2b, 0x39, 0xc2, 0xe5, 0xf0, 0xcc, 0x4a, 0xbe,
	0x95, 0xf8, 0xca, 0xb7, 0x3a, 0x02, 0xf1, 0x31, 0xf0, 0xbd, 0x2e, 0xab, 0x0f, 0x58, 0x1d, 0x39,
	0x0a, 0xe8, 0xb4, 0xc6, 0x3f, 0x20, 0xdf, 0x77, 0x06, 0xf6, 0xc4, 0x0d, 0xbb, 0x6f, 0x6d, 0x77,
	0xe2, 0x28, 0x12, 0x13, 0x6c, 0xcd, 0xc0, 0x37, 0x14, 0x2b, 0x7d, 0x21, 0x5b, 0x48, 0x27, 0x89,
	0x64, 0xd8, 0xb2, 0x3a, 0xa6, 0xda, 0x6d, 0xeb, 0xd7, 0xba, 0x71, 0xab, 0xcb, 0xbf, 0xa1, 0x6d,
	0x90, 0x18, 0x52, 0x37, 0xda, 0xb5, 0xa6, 0x2a, 0x73, 0xa8, 0x00, 0xc0, 0x80, 0xab, 0xa6, 0x51,
	0xb5, 0xe4, 0x54, 0x6c, 0x6b, 0xba, 0xf5, 0xcf, 0xdf, 0x72, 0x3a, 0x76, 0x68, 0x47, 0x00, 0x9f,
	0x14, 0x5c, 0x54, 0xe4, 0x4c, 0x9c, 0xe3, 0x4a, 0xbb, 0x53, 0xeb, 0x44, 0x21, 0x2c, 0x23, 0x44,
	0x93, 0x45, 0x79, 0x10, 0x19, 0x52, 0x33, 0x8c, 0xa6, 0x9c, 0x8b, 0x63, 0xb6, 0x2c, 0xac, 0xe9,
	0x0d, 0x59, 0x8c, 0x63, 0x36, 0xb0, 0xd1, 0x36, 0x65, 0x88, 0x23, 0xdc, 0xa8, 0xad, 0x56, 0xb5,
	0xa1, 0xca, 0x52, 0xac, 0xa8, 0x75, 0x2c, 0xb5, 0x25, 0x6f, 0x2d, 0x95, 0x45, 0x52, 0xe4, 0xe3,
	0x14, 0xaa, 0xde, 0xbe, 0x91, 0x0b, 0x68, 0x07, 0xf2, 0x51, 0x8a, 0x79, 0x11, 0xdb, 0x2b, 0x10,
	0xa9, 0x54, 0x5e, 0x14, 0x12, 0x45, 0xd9, 0x59, 0x02, 0x88, 0x02, 0x95, 0x42, 0x90, 0x12, 0xbb,
	0x85, 0x0e, 0x60, 0xf7, 0xb2, 0x8a, 0xeb, 0x9a, 0x5e, 0x6d, 0x6a, 0x56, 0x27, 0x31, 0x57, 0x05,
	0xf6, 0x92, 0x84, 0x61, 0x5a, 0x9a, 0x41, 0x7e, 0x93, 0x01, 0xaf, 0x30, 0x58, 0xfd, 0xbf, 0xad,
	0x61, 0xb5, 0x4e, 0x46, 0xbd, 0xc6, 0x98, 0x6a, 0xd5, 0x22, 0x4c, 0xba, 0xf4, 0x95, 0x03, 0x5e,
	0x25, 0x9b, 0xba, 0xf1, 0x8c, 0xfc, 0x0b, 0xa2, 0x43, 0xb8, 0xe8, 0xf9, 0xa3, 0x4b, 0x52, 0x5c,
	0x5b, 0x2a, 0xea, 0xcd, 0x96, 0x01, 0x2f, 0xc4, 0xc9, 0x65, 0x4c, 0xff, 0xf4, 0xe1, 0xe0, 0x7f,
	0xed, 0x70, 0x64, 0x5e, 0x77, 0x38, 0x1e, 0x41, 0x8c, 0x5b, 0xd8, 0x38, 0x85, 0xc5, 0x87, 0x9d,
	0x5a, 0xfa, 0xb0, 0x7f, 0xbc, 0xc7, 0xd2, 0x7f, 0x20, 0x44, 0xd0, 0xc6, 0x44, 0x67, 0x90, 0x99,
	0x8f, 0x9a, 0x36, 0xbe, 0xb7, 0x16, 0xae, 0xea, 0x4d, 0x71, 0x24, 0x39, 0x23, 0x17, 0x3e, 0xea,
	0x83, 0x2e, 0x5b, 0xab, 0xa3, 0x5b, 0xd5, 0xbb, 0xae, 0x89, 0x0d, 0xcb, 0xa8, 0x90, 0x15, 0x59,
	0x81, 0x2e, 0x64, 0xae, 0xd6, 0x84, 0xdd, 0x9e, 0xff, 0xbc, 0x1a, 0xb1, 0x26, 0xd2, 0xbf, 0x10,
	0x93, 0x5a, 0x26, 0xf7, 0x81, 0xe3, 0x3e, 0xa6, 0xd2, 0x0d, 0xb3, 0xf6, 0x29, 0x75, 0xdc, 0x88,
	0x74, 0xe6, 0x3c, 0xf3, 0xad, 0xe3, 0xba, 0xd7, 0x9e, 0xff, 0xce, 0xa3, 0xfa, 0xe0, 0x5e, 0x60,
	0x01, 0x2e, 0xbe, 0x05, 0x00, 0x00, 0xff, 0xff, 0x95, 0xbb, 0xeb, 0x52, 0xf3, 0x06, 0x00, 0x00,
}