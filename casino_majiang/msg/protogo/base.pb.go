// Code generated by protoc-gen-go.
// source: base.proto
// DO NOT EDIT!

/*
Package mjproto is a generated protocol buffer package.

It is generated from these files:
	base.proto
	mahjong_desk.proto
	mahjong_hall.proto
	mahjong_play.proto

It has these top-level messages:
	ProtoHeader
	WeixinInfo
	CardInfo
	PlayOptions
	RoomTypeInfo
	ComposePoker
	PlayerPoker
	PlayerInfo
	DeskGameInfo
	Game_DissolveDesk
	Game_AckDissolveDesk
	Game_LeaveDesk
	Game_AckLeaveDesk
	Game_Ready
	Game_AckReady
	Game_Message
	Game_SendMessage
	CardType
	WinCoinInfo
	Game_CurrentResult
	EndLotteryInfo
	Game_SendEndLottery
	Game_QuickConn
	Game_AckQuickConn
	Game_Login
	Game_AckLogin
	Game_Notice
	Game_AckNotice
	Game_GameRecord
	Game_BeanUserRecord
	Game_BeanGameRecord
	Game_AckGameRecord
	Game_Feedback
	Game_CreateRoom
	Game_AckCreateRoom
	Game_EnterRoom
	Game_AckEnterRoom
	Game_Opening
	Game_DealCards
	Game_ExchangeCards
	Game_AckExchangeCards
	Game_DingQue
	Game_BroadcastBeginDingQue
	Game_BroadcastBeginExchange
	Game_GetInCard
	Game_SendOutCard
	Game_AckSendOutCard
	Game_ActPeng
	Game_AckActPeng
	Game_ActGang
	Game_AckActGang
	Game_ActHu
	Game_AckActHu
	Game_ActGuo
	Game_AckActGuo
	Game_OverTurn
	Game_SendGameInfo
*/
package mjproto

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

type EProtoId int32

const (
	EProtoId_PID_QUICK_CONN               EProtoId = 1
	EProtoId_PID_QUICK_CONN_ACK           EProtoId = 2
	EProtoId_PID_GAME_LOGIN               EProtoId = 3
	EProtoId_PID_GAME_LOGIN_ACK           EProtoId = 4
	EProtoId_PID_CREATEROOM               EProtoId = 5
	EProtoId_PID_CREATEROOM_ACK           EProtoId = 6
	EProtoId_PID_ENTER_ROOM               EProtoId = 7
	EProtoId_PID_ENTER_ROOM_ACK           EProtoId = 8
	EProtoId_PID_SEND_GAMEINFO            EProtoId = 9
	EProtoId_PID_READY                    EProtoId = 10
	EProtoId_PID_READY_ACK                EProtoId = 11
	EProtoId_PID_EXCHANGECARDS            EProtoId = 12
	EProtoId_PID_EXCHANGECARDS_ACK        EProtoId = 13
	EProtoId_PID_DINGQUE                  EProtoId = 14
	EProtoId_PID_OPENING                  EProtoId = 15
	EProtoId_PID_DEAL_CARDS               EProtoId = 16
	EProtoId_PID_GET_IN_CARD              EProtoId = 17
	EProtoId_PID_SEND_OUT_CARD            EProtoId = 18
	EProtoId_PID_PENG_CARD                EProtoId = 19
	EProtoId_PID_PENG_CARD_ACK            EProtoId = 20
	EProtoId_PID_GANG_CARD                EProtoId = 21
	EProtoId_PID_GANG_CARD_ACK            EProtoId = 22
	EProtoId_PID_GUO_CARD                 EProtoId = 23
	EProtoId_PID_GUO_CARD_ACK             EProtoId = 24
	EProtoId_PID_HU_CARD                  EProtoId = 25
	EProtoId_PID_HU_CARD_ACK              EProtoId = 26
	EProtoId_PID_BROADCAST_BEGIN_DINGQUE  EProtoId = 27
	EProtoId_PID_BROADCAST_BEGIN_EXCHANGE EProtoId = 28
	EProtoId_PID_OVERTURN                 EProtoId = 29
	EProtoId_PID_CURRENTRESULT            EProtoId = 30
	EProtoId_PID_SENDENDLOTTERY           EProtoId = 31
	EProtoId_PID_DISSOLVE_DESK            EProtoId = 32
	EProtoId_PID_DISSOLVE_DESK_ACK        EProtoId = 33
	EProtoId_PID_LEAVE_DESK               EProtoId = 34
	EProtoId_PID_LEAVE_DESK_ACK           EProtoId = 35
	EProtoId_PID_MESSAGE                  EProtoId = 36
	EProtoId_PID_SEND_MESSAGE             EProtoId = 37
)

