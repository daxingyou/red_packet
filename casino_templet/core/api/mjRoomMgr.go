package api

//麻将的房间管理器
type MjRoomMgr interface {
	GetDesk() MjDesk
	GetRoom(int32, int32) MjRoom
	OnInit() error
	//SetSkeleton(*module.Skeleton)
}
