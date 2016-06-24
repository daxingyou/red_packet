package fruitService

import (
	"casino_server/msg/bbproto"
	"casino_server/utils"
	"casino_server/common/log"
	"reflect"
)

/**

水果机的 index
 */

const (
	S_ORANGE_1  		= ""
	S_BELL_1			= ""
	S_BAR_LITTLE		= ""
	S_BAR			= ""
	S_APPLE_1			= ""
	S_APPLE_LITTLE		= ""
	S_MANGO_1			= ""
	S_WATERMELON		= ""
	S_WATERMELON_LITTLE		= ""
	S_LUCK_1			= ""
	S_APPLE_2			= ""
	S_ORANGE_LITTLE		= ""
	S_ORANGE_2			= ""
	S_BELL_2			= ""
	S_77_LITTLE			= ""
	S_77			= ""
	S_APPLE_3			= ""
	S_MANGO_LITTLE		= ""
	S_MANGO_2			= ""
	S_STAR			= ""
	S_STAR_LITTLE		= ""
	S_LUCK_2			= ""
	S_APPLE_4			= ""
	S_BELL_LITTLE		= ""

)



const(
	INDEX_ORANGE_1 		= 	0
	INDEX_BELL_1		=	1
	INDEX_BAR_LITTLE	=	2
	INDEX_BAR		=	3
	INDEX_APPLE_1		=	4
	INDEX_APPLE_LITTLE	=	5
	INDEX_MANGO_1		=	6
	INDEX_WATERMELON	=	7
	INDEX_WATERMELON_LITTLE =	8
	INDEX_LUCK_1		=	9
	INDEX_APPLE_2		=	10
	INDEX_ORANGE_LITTLE	=	11
	INDEX_ORANGE_2		=	12
	INDEX_BELL_2		=	13
	INDEX_77_LITTLE		=	14
	INDEX_77		=	15
	INDEX_APPLE_3		=	16
	INDEX_MANGO_LITTLE	=	17
	INDEX_MANGO_2		=	18
	INDEX_STAR		=	19
	INDEX_STAR_LITTLE	=	20
	INDEX_LUCK_2		=	21
	INDEX_APPLE_4		=	22
	INDEX_BELL_LITTLE 	=	23
)




/**
水果几的类型,是押注还是赌大小
 */
var SGJ_TYPE_Bet int32 = 1
var SGJ_TYPE_Hilomp int32 = 2


//

//奖励的类型
var SGJ_WIN_TYPE_LUCK0 int32 = 1
var SGJ_WIN_TYPE_LUCK1 int32 = 2
var SGJ_WIN_TYPE_LUCK2 int32 = 3
var SGJ_WIN_TYPE_LUCK3 int32 = 4
var SGJ_WIN_TYPE_DASIXI int32 = 5
var SGJ_WIN_TYPE_DASANYUAN int32 = 6
var SGJ_WIN_TYPE_XIAOSANYUAN int32 = 7
var SGJ_WIN_TYPE_RAND2 int32 = 8
var SGJ_WIN_TYPE_RAND3 int32 = 9
var SGJ_WIN_TYPE_ZONGHENGSIHAI int32 = 10

//奖励的个数
var NUMBER_INT_1 int32 = 1

/**
配置每个水果对应的积分
 */
var SGJV SGJValue = SGJValue{
	Apple                :                        5,
	AppleLittle        :                        3,
	Orange                :                        10,
	OrangeLittle        :                        3,
	Mango                :                        15,
	MangoLittle        :                        3,
	Bell                :                        20,
	BellLittle        :                        3,
	Watermelon        :                        20,
	WatermelonLittle:                        3,
	Star                :                        30,
	StarLittle        :                        3,
	S77                :                        40,
	S77Little        :                        3,
	Bar                :                        120,
	BarLittle        :                        50,
	Litter                :                        3,
	Lucky                :                        0,
}


/**
概率设计
所有的概率都需要以此结构体来做配置
 */
