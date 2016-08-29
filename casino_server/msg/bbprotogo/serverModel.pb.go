// Code generated by protoc-gen-go.
// source: serverModel.proto
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

// Ignoring public import of User from base.proto

// Ignoring public import of EProtoId from base.proto

// Ignoring public import of DDErrorCode from base.proto

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

// Ignoring public import of ThRoom from thPoker.proto

// Ignoring public import of THRoomAddUserBroadcast from thPoker.proto

// Ignoring public import of THPoker from thPoker.proto

// Ignoring public import of THBet from thPoker.proto

// Ignoring public import of THBetBegin from thPoker.proto

// Ignoring public import of THBetBroadcast from thPoker.proto

// Ignoring public import of THUser from thPoker.proto

// Ignoring public import of THLottery from thPoker.proto

// Ignoring public import of ETHType from thPoker.proto

// Ignoring public import of ETHAction from thPoker.proto

// 押注的记录,针对每轮扎金花,只有一次押注记录
type TBetRecord struct {
	UserId           *uint32 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	Betzone          []int32 `protobuf:"varint,2,rep,name=betzone" json:"betzone,omitempty"`
	ZjhRoundNumber   *string `protobuf:"bytes,3,opt,name=ZjhRoundNumber" json:"ZjhRoundNumber,omitempty"`
	WinAmount        *int32  `protobuf:"varint,4,opt,name=winAmount" json:"winAmount,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TBetRecord) Reset()                    { *m = TBetRecord{} }
func (m *TBetRecord) String() string            { return proto.CompactTextString(m) }
func (*TBetRecord) ProtoMessage()               {}
func (*TBetRecord) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *TBetRecord) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *TBetRecord) GetBetzone() []int32 {
	if m != nil {
		return m.Betzone
	}
	return nil
}

func (m *TBetRecord) GetZjhRoundNumber() string {
	if m != nil && m.ZjhRoundNumber != nil {
		return *m.ZjhRoundNumber
	}
	return ""
}

func (m *TBetRecord) GetWinAmount() int32 {
	if m != nil && m.WinAmount != nil {
		return *m.WinAmount
	}
	return 0
}

// 每一轮扎金花的数据
type TZjhRound struct {
	Id               *uint32   `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	BeginTime        *int64    `protobuf:"varint,2,opt,name=beginTime" json:"beginTime,omitempty"`
	BetEndTime       *int64    `protobuf:"varint,3,opt,name=betEndTime" json:"betEndTime,omitempty"`
	LotteryTime      *int64    `protobuf:"varint,4,opt,name=lotteryTime" json:"lotteryTime,omitempty"`
	EndTime          *int64    `protobuf:"varint,5,opt,name=endTime" json:"endTime,omitempty"`
	ZoneAmount       []int32   `protobuf:"varint,6,rep,name=zoneAmount" json:"zoneAmount,omitempty"`
	ZoneWinAmount    []int32   `protobuf:"varint,7,rep,name=zoneWinAmount" json:"zoneWinAmount,omitempty"`
	BankerUserId     *uint32   `protobuf:"varint,8,opt,name=BankerUserId" json:"BankerUserId,omitempty"`
	ZjhPaiList       []*ZjhPai `protobuf:"bytes,9,rep,name=ZjhPaiList" json:"ZjhPaiList,omitempty"`
	Number           *string   `protobuf:"bytes,10,opt,name=Number" json:"Number,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *TZjhRound) Reset()                    { *m = TZjhRound{} }
func (m *TZjhRound) String() string            { return proto.CompactTextString(m) }
func (*TZjhRound) ProtoMessage()               {}
func (*TZjhRound) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *TZjhRound) GetId() uint32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *TZjhRound) GetBeginTime() int64 {
	if m != nil && m.BeginTime != nil {
		return *m.BeginTime
	}
	return 0
}

func (m *TZjhRound) GetBetEndTime() int64 {
	if m != nil && m.BetEndTime != nil {
		return *m.BetEndTime
	}
	return 0
}

func (m *TZjhRound) GetLotteryTime() int64 {
	if m != nil && m.LotteryTime != nil {
		return *m.LotteryTime
	}
	return 0
}

func (m *TZjhRound) GetEndTime() int64 {
	if m != nil && m.EndTime != nil {
		return *m.EndTime
	}
	return 0
}

func (m *TZjhRound) GetZoneAmount() []int32 {
	if m != nil {
		return m.ZoneAmount
	}
	return nil
}

func (m *TZjhRound) GetZoneWinAmount() []int32 {
	if m != nil {
		return m.ZoneWinAmount
	}
	return nil
}

func (m *TZjhRound) GetBankerUserId() uint32 {
	if m != nil && m.BankerUserId != nil {
		return *m.BankerUserId
	}
	return 0
}

func (m *TZjhRound) GetZjhPaiList() []*ZjhPai {
	if m != nil {
		return m.ZjhPaiList
	}
	return nil
}

func (m *TZjhRound) GetNumber() string {
	if m != nil && m.Number != nil {
		return *m.Number
	}
	return ""
}

type TNotice struct {
	Id               *int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	NoticeType       *int32  `protobuf:"varint,2,opt,name=noticeType" json:"noticeType,omitempty"`
	NoticeTitle      *string `protobuf:"bytes,3,opt,name=noticeTitle" json:"noticeTitle,omitempty"`
	NoticeContent    *string `protobuf:"bytes,4,opt,name=noticeContent" json:"noticeContent,omitempty"`
	NoticeMemo       *string `protobuf:"bytes,5,opt,name=noticeMemo" json:"noticeMemo,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TNotice) Reset()                    { *m = TNotice{} }
func (m *TNotice) String() string            { return proto.CompactTextString(m) }
func (*TNotice) ProtoMessage()               {}
func (*TNotice) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *TNotice) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *TNotice) GetNoticeType() int32 {
	if m != nil && m.NoticeType != nil {
		return *m.NoticeType
	}
	return 0
}

