package mongodb

import (
	"testing"
	"net"
	"casino_server/msg/bbprotogo"
	"casino_server/utils/test"
)

func TestLogin(t *testing.T){
	//login2()
login1()
	for ; ;  {
		
	}
}


//指定id登陆
func login1(){
	conn, err := net.Dial(TCP, url)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	var id uint16 = 2
	var uId uint32 = 10005
	var h *bbproto.ProtoHeader = &bbproto.ProtoHeader{
		UserId:&uId,
	}
	var data bbproto.ReqAuthUser
	data.Header = h
	m := test.AssembleData(id,&data)
	conn.Write(m)
}

//游客登陆

func login2(){
	conn, err := net.Dial(TCP, url)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	var id uint16 = 2
	var h *bbproto.ProtoHeader = &bbproto.ProtoHeader{}
	var data bbproto.ReqAuthUser
	var uuidStr = "8029409jowejfiosjljfl"
	data.Header = h
	data.Uuid = &uuidStr
	m := test.AssembleData(id,&data)
	conn.Write(m)
}
