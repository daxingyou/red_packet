// Code generated by protoc-gen-go.
// source: ddz_server.proto
// DO NOT EDIT!

/*
Package doudizhu is a generated protocol buffer package.

It is generated from these files:
	ddz_server.proto

It has these top-level messages:
	PPokerPai
	POutPokerPais
	PDdzDesk
	PGameData
	PDdzUser
	PDdzRoom
	PDdzbak
	PDdzSession
*/
package doudizhu

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

// 牌 ,单张扑克牌
type PPokerPai struct {
	Id               *int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Des              *string `protobuf:"bytes,2,opt,name=des" json:"des,omitempty"`
	Value            *int32  `protobuf:"varint,3,opt,name=value" json:"value,omitempty"`
	Flower           *int32  `protobuf:"varint,4,opt,name=flower" json:"flower,omitempty"`
	Name             *string `protobuf:"bytes,5,opt,name=name" json:"name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PPokerPai) Reset()                    { *m = PPokerPai{} }
func (m *PPokerPai) String() string            { return proto.CompactTextString(m) }
func (*PPokerPai) ProtoMessage()               {}
func (*PPokerPai) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PPokerPai) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *PPokerPai) GetDes() string {
	if m != nil && m.Des != nil {
		return *m.Des
	}
	return ""
}

func (m *PPokerPai) GetValue() int32 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

func (m *PPokerPai) GetFlower() int32 {
	if m != nil && m.Flower != nil {
		return *m.Flower
	}
	return 0
}

func (m *PPokerPai) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

// 打出去的牌
type POutPokerPais struct {
	KeyValue         *int32       `protobuf:"varint,1,opt,name=keyValue" json:"keyValue,omitempty"`
	PokerPais        []*PPokerPai `protobuf:"bytes,2,rep,name=pokerPais" json:"pokerPais,omitempty"`
	Type             *int32       `protobuf:"varint,3,opt,name=type" json:"type,omitempty"`
	IsBomb           *bool        `protobuf:"varint,4,opt,name=isBomb" json:"isBomb,omitempty"`
	CountDuizi       *int32       `protobuf:"varint,5,opt,name=countDuizi" json:"countDuizi,omitempty"`
	CountSanzhang    *int32       `protobuf:"varint,6,opt,name=countSanzhang" json:"countSanzhang,omitempty"`
	CountSizhang     *int32       `protobuf:"varint,7,opt,name=countSizhang" json:"countSizhang,omitempty"`
	CountYizhang     *int32       `protobuf:"varint,8,opt,name=countYizhang" json:"countYizhang,omitempty"`
	UserId           *uint32      `protobuf:"varint,9,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *POutPokerPais) Reset()                    { *m = POutPokerPais{} }
func (m *POutPokerPais) String() string            { return proto.CompactTextString(m) }
func (*POutPokerPais) ProtoMessage()               {}
func (*POutPokerPais) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *POutPokerPais) GetKeyValue() int32 {
	if m != nil && m.KeyValue != nil {
		return *m.KeyValue
	}
	return 0
}

func (m *POutPokerPais) GetPokerPais() []*PPokerPai {
	if m != nil {
		return m.PokerPais
	}
	return nil
}

func (m *POutPokerPais) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *POutPokerPais) GetIsBomb() bool {
	if m != nil && m.IsBomb != nil {
		return *m.IsBomb
	}
	return false
}

func (m *POutPokerPais) GetCountDuizi() int32 {
	if m != nil && m.CountDuizi != nil {
		return *m.CountDuizi
	}
	return 0
}

func (m *POutPokerPais) GetCountSanzhang() int32 {
	if m != nil && m.CountSanzhang != nil {
		return *m.CountSanzhang
	}
	return 0
}

func (m *POutPokerPais) GetCountSizhang() int32 {
	if m != nil && m.CountSizhang != nil {
		return *m.CountSizhang
	}
	return 0
}