func (m *TNotice) GetNoticeTitle() string {
	if m != nil && m.NoticeTitle != nil {
		return *m.NoticeTitle
	}
	return ""
}

func (m *TNotice) GetNoticeContent() string {
	if m != nil && m.NoticeContent != nil {
		return *m.NoticeContent
	}
	return ""
}

func (m *TNotice) GetNoticeMemo() string {
	if m != nil && m.NoticeMemo != nil {
		return *m.NoticeMemo
	}
	return ""
}

// 后台保存数据使用的mode,一般不传递给客户端,这个实体一般是保存在redis当中
// 当用程序崩溃的时候,可以用来恢复用户数据
type ThServerUser struct {
	Seat               *int32  `protobuf:"varint,1,opt,name=Seat" json:"Seat,omitempty"`
	Status             *int32  `protobuf:"varint,2,opt,name=Status" json:"Status,omitempty"`
	BreakStatus        *int32  `protobuf:"varint,3,opt,name=BreakStatus" json:"BreakStatus,omitempty"`
	HandCards          []*Pai  `protobuf:"bytes,4,rep,name=HandCards" json:"HandCards,omitempty"`
	WaiTime            *string `protobuf:"bytes,5,opt,name=waiTime" json:"waiTime,omitempty"`
	WaitUUID           *string `protobuf:"bytes,6,opt,name=waitUUID" json:"waitUUID,omitempty"`
	DeskId             *int32  `protobuf:"varint,7,opt,name=deskId" json:"deskId,omitempty"`
	TotalBet           *int64  `protobuf:"varint,8,opt,name=TotalBet" json:"TotalBet,omitempty"`
	TotalBet4CalcAllin *int64  `protobuf:"varint,9,opt,name=TotalBet4calcAllin" json:"TotalBet4calcAllin,omitempty"`
	WinAmount          *int64  `protobuf:"varint,10,opt,name=winAmount" json:"winAmount,omitempty"`
	TurnCoin           *int64  `protobuf:"varint,11,opt,name=TurnCoin" json:"TurnCoin,omitempty"`
	HandCoin           *int64  `protobuf:"varint,12,opt,name=HandCoin" json:"HandCoin,omitempty"`
	RoomCoin           *int64  `protobuf:"varint,13,opt,name=RoomCoin" json:"RoomCoin,omitempty"`
	WinAmountDetail    []int64 `protobuf:"varint,14,rep,name=winAmountDetail" json:"winAmountDetail,omitempty"`
	UserId             *uint32 `protobuf:"varint,15,opt,name=UserId" json:"UserId,omitempty"`
	GameNumber         *int32  `protobuf:"varint,16,opt,name=GameNumber" json:"GameNumber,omitempty"`
	IsBreak            *bool   `protobuf:"varint,17,opt,name=IsBreak" json:"IsBreak,omitempty"`
	IsLeave            *bool   `protobuf:"varint,18,opt,name=IsLeave" json:"IsLeave,omitempty"`
	XXX_unrecognized   []byte  `json:"-"`
}

func (m *ThServerUser) Reset()                    { *m = ThServerUser{} }
func (m *ThServerUser) String() string            { return proto.CompactTextString(m) }
func (*ThServerUser) ProtoMessage()               {}
func (*ThServerUser) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

