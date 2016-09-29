package room

import (
	"casino_server/common/log"
	"casino_server/msg/bbprotogo"
	"casino_server/service/pokerService"
)

//牌的花色 和值
var (
	POKER_COLOR_DIAMOND int32 = 0    //方片
	POKER_COLOR_CLUB int32 = 1    //梅花
	POKER_COLOR_HEARTS int32 = 2    //红桃
	POKER_COLOR_SPADE int32 = 3    //黑桃
	POKER_COLOR_COUNT int32 = 4
)

var (
	POKER_VALUE_A int32 = 0
	POKER_VALUE_2 int32 = 1
	POKER_VALUE_3 int32 = 2
	POKER_VALUE_4 int32 = 3
	POKER_VALUE_5 int32 = 4
	POKER_VALUE_6 int32 = 5
	POKER_VALUE_7 int32 = 6
	POKER_VALUE_8 int32 = 7
	POKER_VALUE_9 int32 = 8
	POKER_VALUE_10 int32 = 9
	POKER_VALUE_J int32 = 10
	POKER_VALUE_Q int32 = 11
	POKER_VALUE_K int32 = 12
	POKER_VALUE_COUNT int32 = 13
	POKER_VALUE_BACK int32 = 52    // 牌背
	POKER_VALUE_EMPTY int32 = 53    // 没牌
)



//th的牌转换成OG的牌
func ThCard2OGCard(pai *bbproto.Pai) *bbproto.Game_CardInfo {
	result := &bbproto.Game_CardInfo{}
	result.Value = new(int32)
	result.Color = new(int32)

	//log.T("初始化牌的花色:*pai.flower[%v]",*pai.Flower)

	//初始化花色
	switch *pai.Flower {
	case "heart" :
		*result.Color = POKER_COLOR_HEARTS
	case "diamond" :
		*result.Color = POKER_COLOR_DIAMOND
	case "club" :
		*result.Color = POKER_COLOR_CLUB
	case "spade" :
		*result.Color = POKER_COLOR_SPADE
	}

	//初始化值
	switch *pai.Value {
	case 2:
		*result.Value = POKER_VALUE_2
	case 3:
		*result.Value = POKER_VALUE_3
	case 4:
		*result.Value = POKER_VALUE_4
	case 5:
		*result.Value = POKER_VALUE_5
	case 6:
		*result.Value = POKER_VALUE_6
	case 7:
		*result.Value = POKER_VALUE_7
	case 8:
		*result.Value = POKER_VALUE_8
	case 9:
		*result.Value = POKER_VALUE_9
	case 10:
		*result.Value = POKER_VALUE_10
	case 11:                                //j
		*result.Value = POKER_VALUE_J
	case 12:                                //q
		*result.Value = POKER_VALUE_Q
	case 13:                                //k
		*result.Value = POKER_VALUE_K
	case 14:                                //a
		*result.Value = POKER_VALUE_A

	}

	return result
}


//获得前注的信息
func ( t *ThDesk) GetPreCoin() []int64 {
	var result []int64
	for i := 0; i < len(t.Users); i++ {
		u := t.Users[i]
		if u != nil {
			result = append(result, int64(u.PreCoin))
		}
	}
	return result
}

//得到roomCoin
func (mydesk *ThDesk) GetRoomCoin() []int64 {
	var result []int64
	for i := 0; i < len(mydesk.Users); i++ {
		u := mydesk.Users[i]
		if u != nil {
			result = append(result, int64(u.GetRoomCoin()))
		}
	}
	return result
}



//解析每个人下注的金额
func (mydesk *ThDesk) GetHandCoin() []int64 {
	var result []int64
	for i := 0; i < len(mydesk.Users); i++ {
		u := mydesk.Users[i]
		if u != nil {
			//用户手牌
			result = append(result, int64(u.HandCoin))
		}
	}
	return result
}


