package majiang

import (
	"casino_majiang/msg/protogo"
	"errors"
	"github.com/name5566/leaf/gate"
	"casino_majiang/service/lock"
	"casino_majiang/conf/config"
	"casino_majiang/msg/funcsInit"
	"casino_common/utils/numUtils"
	"casino_common/common/Error"
	"casino_common/common/consts"
	"casino_common/utils/rand"
	"casino_common/common/log"
	"casino_common/common/userService"
	"casino_common/utils/db"
)


//普通的麻将房间...
func init() {
	FMJRoomIns = NewMjRoom()
	FMJRoomIns.OnInit()
}

var FMJRoomIns *MjRoom

//初始化
func (r *MjRoom) OnInit() {
	log.T("初始化麻将的room")

}

//更具条件计算创建房间的费用
func (r *MjRoom) CalcCreateFee() int64 {
	return 0;
}

func (r *MjRoom) CreateDesk(m *mjproto.Game_CreateRoom) *MjDesk {
	//create 的时候，是否需要通过type 来判断,怎么样创建房间

	userId := m.GetHeader().GetUserId()
	//首先判断是否已经创建了房间...
	oldDesk := GetMjDeskBySession(userId)
	if oldDesk != nil && oldDesk.GetOwner() == userId {
		//如果房间没有开始游戏..则返回老的房间,否则创建新的房间...
		if !oldDesk.IsBegin() {
			return oldDesk
		}
	}

	//log.T("开始创建房间... ")

	//0,先扣费,添加账单
	var createFee int64 = r.CalcCreateFee()
	remain, err := userService.DECRUserDiamond(userId, createFee)
	if err != nil {
		//扣费失败，创建房间失败
		return nil
	}

	err = userService.CreateDiamonDetail(userId, 0, createFee, remain, "创建麻将desk")
	if err != nil {
		//创建订单的时候失败
		return nil
	}

	//1,创建一个房间，并初始化参数
	desk := NewMjDesk()
	*desk.DeskId, _ = db.GetNextSeq(config.DBT_MJ_DESK)
	*desk.RoomId = r.GetRoomId()
	desk.SetStatus(MJDESK_STATUS_CREATED)        //设置为刚刚创建的状态
	desk.InitUsers(m.GetRoomTypeInfo().GetMjRoomType()) //根据房间类型初始化房间玩家数
	*desk.Password = r.RandRoomKey()
	*desk.Owner = userId        //设置房主
	*desk.CreateFee = createFee
	*desk.MjRoomType = int32(m.GetRoomTypeInfo().GetMjRoomType())        // 房间类型，如：血战到底、三人两房、四人两房、德阳麻将、倒倒胡、血流成河
	desk.InitUserCountAndFangCountByType()        //初始化人数和房数
	*desk.BoardsCout = m.GetRoomTypeInfo().GetBoardsCout()        //局数，如：4局（房卡 × 2）、8局（房卡 × 3）
	*desk.CapMax = m.GetRoomTypeInfo().GetCapMax()
	*desk.CardsNum = m.GetRoomTypeInfo().GetCardsNum()
	*desk.Settlement = m.GetRoomTypeInfo().GetSettlement()
	*desk.BaseValue = m.GetRoomTypeInfo().GetBaseValue()
	*desk.ZiMoRadio = m.GetRoomTypeInfo().GetPlayOptions().GetZiMoRadio()
	desk.OthersCheckBox = m.GetRoomTypeInfo().GetPlayOptions().GetOthersCheckBox()
	*desk.HuRadio = m.GetRoomTypeInfo().GetPlayOptions().GetHuRadio()
	*desk.DianGangHuaRadio = m.GetRoomTypeInfo().GetPlayOptions().GetDianGangHuaRadio()
	*desk.MJPaiCursor = 0
	*desk.TotalPlayCount = m.RoomTypeInfo.GetBoardsCout()
	*desk.CurrPlayCount = 0
	*desk.Banker = userId
	*desk.NextBanker = 0        //游戏过程中动态的计算
	desk.CheckCase = nil        //判断case
	*desk.ActiveUser = userId
	*desk.GameNumber = 0
	*desk.ActUser = userId
	*desk.ActType = MJDESK_ACT_TYPE_MOPAI
	//*desk.BeginTime	//游戏开始时间...
	//*desk.EndTime		//游戏结束时间...
	*desk.NInitActionTime = 30        // 游戏操作的时间

	//把创建的desk加入到room中
	r.AddDesk(desk)
	return desk
}

func (r *MjRoom) RandRoomKey() string {
	a := rand.Rand(100000, 1000000)
	roomKey, _ := numUtils.Int2String(a)
	//1,判断roomKey是否已经存在
	if r.IsRoomKeyExist(roomKey) {
		//log.E("房间密钥[%v]已经存在,创建房间失败,重新创建", roomKey)
		return r.RandRoomKey()
	} else {
		//log.T("最终得到的密钥是[%v]", roomKey)
		return roomKey
	}
	return ""
}

//判断roomkey是否已经存在了
func (r *MjRoom) IsRoomKeyExist(roomkey string) bool {
	ret := false
	for i := 0; i < len(r.Desks); i++ {
		d := r.Desks[i]
		if d != nil && d.GetPassword() == roomkey {
			ret = true
			break
		}
	}

	return ret
}

