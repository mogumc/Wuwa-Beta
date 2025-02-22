// 网关路由初始化与注册
// @author MoGuQAQ
// @version 1.0.0

package router

import (
	"wuwa-beta/api"
	"wuwa-beta/core"
	"wuwa-beta/global"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	if global.Log == nil {
		global.Log = core.InitLogger()
	}
	GinMode := "release"
	gin.SetMode(GinMode)
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())
	Register(router)
	return router
}

func Register(router *gin.Engine) {
	if global.Log == nil {
		global.Log = core.InitLogger()
	}

	router.GET("/launcher/useTaskId", api.UseTaskID)
}
