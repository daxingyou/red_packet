package gate

import (
	"casino_server/msg"
	"casino_server/msg/bbprotogo"
	"casino_server/game"
	"casino_server/login"
	"casino_server/bonus"
)

func init() {
	//指定protobuf格式的路由

	//水果机
	msg.ProtoProcessor.SetRouter(&bbproto.GetIntoRoom{}, game.ChanRPC)
	msg.ProtoProcessor.SetRouter(&bbproto.Shuiguoji{}, game.ChanRPC)                //水果机
	msg.ProtoProcessor.SetRouter(&bbproto.ShuiguojiHilomp{}, game.ChanRPC)        //水果机比大小

	//扎金花
	msg.ProtoProcessor.SetRouter(&bbproto.ZjhRoom{}, game.ChanRPC)                //进入扎金花的房间
	msg.ProtoProcessor.SetRouter(&bbproto.ZjhLottery{}, game.ChanRPC)                //进入扎金花的房间
	msg.ProtoProcessor.SetRouter(&bbproto.ZjhMsg{}, game.ChanRPC)                //进入扎金花的房间
	msg.ProtoProcessor.SetRouter(&bbproto.ZjhBet{}, game.ChanRPC)                //进入扎金花的房间
	msg.ProtoProcessor.SetRouter(&bbproto.ZjhReqSeat{}, game.ChanRPC)                //进入扎金花的房间
	msg.ProtoProcessor.SetRouter(&bbproto.ZjhQueryNoSeatUser{}, game.ChanRPC)                //进入扎金花的房间

	//奖励
	msg.ProtoProcessor.SetRouter(&bbproto.LoginTurntableBonus{}, bonus.ChanRPC)        //登陆奖励
	msg.ProtoProcessor.SetRouter(&bbproto.LoginSignInBonus{}, bonus.ChanRPC)                //连续签到的奖励

	//德州扑克
	msg.ProtoProcessor.SetRouter(&bbproto.ThRoom{}, game.ChanRPC)                //处理德州扑克房间
	msg.ProtoProcessor.SetRouter(&bbproto.THBet{}, game.ChanRPC)                //处理德州扑克押注

	///联众游戏
	msg.ProtoProcessor.SetRouter(&bbproto.REQQuickConn{}, login.ChanRPC)
	msg.ProtoProcessor.SetRouter(&bbproto.NullMsg{}, login.ChanRPC)
	msg.ProtoProcessor.SetRouter(&bbproto.MatchList_ReqMobileMatchList{}, game.ChanRPC)        //快速开始
	msg.ProtoProcessor.SetRouter(&bbproto.Game_LoginGame{}, game.ChanRPC)        //登陆游戏
	msg.ProtoProcessor.SetRouter(&bbproto.Game_Ready{}, game.ChanRPC)        //准备游戏
	msg.ProtoProcessor.SetRouter(&bbproto.Game_Begin{}, game.ChanRPC)        //开始游戏
	msg.ProtoProcessor.SetRouter(&bbproto.Game_EnterMatch{}, game.ChanRPC)        //进入游戏
	msg.ProtoProcessor.SetRouter(&bbproto.Game_FollowBet{}, game.ChanRPC)        //押注的请求
	msg.ProtoProcessor.SetRouter(&bbproto.Game_RaiseBet{}, game.ChanRPC)        //加注的请求
	msg.ProtoProcessor.SetRouter(&bbproto.Game_FoldBet{}, game.ChanRPC)        //弃牌的请求
	msg.ProtoProcessor.SetRouter(&bbproto.Game_CheckBet{}, game.ChanRPC)        //让牌的请求
	msg.ProtoProcessor.SetRouter(&bbproto.Game_Message{}, game.ChanRPC)        //发送信息
	msg.ProtoProcessor.SetRouter(&bbproto.Game_LeaveDesk{}, game.ChanRPC)        //离开游戏房间

	//公告
	msg.ProtoProcessor.SetRouter(&bbproto.Game_Notice{}, login.ChanRPC)        //登陆的时候查看公告信息
	msg.ProtoProcessor.SetRouter(&bbproto.Game_CreateDesk{}, game.ChanRPC)        //创建房间
	msg.ProtoProcessor.SetRouter(&bbproto.Game_DissolveDesk{}, game.ChanRPC)        //解散房间
	msg.ProtoProcessor.SetRouter(&bbproto.Game_GameRecord{}, game.ChanRPC)        //查询战绩的接口


	msg.ProtoProcessor.SetRouter(&bbproto.Game_TounamentBlind{}, game.ChanRPC)
	msg.ProtoProcessor.SetRouter(&bbproto.Game_TounamentRewards{}, game.ChanRPC)
	msg.ProtoProcessor.SetRouter(&bbproto.Game_TounamentRank{}, game.ChanRPC)
	msg.ProtoProcessor.SetRouter(&bbproto.Game_TounamentSummary{}, game.ChanRPC)

	msg.ProtoProcessor.SetRouter(&bbproto.Game_MatchList{}, game.ChanRPC)        //锦标赛 列表
	msg.ProtoProcessor.SetRouter(&bbproto.Game_Rebuy{}, game.ChanRPC)                        //重新买入的协议
	msg.ProtoProcessor.SetRouter(&bbproto.Game_NotRebuy{}, game.ChanRPC)        //不重购
	msg.ProtoProcessor.SetRouter(&bbproto.Game_Login{}, game.ChanRPC)                        //登陆大厅的协议
	msg.ProtoProcessor.SetRouter(&bbproto.Game_Feedback{}, game.ChanRPC)                //游戏回馈
}