func (m *ThServerUser) GetSeat() int32 {
	if m != nil && m.Seat != nil {
		return *m.Seat
	}
	return 0
}

func (m *ThServerUser) GetStatus() int32 {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return 0
}

func (m *ThServerUser) GetBreakStatus() int32 {
	if m != nil && m.BreakStatus != nil {
		return *m.BreakStatus
	}
	return 0
}

func (m *ThServerUser) GetHandCards() []*Pai {
	if m != nil {
		return m.HandCards
	}
	return nil
}

func (m *ThServerUser) GetWaiTime() string {
	if m != nil && m.WaiTime != nil {
		return *m.WaiTime
	}
	return ""
}

func (m *ThServerUser) GetWaitUUID() string {
	if m != nil && m.WaitUUID != nil {
		return *m.WaitUUID
	}
	return ""
}

func (m *ThServerUser) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

func (m *ThServerUser) GetTotalBet() int64 {
	if m != nil && m.TotalBet != nil {
		return *m.TotalBet
	}
	return 0
}

func (m *ThServerUser) GetTotalBet4CalcAllin() int64 {
	if m != nil && m.TotalBet4CalcAllin != nil {
		return *m.TotalBet4CalcAllin
	}
	return 0
}

func (m *ThServerUser) GetWinAmount() int64 {
	if m != nil && m.WinAmount != nil {
		return *m.WinAmount
	}
	return 0
}

func (m *ThServerUser) GetTurnCoin() int64 {
	if m != nil && m.TurnCoin != nil {
		return *m.TurnCoin
	}
	return 0
}

func (m *ThServerUser) GetHandCoin() int64 {
	if m != nil && m.HandCoin != nil {
		return *m.HandCoin
	}
	return 0
}

func (m *ThServerUser) GetRoomCoin() int64 {
	if m != nil && m.RoomCoin != nil {
		return *m.RoomCoin
	}
	return 0
}

func (m *ThServerUser) GetWinAmountDetail() []int64 {
	if m != nil {
		return m.WinAmountDetail
	}
	return nil
}

func (m *ThServerUser) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *ThServerUser) GetGameNumber() int32 {
	if m != nil && m.GameNumber != nil {
		return *m.GameNumber
	}
	return 0
}

func (m *ThServerUser) GetIsBreak() bool {
	if m != nil && m.IsBreak != nil {
		return *m.IsBreak
	}
	return false
}

func (m *ThServerUser) GetIsLeave() bool {
	if m != nil && m.IsLeave != nil {
		return *m.IsLeave
	}
	return false
}

// 保存用户的会话信息
type ThServerUserSession struct {
	UserId           *uint32 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	GameStatus       *int32  `protobuf:"varint,2,opt,name=gameStatus" json:"gameStatus,omitempty"`
	GameType         *int32  `protobuf:"varint,3,opt,name=gameType" json:"gameType,omitempty"`
	DeskId           *int32  `protobuf:"varint,4,opt,name=deskId" json:"deskId,omitempty"`
	MatchId          *int32  `protobuf:"varint,5,opt,name=matchId" json:"matchId,omitempty"`
	RoomKey          *string `protobuf:"bytes,6,opt,name=roomKey" json:"roomKey,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ThServerUserSession) Reset()                    { *m = ThServerUserSession{} }
func (m *ThServerUserSession) String() string            { return proto.CompactTextString(m) }
func (*ThServerUserSession) ProtoMessage()               {}
func (*ThServerUserSession) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{4} }

func (m *ThServerUserSession) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *ThServerUserSession) GetGameStatus() int32 {
	if m != nil && m.GameStatus != nil {
		return *m.GameStatus
	}
	return 0
}

func (m *ThServerUserSession) GetGameType() int32 {
	if m != nil && m.GameType != nil {
		return *m.GameType
	}
	return 0
}

func (m *ThServerUserSession) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

func (m *ThServerUserSession) GetMatchId() int32 {
	if m != nil && m.MatchId != nil {
		return *m.MatchId
	}
	return 0
}

func (m *ThServerUserSession) GetRoomKey() string {
	if m != nil && m.RoomKey != nil {
		return *m.RoomKey
	}
	return ""
}

type ThServerAllInJackpot struct {
	UserId           *uint32 `protobuf:"varint,1,opt,name=UserId" json:"UserId,omitempty"`
	Jackpopt         *int64  `protobuf:"varint,2,opt,name=Jackpopt" json:"Jackpopt,omitempty"`
	ThroundCount     *int32  `protobuf:"varint,3,opt,name=ThroundCount" json:"ThroundCount,omitempty"`
	AllInAmount      *int64  `protobuf:"varint,4,opt,name=AllInAmount" json:"AllInAmount,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ThServerAllInJackpot) Reset()                    { *m = ThServerAllInJackpot{} }
func (m *ThServerAllInJackpot) String() string            { return proto.CompactTextString(m) }
func (*ThServerAllInJackpot) ProtoMessage()               {}
func (*ThServerAllInJackpot) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{5} }

