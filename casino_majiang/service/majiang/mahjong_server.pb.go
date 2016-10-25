// Code generated by protoc-gen-go.
// source: mahjong_server.proto
// DO NOT EDIT!

/*
Package majiang is a generated protocol buffer package.

It is generated from these files:
	mahjong_server.proto

It has these top-level messages:
	MJPai
	MJHandPai
	HuPaiInfo
	GangPaiInfo
	PlayerGameData
	CheckBean
	CheckCase
	MjDesk
	BillBean
	Bill
	StatiscRound
	MjUserStatisc
	MjUser
	MjSession
	MjRoom
	RunningDeskKeys
*/
package majiang

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

// 麻将牌的结构
type MJPai struct {
	Index            *int32  `protobuf:"varint,1,opt,name=index" json:"index,omitempty"`
	Flower           *int32  `protobuf:"varint,2,opt,name=flower" json:"flower,omitempty"`
	Value            *int32  `protobuf:"varint,3,opt,name=value" json:"value,omitempty"`
	Des              *string `protobuf:"bytes,4,opt,name=des" json:"des,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MJPai) Reset()                    { *m = MJPai{} }
func (m *MJPai) String() string            { return proto.CompactTextString(m) }
func (*MJPai) ProtoMessage()               {}
func (*MJPai) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *MJPai) GetIndex() int32 {
	if m != nil && m.Index != nil {
		return *m.Index
	}
	return 0
}

func (m *MJPai) GetFlower() int32 {
	if m != nil && m.Flower != nil {
		return *m.Flower
	}
	return 0
}

func (m *MJPai) GetValue() int32 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

func (m *MJPai) GetDes() string {
	if m != nil && m.Des != nil {
		return *m.Des
	}
	return ""
}

// 手里的一副牌
type MJHandPai struct {
	Pais             []*MJPai `protobuf:"bytes,1,rep,name=Pais" json:"Pais,omitempty"`
	PengPais         []*MJPai `protobuf:"bytes,2,rep,name=PengPais" json:"PengPais,omitempty"`
	GangPais         []*MJPai `protobuf:"bytes,3,rep,name=GangPais" json:"GangPais,omitempty"`
	HuPais           []*MJPai `protobuf:"bytes,4,rep,name=HuPais" json:"HuPais,omitempty"`
	InPai            *MJPai   `protobuf:"bytes,5,opt,name=inPai" json:"inPai,omitempty"`
	QueFlower        *int32   `protobuf:"varint,6,opt,name=queFlower" json:"queFlower,omitempty"`
	OutPais          []*MJPai `protobuf:"bytes,7,rep,name=OutPais" json:"OutPais,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *MJHandPai) Reset()                    { *m = MJHandPai{} }
func (m *MJHandPai) String() string            { return proto.CompactTextString(m) }
func (*MJHandPai) ProtoMessage()               {}
func (*MJHandPai) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *MJHandPai) GetPais() []*MJPai {
	if m != nil {
		return m.Pais
	}
	return nil
}

func (m *MJHandPai) GetPengPais() []*MJPai {
	if m != nil {
		return m.PengPais
	}
	return nil
}

func (m *MJHandPai) GetGangPais() []*MJPai {
	if m != nil {
		return m.GangPais
	}
	return nil
}

func (m *MJHandPai) GetHuPais() []*MJPai {
	if m != nil {
		return m.HuPais
	}
	return nil
}

func (m *MJHandPai) GetInPai() *MJPai {
	if m != nil {
		return m.InPai
	}
	return nil
}

func (m *MJHandPai) GetQueFlower() int32 {
	if m != nil && m.QueFlower != nil {
		return *m.QueFlower
	}
	return 0
}

func (m *MJHandPai) GetOutPais() []*MJPai {
	if m != nil {
		return m.OutPais
	}
	return nil
}

