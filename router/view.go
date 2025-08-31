package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitViewRouter(engine *gin.Engine) {
	//engine.GET("/", tmpl.PageIndex)
	engine.GET("/login", PageLogin)
	engine.GET("/pannel", PagePannel)
	engine.GET("/livechat", PageChat)
	engine.GET("/main", PageMain)
	engine.GET("/chat_main", PageChatMain)
	engine.GET("/setting", PageSetting)
}

// 登陆界面
func PageLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

// 面板界面
func PagePannel(c *gin.Context) {
	c.HTML(http.StatusOK, "pannel.html", nil)
}

// 后台界面
func PageMain(c *gin.Context) {
	c.HTML(http.StatusOK, "main.html", nil)
}

// 咨询界面
func PageChat(c *gin.Context) {
	kefuId := c.Query("user_id")
	refer := c.Query("refer")
	if refer == "" {
		refer = c.Request.Referer()
	}
	if refer == "" {
		refer = "​​Direct Link"
	}
	c.HTML(http.StatusOK, "chat_page.html", gin.H{
		"KEFU_ID": kefuId,
		"Refer":   refer,
	})
}

// 客服界面
func PageChatMain(c *gin.Context) {
	c.HTML(http.StatusOK, "chat_main.html", nil)
}

// 设置界面
func PageSetting(c *gin.Context) {
	c.HTML(http.StatusOK, "setting.html", nil)
}
