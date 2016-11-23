package doudizhu

//备份用
func NewPDdzbak() *PDdzbak {
	ret := new(PDdzbak)
	return ret
}

//Desk
func NewPDdzDesk() *PDdzDesk {
	ret := new(PDdzDesk)
	ret.DeskId = new(int32)
	ret.Key = new(string)
	ret.UserCountLimit = new(int32)
	ret.Owner = new(uint32)
	ret.DizhuPaiUser = new(uint32)
	ret.DiZhuUserId = new(uint32)
	ret.ActiveUserId = new(uint32)
	ret.BaseValue = new(int64)
	ret.QingDizhuValue = new(int64)
	ret.WinValue = new(int64)
	ret.DdzType = new(int32)
	ret.RoomType = new(int32)
	ret.BoardsCount = new(int32)
	ret.CapMax = new(int64)
	ret.IsJiaoFen = new(bool)
	return ret
}

func NewPDdzUser() *PDdzUser {
	ret := new(PDdzUser)
	return ret
}


//New一个Desk
func NewDdzDesk() *DdzDesk {
	desk := new(DdzDesk)
	desk.PDdzDesk = NewPDdzDesk()
	desk.UserCountLimit = new(int32)
	desk.RoomId = new(int32)
	return desk
}

func NewDdzUser() *DdzUser {
	user := new(DdzUser)
	user.PDdzUser = NewPDdzUser()
	user.UserId = new(uint32)
	user.DeskId = new(int32)
	user.RoomId = new(int32)
	user.Status = new(int32)
	user.IsBreak = new(bool)
	user.IsLeave = new(bool)
	user.QiangDiZhuStatus = new(int32)
	return user
}

func NewPPokerPai() *PPokerPai {
	pai := new(PPokerPai)
	pai.Des = new(string)
	pai.Flower = new(int32)
	pai.Id = new(int32)
	pai.Name = new(string)
	pai.Value = new(int32)
	return pai
}

func NewPDdzBillBean() *PDdzBillBean {
	b := new(PDdzBillBean)
	b.Coin = new(int64)
	b.WinUser = new(uint32)
	b.LoseUser = new(uint32)
	b.Desc = new(string)
	return b
}

func NewPDdzSession() *PDdzSession {
	ret := new(PDdzSession)
	ret.UserId = new(uint32)
	ret.DeskId = new(int32)
	ret.RoomId = new(int32)
	ret.GameStatus = new(int32)
	return ret
}