type HuPaiInfo struct {
	SendUserId       *uint32 `protobuf:"varint,1,opt,name=sendUserId" json:"sendUserId,omitempty"`
	GetUserId        *uint32 `protobuf:"varint,2,opt,name=getUserId" json:"getUserId,omitempty"`
	ByWho            *int32  `protobuf:"varint,3,opt,name=byWho" json:"byWho,omitempty"`
	HuType           *int32  `protobuf:"varint,4,opt,name=huType" json:"huType,omitempty"`
	CardType         *int32  `protobuf:"varint,5,opt,name=cardType" json:"cardType,omitempty"`
	HuDesc           *string `protobuf:"bytes,6,opt,name=huDesc" json:"huDesc,omitempty"`
	Pai              *MJPai  `protobuf:"bytes,7,opt,name=pai" json:"pai,omitempty"`
	Fan              *int32  `protobuf:"varint,8,opt,name=fan" json:"fan,omitempty"`
	Score            *int64  `protobuf:"varint,9,opt,name=score" json:"score,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *HuPaiInfo) Reset()                    { *m = HuPaiInfo{} }
func (m *HuPaiInfo) String() string            { return proto.CompactTextString(m) }
func (*HuPaiInfo) ProtoMessage()               {}
func (*HuPaiInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *HuPaiInfo) GetSendUserId() uint32 {
	if m != nil && m.SendUserId != nil {
		return *m.SendUserId
	}
	return 0
}

func (m *HuPaiInfo) GetGetUserId() uint32 {
	if m != nil && m.GetUserId != nil {
		return *m.GetUserId
	}
	return 0
}

func (m *HuPaiInfo) GetByWho() int32 {
	if m != nil && m.ByWho != nil {
		return *m.ByWho
	}
	return 0
}

func (m *HuPaiInfo) GetHuType() int32 {
	if m != nil && m.HuType != nil {
		return *m.HuType
	}
	return 0
}

func (m *HuPaiInfo) GetCardType() int32 {
	if m != nil && m.CardType != nil {
		return *m.CardType
	}
	return 0
}

func (m *HuPaiInfo) GetHuDesc() string {
	if m != nil && m.HuDesc != nil {
		return *m.HuDesc
	}
	return ""
}

func (m *HuPaiInfo) GetPai() *MJPai {
	if m != nil {
		return m.Pai
	}
	return nil
}

func (m *HuPaiInfo) GetFan() int32 {
	if m != nil && m.Fan != nil {
		return *m.Fan
	}
	return 0
}

func (m *HuPaiInfo) GetScore() int64 {
	if m != nil && m.Score != nil {
		return *m.Score
	}
	return 0
}

type GangPaiInfo struct {
	SendUserId       *uint32 `protobuf:"varint,1,opt,name=sendUserId" json:"sendUserId,omitempty"`
	GetUserId        *uint32 `protobuf:"varint,2,opt,name=getUserId" json:"getUserId,omitempty"`
	ByWho            *int32  `protobuf:"varint,3,opt,name=byWho" json:"byWho,omitempty"`
	GangType         *int32  `protobuf:"varint,4,opt,name=gangType" json:"gangType,omitempty"`
	Pai              *MJPai  `protobuf:"bytes,5,opt,name=pai" json:"pai,omitempty"`
	Transfer         *int32  `protobuf:"varint,6,opt,name=transfer" json:"transfer,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GangPaiInfo) Reset()                    { *m = GangPaiInfo{} }
func (m *GangPaiInfo) String() string            { return proto.CompactTextString(m) }
func (*GangPaiInfo) ProtoMessage()               {}
func (*GangPaiInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GangPaiInfo) GetSendUserId() uint32 {
	if m != nil && m.SendUserId != nil {
		return *m.SendUserId
	}
	return 0
}

func (m *GangPaiInfo) GetGetUserId() uint32 {
	if m != nil && m.GetUserId != nil {
		return *m.GetUserId
	}
	return 0
}

func (m *GangPaiInfo) GetByWho() int32 {
	if m != nil && m.ByWho != nil {
		return *m.ByWho
	}
	return 0
}

func (m *GangPaiInfo) GetGangType() int32 {
	if m != nil && m.GangType != nil {
		return *m.GangType
	}
	return 0
}

func (m *GangPaiInfo) GetPai() *MJPai {
	if m != nil {
		return m.Pai
	}
	return nil
}

func (m *GangPaiInfo) GetTransfer() int32 {
	if m != nil && m.Transfer != nil {
		return *m.Transfer
	}
	return 0
}

// 一个玩家的游戏信息
type PlayerGameData struct {
	HandPai          *MJHandPai     `protobuf:"bytes,1,opt,name=handPai" json:"handPai,omitempty"`
	HuInfo           []*HuPaiInfo   `protobuf:"bytes,2,rep,name=huInfo" json:"huInfo,omitempty"`
	DianHuInfo       []*HuPaiInfo   `protobuf:"bytes,3,rep,name=dianHuInfo" json:"dianHuInfo,omitempty"`
	GangInfo         []*GangPaiInfo `protobuf:"bytes,4,rep,name=gangInfo" json:"gangInfo,omitempty"`
	DianGangInfo     []*GangPaiInfo `protobuf:"bytes,5,rep,name=dianGangInfo" json:"dianGangInfo,omitempty"`
	GangTransfer     []*GangPaiInfo `protobuf:"bytes,6,rep,name=gangTransfer" json:"gangTransfer,omitempty"`
	TotalScore       *int32         `protobuf:"varint,7,opt,name=totalScore" json:"totalScore,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *PlayerGameData) Reset()                    { *m = PlayerGameData{} }
func (m *PlayerGameData) String() string            { return proto.CompactTextString(m) }
func (*PlayerGameData) ProtoMessage()               {}
func (*PlayerGameData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *PlayerGameData) GetHandPai() *MJHandPai {
	if m != nil {
		return m.HandPai
	}
	return nil
}

func (m *PlayerGameData) GetHuInfo() []*HuPaiInfo {
	if m != nil {
		return m.HuInfo
	}
	return nil
}

func (m *PlayerGameData) GetDianHuInfo() []*HuPaiInfo {
	if m != nil {
		return m.DianHuInfo
	}
	return nil
}

func (m *PlayerGameData) GetGangInfo() []*GangPaiInfo {
	if m != nil {
		return m.GangInfo
	}
	return nil
}

func (m *PlayerGameData) GetDianGangInfo() []*GangPaiInfo {
	if m != nil {
		return m.DianGangInfo
	}
	return nil
}

func (m *PlayerGameData) GetGangTransfer() []*GangPaiInfo {
	if m != nil {
		return m.GangTransfer
	}
	return nil
}

func (m *PlayerGameData) GetTotalScore() int32 {
	if m != nil && m.TotalScore != nil {
		return *m.TotalScore
	}
	return 0
}

// 需要确认的时间
type CheckBean struct {
	UserId           *uint32 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	CanPeng          *bool   `protobuf:"varint,2,opt,name=CanPeng" json:"CanPeng,omitempty"`
	CanGang          *bool   `protobuf:"varint,3,opt,name=CanGang" json:"CanGang,omitempty"`
	CanHu            *bool   `protobuf:"varint,4,opt,name=CanHu" json:"CanHu,omitempty"`
	CheckStatus      *int32  `protobuf:"varint,5,opt,name=CheckStatus" json:"CheckStatus,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CheckBean) Reset()                    { *m = CheckBean{} }
func (m *CheckBean) String() string            { return proto.CompactTextString(m) }
func (*CheckBean) ProtoMessage()               {}
func (*CheckBean) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *CheckBean) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *CheckBean) GetCanPeng() bool {
	if m != nil && m.CanPeng != nil {
		return *m.CanPeng
	}
	return false
}

func (m *CheckBean) GetCanGang() bool {
	if m != nil && m.CanGang != nil {
		return *m.CanGang
	}
	return false
}

func (m *CheckBean) GetCanHu() bool {
	if m != nil && m.CanHu != nil {
		return *m.CanHu
	}
	return false
}

func (m *CheckBean) GetCheckStatus() int32 {
	if m != nil && m.CheckStatus != nil {
		return *m.CheckStatus
	}
	return 0
}

