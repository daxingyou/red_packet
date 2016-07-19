// Code generated by protoc-gen-go.
// source: thPoker.proto
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

// Ignoring public import of User from user.proto

// Ignoring public import of ZjhRoom from zhajinhua.proto

// Ignoring public import of ZjhBet from zhajinhua.proto

// Ignoring public import of ZjhLottery from zhajinhua.proto

// Ignoring public import of BroadcastBet from zhajinhua.proto

// Ignoring public import of ZjhPai from zhajinhua.proto

// Ignoring public import of Pai from zhajinhua.proto

// Ignoring public import of ZjhQueryNoSeatUser from zhajinhua.proto

// Ignoring public import of ZjhReqSeat from zhajinhua.proto

// Ignoring public import of ZjhMsg from zhajinhua.proto

// Ignoring public import of ZjhBroadcastBeginBet from zhajinhua.proto

// Ignoring public import of EPaiType from zhajinhua.proto

// 德州扑克的
type ETHType int32

const (
	ETHType_TH_TYPE_GAOPAI              ETHType = 1
	ETHType_TH_TYPE_YIDUI               ETHType = 2
	ETHType_TH_TYPE_LIANGDUI            ETHType = 3
	ETHType_TH_TYPE_THREE               ETHType = 4
	ETHType_TH_TYPE_SHUNZI              ETHType = 5
	ETHType_TH_TYPE_TONGHUA             ETHType = 6
	ETHType_TH_TYPE_HULU                ETHType = 7
	ETHType_TH_TYPE_SITIAO              ETHType = 8
	ETHType_TH_TYPE_TONGHUASHUN         ETHType = 9
	ETHType_TH_TYPE_HUANGSHITONGHUASHUN ETHType = 10
)

var ETHType_name = map[int32]string{
	1:  "TH_TYPE_GAOPAI",
	2:  "TH_TYPE_YIDUI",
	3:  "TH_TYPE_LIANGDUI",
	4:  "TH_TYPE_THREE",
	5:  "TH_TYPE_SHUNZI",
	6:  "TH_TYPE_TONGHUA",
	7:  "TH_TYPE_HULU",
	8:  "TH_TYPE_SITIAO",
	9:  "TH_TYPE_TONGHUASHUN",
	10: "TH_TYPE_HUANGSHITONGHUASHUN",
}
var ETHType_value = map[string]int32{
	"TH_TYPE_GAOPAI":              1,
	"TH_TYPE_YIDUI":               2,
	"TH_TYPE_LIANGDUI":            3,
	"TH_TYPE_THREE":               4,
	"TH_TYPE_SHUNZI":              5,
	"TH_TYPE_TONGHUA":             6,
	"TH_TYPE_HULU":                7,
	"TH_TYPE_SITIAO":              8,
	"TH_TYPE_TONGHUASHUN":         9,
	"TH_TYPE_HUANGSHITONGHUASHUN": 10,
}

func (x ETHType) Enum() *ETHType {
	p := new(ETHType)
	*p = x
	return p
}
func (x ETHType) String() string {
	return proto.EnumName(ETHType_name, int32(x))
}
func (x *ETHType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ETHType_value, data, "ETHType")
	if err != nil {
		return err
	}
	*x = ETHType(value)
	return nil
}
func (ETHType) EnumDescriptor() ([]byte, []int) { return fileDescriptor10, []int{0} }

type ETHAction int32

const (
	ETHAction_TH_DESK_BET_TYPE_BET      ETHAction = 1
	ETHAction_TH_DESK_BET_TYPE_CALL     ETHAction = 2
	ETHAction_TH_DESK_BET_TYPE_FOLD     ETHAction = 3
	ETHAction_TH_DESK_BET_TYPE_CHECK    ETHAction = 4
	ETHAction_TH_DESK_BET_TYPE_RAISE    ETHAction = 5
	ETHAction_TH_DESK_BET_TYPE_RERRAISE ETHAction = 6
	ETHAction_TH_DESK_BET_TYPE_ALLIN    ETHAction = 7
)