func (m *POutPokerPais) GetCountYizhang() int32 {
	if m != nil && m.CountYizhang != nil {
		return *m.CountYizhang
	}
	return 0
}

func (m *POutPokerPais) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

// desk
type PDdzDesk struct {
	DeskId           *int32         `protobuf:"varint,1,opt,name=deskId" json:"deskId,omitempty"`
	Key              *string        `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
	UserCountLimit   *int32         `protobuf:"varint,3,opt,name=userCountLimit" json:"userCountLimit,omitempty"`
	AllPokerPai      []*PPokerPai   `protobuf:"bytes,4,rep,name=allPokerPai" json:"allPokerPai,omitempty"`
	DiPokerPai       []*PPokerPai   `protobuf:"bytes,5,rep,name=diPokerPai" json:"diPokerPai,omitempty"`
	OutPai           *POutPokerPais `protobuf:"bytes,6,opt,name=outPai" json:"outPai,omitempty"`
	Owner            *uint32        `protobuf:"varint,7,opt,name=owner" json:"owner,omitempty"`
	DizhuPaiUser     *uint32        `protobuf:"varint,8,opt,name=dizhuPaiUser" json:"dizhuPaiUser,omitempty"`
	Dizhu            *uint32        `protobuf:"varint,9,opt,name=dizhu" json:"dizhu,omitempty"`
	ActiveUser       *uint32        `protobuf:"varint,10,opt,name=activeUser" json:"activeUser,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *PDdzDesk) Reset()                    { *m = PDdzDesk{} }
func (m *PDdzDesk) String() string            { return proto.CompactTextString(m) }
func (*PDdzDesk) ProtoMessage()               {}
func (*PDdzDesk) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PDdzDesk) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

func (m *PDdzDesk) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *PDdzDesk) GetUserCountLimit() int32 {
	if m != nil && m.UserCountLimit != nil {
		return *m.UserCountLimit
	}
	return 0
}

func (m *PDdzDesk) GetAllPokerPai() []*PPokerPai {
	if m != nil {
		return m.AllPokerPai
	}
	return nil
}

func (m *PDdzDesk) GetDiPokerPai() []*PPokerPai {
	if m != nil {
		return m.DiPokerPai
	}
	return nil
}

func (m *PDdzDesk) GetOutPai() *POutPokerPais {
	if m != nil {
		return m.OutPai
	}
	return nil
}

func (m *PDdzDesk) GetOwner() uint32 {
	if m != nil && m.Owner != nil {
		return *m.Owner
	}
	return 0
}

func (m *PDdzDesk) GetDizhuPaiUser() uint32 {
	if m != nil && m.DizhuPaiUser != nil {
		return *m.DizhuPaiUser
	}
	return 0
}

func (m *PDdzDesk) GetDizhu() uint32 {
	if m != nil && m.Dizhu != nil {
		return *m.Dizhu
	}
	return 0
}

func (m *PDdzDesk) GetActiveUser() uint32 {
	if m != nil && m.ActiveUser != nil {
		return *m.ActiveUser
	}
	return 0
}