//
func NewGame_SendOverTurn() *bbproto.Game_SendOverTurn {
	ret := &bbproto.Game_SendOverTurn{}
	ret.NextSeat = new(int32)
	ret.Tableid = new(int32)
	ret.Pool = new(int64)
	ret.MinRaise = new(int64)
	return ret
}

//取边池的大小
func (t *ThDesk) GetSecondPool() []int64 {
	var ret []int64
	for i := 0; i < len(t.AllInJackpot); i++ {
		a := t.AllInJackpot[i]
		if a != nil {
			ret = append(ret, a.Jackpopt)
		}
	}

	//边池
	ret = append(ret, t.EdgeJackpot)
	return ret
}

//发送新增用户的广播
func (t *ThDesk) BroadGameInfo(reqUserId uint32) {
	log.T("开始发送用户gameinfo的广播")
	msg := t.initGameSendgameInfoByDesk(reqUserId)
	for i := 0; i < len(t.Users); i++ {
		u := t.Users[i]
		if u != nil && !u.IsLeave && !u.IsBreak {
			*msg.Seat = t.Users[i].Seat                        //给用户发送广播的时候需要判断自己的座位号是多少
			msg.Handcard = t.GetMyHandCard(t.Users[i].UserId)        //针对用户发送手牌的信息
			//给用户发送信息
			u.WriteMsg(msg)
		}
	}
}



//返回房间的信息 todo 登陆成功的处理,给请求登陆的玩家返回登陆结果的消息
func (mydesk *ThDesk) initGameSendgameInfoByDesk(reqUserId uint32) *bbproto.Game_SendGameInfo {
	result := bbproto.NewGame_SendGameInfo()
	//初始化桌子相关的信息
	*result.Tableid = int32(mydesk.Id)                                //桌子的Id
	*result.TablePlayer = mydesk.GetUserCount()                        //玩家总人数
	*result.BankSeat = mydesk.GetUserSeatByUserId(mydesk.Dealer)        //int32(mydesk.GetUserIndex(mydesk.Dealer))        //庄家
	*result.ChipSeat = mydesk.GetUserSeatByUserId(mydesk.BetUserNow) //int32(mydesk.GetUserIndex(mydesk.BetUserNow))//当前活动玩家
	*result.ActionTime = ThdeskConfig.TH_TIMEOUT_DURATION_INT       //当前操作时间,服务器当前的时间
	*result.DelayTime = int32(1000)                                //当前延时时间
	*result.GameStatus = deskStatus2OG(mydesk)
	*result.Pool = int64(mydesk.Jackpot)                                //奖池
	result.Publiccard = mydesk.ThPublicCard2OGC()                        //公共牌...
	*result.MinRaise = mydesk.GetMinRaise()                                //最低加注金额
	*result.NInitActionTime = ThdeskConfig.TH_TIMEOUT_DURATION_INT
	*result.NInitDelayTime = ThdeskConfig.TH_TIMEOUT_DURATION_INT
	result.HandCoin = mydesk.GetRoomCoin()                        //带入金额
	result.TurnCoin = mydesk.getTurnCoin()
	result.SecondPool = mydesk.GetSecondPool()
	*result.TurnMax = mydesk.BetAmountNow
	result.WeixinInfos = mydesk.GetWeiXinInfos()
	*result.OwnerSeat = mydesk.GetUserSeatByUserId(mydesk.DeskOwner)
	*result.PreCoin = mydesk.PreCoin
	*result.SmallBlind = mydesk.SmallBlindCoin
	*result.BigBlind = mydesk.BigBlindCoin
	*result.RoomType = mydesk.GameType
	*result.InitRoomCoin = mydesk.InitRoomCoin
	*result.CurrPlayCount = mydesk.JuCountNow
	*result.TotalPlayCount = mydesk.JuCount
	*result.SenderUserId = reqUserId


	//循环User来处理
	for i := 0; i < len(mydesk.Users); i++ {
		u := mydesk.Users[i]
		if u != nil {
			//用户当前状态
			result.BAllIn = append(result.BAllIn, isAllIn(u))                        //是否已经全下了
			result.BBreak = append(result.BBreak, isBreak(u))                        //是否
			result.BLeave = append(result.BLeave, isLeave(u))
			result.BEnable = append(result.BEnable, isEnable(u))                        //用户是否可以操作,0表示不能操作,1表示可以操作
			result.BFold = append(result.BFold, isFold(u))        //是否弃牌
			//nickName			//seatId
			result.NickName = append(result.NickName, u.NickName)
			result.SeatId = append(result.SeatId, int32(i))
			result.BReady = append(result.BReady, isReady(u))
		} else {

		}
	}
	return result

}