var ETHAction_name = map[int32]string{
	1: "TH_DESK_BET_TYPE_BET",
	2: "TH_DESK_BET_TYPE_CALL",
	3: "TH_DESK_BET_TYPE_FOLD",
	4: "TH_DESK_BET_TYPE_CHECK",
	5: "TH_DESK_BET_TYPE_RAISE",
	6: "TH_DESK_BET_TYPE_RERRAISE",
	7: "TH_DESK_BET_TYPE_ALLIN",
}
var ETHAction_value = map[string]int32{
	"TH_DESK_BET_TYPE_BET":      1,
	"TH_DESK_BET_TYPE_CALL":     2,
	"TH_DESK_BET_TYPE_FOLD":     3,
	"TH_DESK_BET_TYPE_CHECK":    4,
	"TH_DESK_BET_TYPE_RAISE":    5,
	"TH_DESK_BET_TYPE_RERRAISE": 6,
	"TH_DESK_BET_TYPE_ALLIN":    7,
}

func (x ETHAction) Enum() *ETHAction {
	p := new(ETHAction)
	*p = x
	return p
}
func (x ETHAction) String() string {
	return proto.EnumName(ETHAction_name, int32(x))
}
func (x *ETHAction) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ETHAction_value, data, "ETHAction")
	if err != nil {
		return err
	}
	*x = ETHAction(value)
	return nil
}
func (ETHAction) EnumDescriptor() ([]byte, []int) { return fileDescriptor10, []int{1} }

type ThRoom struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	ReqType          *int32       `protobuf:"varint,2,opt,name=reqType" json:"reqType,omitempty"`
	Jackpot          *int64       `protobuf:"varint,3,opt,name=jackpot" json:"jackpot,omitempty"`
	Users            []*THUser    `protobuf:"bytes,4,rep,name=users" json:"users,omitempty"`
	BetTime          *int32       `protobuf:"varint,5,opt,name=betTime" json:"betTime,omitempty"`
	DeskStatus       *int32       `protobuf:"varint,6,opt,name=deskStatus" json:"deskStatus,omitempty"`
	PublicPais       []*Pai       `protobuf:"bytes,7,rep,name=publicPais" json:"publicPais,omitempty"`
	BetUserId        *uint32      `protobuf:"varint,8,opt,name=betUserId" json:"betUserId,omitempty"`
	BigBlind         *uint32      `protobuf:"varint,9,opt,name=bigBlind" json:"bigBlind,omitempty"`
	SmallBlind       *uint32      `protobuf:"varint,10,opt,name=smallBlind" json:"smallBlind,omitempty"`
	Butten           *uint32      `protobuf:"varint,11,opt,name=butten" json:"butten,omitempty"`
	BetRemainTime    *int32       `protobuf:"varint,12,opt,name=betRemainTime" json:"betRemainTime,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *ThRoom) Reset()                    { *m = ThRoom{} }
func (m *ThRoom) String() string            { return proto.CompactTextString(m) }
func (*ThRoom) ProtoMessage()               {}
func (*ThRoom) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{0} }

func (m *ThRoom) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ThRoom) GetReqType() int32 {
	if m != nil && m.ReqType != nil {
		return *m.ReqType
	}
	return 0
}

func (m *ThRoom) GetJackpot() int64 {
	if m != nil && m.Jackpot != nil {
		return *m.Jackpot
	}
	return 0
}

func (m *ThRoom) GetUsers() []*THUser {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *ThRoom) GetBetTime() int32 {
	if m != nil && m.BetTime != nil {
		return *m.BetTime
	}
	return 0
}

func (m *ThRoom) GetDeskStatus() int32 {
	if m != nil && m.DeskStatus != nil {
		return *m.DeskStatus
	}
	return 0
}

func (m *ThRoom) GetPublicPais() []*Pai {
	if m != nil {
		return m.PublicPais
	}
	return nil
}

func (m *ThRoom) GetBetUserId() uint32 {
	if m != nil && m.BetUserId != nil {
		return *m.BetUserId
	}
	return 0
}

func (m *ThRoom) GetBigBlind() uint32 {
	if m != nil && m.BigBlind != nil {
		return *m.BigBlind
	}
	return 0
}

func (m *ThRoom) GetSmallBlind() uint32 {
	if m != nil && m.SmallBlind != nil {
		return *m.SmallBlind
	}
	return 0
}

func (m *ThRoom) GetButten() uint32 {
	if m != nil && m.Butten != nil {
		return *m.Butten
	}
	return 0
}

func (m *ThRoom) GetBetRemainTime() int32 {
	if m != nil && m.BetRemainTime != nil {
		return *m.BetRemainTime
	}
	return 0
}

// 刚进来的人
type THRoomAddUserBroadcast struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	User             *THUser      `protobuf:"bytes,2,opt,name=user" json:"user,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *THRoomAddUserBroadcast) Reset()                    { *m = THRoomAddUserBroadcast{} }
func (m *THRoomAddUserBroadcast) String() string            { return proto.CompactTextString(m) }
func (*THRoomAddUserBroadcast) ProtoMessage()               {}
func (*THRoomAddUserBroadcast) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{1} }