// 游戏数据
type PGameData struct {
	HandPokers       []*PPokerPai `protobuf:"bytes,4,rep,name=handPokers" json:"handPokers,omitempty"`
	BoomPokers       []*PPokerPai `protobuf:"bytes,5,rep,name=boomPokers" json:"boomPokers,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *PGameData) Reset()                    { *m = PGameData{} }
func (m *PGameData) String() string            { return proto.CompactTextString(m) }
func (*PGameData) ProtoMessage()               {}
func (*PGameData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PGameData) GetHandPokers() []*PPokerPai {
	if m != nil {
		return m.HandPokers
	}
	return nil
}

func (m *PGameData) GetBoomPokers() []*PPokerPai {
	if m != nil {
		return m.BoomPokers
	}
	return nil
}

// user
type PDdzUser struct {
	UserId           *uint32        `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	GameData         *PGameData     `protobuf:"bytes,2,opt,name=gameData" json:"gameData,omitempty"`
	Status           *int32         `protobuf:"varint,3,opt,name=status" json:"status,omitempty"`
	IsBreak          *bool          `protobuf:"varint,4,opt,name=isBreak" json:"isBreak,omitempty"`
	IsLeave          *bool          `protobuf:"varint,5,opt,name=isLeave" json:"isLeave,omitempty"`
	OutPaiList       *POutPokerPais `protobuf:"bytes,6,opt,name=outPaiList" json:"outPaiList,omitempty"`
	QiangDiZhuStatus *int32         `protobuf:"varint,7,opt,name=qiangDiZhuStatus" json:"qiangDiZhuStatus,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *PDdzUser) Reset()                    { *m = PDdzUser{} }
func (m *PDdzUser) String() string            { return proto.CompactTextString(m) }
func (*PDdzUser) ProtoMessage()               {}
func (*PDdzUser) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *PDdzUser) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *PDdzUser) GetGameData() *PGameData {
	if m != nil {
		return m.GameData
	}
	return nil
}

func (m *PDdzUser) GetStatus() int32 {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return 0
}

func (m *PDdzUser) GetIsBreak() bool {
	if m != nil && m.IsBreak != nil {
		return *m.IsBreak
	}
	return false
}

func (m *PDdzUser) GetIsLeave() bool {
	if m != nil && m.IsLeave != nil {
		return *m.IsLeave
	}
	return false
}

func (m *PDdzUser) GetOutPaiList() *POutPokerPais {
	if m != nil {
		return m.OutPaiList
	}
	return nil
}

func (m *PDdzUser) GetQiangDiZhuStatus() int32 {
	if m != nil && m.QiangDiZhuStatus != nil {
		return *m.QiangDiZhuStatus
	}
	return 0
}

// room
type PDdzRoom struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *PDdzRoom) Reset()                    { *m = PDdzRoom{} }
func (m *PDdzRoom) String() string            { return proto.CompactTextString(m) }
func (*PDdzRoom) ProtoMessage()               {}
func (*PDdzRoom) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

// 备份专用...
type PDdzbak struct {
	Desk             *PDdzDesk   `protobuf:"bytes,1,opt,name=desk" json:"desk,omitempty"`
	Users            []*PDdzUser `protobuf:"bytes,2,rep,name=users" json:"users,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *PDdzbak) Reset()                    { *m = PDdzbak{} }
func (m *PDdzbak) String() string            { return proto.CompactTextString(m) }
func (*PDdzbak) ProtoMessage()               {}
func (*PDdzbak) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *PDdzbak) GetDesk() *PDdzDesk {
	if m != nil {
		return m.Desk
	}
	return nil
}

func (m *PDdzbak) GetUsers() []*PDdzUser {
	if m != nil {
		return m.Users
	}
	return nil
}