type CheckCase struct {
	CheckB           []*CheckBean `protobuf:"bytes,1,rep,name=CheckB" json:"CheckB,omitempty"`
	CheckStatus      *int32       `protobuf:"varint,2,opt,name=CheckStatus" json:"CheckStatus,omitempty"`
	CheckMJPai       *MJPai       `protobuf:"bytes,3,opt,name=CheckMJPai" json:"CheckMJPai,omitempty"`
	UserIdOut        *uint32      `protobuf:"varint,4,opt,name=UserIdOut" json:"UserIdOut,omitempty"`
	PreOutGangInfo   *GangPaiInfo `protobuf:"bytes,5,opt,name=PreOutGangInfo" json:"PreOutGangInfo,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *CheckCase) Reset()                    { *m = CheckCase{} }
func (m *CheckCase) String() string            { return proto.CompactTextString(m) }
func (*CheckCase) ProtoMessage()               {}
func (*CheckCase) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *CheckCase) GetCheckB() []*CheckBean {
	if m != nil {
		return m.CheckB
	}
	return nil
}

func (m *CheckCase) GetCheckStatus() int32 {
	if m != nil && m.CheckStatus != nil {
		return *m.CheckStatus
	}
	return 0
}

func (m *CheckCase) GetCheckMJPai() *MJPai {
	if m != nil {
		return m.CheckMJPai
	}
	return nil
}

func (m *CheckCase) GetUserIdOut() uint32 {
	if m != nil && m.UserIdOut != nil {
		return *m.UserIdOut
	}
	return 0
}

func (m *CheckCase) GetPreOutGangInfo() *GangPaiInfo {
	if m != nil {
		return m.PreOutGangInfo
	}
	return nil
}

// 麻将牌的结构
type MjDesk struct {
	DeskId           *int32     `protobuf:"varint,1,opt,name=DeskId" json:"DeskId,omitempty"`
	RoomId           *int32     `protobuf:"varint,2,opt,name=RoomId" json:"RoomId,omitempty"`
	Status           *int32     `protobuf:"varint,3,opt,name=Status" json:"Status,omitempty"`
	Users            []*MjUser  `protobuf:"bytes,4,rep,name=users" json:"users,omitempty"`
	Password         *string    `protobuf:"bytes,5,opt,name=Password" json:"Password,omitempty"`
	Owner            *uint32    `protobuf:"varint,6,opt,name=Owner" json:"Owner,omitempty"`
	CreateFee        *int64     `protobuf:"varint,7,opt,name=CreateFee" json:"CreateFee,omitempty"`
	MjRoomType       *int32     `protobuf:"varint,8,opt,name=mjRoomType" json:"mjRoomType,omitempty"`
	BoardsCout       *int32     `protobuf:"varint,9,opt,name=boardsCout" json:"boardsCout,omitempty"`
	CapMax           *int64     `protobuf:"varint,10,opt,name=capMax" json:"capMax,omitempty"`
	CardsNum         *int32     `protobuf:"varint,11,opt,name=cardsNum" json:"cardsNum,omitempty"`
	Settlement       *int32     `protobuf:"varint,12,opt,name=settlement" json:"settlement,omitempty"`
	BaseValue        *int64     `protobuf:"varint,13,opt,name=baseValue" json:"baseValue,omitempty"`
	ZiMoRadio        *int32     `protobuf:"varint,14,opt,name=ziMoRadio" json:"ziMoRadio,omitempty"`
	DianGangHuaRadio *int32     `protobuf:"varint,15,opt,name=dianGangHuaRadio" json:"dianGangHuaRadio,omitempty"`
	OthersCheckBox   []int32    `protobuf:"varint,16,rep,name=othersCheckBox" json:"othersCheckBox,omitempty"`
	HuRadio          *int32     `protobuf:"varint,17,opt,name=huRadio" json:"huRadio,omitempty"`
	AllMJPai         []*MJPai   `protobuf:"bytes,18,rep,name=AllMJPai" json:"AllMJPai,omitempty"`
	DeskMJPai        []*MJPai   `protobuf:"bytes,19,rep,name=DeskMJPai" json:"DeskMJPai,omitempty"`
	MJPaiCursor      *int32     `protobuf:"varint,20,opt,name=MJPaiCursor" json:"MJPaiCursor,omitempty"`
	TotalPlayCount   *int32     `protobuf:"varint,21,opt,name=TotalPlayCount" json:"TotalPlayCount,omitempty"`
	CurrPlayCount    *int32     `protobuf:"varint,22,opt,name=CurrPlayCount" json:"CurrPlayCount,omitempty"`
	Banker           *uint32    `protobuf:"varint,23,opt,name=Banker" json:"Banker,omitempty"`
	CheckCase        *CheckCase `protobuf:"bytes,24,opt,name=CheckCase" json:"CheckCase,omitempty"`
	ActiveUser       *uint32    `protobuf:"varint,25,opt,name=ActiveUser" json:"ActiveUser,omitempty"`
	GameNumber       *int32     `protobuf:"varint,26,opt,name=GameNumber" json:"GameNumber,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *MjDesk) Reset()                    { *m = MjDesk{} }
func (m *MjDesk) String() string            { return proto.CompactTextString(m) }
func (*MjDesk) ProtoMessage()               {}
func (*MjDesk) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *MjDesk) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

func (m *MjDesk) GetRoomId() int32 {
	if m != nil && m.RoomId != nil {
		return *m.RoomId
	}
	return 0
}

func (m *MjDesk) GetStatus() int32 {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return 0
}

func (m *MjDesk) GetUsers() []*MjUser {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *MjDesk) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

func (m *MjDesk) GetOwner() uint32 {
	if m != nil && m.Owner != nil {
		return *m.Owner
	}
	return 0
}

func (m *MjDesk) GetCreateFee() int64 {
	if m != nil && m.CreateFee != nil {
		return *m.CreateFee
	}
	return 0
}

func (m *MjDesk) GetMjRoomType() int32 {
	if m != nil && m.MjRoomType != nil {
		return *m.MjRoomType
	}
	return 0
}

func (m *MjDesk) GetBoardsCout() int32 {
	if m != nil && m.BoardsCout != nil {
		return *m.BoardsCout
	}
	return 0
}

func (m *MjDesk) GetCapMax() int64 {
	if m != nil && m.CapMax != nil {
		return *m.CapMax
	}
	return 0
}

func (m *MjDesk) GetCardsNum() int32 {
	if m != nil && m.CardsNum != nil {
		return *m.CardsNum
	}
	return 0
}

func (m *MjDesk) GetSettlement() int32 {
	if m != nil && m.Settlement != nil {
		return *m.Settlement
	}
	return 0
}

func (m *MjDesk) GetBaseValue() int64 {
	if m != nil && m.BaseValue != nil {
		return *m.BaseValue
	}
	return 0
}

func (m *MjDesk) GetZiMoRadio() int32 {
	if m != nil && m.ZiMoRadio != nil {
		return *m.ZiMoRadio
	}
	return 0
}

func (m *MjDesk) GetDianGangHuaRadio() int32 {
	if m != nil && m.DianGangHuaRadio != nil {
		return *m.DianGangHuaRadio
	}
	return 0
}

func (m *MjDesk) GetOthersCheckBox() []int32 {
	if m != nil {
		return m.OthersCheckBox
	}
	return nil
}

func (m *MjDesk) GetHuRadio() int32 {
	if m != nil && m.HuRadio != nil {
		return *m.HuRadio
	}
	return 0
}

func (m *MjDesk) GetAllMJPai() []*MJPai {
	if m != nil {
		return m.AllMJPai
	}
	return nil
}

func (m *MjDesk) GetDeskMJPai() []*MJPai {
	if m != nil {
		return m.DeskMJPai
	}
	return nil
}

func (m *MjDesk) GetMJPaiCursor() int32 {
	if m != nil && m.MJPaiCursor != nil {
		return *m.MJPaiCursor
	}
	return 0
}

func (m *MjDesk) GetTotalPlayCount() int32 {
	if m != nil && m.TotalPlayCount != nil {
		return *m.TotalPlayCount
	}
	return 0
}