type ShuiGuoPro struct {
	ORANGE_1 	  int32
	BELL_1	          int32
	BAR_LITTLE        int32
	BAR	          int32
	APPLE_1	          int32
	APPLE_LITTLE      int32
	MANGO_1	          int32
	WATERMELON        int32
	WATERMELON_LITTLE int32
	LUCK_1	          int32
	APPLE_2	          int32
	ORANGE_LITTLE     int32
	ORANGE_2	  int32
	BELL_2	          int32
	S77_LITTLE	  int32
	S77	          int32
	APPLE_3	          int32
	MANGO_LITTLE      int32
	MANGO_2	          int32
	STAR	          int32
	STAR_LITTLE       int32
	LUCK_2	          int32
	APPLE_4	          int32
	BELL_LITTLE	  int32

	DASIXI            int32 //大四喜
	DASANYUAN         int32 //大三元
	XIAOSANYUAN       int32 //小三元
	RAND1		  int32 //随机赠送一炮
	RAND2             int32 //随机赠送两炮
	RAND3             int32 //随机赠送三炮
	RAND4             int32 //随机赠送四炮
	ZONGHENGSIHAI     int32 //纵横四海
	WU		  int32	//没有跑到的概率(注意:只有先跑到luck之后才有这个概率)
}

func ( p *ShuiGuoPro) GetMax() int32{
	v := reflect.ValueOf(*p)
	count := v.NumField()
	var result int32 = int32(v.Field(0).Interface().(int32))
	for i:=0;i<count ;i++  {
		f :=v.Field(i)
		if f.Interface().(int32) > result {
			result = f.Interface().(int32)
		}
	}
	log.T("概率最大的值,",result)
	return result
}





/**
跑中luck 时,其他的概率
 */
var SGJLuckP ShuiGuoPro = ShuiGuoPro{
	ORANGE_1 	             :        5,
	BELL_1	             :        6,
	BAR_LITTLE             :        7,
	BAR	             :        8,
	APPLE_1	          :        15,
	APPLE_LITTLE             :        20,
	MANGO_1	           :        25,
	WATERMELON            :        30,
	WATERMELON_LITTLE          :        35,
	LUCK_1	           :        40,
	APPLE_2	         	:        42,
	ORANGE_LITTLE         :        43,
	ORANGE_2	       :        46,
	BELL_2	           :        47,
	S77_LITTLE	         :        48,
	S77	          :        49,
	APPLE_3	         	:        50,
	MANGO_LITTLE          :        60,
	MANGO_2	         	:        68,
	STAR	             :        70,
	STAR_LITTLE             :        71,
	LUCK_2	             :        72,
	APPLE_4	             :        73,
	BELL_LITTLE	 	:		73,

	DASIXI                        :        74, //大四喜
	DASANYUAN                :        75, //大三元
	XIAOSANYUAN                :        76, //小三元
	RAND1		             :        78, //随机两炮
	RAND2                        :        80, //随机三炮
	RAND3                        :        90, //Bar
	RAND4                        :        100, //纵横四海
	ZONGHENGSIHAI              :        100,

}


/**
水果机单次中的概率
 */
var SGJP ShuiGuoPro = ShuiGuoPro{
	ORANGE_1 	        :        0,
	BELL_1	               	:        0,
	BAR_LITTLE              :        0,
	BAR	                :        0,
	APPLE_1	                :        0,
	APPLE_LITTLE           	:        0,
	MANGO_1	                :        0,
	WATERMELON            	:        0,
	WATERMELON_LITTLE     	:        0,
	LUCK_1	            	:        0,
	APPLE_2	                :        0,
	ORANGE_LITTLE         	:        0,
	ORANGE_2	       	:        0,
	BELL_2	             	:        0,
	S77_LITTLE	       	:        0,
	S77	             	:        0,
	APPLE_3	                :        0,
	MANGO_LITTLE            :        0,
	MANGO_2	                :        0,
	STAR	                :        0,
	STAR_LITTLE             :        0,
	LUCK_2	            	:        0,
	APPLE_4	              	:        0,
	BELL_LITTLE	        :        0,

	DASIXI                  :        0,
	DASANYUAN               :        0, //Bar
	XIAOSANYUAN             :        0, //纵横四海
	RAND1		      	:        0,
	RAND2 			:   	 0,
	RAND3                   :        3, //Bar
	RAND4                   :        0, //纵横四海
	ZONGHENGSIHAI           :        0,
}


