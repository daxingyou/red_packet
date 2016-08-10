package mongodb

import (
	"net"
	"casino_server/msg/bbprotogo"
	"fmt"
	"casino_server/utils/test"
	"testing"
)

func TestOg(t *testing.T) {
	 //game_EnterMatch(10007)
	 //game_EnterMatch(10008)
	 //game_EnterMatch(10009)
	 //game_EnterMatch(10010)
	 //game_EnterMatch(10011)
	//ogbet(0,20)
	//ogbet(1,20)
	//ogbet(2,20)
	//ogbet(3,20)

	//ogRaise(0,20)
	//ogRaise(1,20)
	//ogRaise(2,20)
	//ogRaise(3,20)

	//rEQQuickConn(10006)

	createDesk(10084,"roomkkkk")
	for ; ;  {
	}
}

func game_EnterMatch(userId uint32){
	conn, err := net.Dial(TCP, url)
	if err != nil {
		panic(err)
	}
	ide2 := int32(31)

	fmt.Println("proto 得到的id ", ide2)
	data2 := &bbproto.Game_EnterMatch{}
	data2.UserId = &userId
	m2 := test.AssembleDataNomd5(uint16(ide2), data2)
	conn.Write(m2)
}

//押注
func ogbet(seatId int32,coin int64){
	var tableId int32 = 0
	conn, err := net.Dial(TCP, url)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	ide2 := int32(bbproto.EProtoId_PID_GAME_FOLLOWBET)

	fmt.Println("proto 得到的id ", ide2)
	followData := &bbproto.Game_FollowBet{}
	followData.Tableid = &tableId
	followData.Seat = &seatId
	m2 := test.AssembleDataNomd5(uint16(ide2), followData)
	conn.Write(m2)
	//_ = test.Read(conn).(*bbproto.Game_AckFollowBet)
}



//用户登陆
func rEQQuickConn(userId uint32){
	conn, err := net.Dial(TCP, url)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	ide2 := int32(bbproto.EProtoId_REQQUICKCONN)

	fmt.Println("proto 得到的id ", ide2)
	followData := &bbproto.REQQuickConn{}
	followData.UserId = &userId
	m2 := test.AssembleDataNomd5(uint16(ide2), followData)
	conn.Write(m2)
	_ = test.Read(conn).(*bbproto.ACKQuickConn)

}


func ogRaise(seatId int32,coin int64){
	var tableId int32 = 0
	conn, err := net.Dial(TCP, url)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	ide2 := int32(bbproto.EProtoId_PID_GAME_RAISEBET)

	fmt.Println("proto 得到的id ", ide2)
	raiseData := &bbproto.Game_RaiseBet{}
	raiseData.Tableid = &tableId
	raiseData.Seat = &seatId
	raiseData.Coin = &coin
	m2 := test.AssembleDataNomd5(uint16(ide2), raiseData)
	conn.Write(m2)
	//_ = test.Read(conn).(*bbproto.Game_AckRaiseBet)
}


//创建一个房间
func createDesk(userId uint32,roomKey string){
	pid := int32(bbproto.EProtoId_PID_GAME_GAME_CREATEDESK)
	reqData := &bbproto.Game_CreateDesk{}
	reqData.BigBlind = new(int64)
	reqData.SmallBlind = new(int64)
	reqData.InitCount = new(int32)
	reqData.InitCoin = new(int64)
	reqData.UserId = new(uint32)
	reqData.Password = new(string)

	*reqData.BigBlind = 30
	*reqData.SmallBlind = 15
	*reqData.InitCount = 20
	*reqData.InitCoin = 1000
	*reqData.UserId = userId
	*reqData.Password = roomKey

	conn, err := net.Dial(TCP, url)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	m2 := test.AssembleDataNomd5(uint16(pid), reqData)
	conn.Write(m2)

	//读取放回的信息
	test.Read(conn)

}