// session
type PDdzSession struct {
	UserId           *uint32 `protobuf:"varint,1,opt,name=UserId" json:"UserId,omitempty"`
	RoomId           *int32  `protobuf:"varint,2,opt,name=RoomId" json:"RoomId,omitempty"`
	DeskId           *int32  `protobuf:"varint,3,opt,name=DeskId" json:"DeskId,omitempty"`
	GameStatus       *int32  `protobuf:"varint,4,opt,name=GameStatus" json:"GameStatus,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PDdzSession) Reset()                    { *m = PDdzSession{} }
func (m *PDdzSession) String() string            { return proto.CompactTextString(m) }
func (*PDdzSession) ProtoMessage()               {}
func (*PDdzSession) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *PDdzSession) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *PDdzSession) GetRoomId() int32 {
	if m != nil && m.RoomId != nil {
		return *m.RoomId
	}
	return 0
}

func (m *PDdzSession) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

func (m *PDdzSession) GetGameStatus() int32 {
	if m != nil && m.GameStatus != nil {
		return *m.GameStatus
	}
	return 0
}

func init() {
	proto.RegisterType((*PPokerPai)(nil), "doudizhu.PPokerPai")
	proto.RegisterType((*POutPokerPais)(nil), "doudizhu.POutPokerPais")
	proto.RegisterType((*PDdzDesk)(nil), "doudizhu.PDdzDesk")
	proto.RegisterType((*PGameData)(nil), "doudizhu.PGameData")
	proto.RegisterType((*PDdzUser)(nil), "doudizhu.PDdzUser")
	proto.RegisterType((*PDdzRoom)(nil), "doudizhu.PDdzRoom")
	proto.RegisterType((*PDdzbak)(nil), "doudizhu.PDdzbak")
	proto.RegisterType((*PDdzSession)(nil), "doudizhu.PDdzSession")
}

var fileDescriptor0 = []byte{
	// 526 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x53, 0xdd, 0x8a, 0xd3, 0x40,
	0x14, 0xa6, 0x3f, 0x69, 0xd3, 0x93, 0xcd, 0x5a, 0xc6, 0xbf, 0x5c, 0xae, 0x01, 0xb5, 0x20, 0xf4,
	0xa2, 0x8f, 0xa0, 0x01, 0x59, 0x28, 0x5a, 0x2d, 0x2b, 0x28, 0x88, 0x4c, 0x77, 0xc6, 0xed, 0x90,
	0x36, 0x53, 0x33, 0x93, 0x2e, 0xdb, 0x5b, 0xdf, 0xca, 0x87, 0xf0, 0x99, 0x3c, 0x73, 0x26, 0x59,
	0x53, 0x61, 0xdd, 0xbb, 0xcc, 0x99, 0x93, 0xf3, 0xfd, 0x9c, 0x6f, 0x60, 0x2c, 0xc4, 0xe1, 0x9b,
	0x91, 0xe5, 0x5e, 0x96, 0xd3, 0x5d, 0xa9, 0xad, 0x66, 0xa1, 0xd0, 0x95, 0x50, 0x87, 0x75, 0x95,
	0x7e, 0x80, 0xd1, 0x62, 0xa1, 0x73, 0x59, 0x2e, 0xb8, 0x62, 0x00, 0x5d, 0x25, 0x92, 0xce, 0x59,
	0x67, 0x12, 0xb0, 0x08, 0x7a, 0x42, 0x9a, 0xa4, 0x8b, 0x87, 0x11, 0x8b, 0x21, 0xd8, 0xf3, 0x4d,
	0x25, 0x93, 0x1e, 0xdd, 0x9d, 0xc2, 0xe0, 0xfb, 0x46, 0x5f, 0xcb, 0x32, 0xe9, 0xd3, 0xf9, 0x04,
	0xfa, 0x05, 0xdf, 0xca, 0x24, 0x70, 0xcd, 0xe9, 0xef, 0x0e, 0xc4, 0x8b, 0xf7, 0x95, 0x6d, 0xc6,
	0x1a, 0x36, 0x86, 0x30, 0x97, 0x37, 0x9f, 0x68, 0x82, 0x9f, 0xfe, 0x02, 0x46, 0xbb, 0xe6, 0x1a,
	0x31, 0x7a, 0x93, 0x68, 0xf6, 0x70, 0xda, 0x90, 0x9a, 0xfe, 0x65, 0x84, 0x93, 0xed, 0xcd, 0xae,
	0x85, 0xab, 0xcc, 0x6b, 0xbd, 0x5d, 0x11, 0x6e, 0xc8, 0x90, 0xf0, 0xa5, 0xae, 0x0a, 0x9b, 0x55,
	0xea, 0xa0, 0x08, 0x3d, 0x60, 0x8f, 0x21, 0xa6, 0xda, 0x92, 0x17, 0x87, 0x35, 0x2f, 0xae, 0x92,
	0x01, 0x95, 0x1f, 0xc1, 0x89, 0x2f, 0x2b, 0x5f, 0x1d, 0x1e, 0x55, 0x3f, 0xd7, 0xd5, 0xb0, 0x81,
	0xa9, 0xd0, 0xae, 0x73, 0x91, 0x8c, 0xf0, 0x1c, 0xa7, 0x3f, 0xbb, 0x10, 0x2e, 0x32, 0x71, 0xc8,
	0xa4, 0xc9, 0xdd, 0x25, 0xfa, 0x92, 0x9f, 0xb7, 0x7c, 0x42, 0x6d, 0xb5, 0x4f, 0x4f, 0xe0, 0xd4,
	0xfd, 0xf9, 0xc6, 0xcd, 0x9c, 0xab, 0xad, 0xb2, 0x35, 0xf1, 0x09, 0x44, 0x7c, 0xb3, 0x69, 0x54,
	0x21, 0xfb, 0x3b, 0x05, 0xbf, 0x04, 0x10, 0xea, 0xb6, 0x31, 0xf8, 0x5f, 0xe3, 0x40, 0xa3, 0xc7,
	0xd8, 0xe4, 0x04, 0x46, 0xb3, 0xa7, 0xad, 0xa6, 0x23, 0xf3, 0x71, 0x77, 0xfa, 0xba, 0xc0, 0x5d,
	0x39, 0xc9, 0xb1, 0x93, 0x4c, 0x5d, 0x78, 0x77, 0x81, 0x54, 0x49, 0x72, 0xec, 0x9a, 0xa8, 0xea,
	0x15, 0x3b, 0x63, 0xf9, 0xa5, 0x55, 0x7b, 0x49, 0x2d, 0x40, 0x2e, 0x7c, 0xc5, 0xa4, 0xbc, 0xc5,
	0x2d, 0x67, 0xdc, 0x72, 0x47, 0x13, 0x0d, 0x13, 0x84, 0x62, 0xee, 0xd1, 0xb3, 0xd2, 0x7a, 0x5b,
	0x37, 0xde, 0xad, 0x27, 0xfd, 0xd5, 0xf1, 0x26, 0x3b, 0xc4, 0xd6, 0x06, 0x3a, 0xc4, 0xe7, 0x39,
	0x84, 0x57, 0x35, 0x34, 0x39, 0x7d, 0x3c, 0xe3, 0x96, 0x15, 0xfe, 0x66, 0x2c, 0xb7, 0x95, 0xa9,
	0x6d, 0x7f, 0x00, 0x43, 0xcc, 0x4b, 0x29, 0x79, 0x5e, 0x07, 0x86, 0x0a, 0x73, 0xc9, 0xf7, 0x3e,
	0xab, 0x21, 0x7b, 0x05, 0xe0, 0x5d, 0x9c, 0x2b, 0x63, 0xef, 0x73, 0x32, 0x81, 0xf1, 0x0f, 0x85,
	0x31, 0xc9, 0xd4, 0x97, 0x75, 0xb5, 0xf4, 0x40, 0x94, 0xa3, 0x14, 0x3c, 0xf7, 0x8f, 0xa8, 0x34,
	0x7d, 0x07, 0x43, 0xf7, 0xbd, 0xe2, 0x39, 0x3b, 0x83, 0xbe, 0xcb, 0x0a, 0x89, 0x88, 0x66, 0xac,
	0x35, 0xb7, 0x49, 0xd3, 0x33, 0x08, 0x9c, 0xd0, 0xe6, 0x0d, 0xfc, 0xd3, 0xe2, 0xbc, 0xc0, 0x17,
	0x1a, 0xb9, 0xef, 0xa5, 0x34, 0x46, 0xe9, 0xc2, 0x69, 0xbc, 0x68, 0x5b, 0x83, 0x67, 0x07, 0x8b,
	0xe7, 0x6e, 0x13, 0xde, 0xcc, 0xe7, 0xd3, 0x7b, 0x80, 0xab, 0x74, 0xfe, 0xd4, 0x74, 0xe9, 0xbd,
	0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0x25, 0x17, 0xc0, 0xc3, 0x11, 0x04, 0x00, 0x00,
}