func (m *MjDesk) GetCurrPlayCount() int32 {
	if m != nil && m.CurrPlayCount != nil {
		return *m.CurrPlayCount
	}
	return 0
}

func (m *MjDesk) GetBanker() uint32 {
	if m != nil && m.Banker != nil {
		return *m.Banker
	}
	return 0
}

func (m *MjDesk) GetCheckCase() *CheckCase {
	if m != nil {
		return m.CheckCase
	}
	return nil
}

func (m *MjDesk) GetActiveUser() uint32 {
	if m != nil && m.ActiveUser != nil {
		return *m.ActiveUser
	}
	return 0
}

func (m *MjDesk) GetGameNumber() int32 {
	if m != nil && m.GameNumber != nil {
		return *m.GameNumber
	}
	return 0
}

type BillBean struct {
	UserId           *uint32 `protobuf:"varint,1,opt,name=UserId" json:"UserId,omitempty"`
	OutUserId        *uint32 `protobuf:"varint,2,opt,name=OutUserId" json:"OutUserId,omitempty"`
	Type             *int32  `protobuf:"varint,3,opt,name=type" json:"type,omitempty"`
	Des              *string `protobuf:"bytes,4,opt,name=des" json:"des,omitempty"`
	Amount           *int64  `protobuf:"varint,5,opt,name=amount" json:"amount,omitempty"`
	Pai              *MJPai  `protobuf:"bytes,6,opt,name=pai" json:"pai,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *BillBean) Reset()                    { *m = BillBean{} }
func (m *BillBean) String() string            { return proto.CompactTextString(m) }
func (*BillBean) ProtoMessage()               {}
func (*BillBean) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *BillBean) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *BillBean) GetOutUserId() uint32 {
	if m != nil && m.OutUserId != nil {
		return *m.OutUserId
	}
	return 0
}

func (m *BillBean) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *BillBean) GetDes() string {
	if m != nil && m.Des != nil {
		return *m.Des
	}
	return ""
}

func (m *BillBean) GetAmount() int64 {
	if m != nil && m.Amount != nil {
		return *m.Amount
	}
	return 0
}

func (m *BillBean) GetPai() *MJPai {
	if m != nil {
		return m.Pai
	}
	return nil
}

type Bill struct {
	WinAmount        *int64      `protobuf:"varint,1,opt,name=WinAmount" json:"WinAmount,omitempty"`
	Bills            []*BillBean `protobuf:"bytes,2,rep,name=bills" json:"bills,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *Bill) Reset()                    { *m = Bill{} }
func (m *Bill) String() string            { return proto.CompactTextString(m) }
func (*Bill) ProtoMessage()               {}
func (*Bill) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *Bill) GetWinAmount() int64 {
	if m != nil && m.WinAmount != nil {
		return *m.WinAmount
	}
	return 0
}

func (m *Bill) GetBills() []*BillBean {
	if m != nil {
		return m.Bills
	}
	return nil
}

type StatiscRound struct {
	Round            *int32  `protobuf:"varint,1,opt,name=Round" json:"Round,omitempty"`
	GameNumber       *int32  `protobuf:"varint,2,opt,name=GameNumber" json:"GameNumber,omitempty"`
	Result           *string `protobuf:"bytes,3,opt,name=Result" json:"Result,omitempty"`
	WinAmount        *int64  `protobuf:"varint,4,opt,name=WinAmount" json:"WinAmount,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *StatiscRound) Reset()                    { *m = StatiscRound{} }
func (m *StatiscRound) String() string            { return proto.CompactTextString(m) }
func (*StatiscRound) ProtoMessage()               {}
func (*StatiscRound) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *StatiscRound) GetRound() int32 {
	if m != nil && m.Round != nil {
		return *m.Round
	}
	return 0
}

func (m *StatiscRound) GetGameNumber() int32 {
	if m != nil && m.GameNumber != nil {
		return *m.GameNumber
	}
	return 0
}

func (m *StatiscRound) GetResult() string {
	if m != nil && m.Result != nil {
		return *m.Result
	}
	return ""
}

func (m *StatiscRound) GetWinAmount() int64 {
	if m != nil && m.WinAmount != nil {
		return *m.WinAmount
	}
	return 0
}

type MjUserStatisc struct {
	WinCoin          *int64          `protobuf:"varint,1,opt,name=WinCoin" json:"WinCoin,omitempty"`
	CountHu          *int32          `protobuf:"varint,2,opt,name=CountHu" json:"CountHu,omitempty"`
	CountZiMo        *int32          `protobuf:"varint,3,opt,name=CountZiMo" json:"CountZiMo,omitempty"`
	CountDianPao     *int32          `protobuf:"varint,4,opt,name=CountDianPao" json:"CountDianPao,omitempty"`
	CountAnGang      *int32          `protobuf:"varint,5,opt,name=CountAnGang" json:"CountAnGang,omitempty"`
	CountMingGang    *int32          `protobuf:"varint,6,opt,name=CountMingGang" json:"CountMingGang,omitempty"`
	CountDianGang    *int32          `protobuf:"varint,7,opt,name=CountDianGang" json:"CountDianGang,omitempty"`
	CountChaJiao     *int32          `protobuf:"varint,8,opt,name=CountChaJiao" json:"CountChaJiao,omitempty"`
	RoundBean        []*StatiscRound `protobuf:"bytes,9,rep,name=roundBean" json:"roundBean,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *MjUserStatisc) Reset()                    { *m = MjUserStatisc{} }
func (m *MjUserStatisc) String() string            { return proto.CompactTextString(m) }
func (*MjUserStatisc) ProtoMessage()               {}
func (*MjUserStatisc) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *MjUserStatisc) GetWinCoin() int64 {
	if m != nil && m.WinCoin != nil {
		return *m.WinCoin
	}
	return 0
}

func (m *MjUserStatisc) GetCountHu() int32 {
	if m != nil && m.CountHu != nil {
		return *m.CountHu
	}
	return 0
}

func (m *MjUserStatisc) GetCountZiMo() int32 {
	if m != nil && m.CountZiMo != nil {
		return *m.CountZiMo
	}
	return 0
}

func (m *MjUserStatisc) GetCountDianPao() int32 {
	if m != nil && m.CountDianPao != nil {
		return *m.CountDianPao
	}
	return 0
}

func (m *MjUserStatisc) GetCountAnGang() int32 {
	if m != nil && m.CountAnGang != nil {
		return *m.CountAnGang
	}
	return 0
}

func (m *MjUserStatisc) GetCountMingGang() int32 {
	if m != nil && m.CountMingGang != nil {
		return *m.CountMingGang
	}
	return 0
}

func (m *MjUserStatisc) GetCountDianGang() int32 {
	if m != nil && m.CountDianGang != nil {
		return *m.CountDianGang
	}
	return 0
}

