package internal

import (
	"github.com/name5566/leaf/module"
	"casino_doudizhu/base"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC = skeleton.ChanRPCServer
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton

	//回复数据
	//majiang.RecoverFMJ()        //回复麻将朋友桌的数据

}

func (m *Module) OnDestroy() {

}