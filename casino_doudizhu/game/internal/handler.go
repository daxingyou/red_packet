package internal

import (
	"reflect"
	"casino_doudizhu/msg/protogo"
	"github.com/name5566/leaf/gate"
	"casino_doudizhu/service/DdzService"
)

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)

}

func init() {

	handler(&ddzproto.DdzCreateRoom{}, handlerCreateDesk)    //创建房间
	handler(&ddzproto.DdzEnterRoom{}, HandlerEnterRoom)    //进入房间
	handler(&ddzproto.DdzReady{}, handlerReady)    //准备
	handler(&ddzproto.DdzJiaoDiZhu{}, handlerJiaoDiZhu)    //
	handler(&ddzproto.DdzRobDiZhu{}, handlerQiangDiZhu)    //
	handler(&ddzproto.DdzDouble{}, handlerJiaBei)    //
	handler(&ddzproto.DdzShowHandPokers{}, handlerShowHandPokers)    //
	handler(&ddzproto.DdzMenuZhua{}, handlerMenuZhua)    //
	handler(&ddzproto.DdzSeeCards{}, handlerSeeCards)    //
	handler(&ddzproto.DdzPull{}, handlePull)    //
	handler(&ddzproto.DdzOutCards{}, handlerChuPai)    //
	handler(&ddzproto.DdzActGuo{}, handlerActGuo)    //
	handler(&ddzproto.DdzDissolveDesk{}, handlerDissolveDesk)    //
	handler(&ddzproto.DdzLeaveDesk{}, handlerLeaveDesk)    //
	handler(&ddzproto.DdzMessage{}, handlerMessage)    //
	handler(&ddzproto.DdzGameRecord{}, handlerGameRecord)    //
}


//创建房间
func handlerCreateDesk(args []interface{}) {
	m := args[0].(*ddzproto.DdzCreateRoom)
	a := args[1].(gate.Agent)
	DdzService.HandlerCreateDesk(m.GetHeader().GetUserId(), a)
}

//进入房间
func HandlerEnterRoom(args []interface{}) {
	m := args[0].(*ddzproto.DdzEnterRoom)
	a := args[1].(gate.Agent)

	var enterType int32 = 0
	DdzService.HandlerEnterDesk(m.GetUserId(), m.GetPassWord(), enterType, a)

}

//准备
func handlerReady(args []interface{}) {
	m := args[0].(*ddzproto.DdzReady)
	DdzService.HandlerFDdzReady(m.GetUserId())
}

func handlerJiaoDiZhu(args []interface{}) {
	m := args[0].(*ddzproto.DdzJiaoDiZhu)
	DdzService.HandlerJiaoDiZhu(m.GetUserId())
}

//明牌
func handlerShowHandPokers(args []interface{}) {
	m := args[0].(*ddzproto.DdzShowHandPokers)
	DdzService.HandlerShowHandPokers(m.GetUserId())
}

//闷牌
func handlerMenuZhua(args []interface{}) {
	m := args[0].(*ddzproto.DdzMenuZhua)
	DdzService.HandlerMenuZhua(m.GetUserId())

}
func handlerSeeCards(args []interface{}) {
	m := args[0].(*ddzproto.DdzSeeCards)
	DdzService.HandlerSeeCards(m.GetUserId())
}

func handlePull(args []interface{}) {
	m := args[0].(*ddzproto.DdzPull)
	DdzService.HandlePull(m.GetUserId())
}

//抢地主
func handlerQiangDiZhu(args []interface{}) {
	m := args[0].(*ddzproto.DdzRobDiZhu)
	DdzService.HandlerQiangDiZhu(m.GetUserId())

}

func handlerActGuo(args []interface{}) {
	m := args[0].(*ddzproto.DdzActGuo)
	DdzService.HandlerActPass(m.GetUserId())
}

func handlerDissolveDesk(args []interface{}) {
	m := args[0].(*ddzproto.DdzDissolveDesk)
	DdzService.HandlerDissolveDesk(m.GetUserId())
}

func handlerLeaveDesk(args []interface{}) {
	m := args[0].(*ddzproto.DdzLeaveDesk)
	DdzService.HandlerLeaveDesk(m.GetHeader().GetUserId())
}

func handlerMessage(args []interface{}) {
	//m := args[0].(*ddzproto.Game_Message)
	DdzService.HandlerMessage()

}
func handlerGameRecord(args []interface{}) {
	//m := args[0].(*ddzproto.Game_GameRecord)
	DdzService.HandlerGameRecord()

}

//加倍
func handlerJiaBei(args []interface{}) {
	//m := args[0].(*ddzproto.Game_Double)
	DdzService.HandlerJiaBei()
}

//出牌
func handlerChuPai(args []interface{}) {
	m := args[0].(*ddzproto.DdzOutCards)
	DdzService.HandlerActOut(m.GetHeader().GetUserId())
}