func isReady(u *ThUser) int32 {
	if u.IsReady() {
		return 1
	} else {
		return 0
	}
}

//判断是否allIn
func isAllIn(u *ThUser) int32 {
	if u.Status == TH_USER_STATUS_ALLINING {
		return 1
	} else {
		return 0
	}
}


//判断是否是掉线
func isBreak(u *ThUser) int32 {
	if u.IsBreak {
		return 1
	} else {
		return 0
	}
}

//判断是否离开房间
func isLeave(u *ThUser) int32 {
	if u.IsLeave {
		return 1
	} else {
		return 0
	}
}

//判断是否allIn
func isFinalAddon(u *ThUser) int32 {
	return 0
}

//判断是否allIn
func isEnable(u *ThUser) int32 {
	if u.Status == TH_USER_STATUS_BETING {
		//游戏中
		return 1
	} else {
		return 0
	}
}

//判断是否allIn
func isFold(u *ThUser) int32 {
	if u.Status == TH_USER_STATUS_FOLDED {
		//游戏中
		return 1
	} else {
		return 0
	}
}

//解析TurnCoin
func (mydesk *ThDesk) getTurnCoin() []int64 {
	var result []int64
	for i := 0; i < len(mydesk.Users); i++ {
		u := mydesk.Users[i]
		if u != nil {
			//用户手牌
			result = append(result, int64(u.TurnCoin))
		}
	}
	return result
}


//游戏状态的转换
func deskStatus2OG(desk *ThDesk) int32 {
	var result int32 = GAME_STATUS_READY
	//status := desk.Status
	round := desk.RoundCount

	if desk.IsStop() {
		//没有开始
		result = GAME_STATUS_READY
	} else if desk.IsRun() || desk.IsLottery() {
		switch round {
		case TH_DESK_ROUND1:
			result = GAME_STATUS_FIRST_TURN
		case TH_DESK_ROUND2:
			result = GAME_STATUS_SECOND_TURN
		case TH_DESK_ROUND3:
			result = GAME_STATUS_THIRD_TURN
		case TH_DESK_ROUND4:
			result = GAME_STATUS_LAST_TURN
		default:
			result = GAME_STATUS_DEAL_CARDS
		}
	} else if desk.IsLottery() {
		result = GAME_STATUS_SHOW_RESULT
	}

	return result
}



//发送牌的时候
func (t *ThDesk) OGTHBroadInitCard(msg *bbproto.Game_InitCard) {
	for i := 0; i < len(t.Users); i++ {
		if t.Users[i] != nil {
			//给用户发送广播的时候需要判断自己的座位号是多少
			a := t.Users[i].Agent
			msg.HandCard = t.GetMyHandCard(t.Users[i].UserId)
			a.WriteMsg(msg)
		}
	}
}

//解析手牌
func (mydesk *ThDesk ) GetHandCard() []*bbproto.Game_CardInfo {
	//log.T("把desk的手牌,转化为og的手牌")
	var handCard []*bbproto.Game_CardInfo
	for i := 0; i < len(mydesk.Users); i++ {
		u := mydesk.Users[i]
		if u != nil {
			//log.T("开始给玩家[%v]解析手牌", u.UserId)
			result := make([]*bbproto.Game_CardInfo, 0)
			//用户手牌
			if len(u.HandCards) == 2 {
				for i := 0; i < len(u.HandCards); i++ {
					c := u.HandCards[i]
					gc := ThCard2OGCard(c)
					//增加到数组中
					result = append(result, gc)
				}
			} else {
				log.T("玩家[%v]的手牌为空", u.UserId)
				gc := &bbproto.Game_CardInfo{}
				gc.Color = new(int32)
				gc.Value = new(int32)
				*gc.Color = POKER_COLOR_COUNT
				*gc.Value = POKER_VALUE_EMPTY
				result = append(result, gc)
				result = append(result, gc)
			}
			handCard = append(handCard, result...)
		}

	}
	return handCard
}