func (m *MjUserStatisc) GetCountChaJiao() int32 {
	if m != nil && m.CountChaJiao != nil {
		return *m.CountChaJiao
	}
	return 0
}

func (m *MjUserStatisc) GetRoundBean() []*StatiscRound {
	if m != nil {
		return m.RoundBean
	}
	return nil
}

type MjUser struct {
	UserId           *uint32         `protobuf:"varint,1,opt,name=UserId" json:"UserId,omitempty"`
	Coin             *int64          `protobuf:"varint,2,opt,name=Coin" json:"Coin,omitempty"`
	GameData         *PlayerGameData `protobuf:"bytes,3,opt,name=GameData" json:"GameData,omitempty"`
	Status           *int32          `protobuf:"varint,4,opt,name=Status" json:"Status,omitempty"`
	IsBreak          *bool           `protobuf:"varint,5,opt,name=IsBreak" json:"IsBreak,omitempty"`
	IsLeave          *bool           `protobuf:"varint,6,opt,name=IsLeave" json:"IsLeave,omitempty"`
	DeskId           *int32          `protobuf:"varint,7,opt,name=DeskId" json:"DeskId,omitempty"`
	RoomId           *int32          `protobuf:"varint,8,opt,name=RoomId" json:"RoomId,omitempty"`
	DingQue          *bool           `protobuf:"varint,9,opt,name=DingQue" json:"DingQue,omitempty"`
	Exchanged        *bool           `protobuf:"varint,10,opt,name=Exchanged" json:"Exchanged,omitempty"`
	Ready            *bool           `protobuf:"varint,11,opt,name=Ready" json:"Ready,omitempty"`
	IsBanker         *bool           `protobuf:"varint,12,opt,name=IsBanker" json:"IsBanker,omitempty"`
	PreMoGangInfo    *GangPaiInfo    `protobuf:"bytes,13,opt,name=PreMoGangInfo" json:"PreMoGangInfo,omitempty"`
	Bill             *Bill           `protobuf:"bytes,14,opt,name=bill" json:"bill,omitempty"`
	Statisc          *MjUserStatisc  `protobuf:"bytes,15,opt,name=statisc" json:"statisc,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *MjUser) Reset()                    { *m = MjUser{} }
func (m *MjUser) String() string            { return proto.CompactTextString(m) }
func (*MjUser) ProtoMessage()               {}
func (*MjUser) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *MjUser) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *MjUser) GetCoin() int64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *MjUser) GetGameData() *PlayerGameData {
	if m != nil {
		return m.GameData
	}
	return nil
}

func (m *MjUser) GetStatus() int32 {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return 0
}

func (m *MjUser) GetIsBreak() bool {
	if m != nil && m.IsBreak != nil {
		return *m.IsBreak
	}
	return false
}

func (m *MjUser) GetIsLeave() bool {
	if m != nil && m.IsLeave != nil {
		return *m.IsLeave
	}
	return false
}

func (m *MjUser) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

func (m *MjUser) GetRoomId() int32 {
	if m != nil && m.RoomId != nil {
		return *m.RoomId
	}
	return 0
}

func (m *MjUser) GetDingQue() bool {
	if m != nil && m.DingQue != nil {
		return *m.DingQue
	}
	return false
}

func (m *MjUser) GetExchanged() bool {
	if m != nil && m.Exchanged != nil {
		return *m.Exchanged
	}
	return false
}

func (m *MjUser) GetReady() bool {
	if m != nil && m.Ready != nil {
		return *m.Ready
	}
	return false
}

func (m *MjUser) GetIsBanker() bool {
	if m != nil && m.IsBanker != nil {
		return *m.IsBanker
	}
	return false
}

func (m *MjUser) GetPreMoGangInfo() *GangPaiInfo {
	if m != nil {
		return m.PreMoGangInfo
	}
	return nil
}

func (m *MjUser) GetBill() *Bill {
	if m != nil {
		return m.Bill
	}
	return nil
}

func (m *MjUser) GetStatisc() *MjUserStatisc {
	if m != nil {
		return m.Statisc
	}
	return nil
}

// 用户的session
type MjSession struct {
	UserId           *uint32 `protobuf:"varint,1,opt,name=UserId" json:"UserId,omitempty"`
	RoomId           *int32  `protobuf:"varint,2,opt,name=RoomId" json:"RoomId,omitempty"`
	DeskId           *int32  `protobuf:"varint,3,opt,name=DeskId" json:"DeskId,omitempty"`
	GameStatus       *int32  `protobuf:"varint,4,opt,name=GameStatus" json:"GameStatus,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MjSession) Reset()                    { *m = MjSession{} }
func (m *MjSession) String() string            { return proto.CompactTextString(m) }
func (*MjSession) ProtoMessage()               {}
func (*MjSession) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *MjSession) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *MjSession) GetRoomId() int32 {
	if m != nil && m.RoomId != nil {
		return *m.RoomId
	}
	return 0
}

func (m *MjSession) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

func (m *MjSession) GetGameStatus() int32 {
	if m != nil && m.GameStatus != nil {
		return *m.GameStatus
	}
	return 0
}

// 一个麻将room
type MjRoom struct {
	RoomId           *int32    `protobuf:"varint,1,opt,name=RoomId" json:"RoomId,omitempty"`
	RoomType         *int32    `protobuf:"varint,2,opt,name=RoomType" json:"RoomType,omitempty"`
	ReadyTime        *int64    `protobuf:"varint,3,opt,name=ReadyTime" json:"ReadyTime,omitempty"`
	BeginTime        *int64    `protobuf:"varint,4,opt,name=BeginTime" json:"BeginTime,omitempty"`
	EndTIme          *int64    `protobuf:"varint,5,opt,name=EndTIme" json:"EndTIme,omitempty"`
	Desks            []*MjDesk `protobuf:"bytes,6,rep,name=Desks" json:"Desks,omitempty"`
	PaiIndexNow      *int32    `protobuf:"varint,7,opt,name=PaiIndexNow" json:"PaiIndexNow,omitempty"`
	MjPais           []*MJPai  `protobuf:"bytes,8,rep,name=MjPais" json:"MjPais,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *MjRoom) Reset()                    { *m = MjRoom{} }
func (m *MjRoom) String() string            { return proto.CompactTextString(m) }
func (*MjRoom) ProtoMessage()               {}
func (*MjRoom) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *MjRoom) GetRoomId() int32 {
	if m != nil && m.RoomId != nil {
		return *m.RoomId
	}
	return 0
}

func (m *MjRoom) GetRoomType() int32 {
	if m != nil && m.RoomType != nil {
		return *m.RoomType
	}
	return 0
}

func (m *MjRoom) GetReadyTime() int64 {
	if m != nil && m.ReadyTime != nil {
		return *m.ReadyTime
	}
	return 0
}

func (m *MjRoom) GetBeginTime() int64 {
	if m != nil && m.BeginTime != nil {
		return *m.BeginTime
	}
	return 0
}

func (m *MjRoom) GetEndTIme() int64 {
	if m != nil && m.EndTIme != nil {
		return *m.EndTIme
	}
	return 0
}

func (m *MjRoom) GetDesks() []*MjDesk {
	if m != nil {
		return m.Desks
	}
	return nil
}

func (m *MjRoom) GetPaiIndexNow() int32 {
	if m != nil && m.PaiIndexNow != nil {
		return *m.PaiIndexNow
	}
	return 0
}

func (m *MjRoom) GetMjPais() []*MJPai {
	if m != nil {
		return m.MjPais
	}
	return nil
}

//
type RunningDeskKeys struct {
	Keys             []int32 `protobuf:"varint,1,rep,name=keys" json:"keys,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RunningDeskKeys) Reset()                    { *m = RunningDeskKeys{} }
func (m *RunningDeskKeys) String() string            { return proto.CompactTextString(m) }
func (*RunningDeskKeys) ProtoMessage()               {}
func (*RunningDeskKeys) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *RunningDeskKeys) GetKeys() []int32 {
	if m != nil {
		return m.Keys
	}
	return nil
}

