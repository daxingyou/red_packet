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

type EUnitType int32

const (
	EUnitType_UALL   EUnitType = 0
	EUnitType_UFIRE  EUnitType = 1
	EUnitType_UWATER EUnitType = 2
	EUnitType_UWIND  EUnitType = 3
	EUnitType_ULIGHT EUnitType = 4
	EUnitType_UDARK  EUnitType = 5
	EUnitType_UNONE  EUnitType = 6
	EUnitType_UHeart EUnitType = 7
)

var EUnitType_name = map[int32]string{
	0: "UALL",
	1: "UFIRE",
	2: "UWATER",
	3: "UWIND",
	4: "ULIGHT",
	5: "UDARK",
	6: "UNONE",
	7: "UHeart",
}
var EUnitType_value = map[string]int32{
	"UALL":   0,
	"UFIRE":  1,
	"UWATER": 2,
	"UWIND":  3,
	"ULIGHT": 4,
	"UDARK":  5,
	"UNONE":  6,
	"UHeart": 7,
}

func (x EUnitType) Enum() *EUnitType {
	p := new(EUnitType)
	*p = x
	return p
}
func (x EUnitType) String() string {
	return proto.EnumName(EUnitType_name, int32(x))
}
func (x *EUnitType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EUnitType_value, data, "EUnitType")
	if err != nil {
		return err
	}
	*x = EUnitType(value)
	return nil
}
func (EUnitType) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type EUnitRace int32

const (
	EUnitRace_ALL          EUnitRace = 0
	EUnitRace_HUMAN        EUnitRace = 1
	EUnitRace_UNDEAD       EUnitRace = 2
	EUnitRace_MYTHIC       EUnitRace = 3
	EUnitRace_BEAST        EUnitRace = 4
	EUnitRace_MONSTER      EUnitRace = 5
	EUnitRace_LEGEND       EUnitRace = 6
	EUnitRace_DRAGON       EUnitRace = 7
	EUnitRace_SCREAMCHEESE EUnitRace = 8
	EUnitRace_EVOLVEPARTS  EUnitRace = 9
)

var EUnitRace_name = map[int32]string{
	0: "ALL",
	1: "HUMAN",
	2: "UNDEAD",
	3: "MYTHIC",
	4: "BEAST",
	5: "MONSTER",
	6: "LEGEND",
	7: "DRAGON",
	8: "SCREAMCHEESE",
	9: "EVOLVEPARTS",
}
var EUnitRace_value = map[string]int32{
	"ALL":          0,
	"HUMAN":        1,
	"UNDEAD":       2,
	"MYTHIC":       3,
	"BEAST":        4,
	"MONSTER":      5,
	"LEGEND":       6,
	"DRAGON":       7,
	"SCREAMCHEESE": 8,
	"EVOLVEPARTS":  9,
}

func (x EUnitRace) Enum() *EUnitRace {
	p := new(EUnitRace)
	*p = x
	return p
}
func (x EUnitRace) String() string {
	return proto.EnumName(EUnitRace_name, int32(x))
}
func (x *EUnitRace) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EUnitRace_value, data, "EUnitRace")
	if err != nil {
		return err
	}
	*x = EUnitRace(value)
	return nil
}
func (EUnitRace) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type EUnitProtoId int32

const (
	EUnitProtoId_TESTP1          EUnitProtoId = 0
	EUnitProtoId_REG             EUnitProtoId = 1
	EUnitProtoId_REQAUTHUSER     EUnitProtoId = 2
	EUnitProtoId_HEATBEAT        EUnitProtoId = 3
	EUnitProtoId_GETINTOROOM     EUnitProtoId = 4
	EUnitProtoId_ROOMMSG         EUnitProtoId = 5
	EUnitProtoId_GETREWARDS      EUnitProtoId = 6
	EUnitProtoId_SHUIGUOJI       EUnitProtoId = 7
	EUnitProtoId_SHUIGUOJIHILOMP EUnitProtoId = 8
	EUnitProtoId_SHUIGUOJIRES    EUnitProtoId = 9
	// 扎金花相关的逻辑
	EUnitProtoId_ZJHROOM            EUnitProtoId = 10
	EUnitProtoId_ZJHBET             EUnitProtoId = 11
	EUnitProtoId_ZJHMSG             EUnitProtoId = 12
	EUnitProtoId_ZJHQUERYNOSEATUSER EUnitProtoId = 13
	EUnitProtoId_ZJHREQSEAT         EUnitProtoId = 14
	EUnitProtoId_ZJHLOTTERY         EUnitProtoId = 15
)

