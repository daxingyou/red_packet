package doudizhu

import (
	"sync"
	"casino_server/common/log"
	"casino_server/common/Error"
)

var (
	DDZUSER_STATUS_READY int32 = 2        //已经准备

)

var (
	DDZUSER_QIANGDIZHU_STATUS_NULL int32 = 0 //没操作
	DDZUSER_QIANGDIZHU_STATUS_QIANG int32 = 1 //抢地主
	DDZUSER_QIANGDIZHU_STATUS_PASS int32 = 2 //不叫
)

type DdzUser struct {
	sync.Mutex
	*PDdzUser
}

//清楚session
func (u *DdzUser)ClearAgentGameData() {

}

func (u *DdzUser) SetOnline() error {
	*u.IsBreak = false
	return nil
}

func (u *DdzUser) UpdateSession() {
	//更新session的信息
}

//设置庄太
func (u *DdzUser) SetStatus(s int32) {
	*u.Status = s
}

func (u *DdzUser) SetQiangDiZhuStatus(s int32) {
	*u.QiangDiZhuStatus = s
}

func (u *DdzUser) IsReady() bool {
	return u.GetStatus() == DDZUSER_STATUS_READY
}

func (u *DdzUser) IsNotReady() bool {
	return !u.IsReady()
}

//是否抢地主
func (u *DdzUser) IsQiangDiZhu() bool {
	return u.GetQiangDiZhuStatus() == DDZUSER_QIANGDIZHU_STATUS_QIANG
}

//是否不叫
func (u *DdzUser) IsBuJiao() bool {
	return u.GetQiangDiZhuStatus() == DDZUSER_QIANGDIZHU_STATUS_PASS
}

func (u *DdzUser) DelHandlPai(pais *PPokerPai) {
	index := -1
	for i, pai := range u.GameData.HandPokers {
		if pai != nil && pai.GetId() == pais.GetId() {
			index = i
			break
		}
	}
	if index > -1 {
		u.GameData.HandPokers = append(u.GameData.HandPokers[:index], u.GameData.HandPokers[index + 1:]...)
		return nil

	} else {
		log.E("服务器错误：删除手牌的时候出错，没有找到对应的手牌[%v]", pais)
		return Error.NewError(-1, "删除手牌时出错，没有找到对应的手牌...")
	}

}

//增加出牌
func (u *DdzUser) DOPoutPokerPais(out *POutPokerPais) error {
	//1，增加出的牌
	u.OutPaiList = append(u.OutPaiList, out)

	//2，删除手牌
	for _, p := range out.PokerPais {
		u.DelHandlPai(p)
	}

	return nil
}

//得到玩家手牌的张数
func (u *DdzUser) GetHandPaiCount() int32 {
	return len(u.GameData.HandPokers)
}

func (u *DdzUser) beginInit() error{
	return nil

}