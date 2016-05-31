package internal

import (
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/gate"
	"reflect"
	"casino_server/msg/bbproto"
	"casino_server/msg"
)

func init() {
	// 向当前模块（game 模块）注册 Hello 消息的消息处理函数 handleHello
	handler(&msg.Hello{}, handleHello)

	handler(&bbproto.TestP1{},handleTestP1)
	handler(&bbproto.N{},handleProtHello)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleHello(args []interface{}) {
	// 收到的 Hello 消息
	m := args[0].(*msg.Hello)
	// 消息的发送者
	a := args[1].(gate.Agent)

	// 输出收到的消息的内容
	log.Debug("hello %v", m.Name)

	// 给发送者回应一个 Hello 消息

	a.WriteMsg(&msg.Hello{
		Name: "bbgogogogogogo",
	})

}

func handleProtHello(args []interface{}){
	log.Debug("进入handleProtHello()")
	// 收到的 Hello 消息
	m := args[0].(*bbproto.N)
	// 消息的发送者
	a := args[1].(gate.Agent)

	// 输出收到的消息的内容
	log.Debug("接收到的name %v", *m.Name)
	 //给发送者回应一个 Hello 消息
	var data bbproto.N
	//var n string = "a"+ time.Now().String()
	var n string = "hi leaf"
	data.Name = &n
	a.WriteMsg(&data)
}

func handleTestP1(args[]interface{}){
	log.Debug("进入handleTestP1()")
	// 收到的 Hello 消息
	m := args[0].(*bbproto.TestP1)
	// 消息的发送者
	a := args[1].(gate.Agent)

	// 输出收到的消息的内容
	log.Debug("接收到的name %v", *m.Name2)
	//给发送者回应一个 Hello 消息
	var data bbproto.TestP1
	var n string = "hi leaf testp2"
	data.Name2 = &n
	a.WriteMsg(&data)
}