//解析手牌
func (mydesk *ThDesk ) GetMyHandCard(userId uint32) []*bbproto.Game_CardInfo {
	//log.T("把desk的手牌,转化为og的手牌")
	var handCard []*bbproto.Game_CardInfo
	for i := 0; i < len(mydesk.Users); i++ {
		u := mydesk.Users[i]
		if u != nil {
			result := make([]*bbproto.Game_CardInfo, 0)
			//用户手牌
			if len(u.HandCards) == 2 && u.UserId == userId {
				log.T("开始转化user[%v]的手牌", userId)
				for i := 0; i < len(u.HandCards); i++ {
					c := u.HandCards[i]
					gc := ThCard2OGCard(c)
					//增加到数组中
					result = append(result, gc)
				}
			} else {
				//log.T("玩家[%v]的手牌为空", u.UserId)
				gc := &bbproto.Game_CardInfo{}
				gc.Color = new(int32)
				gc.Value = new(int32)
				*gc.Color = POKER_COLOR_COUNT
				*gc.Value = POKER_VALUE_EMPTY
				result = append(result, gc)
				result = append(result, gc)
			}
			handCard = append(handCard, result...)
		} else {

		}

	}
	return handCard
}

func NewGame_WinCoin() *bbproto.Game_WinCoin {
	gwc := &bbproto.Game_WinCoin{}
	gwc.Card1 = new(int32)
	gwc.Card2 = new(int32)
	gwc.Card3 = new(int32)
	gwc.Card4 = new(int32)
	gwc.Card5 = new(int32)
	gwc.Cardtype = new(int32)
	gwc.Coin = new(int64)
	gwc.Seat = new(int32)
	gwc.PoolIndex = new(int32)
	gwc.IsDiscard = new(bool)
	gwc.IsHideCard = new(bool)
	return gwc
}

//
func (t *ThDesk) getWinCoinInfo() []*bbproto.Game_WinCoin {
	var ret []*bbproto.Game_WinCoin
	//对每个人做计算
	for i := 0; i < len(t.Users); i++ {
		u := t.Users[i]
		if u != nil && u.IsClose() && u.winAmount > 0  && u.thCards != nil && u.HandCards != nil {
			//开是对这个人计算
			gwc := GetWinCoinCardInfo(u, 0, t)
			*gwc.Coin = u.winAmount                                        //这个表示的是赢得的底池多少钱
			ret = append(ret, gwc)
		}
	}
	return ret
}


//得到coinInfo 的数据
func (t *ThDesk) getCoinInfo(buser *ThUser) []*bbproto.Game_WinCoin {
	var ret []*bbproto.Game_WinCoin
	//对每个人做计算
	for i := 0; i < len(t.Users); i++ {
		u := t.Users[i]
		if u != nil && u.IsClose() && u.thCards != nil && u.HandCards != nil {
			//log.T("开始计算user[%v]一局比赛的得失的coinInfo,status[%v],thcards[%v],t.RoundCount[%v],u.winAmount[%v],u.totalBet[%v],u.isAllin[%v],u.isBettion[%v]",
			//	u.UserId, u.Status, u.thCards, t.RoundCount, u.winAmount, u.TotalBet, u.IsAllIn(), u.IsBetting())
			//开是对这个人计算
			gwc := GetWinCoinCardInfo(u, buser.UserId, t)
			//这里需要先对牌进行排序
			*gwc.Coin = u.winAmount - u.TotalBet                      //表示除去押注的金额,净赚多少钱
			gwc.Rolename = []byte(u.NickName)
			ret = append(ret, gwc)
		}
	}
	return ret
}

