// Code generated by protoc-gen-go.
// source: user.proto
// DO NOT EDIT!

package bbproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of ProtoHeader from base.proto

// Ignoring public import of TerminalInfo from base.proto

// Ignoring public import of EUnitType from base.proto

// Ignoring public import of EUnitRace from base.proto

// Ignoring public import of EUnitProtoId from base.proto

type User struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Name             *string      `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Balance          *int32       `protobuf:"varint,3,opt,name=balance" json:"balance,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{0} }

func (m *User) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *User) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *User) GetBalance() int32 {
	if m != nil && m.Balance != nil {
		return *m.Balance
	}
	return 0
}

func init() {
	proto.RegisterType((*User)(nil), "bbproto.User")
}

var fileDescriptor9 = []byte{
	// 115 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x2d, 0x4e, 0x2d,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4f, 0x4a, 0x02, 0x33, 0xa4, 0xb8, 0x92, 0x12,
	0x8b, 0x53, 0x21, 0x82, 0x4a, 0xde, 0x5c, 0x2c, 0xa1, 0x40, 0x25, 0x42, 0x2a, 0x5c, 0x6c, 0x19,
	0xa9, 0x89, 0x29, 0xa9, 0x45, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x22, 0x7a, 0x50, 0xd5,
	0x7a, 0x01, 0x20, 0xd2, 0x03, 0x2c, 0x27, 0xc4, 0xc3, 0xc5, 0x92, 0x97, 0x98, 0x9b, 0x2a, 0xc1,
	0x04, 0x54, 0xc3, 0x29, 0xc4, 0xcf, 0xc5, 0x9e, 0x94, 0x98, 0x93, 0x98, 0x97, 0x9c, 0x2a, 0xc1,
	0x0c, 0x14, 0x60, 0x0d, 0x60, 0x00, 0x04, 0x00, 0x00, 0xff, 0xff, 0xce, 0xd4, 0xcf, 0x74, 0x70,
	0x00, 0x00, 0x00,
}
