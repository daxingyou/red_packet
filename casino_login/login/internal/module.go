package internal

import (
	"github.com/name5566/leaf/module"
	"fmt"
	"casino_login/base"
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
}

func (m *Module) OnDestroy() {
	fmt.Println("")

}
