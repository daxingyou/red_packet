// Code generated by protoc-gen-go.
// source: base.proto
// DO NOT EDIT!

package mjproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type EProtoId int32

const (
	EProtoId_PID_GAME_CREATEROOM EProtoId = 1
)

var EProtoId_name = map[int32]string{
	1: "PID_GAME_CREATEROOM",
}
var EProtoId_value = map[string]int32{
	"PID_GAME_CREATEROOM": 1,
}

func (x EProtoId) Enum() *EProtoId {
	p := new(EProtoId)
	*p = x
	return p
}
func (x EProtoId) String() string {
	return proto.EnumName(EProtoId_name, int32(x))
}
func (x *EProtoId) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EProtoId_value, data, "EProtoId")
	if err != nil {
		return err
	}
	*x = EProtoId(value)
	return nil
}
func (EProtoId) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type DDErrorCode int32

const (
	DDErrorCode_ERRORCODE_SUCC DDErrorCode = 0
	// -101   -200	游戏异常
	DDErrorCode_ERRORCODE_CREATE_DESK_DIAMOND_NOTENOUGH DDErrorCode = -101
	DDErrorCode_ERRORCODE_CREATE_DESK_USER_NOTFOUND     DDErrorCode = -102
	DDErrorCode_ERRORCODE_INTO_DESK_NOTFOUND            DDErrorCode = -103
	DDErrorCode_ERRORCODE_GAME_READY_REPEAT             DDErrorCode = -110
	DDErrorCode_ERRORCODE_GAME_READY_CHIP_NOT_ENOUGH    DDErrorCode = -111
)

var DDErrorCode_name = map[int32]string{
	0:    "ERRORCODE_SUCC",
	-101: "ERRORCODE_CREATE_DESK_DIAMOND_NOTENOUGH",
	-102: "ERRORCODE_CREATE_DESK_USER_NOTFOUND",
	-103: "ERRORCODE_INTO_DESK_NOTFOUND",
	-110: "ERRORCODE_GAME_READY_REPEAT",
	-111: "ERRORCODE_GAME_READY_CHIP_NOT_ENOUGH",
}
var DDErrorCode_value = map[string]int32{
	"ERRORCODE_SUCC":                          0,
	"ERRORCODE_CREATE_DESK_DIAMOND_NOTENOUGH": -101,
	"ERRORCODE_CREATE_DESK_USER_NOTFOUND":     -102,
	"ERRORCODE_INTO_DESK_NOTFOUND":            -103,
	"ERRORCODE_GAME_READY_REPEAT":             -110,
	"ERRORCODE_GAME_READY_CHIP_NOT_ENOUGH":    -111,
}

func (x DDErrorCode) Enum() *DDErrorCode {
	p := new(DDErrorCode)
	*p = x
	return p
}
func (x DDErrorCode) String() string {
	return proto.EnumName(DDErrorCode_name, int32(x))
}
func (x *DDErrorCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DDErrorCode_value, data, "DDErrorCode")
	if err != nil {
		return err
	}
	*x = DDErrorCode(value)
	return nil
}
func (DDErrorCode) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

// ProtoHeader 需要在每个 Message 中作为第一个字段
type ProtoHeader struct {
	Version          *string `protobuf:"bytes,1,opt,name=version" json:"version,omitempty"`
	UserId           *uint32 `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	Code             *int32  `protobuf:"varint,3,opt,name=code" json:"code,omitempty"`
	Error            *string `protobuf:"bytes,4,opt,name=error" json:"error,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ProtoHeader) Reset()                    { *m = ProtoHeader{} }
func (m *ProtoHeader) String() string            { return proto.CompactTextString(m) }
func (*ProtoHeader) ProtoMessage()               {}
func (*ProtoHeader) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *ProtoHeader) GetVersion() string {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return ""
}

func (m *ProtoHeader) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *ProtoHeader) GetCode() int32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return 0
}