/**
	随机赠送的时候的概率
 */
var SGJPRand ShuiGuoPro = ShuiGuoPro{
	ORANGE_1 	        :        5,
	BELL_1	               	:        15,
	BAR_LITTLE              :        20,
	BAR	                :        25,
	APPLE_1	                :        30,
	APPLE_LITTLE           	:        35,
	MANGO_1	                :        40,
	WATERMELON            	:        42,
	WATERMELON_LITTLE     	:        43,
	LUCK_1	            	:        0,
	APPLE_2	                :        47,
	ORANGE_LITTLE         	:        48,
	ORANGE_2	       	:        49,
	BELL_2	             	:        50,
	S77_LITTLE	       	:        60,
	S77	             	:        68,
	APPLE_3	                :        70,
	MANGO_LITTLE            :        71,
	MANGO_2	                :        72,
	STAR	                :        73,
	STAR_LITTLE             :        74,
	LUCK_2	            	:        0,
	APPLE_4	              	:        76,
	BELL_LITTLE	        :        100,
	DASIXI                  :        0,
	DASANYUAN               :        0,
	XIAOSANYUAN             :        0,
	RAND1		      	:        0,
	RAND2 			:	 0,
	RAND3                   :        0, //Bar
	RAND4                   :        0, //纵横四海
	ZONGHENGSIHAI           :        0,
}



/**

水果机押注对应的积分的数据
 */
type SGJValue struct {
	Apple            int32
	AppleLittle      int32
	Orange           int32
	OrangeLittle     int32
	Mango            int32
	MangoLittle      int32
	Bell             int32
	BellLittle       int32
	Watermelon       int32
	WatermelonLittle int32
	Star             int32
	StarLittle       int32
	S77              int32
	S77Little        int32
	Bar              int32
	BarLittle        int32
	Litter           int32
	Lucky            int32
}


/**
	得到一次的结果
	水果机器的结果有可能有很多种,这里需要什么策略来返回结果?
 */