func (m *ThServerAllInJackpot) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *ThServerAllInJackpot) GetJackpopt() int64 {
	if m != nil && m.Jackpopt != nil {
		return *m.Jackpopt
	}
	return 0
}

func (m *ThServerAllInJackpot) GetThroundCount() int32 {
	if m != nil && m.ThroundCount != nil {
		return *m.ThroundCount
	}
	return 0
}

func (m *ThServerAllInJackpot) GetAllInAmount() int64 {
	if m != nil && m.AllInAmount != nil {
		return *m.AllInAmount
	}
	return 0
}

// 服务端保存桌子的状态
type ThServerDesk struct {
	Id                   *int32        `protobuf:"varint,1,opt,name=Id" json:"Id,omitempty"`
	DeskOwner            *uint32       `protobuf:"varint,2,opt,name=deskOwner" json:"deskOwner,omitempty"`
	RoomKey              *string       `protobuf:"bytes,3,opt,name=RoomKey" json:"RoomKey,omitempty"`
	DeskType             *int32        `protobuf:"varint,4,opt,name=DeskType" json:"DeskType,omitempty"`
	InitRoomCoin         *int32        `protobuf:"varint,5,opt,name=InitRoomCoin" json:"InitRoomCoin,omitempty"`
	JuCount              *int32        `protobuf:"varint,6,opt,name=JuCount" json:"JuCount,omitempty"`
	SmallBlindCoin       *int32        `protobuf:"varint,7,opt,name=SmallBlindCoin" json:"SmallBlindCoin,omitempty"`
	BigBlindCoin         *int32        `protobuf:"varint,8,opt,name=BigBlindCoin" json:"BigBlindCoin,omitempty"`
	Dealer               *uint32       `protobuf:"varint,9,opt,name=Dealer" json:"Dealer,omitempty"`
	BigBlind             *uint32       `protobuf:"varint,10,opt,name=BigBlind" json:"BigBlind,omitempty"`
	SmallBlind           *uint32       `protobuf:"varint,11,opt,name=SmallBlind" json:"SmallBlind,omitempty"`
	RaiseUserId          *uint32       `protobuf:"varint,12,opt,name=RaiseUserId" json:"RaiseUserId,omitempty"`
	NewRoundFirstBetUser *uint32       `protobuf:"varint,13,opt,name=NewRoundFirstBetUser" json:"NewRoundFirstBetUser,omitempty"`
	BetUserNow           *uint32       `protobuf:"varint,14,opt,name=BetUserNow" json:"BetUserNow,omitempty"`
	GameNumber           *ThServerUser `protobuf:"bytes,15,opt,name=GameNumber" json:"GameNumber,omitempty"`
	//    repeated int32	Users			    =16;	//座位号  //这里通过什么来存取比较好?,user放在thdesk中比较大
	PublicPai        []*Pai                `protobuf:"bytes,17,rep,name=PublicPai" json:"PublicPai,omitempty"`
	UserCount        *int32                `protobuf:"varint,18,opt,name=UserCount" json:"UserCount,omitempty"`
	UserCountOnline  *int32                `protobuf:"varint,19,opt,name=UserCountOnline" json:"UserCountOnline,omitempty"`
	Status           *int32                `protobuf:"varint,20,opt,name=Status" json:"Status,omitempty"`
	BetAmountNow     *int64                `protobuf:"varint,21,opt,name=BetAmountNow" json:"BetAmountNow,omitempty"`
	RoundCount       *int32                `protobuf:"varint,22,opt,name=RoundCount" json:"RoundCount,omitempty"`
	Jackpot          *int64                `protobuf:"varint,23,opt,name=Jackpot" json:"Jackpot,omitempty"`
	EdgeJackpot      *int64                `protobuf:"varint,24,opt,name=edgeJackpot" json:"edgeJackpot,omitempty"`
	MinRaise         *int64                `protobuf:"varint,25,opt,name=MinRaise" json:"MinRaise,omitempty"`
	AllInJackpot     *ThServerAllInJackpot `protobuf:"bytes,26,opt,name=AllInJackpot" json:"AllInJackpot,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *ThServerDesk) Reset()                    { *m = ThServerDesk{} }
func (m *ThServerDesk) String() string            { return proto.CompactTextString(m) }
func (*ThServerDesk) ProtoMessage()               {}
func (*ThServerDesk) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{6} }

func (m *ThServerDesk) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *ThServerDesk) GetDeskOwner() uint32 {
	if m != nil && m.DeskOwner != nil {
		return *m.DeskOwner
	}
	return 0
}

func (m *ThServerDesk) GetRoomKey() string {
	if m != nil && m.RoomKey != nil {
		return *m.RoomKey
	}
	return ""
}

func (m *ThServerDesk) GetDeskType() int32 {
	if m != nil && m.DeskType != nil {
		return *m.DeskType
	}
	return 0
}

func (m *ThServerDesk) GetInitRoomCoin() int32 {
	if m != nil && m.InitRoomCoin != nil {
		return *m.InitRoomCoin
	}
	return 0
}

func (m *ThServerDesk) GetJuCount() int32 {
	if m != nil && m.JuCount != nil {
		return *m.JuCount
	}
	return 0
}

func (m *ThServerDesk) GetSmallBlindCoin() int32 {
	if m != nil && m.SmallBlindCoin != nil {
		return *m.SmallBlindCoin
	}
	return 0
}

func (m *ThServerDesk) GetBigBlindCoin() int32 {
	if m != nil && m.BigBlindCoin != nil {
		return *m.BigBlindCoin
	}
	return 0
}

func (m *ThServerDesk) GetDealer() uint32 {
	if m != nil && m.Dealer != nil {
		return *m.Dealer
	}
	return 0
}

func (m *ThServerDesk) GetBigBlind() uint32 {
	if m != nil && m.BigBlind != nil {
		return *m.BigBlind
	}
	return 0
}

func (m *ThServerDesk) GetSmallBlind() uint32 {
	if m != nil && m.SmallBlind != nil {
		return *m.SmallBlind
	}
	return 0
}

func (m *ThServerDesk) GetRaiseUserId() uint32 {
	if m != nil && m.RaiseUserId != nil {
		return *m.RaiseUserId
	}
	return 0
}

func (m *ThServerDesk) GetNewRoundFirstBetUser() uint32 {
	if m != nil && m.NewRoundFirstBetUser != nil {
		return *m.NewRoundFirstBetUser
	}
	return 0
}

func (m *ThServerDesk) GetBetUserNow() uint32 {
	if m != nil && m.BetUserNow != nil {
		return *m.BetUserNow
	}
	return 0
}

func (m *ThServerDesk) GetGameNumber() *ThServerUser {
	if m != nil {
		return m.GameNumber
	}
	return nil
}

func (m *ThServerDesk) GetPublicPai() []*Pai {
	if m != nil {
		return m.PublicPai
	}
	return nil
}

func (m *ThServerDesk) GetUserCount() int32 {
	if m != nil && m.UserCount != nil {
		return *m.UserCount
	}
	return 0
}

func (m *ThServerDesk) GetUserCountOnline() int32 {
	if m != nil && m.UserCountOnline != nil {
		return *m.UserCountOnline
	}
	return 0
}

func (m *ThServerDesk) GetStatus() int32 {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return 0
}

func (m *ThServerDesk) GetBetAmountNow() int64 {
	if m != nil && m.BetAmountNow != nil {
		return *m.BetAmountNow
	}
	return 0
}

func (m *ThServerDesk) GetRoundCount() int32 {
	if m != nil && m.RoundCount != nil {
		return *m.RoundCount
	}
	return 0
}

func (m *ThServerDesk) GetJackpot() int64 {
	if m != nil && m.Jackpot != nil {
		return *m.Jackpot
	}
	return 0
}

func (m *ThServerDesk) GetEdgeJackpot() int64 {
	if m != nil && m.EdgeJackpot != nil {
		return *m.EdgeJackpot
	}
	return 0
}

func (m *ThServerDesk) GetMinRaise() int64 {
	if m != nil && m.MinRaise != nil {
		return *m.MinRaise
	}
	return 0
}

func (m *ThServerDesk) GetAllInJackpot() *ThServerAllInJackpot {
	if m != nil {
		return m.AllInJackpot
	}
	return nil
}

//
type CsThRankInfo struct {
	UserId           *uint32 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	MatchId          *int32  `protobuf:"varint,2,opt,name=matchId" json:"matchId,omitempty"`
	Balance          *int64  `protobuf:"varint,3,opt,name=balance" json:"balance,omitempty"`
	EndTime          *int64  `protobuf:"varint,4,opt,name=endTime" json:"endTime,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CsThRankInfo) Reset()                    { *m = CsThRankInfo{} }
func (m *CsThRankInfo) String() string            { return proto.CompactTextString(m) }
func (*CsThRankInfo) ProtoMessage()               {}
func (*CsThRankInfo) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{7} }

