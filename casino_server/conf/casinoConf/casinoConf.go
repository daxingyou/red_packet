package casinoConf


//------------------------------mongo-数据库相关--------------------------
const (
	DB_NAME = "test"                //数据库名字
	DB_ENSURECOUNTER_KEY = "id"                        //自增键

	DBT_T_USER = "t_user"                //user表名字
	DBT_T_TEST = "t_test"                //user表名字
	DBT_T_SUB2 = "t_test_sub2"                //user表名字
	DBT_T_ZJH_ROUND = "t_zjh_round"                //每局炸扎金花的数据
	DBT_T_BONUS_TURNTABLE = "t_bonus_turntable"        //转盘奖励的表
	DBT_T_TH_GAMENUMBER_SEQ = "t_th_gamenumber_seq"                //德州扑克
	DBT_T_TH_RECORD = "t_th_record"                //德州扑克
	DBT_T_CS_TH_RECORD = "t_cs_th_rercord"        //锦标赛的记录
	DBT_T_CS_TH_DESK_RECORD = "t_cs_th_desk_record"        //德州扑克
	DBT_T_TH_DESK_RECORD = "t_th_desk_record"        //德州扑克
	DBT_T_TH_DESK = "t_th_desk"                //德州桌子
	DBT_T_TH_NOTICE = "t_th_notice"
	DBT_T_USER_DIAMOND_DETAILS = "t_user_diamond_details"        //用户钻石的交易记录

	//数据库的常用设置
	MIN_USER_ID = 10000                //USER_ID 起始
	MAX_USER_ID = 100000000                //USER_ID 限制


	//------------------------------redis-数据库相关--------------------------

	REDIS_IP = "127.0.0.1"
	REDIS_PORT = "6379"
	REDIS_DB_NAME = "test"

	//-------------------------------游戏设置---------------------------------
	SWITCH_ZJH bool = false
)
