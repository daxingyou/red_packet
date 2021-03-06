package config

import (
	"casino_redpack/modules"
	"casino_common/common/model"
	"github.com/go-macaron/binding"
	"casino_common/utils/db"
	"casino_common/common/consts/tableName"
	"gopkg.in/mgo.v2/bson"
	"casino_redpack/model/configModel"
)

//商品配置列表
func GoodsListHandler(ctx *modules.Context) {
	list := []*model.T_Goods_Row{}
	cate_id := ctx.QueryInt("cate_id")
	query := bson.M{}
	if cate_id > 0 {
		query["category"] = cate_id
	}
	db.C(tableName.DBT_T_GOODS_INFO).FindAll(query, &list)
	ctx.Data["cate_id"] = cate_id
	ctx.Data["list"] = list
	ctx.HTML(200, "admin/config/goods/index")
}



//编辑--一条数据
func GoodsEditPost(ctx *modules.Context) {
	id :=ctx.Query("id")
	obj_id := bson.ObjectIdHex(id)
	goods_info := configModel.GoodsInfoOne(obj_id)
	ctx.Data["info"] = goods_info
	ctx.HTML(200,"admin/config/goods/edit")
}

//编辑--提交
func GoodsEditUpdate(ctx *modules.Context) {
	id :=ctx.Query("id")
	obj_id := bson.ObjectIdHex(id)
	Goodsid :=ctx.QueryFloat64("Goodsid")
	Name := ctx.Query("Name")
	Category := ctx.QueryFloat64("Category")
	PriceType := ctx.QueryFloat64("PriceType")
	Price := ctx.QueryFloat64("Price")
	Amount := ctx.QueryFloat64("Amount")
	GoodsType := ctx.QueryFloat64("GoodsType")
	Discount := ctx.Query("Discount")
	Image := ctx.Query("Image")
	Isshow := ctx.QueryBool("IsShow")
	Sort := ctx.QueryFloat64("Sort")
	err :=configModel.GoodsEditUpdate(obj_id,Goodsid,Name,Category,PriceType,Price,GoodsType,Amount,Discount,Image,Isshow,Sort)
	if err != nil {
		ctx.Ajax(-1, "编辑失败！", nil)
		return
	}
	ctx.Ajax(1, "编辑成功！", nil)
}

//新增
func GoodsInsertPost(ctx *modules.Context, form model.T_Goods_Row, errs binding.Errors) {
	if errs.Len() > 0 {
		ctx.Ajax(-1, "表单参数错误！", nil)
		return
	}
	err := form.Insert()
	if err != nil {
		ctx.Ajax(-1, "新增配置失败！", nil)
		return
	}
	ctx.Ajax(1, "新增配置成功！", nil)
}

//删除
func GoodsRemoveHnadler(ctx *modules.Context) {
	id := ctx.Query("id")
	obj_id := bson.ObjectIdHex(id)
	err := db.C(tableName.DBT_T_GOODS_INFO).Remove(bson.M{"_id": obj_id})
	if err != nil {
		ctx.Ajax(-1, "删除失败！", nil)
		return
	}
	ctx.Ajax(1, "删除成功！", nil)
}
