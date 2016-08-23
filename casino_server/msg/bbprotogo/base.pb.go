// Code generated by protoc-gen-go.
// source: base.proto
// DO NOT EDIT!

package bbproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type EProtoId int32

const (
	EProtoId_GETINTOROOM     EProtoId = 1
	EProtoId_SHUIGUOJI       EProtoId = 2
	EProtoId_SHUIGUOJIHILOMP EProtoId = 3
	EProtoId_SHUIGUOJIRES    EProtoId = 4
	// 扎金花相关的逻辑
	EProtoId_ZJHROOM                          EProtoId = 5
	EProtoId_ZJHBET                           EProtoId = 6
	EProtoId_ZJHMSG                           EProtoId = 7
	EProtoId_ZJHQUERYNOSEATUSER               EProtoId = 8
	EProtoId_ZJHREQSEAT                       EProtoId = 9
	EProtoId_ZJHLOTTERY                       EProtoId = 10
	EProtoId_ZJHBROADCASTBEGINBET             EProtoId = 11
	EProtoId_LOGINSIGNINBONUS                 EProtoId = 12
	EProtoId_LOGINTURNTABLEBONUS              EProtoId = 13
	EProtoId_OLINEBONUS                       EProtoId = 14
	EProtoId_TIMINGBONUS                      EProtoId = 15
	EProtoId_THROOM                           EProtoId = 16
	EProtoId_THBET                            EProtoId = 17
	EProtoId_THBETBEGIN                       EProtoId = 18
	EProtoId_THBETBROADCAST                   EProtoId = 19
	EProtoId_THROOMADDUSERBROADCAST           EProtoId = 20
	EProtoId_PID_REQQUICKCONN                 EProtoId = 21
	EProtoId_PID_ACKQUICKCONN                 EProtoId = 22
	EProtoId_PID_NULLMSG                      EProtoId = 23
	EProtoId_PID_MATCHLIST_REQMOBILEMATCHLIST EProtoId = 24
	EProtoId_PID_GAME_LOGINGAME               EProtoId = 25
	EProtoId_PID_GAME_ENTERMATCH              EProtoId = 26
	EProtoId_PID_GAME_ACKENTERMATCH           EProtoId = 27
	EProtoId_PID_GAME_SENDGAMEINFO            EProtoId = 28
	EProtoId_PID_GAME_BLINDCOIN               EProtoId = 29
	EProtoId_PID_GAME_INITCARD                EProtoId = 30
	EProtoId_PID_GAME_SENDFLOPCARD            EProtoId = 31
	EProtoId_PID_GAME_SENDTURNCARD            EProtoId = 32
	EProtoId_PID_GAME_SENDRIVERCARD           EProtoId = 33
	EProtoId_PID_GAME_RAISEBET                EProtoId = 34
	EProtoId_PID_GAME_ACKRAISEBET             EProtoId = 35
	EProtoId_PID_GAME_FOLLOWBET               EProtoId = 36
	EProtoId_PID_GAME_ACKFOLLOWBET            EProtoId = 37
	EProtoId_PID_GAME_FOLDBET                 EProtoId = 38
	EProtoId_PID_GAME_ACKFOLDBET              EProtoId = 39
	EProtoId_PID_GAME_CHECKBET                EProtoId = 40
	EProtoId_PID_GAME_ACKCHECKBET             EProtoId = 41
	EProtoId_PID_GAME_SENDOVERTURN            EProtoId = 42
	EProtoId_PID_GAME_SENDADDUSER             EProtoId = 43
	EProtoId_PID_GAME_GAME_SHOWCARD           EProtoId = 44
	EProtoId_PID_GAME_GAME_ACKSHOWCARD        EProtoId = 45
	EProtoId_PID_GAME_GAME_TESTRESULT         EProtoId = 46
	EProtoId_PID_GAME_PRECOIN                 EProtoId = 47
	EProtoId_PID_GAME_GAMENOTICE              EProtoId = 48
	EProtoId_PID_GAME_GAME_ACKNOTICE          EProtoId = 49
	EProtoId_PID_GAME_GAME_CREATEDESK         EProtoId = 50
	EProtoId_PID_GAME_GAME_ACKCREATEDESK      EProtoId = 51
	EProtoId_PID_GAME_GAME_READY              EProtoId = 52
	EProtoId_PID_GAME_GAME_ACKREADY           EProtoId = 53
	EProtoId_PID_GAME_GAME_BEGIN              EProtoId = 54
	EProtoId_PID_GAME_GAME_GAMERECORD         EProtoId = 55
	EProtoId_PID_GAME_GAME_ACKGAMERECORD      EProtoId = 56
	EProtoId_PID_GAME_GAME_BEANGAMERECORD     EProtoId = 57
	EProtoId_PID_GAME_GAME_DISSOLVEDESK       EProtoId = 58
	EProtoId_PID_GAME_GAME_ACKDISSOLVEDESK    EProtoId = 59
	EProtoId_PID_GAME_LEAVEDESK               EProtoId = 60
	EProtoId_PID_GAME_ACKLEAVEDESK            EProtoId = 61
	EProtoId_PID_GAME_SENDDESKENDLOTTERY      EProtoId = 62
	EProtoId_PID_GAME_MESSAGE                 EProtoId = 63
	EProtoId_PID_GAME_SENDMESSAGE             EProtoId = 64
	EProtoId_PID_GAME_TOUNAMENTBLIND          EProtoId = 65
	EProtoId_PID_GAME_TOUNAMENTREWARDS        EProtoId = 66
	EProtoId_PID_GAME_TOUNAMENTRANK           EProtoId = 67
	EProtoId_PID_GAME_TOUNAMENTSUMMARY        EProtoId = 68
	EProtoId_PID_GAME_MATCHLIST               EProtoId = 69
)