func GetWinCoinCardInfo(u *ThUser, buserId uint32, desk *ThDesk) *bbproto.Game_WinCoin {
	gwc := NewGame_WinCoin()

	//gettingOrALlingCount := desk.GetBettingOrAllinCount()
	//如果需要两排，或者发送广播的目标是自己,那么直接两排。
	if u.IsShowCard || u.UserId == buserId {
		//

		//这里的代码不能删除，开总会让这里的逻辑回滚，我猜的...
		//只有在allin或者betting的人才显示5张牌
		//if (u.IsAllIn() || u.IsBetting()) && gettingOrALlingCount >= 2 {
		//	paicards := OGTHCardPaixu(u.thCards)
		//	*gwc.Card1 = paicards[0].GetMapKey()
		//	*gwc.Card2 = paicards[1].GetMapKey()
		//	*gwc.Card3 = paicards[2].GetMapKey()
		//	*gwc.Card4 = paicards[3].GetMapKey()
		//	*gwc.Card5 = paicards[4].GetMapKey()
		//	*gwc.Cardtype = u.thCards.GetOGCardType()
		//} else {
		//	//只发送手牌
		//	*gwc.Card1 = u.HandCards[0].GetMapKey()
		//	*gwc.Card2 = u.HandCards[1].GetMapKey()
		//	*gwc.Card3 = 53
		//	*gwc.Card4 = 53
		//	*gwc.Card5 = 53
		//	*gwc.Cardtype = 0
		//}

		paiCount := 0
		if desk.SendFlop {
			paiCount = 3
		}

		if desk.SendTurn {
			paiCount = 4
		}

		if desk.SendRive {
			paiCount = 5
		}

		if paiCount > 0 {
			thcards := pokerService.GetTHPoker(u.HandCards, desk.PublicPai[:paiCount], 5)
			paicards := OGTHCardPaixu(thcards)
			*gwc.Card1 = paicards[0].GetMapKey()
			*gwc.Card2 = paicards[1].GetMapKey()
			*gwc.Card3 = paicards[2].GetMapKey()
			*gwc.Card4 = paicards[3].GetMapKey()
			*gwc.Card5 = paicards[4].GetMapKey()
			*gwc.Cardtype = thcards.GetOGCardType()
			*gwc.IsHideCard = false
		} else {
			*gwc.Card1 = u.HandCards[0].GetMapKey()
			*gwc.Card2 = u.HandCards[1].GetMapKey()
			*gwc.Card3 = 53
			*gwc.Card4 = 53
			*gwc.Card5 = 53
			*gwc.Cardtype = 0
			*gwc.IsHideCard = true
		}

	} else {
		*gwc.Card1 = 53
		*gwc.Card2 = 53
		*gwc.Card3 = 53
		*gwc.Card4 = 53
		*gwc.Card5 = 53
		*gwc.Cardtype = 9
		*gwc.IsHideCard = true

	}

	*gwc.Seat = u.Seat
	*gwc.PoolIndex = int32(0)
	*gwc.IsDiscard = u.IsFold()        //是否弃牌

	return gwc

}

func (t *ThDesk) GetBshowCard() []int32 {
	var ret []int32
	for i := 0; i < len(t.Users); i++ {
		u := t.Users[i]
		if u != nil {
			ret = append(ret, int32(1))
		}
	}

	return ret
}

