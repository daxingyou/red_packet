package routers

import (
	"gopkg.in/macaron.v1"
	"casino_super/handler/admin"
	"github.com/go-macaron/binding"
	"casino_super/model/logDao"
	"casino_super/handler/logHandler"
	"casino_super/modules"
)

//注册路由
func Regist(m *macaron.Macaron) {
	//日志
	m.Post("/log", binding.Json(logDao.ReqLog{}), logHandler.Post)
	m.Delete("/logs", binding.Json(logHandler.CodeValidate{}), logHandler.Delete)
	m.Get("/logs/:page", logHandler.Get)
	m.Get("/logs", logHandler.Get)

	//后台
	m.Group("/admin", func() {
		m.Get("/", admin.IndexHandler)
	})

	//首页
	m.Get("/", func(alert modules.Alert) {
		alert.Success("即将跳转至后台！", "/admin", 3)
	})
}
