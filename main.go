// 启动文件
// @author MoGuQAQ
// @version 1.0.0

package main

import (
	"wuwa-beta/core"
	"wuwa-beta/global"
	"wuwa-beta/router"
)

func main() {
	global.Log = core.InitLogger()
	global.Log.Infof("程序开始运行...")
	global.Log.Infof("初始化网关")
	router := router.InitRouter()
	global.Log.Infof("注册API路由")
	address := "127.0.0.1:1088"
	global.Log.Infof("网关启动成功 运行在: %s", address)
	router.Run(address)
}