func HandlerShuiguoji(m *bbproto.Shuiguoji) (*bbproto.ShuiguojiRes, error) {
	//1,检测参数并且根据押注的内容选择处理方式
	if m == nil {
		return nil,nil
	}

	//2,活的返回值
	result := &bbproto.ShuiguojiRes{}
	result.Result = make([]int32,24)
	p := SGJP
	err := BetResultWin(m,result, &p)
	if err != nil {
		log.E(err.Error())
		log.E("获取水果机结果的时候出错")
	}

	//跑到的结棍
	log.T("押注:%v",m)
	log.T("跑到的结果:%v",result.Result)

	//计算得分结果
	var scoresTotal int32 = 0

	if result.Result[INDEX_ORANGE_1] > 0 {
		scoresTotal += (m.GetScoresOrange() * SGJV.Orange)
	}

	if result.Result[INDEX_BELL_1] > 0 {
		scoresTotal += (m.GetScoresBell() * SGJV.Bell)
	}

	if result.Result[INDEX_BAR_LITTLE] > 0 {
		scoresTotal += (m.GetScoresBar() * SGJV.BarLittle)
	}

	if result.Result[INDEX_BAR] > 0 {
		scoresTotal += (m.GetScoresBar()  * SGJV.Bar)
	}

	if result.Result[INDEX_APPLE_1] > 0 {
		scoresTotal += (m.GetScoresApple()  * SGJV.Apple)

	}

	if result.Result[INDEX_APPLE_LITTLE] > 0 {
		scoresTotal += (m.GetScoresApple()  * SGJV.AppleLittle)
	}

	if result.Result[INDEX_MANGO_1] > 0 {
		scoresTotal += (m.GetScoresMango()  * SGJV.Mango)
	}

	if result.Result[INDEX_WATERMELON] > 0 {
		scoresTotal += (m.GetScoresWatermelon()  * SGJV.Watermelon)
	}

	if result.Result[INDEX_WATERMELON_LITTLE] > 0 {
		scoresTotal += (m.GetScoresWatermelon()  * SGJV.WatermelonLittle)
	}

	if result.Result[INDEX_APPLE_2] > 0 {
		scoresTotal += (m.GetScoresApple()  * SGJV.Apple)
		log.N("计算得到的总分是INDEX_APPLE_2%v", scoresTotal)
	}

	if result.Result[INDEX_ORANGE_LITTLE] > 0 {
		scoresTotal += (m.GetScoresOrange()  * SGJV.OrangeLittle)
	}

	if result.Result[INDEX_ORANGE_2] > 0 {
		scoresTotal += (m.GetScoresOrange()  * SGJV.Orange)
	}

	if result.Result[INDEX_BELL_2] > 0 {
		scoresTotal += (m.GetScoresBell()  * SGJV.Bell)
	}

	if result.Result[INDEX_77_LITTLE] > 0 {
		scoresTotal += (m.GetScores77()  * SGJV.S77Little)
	}
	if result.Result[INDEX_77] > 0 {
		scoresTotal += (m.GetScores77()  * SGJV.S77)
	}
	if result.Result[INDEX_APPLE_3] > 0 {
		scoresTotal += (m.GetScoresApple()  * SGJV.Apple)
		log.N("计算得到的总分是INDEX_APPLE_3%v", scoresTotal)
	}
	if result.Result[INDEX_MANGO_LITTLE] > 0 {
		scoresTotal += (m.GetScoresMango()  * SGJV.MangoLittle)
		log.N("计算得到的总分是INDEX_MANGO_LITTLE%v", scoresTotal)
	}
	if result.Result[INDEX_MANGO_2] > 0 {
		scoresTotal += (m.GetScoresMango()  * SGJV.Mango)
		log.N("计算得到的总分是INDEX_MANGO_2%v", scoresTotal)
	}
	if result.Result[INDEX_STAR] > 0 {
		scoresTotal += (m.GetScoresStar()  * SGJV.Star)
		log.N("计算得到的总分是INDEX_STAR%v", scoresTotal)
	}

	if result.Result[INDEX_STAR_LITTLE] > 0 {
		scoresTotal += (m.GetScoresStar()  * SGJV.StarLittle)
		log.N("计算得到的总分是INDEX_STAR_LITTLE%v", scoresTotal)
	}

	if result.Result[INDEX_APPLE_4] > 0 {
		scoresTotal += (m.GetScoresApple()  * SGJV.Apple)
		log.N("计算得到的总分是INDEX_APPLE_4%v", scoresTotal)
	}
	if result.Result[INDEX_BELL_LITTLE] > 0 {
		scoresTotal += (m.GetScoresBell()  * SGJV.BellLittle)
	}

	result.ScoresWin = &scoresTotal
	log.N("计算得到的总分是%v", result.GetScoresWin())

	//返回值
	return result, err

}


/**

水果机比大小的业务

 */
func HandlerShuiguojiHilomp(m *bbproto.Shuiguoji) (*bbproto.ShuiguojiHilomp, error) {
	return HilompResult(m.GetProtoHeader().GetUserId())
}



/**
得到跑到的结果
 */
