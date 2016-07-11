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

// Ignoring public import of EProtoId from base.proto

type User struct {
	Mid                *string `protobuf:"bytes,1,opt,name=mid" json:"mid,omitempty"`
	Name               *string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Coin               *int32  `protobuf:"varint,3,opt,name=coin" json:"coin,omitempty"`
	Id                 *uint32 `protobuf:"varint,4,opt,name=id" json:"id,omitempty"`
	NickName           *string `protobuf:"bytes,5,opt,name=nickName" json:"nickName,omitempty"`
	LoginTurntable     *bool   `protobuf:"varint,6,opt,name=loginTurntable" json:"loginTurntable,omitempty"`
	LoginTurntableTime *string `protobuf:"bytes,7,opt,name=loginTurntableTime" json:"loginTurntableTime,omitempty"`
	Scores             *int32  `protobuf:"varint,8,opt,name=scores" json:"scores,omitempty"`
	LastSignTime       *string `protobuf:"bytes,9,opt,name=lastSignTime" json:"lastSignTime,omitempty"`
	SignCount          *int32  `protobuf:"varint,10,opt,name=signCount" json:"signCount,omitempty"`
	XXX_unrecognized   []byte  `json:"-"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{0} }

func (m *User) GetMid() string {
	if m != nil && m.Mid != nil {
		return *m.Mid
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *User) GetCoin() int32 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *User) GetId() uint32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *User) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *User) GetLoginTurntable() bool {
	if m != nil && m.LoginTurntable != nil {
		return *m.LoginTurntable
	}
	return false
}

func (m *User) GetLoginTurntableTime() string {
	if m != nil && m.LoginTurntableTime != nil {
		return *m.LoginTurntableTime
	}
	return ""
}

func (m *User) GetScores() int32 {
	if m != nil && m.Scores != nil {
		return *m.Scores
	}
	return 0
}

func (m *User) GetLastSignTime() string {
	if m != nil && m.LastSignTime != nil {
		return *m.LastSignTime
	}
	return ""
}

func (m *User) GetSignCount() int32 {
	if m != nil && m.SignCount != nil {
		return *m.SignCount
	}
	return 0
}

func init() {
	proto.RegisterType((*User)(nil), "bbproto.User")
}

var fileDescriptor10 = []byte{
	// 188 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0xcd, 0x4f, 0x8a, 0xc2, 0x30,
	0x14, 0xc7, 0xf1, 0xe9, 0xff, 0xf6, 0x37, 0x9d, 0x32, 0x06, 0x91, 0xd0, 0x95, 0xb8, 0x72, 0xe5,
	0x25, 0xdc, 0x8b, 0x60, 0x3d, 0x40, 0x52, 0x43, 0x09, 0xb6, 0x89, 0x24, 0xe9, 0xe5, 0x3c, 0x9d,
	0x69, 0xc4, 0x85, 0xbb, 0xf7, 0xbe, 0x7c, 0x1e, 0x0f, 0x98, 0xad, 0x30, 0x87, 0x87, 0xd1, 0x4e,
	0x93, 0x82, 0xf3, 0x30, 0xb4, 0xe0, 0xcc, 0x8a, 0x77, 0xdc, 0x3d, 0x23, 0xa4, 0x57, 0x6f, 0xc8,
	0x2f, 0x92, 0x49, 0xde, 0x68, 0xb4, 0x8d, 0xf6, 0x15, 0xa9, 0x91, 0x2a, 0x36, 0x09, 0x1a, 0x7f,
	0xb6, 0x5e, 0x4b, 0x45, 0x13, 0xbf, 0x65, 0x04, 0x88, 0xbd, 0x4b, 0xfd, 0xfc, 0x47, 0xfe, 0x51,
	0x2a, 0xd9, 0xdf, 0x4f, 0x8b, 0xcd, 0x82, 0xdd, 0xa0, 0x19, 0xf5, 0x20, 0x55, 0x37, 0x1b, 0xe5,
	0x18, 0x1f, 0x05, 0xcd, 0x7d, 0x2f, 0x49, 0x0b, 0xf2, 0xdd, 0x3b, 0xe9, 0x6f, 0x8a, 0x70, 0xd3,
	0x20, 0xb7, 0xbd, 0x36, 0xc2, 0xd2, 0x32, 0x7c, 0x58, 0xa3, 0x1e, 0x99, 0x75, 0x17, 0x39, 0xa8,
	0xa0, 0xaa, 0xa0, 0x56, 0xa8, 0xac, 0x2f, 0x47, 0x3d, 0x2b, 0x47, 0xb1, 0xc0, 0xf3, 0xcf, 0x2b,
	0x00, 0x00, 0xff, 0xff, 0x01, 0xf6, 0xb5, 0x87, 0xe0, 0x00, 0x00, 0x00,
}