var EUnitProtoId_name = map[int32]string{
	0:  "TESTP1",
	1:  "REG",
	2:  "REQAUTHUSER",
	3:  "HEATBEAT",
	4:  "GETINTOROOM",
	5:  "ROOMMSG",
	6:  "GETREWARDS",
	7:  "SHUIGUOJI",
	8:  "SHUIGUOJIHILOMP",
	9:  "SHUIGUOJIRES",
	10: "ZJHROOM",
	11: "ZJHBET",
	12: "ZJHMSG",
	13: "ZJHQUERYNOSEATUSER",
	14: "ZJHREQSEAT",
	15: "ZJHLOTTERY",
}
var EUnitProtoId_value = map[string]int32{
	"TESTP1":             0,
	"REG":                1,
	"REQAUTHUSER":        2,
	"HEATBEAT":           3,
	"GETINTOROOM":        4,
	"ROOMMSG":            5,
	"GETREWARDS":         6,
	"SHUIGUOJI":          7,
	"SHUIGUOJIHILOMP":    8,
	"SHUIGUOJIRES":       9,
	"ZJHROOM":            10,
	"ZJHBET":             11,
	"ZJHMSG":             12,
	"ZJHQUERYNOSEATUSER": 13,
	"ZJHREQSEAT":         14,
	"ZJHLOTTERY":         15,
}

func (x EUnitProtoId) Enum() *EUnitProtoId {
	p := new(EUnitProtoId)
	*p = x
	return p
}
func (x EUnitProtoId) String() string {
	return proto.EnumName(EUnitProtoId_name, int32(x))
}
func (x *EUnitProtoId) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EUnitProtoId_value, data, "EUnitProtoId")
	if err != nil {
		return err
	}
	*x = EUnitProtoId(value)
	return nil
}
func (EUnitProtoId) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

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
func (*ProtoHeader) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

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
func (*TerminalInfo) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

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

func init() {
	proto.RegisterType((*ProtoHeader)(nil), "bbproto.ProtoHeader")
	proto.RegisterType((*TerminalInfo)(nil), "bbproto.TerminalInfo")
	proto.RegisterEnum("bbproto.EUnitType", EUnitType_name, EUnitType_value)
	proto.RegisterEnum("bbproto.EUnitRace", EUnitRace_name, EUnitRace_value)
	proto.RegisterEnum("bbproto.EUnitProtoId", EUnitProtoId_name, EUnitProtoId_value)
}