func BetResultWin(m *bbproto.Shuiguoji, res *bbproto.ShuiguojiRes,p *ShuiGuoPro) (error) {
	r := utils.Randn(p.GetMax())
	log.T("水果机 随机数 %v", r)
	log.T("水果机 概率 %v", *p)

	if r < p.ORANGE_1 {
		res.Result[INDEX_ORANGE_1] = 1
		p.ORANGE_1 = 0

	} else if r < p.BELL_1 {
		res.Result[INDEX_BELL_1] = 1
		p.BELL_1 = 0

	} else if r < p.BAR_LITTLE {
		res.Result[INDEX_BAR_LITTLE] = 1
		p.BAR_LITTLE = 0

	} else if r < p.BAR {
		res.Result[INDEX_BAR] = 1
		p.BAR = 0

	} else if r < p.APPLE_1 {
		res.Result[INDEX_APPLE_1] = 1
		p.APPLE_1 = 0

	} else if r < p.WATERMELON {
		res.Result[INDEX_WATERMELON] = 1
		p.WATERMELON = 0

	} else if r < p.WATERMELON_LITTLE {
		res.Result[INDEX_WATERMELON_LITTLE] = 1
		p.WATERMELON_LITTLE = 0
	} else if r < p.LUCK_1 {
		res.Result[INDEX_LUCK_1] = 1
		p.LUCK_1 = 0
	} else if r < p.APPLE_2 {
		res.Result[INDEX_APPLE_2] = 1
		p.APPLE_2 = 0
	} else if r < p.ORANGE_LITTLE {
		res.Result[INDEX_ORANGE_LITTLE] = 1
		p.ORANGE_LITTLE = 0
	} else if r < p.ORANGE_2 {
		res.Result[INDEX_ORANGE_2] = 1
		p.ORANGE_2 = 0
	} else if r < p.BELL_2 {
		res.Result[INDEX_BELL_2] = 1
		p.BELL_2 = 0
	} else if r < p.S77_LITTLE {
		res.Result[INDEX_77_LITTLE] = 1
		p.S77_LITTLE = 0
	} else if r < p.S77 {
		res.Result[INDEX_77] = 1
		p.S77 = 0

	} else if r < p.APPLE_3 {
		res.Result[INDEX_APPLE_3] = 1
		p.APPLE_3 = 0

	} else if r < p.MANGO_LITTLE {
		res.Result[INDEX_MANGO_LITTLE] = 1
		p.MANGO_LITTLE = 0

	} else if r < p.MANGO_2 {
		res.Result[INDEX_MANGO_2] = 1
		p.MANGO_2 = 0

	} else if r < p.STAR {
		res.Result[INDEX_STAR] = 1
		p.STAR = 0

	} else if r < p.STAR_LITTLE {
		res.Result[INDEX_STAR_LITTLE] = 1
		p.STAR_LITTLE = 0

	} else if r < p.LUCK_2 {
		res.Result[INDEX_LUCK_2] = 1
		p.LUCK_2 = 0

	} else if r < p.APPLE_4 {
		res.Result[INDEX_APPLE_4] = 1
		p.APPLE_4 = 0

	} else if r < p.BELL_LITTLE {
		res.Result[INDEX_BELL_LITTLE] = 1
		p.BELL_LITTLE = 0

	} else if r < p.DASIXI {	//大四喜
		p.DASIXI = 0
		res.Result[INDEX_APPLE_1] = 1
		res.Result[INDEX_APPLE_2] = 2
		res.Result[INDEX_APPLE_3] = 3
		res.Result[INDEX_APPLE_4] = 4

	} else if r < p.DASANYUAN {	//大三元
		res.Result[INDEX_WATERMELON] = 1
		res.Result[INDEX_STAR]  = 2
		res.Result[INDEX_77]   = 3

	} else if r < p.XIAOSANYUAN {	//小三元
		res.Result[INDEX_ORANGE_1] = 1
		res.Result[INDEX_MANGO_1]  = 2
		res.Result[INDEX_BELL_1]   = 3

	} else if r < p.RAND1 {	//随机赠送一炮
		log.T("彩蛋,随机赠送一炮")
		pr1 := SGJPRand
		BetResultWin(m,res,&pr1)
		BetResultWin(m,res,&pr1)
	} else if r < p.RAND2 {
		log.T("彩蛋,随机赠送二炮")
		pr2 := SGJPRand
		BetResultWin(m,res,&pr2)
		BetResultWin(m,res,&pr2)
		BetResultWin(m,res,&pr2)
	} else if r < p.RAND3 {
		log.T("彩蛋,随机赠送三炮")
		pr3 := SGJPRand
		BetResultWin(m, res, &pr3)
		BetResultWin(m, res, &pr3)
		BetResultWin(m, res, &pr3)
		BetResultWin(m, res, &pr3)
	}else if r < p.ZONGHENGSIHAI {

	} else if r < p.WU {
		res.Result[INDEX_LUCK_2]   = 1
	}

	return  nil

}

/**
	比大小的结果
 */
func HilompResult(id uint32) (*bbproto.ShuiguojiHilomp, error) {
	result := &bbproto.ShuiguojiHilomp{}
	return result, nil
}