func (m *THRoomAddUserBroadcast) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *THRoomAddUserBroadcast) GetUser() *THUser {
	if m != nil {
		return m.User
	}
	return nil
}

type THPoker struct {
	ThType           *int32 `protobuf:"varint,1,opt,name=thType" json:"thType,omitempty"`
	Pais             []*Pai `protobuf:"bytes,2,rep,name=pais" json:"pais,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *THPoker) Reset()                    { *m = THPoker{} }
func (m *THPoker) String() string            { return proto.CompactTextString(m) }
func (*THPoker) ProtoMessage()               {}
func (*THPoker) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{2} }

func (m *THPoker) GetThType() int32 {
	if m != nil && m.ThType != nil {
		return *m.ThType
	}
	return 0
}

func (m *THPoker) GetPais() []*Pai {
	if m != nil {
		return m.Pais
	}
	return nil
}

type THBet struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	BetType          *int32       `protobuf:"varint,2,opt,name=betType" json:"betType,omitempty"`
	BetAmount        *int32       `protobuf:"varint,3,opt,name=betAmount" json:"betAmount,omitempty"`
	NextBetUser      *uint32      `protobuf:"varint,4,opt,name=nextBetUser" json:"nextBetUser,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *THBet) Reset()                    { *m = THBet{} }
func (m *THBet) String() string            { return proto.CompactTextString(m) }
func (*THBet) ProtoMessage()               {}
func (*THBet) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{3} }

func (m *THBet) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *THBet) GetBetType() int32 {
	if m != nil && m.BetType != nil {
		return *m.BetType
	}
	return 0
}

func (m *THBet) GetBetAmount() int32 {
	if m != nil && m.BetAmount != nil {
		return *m.BetAmount
	}
	return 0
}

func (m *THBet) GetNextBetUser() uint32 {
	if m != nil && m.NextBetUser != nil {
		return *m.NextBetUser
	}
	return 0
}

type THBetBegin struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Dealer           *uint32      `protobuf:"varint,2,opt,name=dealer" json:"dealer,omitempty"`
	BigBlind         *THUser      `protobuf:"bytes,3,opt,name=bigBlind" json:"bigBlind,omitempty"`
	SmallBlind       *THUser      `protobuf:"bytes,4,opt,name=smallBlind" json:"smallBlind,omitempty"`
	BetUserNow       *uint32      `protobuf:"varint,5,opt,name=betUserNow" json:"betUserNow,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *THBetBegin) Reset()                    { *m = THBetBegin{} }
func (m *THBetBegin) String() string            { return proto.CompactTextString(m) }
func (*THBetBegin) ProtoMessage()               {}
func (*THBetBegin) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{4} }

func (m *THBetBegin) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *THBetBegin) GetDealer() uint32 {
	if m != nil && m.Dealer != nil {
		return *m.Dealer
	}
	return 0
}

func (m *THBetBegin) GetBigBlind() *THUser {
	if m != nil {
		return m.BigBlind
	}
	return nil
}

func (m *THBetBegin) GetSmallBlind() *THUser {
	if m != nil {
		return m.SmallBlind
	}
	return nil
}

func (m *THBetBegin) GetBetUserNow() uint32 {
	if m != nil && m.BetUserNow != nil {
		return *m.BetUserNow
	}
	return 0
}

type THBetBroadcast struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	BetType          *int32       `protobuf:"varint,2,opt,name=betType" json:"betType,omitempty"`
	BetAmount        *int32       `protobuf:"varint,3,opt,name=betAmount" json:"betAmount,omitempty"`
	User             *THUser      `protobuf:"bytes,4,opt,name=user" json:"user,omitempty"`
	NextBetUserId    *uint32      `protobuf:"varint,5,opt,name=nextBetUserId" json:"nextBetUserId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *THBetBroadcast) Reset()                    { *m = THBetBroadcast{} }
func (m *THBetBroadcast) String() string            { return proto.CompactTextString(m) }
func (*THBetBroadcast) ProtoMessage()               {}
func (*THBetBroadcast) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{5} }