func init() {
	proto.RegisterType((*MJPai)(nil), "majiang.MJPai")
	proto.RegisterType((*MJHandPai)(nil), "majiang.MJHandPai")
	proto.RegisterType((*HuPaiInfo)(nil), "majiang.HuPaiInfo")
	proto.RegisterType((*GangPaiInfo)(nil), "majiang.GangPaiInfo")
	proto.RegisterType((*PlayerGameData)(nil), "majiang.PlayerGameData")
	proto.RegisterType((*CheckBean)(nil), "majiang.CheckBean")
	proto.RegisterType((*CheckCase)(nil), "majiang.CheckCase")
	proto.RegisterType((*MjDesk)(nil), "majiang.MjDesk")
	proto.RegisterType((*BillBean)(nil), "majiang.BillBean")
	proto.RegisterType((*Bill)(nil), "majiang.Bill")
	proto.RegisterType((*StatiscRound)(nil), "majiang.StatiscRound")
	proto.RegisterType((*MjUserStatisc)(nil), "majiang.MjUserStatisc")
	proto.RegisterType((*MjUser)(nil), "majiang.MjUser")
	proto.RegisterType((*MjSession)(nil), "majiang.MjSession")
	proto.RegisterType((*MjRoom)(nil), "majiang.MjRoom")
	proto.RegisterType((*RunningDeskKeys)(nil), "majiang.RunningDeskKeys")
}