//通过房间号码得到desk
func (r *MjRoom) GetDeskByPassword(key string) *MjDesk {
	//如果找到对应的房间，则返回
	for _, d := range r.GetDesks() {
		if d != nil && d.GetPassword() == key {
			return d
		}
	}

	//如果没有找到，则返回nil
	return nil
}

//通过房间号码得到desk
func (r *MjRoom) GetDeskByDeskId(id int32) *MjDesk {
	//log.T("通过deskId【%v】查询desk", id)
	//如果找到对应的房间，则返回
	for _, d := range r.GetDesks() {
		if d != nil && d.GetDeskId() == id {
			//log.T("通过id[%v]找到desk----d.getDeskId()[%v]", id, d.GetDeskId())
			return d
		}
	}
	//如果没有找到，则返回nil
	return nil
}



//进入房间
//进入的时候，需要判断牌房间的类型...
func (r *MjRoom) EnterRoom(key string, userId uint32, a gate.Agent) (*MjDesk, mjproto.RECONNECT_TYPE, error) {

	var desk *MjDesk
	var reconnect mjproto.RECONNECT_TYPE = mjproto.RECONNECT_TYPE_NORMAL
	//如果是朋友桌,需要通过房间好来找到desk
	if r.IsFriend() {
		desk = r.GetDeskByPassword(key)
		if desk == nil {
			log.T("通过key[%v]没有找到对应的desk", key)
			return nil, mjproto.RECONNECT_TYPE_NORMAL, Error.NewError(consts.ACK_RESULT_FAIL, "房间号输入错误")
		} else {
			var addErr error
			reconnect, addErr = desk.addNewUserFriend(userId, a)
			if addErr != nil {
				//用户加入房间失败...
				log.E("玩家[%v]加入房间失败errMsg[%v]", userId, addErr)
				return nil, reconnect, Error.NewFailError(Error.GetErrorMsg(addErr))
			}
		}

		return desk, reconnect, nil        //朋友桌 数据返回...
	}

	return nil, reconnect, nil

	//如果是锦标赛

	//返回结果...
}

func (r *MjRoom) IsFriend() bool {
	return true
}

func (r *MjRoom) AddDesk(desk *MjDesk) error {
	r.Desks = append(r.Desks, desk)

	//为桌子增加lock ，回复数据的时候，也需要回 lock
	lock.NewDeskLock(desk.GetDeskId())

	//加入之后需要更新数据到redis
	desk.updateRedis()
	AddRunningDeskKey(desk.GetDeskId())
	return nil
}

// room 解散房间...
func (r *MjRoom)DissolveDesk(desk *MjDesk, sendMsg bool) error {
	//清楚数据,1,session相关。2,
	log.T("开始解散desk[%v]...", desk.GetDeskId())
	log.T("开始解散desk[%v]user的session数据...", desk.GetDeskId())
	for _, user := range desk.GetUsers() {
		if user != nil {
			user.ClearAgentGameData()
		}
	}

	log.T("开始删除desk[%v]...", desk.GetDeskId())

	//发送解散房间的广播
	rmErr := r.RmDesk(desk)
	if rmErr != nil {
		log.E("删除房间失败,errmsg[%v]", rmErr)
		return rmErr
	}

	//删除锁
	lock.DelDeskLock(desk.GetDeskId())
	//删除reids
	DelMjDeskRedis(desk)

	//删除房间
	log.T("删除desk[%v]之后，发送删除的广播...", desk.GetDeskId())
	if sendMsg {
		//发送解散房间的广播
		dissolve := newProto.NewGame_AckDissolveDesk()
		*dissolve.DeskId = desk.GetDeskId()
		*dissolve.PassWord = desk.GetPassword()
		*dissolve.UserId = desk.GetOwner()
		desk.BroadCastProto(dissolve)
	}

	return nil

}

func (r *MjRoom) RmDesk(desk *MjDesk) error {
	index := -1
	for i, d := range r.Desks {
		if d != nil && d.GetDeskId() == desk.GetDeskId() {
			index = i
			break
		}
	}

	if index >= 0 {
		r.Desks = append(r.Desks[:index], r.Desks[index + 1:]...)
		return nil
	} else {
		return errors.New("删除失败，没有找到对应的desk")
	}

}

func GetFMJRoom() *MjRoom {
	//暂时返回朋友桌
	return FMJRoomIns
}
//通过用户的session 找到mjroom
func GetMjroomBySession(userId uint32) *MjRoom {
	session := GetSession(userId)
	if session == nil {
		return nil
	}

	session.GetRoomId()
	session.GetDeskId()

	//目前暂时返回一个房间，方便测试 todo
	return FMJRoomIns

}

func GetMjDeskBySession(userId uint32) *MjDesk {
	//得到session
	session := GetSession(userId)
	if session == nil {
		return nil
	}
	log.T("得到用户[%v]的session[%v]", userId, session)

	//得到room
	room := GetMjroomBySession(userId)
	if room == nil {
		return nil
	}

	//返回desk
	return room.GetDeskByDeskId(session.GetDeskId())
}