func (m *THBetBroadcast) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *THBetBroadcast) GetBetType() int32 {
	if m != nil && m.BetType != nil {
		return *m.BetType
	}
	return 0
}

func (m *THBetBroadcast) GetBetAmount() int32 {
	if m != nil && m.BetAmount != nil {
		return *m.BetAmount
	}
	return 0
}

func (m *THBetBroadcast) GetUser() *THUser {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *THBetBroadcast) GetNextBetUserId() uint32 {
	if m != nil && m.NextBetUserId != nil {
		return *m.NextBetUserId
	}
	return 0
}

// 德州扑克中的User
type THUser struct {
	User             *User    `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	HandPais         []*Pai   `protobuf:"bytes,2,rep,name=handPais" json:"handPais,omitempty"`
	Status           *int32   `protobuf:"varint,3,opt,name=status" json:"status,omitempty"`
	SeatNumber       *int32   `protobuf:"varint,4,opt,name=seatNumber" json:"seatNumber,omitempty"`
	Thpoker          *THPoker `protobuf:"bytes,5,opt,name=thpoker" json:"thpoker,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *THUser) Reset()                    { *m = THUser{} }
func (m *THUser) String() string            { return proto.CompactTextString(m) }
func (*THUser) ProtoMessage()               {}
func (*THUser) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{6} }

func (m *THUser) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *THUser) GetHandPais() []*Pai {
	if m != nil {
		return m.HandPais
	}
	return nil
}

func (m *THUser) GetStatus() int32 {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return 0
}

func (m *THUser) GetSeatNumber() int32 {
	if m != nil && m.SeatNumber != nil {
		return *m.SeatNumber
	}
	return 0
}

func (m *THUser) GetThpoker() *THPoker {
	if m != nil {
		return m.Thpoker
	}
	return nil
}