var EProtoId_name = map[int32]string{
	1:  "PID_QUICK_CONN",
	2:  "PID_QUICK_CONN_ACK",
	3:  "PID_GAME_LOGIN",
	4:  "PID_GAME_LOGIN_ACK",
	5:  "PID_CREATEROOM",
	6:  "PID_CREATEROOM_ACK",
	7:  "PID_ENTER_ROOM",
	8:  "PID_ENTER_ROOM_ACK",
	9:  "PID_SEND_GAMEINFO",
	10: "PID_READY",
	11: "PID_READY_ACK",
	12: "PID_EXCHANGECARDS",
	13: "PID_EXCHANGECARDS_ACK",
	14: "PID_DINGQUE",
	15: "PID_OPENING",
	16: "PID_DEAL_CARDS",
	17: "PID_GET_IN_CARD",
	18: "PID_SEND_OUT_CARD",
	19: "PID_PENG_CARD",
	20: "PID_PENG_CARD_ACK",
	21: "PID_GANG_CARD",
	22: "PID_GANG_CARD_ACK",
	23: "PID_GUO_CARD",
	24: "PID_GUO_CARD_ACK",
	25: "PID_HU_CARD",
	26: "PID_HU_CARD_ACK",
	27: "PID_BROADCAST_BEGIN_DINGQUE",
	28: "PID_BROADCAST_BEGIN_EXCHANGE",
	29: "PID_OVERTURN",
	30: "PID_CURRENTRESULT",
	31: "PID_SENDENDLOTTERY",
	32: "PID_DISSOLVE_DESK",
	33: "PID_DISSOLVE_DESK_ACK",
	34: "PID_LEAVE_DESK",
	35: "PID_LEAVE_DESK_ACK",
	36: "PID_MESSAGE",
	37: "PID_SEND_MESSAGE",
}
var EProtoId_value = map[string]int32{
	"PID_QUICK_CONN":               1,
	"PID_QUICK_CONN_ACK":           2,
	"PID_GAME_LOGIN":               3,
	"PID_GAME_LOGIN_ACK":           4,
	"PID_CREATEROOM":               5,
	"PID_CREATEROOM_ACK":           6,
	"PID_ENTER_ROOM":               7,
	"PID_ENTER_ROOM_ACK":           8,
	"PID_SEND_GAMEINFO":            9,
	"PID_READY":                    10,
	"PID_READY_ACK":                11,
	"PID_EXCHANGECARDS":            12,
	"PID_EXCHANGECARDS_ACK":        13,
	"PID_DINGQUE":                  14,
	"PID_OPENING":                  15,
	"PID_DEAL_CARDS":               16,
	"PID_GET_IN_CARD":              17,
	"PID_SEND_OUT_CARD":            18,
	"PID_PENG_CARD":                19,
	"PID_PENG_CARD_ACK":            20,
	"PID_GANG_CARD":                21,
	"PID_GANG_CARD_ACK":            22,
	"PID_GUO_CARD":                 23,
	"PID_GUO_CARD_ACK":             24,
	"PID_HU_CARD":                  25,
	"PID_HU_CARD_ACK":              26,
	"PID_BROADCAST_BEGIN_DINGQUE":  27,
	"PID_BROADCAST_BEGIN_EXCHANGE": 28,
	"PID_OVERTURN":                 29,
	"PID_CURRENTRESULT":            30,
	"PID_SENDENDLOTTERY":           31,
	"PID_DISSOLVE_DESK":            32,
	"PID_DISSOLVE_DESK_ACK":        33,
	"PID_LEAVE_DESK":               34,
	"PID_LEAVE_DESK_ACK":           35,
	"PID_MESSAGE":                  36,
	"PID_SEND_MESSAGE":             37,
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
func (EProtoId) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

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
func (DDErrorCode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// 房间类型信息：包含房间类型和对应的局数、封顶、玩法等信息
// 房间类型枚举
type MJRoomType int32

const (
	MJRoomType_roomType_xueZhanDaoDi    MJRoomType = 0
	MJRoomType_roomType_sanRenLiangFang MJRoomType = 1
	MJRoomType_roomType_siRenLiangFang  MJRoomType = 2
	MJRoomType_roomType_deYangMaJiang   MJRoomType = 3
	MJRoomType_roomType_daoDaoHu        MJRoomType = 4
	MJRoomType_roomType_xueLiuChengHe   MJRoomType = 5
)

var MJRoomType_name = map[int32]string{
	0: "roomType_xueZhanDaoDi",
	1: "roomType_sanRenLiangFang",
	2: "roomType_siRenLiangFang",
	3: "roomType_deYangMaJiang",
	4: "roomType_daoDaoHu",
	5: "roomType_xueLiuChengHe",
}
var MJRoomType_value = map[string]int32{
	"roomType_xueZhanDaoDi":    0,
	"roomType_sanRenLiangFang": 1,
	"roomType_siRenLiangFang":  2,
	"roomType_deYangMaJiang":   3,
	"roomType_daoDaoHu":        4,
	"roomType_xueLiuChengHe":   5,
}

func (x MJRoomType) Enum() *MJRoomType {
	p := new(MJRoomType)
	*p = x
	return p
}
func (x MJRoomType) String() string {
	return proto.EnumName(MJRoomType_name, int32(x))
}
func (x *MJRoomType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MJRoomType_value, data, "MJRoomType")
	if err != nil {
		return err
	}
	*x = MJRoomType(value)
	return nil
}
func (MJRoomType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

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
func (*ProtoHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

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

// 微信信息
type WeixinInfo struct {
	OpenId           *string `protobuf:"bytes,1,opt,name=openId" json:"openId,omitempty"`
	NickName         *string `protobuf:"bytes,2,opt,name=nickName" json:"nickName,omitempty"`
	HeadUrl          *string `protobuf:"bytes,3,opt,name=headUrl" json:"headUrl,omitempty"`
	Sex              *int32  `protobuf:"varint,4,opt,name=sex" json:"sex,omitempty"`
	City             *string `protobuf:"bytes,5,opt,name=city" json:"city,omitempty"`
	UnionId          *string `protobuf:"bytes,6,opt,name=unionId" json:"unionId,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *WeixinInfo) Reset()                    { *m = WeixinInfo{} }
func (m *WeixinInfo) String() string            { return proto.CompactTextString(m) }
func (*WeixinInfo) ProtoMessage()               {}
func (*WeixinInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *WeixinInfo) GetOpenId() string {
	if m != nil && m.OpenId != nil {
		return *m.OpenId
	}
	return ""
}

func (m *WeixinInfo) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *WeixinInfo) GetHeadUrl() string {
	if m != nil && m.HeadUrl != nil {
		return *m.HeadUrl
	}
	return ""
}

func (m *WeixinInfo) GetSex() int32 {
	if m != nil && m.Sex != nil {
		return *m.Sex
	}
	return 0
}

func (m *WeixinInfo) GetCity() string {
	if m != nil && m.City != nil {
		return *m.City
	}
	return ""
}

func (m *WeixinInfo) GetUnionId() string {
	if m != nil && m.UnionId != nil {
		return *m.UnionId
	}
	return ""
}

// 麻将牌
type CardInfo struct {
	Type             *int32 `protobuf:"varint,1,opt,name=type" json:"type,omitempty"`
	Value            *int32 `protobuf:"varint,2,opt,name=value" json:"value,omitempty"`
	Id               *int32 `protobuf:"varint,3,opt,name=id" json:"id,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CardInfo) Reset()                    { *m = CardInfo{} }
func (m *CardInfo) String() string            { return proto.CompactTextString(m) }
func (*CardInfo) ProtoMessage()               {}
func (*CardInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CardInfo) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *CardInfo) GetValue() int32 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

func (m *CardInfo) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

// 玩法：包括自摸、点炮、以及可多选的玩法
type PlayOptions struct {
	ZiMoRadio        *int32  `protobuf:"varint,1,opt,name=ziMoRadio" json:"ziMoRadio,omitempty"`
	DianGangHuaRadio *int32  `protobuf:"varint,2,opt,name=dianGangHuaRadio" json:"dianGangHuaRadio,omitempty"`
	OthersCheckBox   []int32 `protobuf:"varint,3,rep,name=othersCheckBox" json:"othersCheckBox,omitempty"`
	HuRadio          *int32  `protobuf:"varint,4,opt,name=huRadio" json:"huRadio,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PlayOptions) Reset()                    { *m = PlayOptions{} }
func (m *PlayOptions) String() string            { return proto.CompactTextString(m) }
func (*PlayOptions) ProtoMessage()               {}
func (*PlayOptions) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PlayOptions) GetZiMoRadio() int32 {
	if m != nil && m.ZiMoRadio != nil {
		return *m.ZiMoRadio
	}
	return 0
}

func (m *PlayOptions) GetDianGangHuaRadio() int32 {
	if m != nil && m.DianGangHuaRadio != nil {
		return *m.DianGangHuaRadio
	}
	return 0
}

func (m *PlayOptions) GetOthersCheckBox() []int32 {
	if m != nil {
		return m.OthersCheckBox
	}
	return nil
}

func (m *PlayOptions) GetHuRadio() int32 {
	if m != nil && m.HuRadio != nil {
		return *m.HuRadio
	}
	return 0
}

type RoomTypeInfo struct {
	MjRoomType       *MJRoomType  `protobuf:"varint,1,opt,name=mjRoomType,enum=mjproto.MJRoomType" json:"mjRoomType,omitempty"`
	BoardsCout       *int32       `protobuf:"varint,2,opt,name=boardsCout" json:"boardsCout,omitempty"`
	CapMax           *int64       `protobuf:"varint,3,opt,name=capMax" json:"capMax,omitempty"`
	PlayOptions      *PlayOptions `protobuf:"bytes,4,opt,name=playOptions" json:"playOptions,omitempty"`
	CardsNum         *int32       `protobuf:"varint,5,opt,name=cardsNum" json:"cardsNum,omitempty"`
	Settlement       *int32       `protobuf:"varint,6,opt,name=settlement" json:"settlement,omitempty"`
	BaseValue        *int64       `protobuf:"varint,7,opt,name=baseValue" json:"baseValue,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *RoomTypeInfo) Reset()                    { *m = RoomTypeInfo{} }
func (m *RoomTypeInfo) String() string            { return proto.CompactTextString(m) }
func (*RoomTypeInfo) ProtoMessage()               {}
func (*RoomTypeInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RoomTypeInfo) GetMjRoomType() MJRoomType {
	if m != nil && m.MjRoomType != nil {
		return *m.MjRoomType
	}
	return MJRoomType_roomType_xueZhanDaoDi
}

func (m *RoomTypeInfo) GetBoardsCout() int32 {
	if m != nil && m.BoardsCout != nil {
		return *m.BoardsCout
	}
	return 0
}

func (m *RoomTypeInfo) GetCapMax() int64 {
	if m != nil && m.CapMax != nil {
		return *m.CapMax
	}
	return 0
}

func (m *RoomTypeInfo) GetPlayOptions() *PlayOptions {
	if m != nil {
		return m.PlayOptions
	}
	return nil
}

func (m *RoomTypeInfo) GetCardsNum() int32 {
	if m != nil && m.CardsNum != nil {
		return *m.CardsNum
	}
	return 0
}

func (m *RoomTypeInfo) GetSettlement() int32 {
	if m != nil && m.Settlement != nil {
		return *m.Settlement
	}
	return 0
}

func (m *RoomTypeInfo) GetBaseValue() int64 {
	if m != nil && m.BaseValue != nil {
		return *m.BaseValue
	}
	return 0
}

type ComposePoker struct {
	Num              *int32 `protobuf:"varint,1,opt,name=num" json:"num,omitempty"`
	Type             *int32 `protobuf:"varint,2,opt,name=type" json:"type,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ComposePoker) Reset()                    { *m = ComposePoker{} }
func (m *ComposePoker) String() string            { return proto.CompactTextString(m) }
func (*ComposePoker) ProtoMessage()               {}
func (*ComposePoker) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ComposePoker) GetNum() int32 {
	if m != nil && m.Num != nil {
		return *m.Num
	}
	return 0
}

func (m *ComposePoker) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

type PlayerPoker struct {
	HandPoker        []int32         `protobuf:"varint,1,rep,name=handPoker" json:"handPoker,omitempty"`
	ComposePoker     []*ComposePoker `protobuf:"bytes,2,rep,name=composePoker" json:"composePoker,omitempty"`
	OutPoker         []int32         `protobuf:"varint,3,rep,name=outPoker" json:"outPoker,omitempty"`
	HuPoker          *int32          `protobuf:"varint,4,opt,name=huPoker" json:"huPoker,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *PlayerPoker) Reset()                    { *m = PlayerPoker{} }
func (m *PlayerPoker) String() string            { return proto.CompactTextString(m) }
func (*PlayerPoker) ProtoMessage()               {}
func (*PlayerPoker) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *PlayerPoker) GetHandPoker() []int32 {
	if m != nil {
		return m.HandPoker
	}
	return nil
}

func (m *PlayerPoker) GetComposePoker() []*ComposePoker {
	if m != nil {
		return m.ComposePoker
	}
	return nil
}

func (m *PlayerPoker) GetOutPoker() []int32 {
	if m != nil {
		return m.OutPoker
	}
	return nil
}

func (m *PlayerPoker) GetHuPoker() int32 {
	if m != nil && m.HuPoker != nil {
		return *m.HuPoker
	}
	return 0
}

type PlayerInfo struct {
	IsBanker    *bool        `protobuf:"varint,1,opt,name=isBanker" json:"isBanker,omitempty"`
	PlayerPoker *PlayerPoker `protobuf:"bytes,2,opt,name=playerPoker" json:"playerPoker,omitempty"`
	Coin        *int64       `protobuf:"varint,3,opt,name=coin" json:"coin,omitempty"`
	NickName    *string      `protobuf:"bytes,4,opt,name=nickName" json:"nickName,omitempty"`
	SeatId      *int32       `protobuf:"varint,5,opt,name=seatId" json:"seatId,omitempty"`
	// ① 新增字段： repeated WeixinInfo  = 6;  //微信用户信息
	// ② EnterMatch时、广播新进用户时，才需返回 WeixinInfo；其它广播不需要。
	WeixinInfos      *WeixinInfo `protobuf:"bytes,6,opt,name=weixinInfos" json:"weixinInfos,omitempty"`
	IsOwner          *bool       `protobuf:"varint,7,opt,name=isOwner" json:"isOwner,omitempty"`
	BReady           *int32      `protobuf:"varint,8,opt,name=bReady" json:"bReady,omitempty"`
	BDingQue         *int32      `protobuf:"varint,9,opt,name=bDingQue" json:"bDingQue,omitempty"`
	BExchanged       *int32      `protobuf:"varint,10,opt,name=bExchanged" json:"bExchanged,omitempty"`
	NHuPai           *int32      `protobuf:"varint,11,opt,name=nHuPai" json:"nHuPai,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *PlayerInfo) Reset()                    { *m = PlayerInfo{} }
func (m *PlayerInfo) String() string            { return proto.CompactTextString(m) }
func (*PlayerInfo) ProtoMessage()               {}
func (*PlayerInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *PlayerInfo) GetIsBanker() bool {
	if m != nil && m.IsBanker != nil {
		return *m.IsBanker
	}
	return false
}

func (m *PlayerInfo) GetPlayerPoker() *PlayerPoker {
	if m != nil {
		return m.PlayerPoker
	}
	return nil
}

func (m *PlayerInfo) GetCoin() int64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *PlayerInfo) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *PlayerInfo) GetSeatId() int32 {
	if m != nil && m.SeatId != nil {
		return *m.SeatId
	}
	return 0
}

func (m *PlayerInfo) GetWeixinInfos() *WeixinInfo {
	if m != nil {
		return m.WeixinInfos
	}
	return nil
}

func (m *PlayerInfo) GetIsOwner() bool {
	if m != nil && m.IsOwner != nil {
		return *m.IsOwner
	}
	return false
}

func (m *PlayerInfo) GetBReady() int32 {
	if m != nil && m.BReady != nil {
		return *m.BReady
	}
	return 0
}

func (m *PlayerInfo) GetBDingQue() int32 {
	if m != nil && m.BDingQue != nil {
		return *m.BDingQue
	}
	return 0
}

func (m *PlayerInfo) GetBExchanged() int32 {
	if m != nil && m.BExchanged != nil {
		return *m.BExchanged
	}
	return 0
}

func (m *PlayerInfo) GetNHuPai() int32 {
	if m != nil && m.NHuPai != nil {
		return *m.NHuPai
	}
	return 0
}

type DeskGameInfo struct {
	GameStatus       *int32        `protobuf:"varint,1,opt,name=GameStatus" json:"GameStatus,omitempty"`
	RoomTypeInfo     *RoomTypeInfo `protobuf:"bytes,2,opt,name=roomTypeInfo" json:"roomTypeInfo,omitempty"`
	PlayerNum        *int32        `protobuf:"varint,3,opt,name=playerNum" json:"playerNum,omitempty"`
	ActiveSeat       *int32        `protobuf:"varint,4,opt,name=activeSeat" json:"activeSeat,omitempty"`
	ActionTime       *int32        `protobuf:"varint,5,opt,name=actionTime" json:"actionTime,omitempty"`
	DelayTime        *int32        `protobuf:"varint,6,opt,name=delayTime" json:"delayTime,omitempty"`
	NRebuyCount      *int32        `protobuf:"varint,7,opt,name=nRebuyCount" json:"nRebuyCount,omitempty"`
	NInitActionTime  *int32        `protobuf:"varint,8,opt,name=nInitActionTime" json:"nInitActionTime,omitempty"`
	NInitDelayTime   *int32        `protobuf:"varint,9,opt,name=nInitDelayTime" json:"nInitDelayTime,omitempty"`
	InitRoomCoin     *int64        `protobuf:"varint,10,opt,name=initRoomCoin" json:"initRoomCoin,omitempty"`
	CurrPlayCount    *int32        `protobuf:"varint,11,opt,name=currPlayCount" json:"currPlayCount,omitempty"`
	TotalPlayCount   *int32        `protobuf:"varint,12,opt,name=totalPlayCount" json:"totalPlayCount,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *DeskGameInfo) Reset()                    { *m = DeskGameInfo{} }
func (m *DeskGameInfo) String() string            { return proto.CompactTextString(m) }
func (*DeskGameInfo) ProtoMessage()               {}
func (*DeskGameInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *DeskGameInfo) GetGameStatus() int32 {
	if m != nil && m.GameStatus != nil {
		return *m.GameStatus
	}
	return 0
}

func (m *DeskGameInfo) GetRoomTypeInfo() *RoomTypeInfo {
	if m != nil {
		return m.RoomTypeInfo
	}
	return nil
}

func (m *DeskGameInfo) GetPlayerNum() int32 {
	if m != nil && m.PlayerNum != nil {
		return *m.PlayerNum
	}
	return 0
}

func (m *DeskGameInfo) GetActiveSeat() int32 {
	if m != nil && m.ActiveSeat != nil {
		return *m.ActiveSeat
	}
	return 0
}

func (m *DeskGameInfo) GetActionTime() int32 {
	if m != nil && m.ActionTime != nil {
		return *m.ActionTime
	}
	return 0
}

func (m *DeskGameInfo) GetDelayTime() int32 {
	if m != nil && m.DelayTime != nil {
		return *m.DelayTime
	}
	return 0
}

func (m *DeskGameInfo) GetNRebuyCount() int32 {
	if m != nil && m.NRebuyCount != nil {
		return *m.NRebuyCount
	}
	return 0
}

func (m *DeskGameInfo) GetNInitActionTime() int32 {
	if m != nil && m.NInitActionTime != nil {
		return *m.NInitActionTime
	}
	return 0
}

func (m *DeskGameInfo) GetNInitDelayTime() int32 {
	if m != nil && m.NInitDelayTime != nil {
		return *m.NInitDelayTime
	}
	return 0
}

func (m *DeskGameInfo) GetInitRoomCoin() int64 {
	if m != nil && m.InitRoomCoin != nil {
		return *m.InitRoomCoin
	}
	return 0
}

func (m *DeskGameInfo) GetCurrPlayCount() int32 {
	if m != nil && m.CurrPlayCount != nil {
		return *m.CurrPlayCount
	}
	return 0
}

func (m *DeskGameInfo) GetTotalPlayCount() int32 {
	if m != nil && m.TotalPlayCount != nil {
		return *m.TotalPlayCount
	}
	return 0
}

func init() {
	proto.RegisterType((*ProtoHeader)(nil), "mjproto.ProtoHeader")
	proto.RegisterType((*WeixinInfo)(nil), "mjproto.WeixinInfo")
	proto.RegisterType((*CardInfo)(nil), "mjproto.CardInfo")
	proto.RegisterType((*PlayOptions)(nil), "mjproto.PlayOptions")
	proto.RegisterType((*RoomTypeInfo)(nil), "mjproto.RoomTypeInfo")
	proto.RegisterType((*ComposePoker)(nil), "mjproto.ComposePoker")
	proto.RegisterType((*PlayerPoker)(nil), "mjproto.PlayerPoker")
	proto.RegisterType((*PlayerInfo)(nil), "mjproto.PlayerInfo")
	proto.RegisterType((*DeskGameInfo)(nil), "mjproto.DeskGameInfo")
	proto.RegisterEnum("mjproto.EProtoId", EProtoId_name, EProtoId_value)
	proto.RegisterEnum("mjproto.DDErrorCode", DDErrorCode_name, DDErrorCode_value)
	proto.RegisterEnum("mjproto.MJRoomType", MJRoomType_name, MJRoomType_value)
}

var fileDescriptor0 = []byte{
	// 1276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x56, 0xdd, 0x72, 0xdb, 0x44,
	0x14, 0xc6, 0x71, 0x9c, 0xd8, 0x6b, 0x27, 0xd9, 0xa8, 0x4d, 0xeb, 0x36, 0x81, 0x16, 0xb7, 0x4c,
	0xd3, 0x32, 0xd3, 0x81, 0xc2, 0x0b, 0x28, 0xd2, 0xd6, 0x56, 0x63, 0x4b, 0xae, 0x64, 0x15, 0xca,
	0x8d, 0x47, 0xb1, 0x96, 0x44, 0x4d, 0x2c, 0x79, 0xf4, 0xd3, 0x26, 0xbc, 0x45, 0xb9, 0x03, 0x5e,
	0x80, 0x5b, 0x78, 0x01, 0xde, 0x8b, 0x1b, 0x38, 0x7b, 0x56, 0x6b, 0x39, 0x90, 0x4c, 0x3c, 0x23,
	0x7d, 0xfb, 0x9d, 0x73, 0xbe, 0xfd, 0xce, 0xd1, 0x4a, 0x84, 0x9c, 0x04, 0x19, 0x7f, 0xbe, 0x48,
	0x93, 0x3c, 0xd1, 0x36, 0xe7, 0xef, 0xf0, 0xa2, 0x77, 0x4c, 0xda, 0x63, 0x71, 0x31, 0xe0, 0x41,
	0xc8, 0x53, 0x6d, 0x87, 0x6c, 0xbe, 0xe7, 0x69, 0x16, 0x25, 0x71, 0xb7, 0xf6, 0xb0, 0x76, 0xd8,
	0xd2, 0xb6, 0xc9, 0x46, 0x91, 0xf1, 0xd4, 0x0a, 0xbb, 0x6b, 0x70, 0xbf, 0xa5, 0x75, 0xc8, 0xfa,
	0x2c, 0x09, 0x79, 0xb7, 0x0e, 0x77, 0x0d, 0x6d, 0x8b, 0x34, 0x78, 0x9a, 0x26, 0x69, 0x77, 0x5d,
	0x90, 0x7b, 0xe7, 0x84, 0x7c, 0xc7, 0xa3, 0xcb, 0x28, 0xb6, 0xe2, 0x1f, 0x13, 0x11, 0x9a, 0x2c,
	0x78, 0x0c, 0xa1, 0x32, 0x15, 0x25, 0xcd, 0x38, 0x9a, 0x9d, 0xdb, 0xc1, 0x9c, 0x63, 0xb2, 0x96,
	0xa8, 0x76, 0x06, 0x75, 0xfd, 0xf4, 0x02, 0xf3, 0xb5, 0xb4, 0x36, 0xa9, 0x67, 0xfc, 0x12, 0xb3,
	0x35, 0xb0, 0x54, 0x94, 0x5f, 0x75, 0x1b, 0x8a, 0x5b, 0xc4, 0xa0, 0x0b, 0xd2, 0x6d, 0x60, 0xb1,
	0x6f, 0x48, 0xd3, 0x08, 0xd2, 0x10, 0x4b, 0x01, 0x35, 0xbf, 0x5a, 0x70, 0x2c, 0x84, 0xaa, 0xde,
	0x07, 0x17, 0x85, 0xac, 0xd2, 0xd0, 0x08, 0x59, 0x8b, 0x42, 0x29, 0xb8, 0x37, 0x83, 0xed, 0x5e,
	0x04, 0x57, 0xce, 0x22, 0x87, 0x54, 0x99, 0xb6, 0x4b, 0x5a, 0x3f, 0x45, 0xa3, 0xc4, 0x0d, 0xc2,
	0x28, 0x29, 0x83, 0xbb, 0x84, 0x86, 0x51, 0x10, 0xf7, 0x83, 0xf8, 0x74, 0x50, 0x04, 0x72, 0x45,
	0xe6, 0xb9, 0x43, 0xb6, 0x93, 0xfc, 0x0c, 0xdc, 0x31, 0xce, 0xf8, 0xec, 0xfc, 0x28, 0xb9, 0x84,
	0x9c, 0x75, 0xc0, 0xc5, 0x2e, 0x0a, 0x49, 0x44, 0xe1, 0xbd, 0xbf, 0x6a, 0xa4, 0xe3, 0x26, 0xc9,
	0x7c, 0x02, 0x92, 0x50, 0xde, 0x13, 0x42, 0xe6, 0xef, 0x14, 0x82, 0x75, 0xb6, 0x5f, 0xdc, 0x7a,
	0x5e, 0xb6, 0xe0, 0xf9, 0xe8, 0x95, 0x5a, 0xd2, 0x40, 0xeb, 0x49, 0x02, 0x9b, 0xca, 0x8c, 0xa4,
	0xc8, 0xcb, 0xb2, 0x60, 0xe3, 0x2c, 0x58, 0x8c, 0x82, 0x4b, 0xdc, 0x42, 0x5d, 0x7b, 0x4a, 0xda,
	0x8b, 0x6a, 0x0b, 0x58, 0xb2, 0xfd, 0xe2, 0xf6, 0x32, 0xdb, 0xea, 0xf6, 0xc0, 0xf1, 0x99, 0xc8,
	0x66, 0x17, 0x73, 0x74, 0xb1, 0x21, 0x0a, 0x64, 0x3c, 0xcf, 0x2f, 0xf8, 0x9c, 0xc7, 0x39, 0x1a,
	0xd9, 0x10, 0x26, 0x88, 0xc9, 0x78, 0x83, 0x96, 0x6d, 0x8a, 0x1a, 0xbd, 0xa7, 0xa4, 0x63, 0x24,
	0xf3, 0x45, 0x92, 0xf1, 0x71, 0x72, 0x0e, 0x63, 0x01, 0x7d, 0x89, 0x21, 0x47, 0x4d, 0xf5, 0x05,
	0xcd, 0x46, 0x79, 0xbd, 0x85, 0x74, 0x94, 0xa7, 0x92, 0x09, 0xc9, 0xce, 0x82, 0x38, 0xc4, 0x1b,
	0xe0, 0x0b, 0x7f, 0xbe, 0x24, 0x9d, 0xd9, 0x4a, 0x32, 0x88, 0xab, 0x83, 0xe2, 0xbd, 0xa5, 0xe2,
	0x6b, 0x95, 0x40, 0x32, 0x6c, 0x5d, 0x12, 0x57, 0xec, 0x95, 0x80, 0xb4, 0xf7, 0xef, 0x1a, 0x21,
	0xb2, 0x24, 0x9a, 0x0b, 0x11, 0x51, 0x76, 0x14, 0xc4, 0xb2, 0x60, 0xed, 0xb0, 0xa9, 0x1c, 0x2a,
	0x25, 0xa1, 0xce, 0xff, 0x3a, 0xa4, 0xe4, 0xe2, 0x38, 0x47, 0x71, 0x69, 0xed, 0xea, 0x84, 0xae,
	0xab, 0xf1, 0xcf, 0x78, 0x90, 0xc3, 0xd0, 0x49, 0xff, 0x0e, 0x49, 0xfb, 0xc3, 0x72, 0xc2, 0x33,
	0x34, 0xb0, 0xbd, 0xd2, 0xca, 0x95, 0xe9, 0x07, 0xd9, 0x51, 0xe6, 0x7c, 0x88, 0x41, 0xc0, 0x26,
	0xaa, 0x82, 0x54, 0x27, 0x2e, 0x4c, 0xfb, 0x55, 0xb7, 0x89, 0xa9, 0xa0, 0xd8, 0x89, 0x19, 0xc5,
	0xa7, 0xaf, 0xc1, 0xf5, 0x96, 0x6a, 0xce, 0x09, 0xbb, 0x9c, 0x81, 0x7f, 0xa7, 0x3c, 0xec, 0x12,
	0xd5, 0xfd, 0x78, 0x50, 0x8c, 0x83, 0xa8, 0xdb, 0xc6, 0xcd, 0xff, 0xbe, 0x46, 0x3a, 0x26, 0xcf,
	0xce, 0xfb, 0xa0, 0x11, 0xeb, 0x40, 0x90, 0xb8, 0xf6, 0xf2, 0x20, 0x2f, 0xb2, 0xb2, 0x43, 0xe0,
	0x78, 0xba, 0x32, 0x7f, 0xa5, 0x03, 0x95, 0xe3, 0xd7, 0x86, 0x13, 0x3a, 0x26, 0xdd, 0x12, 0x53,
	0x52, 0x57, 0x42, 0x82, 0x59, 0x1e, 0xbd, 0xe7, 0x1e, 0xec, 0xbd, 0x7c, 0x1a, 0x4b, 0x2c, 0x89,
	0x27, 0x11, 0xb8, 0xd3, 0x50, 0x93, 0x13, 0x72, 0x88, 0x45, 0x48, 0x0e, 0xd3, 0x2d, 0xd2, 0x8e,
	0x5d, 0x7e, 0x52, 0x5c, 0xc1, 0x04, 0xc3, 0x84, 0x6d, 0x22, 0x78, 0x97, 0xec, 0x80, 0x29, 0x51,
	0xae, 0x57, 0x09, 0x9a, 0xea, 0x91, 0xc2, 0x05, 0x73, 0x99, 0x45, 0x3a, 0x71, 0x9b, 0x74, 0x22,
	0x80, 0x85, 0x4e, 0x43, 0xb4, 0x87, 0x60, 0x7b, 0xf6, 0xc8, 0xd6, 0xac, 0x48, 0x53, 0xd1, 0x3f,
	0x99, 0xbd, 0xad, 0x92, 0xe4, 0x49, 0x1e, 0x5c, 0x54, 0x78, 0x47, 0xe0, 0xcf, 0xfe, 0xdc, 0x20,
	0x4d, 0x86, 0x87, 0x9b, 0x15, 0x82, 0xfc, 0xed, 0xb1, 0x65, 0x4e, 0x5f, 0xfb, 0x96, 0x71, 0x3c,
	0x35, 0x1c, 0xdb, 0xa6, 0x35, 0x08, 0xd4, 0xae, 0x63, 0x53, 0xdd, 0x38, 0xa6, 0x6b, 0x8a, 0xdb,
	0xd7, 0x47, 0x6c, 0x3a, 0x74, 0xfa, 0x96, 0x4d, 0xeb, 0x8a, 0x5b, 0x61, 0xc8, 0x5d, 0x57, 0x5c,
	0xc3, 0x65, 0xfa, 0x84, 0xb9, 0x8e, 0x33, 0xa2, 0x0d, 0xc5, 0xad, 0x30, 0xe4, 0x6e, 0x28, 0x2e,
	0xb3, 0x01, 0x9d, 0x22, 0x77, 0x53, 0x71, 0x2b, 0x0c, 0xb9, 0x4d, 0xd8, 0xeb, 0xae, 0xc0, 0x3d,
	0x66, 0xcb, 0xa2, 0x96, 0xfd, 0xd2, 0xa1, 0x2d, 0x38, 0xda, 0x5a, 0x02, 0x86, 0xcc, 0xe6, 0x5b,
	0x4a, 0xa0, 0x01, 0x5b, 0xcb, 0x5b, 0x0c, 0x6c, 0xab, 0x40, 0xf6, 0xbd, 0x31, 0xd0, 0xed, 0x3e,
	0x33, 0x74, 0xd7, 0xf4, 0x68, 0x47, 0xbb, 0x47, 0xf6, 0xfe, 0x07, 0x63, 0xc4, 0x16, 0x4c, 0x6a,
	0x5b, 0x2c, 0x99, 0x96, 0xdd, 0x7f, 0xed, 0x33, 0xba, 0xad, 0x00, 0x67, 0xcc, 0x6c, 0xc0, 0xe8,
	0x8e, 0x12, 0x6e, 0x32, 0x7d, 0x38, 0x95, 0x09, 0x29, 0x34, 0x7a, 0x07, 0x0d, 0x61, 0x93, 0x29,
	0x98, 0x21, 0x50, 0xba, 0x7b, 0x4d, 0xb5, 0xe3, 0x4f, 0x24, 0xac, 0x29, 0x99, 0x90, 0xaf, 0x2f,
	0xa1, 0x5b, 0x8a, 0xb9, 0x84, 0x50, 0xcb, 0x6d, 0xc5, 0xec, 0xeb, 0x8a, 0xb9, 0xa7, 0x98, 0x4b,
	0x08, 0x99, 0x77, 0xe0, 0xf1, 0xe9, 0x20, 0xec, 0x3b, 0x92, 0x78, 0x17, 0x86, 0x86, 0xae, 0x22,
	0xc8, 0xeb, 0xaa, 0xcd, 0x0c, 0x7c, 0x49, 0xbb, 0xa7, 0x84, 0x97, 0x00, 0xb2, 0xee, 0x6b, 0x0f,
	0xc8, 0xbe, 0x00, 0x8f, 0x5c, 0x47, 0x37, 0x0d, 0xdd, 0x9b, 0x4c, 0x8f, 0x98, 0xe8, 0xb1, 0xf2,
	0x64, 0x5f, 0x7b, 0x48, 0x0e, 0x6e, 0x22, 0x28, 0x3f, 0xe9, 0x81, 0x12, 0xe4, 0xbc, 0x61, 0xee,
	0xc4, 0x77, 0x6d, 0xfa, 0xa9, 0x52, 0x6e, 0xf8, 0xae, 0x0b, 0xfd, 0x75, 0x99, 0xe7, 0x0f, 0x27,
	0xf4, 0x33, 0xd5, 0x72, 0x61, 0x12, 0xfc, 0x0f, 0x9d, 0x09, 0xb4, 0xfe, 0x2d, 0x7d, 0xa0, 0xe8,
	0xa6, 0xe5, 0x79, 0xce, 0xf0, 0x0d, 0x03, 0xbb, 0xbd, 0x63, 0xfa, 0x50, 0x75, 0xee, 0x1a, 0x8c,
	0xaa, 0x3f, 0x57, 0x7d, 0x19, 0x32, 0x5d, 0xd1, 0x7b, 0x2a, 0x7b, 0x85, 0x21, 0xf7, 0x91, 0xf2,
	0x61, 0xc4, 0x3c, 0x4f, 0x07, 0xbd, 0x8f, 0x95, 0x5d, 0xd8, 0x2b, 0x85, 0x7e, 0xf1, 0xec, 0xe3,
	0x1a, 0x69, 0x9b, 0x26, 0x13, 0x2f, 0x75, 0x03, 0xde, 0xf3, 0xa2, 0x04, 0x73, 0x5d, 0xc7, 0x35,
	0x1c, 0x93, 0x4d, 0x3d, 0xdf, 0x30, 0xe8, 0x27, 0xda, 0xb7, 0xe4, 0x49, 0x85, 0xc9, 0x29, 0x97,
	0x95, 0x4c, 0x4b, 0x1f, 0x39, 0x90, 0xce, 0x76, 0x26, 0xcc, 0x76, 0xfc, 0xfe, 0x80, 0xfe, 0xf6,
	0x4f, 0xf9, 0x57, 0xd3, 0xbe, 0x22, 0x8f, 0x6e, 0x8e, 0xf2, 0x3d, 0x18, 0x7d, 0x08, 0x79, 0xe9,
	0xf8, 0xb6, 0x49, 0x7f, 0xad, 0x22, 0x9e, 0x92, 0x83, 0x2a, 0xc2, 0xb2, 0x27, 0x8e, 0xe4, 0x2f,
	0xa9, 0xbf, 0x54, 0xd4, 0x43, 0xb2, 0x5f, 0x51, 0xf1, 0x21, 0x95, 0xcf, 0x84, 0xcb, 0xc6, 0x50,
	0x88, 0xfe, 0x5c, 0x31, 0xbf, 0x26, 0x8f, 0x6f, 0x64, 0x1a, 0x03, 0x6b, 0x2c, 0x52, 0x4f, 0x4b,
	0xe5, 0x1f, 0x97, 0x21, 0xcf, 0xfe, 0x80, 0x17, 0xce, 0xca, 0x4b, 0x1a, 0x1a, 0xa2, 0x4e, 0xd7,
	0xe9, 0x65, 0xc1, 0x7f, 0x80, 0xf3, 0xda, 0x0c, 0x12, 0x33, 0x02, 0x67, 0x0e, 0x48, 0x77, 0xb9,
	0x94, 0x05, 0x70, 0x10, 0xc6, 0x43, 0xf8, 0x96, 0x38, 0x7d, 0x09, 0x3f, 0x38, 0x6f, 0xf6, 0xc9,
	0xdd, 0x6a, 0x35, 0xba, 0xb6, 0xb8, 0xa6, 0xdd, 0x27, 0x77, 0x96, 0x8b, 0x21, 0x7f, 0x0b, 0xe0,
	0x28, 0x78, 0x25, 0xd6, 0xa9, 0x38, 0xf8, 0x76, 0xab, 0x35, 0x28, 0x15, 0x24, 0x83, 0x02, 0xce,
	0x9e, 0xd5, 0x10, 0x10, 0x32, 0x8c, 0x0a, 0xf8, 0x32, 0x81, 0xcf, 0x16, 0x4e, 0x1b, 0xff, 0x06,
	0x00, 0x00, 0xff, 0xff, 0xa6, 0x27, 0x27, 0xb5, 0xed, 0x09, 0x00, 0x00,
}
