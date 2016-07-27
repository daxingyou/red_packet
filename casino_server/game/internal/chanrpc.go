package internal

import (
	"github.com/name5566/leaf/gate"
	"casino_server/common/log"
	"casino_server/gamedata"
	"casino_server/service/room"
)



func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
}

func rpcNewAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	_ = a
}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)

	log.Normal("调用了.....game.chanrpc.rpcCloseAgent")
	//对数据做异常退出时的保存工作
	//1,保存游戏数据
	//2,删除连接中管理的agent

	//用户掉线的处理--------------------------------------begin-----------------------------------
	//,如果UserId是10006的话,连接断开的时候,desk 删除user
	agentData := a.UserData()
	if agentData == nil {
		log.T("通过agent[%v]取出来的userData 是nil",a)
	}else{
		//用户数据还在,设置用户为掉线的状态
		userData := agentData.(*gamedata.AgentUserData)
		log.T("用户[%v]现在掉线了,现在设置用户为掉线的状态",userData.UserId)
		desk := room.ThGameRoomIns.GetDeskByUserId(userData.UserId)
		desk.SetOfflineStatus(userData.UserId)
	}
	//用户掉线的处理--------------------------------------end--------------------------------------

}