var fileDescriptor0 = []byte{
	// 1256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x56, 0x4b, 0x73, 0xe3, 0x44,
	0x10, 0x2e, 0x3f, 0x64, 0xcb, 0x63, 0x3b, 0x0f, 0x6d, 0x92, 0x15, 0xaf, 0x6c, 0x10, 0x05, 0x84,
	0x47, 0xe5, 0x90, 0x2b, 0xa7, 0xd8, 0xd9, 0xcd, 0x03, 0xbc, 0x09, 0x4e, 0x60, 0xab, 0xb8, 0x50,
	0x13, 0x6b, 0x62, 0xcb, 0xb1, 0x67, 0x82, 0x46, 0xca, 0x83, 0x33, 0x7f, 0x81, 0x1f, 0x41, 0x15,
	0x67, 0x2e, 0xfc, 0x14, 0x8a, 0x2a, 0x7e, 0x0a, 0xdd, 0x3d, 0x23, 0x5b, 0x71, 0x94, 0xbd, 0xec,
	0xc9, 0x56, 0x4f, 0x4f, 0x4f, 0xf7, 0xd7, 0x5f, 0x7f, 0x33, 0x6c, 0x6d, 0xca, 0x47, 0x63, 0x25,
	0x87, 0x3f, 0x6b, 0x11, 0xdf, 0x88, 0x78, 0xe7, 0x3a, 0x56, 0x89, 0xf2, 0xea, 0x53, 0x3e, 0x8e,
	0xb8, 0x1c, 0x06, 0x1d, 0xe6, 0xf4, 0x8e, 0x4f, 0x79, 0xe4, 0xb5, 0x99, 0x13, 0xc9, 0x50, 0xdc,
	0xf9, 0xa5, 0xad, 0xd2, 0xb6, 0xe3, 0x2d, 0xb1, 0xda, 0xe5, 0x44, 0xdd, 0x8a, 0xd8, 0x2f, 0xd3,
	0x37, 0x2c, 0xdf, 0xf0, 0x49, 0x2a, 0xfc, 0x0a, 0x7d, 0x36, 0x59, 0x25, 0x14, 0xda, 0xaf, 0xc2,
	0x47, 0x23, 0xf8, 0xaf, 0xc4, 0x1a, 0xbd, 0xe3, 0x43, 0x2e, 0x43, 0x0c, 0xf4, 0x21, 0xab, 0xc2,
	0x8f, 0x86, 0x38, 0x95, 0xed, 0xe6, 0xee, 0xd2, 0x8e, 0x3d, 0x69, 0xc7, 0x1c, 0xb3, 0xc5, 0xdc,
	0x53, 0x21, 0x87, 0xe4, 0x51, 0x7e, 0xca, 0xe3, 0x80, 0x5b, 0x8f, 0x4a, 0xa1, 0xc7, 0x26, 0xab,
	0x1d, 0xa6, 0xb4, 0x5e, 0x2d, 0x5c, 0xff, 0x08, 0x4b, 0x81, 0x3f, 0xbe, 0x03, 0xe9, 0x3d, 0x5e,
	0x5e, 0x65, 0x8d, 0x5f, 0x52, 0xf1, 0xca, 0x54, 0x57, 0xa3, 0x72, 0x5e, 0xb0, 0xfa, 0x49, 0x9a,
	0x50, 0xc8, 0x7a, 0x51, 0xc8, 0xe0, 0x4f, 0x28, 0x91, 0xce, 0x3c, 0x92, 0x97, 0xca, 0xf3, 0x18,
	0xd3, 0x42, 0x86, 0x3f, 0x00, 0xa4, 0x47, 0x21, 0x01, 0xd6, 0xc6, 0xa8, 0x43, 0x91, 0x58, 0x53,
	0x99, 0x4c, 0x80, 0xd9, 0xc5, 0xfd, 0x9b, 0x91, 0xb2, 0x98, 0x01, 0xa4, 0xa3, 0xf4, 0xfc, 0xfe,
	0x5a, 0x10, 0x6c, 0x8e, 0xb7, 0xc2, 0xdc, 0x01, 0x8f, 0x43, 0xb2, 0x38, 0x73, 0x8f, 0x7d, 0xa1,
	0x07, 0x94, 0x56, 0xc3, 0xfb, 0x80, 0x55, 0xae, 0xa1, 0x8c, 0x7a, 0x61, 0x19, 0xd0, 0x82, 0x4b,
	0x2e, 0x7d, 0x37, 0x6b, 0x8f, 0x1e, 0xa8, 0x58, 0xf8, 0x0d, 0xf8, 0xac, 0x04, 0xbf, 0x95, 0x58,
	0xd3, 0x82, 0xf8, 0x0e, 0x09, 0x43, 0x82, 0x43, 0x08, 0x92, 0x4b, 0xd9, 0x26, 0x54, 0x8c, 0x2b,
	0xb8, 0x27, 0x31, 0x97, 0xfa, 0x32, 0x83, 0x35, 0xf8, 0xbd, 0xcc, 0x96, 0x4e, 0x27, 0xfc, 0x5e,
	0xc4, 0x07, 0x7c, 0x2a, 0xf6, 0x79, 0xc2, 0xbd, 0x4f, 0x58, 0x7d, 0x64, 0x88, 0x42, 0x69, 0x34,
	0x77, 0xbd, 0x5c, 0x94, 0x8c, 0x42, 0x01, 0xe2, 0x80, 0x89, 0x5b, 0x8a, 0xcc, 0x7d, 0xe6, 0x3d,
	0xf8, 0x8c, 0xb1, 0x10, 0x4c, 0x87, 0xc6, 0xaf, 0xf2, 0x16, 0x3f, 0x2a, 0x82, 0xbc, 0x0c, 0x5d,
	0xd6, 0x66, 0x5e, 0x79, 0x88, 0xbe, 0x64, 0x2d, 0x8c, 0x77, 0x90, 0xf9, 0x3a, 0x6f, 0xf7, 0x25,
	0x60, 0xe6, 0xd5, 0x3e, 0xed, 0x0b, 0xd0, 0x27, 0x2a, 0xe1, 0x93, 0x33, 0x6a, 0x4f, 0x9d, 0x70,
	0xe1, 0xac, 0xd1, 0x1d, 0x89, 0xc1, 0x55, 0x47, 0x70, 0x89, 0x4d, 0x4f, 0xf3, 0x7d, 0x59, 0x66,
	0xf5, 0x2e, 0x97, 0x38, 0x24, 0xd4, 0x15, 0xd7, 0x1a, 0x30, 0x26, 0xf5, 0xc5, 0xc5, 0x36, 0x75,
	0xb1, 0x72, 0x6a, 0x8a, 0xeb, 0x3d, 0x63, 0x4d, 0x8a, 0x76, 0x96, 0xf0, 0x24, 0xd5, 0x86, 0x4a,
	0xc1, 0x1f, 0x25, 0x7b, 0x46, 0x97, 0x6b, 0x81, 0x80, 0x9a, 0x03, 0xed, 0x54, 0xce, 0x81, 0x9a,
	0xe7, 0xb1, 0x10, 0xc6, 0x8c, 0x7d, 0xc0, 0x18, 0x19, 0xa9, 0xc3, 0x74, 0x7c, 0xe1, 0x3c, 0x19,
	0x16, 0xc1, 0x08, 0x51, 0x4a, 0x6d, 0xef, 0x6b, 0xe8, 0x7b, 0x2c, 0xe0, 0x3b, 0x07, 0x67, 0xe9,
	0x29, 0x88, 0x82, 0xbf, 0xab, 0xac, 0xd6, 0x1b, 0x03, 0xef, 0xaf, 0x10, 0x0c, 0xfc, 0xb5, 0x60,
	0xd0, 0x44, 0xf4, 0x95, 0x9a, 0x5a, 0x86, 0xd2, 0xb7, 0xcd, 0xcf, 0x50, 0x74, 0x93, 0x39, 0x08,
	0x5e, 0xa6, 0x04, 0xcb, 0xf3, 0xd4, 0xc6, 0x98, 0x13, 0x72, 0xf2, 0x94, 0x6b, 0x7d, 0xab, 0xe2,
	0x90, 0x52, 0x68, 0x20, 0x78, 0x27, 0xb7, 0xd2, 0x52, 0x94, 0xa6, 0xa0, 0x1b, 0x0b, 0x9e, 0x88,
	0x57, 0xc2, 0x74, 0xa7, 0x82, 0x1d, 0x9b, 0x8e, 0xf1, 0x54, 0x22, 0xbe, 0x99, 0x2f, 0xb0, 0x5d,
	0x28, 0x18, 0x56, 0xdd, 0x55, 0x50, 0x64, 0x23, 0xcb, 0x65, 0xc0, 0xaf, 0x7b, 0xfc, 0xce, 0x67,
	0xb4, 0xcf, 0xce, 0xb3, 0x7e, 0x9d, 0x4e, 0xfd, 0x66, 0xb6, 0x4b, 0x8b, 0x24, 0x99, 0x88, 0xa9,
	0x90, 0x89, 0xdf, 0x22, 0x1b, 0x1c, 0x78, 0x01, 0x2d, 0xf9, 0x91, 0xc4, 0xb4, 0x4d, 0x1b, 0xc1,
	0xf4, 0x6b, 0xd4, 0x53, 0x7d, 0x1e, 0x46, 0xca, 0x5f, 0x22, 0x2f, 0x9f, 0xad, 0x64, 0x6c, 0x3c,
	0x4c, 0xb9, 0x59, 0x59, 0xa6, 0x95, 0x0d, 0xb6, 0xa4, 0x92, 0x11, 0x94, 0x6c, 0x3a, 0xa7, 0xee,
	0xfc, 0x15, 0x28, 0xdd, 0x41, 0x96, 0x8c, 0x52, 0xe3, 0xb8, 0x4a, 0x8e, 0xa0, 0xa3, 0x7b, 0x93,
	0x89, 0x69, 0x9c, 0x57, 0xa8, 0x93, 0x1f, 0xb3, 0x06, 0x82, 0x6d, 0x5c, 0x9e, 0x15, 0xba, 0x00,
	0x29, 0xe8, 0x4f, 0x37, 0x8d, 0xb5, 0x8a, 0xfd, 0xb5, 0x2c, 0x85, 0x73, 0xa4, 0x34, 0x8e, 0x36,
	0xe0, 0x01, 0xa5, 0xad, 0x93, 0x7d, 0x9d, 0xb5, 0xc1, 0x2f, 0x9e, 0x9b, 0x37, 0x32, 0x9c, 0x3a,
	0x5c, 0x5e, 0x01, 0xe4, 0xcf, 0x09, 0xf2, 0x4f, 0x73, 0xcc, 0xf4, 0xfd, 0x05, 0x11, 0x98, 0x73,
	0x16, 0xc0, 0xdb, 0x1b, 0x24, 0xd1, 0x8d, 0xc0, 0x46, 0xfa, 0xef, 0xd1, 0x56, 0xb0, 0xa1, 0x92,
	0x00, 0xc2, 0x17, 0x60, 0x7b, 0x9f, 0x98, 0x9e, 0x32, 0xb7, 0x13, 0x4d, 0x26, 0xd9, 0x2c, 0x2d,
	0x6a, 0x1c, 0x90, 0xf0, 0x81, 0xc6, 0xb5, 0x58, 0x35, 0xc1, 0xbe, 0x3e, 0xbe, 0xc7, 0x70, 0x37,
	0x9f, 0x52, 0xe2, 0x0e, 0xf5, 0xc5, 0xaa, 0x5d, 0xad, 0x88, 0xf5, 0xc1, 0x37, 0xac, 0x8a, 0xc7,
	0xe2, 0x11, 0x6f, 0x22, 0xb9, 0x67, 0xf6, 0x95, 0x68, 0xdf, 0x16, 0xc8, 0x28, 0x2c, 0x65, 0x17,
	0xdc, 0xea, 0x6c, 0x67, 0x96, 0x67, 0x70, 0xce, 0x5a, 0x48, 0xe3, 0x48, 0x0f, 0xfa, 0xb0, 0x2f,
	0x44, 0x52, 0xd2, 0x1f, 0xcb, 0xfa, 0x87, 0x65, 0xce, 0x98, 0xdf, 0x17, 0x3a, 0x9d, 0x24, 0x94,
	0x79, 0xe3, 0xe1, 0xb9, 0x55, 0x52, 0xfd, 0x7f, 0x4b, 0xac, 0x6d, 0x78, 0x6f, 0x83, 0x23, 0x29,
	0xc0, 0xa9, 0xab, 0x22, 0x69, 0x53, 0x43, 0x2d, 0xc1, 0x1d, 0x20, 0x1e, 0xe5, 0x8c, 0x8e, 0x64,
	0xf8, 0x09, 0x08, 0x68, 0x31, 0x59, 0x63, 0x2d, 0x32, 0xed, 0x43, 0xce, 0xa7, 0x5c, 0x59, 0xe9,
	0x47, 0x79, 0x40, 0xeb, 0x9e, 0x51, 0x22, 0x67, 0xd6, 0x71, 0x34, 0xf6, 0x22, 0x39, 0x24, 0x73,
	0xed, 0x81, 0x79, 0xdf, 0x52, 0xd8, 0xc8, 0xde, 0x2c, 0x70, 0x77, 0xc4, 0x8f, 0x23, 0x08, 0x6c,
	0x46, 0x6b, 0x9b, 0x35, 0x62, 0xac, 0x1d, 0x81, 0x81, 0xc9, 0x42, 0xc4, 0xd6, 0x67, 0x88, 0xe5,
	0x51, 0x0a, 0xfe, 0x29, 0xa3, 0x4e, 0xd0, 0x5c, 0x2f, 0x36, 0x1a, 0xba, 0x4a, 0x55, 0x96, 0xa9,
	0xca, 0x2f, 0xf0, 0x09, 0x61, 0x2e, 0x1c, 0xab, 0x59, 0xcf, 0x67, 0x11, 0x17, 0xee, 0xa3, 0xb9,
	0xa0, 0x98, 0x32, 0x01, 0xa0, 0x23, 0xdd, 0x01, 0x45, 0xb8, 0xa2, 0x12, 0x5d, 0x63, 0xf8, 0x4e,
	0xf0, 0x1b, 0x41, 0xc5, 0xb9, 0x39, 0x89, 0xaa, 0x2f, 0x48, 0x94, 0x9b, 0x45, 0xd8, 0x07, 0x38,
	0xbe, 0x4f, 0xcd, 0x65, 0xec, 0x22, 0xc4, 0x2f, 0xef, 0x06, 0x70, 0xeb, 0x0d, 0x45, 0x48, 0x52,
	0x41, 0x0a, 0xde, 0x17, 0x3c, 0xbc, 0x27, 0x9d, 0x70, 0x51, 0x39, 0xe0, 0x50, 0x33, 0x23, 0x2d,
	0xb2, 0x7c, 0xc5, 0xda, 0x20, 0xa0, 0x3d, 0x35, 0xd3, 0xcf, 0xf6, 0xd3, 0xfa, 0x09, 0x3c, 0xad,
	0x22, 0xdf, 0x48, 0x3a, 0x9a, 0xbb, 0xed, 0x07, 0x74, 0xf3, 0x3e, 0x67, 0x75, 0x6d, 0x40, 0x24,
	0x01, 0x69, 0xee, 0x6e, 0x2c, 0x68, 0xa4, 0x85, 0x38, 0x38, 0x81, 0x47, 0xdc, 0xf8, 0x4c, 0x68,
	0x1d, 0xa9, 0xc7, 0x83, 0x54, 0xa0, 0xc3, 0x16, 0x84, 0x4a, 0x9e, 0xb1, 0x79, 0x28, 0x83, 0xbf,
	0x4a, 0xd8, 0x2e, 0xdc, 0x96, 0xdb, 0x5e, 0xca, 0x5e, 0x16, 0x33, 0x81, 0x9d, 0xf1, 0x90, 0x10,
	0x39, 0x8f, 0xa6, 0x66, 0x36, 0x49, 0x16, 0x3b, 0x62, 0x18, 0x49, 0x32, 0x55, 0x33, 0xfa, 0xbe,
	0x94, 0xe1, 0xf9, 0xd1, 0x54, 0xd8, 0x11, 0x05, 0xfd, 0xc7, 0x3c, 0xb4, 0xbd, 0x82, 0xf3, 0xfa,
	0x4f, 0xf7, 0x09, 0xb0, 0x96, 0x50, 0x82, 0x87, 0xed, 0x6b, 0x75, 0x6b, 0x3b, 0xb6, 0x89, 0x79,
	0xd1, 0x63, 0xcf, 0x2d, 0x7c, 0xec, 0xbd, 0x60, 0xcb, 0xfd, 0x54, 0x4a, 0x68, 0x22, 0xc6, 0xf8,
	0x56, 0xdc, 0x6b, 0xe4, 0xd7, 0x15, 0xfc, 0xd2, 0xf5, 0xe9, 0xfc, 0x1f, 0x00, 0x00, 0xff, 0xff,
	0x7f, 0x11, 0x39, 0xbf, 0x54, 0x0b, 0x00, 0x00,
}
