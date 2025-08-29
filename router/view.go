package router

import (
	"github.com/gin-gonic/gin"
	"goflylivechat/middleware"
	"goflylivechat/tmpl"
)

func InitViewRouter(engine *gin.Engine) {
	//engine.GET("/", tmpl.PageIndex)

	engine.GET("/login", tmpl.PageLogin)
	engine.GET("/pannel", tmpl.PagePannel)
	engine.GET("/chatIndex", tmpl.PageChat)
	engine.GET("/livechat", tmpl.PageChat)
	engine.GET("/main", middleware.JwtPageMiddleware, tmpl.PageMain)
	engine.GET("/chat_main", middleware.JwtPageMiddleware, middleware.DomainLimitMiddleware, tmpl.PageChatMain)
	engine.GET("/setting", middleware.DomainLimitMiddleware, tmpl.PageSetting)
	engine.GET("/setting_statistics", tmpl.PageSettingStatis)
	engine.GET("/setting_welcome", tmpl.PageSettingWelcome)
	engine.GET("/setting_deploy", tmpl.PageSettingDeploy)
	engine.GET("/setting_kefu_list", tmpl.PageKefuList)
	engine.GET("/setting_avator", tmpl.PageAvator)
	engine.GET("/setting_modifypass", tmpl.PageModifypass)
	engine.GET("/setting_ipblack", tmpl.PageIpblack)
	engine.GET("/setting_config", tmpl.PageConfig)
}