func (m *ProtoHeader) GetError() string {
	if m != nil && m.Error != nil {
		return *m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*ProtoHeader)(nil), "mjproto.ProtoHeader")
	proto.RegisterEnum("mjproto.EProtoId", EProtoId_name, EProtoId_value)
	proto.RegisterEnum("mjproto.DDErrorCode", DDErrorCode_name, DDErrorCode_value)
}

var fileDescriptor1 = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x8f, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x49, 0x69, 0x29, 0x5c, 0x69, 0x89, 0xcc, 0x40, 0x24, 0x18, 0x10, 0x45, 0xa2, 0x74,
	0x40, 0x20, 0xf1, 0x02, 0x95, 0x7d, 0x34, 0x56, 0x95, 0x38, 0x72, 0x92, 0x81, 0xc9, 0x2a, 0xc4,
	0x03, 0x48, 0x10, 0xe4, 0x00, 0xef, 0x51, 0x36, 0xe0, 0x61, 0xc1, 0x49, 0xab, 0x7a, 0xa9, 0x07,
	0xeb, 0xce, 0xfa, 0xbe, 0xff, 0x7c, 0x00, 0x0f, 0xf3, 0x4a, 0x5f, 0xbd, 0x99, 0xf2, 0xbd, 0x24,
	0xdd, 0x97, 0xe7, 0xa6, 0x38, 0x9b, 0x41, 0x2f, 0xa9, 0x8b, 0x50, 0xcf, 0x0b, 0x6d, 0xc8, 0x01,
	0x74, 0x3f, 0xb5, 0xa9, 0x9e, 0xca, 0xd7, 0xc0, 0x3b, 0xf5, 0x46, 0x7b, 0x64, 0x00, 0x3b, 0x1f,
	0x95, 0x36, 0xbc, 0x08, 0x5a, 0xb6, 0xef, 0x93, 0x7d, 0x68, 0x3f, 0x96, 0x85, 0x0e, 0xb6, 0x6d,
	0xd7, 0x21, 0x7d, 0xe8, 0x68, 0x63, 0x4a, 0x13, 0xb4, 0x6b, 0x78, 0x3c, 0x84, 0x5d, 0x6c, 0xd2,
	0x78, 0x41, 0x8e, 0xe0, 0x30, 0xe1, 0x4c, 0x4d, 0x27, 0x11, 0x2a, 0x2a, 0x71, 0x92, 0xa1, 0x14,
	0x22, 0xf2, 0xbd, 0xf1, 0xa2, 0x05, 0x3d, 0xc6, 0xb0, 0xd6, 0xa8, 0x4d, 0x22, 0x04, 0x06, 0x28,
	0xa5, 0x90, 0x54, 0x30, 0x54, 0x69, 0x4e, 0xa9, 0xbf, 0x45, 0x6e, 0xe1, 0xc2, 0xbd, 0x2d, 0x6d,
	0xc5, 0x30, 0x9d, 0x29, 0xc6, 0x27, 0x91, 0x88, 0x99, 0x8a, 0x45, 0x86, 0xb1, 0xc8, 0xa7, 0xa1,
	0xff, 0xfb, 0xb7, 0x3a, 0x1e, 0xb9, 0x86, 0xe1, 0x66, 0x2b, 0x4f, 0x51, 0xd6, 0xca, 0x9d, 0xc8,
	0x63, 0xe6, 0xff, 0x38, 0xe3, 0x12, 0x4e, 0x9c, 0xc1, 0xe3, 0x4c, 0x2c, 0xf9, 0x35, 0xfa, 0xed,
	0xd0, 0x11, 0x1c, 0x3b, 0xb4, 0xd9, 0xca, 0x0e, 0x60, 0xf7, 0xf6, 0x4e, 0xec, 0x20, 0xff, 0xcb,
	0x91, 0x37, 0x70, 0xbe, 0x91, 0xa4, 0x21, 0x4f, 0xea, 0x68, 0xb5, 0xfa, 0xf9, 0x62, 0xad, 0xfc,
	0x07, 0x00, 0x00, 0xff, 0xff, 0x10, 0xac, 0xc7, 0xd0, 0x9b, 0x01, 0x00, 0x00,
}