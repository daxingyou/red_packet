package majiang

import (
	"casino_majiang/msg/funcsInit"
	"casino_majiang/conf/log"
	"casino_server/utils"
	"casino_server/utils/numUtils"
)


//普通的麻将房间...

//room的结构定义在proto中
//type MjRoom struct {
//	roomType int32
//
//}


func init() {
	FMJRoomIns = newProto.NewMjRoom()
	FMJRoomIns.OnInit()
}

var FMJRoomIns *MjRoom

//初始化
func (r *MjRoom) OnInit() {
	log.T("初始化麻将的room")

}

func (r *MjRoom) CreateDesk() *MjDesk {
	//create 的时候，是否需要通过type 来判断,怎么样创建房间

	//1,创建一个房间，并初始化参数
	desk := newProto.NewMjDesk()
	*desk.Password = r.RandRoomKey()

	//把创建的desk加入到room中
	r.AddDesk(desk)
	return desk
}

func (r *MjRoom) RandRoomKey() string {
	a := utils.Rand(100000, 1000000)
	roomKey, _ := numUtils.Int2String(a)
	//1,判断roomKey是否已经存在
	if r.IsRoomKeyExist(roomKey) {
		log.E("房间密钥[%v]已经存在,创建房间失败,重新创建", roomKey)
		return r.RandRoomKey()
	} else {
		log.T("最终得到的密钥是[%v]", roomKey)
		return roomKey
	}
	return nil
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

func (r *MjRoom) AddDesk(desk *MjDesk) error {
	r.Desks = append(r.Desks, desk)
	return nil
}


