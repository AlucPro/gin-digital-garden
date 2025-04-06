package main

import (
	"gin-digital-garden/core"
	"gin-digital-garden/global"
	"gin-digital-garden/initialize"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

const AppMode = "debug" // 运行环境，主要有三种：debug、test、release

func main() {
	gin.SetMode(AppMode)

	// 1.配置初始化
	global.GLOBAL_VP = core.InitializeViper()

	// 2.日志
	global.GLOBAL_LOG = core.InitializeZap()
	zap.ReplaceGlobals(global.GLOBAL_LOG)

	global.GLOBAL_LOG.Info("server run success on ", zap.String("zap_log", "zap_log"))

	//  3.数据库连接
	global.GLOBAL_DB = initialize.Gorm()

	// 4.其他初始化
	initialize.OtherInit()

	// 5.启动服务
	core.RunServer()
}
