package gate

import (
)
import (
	"casino_server/msg"
	"casino_server/msg/bbproto"
	"casino_server/game"
	"casino_server/login"
	"casino_server/system"
)

func init() {
	//指定protobuf格式的路由
	msg.PortoProcessor.SetRouter(&bbproto.TestP1{},game.ChanRPC)
	msg.PortoProcessor.SetRouter(&bbproto.Reg{},login.ChanRPC)
	msg.PortoProcessor.SetRouter(&bbproto.ReqAuthUser{},login.ChanRPC)
	msg.PortoProcessor.SetRouter(&bbproto.HeatBeat{},system.ChanRPC)
	msg.PortoProcessor.SetRouter(&bbproto.GetIntoRoom{},game.ChanRPC)
}