// 开奖时需要用到的协议
type THLottery struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Users            []*THUser    `protobuf:"bytes,2,rep,name=users" json:"users,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *THLottery) Reset()                    { *m = THLottery{} }
func (m *THLottery) String() string            { return proto.CompactTextString(m) }
func (*THLottery) ProtoMessage()               {}
func (*THLottery) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{7} }

func (m *THLottery) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *THLottery) GetUsers() []*THUser {
	if m != nil {
		return m.Users
	}
	return nil
}

func init() {
	proto.RegisterType((*ThRoom)(nil), "bbproto.ThRoom")
	proto.RegisterType((*THRoomAddUserBroadcast)(nil), "bbproto.THRoomAddUserBroadcast")
	proto.RegisterType((*THPoker)(nil), "bbproto.THPoker")
	proto.RegisterType((*THBet)(nil), "bbproto.THBet")
	proto.RegisterType((*THBetBegin)(nil), "bbproto.THBetBegin")
	proto.RegisterType((*THBetBroadcast)(nil), "bbproto.THBetBroadcast")
	proto.RegisterType((*THUser)(nil), "bbproto.THUser")
	proto.RegisterType((*THLottery)(nil), "bbproto.THLottery")
	proto.RegisterEnum("bbproto.ETHType", ETHType_name, ETHType_value)
	proto.RegisterEnum("bbproto.ETHAction", ETHAction_name, ETHAction_value)
}

var fileDescriptor10 = []byte{
	// 726 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0xcd, 0x6e, 0xda, 0x4a,
	0x14, 0xbe, 0xe6, 0xcf, 0x70, 0x08, 0xc1, 0x77, 0x92, 0xdc, 0xeb, 0x10, 0x35, 0x4d, 0x69, 0x17,
	0x55, 0x16, 0x2c, 0x22, 0xf5, 0x01, 0x4c, 0xe2, 0xc6, 0x56, 0x10, 0xb8, 0x60, 0x16, 0xa9, 0x54,
	0x45, 0x63, 0x3c, 0x8d, 0x9d, 0x80, 0x4d, 0xf1, 0xa0, 0x36, 0x7d, 0x86, 0x2e, 0xba, 0xeb, 0xae,
	0x8f, 0xd4, 0x87, 0xe8, 0x93, 0xf4, 0xcc, 0x18, 0x08, 0x09, 0x44, 0xa2, 0x9b, 0xc8, 0x39, 0x3f,
	0xdf, 0x7c, 0xe7, 0xfb, 0xce, 0x01, 0x2a, 0x3c, 0x70, 0xe2, 0x5b, 0x36, 0x69, 0x8c, 0x27, 0x31,
	0x8f, 0x89, 0xea, 0x79, 0xf2, 0xa3, 0x06, 0x1e, 0x4d, 0x58, 0x63, 0xf6, 0x3d, 0x4d, 0xe6, 0x05,
	0xb5, 0xea, 0xd7, 0x80, 0xde, 0x84, 0x51, 0x30, 0xa5, 0x69, 0xa0, 0xfe, 0x23, 0x03, 0x05, 0x37,
	0xe8, 0xc6, 0xf1, 0x88, 0xbc, 0x82, 0x42, 0xc0, 0xa8, 0xcf, 0x26, 0xba, 0x72, 0xa4, 0xbc, 0x2e,
	0x9f, 0xec, 0x36, 0x66, 0x68, 0x0d, 0x47, 0xfc, 0xb5, 0x64, 0x8e, 0x54, 0x41, 0x9d, 0xb0, 0x4f,
	0xee, 0xdd, 0x98, 0xe9, 0x19, 0x2c, 0xcb, 0x8b, 0xc0, 0x0d, 0x1d, 0xdc, 0x8e, 0x63, 0xae, 0x67,
	0x31, 0x90, 0x25, 0x87, 0x90, 0x17, 0x2f, 0x26, 0x7a, 0xee, 0x28, 0x8b, 0x30, 0xd5, 0x05, 0x8c,
	0x6b, 0xf5, 0x93, 0x14, 0xc1, 0x63, 0xdc, 0x0d, 0x47, 0x4c, 0xcf, 0x4b, 0x04, 0x02, 0xe0, 0xb3,
	0xe4, 0xb6, 0xc7, 0x29, 0x9f, 0x26, 0x7a, 0x41, 0xc6, 0x8e, 0x00, 0xc6, 0x53, 0x6f, 0x18, 0x0e,
	0x1c, 0x1a, 0x26, 0xba, 0x2a, 0x91, 0xb6, 0xee, 0x09, 0xd1, 0x90, 0xfc, 0x0b, 0x25, 0x84, 0x11,
	0x88, 0xb6, 0xaf, 0x17, 0xb1, 0xa9, 0x42, 0x34, 0x28, 0x7a, 0xe1, 0x75, 0x73, 0x18, 0x46, 0xbe,
	0x5e, 0x92, 0x11, 0x84, 0x4e, 0x46, 0x74, 0x38, 0x4c, 0x63, 0x20, 0x63, 0xdb, 0x50, 0xf0, 0xa6,
	0x9c, 0xb3, 0x48, 0x2f, 0xcb, 0xff, 0xf7, 0xa0, 0x82, 0x40, 0x5d, 0x36, 0xa2, 0x61, 0x24, 0x59,
	0x6d, 0x09, 0x06, 0xf5, 0x0f, 0xf0, 0x9f, 0x6b, 0x09, 0x61, 0x0c, 0xdf, 0x17, 0xaf, 0x34, 0x27,
	0x31, 0xf5, 0x07, 0x34, 0xe1, 0x1b, 0x0a, 0xf5, 0x0c, 0x72, 0x42, 0x06, 0xa9, 0xd2, 0xaa, 0x0a,
	0xf5, 0x37, 0xa0, 0xba, 0x96, 0xf4, 0x4e, 0x10, 0xe2, 0x81, 0x54, 0x54, 0x91, 0xb3, 0xd7, 0x20,
	0x37, 0x16, 0x53, 0x67, 0x56, 0xa7, 0xae, 0x7f, 0x84, 0xbc, 0x6b, 0x35, 0x19, 0xdf, 0xdc, 0x2d,
	0xa1, 0xf5, 0xbd, 0x5b, 0xa9, 0x6a, 0xc6, 0x28, 0x9e, 0x46, 0xa9, 0x5f, 0x79, 0xb2, 0x03, 0xe5,
	0x88, 0x7d, 0xe1, 0xcd, 0x54, 0x4c, 0x74, 0x0d, 0x45, 0xa9, 0xff, 0x54, 0x00, 0xe4, 0x43, 0x4d,
	0x76, 0x1d, 0x46, 0x1b, 0xbe, 0x86, 0x83, 0xf8, 0x8c, 0x0e, 0x67, 0x43, 0x57, 0xc8, 0x8b, 0x25,
	0x3f, 0xb2, 0x6b, 0x65, 0x20, 0x2f, 0x1f, 0x18, 0x94, 0x5b, 0x5f, 0x84, 0x2e, 0xce, 0xac, 0x6e,
	0xc7, 0x9f, 0xe5, 0xd2, 0x54, 0xea, 0xdf, 0x15, 0xd8, 0x4e, 0x09, 0xfe, 0xa5, 0x2f, 0x9b, 0x48,
	0x32, 0xf7, 0xee, 0x09, 0x3e, 0xb8, 0x31, 0x4b, 0x8a, 0xe1, 0xfa, 0xa5, 0x94, 0xbe, 0x29, 0x78,
	0x4b, 0x69, 0xc5, 0xc1, 0x0c, 0x20, 0x25, 0x52, 0x59, 0x00, 0xc8, 0xe4, 0x21, 0x14, 0x03, 0x1a,
	0xf9, 0xce, 0x13, 0x1e, 0x0b, 0x19, 0x93, 0xf4, 0x16, 0xb2, 0xf3, 0xfb, 0x48, 0x18, 0xe5, 0xed,
	0xe9, 0xc8, 0x9b, 0x71, 0xca, 0xa3, 0xb4, 0x2a, 0x0f, 0xc6, 0x62, 0x7d, 0xe4, 0xe3, 0xe5, 0x13,
	0x6d, 0x89, 0xa4, 0x5c, 0xab, 0xfa, 0x3b, 0x28, 0xb9, 0x56, 0x2b, 0xc6, 0x4d, 0x9f, 0xdc, 0x6d,
	0xa8, 0xcd, 0xe2, 0x74, 0x33, 0x6b, 0x4f, 0xf7, 0xf8, 0xb7, 0x02, 0xaa, 0xe9, 0x5a, 0x42, 0x3c,
	0x64, 0x85, 0xfa, 0x5f, 0xb9, 0x97, 0x8e, 0x79, 0x75, 0x6e, 0x74, 0x1c, 0xc3, 0xd6, 0x14, 0x94,
	0xb2, 0x32, 0x8f, 0x5d, 0xda, 0x67, 0x7d, 0x5b, 0xcb, 0x90, 0x5d, 0xd0, 0xe6, 0xa1, 0x96, 0x6d,
	0xb4, 0xcf, 0x45, 0x34, 0xbb, 0x5c, 0x88, 0x47, 0x66, 0x9a, 0x5a, 0x6e, 0x19, 0xaf, 0x67, 0xf5,
	0xdb, 0xef, 0x6d, 0x4d, 0xac, 0x66, 0x75, 0x51, 0xd6, 0x69, 0x9f, 0x5b, 0x7d, 0x43, 0x2b, 0xe0,
	0x95, 0x6f, 0xcd, 0x83, 0x56, 0xbf, 0xd5, 0xd7, 0xd4, 0x07, 0xad, 0xb6, 0x6b, 0x1b, 0x1d, 0xad,
	0x48, 0xfe, 0x87, 0x9d, 0x47, 0xad, 0x02, 0x55, 0x2b, 0x91, 0xe7, 0x70, 0x70, 0xdf, 0x8e, 0x84,
	0x7a, 0x96, 0xbd, 0x5c, 0x00, 0xc7, 0xbf, 0x14, 0x28, 0xe1, 0x90, 0xc6, 0x80, 0x87, 0x71, 0x44,
	0x74, 0xd8, 0xc5, 0xf2, 0x33, 0xb3, 0x77, 0x71, 0xd5, 0x34, 0xdd, 0xb4, 0x0f, 0x3f, 0x70, 0xd8,
	0x7d, 0xd8, 0x5b, 0xc9, 0x9c, 0x1a, 0xad, 0x16, 0x0e, 0xbd, 0x2e, 0xf5, 0xb6, 0xd3, 0x3a, 0xc3,
	0xc9, 0x6b, 0xe2, 0x67, 0xe5, 0x71, 0x97, 0x65, 0x9e, 0x5e, 0xa0, 0x04, 0xeb, 0x72, 0x5d, 0xc3,
	0xee, 0x99, 0x9a, 0x58, 0xc9, 0xfd, 0xd5, 0x9c, 0xd9, 0x4d, 0xd3, 0x85, 0xb5, 0xad, 0xc8, 0xc5,
	0x6e, 0x6b, 0xaa, 0xf3, 0x8f, 0xa3, 0x38, 0x99, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb0, 0xe5,
	0xd2, 0xec, 0x2b, 0x06, 0x00, 0x00,
}