var fileDescriptor1 = []byte{
	// 518 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x3c, 0x91, 0xcf, 0x6f, 0xd3, 0x4e,
	0x10, 0xc5, 0xbf, 0x49, 0xe3, 0x38, 0x9e, 0x38, 0xed, 0x7e, 0x17, 0x09, 0xe5, 0x88, 0x7a, 0x42,
	0x3d, 0x20, 0xf1, 0x27, 0x6c, 0xeb, 0xc1, 0x76, 0xf1, 0x8f, 0x74, 0xbd, 0x4e, 0x14, 0x6e, 0x4e,
	0xb2, 0x80, 0x45, 0x6b, 0x47, 0x76, 0x82, 0xe0, 0x88, 0xb8, 0xf1, 0x57, 0x33, 0xbb, 0x35, 0xbd,
	0xbd, 0x79, 0x7e, 0xb3, 0xef, 0x93, 0x09, 0xc0, 0xae, 0xea, 0xf5, 0xbb, 0x63, 0xd7, 0x9e, 0x5a,
	0xee, 0xee, 0x76, 0x56, 0x5c, 0xff, 0x1a, 0xc1, 0x7c, 0x65, 0x54, 0xa4, 0xab, 0x83, 0xee, 0xf8,
	0x25, 0x4c, 0xab, 0x63, 0xbd, 0xd6, 0xdd, 0x72, 0xf4, 0x66, 0xf4, 0xd6, 0xe3, 0xff, 0x83, 0xd7,
	0xeb, 0xbe, 0xaf, 0xdb, 0x26, 0x3e, 0x2c, 0xc7, 0xd6, 0xa2, 0xc8, 0xb9, 0xd7, 0x1d, 0xcd, 0x17,
	0x34, 0x2f, 0x38, 0x83, 0xd9, 0xb1, 0xda, 0x7f, 0xd3, 0x27, 0x72, 0x26, 0xe4, 0x38, 0xdc, 0x87,
	0xc9, 0xbe, 0x3d, 0xe8, 0xa5, 0x63, 0xa7, 0x05, 0x38, 0xba, 0xeb, 0xda, 0x6e, 0x39, 0xb5, 0xeb,
	0x14, 0xd7, 0x3f, 0x4e, 0x5d, 0xa5, 0xaa, 0x2f, 0x4b, 0xd7, 0x04, 0xae, 0x53, 0xf0, 0x95, 0xee,
	0x9e, 0xea, 0xa6, 0x7a, 0x8c, 0x9b, 0xcf, 0x2d, 0xbf, 0x02, 0x77, 0xff, 0xb5, 0x6a, 0x1a, 0xfd,
	0x38, 0x40, 0x70, 0x80, 0x83, 0xfe, 0x5e, 0xef, 0x75, 0x56, 0x3d, 0xe9, 0x81, 0x82, 0x3a, 0xce,
	0xe7, 0xfa, 0x99, 0xc1, 0xa3, 0xc0, 0xb8, 0xed, 0x6d, 0xbb, 0x77, 0xa3, 0xc1, 0xc3, 0xb2, 0xa9,
	0x4f, 0xea, 0xe7, 0x51, 0xf3, 0x19, 0x4c, 0x4a, 0x91, 0x24, 0xec, 0x3f, 0xee, 0x81, 0x53, 0x7e,
	0x88, 0x25, 0xb2, 0x11, 0xa5, 0xa7, 0xe5, 0x46, 0x28, 0x94, 0x6c, 0x6c, 0xed, 0x4d, 0x9c, 0x05,
	0xec, 0xc2, 0xda, 0x49, 0x1c, 0x46, 0x8a, 0x4d, 0xac, 0x1d, 0x08, 0xf9, 0x91, 0x39, 0x56, 0x66,
	0x79, 0x86, 0x6c, 0x6a, 0x13, 0x74, 0xa8, 0xee, 0xc4, 0xdc, 0x9b, 0x3f, 0xa3, 0xa1, 0x47, 0x56,
	0x7b, 0xcd, 0x5d, 0xb8, 0x78, 0xa9, 0x89, 0xca, 0x54, 0x64, 0x43, 0x4d, 0x16, 0xa0, 0x08, 0xa8,
	0x86, 0x74, 0xba, 0x55, 0x51, 0x7c, 0x47, 0x3d, 0x14, 0xb9, 0x45, 0x51, 0x98, 0x9a, 0x39, 0xb8,
	0x69, 0x9e, 0x15, 0x06, 0xc5, 0x31, 0x99, 0x04, 0x43, 0x24, 0x16, 0xdb, 0x14, 0x48, 0x11, 0xe6,
	0x19, 0x73, 0xe9, 0x62, 0x7e, 0x71, 0x27, 0x51, 0xa4, 0x77, 0x11, 0x62, 0x81, 0x6c, 0x46, 0x17,
	0x9a, 0xe3, 0x3a, 0x4f, 0xd6, 0xb8, 0x12, 0x52, 0x15, 0xcc, 0xbb, 0xf9, 0x3d, 0x06, 0xdf, 0xc2,
	0xd8, 0xff, 0x32, 0x3e, 0x98, 0x7d, 0x85, 0x85, 0x5a, 0xbd, 0x27, 0x24, 0x62, 0x93, 0x18, 0x12,
	0x10, 0xad, 0x49, 0x7c, 0x10, 0xa5, 0x8a, 0xca, 0xc2, 0xfe, 0x78, 0x1f, 0x66, 0x11, 0x0a, 0x45,
	0x34, 0x8a, 0xb8, 0xe8, 0x73, 0x88, 0x2a, 0xce, 0x54, 0x2e, 0xf3, 0x3c, 0x7d, 0xa6, 0x33, 0x2a,
	0x2d, 0x42, 0xa2, 0xbb, 0x04, 0xa0, 0xaf, 0x12, 0x37, 0x42, 0x06, 0x05, 0x11, 0x2e, 0xc0, 0x2b,
	0xa2, 0x32, 0x0e, 0xcb, 0xfc, 0x3e, 0x26, 0xc8, 0x57, 0x70, 0xf5, 0x32, 0x46, 0x71, 0x92, 0xa7,
	0x2b, 0xe2, 0x34, 0xe4, 0xff, 0x4c, 0x89, 0x04, 0x6a, 0x9e, 0xfc, 0x74, 0x1f, 0xd9, 0xf7, 0xc1,
	0x40, 0xd2, 0x70, 0x8b, 0x8a, 0xcd, 0x07, 0x6d, 0xaa, 0x7c, 0xfe, 0x1a, 0x38, 0xe9, 0x87, 0x12,
	0xe5, 0x36, 0xcb, 0x0b, 0x82, 0xb3, 0xb8, 0x0b, 0x83, 0x60, 0x96, 0xf1, 0xc1, 0x78, 0xec, 0x72,
	0x98, 0x93, 0x5c, 0xd1, 0xfd, 0xb6, 0xec, 0xea, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcb, 0x5f,
	0x39, 0xce, 0xe2, 0x02, 0x00, 0x00,
}