//把结果牌 针对og来排序
func OGTHCardPaixu(s *pokerService.ThCards) []*bbproto.Pai {

	//判断参数是否nil
	if s == nil {
		return nil
	}

	lensCards := len(s.Cards)
	ret := make([]*bbproto.Pai, lensCards)

	for i := 0; i < lensCards; i++ {
		//log.T("og 排序之前[%v]----[%v]:", i, s.Cards[i].GetValue())
	}

	if s.ThType == pokerService.THPOKER_TYPE_SITIAO {
		//如果是四条的话,前四张是四条
		if s.KeyValue[0] == s.Cards[0].GetValue() {
			copy(ret, s.Cards)
		} else {
			copy(ret[0:4], s.Cards[1:5])
			copy(ret[4:5], s.Cards[0:1])
		}
	} else if s.ThType == pokerService.THPOKER_TYPE_HULU {
		//如果是葫芦的话,前三张是三条
		if s.KeyValue[0] == s.Cards[0].GetValue() {
			copy(ret, s.Cards)
		} else {
			copy(ret[0:3], s.Cards[2:5])
			copy(ret[3:5], s.Cards[0:2])
		}
	} else if s.ThType == pokerService.THPOKER_TYPE_SANTIAO {
		//如果是三条的话,前三张是三条
		santiaoIndex := 0
		for i := 0; i < lensCards; i++ {
			if s.Cards[i].GetValue() == s.KeyValue[0] {
				santiaoIndex = i
				break
			}
		}
		ret[0] = s.Cards[santiaoIndex]
		ret[1] = s.Cards[santiaoIndex + 1]
		ret[2] = s.Cards[santiaoIndex + 2]

		log.T("排序葫芦1:[%v]", ret)
		tempS := make([]*bbproto.Pai, lensCards)
		copy(tempS, s.Cards)
		for i := 3; i < 5; i++ {
			for j := 0; j < lensCards; j++ {
				ts := tempS[j]
				if ts != nil && ts.GetValue() != ret[0].GetValue() {
					ret[i] = ts
					log.T("排序葫芦[%v]:[%v]", i, ret[i])
					tempS[j] = nil
					break
				}
			}
		}
	} else if s.ThType == pokerService.THPOKER_TYPE_LIANGDUI {
		//如果是两队的话,

		yiduiIndex, liangduiIndex, danpai := 0, 0, 0

		for i := 0; i < lensCards; i++ {
			if s.Cards[i].GetValue() == s.KeyValue[0] {
				yiduiIndex = i
				break
			}
		}

		for i := 0; i < lensCards; i++ {
			if s.Cards[i].GetValue() == s.KeyValue[1] {
				liangduiIndex = i
				break
			}
		}

		for i := 0; i < lensCards; i++ {
			if s.Cards[i].GetValue() == s.KeyValue[2] {
				danpai = i
				break
			}
		}

		copy(ret[0:2], s.Cards[yiduiIndex:yiduiIndex + 2])
		copy(ret[2:4], s.Cards[liangduiIndex:liangduiIndex + 2])
		copy(ret[4:lensCards], s.Cards[danpai:danpai + 1])

	} else if s.ThType == pokerService.THPOKER_TYPE_YIDUI {
		//如果是一对的话,
		//1,找到一对的坐标
		tempS := make([]*bbproto.Pai, lensCards)
		copy(tempS, s.Cards)
		log.T("temps:[%v]", tempS)

		yiduiIndex := 0
		for i := 0; i < lensCards; i++ {
			if s.Cards[i].GetValue() == s.KeyValue[0] {
				yiduiIndex = i
				break
			}
		}
		ret[0] = tempS[yiduiIndex]
		tempS[yiduiIndex] = nil
		ret[1] = tempS[yiduiIndex + 1]
		tempS[yiduiIndex + 1] = nil

		for i := 2; i < 5; i++ {
			for j := 0; j < lensCards; j++ {
				ts := tempS[j]
				if ts != nil {
					ret[i] = ts
					tempS[j] = nil
					//log.T("排序一对[%v],[%v]", i, ret[i])
					break
				}
			}
		}

	} else if s.ThType == pokerService.THPOKER_TYPE_SHUNZI && s.IsWheel {
		copy(ret[0:4], s.Cards[1:5])
		copy(ret[4:5], s.Cards[0:1])

	} else {
		copy(ret, s.Cards)
	}

	for i := 0; i < lensCards; i++ {
		//log.T("og 排序之后[%v]----[%v]:", i, ret[i].GetValue())
	}
	return ret
}