var EProtoId_name = map[int32]string{
	1:  "GETINTOROOM",
	2:  "SHUIGUOJI",
	3:  "SHUIGUOJIHILOMP",
	4:  "SHUIGUOJIRES",
	5:  "ZJHROOM",
	6:  "ZJHBET",
	7:  "ZJHMSG",
	8:  "ZJHQUERYNOSEATUSER",
	9:  "ZJHREQSEAT",
	10: "ZJHLOTTERY",
	11: "ZJHBROADCASTBEGINBET",
	12: "LOGINSIGNINBONUS",
	13: "LOGINTURNTABLEBONUS",
	14: "OLINEBONUS",
	15: "TIMINGBONUS",
	16: "THROOM",
	17: "THBET",
	18: "THBETBEGIN",
	19: "THBETBROADCAST",
	20: "THROOMADDUSERBROADCAST",
	21: "PID_REQQUICKCONN",
	22: "PID_ACKQUICKCONN",
	23: "PID_NULLMSG",
	24: "PID_MATCHLIST_REQMOBILEMATCHLIST",
	25: "PID_GAME_LOGINGAME",
	26: "PID_GAME_ENTERMATCH",
	27: "PID_GAME_ACKENTERMATCH",
	28: "PID_GAME_SENDGAMEINFO",
	29: "PID_GAME_BLINDCOIN",
	30: "PID_GAME_INITCARD",
	31: "PID_GAME_SENDFLOPCARD",
	32: "PID_GAME_SENDTURNCARD",
	33: "PID_GAME_SENDRIVERCARD",
	34: "PID_GAME_RAISEBET",
	35: "PID_GAME_ACKRAISEBET",
	36: "PID_GAME_FOLLOWBET",
	37: "PID_GAME_ACKFOLLOWBET",
	38: "PID_GAME_FOLDBET",
	39: "PID_GAME_ACKFOLDBET",
	40: "PID_GAME_CHECKBET",
	41: "PID_GAME_ACKCHECKBET",
	42: "PID_GAME_SENDOVERTURN",
	43: "PID_GAME_SENDADDUSER",
	44: "PID_GAME_GAME_SHOWCARD",
	45: "PID_GAME_GAME_ACKSHOWCARD",
	46: "PID_GAME_GAME_TESTRESULT",
	47: "PID_GAME_PRECOIN",
	48: "PID_GAME_GAMENOTICE",
	49: "PID_GAME_GAME_ACKNOTICE",
	50: "PID_GAME_GAME_CREATEDESK",
	51: "PID_GAME_GAME_ACKCREATEDESK",
	52: "PID_GAME_GAME_READY",
	53: "PID_GAME_GAME_ACKREADY",
	54: "PID_GAME_GAME_BEGIN",
	55: "PID_GAME_GAME_GAMERECORD",
	56: "PID_GAME_GAME_ACKGAMERECORD",
	57: "PID_GAME_GAME_BEANGAMERECORD",
	58: "PID_GAME_GAME_DISSOLVEDESK",
	59: "PID_GAME_GAME_ACKDISSOLVEDESK",
	60: "PID_GAME_LEAVEDESK",
	61: "PID_GAME_ACKLEAVEDESK",
	62: "PID_GAME_SENDDESKENDLOTTERY",
	63: "PID_GAME_MESSAGE",
	64: "PID_GAME_SENDMESSAGE",
	65: "PID_GAME_TOUNAMENTBLIND",
	66: "PID_GAME_TOUNAMENTREWARDS",
	67: "PID_GAME_TOUNAMENTRANK",
	68: "PID_GAME_TOUNAMENTSUMMARY",
	69: "PID_GAME_MATCHLIST",
}
var EProtoId_value = map[string]int32{
	"GETINTOROOM":                      1,
	"SHUIGUOJI":                        2,
	"SHUIGUOJIHILOMP":                  3,
	"SHUIGUOJIRES":                     4,
	"ZJHROOM":                          5,
	"ZJHBET":                           6,
	"ZJHMSG":                           7,
	"ZJHQUERYNOSEATUSER":               8,
	"ZJHREQSEAT":                       9,
	"ZJHLOTTERY":                       10,
	"ZJHBROADCASTBEGINBET":             11,
	"LOGINSIGNINBONUS":                 12,
	"LOGINTURNTABLEBONUS":              13,
	"OLINEBONUS":                       14,
	"TIMINGBONUS":                      15,
	"THROOM":                           16,
	"THBET":                            17,
	"THBETBEGIN":                       18,
	"THBETBROADCAST":                   19,
	"THROOMADDUSERBROADCAST":           20,
	"PID_REQQUICKCONN":                 21,
	"PID_ACKQUICKCONN":                 22,
	"PID_NULLMSG":                      23,
	"PID_MATCHLIST_REQMOBILEMATCHLIST": 24,
	"PID_GAME_LOGINGAME":               25,
	"PID_GAME_ENTERMATCH":              26,
	"PID_GAME_ACKENTERMATCH":           27,
	"PID_GAME_SENDGAMEINFO":            28,
	"PID_GAME_BLINDCOIN":               29,
	"PID_GAME_INITCARD":                30,
	"PID_GAME_SENDFLOPCARD":            31,
	"PID_GAME_SENDTURNCARD":            32,
	"PID_GAME_SENDRIVERCARD":           33,
	"PID_GAME_RAISEBET":                34,
	"PID_GAME_ACKRAISEBET":             35,
	"PID_GAME_FOLLOWBET":               36,
	"PID_GAME_ACKFOLLOWBET":            37,
	"PID_GAME_FOLDBET":                 38,
	"PID_GAME_ACKFOLDBET":              39,
	"PID_GAME_CHECKBET":                40,
	"PID_GAME_ACKCHECKBET":             41,
	"PID_GAME_SENDOVERTURN":            42,
	"PID_GAME_SENDADDUSER":             43,
	"PID_GAME_GAME_SHOWCARD":           44,
	"PID_GAME_GAME_ACKSHOWCARD":        45,
	"PID_GAME_GAME_TESTRESULT":         46,
	"PID_GAME_PRECOIN":                 47,
	"PID_GAME_GAMENOTICE":              48,
	"PID_GAME_GAME_ACKNOTICE":          49,
	"PID_GAME_GAME_CREATEDESK":         50,
	"PID_GAME_GAME_ACKCREATEDESK":      51,
	"PID_GAME_GAME_READY":              52,
	"PID_GAME_GAME_ACKREADY":           53,
	"PID_GAME_GAME_BEGIN":              54,
	"PID_GAME_GAME_GAMERECORD":         55,
	"PID_GAME_GAME_ACKGAMERECORD":      56,
	"PID_GAME_GAME_BEANGAMERECORD":     57,
	"PID_GAME_GAME_DISSOLVEDESK":       58,
	"PID_GAME_GAME_ACKDISSOLVEDESK":    59,
	"PID_GAME_LEAVEDESK":               60,
	"PID_GAME_ACKLEAVEDESK":            61,
	"PID_GAME_SENDDESKENDLOTTERY":      62,
	"PID_GAME_MESSAGE":                 63,
	"PID_GAME_SENDMESSAGE":             64,
	"PID_GAME_TOUNAMENTBLIND":          65,
	"PID_GAME_TOUNAMENTREWARDS":        66,
	"PID_GAME_TOUNAMENTRANK":           67,
	"PID_GAME_TOUNAMENTSUMMARY":        68,
	"PID_GAME_MATCHLIST":               69,
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
func (EProtoId) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

type DDErrorCode int32

const (
	DDErrorCode_ERRORCODE_SUCC DDErrorCode = 0
	// -101   -200	游戏异常
	DDErrorCode_ERRORCODE_CREATE_DESK_DIAMOND_NOTENOUGH DDErrorCode = -101
	DDErrorCode_ERRORCODE_CREATE_DESK_USER_NOTFOUND     DDErrorCode = -102
	DDErrorCode_ERRORCODE_GAME_READY_REPEAT             DDErrorCode = -110
)

var DDErrorCode_name = map[int32]string{
	0:    "ERRORCODE_SUCC",
	-101: "ERRORCODE_CREATE_DESK_DIAMOND_NOTENOUGH",
	-102: "ERRORCODE_CREATE_DESK_USER_NOTFOUND",
	-110: "ERRORCODE_GAME_READY_REPEAT",
}
var DDErrorCode_value = map[string]int32{
	"ERRORCODE_SUCC":                          0,
	"ERRORCODE_CREATE_DESK_DIAMOND_NOTENOUGH": -101,
	"ERRORCODE_CREATE_DESK_USER_NOTFOUND":     -102,
	"ERRORCODE_GAME_READY_REPEAT":             -110,
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
func (DDErrorCode) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

// general response protocol
type ProtoHeader struct {
	ApiVer           *string `protobuf:"bytes,1,opt,name=apiVer" json:"apiVer,omitempty"`
	SessionId        *string `protobuf:"bytes,2,opt,name=sessionId" json:"sessionId,omitempty"`
	UserId           *uint32 `protobuf:"varint,3,opt,name=userId" json:"userId,omitempty"`
	PacketId         *int32  `protobuf:"varint,4,opt,name=packetId" json:"packetId,omitempty"`
	Code             *int32  `protobuf:"varint,5,opt,name=code" json:"code,omitempty"`
	Error            *string `protobuf:"bytes,6,opt,name=error" json:"error,omitempty"`
	ExtraTag         *int32  `protobuf:"varint,7,opt,name=extraTag" json:"extraTag,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ProtoHeader) Reset()                    { *m = ProtoHeader{} }
func (m *ProtoHeader) String() string            { return proto.CompactTextString(m) }
func (*ProtoHeader) ProtoMessage()               {}
func (*ProtoHeader) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *ProtoHeader) GetApiVer() string {
	if m != nil && m.ApiVer != nil {
		return *m.ApiVer
	}
	return ""
}

func (m *ProtoHeader) GetSessionId() string {
	if m != nil && m.SessionId != nil {
		return *m.SessionId
	}
	return ""
}

func (m *ProtoHeader) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *ProtoHeader) GetPacketId() int32 {
	if m != nil && m.PacketId != nil {
		return *m.PacketId
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

func (m *ProtoHeader) GetExtraTag() int32 {
	if m != nil && m.ExtraTag != nil {
		return *m.ExtraTag
	}
	return 0
}

type TerminalInfo struct {
	Channel          *string `protobuf:"bytes,1,opt,name=channel" json:"channel,omitempty"`
	DeviceName       *string `protobuf:"bytes,2,opt,name=deviceName" json:"deviceName,omitempty"`
	Uuid             *string `protobuf:"bytes,3,opt,name=uuid" json:"uuid,omitempty"`
	Os               *string `protobuf:"bytes,4,opt,name=os" json:"os,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TerminalInfo) Reset()                    { *m = TerminalInfo{} }
func (m *TerminalInfo) String() string            { return proto.CompactTextString(m) }
func (*TerminalInfo) ProtoMessage()               {}
func (*TerminalInfo) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *TerminalInfo) GetChannel() string {
	if m != nil && m.Channel != nil {
		return *m.Channel
	}
	return ""
}

func (m *TerminalInfo) GetDeviceName() string {
	if m != nil && m.DeviceName != nil {
		return *m.DeviceName
	}
	return ""
}

func (m *TerminalInfo) GetUuid() string {
	if m != nil && m.Uuid != nil {
		return *m.Uuid
	}
	return ""
}

func (m *TerminalInfo) GetOs() string {
	if m != nil && m.Os != nil {
		return *m.Os
	}
	return ""
}

type User struct {
	Mid                *string `protobuf:"bytes,1,opt,name=mid" json:"mid,omitempty"`
	Pwd                *string `protobuf:"bytes,2,opt,name=pwd" json:"pwd,omitempty"`
	Coin               *int64  `protobuf:"varint,3,opt,name=coin" json:"coin,omitempty"`
	Id                 *uint32 `protobuf:"varint,4,opt,name=id" json:"id,omitempty"`
	NickName           *string `protobuf:"bytes,5,opt,name=nickName" json:"nickName,omitempty"`
	LoginTurntable     *bool   `protobuf:"varint,6,opt,name=loginTurntable" json:"loginTurntable,omitempty"`
	LoginTurntableTime *string `protobuf:"bytes,7,opt,name=loginTurntableTime" json:"loginTurntableTime,omitempty"`
	Scores             *int32  `protobuf:"varint,8,opt,name=scores" json:"scores,omitempty"`
	LastSignTime       *string `protobuf:"bytes,9,opt,name=lastSignTime" json:"lastSignTime,omitempty"`
	SignCount          *int32  `protobuf:"varint,10,opt,name=signCount" json:"signCount,omitempty"`
	Diamond            *int64  `protobuf:"varint,11,opt,name=Diamond" json:"Diamond,omitempty"`
	OpenId             *string `protobuf:"bytes,12,opt,name=openId" json:"openId,omitempty"`
	HeadUrl            *string `protobuf:"bytes,13,opt,name=headUrl" json:"headUrl,omitempty"`
	XXX_unrecognized   []byte  `json:"-"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *User) GetMid() string {
	if m != nil && m.Mid != nil {
		return *m.Mid
	}
	return ""
}

func (m *User) GetPwd() string {
	if m != nil && m.Pwd != nil {
		return *m.Pwd
	}
	return ""
}

func (m *User) GetCoin() int64 {
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

func (m *User) GetDiamond() int64 {
	if m != nil && m.Diamond != nil {
		return *m.Diamond
	}
	return 0
}

func (m *User) GetOpenId() string {
	if m != nil && m.OpenId != nil {
		return *m.OpenId
	}
	return ""
}

func (m *User) GetHeadUrl() string {
	if m != nil && m.HeadUrl != nil {
		return *m.HeadUrl
	}
	return ""
}

func init() {
	proto.RegisterType((*ProtoHeader)(nil), "bbproto.ProtoHeader")
	proto.RegisterType((*TerminalInfo)(nil), "bbproto.TerminalInfo")
	proto.RegisterType((*User)(nil), "bbproto.User")
	proto.RegisterEnum("bbproto.EProtoId", EProtoId_name, EProtoId_value)
	proto.RegisterEnum("bbproto.DDErrorCode", DDErrorCode_name, DDErrorCode_value)
}

var fileDescriptor2 = []byte{
	// 1117 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x56, 0xdb, 0x56, 0xdb, 0x46,
	0x17, 0xfe, 0x1d, 0x30, 0xc6, 0x83, 0x21, 0x93, 0xc9, 0x01, 0x71, 0x4a, 0x08, 0xc9, 0xdf, 0x50,
	0xda, 0xa6, 0x69, 0x9b, 0x9e, 0x8f, 0xb2, 0x34, 0xd8, 0x0a, 0x92, 0xc6, 0x48, 0x23, 0x58, 0xe9,
	0x0d, 0x4b, 0xd8, 0x2a, 0xd1, 0x8a, 0x91, 0x58, 0xb2, 0x69, 0x7b, 0xdb, 0x57, 0xe8, 0x65, 0xbb,
	0xfa, 0x14, 0x7d, 0xac, 0x3e, 0x44, 0xbb, 0xf7, 0xc8, 0x96, 0x11, 0x82, 0x0b, 0x31, 0xf3, 0x7d,
	0x7b, 0xcf, 0xfe, 0xf6, 0x61, 0x64, 0x11, 0x72, 0x1a, 0x8e, 0xa2, 0xe7, 0x17, 0x59, 0x3a, 0x4e,
	0x59, 0xe3, 0xf4, 0x54, 0x2d, 0x76, 0x7e, 0xab, 0x91, 0xa5, 0x1e, 0xae, 0xba, 0x51, 0x38, 0x88,
	0x32, 0xb6, 0x42, 0x16, 0xc2, 0x8b, 0xf8, 0x28, 0xca, 0xb4, 0xda, 0x76, 0x6d, 0xb7, 0xc9, 0xee,
	0x90, 0xe6, 0x28, 0x1a, 0x8d, 0xe2, 0x34, 0xb1, 0x06, 0xda, 0x2d, 0x05, 0x81, 0xc9, 0xe5, 0x28,
	0xca, 0x60, 0x3f, 0x07, 0xfb, 0x65, 0x46, 0xc9, 0xe2, 0x45, 0xd8, 0x7f, 0x1b, 0x8d, 0x01, 0x99,
	0x07, 0xa4, 0xce, 0x5a, 0x64, 0xbe, 0x9f, 0x0e, 0x22, 0xad, 0xae, 0x76, 0xcb, 0xa4, 0x1e, 0x65,
	0x59, 0x9a, 0x69, 0x0b, 0xca, 0x1d, 0xcc, 0xa3, 0x5f, 0xc7, 0x59, 0x28, 0xc3, 0x33, 0xad, 0x81,
	0x06, 0x3b, 0x0e, 0x69, 0xc9, 0x28, 0x3b, 0x8f, 0x93, 0x70, 0x68, 0x25, 0x3f, 0xa5, 0xec, 0x36,
	0x69, 0xf4, 0xdf, 0x84, 0x49, 0x12, 0x0d, 0x27, 0x22, 0x18, 0x21, 0x83, 0xe8, 0xe7, 0xb8, 0x1f,
	0xb9, 0xe1, 0x79, 0x34, 0x51, 0x01, 0x31, 0x2e, 0x2f, 0xe3, 0x5c, 0x43, 0x13, 0x0c, 0x6e, 0xa5,
	0x23, 0x15, 0xbd, 0xb9, 0xf3, 0x4f, 0x8d, 0xcc, 0x07, 0x20, 0x90, 0x2d, 0x91, 0xb9, 0x73, 0xb0,
	0xc8, 0xcf, 0x80, 0xcd, 0xc5, 0x2f, 0x83, 0x99, 0x73, 0x3f, 0x8d, 0x13, 0xe5, 0x3c, 0x87, 0xce,
	0x71, 0x2e, 0x5d, 0x25, 0x93, 0xc4, 0xfd, 0xb7, 0x2a, 0x50, 0x5d, 0xd9, 0x3e, 0x20, 0x2b, 0xc3,
	0xf4, 0x2c, 0x4e, 0xe4, 0x65, 0x96, 0x8c, 0xc3, 0xd3, 0x61, 0xa4, 0xf2, 0x58, 0x64, 0xeb, 0x84,
	0x95, 0x71, 0x19, 0x83, 0x4f, 0x63, 0x5a, 0xa2, 0x51, 0x3f, 0xcd, 0xa2, 0x91, 0xb6, 0xa8, 0x4a,
	0x70, 0x8f, 0xb4, 0x86, 0xe1, 0x68, 0xec, 0xc7, 0x67, 0x89, 0xb2, 0x6a, 0x16, 0xb5, 0x05, 0xc4,
	0x48, 0x2f, 0x93, 0xb1, 0x46, 0x94, 0x21, 0xa4, 0x6e, 0xc6, 0xe1, 0x79, 0x9a, 0x0c, 0xb4, 0x25,
	0xa5, 0x0d, 0x4e, 0x4a, 0x2f, 0x22, 0x2c, 0x7e, 0x4b, 0xf9, 0x80, 0xc1, 0x1b, 0xe8, 0x54, 0x90,
	0x0d, 0xb5, 0x65, 0x04, 0xf6, 0xfe, 0x5a, 0x26, 0x8b, 0x5c, 0x75, 0xd0, 0x1a, 0x00, 0xbb, 0xd4,
	0xe1, 0xd2, 0x72, 0xa5, 0xf0, 0x84, 0x70, 0x68, 0x0d, 0x6a, 0xdf, 0xf4, 0xbb, 0x81, 0xd5, 0x09,
	0xc4, 0x2b, 0x8b, 0xde, 0x62, 0x77, 0xc9, 0xed, 0x62, 0xdb, 0xb5, 0x6c, 0xe1, 0xf4, 0xe8, 0x1c,
	0xa4, 0xdc, 0x2a, 0x40, 0x8f, 0xfb, 0x74, 0x1e, 0x6a, 0xd5, 0xf8, 0xf1, 0x55, 0x57, 0x1d, 0x51,
	0x87, 0xea, 0x2c, 0xc0, 0xa6, 0xcd, 0x25, 0x5d, 0x98, 0xac, 0x1d, 0xbf, 0x43, 0x1b, 0x50, 0x17,
	0x06, 0xeb, 0xc3, 0x80, 0x7b, 0xaf, 0x5d, 0xe1, 0x73, 0x5d, 0x06, 0x3e, 0xf7, 0xe8, 0x22, 0x28,
	0x26, 0xe8, 0xcc, 0x0f, 0x11, 0xa3, 0xcd, 0xc9, 0xde, 0x16, 0x52, 0x82, 0x25, 0x25, 0x4c, 0x23,
	0xf7, 0xf0, 0x3c, 0x4f, 0xe8, 0xa6, 0xa1, 0xfb, 0xb2, 0xcd, 0x3b, 0x96, 0x8b, 0xa7, 0x2f, 0x41,
	0x95, 0xa8, 0x2d, 0x60, 0xe7, 0x5b, 0x1d, 0x17, 0x30, 0xe1, 0x06, 0x3e, 0x6d, 0xb1, 0x55, 0x72,
	0x57, 0xa1, 0x32, 0xf0, 0x5c, 0xa9, 0xb7, 0x6d, 0x9e, 0x13, 0xcb, 0x78, 0xb0, 0xb0, 0x2d, 0x77,
	0xb2, 0x5f, 0xc1, 0xe4, 0xa5, 0xe5, 0x58, 0x6e, 0x27, 0x07, 0x6e, 0xa3, 0x5a, 0x99, 0x67, 0x41,
	0x59, 0x93, 0xd4, 0xa5, 0x4a, 0xe2, 0x0e, 0xfa, 0xa9, 0xa5, 0x8a, 0x4c, 0x19, 0x4c, 0xd7, 0x4a,
	0xbe, 0x9f, 0x4a, 0xa2, 0x77, 0xa1, 0xb9, 0x0f, 0x72, 0x57, 0xdd, 0x34, 0x31, 0xaf, 0x19, 0x77,
	0x0f, 0x65, 0xf6, 0x2c, 0xf3, 0x04, 0x32, 0x3c, 0x0c, 0x2c, 0xe3, 0xc0, 0x10, 0xae, 0x4b, 0xef,
	0x4f, 0x51, 0xdd, 0x38, 0x98, 0xa1, 0x0f, 0x50, 0x13, 0xa2, 0x6e, 0x60, 0xdb, 0x58, 0xb5, 0x55,
	0xf6, 0x94, 0x6c, 0x23, 0xe0, 0xe8, 0xd2, 0xe8, 0xda, 0x96, 0x2f, 0xf1, 0x18, 0x47, 0xb4, 0x2d,
	0x9b, 0x17, 0x10, 0xd5, 0xb0, 0xb6, 0x68, 0xd5, 0xd1, 0x1d, 0x7e, 0xa2, 0x92, 0xc7, 0x15, 0x5d,
	0xc3, 0x5a, 0x14, 0x38, 0x77, 0xa1, 0xa0, 0xca, 0x89, 0xae, 0xa3, 0xde, 0x82, 0x00, 0x09, 0x57,
	0xb8, 0x0d, 0xb6, 0x46, 0xee, 0x17, 0x9c, 0xcf, 0x5d, 0x13, 0x17, 0x96, 0xbb, 0x2f, 0xe8, 0x66,
	0x29, 0x4e, 0x1b, 0x6a, 0x69, 0x1a, 0x02, 0x4a, 0xb2, 0xc5, 0xee, 0x93, 0x3b, 0x05, 0x6e, 0xb9,
	0x96, 0x34, 0x74, 0xcf, 0xa4, 0x0f, 0x2b, 0x27, 0xed, 0xdb, 0xa2, 0xa7, 0xa8, 0x47, 0x15, 0x0a,
	0xbb, 0xa5, 0xa8, 0xed, 0x92, 0x36, 0xa4, 0x3c, 0xeb, 0x88, 0x7b, 0x8a, 0x7b, 0x5c, 0x0a, 0xe4,
	0xe9, 0x96, 0xcf, 0xb1, 0x45, 0x3b, 0x38, 0x23, 0x57, 0xd3, 0x29, 0x98, 0x27, 0x25, 0xc5, 0xfb,
	0xc2, 0xb6, 0xc5, 0x31, 0xe2, 0x4f, 0x4b, 0xf1, 0xc1, 0x63, 0x46, 0xfd, 0x7f, 0xda, 0x99, 0xa9,
	0x8b, 0x89, 0xe8, 0x3b, 0xa5, 0x52, 0xe6, 0x0e, 0x8a, 0x78, 0x56, 0x92, 0x64, 0x74, 0xb9, 0x71,
	0x80, 0xf0, 0xee, 0x75, 0x49, 0x05, 0xf3, 0x6e, 0x25, 0x75, 0x01, 0xe9, 0x61, 0xfa, 0x74, 0xaf,
	0xe4, 0x84, 0xd4, 0x64, 0x9a, 0xe8, 0x7b, 0xa5, 0xa2, 0xe4, 0x74, 0x57, 0x1c, 0xab, 0xa2, 0xbc,
	0xcf, 0xb6, 0xc8, 0x5a, 0x99, 0x83, 0x78, 0x05, 0xfd, 0x01, 0xdb, 0x24, 0x5a, 0x99, 0x96, 0xdc,
	0x97, 0x70, 0x71, 0x03, 0x5b, 0xd2, 0xe7, 0xa5, 0x6c, 0x7b, 0x1e, 0x57, 0x0d, 0xfd, 0xb0, 0x94,
	0x2d, 0x3e, 0x5c, 0x21, 0x2d, 0x83, 0xd3, 0x17, 0x6c, 0x83, 0xac, 0x56, 0x62, 0x4d, 0xc8, 0x8f,
	0xaa, 0x91, 0x0c, 0x0f, 0xee, 0x34, 0x37, 0xb9, 0x7f, 0x40, 0x3f, 0x66, 0x8f, 0xc8, 0x46, 0xc5,
	0xf5, 0x8a, 0xc1, 0x27, 0x95, 0xa0, 0x30, 0xeb, 0xba, 0xf9, 0x9a, 0xbe, 0xac, 0x26, 0x8f, 0x3d,
	0x56, 0xdc, 0xa7, 0x55, 0xa7, 0xfc, 0x9a, 0x7e, 0x56, 0x15, 0x83, 0x0f, 0x4c, 0x10, 0x8a, 0xf2,
	0xf9, 0x8d, 0x62, 0xae, 0x18, 0x7c, 0xc1, 0xb6, 0xc9, 0xe6, 0xf5, 0x73, 0x75, 0xf7, 0x8a, 0xc5,
	0x97, 0xec, 0x21, 0x59, 0x2f, 0x5b, 0x98, 0x96, 0xef, 0x0b, 0xfb, 0x28, 0x4f, 0xe7, 0x2b, 0xf6,
	0x98, 0x6c, 0x55, 0x42, 0x94, 0x4c, 0xbe, 0x2e, 0xdf, 0x5b, 0xae, 0x4f, 0xf0, 0x6f, 0xae, 0x4f,
	0xe7, 0x8c, 0xfa, 0xb6, 0x24, 0x1c, 0x47, 0x04, 0x61, 0xf8, 0x37, 0x7d, 0x5f, 0x7e, 0x57, 0x6a,
	0xa8, 0xc3, 0x7d, 0x5f, 0xef, 0x70, 0xfa, 0x7d, 0x65, 0xb2, 0xa6, 0xcc, 0x0f, 0xa5, 0x8e, 0x4a,
	0x11, 0xb8, 0xd8, 0x6d, 0xa9, 0x2e, 0x37, 0xd5, 0x4b, 0xa3, 0x55, 0x90, 0x1e, 0x3f, 0x86, 0xc9,
	0xf2, 0x69, 0xbb, 0xd4, 0x98, 0x19, 0xad, 0xbb, 0x07, 0xd4, 0xb8, 0xd9, 0xd5, 0x0f, 0x1c, 0x47,
	0x07, 0x99, 0x66, 0x29, 0xf5, 0xd9, 0xab, 0x8c, 0xef, 0xfd, 0x0d, 0x1f, 0x18, 0x90, 0x14, 0x7e,
	0x00, 0x18, 0xf0, 0x4d, 0x80, 0x6f, 0x5b, 0xee, 0x79, 0xc2, 0x33, 0x84, 0x09, 0xca, 0x03, 0xc3,
	0xa0, 0xff, 0x63, 0x2f, 0xc9, 0xb3, 0x19, 0x96, 0x8f, 0xd0, 0x09, 0xd6, 0x01, 0x3a, 0xa0, 0x3b,
	0xc2, 0x85, 0xf7, 0xa7, 0x90, 0x30, 0xb2, 0x41, 0xa7, 0x4b, 0xff, 0xfc, 0x77, 0xf2, 0x57, 0x63,
	0x2f, 0xc8, 0x93, 0x9b, 0xbd, 0xf0, 0x8a, 0xa1, 0xcb, 0x3e, 0xa8, 0x34, 0xe9, 0x1f, 0x33, 0x8f,
	0x5d, 0xb2, 0x31, 0xf3, 0x98, 0x4d, 0x24, 0x3c, 0x7b, 0xf8, 0x5b, 0xf5, 0x7b, 0x61, 0xf9, 0x5f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x78, 0xee, 0xee, 0x66, 0x2c, 0x09, 0x00, 0x00,
}
