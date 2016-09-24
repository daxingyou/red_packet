package main

import (
	"github.com/name5566/leaf"
	lconf "github.com/name5566/leaf/conf"
	"casino_majiang/conf"
	"casino_majiang/game"
	"casino_majiang/gate"
	"casino_majiang/login"
)

func init() {
	//系统启动时候的初始化操作...
	conf.InitSystem()
}

func main() {
	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath

	leaf.Run(
		game.Module,
		gate.Module,
		login.Module,
	)
}
