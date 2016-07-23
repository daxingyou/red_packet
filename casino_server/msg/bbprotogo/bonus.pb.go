// Code generated by protoc-gen-go.
// source: bonus.proto
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

// 登陆转盘奖励,每天最多一次
type LoginTurntableBonus struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	SuccIndex        *int32       `protobuf:"varint,2,opt,name=succIndex" json:"succIndex,omitempty"`
	BonusAmount      *int32       `protobuf:"varint,3,opt,name=bonusAmount" json:"bonusAmount,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *LoginTurntableBonus) Reset()                    { *m = LoginTurntableBonus{} }
func (m *LoginTurntableBonus) String() string            { return proto.CompactTextString(m) }
func (*LoginTurntableBonus) ProtoMessage()               {}
func (*LoginTurntableBonus) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *LoginTurntableBonus) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *LoginTurntableBonus) GetSuccIndex() int32 {
	if m != nil && m.SuccIndex != nil {
		return *m.SuccIndex
	}
	return 0
}

func (m *LoginTurntableBonus) GetBonusAmount() int32 {
	if m != nil && m.BonusAmount != nil {
		return *m.BonusAmount
	}
	return 0
}

// 每日签到的签到奖励
type LoginSignInBonus struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Count            *int32       `protobuf:"varint,2,opt,name=count" json:"count,omitempty"`
	Coin             *int32       `protobuf:"varint,3,opt,name=coin" json:"coin,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *LoginSignInBonus) Reset()                    { *m = LoginSignInBonus{} }
func (m *LoginSignInBonus) String() string            { return proto.CompactTextString(m) }
func (*LoginSignInBonus) ProtoMessage()               {}
func (*LoginSignInBonus) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *LoginSignInBonus) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *LoginSignInBonus) GetCount() int32 {
	if m != nil && m.Count != nil {
		return *m.Count
	}
	return 0
}

func (m *LoginSignInBonus) GetCoin() int32 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

// 在线奖励,用户在线的时候,系统倒计时
type OlineBonus struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *OlineBonus) Reset()                    { *m = OlineBonus{} }
func (m *OlineBonus) String() string            { return proto.CompactTextString(m) }
func (*OlineBonus) ProtoMessage()               {}
func (*OlineBonus) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *OlineBonus) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

// 定时奖励,在特定的时间段可以领取特定的奖励
type TimingBonus struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *TimingBonus) Reset()                    { *m = TimingBonus{} }
func (m *TimingBonus) String() string            { return proto.CompactTextString(m) }
func (*TimingBonus) ProtoMessage()               {}
func (*TimingBonus) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

func (m *TimingBonus) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func init() {
	proto.RegisterType((*LoginTurntableBonus)(nil), "bbproto.LoginTurntableBonus")
	proto.RegisterType((*LoginSignInBonus)(nil), "bbproto.LoginSignInBonus")
	proto.RegisterType((*OlineBonus)(nil), "bbproto.OlineBonus")
	proto.RegisterType((*TimingBonus)(nil), "bbproto.TimingBonus")
}

var fileDescriptor4 = []byte{
	// 187 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0xca, 0xcf, 0x2b,
	0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4f, 0x4a, 0x02, 0x33, 0xa4, 0xb8, 0x92,
	0x12, 0x8b, 0x53, 0x21, 0x82, 0x4a, 0xc9, 0x5c, 0xc2, 0x3e, 0xf9, 0xe9, 0x99, 0x79, 0x21, 0xa5,
	0x45, 0x79, 0x25, 0x89, 0x49, 0x39, 0xa9, 0x4e, 0x20, 0x1d, 0x42, 0x2a, 0x5c, 0x6c, 0x19, 0xa9,
	0x89, 0x29, 0xa9, 0x45, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x22, 0x7a, 0x50, 0xcd, 0x7a,
	0x01, 0x20, 0xd2, 0x03, 0x2c, 0x27, 0x24, 0xc8, 0xc5, 0x59, 0x5c, 0x9a, 0x9c, 0xec, 0x99, 0x97,
	0x92, 0x5a, 0x21, 0xc1, 0x04, 0x54, 0xc8, 0x2a, 0x24, 0x0c, 0xb5, 0xd3, 0x31, 0x37, 0xbf, 0x34,
	0xaf, 0x44, 0x82, 0x19, 0x24, 0xa8, 0x14, 0xca, 0x25, 0x00, 0xb6, 0x24, 0x38, 0x33, 0x3d, 0xcf,
	0x33, 0x8f, 0x14, 0x1b, 0x78, 0xb9, 0x58, 0x93, 0xc1, 0x06, 0x41, 0x4c, 0xe7, 0xe1, 0x62, 0x49,
	0xce, 0xcf, 0xcc, 0x83, 0x1a, 0x6b, 0xc4, 0xc5, 0xe5, 0x9f, 0x93, 0x99, 0x47, 0x8a, 0x93, 0x95,
	0x8c, 0xb9, 0xb8, 0x43, 0x32, 0x73, 0x33, 0xf3, 0xd2, 0x49, 0xd0, 0x14, 0xc0, 0x00, 0x08, 0x00,
	0x00, 0xff, 0xff, 0xe5, 0x8f, 0x55, 0x76, 0x49, 0x01, 0x00, 0x00,
}