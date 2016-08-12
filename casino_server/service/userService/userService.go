package userService

import (
	"casino_server/conf/casinoConf"
	"casino_server/msg/bbprotogo"
	"fmt"
	"github.com/name5566/leaf/db/mongodb"
	"casino_server/common/config"
	"gopkg.in/mgo.v2/bson"
	"casino_server/common/log"
	"casino_server/conf/intCons"
	"casino_server/utils/redis"
	"casino_server/utils/numUtils"
	"strings"
	"casino_server/mode"
	"time"
)


//const-config
var(
	//更新的方式,是只更新到redis 还是redis和mongo都需要更新
	UPDATE_TYPE_ONLY_REDIS	int  = 1
	UPDATE_TYPE_REAIS_MONGO int  = 2
)


/**

判断id是否正确
 */
func CheckUserId(userId uint32) int8 {
	if userId > casinoConf.MAX_USER_ID || userId < casinoConf.MIN_USER_ID {
		return intCons.LOGIN_WAY_QUICK
	} else {
		return intCons.LOGIN_WAY_LOGIN
	}
}

/**
	1,create 一个user
	2,保存mongo
	3,缓存到redis
 */
func NewUserAndSave() (*bbproto.User, error) {
	//1,获取数据库连接和回话
	c, err := mongodb.Dial(casinoConf.DB_IP, casinoConf.DB_PORT)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer c.Close()

	s := c.Ref()
	defer c.UnRef(s)

	//2,创建user获得自增主键
	id, err := c.NextSeq(casinoConf.DB_NAME, casinoConf.DBT_T_USER, casinoConf.DB_ENSURECOUNTER_KEY)
	if err != nil {
		return nil, err
	}
	userId := uint32(id)
	Nickname := config.RandNickname()

	//构造user
	nuser := &mode.T_user{}
	nuser.Mid = bson.NewObjectId()
	nuser.Id = userId
	nuser.NickName = Nickname
	nuser.Coin = int64(10000)

	err = s.DB(casinoConf.DB_NAME).C(casinoConf.DBT_T_USER).Insert(nuser)
	if err != nil {
		log.E("保存用户的时候失败 error【%v】",err.Error())
		return nil,err
	}

	result,_ := Tuser2Ruser(nuser)
	return result, nil

}


func GetRedisUserKey(id uint32) string {
	idStr, _ := numUtils.Uint2String(id)
	return strings.Join([]string{casinoConf.DBT_T_USER,idStr}, "-")
}

/**
	根据用户id得到User的id
	1,首先从redis中查询user信息
	2,如果redis中不存在,则从mongo中查询
	3,如果mongo不存在,返回错误信息,客户端跳转到登陆界面

 */
func GetUserById(id uint32) *bbproto.User {
	//1,首先在 redis中去的数据
	conn := data.Data{}
	conn.Open(casinoConf.REDIS_DB_NAME)
	defer conn.Close()

	key := GetRedisUserKey(id)
	result := &bbproto.User{}
	conn.GetObj(key, result)
	if result == nil || result.GetId() == 0 {
		log.E("redis中没有找到user[%v],需要在mongo中查询,并且缓存在redis中。", id)
		// 获取连接 connection
		c, err := mongodb.Dial(casinoConf.DB_IP, casinoConf.DB_PORT)
		if err != nil {
			result = nil
		}
		defer c.Close()
		s := c.Ref()
		defer c.UnRef(s)

		//从数据库中查询user
		tuser := &mode.T_user{}
		s.DB(casinoConf.DB_NAME).C(casinoConf.DBT_T_USER).Find(bson.M{"id": id}).One(tuser)
		if tuser.Id < casinoConf.MIN_USER_ID {
			log.T("在mongo中没有查询到user[%v].", id)
			result = nil
		}else{
			log.T("在mongo中查询到了user[%v],现在开始缓存",tuser)
			//把从数据获得的结果填充到redis的model中
			result,_ = Tuser2Ruser(tuser)
			if result!=nil {
				SaveUser2Redis(result)
			}
		}
	}

	//判断用户是否存在,如果不存在,则返回空
	if result == nil {
		return nil
	}else{
		result.OninitLoginTurntableState()	//初始化登录转盘之后的奖励
		return result
	}

}

/**
	将用户model保存在redis中
 */
func SaveUser2Redis(u *bbproto.User) {
	conn := data.Data{}
	conn.Open(casinoConf.REDIS_DB_NAME)
	defer conn.Close()
	key := GetRedisUserKey(u.GetId())
	conn.SetObj(key, u)
}

/**
	保存数据到redis和mongo中
 */
func SaveUser2RedisAndMongo(u *bbproto.User){
	SaveUser2Redis(u)
	UpsertUser2Mongo(u)
}


//把redis中的数据刷新到数据库
func FlashUser2Mongo(userId uint32) error{
	u := GetUserById(userId)
	UpsertUser2Mongo(u)
	return nil
}