func (m *CsThRankInfo) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *CsThRankInfo) GetMatchId() int32 {
	if m != nil && m.MatchId != nil {
		return *m.MatchId
	}
	return 0
}

func (m *CsThRankInfo) GetBalance() int64 {
	if m != nil && m.Balance != nil {
		return *m.Balance
	}
	return 0
}

func (m *CsThRankInfo) GetEndTime() int64 {
	if m != nil && m.EndTime != nil {
		return *m.EndTime
	}
	return 0
}

func init() {
	proto.RegisterType((*TBetRecord)(nil), "bbproto.TBetRecord")
	proto.RegisterType((*TZjhRound)(nil), "bbproto.TZjhRound")
	proto.RegisterType((*TNotice)(nil), "bbproto.TNotice")
	proto.RegisterType((*ThServerUser)(nil), "bbproto.ThServerUser")
	proto.RegisterType((*ThServerUserSession)(nil), "bbproto.ThServerUserSession")
	proto.RegisterType((*ThServerAllInJackpot)(nil), "bbproto.ThServerAllInJackpot")
	proto.RegisterType((*ThServerDesk)(nil), "bbproto.ThServerDesk")
	proto.RegisterType((*CsThRankInfo)(nil), "bbproto.CsThRankInfo")
}

var fileDescriptor4 = []byte{
	// 876 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x55, 0xdb, 0x72, 0xdb, 0x36,
	0x10, 0xad, 0x44, 0x29, 0x32, 0x61, 0xc9, 0x8a, 0x69, 0x39, 0x61, 0x3d, 0xed, 0x34, 0xa3, 0xbe,
	0xb8, 0x2f, 0x7e, 0x48, 0xfb, 0x03, 0x91, 0xdc, 0x8b, 0xd2, 0xd8, 0x51, 0x25, 0x7a, 0xda, 0xe9,
	0x1b, 0x48, 0x6e, 0x2d, 0xc4, 0x24, 0xe0, 0x21, 0xc1, 0x68, 0x9c, 0x99, 0x7e, 0x5c, 0x7f, 0xa3,
	0xbf, 0xd0, 0x9f, 0xe8, 0x62, 0x01, 0x50, 0x72, 0xdd, 0x37, 0xfb, 0x00, 0xdc, 0x3d, 0xe7, 0xec,
	0x59, 0x88, 0x1d, 0xd7, 0x50, 0x7d, 0x84, 0xea, 0x4a, 0xe5, 0x50, 0x5c, 0xdc, 0x57, 0x4a, 0xab,
	0x68, 0x90, 0xa6, 0xf4, 0xc7, 0x19, 0x4b, 0x79, 0x0d, 0x16, 0x3c, 0x1b, 0x7f, 0xda, 0xf0, 0x0f,
	0x42, 0x6e, 0x1a, 0xee, 0x80, 0x91, 0xde, 0x2c, 0xd5, 0x1d, 0x54, 0xf6, 0xdf, 0xe9, 0x6f, 0x8c,
	0x25, 0x33, 0xd0, 0x2b, 0xc8, 0x54, 0x95, 0x47, 0x47, 0xec, 0x59, 0x83, 0x85, 0x17, 0x79, 0xdc,
	0x79, 0xd5, 0x39, 0x1f, 0x45, 0x63, 0x36, 0x48, 0x41, 0x7f, 0x52, 0x12, 0xe2, 0xee, 0xab, 0xe0,
	0xbc, 0x1f, 0xbd, 0x60, 0x47, 0xbf, 0x7f, 0xd8, 0xac, 0x54, 0x23, 0xf3, 0xeb, 0xa6, 0x4c, 0xa1,
	0x8a, 0x03, 0xbc, 0x18, 0x46, 0xc7, 0x2c, 0xdc, 0x0a, 0xf9, 0xa6, 0xc4, 0x03, 0x1d, 0xf7, 0x10,
	0xea, 0x4f, 0xff, 0xee, 0xb0, 0x30, 0xf1, 0x97, 0x23, 0xc6, 0xba, 0xc2, 0x57, 0xc5, 0xcb, 0x29,
	0xdc, 0x0a, 0x99, 0x88, 0xd2, 0xd4, 0xed, 0x9c, 0x07, 0x11, 0x9e, 0x63, 0xa3, 0xef, 0x65, 0x4e,
	0x58, 0x40, 0xd8, 0x09, 0x3b, 0x2c, 0x94, 0xd6, 0x50, 0x3d, 0x10, 0xd8, 0x23, 0x10, 0x19, 0x81,
	0xbb, 0xd5, 0xf7, 0x5f, 0x1a, 0x7e, 0xae, 0xf5, 0x33, 0x62, 0x79, 0xca, 0x46, 0x06, 0xfb, 0xb5,
	0x65, 0x34, 0x20, 0x78, 0xc2, 0x86, 0x33, 0x2e, 0x51, 0xfb, 0x8d, 0xd5, 0x78, 0x40, 0x6c, 0xbe,
	0x66, 0x0c, 0x59, 0x2e, 0xb9, 0x78, 0x27, 0x6a, 0x1d, 0x87, 0x78, 0xf3, 0xf0, 0xf5, 0xf8, 0xc2,
	0x79, 0x79, 0x61, 0x8f, 0x8c, 0x31, 0x4e, 0x2f, 0x33, 0x7a, 0xa7, 0x82, 0x0d, 0x92, 0x6b, 0xa5,
	0x45, 0x06, 0x7b, 0xca, 0xfa, 0x86, 0x8c, 0x24, 0x34, 0x79, 0xb8, 0xb7, 0xd2, 0xfa, 0x46, 0x86,
	0xc3, 0x84, 0x2e, 0xc0, 0xf9, 0x85, 0x0c, 0x2d, 0x38, 0x57, 0x52, 0x83, 0xf3, 0x2c, 0xdc, 0x7d,
	0x7f, 0x05, 0xa5, 0x22, 0x81, 0xe1, 0xf4, 0x9f, 0x2e, 0x1b, 0x26, 0x9b, 0x35, 0x8d, 0xdb, 0x10,
	0x8f, 0x86, 0xac, 0xb7, 0x06, 0xae, 0x5d, 0x4b, 0x64, 0xb6, 0xd6, 0x5c, 0x37, 0xf5, 0xae, 0xdd,
	0xac, 0x02, 0x7e, 0xe7, 0xc0, 0x80, 0xc0, 0xaf, 0x58, 0xf8, 0x13, 0x97, 0xf9, 0x9c, 0x57, 0x79,
	0x8d, 0xad, 0x8c, 0xc4, 0x61, 0x2b, 0xd1, 0xe8, 0x43, 0x5b, 0xb7, 0x5c, 0xb4, 0xb6, 0x86, 0xd1,
	0x73, 0x76, 0x80, 0x80, 0xbe, 0xb9, 0x59, 0x5c, 0xa2, 0xa9, 0x06, 0xc1, 0x46, 0x39, 0xd4, 0x77,
	0xe8, 0xdb, 0x80, 0x6a, 0xe2, 0x8d, 0x44, 0x69, 0x5e, 0x60, 0x7a, 0xc8, 0xc9, 0x20, 0x3a, 0x63,
	0x91, 0x47, 0xbe, 0xcb, 0x78, 0x91, 0xbd, 0x29, 0x0a, 0x21, 0xd1, 0x51, 0x73, 0xf6, 0x28, 0x20,
	0x8c, 0x20, 0x53, 0xa0, 0xa9, 0xe4, 0x5c, 0xe1, 0xa5, 0x43, 0x8f, 0x10, 0x4d, 0x83, 0x0c, 0x3d,
	0xb2, 0x52, 0xaa, 0x24, 0x64, 0x44, 0xc8, 0x4b, 0x36, 0x6e, 0x0b, 0x5d, 0x82, 0xe6, 0xa2, 0x88,
	0x8f, 0x50, 0x50, 0x60, 0xf8, 0xb9, 0xb9, 0x8e, 0x69, 0xae, 0xe8, 0xe5, 0x8f, 0xbc, 0x04, 0x37,
	0xb6, 0xe7, 0xc4, 0x19, 0x65, 0x2e, 0x6a, 0xb2, 0x27, 0x3e, 0x46, 0xe0, 0xc0, 0x02, 0xef, 0x80,
	0x7f, 0x84, 0x38, 0x32, 0xc0, 0xf4, 0x4f, 0x76, 0xb2, 0x6f, 0xf6, 0x1a, 0xea, 0x5a, 0x28, 0xf9,
	0x64, 0x31, 0xb0, 0xf8, 0x2d, 0x16, 0x7f, 0xe4, 0x3c, 0x72, 0x35, 0x18, 0x8d, 0x3e, 0xf0, 0xb3,
	0x71, 0x96, 0xf5, 0x7c, 0xfb, 0x92, 0xeb, 0x6c, 0x83, 0x40, 0xdf, 0x03, 0x15, 0xca, 0xfb, 0x19,
	0x1e, 0xac, 0xc9, 0xd3, 0x8c, 0x4d, 0x7c, 0x7b, 0x74, 0x6f, 0x21, 0xdf, 0xf2, 0xec, 0xee, 0x5e,
	0xe9, 0x3d, 0x71, 0xb6, 0x3f, 0xf6, 0xb2, 0x47, 0xf7, 0xda, 0x6d, 0xd0, 0xc4, 0xa4, 0xa4, 0x32,
	0xbb, 0x36, 0x27, 0x8f, 0x03, 0x9f, 0x06, 0xaa, 0xb3, 0xb7, 0x99, 0xc1, 0xf4, 0xaf, 0xde, 0x2e,
	0x51, 0x97, 0xc8, 0xcf, 0x44, 0x78, 0xe1, 0x23, 0x8c, 0x83, 0x32, 0x9c, 0xdf, 0x6f, 0x25, 0xba,
	0xd6, 0xf5, 0xaf, 0xc0, 0xca, 0xb1, 0x0c, 0x7c, 0x38, 0xcc, 0x77, 0xa4, 0xd4, 0x2a, 0xc3, 0xee,
	0x0b, 0x29, 0x74, 0x3b, 0xab, 0x56, 0xde, 0xdb, 0x66, 0xee, 0x16, 0xb3, 0x63, 0x9f, 0x8f, 0x75,
	0xc9, 0x8b, 0x62, 0x86, 0xc1, 0xb0, 0x63, 0x1e, 0xf8, 0xcf, 0x67, 0xe2, 0x76, 0x87, 0x1e, 0x78,
	0xfb, 0x2e, 0x81, 0x17, 0xc8, 0x23, 0xf4, 0xa2, 0xfd, 0x2d, 0x8a, 0x10, 0x8d, 0x61, 0x57, 0x8f,
	0x42, 0x34, 0x32, 0x92, 0x57, 0x5c, 0xd4, 0xe0, 0xfc, 0x1a, 0x12, 0xf8, 0x05, 0x9b, 0x5c, 0xc3,
	0x96, 0x9e, 0xa2, 0x1f, 0x44, 0x55, 0x6b, 0x8c, 0xa8, 0x39, 0xa7, 0x4c, 0x51, 0x19, 0x07, 0x5c,
	0xab, 0x2d, 0xc6, 0xc9, 0x60, 0xdf, 0x3c, 0x8a, 0x8f, 0x89, 0xd4, 0xe1, 0xeb, 0xd3, 0x76, 0x67,
	0x1e, 0x2d, 0x24, 0x6e, 0xd7, 0xb2, 0x49, 0x0b, 0x91, 0xe1, 0x26, 0x61, 0xae, 0x9e, 0x6e, 0x17,
	0x7a, 0x6a, 0x2e, 0x5a, 0x27, 0x22, 0xd2, 0x86, 0x31, 0x6e, 0xa1, 0xf7, 0x12, 0xe9, 0x43, 0x7c,
	0xf2, 0x9f, 0x7d, 0x9e, 0xb4, 0xd6, 0x80, 0xb6, 0xf3, 0x33, 0xec, 0x4e, 0xfd, 0xab, 0xb7, 0xda,
	0xcd, 0xfa, 0x45, 0xeb, 0xb6, 0x8d, 0x4b, 0xfc, 0xd2, 0x3f, 0xa0, 0x90, 0xdf, 0x82, 0x07, 0x63,
	0xbf, 0x51, 0x57, 0x42, 0x92, 0x43, 0xf1, 0xe7, 0x84, 0x7c, 0xcb, 0x86, 0xfb, 0x59, 0x8b, 0xcf,
	0x48, 0xeb, 0x97, 0x4f, 0xb4, 0xee, 0x5f, 0x9a, 0xfe, 0xc2, 0x86, 0xf3, 0x3a, 0xd9, 0xac, 0xf0,
	0x3d, 0x5d, 0xc8, 0x3f, 0xd4, 0xff, 0xfd, 0x72, 0xf8, 0xa8, 0x77, 0x3d, 0xbb, 0x94, 0x17, 0x5c,
	0x66, 0xfe, 0x79, 0xdf, 0x7b, 0xc9, 0x29, 0x96, 0xcb, 0xcf, 0x96, 0x9d, 0x65, 0xf7, 0xdf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x4d, 0x30, 0xb5, 0x63, 0xd9, 0x06, 0x00, 0x00,
}