func UpsertUser2Mongo(u *bbproto.User){
	//得到数据库连接池
	c, err := mongodb.Dial(casinoConf.DB_IP, casinoConf.DB_PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()

	s := c.Ref()
	defer c.UnRef(s)

	//把bbproto.User转化为  model.User
	tuser,_:=Ruser2Tuser(u)	//
	log.T("把user[%v]保存到数据库]",tuser)
	if tuser.Mid == ""{
		s.DB(casinoConf.DB_NAME).C(casinoConf.DBT_T_USER).Insert(tuser)
	}else{
		s.DB(casinoConf.DB_NAME).C(casinoConf.DBT_T_USER).Update(bson.M{"_id": tuser.Mid},tuser)
	}

}
/**
	更新用用户余额的信息
 */
func UpUserBalance(userId uint32, amount int64,utype int) error {

	//1,获得锁
	l := UserLockPools.GetUserLockByUserId(userId)
	l.Lock()
	defer l.Unlock()

	//2,跟新redis中的值
	//由于用户user相关的都会存在redis 中的,所以肯定会更新redis
	user := GetUserById(userId)
	*user.Coin += amount
	SaveUser2Redis(user)	//保存user

	//3,更新mongo 中的值
	if utype == UPDATE_TYPE_REAIS_MONGO {
		UpsertUser2Mongo(user)
	}

	return nil
}


/**
	mongo中User模型转化为 redis中的user模型
 */
func Tuser2Ruser(tu *mode.T_user)(*bbproto.User,error){
	result := &bbproto.User{}
	if tu.Mid.Hex() != "" {
		hesStr := tu.Mid.Hex()
		result.Mid = &hesStr
		//log.T("获得t_user.mid %v",hesStr)
	}

	result.Name = &tu.Name
	result.Id = &tu.Id
	result.NickName = &tu.NickName
	result.Coin = &tu.Coin
	result.Diamond = &tu.Diamond
	return result,nil
}

/**
	redis中的user模型转化为mongdo的User模型
	把Redis_user 转化为mongo_t_user的时候喂自动为其分配objectId,方存储
 */

func Ruser2Tuser(ru *bbproto.User) (*mode.T_user,error){
	result := &mode.T_user{}

	if ru.Mid != nil {
		result.Mid = bson.ObjectIdHex(ru.GetMid())
	}else{
		result.Mid = bson.NewObjectId()
	}

	result.Id = ru.GetId()
	result.Name = ru.GetName()
	result.NickName = ru.GetNickName()
	result.Coin = ru.GetCoin()
	return result,nil
}

/**
	判断用户id是否合法,todo 这里是否在数据库中判断?由于之后的查询会在数据库中查询,所以这里可以先不用判断
	if userId  > casinoConf.MAX_USER_ID || userId < casinoConf.MIN_USER_ID {

 */
func CheckUserIdRightful(userId uint32) bool{
	u := GetUserById(userId)
	if u == nil {
		return false
	}else{
		return  true
	}
}

//增加用户的coin
func IncreasUserCoin(userId uint32,coin int64) error{
	DecreaseUserCoin(userId,(0-coin))
	return nil
}

//减少用户的余额
func DecreaseUserCoin(userId uint32,coin int64) error{
	//lock := UserLockPools.GetUserLockByUserId(userId)
	//lock.Lock()
	//defer lock.Unlock()

	//开是减少用户的金币
	user := GetUserById(userId)
	*user.Coin -= coin
	SaveUser2Redis(user)
	return nil
}


//更新用户的钻石之后,在放回用户当前的余额,更新用户钻石需要同事更新redis和mongo的数据
func UpdateUserDiamond(userId uint32,diamond int64) int64{
	//1,获取锁
	//lock := UserLockPools.GetUserLockByUserId(userId)
	//lock.Lock()
	//defer lock.Unlock()

	//2,修改用户redis和mongo中的数据
	user := GetUserById(userId)
	if user == nil {
		return -1
	}

	*user.Diamond += diamond
	SaveUser2RedisAndMongo(user)

	//3,返回数据
	return user.GetDiamond()
}

func CreateDiamonDetail(userId uint32,detailsType int32,diamond int64,remainDiamond int64,memo string) error{

	//1,获取数据库连接
	c, err := mongodb.Dial(casinoConf.DB_IP, casinoConf.DB_PORT)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer c.Close()

	s := c.Ref()
	defer c.UnRef(s)

	//2,活的交易记录自增主键
	id, err := c.NextSeq(casinoConf.DB_NAME, casinoConf.DBT_T_USER_DIAMOND_DETAILS, casinoConf.DB_ENSURECOUNTER_KEY)
	if err != nil {
		return err
	}

	//构造交易记录
	detail := &mode.T_user_diamond_details{}
	detail.Id = uint32(id)
	detail.UserId = userId
	detail.Diamond = diamond
	detail.ReaminDiamond = remainDiamond
	detail.DetailsType = detailsType
	detail.DetailTime = time.Now()
	detail.Memo = memo

	err = s.DB(casinoConf.DB_NAME).C(casinoConf.DBT_T_USER_DIAMOND_DETAILS).Insert(detail)
	if err != nil {
		log.E("保存用户交易记录的时候失败 error【%v】",err.Error())
		return err
	}
	return  nil
}


